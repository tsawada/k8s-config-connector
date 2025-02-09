// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/osconfig/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOSConfigPatchDeployments implements OSConfigPatchDeploymentInterface
type FakeOSConfigPatchDeployments struct {
	Fake *FakeOsconfigV1alpha1
	ns   string
}

var osconfigpatchdeploymentsResource = schema.GroupVersionResource{Group: "osconfig.cnrm.cloud.google.com", Version: "v1alpha1", Resource: "osconfigpatchdeployments"}

var osconfigpatchdeploymentsKind = schema.GroupVersionKind{Group: "osconfig.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "OSConfigPatchDeployment"}

// Get takes name of the oSConfigPatchDeployment, and returns the corresponding oSConfigPatchDeployment object, and an error if there is any.
func (c *FakeOSConfigPatchDeployments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.OSConfigPatchDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(osconfigpatchdeploymentsResource, c.ns, name), &v1alpha1.OSConfigPatchDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OSConfigPatchDeployment), err
}

// List takes label and field selectors, and returns the list of OSConfigPatchDeployments that match those selectors.
func (c *FakeOSConfigPatchDeployments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.OSConfigPatchDeploymentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(osconfigpatchdeploymentsResource, osconfigpatchdeploymentsKind, c.ns, opts), &v1alpha1.OSConfigPatchDeploymentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OSConfigPatchDeploymentList{ListMeta: obj.(*v1alpha1.OSConfigPatchDeploymentList).ListMeta}
	for _, item := range obj.(*v1alpha1.OSConfigPatchDeploymentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested oSConfigPatchDeployments.
func (c *FakeOSConfigPatchDeployments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(osconfigpatchdeploymentsResource, c.ns, opts))

}

// Create takes the representation of a oSConfigPatchDeployment and creates it.  Returns the server's representation of the oSConfigPatchDeployment, and an error, if there is any.
func (c *FakeOSConfigPatchDeployments) Create(ctx context.Context, oSConfigPatchDeployment *v1alpha1.OSConfigPatchDeployment, opts v1.CreateOptions) (result *v1alpha1.OSConfigPatchDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(osconfigpatchdeploymentsResource, c.ns, oSConfigPatchDeployment), &v1alpha1.OSConfigPatchDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OSConfigPatchDeployment), err
}

// Update takes the representation of a oSConfigPatchDeployment and updates it. Returns the server's representation of the oSConfigPatchDeployment, and an error, if there is any.
func (c *FakeOSConfigPatchDeployments) Update(ctx context.Context, oSConfigPatchDeployment *v1alpha1.OSConfigPatchDeployment, opts v1.UpdateOptions) (result *v1alpha1.OSConfigPatchDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(osconfigpatchdeploymentsResource, c.ns, oSConfigPatchDeployment), &v1alpha1.OSConfigPatchDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OSConfigPatchDeployment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOSConfigPatchDeployments) UpdateStatus(ctx context.Context, oSConfigPatchDeployment *v1alpha1.OSConfigPatchDeployment, opts v1.UpdateOptions) (*v1alpha1.OSConfigPatchDeployment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(osconfigpatchdeploymentsResource, "status", c.ns, oSConfigPatchDeployment), &v1alpha1.OSConfigPatchDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OSConfigPatchDeployment), err
}

// Delete takes name of the oSConfigPatchDeployment and deletes it. Returns an error if one occurs.
func (c *FakeOSConfigPatchDeployments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(osconfigpatchdeploymentsResource, c.ns, name, opts), &v1alpha1.OSConfigPatchDeployment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOSConfigPatchDeployments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(osconfigpatchdeploymentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.OSConfigPatchDeploymentList{})
	return err
}

// Patch applies the patch and returns the patched oSConfigPatchDeployment.
func (c *FakeOSConfigPatchDeployments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.OSConfigPatchDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(osconfigpatchdeploymentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.OSConfigPatchDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OSConfigPatchDeployment), err
}
