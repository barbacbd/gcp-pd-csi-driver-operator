// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	operatorv1 "github.com/openshift/api/operator/v1"
)

// AWSNetworkLoadBalancerParametersApplyConfiguration represents a declarative configuration of the AWSNetworkLoadBalancerParameters type for use
// with apply.
type AWSNetworkLoadBalancerParametersApplyConfiguration struct {
	Subnets        *AWSSubnetsApplyConfiguration `json:"subnets,omitempty"`
	EIPAllocations []operatorv1.EIPAllocation    `json:"eipAllocations,omitempty"`
}

// AWSNetworkLoadBalancerParametersApplyConfiguration constructs a declarative configuration of the AWSNetworkLoadBalancerParameters type for use with
// apply.
func AWSNetworkLoadBalancerParameters() *AWSNetworkLoadBalancerParametersApplyConfiguration {
	return &AWSNetworkLoadBalancerParametersApplyConfiguration{}
}

// WithSubnets sets the Subnets field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Subnets field is set to the value of the last call.
func (b *AWSNetworkLoadBalancerParametersApplyConfiguration) WithSubnets(value *AWSSubnetsApplyConfiguration) *AWSNetworkLoadBalancerParametersApplyConfiguration {
	b.Subnets = value
	return b
}

// WithEIPAllocations adds the given value to the EIPAllocations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the EIPAllocations field.
func (b *AWSNetworkLoadBalancerParametersApplyConfiguration) WithEIPAllocations(values ...operatorv1.EIPAllocation) *AWSNetworkLoadBalancerParametersApplyConfiguration {
	for i := range values {
		b.EIPAllocations = append(b.EIPAllocations, values[i])
	}
	return b
}
