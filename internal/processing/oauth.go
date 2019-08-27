package processing

import (
	"context"
	"encoding/json"
	"fmt"
	gatewayv2alpha1 "github.com/kyma-incubator/api-gateway/api/v2alpha1"
	rulev1alpha1 "github.com/ory/oathkeeper-maester/api/v1alpha1"
	"github.com/pkg/errors"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	k8sMeta "k8s.io/apimachinery/pkg/apis/meta/v1"
	internalTypes "github.com/kyma-incubator/api-gateway/internal/types/ory"
	"k8s.io/apimachinery/pkg/runtime"
	"knative.dev/pkg/apis/istio/common/v1alpha1"
	networkingv1alpha3 "knative.dev/pkg/apis/istio/v1alpha3"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type oauth struct {
	client.Client
}

func (o *oauth) Process(ctx context.Context, api *gatewayv2alpha1.Gate) error {
	fmt.Println("Processing API")

	//oldVS, err := o.getVirtualService(ctx, api)
	//if err != nil {
	//	return err
	//}
	//
	//if oldVS != nil {
	//	newVS := o.prepareVirtualService(api, oldVS)
	//	return o.update(ctx, newVS)
	//} else {
	//	vs := generateVirtualService(api)
	//	return o.create(ctx, vs)
	//}

	oauthConfig, err := generateOauthConfig(api)
	if err!=nil {
		return err
	}

	requiredScopesJSON, err := generateRequiredScopesJSON(oauthConfig)
	if err!=nil {
		return err
	}

	ar := generateAccessRule(api, oauthConfig, requiredScopesJSON)
	return o.create(ctx, ar)
}

func (o *oauth) getVirtualService(ctx context.Context, api *gatewayv2alpha1.Gate) (*networkingv1alpha3.VirtualService, error) {
	virtualServiceName := fmt.Sprintf("%s-%s", api.ObjectMeta.Name, *api.Spec.Service.Name)
	namespacedName := client.ObjectKey{Namespace: api.GetNamespace(), Name: virtualServiceName}
	var vs networkingv1alpha3.VirtualService

	err := o.Client.Get(ctx, namespacedName, &vs)
	if err != nil {
		if apierrs.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	return &vs, nil
}

func (o *oauth) create(ctx context.Context, obj runtime.Object) error {
	return o.Client.Create(ctx, obj)
}

func (o *oauth) prepareVirtualService(api *gatewayv2alpha1.Gate, vs *networkingv1alpha3.VirtualService) *networkingv1alpha3.VirtualService {
	virtualServiceName := fmt.Sprintf("%s-%s", api.ObjectMeta.Name, *api.Spec.Service.Name)
	controller := true

	ownerRef := &k8sMeta.OwnerReference{
		Name:       api.ObjectMeta.Name,
		APIVersion: api.TypeMeta.APIVersion,
		Kind:       api.TypeMeta.Kind,
		UID:        api.ObjectMeta.UID,
		Controller: &controller,
	}

	vs.ObjectMeta.OwnerReferences = []k8sMeta.OwnerReference{*ownerRef}
	vs.ObjectMeta.Name = virtualServiceName
	vs.ObjectMeta.Namespace = api.ObjectMeta.Namespace

	match := &networkingv1alpha3.HTTPMatchRequest{
		URI: &v1alpha1.StringMatch{
			Regex: "/.*",
		},
	}
	route := &networkingv1alpha3.HTTPRouteDestination{
		Destination: networkingv1alpha3.Destination{
			Host: fmt.Sprintf("%s.%s.svc.cluster.local", *api.Spec.Service.Name, api.ObjectMeta.Namespace),
			Port: networkingv1alpha3.PortSelector{
				Number: uint32(*api.Spec.Service.Port),
			},
		},
	}

	spec := &networkingv1alpha3.VirtualServiceSpec{
		Hosts:    []string{*api.Spec.Service.Host},
		Gateways: []string{*api.Spec.Gateway},
		HTTP: []networkingv1alpha3.HTTPRoute{
			{
				Match: []networkingv1alpha3.HTTPMatchRequest{*match},
				Route: []networkingv1alpha3.HTTPRouteDestination{*route},
			},
		},
	}

	vs.Spec = *spec

	return vs

}

func (o *oauth) update(ctx context.Context, obj runtime.Object) error {
	return o.Client.Update(ctx, obj)
}

func generateObjectMeta(api *gatewayv2alpha1.Gate) k8sMeta.ObjectMeta {
	objName := fmt.Sprintf("%s-%s", api.ObjectMeta.Name, *api.Spec.Service.Name)
	controller := true

	ownerRef := &k8sMeta.OwnerReference{
		Name:       api.ObjectMeta.Name,
		APIVersion: api.TypeMeta.APIVersion,
		Kind:       api.TypeMeta.Kind,
		UID:        api.ObjectMeta.UID,
		Controller: &controller,
	}

	objectMeta := k8sMeta.ObjectMeta{
		Name:            objName,
		Namespace:       api.ObjectMeta.Namespace,
		OwnerReferences: []k8sMeta.OwnerReference{*ownerRef},
	}

	return objectMeta
}

func generateVirtualService(api *gatewayv2alpha1.Gate) *networkingv1alpha3.VirtualService {
	objectMeta := generateObjectMeta(api)

	match := &networkingv1alpha3.HTTPMatchRequest{
		URI: &v1alpha1.StringMatch{
			Regex: "/.*",
		},
	}
	route := &networkingv1alpha3.HTTPRouteDestination{
		Destination: networkingv1alpha3.Destination{
			Host: fmt.Sprintf("%s.%s.svc.cluster.local", *api.Spec.Service.Name, api.ObjectMeta.Namespace),
			Port: networkingv1alpha3.PortSelector{
				Number: uint32(*api.Spec.Service.Port),
			},
		},
	}

	spec := &networkingv1alpha3.VirtualServiceSpec{
		Hosts:    []string{*api.Spec.Service.Host},
		Gateways: []string{*api.Spec.Gateway},
		HTTP: []networkingv1alpha3.HTTPRoute{
			{
				Match: []networkingv1alpha3.HTTPMatchRequest{*match},
				Route: []networkingv1alpha3.HTTPRouteDestination{*route},
			},
		},
	}

	vs := &networkingv1alpha3.VirtualService{
		ObjectMeta: objectMeta,
		Spec:       *spec,
	}

	return vs
}

func generateRequiredScopesJSON(config *gatewayv2alpha1.OauthModeConfig) ([]byte, error){
	requiredScopes := &internalTypes.OauthIntrospectionConfig{config.Paths[0].Scopes}
	return json.Marshal(requiredScopes)
}

func generateOauthConfig(api *gatewayv2alpha1.Gate)(*gatewayv2alpha1.OauthModeConfig, error){
	apiConfig := api.Spec.Auth.Config
	var oauthConfig gatewayv2alpha1.OauthModeConfig

	err := json.Unmarshal(apiConfig.Raw, &oauthConfig)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &oauthConfig, nil
}

func generateAccessRule(api *gatewayv2alpha1.Gate, config *gatewayv2alpha1.OauthModeConfig, requiredScopes []byte) *rulev1alpha1.Rule {
	objectMeta := generateObjectMeta(api)

	rawConfig := &runtime.RawExtension{
		Raw: requiredScopes,
	}

	spec := &rulev1alpha1.RuleSpec{
		Upstream:&rulev1alpha1.Upstream{
			URL: fmt.Sprintf("http://%s.%s.svc.cluster.local:%d", *api.Spec.Service.Name, api.ObjectMeta.Namespace, int(*api.Spec.Service.Port)),
		},
		Match:&rulev1alpha1.Match{
			Methods: config.Paths[0].Methods,
			URL: fmt.Sprintf("<http|https>://%s/<.*>", *api.Spec.Service.Host),
		},
		Authorizer:&rulev1alpha1.Authorizer{
			Handler: &rulev1alpha1.Handler{
				Name: "allow",
			},
		},
		Authenticators:[]*rulev1alpha1.Authenticator{
			{
				Handler: &rulev1alpha1.Handler{
					Name:"oauth2_introspection",
					Config: rawConfig,
				},
			},
		},
	}

	rule := &rulev1alpha1.Rule{
		ObjectMeta: objectMeta,
		Spec:       *spec,
	}

	return rule
}
