// +build !ignore_autogenerated_openshift

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	api_v1 "k8s.io/kubernetes/pkg/api/v1"
	conversion "k8s.io/kubernetes/pkg/conversion"
	runtime "k8s.io/kubernetes/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_Route, InType: reflect.TypeOf(&Route{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_RouteIngress, InType: reflect.TypeOf(&RouteIngress{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_RouteIngressCondition, InType: reflect.TypeOf(&RouteIngressCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_RouteList, InType: reflect.TypeOf(&RouteList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_RoutePort, InType: reflect.TypeOf(&RoutePort{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_RouteSpec, InType: reflect.TypeOf(&RouteSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_RouteStatus, InType: reflect.TypeOf(&RouteStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_RouteTargetReference, InType: reflect.TypeOf(&RouteTargetReference{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_RouterShard, InType: reflect.TypeOf(&RouterShard{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1_TLSConfig, InType: reflect.TypeOf(&TLSConfig{})},
	)
}

func DeepCopy_v1_Route(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Route)
		out := out.(*Route)
		out.TypeMeta = in.TypeMeta
		if err := api_v1.DeepCopy_v1_ObjectMeta(&in.ObjectMeta, &out.ObjectMeta, c); err != nil {
			return err
		}
		if err := DeepCopy_v1_RouteSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_v1_RouteStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1_RouteIngress(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RouteIngress)
		out := out.(*RouteIngress)
		out.Host = in.Host
		out.RouterName = in.RouterName
		if in.Conditions != nil {
			in, out := &in.Conditions, &out.Conditions
			*out = make([]RouteIngressCondition, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_RouteIngressCondition(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Conditions = nil
		}
		out.WildcardPolicy = in.WildcardPolicy
		return nil
	}
}

func DeepCopy_v1_RouteIngressCondition(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RouteIngressCondition)
		out := out.(*RouteIngressCondition)
		out.Type = in.Type
		out.Status = in.Status
		out.Reason = in.Reason
		out.Message = in.Message
		if in.LastTransitionTime != nil {
			in, out := &in.LastTransitionTime, &out.LastTransitionTime
			*out = new(unversioned.Time)
			**out = (*in).DeepCopy()
		} else {
			out.LastTransitionTime = nil
		}
		return nil
	}
}

func DeepCopy_v1_RouteList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RouteList)
		out := out.(*RouteList)
		out.TypeMeta = in.TypeMeta
		out.ListMeta = in.ListMeta
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Route, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_Route(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Items = nil
		}
		return nil
	}
}

func DeepCopy_v1_RoutePort(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RoutePort)
		out := out.(*RoutePort)
		out.TargetPort = in.TargetPort
		return nil
	}
}

func DeepCopy_v1_RouteSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RouteSpec)
		out := out.(*RouteSpec)
		out.Host = in.Host
		out.Path = in.Path
		if err := DeepCopy_v1_RouteTargetReference(&in.To, &out.To, c); err != nil {
			return err
		}
		if in.AlternateBackends != nil {
			in, out := &in.AlternateBackends, &out.AlternateBackends
			*out = make([]RouteTargetReference, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_RouteTargetReference(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.AlternateBackends = nil
		}
		if in.Port != nil {
			in, out := &in.Port, &out.Port
			*out = new(RoutePort)
			**out = **in
		} else {
			out.Port = nil
		}
		if in.TLS != nil {
			in, out := &in.TLS, &out.TLS
			*out = new(TLSConfig)
			**out = **in
		} else {
			out.TLS = nil
		}
		out.WildcardPolicy = in.WildcardPolicy
		return nil
	}
}

func DeepCopy_v1_RouteStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RouteStatus)
		out := out.(*RouteStatus)
		if in.Ingress != nil {
			in, out := &in.Ingress, &out.Ingress
			*out = make([]RouteIngress, len(*in))
			for i := range *in {
				if err := DeepCopy_v1_RouteIngress(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		} else {
			out.Ingress = nil
		}
		return nil
	}
}

func DeepCopy_v1_RouteTargetReference(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RouteTargetReference)
		out := out.(*RouteTargetReference)
		out.Kind = in.Kind
		out.Name = in.Name
		if in.Weight != nil {
			in, out := &in.Weight, &out.Weight
			*out = new(int32)
			**out = **in
		} else {
			out.Weight = nil
		}
		return nil
	}
}

func DeepCopy_v1_RouterShard(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*RouterShard)
		out := out.(*RouterShard)
		out.ShardName = in.ShardName
		out.DNSSuffix = in.DNSSuffix
		return nil
	}
}

func DeepCopy_v1_TLSConfig(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*TLSConfig)
		out := out.(*TLSConfig)
		out.Termination = in.Termination
		out.Certificate = in.Certificate
		out.Key = in.Key
		out.CACertificate = in.CACertificate
		out.DestinationCACertificate = in.DestinationCACertificate
		out.InsecureEdgeTerminationPolicy = in.InsecureEdgeTerminationPolicy
		return nil
	}
}
