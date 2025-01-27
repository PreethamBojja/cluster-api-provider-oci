//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) 2022, Oracle and/or its affiliates.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta2

import (
	apiv1beta2 "github.com/oracle/cluster-api-provider-oci/api/v1beta2"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/cluster-api/api/v1beta1"
	"sigs.k8s.io/cluster-api/errors"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceConfiguration) DeepCopyInto(out *InstanceConfiguration) {
	*out = *in
	if in.InstanceConfigurationId != nil {
		in, out := &in.InstanceConfigurationId, &out.InstanceConfigurationId
		*out = new(string)
		**out = **in
	}
	if in.Shape != nil {
		in, out := &in.Shape, &out.Shape
		*out = new(string)
		**out = **in
	}
	if in.ShapeConfig != nil {
		in, out := &in.ShapeConfig, &out.ShapeConfig
		*out = new(ShapeConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.InstanceVnicConfiguration != nil {
		in, out := &in.InstanceVnicConfiguration, &out.InstanceVnicConfiguration
		*out = new(apiv1beta2.NetworkDetails)
		(*in).DeepCopyInto(*out)
	}
	if in.PlatformConfig != nil {
		in, out := &in.PlatformConfig, &out.PlatformConfig
		*out = new(apiv1beta2.PlatformConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.AgentConfig != nil {
		in, out := &in.AgentConfig, &out.AgentConfig
		*out = new(apiv1beta2.LaunchInstanceAgentConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.PreemptibleInstanceConfig != nil {
		in, out := &in.PreemptibleInstanceConfig, &out.PreemptibleInstanceConfig
		*out = new(apiv1beta2.PreemptibleInstanceConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.AvailabilityConfig != nil {
		in, out := &in.AvailabilityConfig, &out.AvailabilityConfig
		*out = new(apiv1beta2.LaunchInstanceAvailabilityConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.DedicatedVmHostId != nil {
		in, out := &in.DedicatedVmHostId, &out.DedicatedVmHostId
		*out = new(string)
		**out = **in
	}
	if in.LaunchOptions != nil {
		in, out := &in.LaunchOptions, &out.LaunchOptions
		*out = new(apiv1beta2.LaunchOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.InstanceOptions != nil {
		in, out := &in.InstanceOptions, &out.InstanceOptions
		*out = new(apiv1beta2.InstanceOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.IsPvEncryptionInTransitEnabled != nil {
		in, out := &in.IsPvEncryptionInTransitEnabled, &out.IsPvEncryptionInTransitEnabled
		*out = new(bool)
		**out = **in
	}
	if in.InstanceSourceViaImageDetails != nil {
		in, out := &in.InstanceSourceViaImageDetails, &out.InstanceSourceViaImageDetails
		*out = new(InstanceSourceViaImageConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.CapacityReservationId != nil {
		in, out := &in.CapacityReservationId, &out.CapacityReservationId
		*out = new(string)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceConfiguration.
func (in *InstanceConfiguration) DeepCopy() *InstanceConfiguration {
	if in == nil {
		return nil
	}
	out := new(InstanceConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSourceViaImageConfig) DeepCopyInto(out *InstanceSourceViaImageConfig) {
	*out = *in
	if in.ImageId != nil {
		in, out := &in.ImageId, &out.ImageId
		*out = new(string)
		**out = **in
	}
	if in.KmsKeyId != nil {
		in, out := &in.KmsKeyId, &out.KmsKeyId
		*out = new(string)
		**out = **in
	}
	if in.BootVolumeSizeInGBs != nil {
		in, out := &in.BootVolumeSizeInGBs, &out.BootVolumeSizeInGBs
		*out = new(int64)
		**out = **in
	}
	if in.BootVolumeVpusPerGB != nil {
		in, out := &in.BootVolumeVpusPerGB, &out.BootVolumeVpusPerGB
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSourceViaImageConfig.
func (in *InstanceSourceViaImageConfig) DeepCopy() *InstanceSourceViaImageConfig {
	if in == nil {
		return nil
	}
	out := new(InstanceSourceViaImageConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceVnicConfiguration) DeepCopyInto(out *InstanceVnicConfiguration) {
	*out = *in
	if in.NSGId != nil {
		in, out := &in.NSGId, &out.NSGId
		*out = new(string)
		**out = **in
	}
	if in.SkipSourceDestCheck != nil {
		in, out := &in.SkipSourceDestCheck, &out.SkipSourceDestCheck
		*out = new(bool)
		**out = **in
	}
	if in.NsgNames != nil {
		in, out := &in.NsgNames, &out.NsgNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.HostnameLabel != nil {
		in, out := &in.HostnameLabel, &out.HostnameLabel
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.AssignPrivateDnsRecord != nil {
		in, out := &in.AssignPrivateDnsRecord, &out.AssignPrivateDnsRecord
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceVnicConfiguration.
func (in *InstanceVnicConfiguration) DeepCopy() *InstanceVnicConfiguration {
	if in == nil {
		return nil
	}
	out := new(InstanceVnicConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyValue) DeepCopyInto(out *KeyValue) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyValue.
func (in *KeyValue) DeepCopy() *KeyValue {
	if in == nil {
		return nil
	}
	out := new(KeyValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchDetails) DeepCopyInto(out *LaunchDetails) {
	*out = *in
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchDetails.
func (in *LaunchDetails) DeepCopy() *LaunchDetails {
	if in == nil {
		return nil
	}
	out := new(LaunchDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeEvictionNodePoolSettings) DeepCopyInto(out *NodeEvictionNodePoolSettings) {
	*out = *in
	if in.EvictionGraceDuration != nil {
		in, out := &in.EvictionGraceDuration, &out.EvictionGraceDuration
		*out = new(string)
		**out = **in
	}
	if in.IsForceDeleteAfterGraceDuration != nil {
		in, out := &in.IsForceDeleteAfterGraceDuration, &out.IsForceDeleteAfterGraceDuration
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeEvictionNodePoolSettings.
func (in *NodeEvictionNodePoolSettings) DeepCopy() *NodeEvictionNodePoolSettings {
	if in == nil {
		return nil
	}
	out := new(NodeEvictionNodePoolSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolCyclingDetails) DeepCopyInto(out *NodePoolCyclingDetails) {
	*out = *in
	if in.IsNodeCyclingEnabled != nil {
		in, out := &in.IsNodeCyclingEnabled, &out.IsNodeCyclingEnabled
		*out = new(bool)
		**out = **in
	}
	if in.MaximumSurge != nil {
		in, out := &in.MaximumSurge, &out.MaximumSurge
		*out = new(string)
		**out = **in
	}
	if in.MaximumUnavailable != nil {
		in, out := &in.MaximumUnavailable, &out.MaximumUnavailable
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolCyclingDetails.
func (in *NodePoolCyclingDetails) DeepCopy() *NodePoolCyclingDetails {
	if in == nil {
		return nil
	}
	out := new(NodePoolCyclingDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolNodeConfig) DeepCopyInto(out *NodePoolNodeConfig) {
	*out = *in
	if in.IsPvEncryptionInTransitEnabled != nil {
		in, out := &in.IsPvEncryptionInTransitEnabled, &out.IsPvEncryptionInTransitEnabled
		*out = new(bool)
		**out = **in
	}
	if in.KmsKeyId != nil {
		in, out := &in.KmsKeyId, &out.KmsKeyId
		*out = new(string)
		**out = **in
	}
	if in.PlacementConfigs != nil {
		in, out := &in.PlacementConfigs, &out.PlacementConfigs
		*out = make([]PlacementConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NsgNames != nil {
		in, out := &in.NsgNames, &out.NsgNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NodePoolPodNetworkOptionDetails != nil {
		in, out := &in.NodePoolPodNetworkOptionDetails, &out.NodePoolPodNetworkOptionDetails
		*out = new(NodePoolPodNetworkOptionDetails)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolNodeConfig.
func (in *NodePoolNodeConfig) DeepCopy() *NodePoolNodeConfig {
	if in == nil {
		return nil
	}
	out := new(NodePoolNodeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolPodNetworkOptionDetails) DeepCopyInto(out *NodePoolPodNetworkOptionDetails) {
	*out = *in
	in.VcnIpNativePodNetworkOptions.DeepCopyInto(&out.VcnIpNativePodNetworkOptions)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolPodNetworkOptionDetails.
func (in *NodePoolPodNetworkOptionDetails) DeepCopy() *NodePoolPodNetworkOptionDetails {
	if in == nil {
		return nil
	}
	out := new(NodePoolPodNetworkOptionDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeShapeConfig) DeepCopyInto(out *NodeShapeConfig) {
	*out = *in
	if in.MemoryInGBs != nil {
		in, out := &in.MemoryInGBs, &out.MemoryInGBs
		*out = new(string)
		**out = **in
	}
	if in.Ocpus != nil {
		in, out := &in.Ocpus, &out.Ocpus
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeShapeConfig.
func (in *NodeShapeConfig) DeepCopy() *NodeShapeConfig {
	if in == nil {
		return nil
	}
	out := new(NodeShapeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeSourceViaImage) DeepCopyInto(out *NodeSourceViaImage) {
	*out = *in
	if in.BootVolumeSizeInGBs != nil {
		in, out := &in.BootVolumeSizeInGBs, &out.BootVolumeSizeInGBs
		*out = new(int64)
		**out = **in
	}
	if in.ImageId != nil {
		in, out := &in.ImageId, &out.ImageId
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeSourceViaImage.
func (in *NodeSourceViaImage) DeepCopy() *NodeSourceViaImage {
	if in == nil {
		return nil
	}
	out := new(NodeSourceViaImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCIMachinePool) DeepCopyInto(out *OCIMachinePool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIMachinePool.
func (in *OCIMachinePool) DeepCopy() *OCIMachinePool {
	if in == nil {
		return nil
	}
	out := new(OCIMachinePool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OCIMachinePool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCIMachinePoolList) DeepCopyInto(out *OCIMachinePoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OCIMachinePool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIMachinePoolList.
func (in *OCIMachinePoolList) DeepCopy() *OCIMachinePoolList {
	if in == nil {
		return nil
	}
	out := new(OCIMachinePoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OCIMachinePoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCIMachinePoolSpec) DeepCopyInto(out *OCIMachinePoolSpec) {
	*out = *in
	if in.ProviderID != nil {
		in, out := &in.ProviderID, &out.ProviderID
		*out = new(string)
		**out = **in
	}
	if in.OCID != nil {
		in, out := &in.OCID, &out.OCID
		*out = new(string)
		**out = **in
	}
	if in.PlacementDetails != nil {
		in, out := &in.PlacementDetails, &out.PlacementDetails
		*out = make([]PlacementDetails, len(*in))
		copy(*out, *in)
	}
	in.InstanceConfiguration.DeepCopyInto(&out.InstanceConfiguration)
	if in.ProviderIDList != nil {
		in, out := &in.ProviderIDList, &out.ProviderIDList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIMachinePoolSpec.
func (in *OCIMachinePoolSpec) DeepCopy() *OCIMachinePoolSpec {
	if in == nil {
		return nil
	}
	out := new(OCIMachinePoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCIMachinePoolStatus) DeepCopyInto(out *OCIMachinePoolStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(v1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FailureReason != nil {
		in, out := &in.FailureReason, &out.FailureReason
		*out = new(errors.MachineStatusError)
		**out = **in
	}
	if in.FailureMessage != nil {
		in, out := &in.FailureMessage, &out.FailureMessage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIMachinePoolStatus.
func (in *OCIMachinePoolStatus) DeepCopy() *OCIMachinePoolStatus {
	if in == nil {
		return nil
	}
	out := new(OCIMachinePoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCIManagedMachinePool) DeepCopyInto(out *OCIManagedMachinePool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIManagedMachinePool.
func (in *OCIManagedMachinePool) DeepCopy() *OCIManagedMachinePool {
	if in == nil {
		return nil
	}
	out := new(OCIManagedMachinePool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OCIManagedMachinePool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCIManagedMachinePoolList) DeepCopyInto(out *OCIManagedMachinePoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OCIManagedMachinePool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIManagedMachinePoolList.
func (in *OCIManagedMachinePoolList) DeepCopy() *OCIManagedMachinePoolList {
	if in == nil {
		return nil
	}
	out := new(OCIManagedMachinePoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OCIManagedMachinePoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCIManagedMachinePoolSpec) DeepCopyInto(out *OCIManagedMachinePoolSpec) {
	*out = *in
	if in.ProviderID != nil {
		in, out := &in.ProviderID, &out.ProviderID
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.NodePoolNodeConfig != nil {
		in, out := &in.NodePoolNodeConfig, &out.NodePoolNodeConfig
		*out = new(NodePoolNodeConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeEvictionNodePoolSettings != nil {
		in, out := &in.NodeEvictionNodePoolSettings, &out.NodeEvictionNodePoolSettings
		*out = new(NodeEvictionNodePoolSettings)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeShapeConfig != nil {
		in, out := &in.NodeShapeConfig, &out.NodeShapeConfig
		*out = new(NodeShapeConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSourceViaImage != nil {
		in, out := &in.NodeSourceViaImage, &out.NodeSourceViaImage
		*out = new(NodeSourceViaImage)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeMetadata != nil {
		in, out := &in.NodeMetadata, &out.NodeMetadata
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.InitialNodeLabels != nil {
		in, out := &in.InitialNodeLabels, &out.InitialNodeLabels
		*out = make([]KeyValue, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodePoolCyclingDetails != nil {
		in, out := &in.NodePoolCyclingDetails, &out.NodePoolCyclingDetails
		*out = new(NodePoolCyclingDetails)
		(*in).DeepCopyInto(*out)
	}
	if in.ProviderIDList != nil {
		in, out := &in.ProviderIDList, &out.ProviderIDList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIManagedMachinePoolSpec.
func (in *OCIManagedMachinePoolSpec) DeepCopy() *OCIManagedMachinePoolSpec {
	if in == nil {
		return nil
	}
	out := new(OCIManagedMachinePoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCIManagedMachinePoolStatus) DeepCopyInto(out *OCIManagedMachinePoolStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(v1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FailureReason != nil {
		in, out := &in.FailureReason, &out.FailureReason
		*out = new(errors.MachineStatusError)
		**out = **in
	}
	if in.FailureMessages != nil {
		in, out := &in.FailureMessages, &out.FailureMessages
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIManagedMachinePoolStatus.
func (in *OCIManagedMachinePoolStatus) DeepCopy() *OCIManagedMachinePoolStatus {
	if in == nil {
		return nil
	}
	out := new(OCIManagedMachinePoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCIManagedMachinePoolTemplate) DeepCopyInto(out *OCIManagedMachinePoolTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIManagedMachinePoolTemplate.
func (in *OCIManagedMachinePoolTemplate) DeepCopy() *OCIManagedMachinePoolTemplate {
	if in == nil {
		return nil
	}
	out := new(OCIManagedMachinePoolTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OCIManagedMachinePoolTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCIManagedMachinePoolTemplateList) DeepCopyInto(out *OCIManagedMachinePoolTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OCIManagedMachinePoolTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIManagedMachinePoolTemplateList.
func (in *OCIManagedMachinePoolTemplateList) DeepCopy() *OCIManagedMachinePoolTemplateList {
	if in == nil {
		return nil
	}
	out := new(OCIManagedMachinePoolTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OCIManagedMachinePoolTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCIManagedMachinePoolTemplateResource) DeepCopyInto(out *OCIManagedMachinePoolTemplateResource) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIManagedMachinePoolTemplateResource.
func (in *OCIManagedMachinePoolTemplateResource) DeepCopy() *OCIManagedMachinePoolTemplateResource {
	if in == nil {
		return nil
	}
	out := new(OCIManagedMachinePoolTemplateResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCIManagedMachinePoolTemplateSpec) DeepCopyInto(out *OCIManagedMachinePoolTemplateSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIManagedMachinePoolTemplateSpec.
func (in *OCIManagedMachinePoolTemplateSpec) DeepCopy() *OCIManagedMachinePoolTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(OCIManagedMachinePoolTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCIVirtualMachinePool) DeepCopyInto(out *OCIVirtualMachinePool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIVirtualMachinePool.
func (in *OCIVirtualMachinePool) DeepCopy() *OCIVirtualMachinePool {
	if in == nil {
		return nil
	}
	out := new(OCIVirtualMachinePool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OCIVirtualMachinePool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCIVirtualMachinePoolList) DeepCopyInto(out *OCIVirtualMachinePoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OCIVirtualMachinePool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIVirtualMachinePoolList.
func (in *OCIVirtualMachinePoolList) DeepCopy() *OCIVirtualMachinePoolList {
	if in == nil {
		return nil
	}
	out := new(OCIVirtualMachinePoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OCIVirtualMachinePoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCIVirtualMachinePoolSpec) DeepCopyInto(out *OCIVirtualMachinePoolSpec) {
	*out = *in
	if in.ProviderID != nil {
		in, out := &in.ProviderID, &out.ProviderID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.PlacementConfigs != nil {
		in, out := &in.PlacementConfigs, &out.PlacementConfigs
		*out = make([]VirtualNodepoolPlacementConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NsgNames != nil {
		in, out := &in.NsgNames, &out.NsgNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.PodConfiguration.DeepCopyInto(&out.PodConfiguration)
	if in.Taints != nil {
		in, out := &in.Taints, &out.Taints
		*out = make([]Taint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InitialVirtualNodeLabels != nil {
		in, out := &in.InitialVirtualNodeLabels, &out.InitialVirtualNodeLabels
		*out = make([]KeyValue, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ProviderIDList != nil {
		in, out := &in.ProviderIDList, &out.ProviderIDList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIVirtualMachinePoolSpec.
func (in *OCIVirtualMachinePoolSpec) DeepCopy() *OCIVirtualMachinePoolSpec {
	if in == nil {
		return nil
	}
	out := new(OCIVirtualMachinePoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCIVirtualMachinePoolStatus) DeepCopyInto(out *OCIVirtualMachinePoolStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(v1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FailureReason != nil {
		in, out := &in.FailureReason, &out.FailureReason
		*out = new(errors.MachinePoolStatusFailure)
		**out = **in
	}
	if in.FailureMessages != nil {
		in, out := &in.FailureMessages, &out.FailureMessages
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIVirtualMachinePoolStatus.
func (in *OCIVirtualMachinePoolStatus) DeepCopy() *OCIVirtualMachinePoolStatus {
	if in == nil {
		return nil
	}
	out := new(OCIVirtualMachinePoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlacementConfig) DeepCopyInto(out *PlacementConfig) {
	*out = *in
	if in.AvailabilityDomain != nil {
		in, out := &in.AvailabilityDomain, &out.AvailabilityDomain
		*out = new(string)
		**out = **in
	}
	if in.CapacityReservationId != nil {
		in, out := &in.CapacityReservationId, &out.CapacityReservationId
		*out = new(string)
		**out = **in
	}
	if in.FaultDomains != nil {
		in, out := &in.FaultDomains, &out.FaultDomains
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SubnetName != nil {
		in, out := &in.SubnetName, &out.SubnetName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlacementConfig.
func (in *PlacementConfig) DeepCopy() *PlacementConfig {
	if in == nil {
		return nil
	}
	out := new(PlacementConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlacementDetails) DeepCopyInto(out *PlacementDetails) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlacementDetails.
func (in *PlacementDetails) DeepCopy() *PlacementDetails {
	if in == nil {
		return nil
	}
	out := new(PlacementDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodConfig) DeepCopyInto(out *PodConfig) {
	*out = *in
	if in.NsgNames != nil {
		in, out := &in.NsgNames, &out.NsgNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Shape != nil {
		in, out := &in.Shape, &out.Shape
		*out = new(string)
		**out = **in
	}
	if in.SubnetName != nil {
		in, out := &in.SubnetName, &out.SubnetName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodConfig.
func (in *PodConfig) DeepCopy() *PodConfig {
	if in == nil {
		return nil
	}
	out := new(PodConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShapeConfig) DeepCopyInto(out *ShapeConfig) {
	*out = *in
	if in.Ocpus != nil {
		in, out := &in.Ocpus, &out.Ocpus
		*out = new(string)
		**out = **in
	}
	if in.MemoryInGBs != nil {
		in, out := &in.MemoryInGBs, &out.MemoryInGBs
		*out = new(string)
		**out = **in
	}
	if in.Nvmes != nil {
		in, out := &in.Nvmes, &out.Nvmes
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShapeConfig.
func (in *ShapeConfig) DeepCopy() *ShapeConfig {
	if in == nil {
		return nil
	}
	out := new(ShapeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Taint) DeepCopyInto(out *Taint) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	if in.Effect != nil {
		in, out := &in.Effect, &out.Effect
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Taint.
func (in *Taint) DeepCopy() *Taint {
	if in == nil {
		return nil
	}
	out := new(Taint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VcnIpNativePodNetworkOptions) DeepCopyInto(out *VcnIpNativePodNetworkOptions) {
	*out = *in
	if in.MaxPodsPerNode != nil {
		in, out := &in.MaxPodsPerNode, &out.MaxPodsPerNode
		*out = new(int)
		**out = **in
	}
	if in.NSGNames != nil {
		in, out := &in.NSGNames, &out.NSGNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SubnetNames != nil {
		in, out := &in.SubnetNames, &out.SubnetNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VcnIpNativePodNetworkOptions.
func (in *VcnIpNativePodNetworkOptions) DeepCopy() *VcnIpNativePodNetworkOptions {
	if in == nil {
		return nil
	}
	out := new(VcnIpNativePodNetworkOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualNodepoolPlacementConfig) DeepCopyInto(out *VirtualNodepoolPlacementConfig) {
	*out = *in
	if in.AvailabilityDomain != nil {
		in, out := &in.AvailabilityDomain, &out.AvailabilityDomain
		*out = new(string)
		**out = **in
	}
	if in.FaultDomains != nil {
		in, out := &in.FaultDomains, &out.FaultDomains
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SubnetName != nil {
		in, out := &in.SubnetName, &out.SubnetName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualNodepoolPlacementConfig.
func (in *VirtualNodepoolPlacementConfig) DeepCopy() *VirtualNodepoolPlacementConfig {
	if in == nil {
		return nil
	}
	out := new(VirtualNodepoolPlacementConfig)
	in.DeepCopyInto(out)
	return out
}
