package builders

import (
	k8sMeta "k8s.io/apimachinery/pkg/apis/meta/v1"
	networkingv1alpha1 "knative.dev/pkg/apis/istio/common/v1alpha1"
	networkingv1alpha3 "knative.dev/pkg/apis/istio/v1alpha3"
)

// VirtualService creates builder for knative.dev/pkg/apis/istio/v1alpha3/VirtualService type
func VirtualService(name string) *virtualService {
	return &virtualService{
		name: name,
	}
}

type virtualService struct {
	name      string
	namespace string
	owner     *ownerReference
	spec      *virtualServiceSpec
}

func (b *virtualService) Namespace(ns string) *virtualService {
	b.namespace = ns
	return b
}

func (b *virtualService) Owner(o *ownerReference) *virtualService {
	b.owner = o
	return b
}

func (b *virtualService) Spec(s *virtualServiceSpec) *virtualService {
	b.spec = s
	return b
}

func (b *virtualService) Get() *networkingv1alpha3.VirtualService {
	objectMeta := k8sMeta.ObjectMeta{
		Name:      b.name,
		Namespace: b.namespace,
	}

	if b.owner != nil {
		objectMeta.OwnerReferences = []k8sMeta.OwnerReference{*b.owner.Get()}
	}

	vs := &networkingv1alpha3.VirtualService{
		ObjectMeta: objectMeta,
		Spec:       *b.spec.Get(),
	}

	return vs
}

// VirtualServiceSpec creates builder for knative.dev/pkg/apis/istio/v1alpha3/VirtualServiceSpec type
func VirtualServiceSpec() *virtualServiceSpec {
	return &virtualServiceSpec{}
}

type virtualServiceSpec struct {
	hosts     []string
	gateways  []string
	matchReq  *matchRequest
	routeDest *routeDestination
}

func (b *virtualServiceSpec) Host(host string) *virtualServiceSpec {
	b.hosts = []string{host}
	return b
}

func (b *virtualServiceSpec) Gateway(gw string) *virtualServiceSpec {
	b.gateways = append(b.gateways, gw)
	return b
}

func (b *virtualServiceSpec) HTTP(mr *matchRequest, rd *routeDestination) *virtualServiceSpec {
	b.matchReq = mr
	b.routeDest = rd
	return b
}

func (b *virtualServiceSpec) Get() *networkingv1alpha3.VirtualServiceSpec {

	var httpMatch []networkingv1alpha3.HTTPMatchRequest
	var routeDest []networkingv1alpha3.HTTPRouteDestination

	if b.matchReq != nil {
		httpMatch = append(httpMatch, *b.matchReq.Get())
	}

	if b.routeDest != nil {
		routeDest = append(routeDest, *b.routeDest.Get())
	}

	spec := &networkingv1alpha3.VirtualServiceSpec{
		Hosts:    b.hosts,
		Gateways: b.gateways,
		HTTP: []networkingv1alpha3.HTTPRoute{
			{
				Match: httpMatch,
				Route: routeDest,
			},
		},
	}

	return spec
}

// MatchRequest creates builder for knative.dev/pkg/apis/istio/v1alpha3/HTTPMatchRequest type
func MatchRequest() *matchRequest {
	return &matchRequest{}
}

type matchRequest struct {
	data *networkingv1alpha3.HTTPMatchRequest
}

func (mr *matchRequest) Get() *networkingv1alpha3.HTTPMatchRequest {
	return mr.data
}

func (mr *matchRequest) URI() *stringMatch {
	mr.data = &networkingv1alpha3.HTTPMatchRequest{}
	mr.data.URI = &networkingv1alpha1.StringMatch{}
	return &stringMatch{mr.data.URI, func() *matchRequest { return mr }}
}

type stringMatch struct {
	data   *networkingv1alpha1.StringMatch
	parent func() *matchRequest
}

func (st *stringMatch) Regex(value string) *matchRequest {
	st.data.Regex = value
	return st.parent()
}

// RouteDestination creates builder for knative.dev/pkg/apis/istio/v1alpha3/HTTPRouteDestination type
func RouteDestination() *routeDestination {
	return &routeDestination{&networkingv1alpha3.HTTPRouteDestination{}}
}

type routeDestination struct {
	data *networkingv1alpha3.HTTPRouteDestination
}

func (rd *routeDestination) Host(value string) *routeDestination {
	rd.data.Destination.Host = value
	return rd
}
func (rd *routeDestination) Port(value uint32) *routeDestination {
	rd.data.Destination.Port.Number = value
	return rd
}
func (rd *routeDestination) Get() *networkingv1alpha3.HTTPRouteDestination {
	return rd.data
}
