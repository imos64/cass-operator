// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1beta1

import (
	json "encoding/json"

	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraDatacenter) DeepCopyInto(out *CassandraDatacenter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraDatacenter.
func (in *CassandraDatacenter) DeepCopy() *CassandraDatacenter {
	if in == nil {
		return nil
	}
	out := new(CassandraDatacenter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CassandraDatacenter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraDatacenterList) DeepCopyInto(out *CassandraDatacenterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CassandraDatacenter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraDatacenterList.
func (in *CassandraDatacenterList) DeepCopy() *CassandraDatacenterList {
	if in == nil {
		return nil
	}
	out := new(CassandraDatacenterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CassandraDatacenterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraDatacenterSpec) DeepCopyInto(out *CassandraDatacenterSpec) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
	in.ManagementApiAuth.DeepCopyInto(&out.ManagementApiAuth)
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Racks != nil {
		in, out := &in.Racks, &out.Racks
		*out = make([]Rack, len(*in))
		copy(*out, *in)
	}
	in.StorageConfig.DeepCopyInto(&out.StorageConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraDatacenterSpec.
func (in *CassandraDatacenterSpec) DeepCopy() *CassandraDatacenterSpec {
	if in == nil {
		return nil
	}
	out := new(CassandraDatacenterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraDatacenterStatus) DeepCopyInto(out *CassandraDatacenterStatus) {
	*out = *in
	in.SuperUserUpserted.DeepCopyInto(&out.SuperUserUpserted)
	in.LastServerNodeStarted.DeepCopyInto(&out.LastServerNodeStarted)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraDatacenterStatus.
func (in *CassandraDatacenterStatus) DeepCopy() *CassandraDatacenterStatus {
	if in == nil {
		return nil
	}
	out := new(CassandraDatacenterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementApiAuthConfig) DeepCopyInto(out *ManagementApiAuthConfig) {
	*out = *in
	if in.Insecure != nil {
		in, out := &in.Insecure, &out.Insecure
		*out = new(ManagementApiAuthInsecureConfig)
		**out = **in
	}
	if in.Manual != nil {
		in, out := &in.Manual, &out.Manual
		*out = new(ManagementApiAuthManualConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementApiAuthConfig.
func (in *ManagementApiAuthConfig) DeepCopy() *ManagementApiAuthConfig {
	if in == nil {
		return nil
	}
	out := new(ManagementApiAuthConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementApiAuthInsecureConfig) DeepCopyInto(out *ManagementApiAuthInsecureConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementApiAuthInsecureConfig.
func (in *ManagementApiAuthInsecureConfig) DeepCopy() *ManagementApiAuthInsecureConfig {
	if in == nil {
		return nil
	}
	out := new(ManagementApiAuthInsecureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementApiAuthManualConfig) DeepCopyInto(out *ManagementApiAuthManualConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementApiAuthManualConfig.
func (in *ManagementApiAuthManualConfig) DeepCopy() *ManagementApiAuthManualConfig {
	if in == nil {
		return nil
	}
	out := new(ManagementApiAuthManualConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rack) DeepCopyInto(out *Rack) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rack.
func (in *Rack) DeepCopy() *Rack {
	if in == nil {
		return nil
	}
	out := new(Rack)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageConfig) DeepCopyInto(out *StorageConfig) {
	*out = *in
	if in.CassandraDataVolumeClaimSpec != nil {
		in, out := &in.CassandraDataVolumeClaimSpec, &out.CassandraDataVolumeClaimSpec
		*out = new(v1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageConfig.
func (in *StorageConfig) DeepCopy() *StorageConfig {
	if in == nil {
		return nil
	}
	out := new(StorageConfig)
	in.DeepCopyInto(out)
	return out
}