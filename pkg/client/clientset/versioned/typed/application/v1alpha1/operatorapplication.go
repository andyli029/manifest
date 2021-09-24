/*
Copyright 2021.

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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "manifest/api/application/v1alpha1"
	scheme "manifest/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// OperatorApplicationsGetter has a method to return a OperatorApplicationInterface.
// A group's client should implement this interface.
type OperatorApplicationsGetter interface {
	OperatorApplications() OperatorApplicationInterface
}

// OperatorApplicationInterface has methods to work with OperatorApplication resources.
type OperatorApplicationInterface interface {
	Create(ctx context.Context, operatorApplication *v1alpha1.OperatorApplication, opts v1.CreateOptions) (*v1alpha1.OperatorApplication, error)
	Update(ctx context.Context, operatorApplication *v1alpha1.OperatorApplication, opts v1.UpdateOptions) (*v1alpha1.OperatorApplication, error)
	UpdateStatus(ctx context.Context, operatorApplication *v1alpha1.OperatorApplication, opts v1.UpdateOptions) (*v1alpha1.OperatorApplication, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.OperatorApplication, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.OperatorApplicationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.OperatorApplication, err error)
	OperatorApplicationExpansion
}

// operatorApplications implements OperatorApplicationInterface
type operatorApplications struct {
	client rest.Interface
}

// newOperatorApplications returns a OperatorApplications
func newOperatorApplications(c *ApplicationV1alpha1Client) *operatorApplications {
	return &operatorApplications{
		client: c.RESTClient(),
	}
}

// Get takes name of the operatorApplication, and returns the corresponding operatorApplication object, and an error if there is any.
func (c *operatorApplications) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.OperatorApplication, err error) {
	result = &v1alpha1.OperatorApplication{}
	err = c.client.Get().
		Resource("operatorapplications").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of OperatorApplications that match those selectors.
func (c *operatorApplications) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.OperatorApplicationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.OperatorApplicationList{}
	err = c.client.Get().
		Resource("operatorapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested operatorApplications.
func (c *operatorApplications) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("operatorapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a operatorApplication and creates it.  Returns the server's representation of the operatorApplication, and an error, if there is any.
func (c *operatorApplications) Create(ctx context.Context, operatorApplication *v1alpha1.OperatorApplication, opts v1.CreateOptions) (result *v1alpha1.OperatorApplication, err error) {
	result = &v1alpha1.OperatorApplication{}
	err = c.client.Post().
		Resource("operatorapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(operatorApplication).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a operatorApplication and updates it. Returns the server's representation of the operatorApplication, and an error, if there is any.
func (c *operatorApplications) Update(ctx context.Context, operatorApplication *v1alpha1.OperatorApplication, opts v1.UpdateOptions) (result *v1alpha1.OperatorApplication, err error) {
	result = &v1alpha1.OperatorApplication{}
	err = c.client.Put().
		Resource("operatorapplications").
		Name(operatorApplication.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(operatorApplication).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *operatorApplications) UpdateStatus(ctx context.Context, operatorApplication *v1alpha1.OperatorApplication, opts v1.UpdateOptions) (result *v1alpha1.OperatorApplication, err error) {
	result = &v1alpha1.OperatorApplication{}
	err = c.client.Put().
		Resource("operatorapplications").
		Name(operatorApplication.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(operatorApplication).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the operatorApplication and deletes it. Returns an error if one occurs.
func (c *operatorApplications) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("operatorapplications").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *operatorApplications) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("operatorapplications").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched operatorApplication.
func (c *operatorApplications) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.OperatorApplication, err error) {
	result = &v1alpha1.OperatorApplication{}
	err = c.client.Patch(pt).
		Resource("operatorapplications").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
