// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ComputeAddressSpec struct {
	/* Immutable. The static external IP address represented by this resource. Only
	IPv4 is supported. An address may only be specified for INTERNAL
	address types. The IP address must be inside the specified subnetwork,
	if any. Set by the API if undefined. */
	// +optional
	Address *string `json:"address,omitempty"`

	/* Immutable. The type of address to reserve.
	Note: if you set this argument's value as 'INTERNAL' you need to leave the 'network_tier' argument unset in that resource block. Default value: "EXTERNAL" Possible values: ["INTERNAL", "EXTERNAL"]. */
	// +optional
	AddressType *string `json:"addressType,omitempty"`

	/* Immutable. An optional description of this resource. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Immutable. The IP Version that will be used by this address. The default value is 'IPV4'. Possible values: ["IPV4", "IPV6"]. This field can only be specified for a global address. */
	// +optional
	IpVersion *string `json:"ipVersion,omitempty"`

	/* Location represents the geographical location of the ComputeAddress. Specify a region name or "global" for global resources. Reference: GCP definition of regions/zones (https://cloud.google.com/compute/docs/regions-zones/) */
	Location string `json:"location"`

	/* The network in which to reserve the address. If global, the address
	must be within the RFC1918 IP space. The network cannot be deleted
	if there are any reserved IP ranges referring to it. This field can
	only be used with INTERNAL type with the VPC_PEERING and
	IPSEC_INTERCONNECT purposes. */
	// +optional
	NetworkRef *v1alpha1.ResourceRef `json:"networkRef,omitempty"`

	/* Immutable. The networking tier used for configuring this address. If this field is not
	specified, it is assumed to be PREMIUM.
	This argument should not be used when configuring Internal addresses, because [network tier cannot be set for internal traffic; it's always Premium](https://cloud.google.com/network-tiers/docs/overview). Possible values: ["PREMIUM", "STANDARD"]. */
	// +optional
	NetworkTier *string `json:"networkTier,omitempty"`

	/* Immutable. The prefix length if the resource represents an IP range. */
	// +optional
	PrefixLength *int `json:"prefixLength,omitempty"`

	/* Immutable. The purpose of this resource, which can be one of the following values.

	* GCE_ENDPOINT for addresses that are used by VM instances, alias IP
	ranges, load balancers, and similar resources.

	* SHARED_LOADBALANCER_VIP for an address that can be used by multiple
	internal load balancers.

	* VPC_PEERING for addresses that are reserved for VPC peer networks.

	* IPSEC_INTERCONNECT for addresses created from a private IP range that
	are reserved for a VLAN attachment in an HA VPN over Cloud Interconnect
	configuration. These addresses are regional resources.

	* PRIVATE_SERVICE_CONNECT for a private network address that is used to
	configure Private Service Connect. Only global internal addresses can use
	this purpose.


	This should only be set when using an Internal address. */
	// +optional
	Purpose *string `json:"purpose,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* The subnetwork in which to reserve the address. If an IP address is
	specified, it must be within the subnetwork's IP range.  This field
	can only be used with INTERNAL type with GCE_ENDPOINT/DNS_RESOLVER
	purposes. */
	// +optional
	SubnetworkRef *v1alpha1.ResourceRef `json:"subnetworkRef,omitempty"`
}

type ComputeAddressStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeAddress's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Creation timestamp in RFC3339 text format. */
	// +optional
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	/* The fingerprint used for optimistic locking of this resource.  Used
	internally during updates. */
	// +optional
	LabelFingerprint *string `json:"labelFingerprint,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	// +optional
	SelfLink *string `json:"selfLink,omitempty"`

	/* The URLs of the resources that are using this address. */
	// +optional
	Users []string `json:"users,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeAddress is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeAddress struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeAddressSpec   `json:"spec,omitempty"`
	Status ComputeAddressStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeAddressList contains a list of ComputeAddress
type ComputeAddressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeAddress `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeAddress{}, &ComputeAddressList{})
}
