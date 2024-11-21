// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// IngressControllerHTTPHeaderApplyConfiguration represents a declarative configuration of the IngressControllerHTTPHeader type for use
// with apply.
type IngressControllerHTTPHeaderApplyConfiguration struct {
	Name   *string                                                   `json:"name,omitempty"`
	Action *IngressControllerHTTPHeaderActionUnionApplyConfiguration `json:"action,omitempty"`
}

// IngressControllerHTTPHeaderApplyConfiguration constructs a declarative configuration of the IngressControllerHTTPHeader type for use with
// apply.
func IngressControllerHTTPHeader() *IngressControllerHTTPHeaderApplyConfiguration {
	return &IngressControllerHTTPHeaderApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *IngressControllerHTTPHeaderApplyConfiguration) WithName(value string) *IngressControllerHTTPHeaderApplyConfiguration {
	b.Name = &value
	return b
}

// WithAction sets the Action field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Action field is set to the value of the last call.
func (b *IngressControllerHTTPHeaderApplyConfiguration) WithAction(value *IngressControllerHTTPHeaderActionUnionApplyConfiguration) *IngressControllerHTTPHeaderApplyConfiguration {
	b.Action = value
	return b
}
