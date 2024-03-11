/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "github.com/jacob-delgado/demo/pkg/apis/foo/v1"
	foov1 "github.com/jacob-delgado/demo/pkg/client/applyconfiguration/foo/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHelloTypes implements HelloTypeInterface
type FakeHelloTypes struct {
	Fake *FakeFooV1
	ns   string
}

var hellotypesResource = v1.SchemeGroupVersion.WithResource("hellotypes")

var hellotypesKind = v1.SchemeGroupVersion.WithKind("HelloType")

// Get takes name of the helloType, and returns the corresponding helloType object, and an error if there is any.
func (c *FakeHelloTypes) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.HelloType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(hellotypesResource, c.ns, name), &v1.HelloType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.HelloType), err
}

// List takes label and field selectors, and returns the list of HelloTypes that match those selectors.
func (c *FakeHelloTypes) List(ctx context.Context, opts metav1.ListOptions) (result *v1.HelloTypeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(hellotypesResource, hellotypesKind, c.ns, opts), &v1.HelloTypeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.HelloTypeList{ListMeta: obj.(*v1.HelloTypeList).ListMeta}
	for _, item := range obj.(*v1.HelloTypeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested helloTypes.
func (c *FakeHelloTypes) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(hellotypesResource, c.ns, opts))

}

// Create takes the representation of a helloType and creates it.  Returns the server's representation of the helloType, and an error, if there is any.
func (c *FakeHelloTypes) Create(ctx context.Context, helloType *v1.HelloType, opts metav1.CreateOptions) (result *v1.HelloType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(hellotypesResource, c.ns, helloType), &v1.HelloType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.HelloType), err
}

// Update takes the representation of a helloType and updates it. Returns the server's representation of the helloType, and an error, if there is any.
func (c *FakeHelloTypes) Update(ctx context.Context, helloType *v1.HelloType, opts metav1.UpdateOptions) (result *v1.HelloType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(hellotypesResource, c.ns, helloType), &v1.HelloType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.HelloType), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHelloTypes) UpdateStatus(ctx context.Context, helloType *v1.HelloType, opts metav1.UpdateOptions) (*v1.HelloType, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(hellotypesResource, "status", c.ns, helloType), &v1.HelloType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.HelloType), err
}

// Delete takes name of the helloType and deletes it. Returns an error if one occurs.
func (c *FakeHelloTypes) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(hellotypesResource, c.ns, name, opts), &v1.HelloType{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHelloTypes) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(hellotypesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.HelloTypeList{})
	return err
}

// Patch applies the patch and returns the patched helloType.
func (c *FakeHelloTypes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.HelloType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(hellotypesResource, c.ns, name, pt, data, subresources...), &v1.HelloType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.HelloType), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied helloType.
func (c *FakeHelloTypes) Apply(ctx context.Context, helloType *foov1.HelloTypeApplyConfiguration, opts metav1.ApplyOptions) (result *v1.HelloType, err error) {
	if helloType == nil {
		return nil, fmt.Errorf("helloType provided to Apply must not be nil")
	}
	data, err := json.Marshal(helloType)
	if err != nil {
		return nil, err
	}
	name := helloType.Name
	if name == nil {
		return nil, fmt.Errorf("helloType.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(hellotypesResource, c.ns, *name, types.ApplyPatchType, data), &v1.HelloType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.HelloType), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeHelloTypes) ApplyStatus(ctx context.Context, helloType *foov1.HelloTypeApplyConfiguration, opts metav1.ApplyOptions) (result *v1.HelloType, err error) {
	if helloType == nil {
		return nil, fmt.Errorf("helloType provided to Apply must not be nil")
	}
	data, err := json.Marshal(helloType)
	if err != nil {
		return nil, err
	}
	name := helloType.Name
	if name == nil {
		return nil, fmt.Errorf("helloType.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(hellotypesResource, c.ns, *name, types.ApplyPatchType, data, "status"), &v1.HelloType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.HelloType), err
}
