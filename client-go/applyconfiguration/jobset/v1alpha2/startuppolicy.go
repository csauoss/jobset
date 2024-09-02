/*
Copyright 2023 The Kubernetes Authors.
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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha2

import (
	v1alpha2 "sigs.k8s.io/jobset/api/jobset/v1alpha2"
)

// StartupPolicyApplyConfiguration represents a declarative configuration of the StartupPolicy type for use
// with apply.
type StartupPolicyApplyConfiguration struct {
	StartupPolicyOrder *v1alpha2.StartupPolicyOptions `json:"startupPolicyOrder,omitempty"`
}

// StartupPolicyApplyConfiguration constructs a declarative configuration of the StartupPolicy type for use with
// apply.
func StartupPolicy() *StartupPolicyApplyConfiguration {
	return &StartupPolicyApplyConfiguration{}
}

// WithStartupPolicyOrder sets the StartupPolicyOrder field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StartupPolicyOrder field is set to the value of the last call.
func (b *StartupPolicyApplyConfiguration) WithStartupPolicyOrder(value v1alpha2.StartupPolicyOptions) *StartupPolicyApplyConfiguration {
	b.StartupPolicyOrder = &value
	return b
}
