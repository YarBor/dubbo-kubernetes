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

package v1alpha1

import (
	"context"
	v1alpha1 "github.com/apache/dubbo-kubernetes/pkg/core/gen/apis/dubbo.apache.org/v1alpha1"
	scheme "github.com/apache/dubbo-kubernetes/pkg/core/gen/generated/clientset/versioned/scheme"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ServiceNameMappingsGetter has a method to return a ServiceNameMappingInterface.
// A group's client should implement this interface.
type ServiceNameMappingsGetter interface {
	ServiceNameMappings(namespace string) ServiceNameMappingInterface
}

// ServiceNameMappingInterface has methods to work with ServiceNameMapping resources.
type ServiceNameMappingInterface interface {
	Create(ctx context.Context, serviceNameMapping *v1alpha1.ServiceNameMapping, opts v1.CreateOptions) (*v1alpha1.ServiceNameMapping, error)
	Update(ctx context.Context, serviceNameMapping *v1alpha1.ServiceNameMapping, opts v1.UpdateOptions) (*v1alpha1.ServiceNameMapping, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ServiceNameMapping, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ServiceNameMappingList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServiceNameMapping, err error)
	ServiceNameMappingExpansion
}

// serviceNameMappings implements ServiceNameMappingInterface
type serviceNameMappings struct {
	client rest.Interface
	ns     string
}

// newServiceNameMappings returns a ServiceNameMappings
func newServiceNameMappings(c *DubboV1alpha1Client, namespace string) *serviceNameMappings {
	return &serviceNameMappings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the serviceNameMapping, and returns the corresponding serviceNameMapping object, and an error if there is any.
func (c *serviceNameMappings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ServiceNameMapping, err error) {
	result = &v1alpha1.ServiceNameMapping{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("servicenamemappings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ServiceNameMappings that match those selectors.
func (c *serviceNameMappings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ServiceNameMappingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ServiceNameMappingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("servicenamemappings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested serviceNameMappings.
func (c *serviceNameMappings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("servicenamemappings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a serviceNameMapping and creates it.  Returns the server's representation of the serviceNameMapping, and an error, if there is any.
func (c *serviceNameMappings) Create(ctx context.Context, serviceNameMapping *v1alpha1.ServiceNameMapping, opts v1.CreateOptions) (result *v1alpha1.ServiceNameMapping, err error) {
	result = &v1alpha1.ServiceNameMapping{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("servicenamemappings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(serviceNameMapping).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a serviceNameMapping and updates it. Returns the server's representation of the serviceNameMapping, and an error, if there is any.
func (c *serviceNameMappings) Update(ctx context.Context, serviceNameMapping *v1alpha1.ServiceNameMapping, opts v1.UpdateOptions) (result *v1alpha1.ServiceNameMapping, err error) {
	result = &v1alpha1.ServiceNameMapping{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("servicenamemappings").
		Name(serviceNameMapping.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(serviceNameMapping).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the serviceNameMapping and deletes it. Returns an error if one occurs.
func (c *serviceNameMappings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servicenamemappings").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *serviceNameMappings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servicenamemappings").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched serviceNameMapping.
func (c *serviceNameMappings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServiceNameMapping, err error) {
	result = &v1alpha1.ServiceNameMapping{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("servicenamemappings").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
