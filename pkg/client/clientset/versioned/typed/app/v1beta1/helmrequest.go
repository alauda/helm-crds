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

package v1beta1

import (
	"time"

	v1beta1 "github.com/alauda/helm-crds/pkg/apis/app/v1beta1"
	scheme "github.com/alauda/helm-crds/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// HelmRequestsGetter has a method to return a HelmRequestInterface.
// A group's client should implement this interface.
type HelmRequestsGetter interface {
	HelmRequests(namespace string) HelmRequestInterface
}

// HelmRequestInterface has methods to work with HelmRequest resources.
type HelmRequestInterface interface {
	Create(*v1beta1.HelmRequest) (*v1beta1.HelmRequest, error)
	Update(*v1beta1.HelmRequest) (*v1beta1.HelmRequest, error)
	UpdateStatus(*v1beta1.HelmRequest) (*v1beta1.HelmRequest, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.HelmRequest, error)
	List(opts v1.ListOptions) (*v1beta1.HelmRequestList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.HelmRequest, err error)
	HelmRequestExpansion
}

// helmRequests implements HelmRequestInterface
type helmRequests struct {
	client rest.Interface
	ns     string
}

// newHelmRequests returns a HelmRequests
func newHelmRequests(c *AppV1beta1Client, namespace string) *helmRequests {
	return &helmRequests{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the helmRequest, and returns the corresponding helmRequest object, and an error if there is any.
func (c *helmRequests) Get(name string, options v1.GetOptions) (result *v1beta1.HelmRequest, err error) {
	result = &v1beta1.HelmRequest{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("helmrequests").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of HelmRequests that match those selectors.
func (c *helmRequests) List(opts v1.ListOptions) (result *v1beta1.HelmRequestList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.HelmRequestList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("helmrequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested helmRequests.
func (c *helmRequests) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("helmrequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a helmRequest and creates it.  Returns the server's representation of the helmRequest, and an error, if there is any.
func (c *helmRequests) Create(helmRequest *v1beta1.HelmRequest) (result *v1beta1.HelmRequest, err error) {
	result = &v1beta1.HelmRequest{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("helmrequests").
		Body(helmRequest).
		Do().
		Into(result)
	return
}

// Update takes the representation of a helmRequest and updates it. Returns the server's representation of the helmRequest, and an error, if there is any.
func (c *helmRequests) Update(helmRequest *v1beta1.HelmRequest) (result *v1beta1.HelmRequest, err error) {
	result = &v1beta1.HelmRequest{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("helmrequests").
		Name(helmRequest.Name).
		Body(helmRequest).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *helmRequests) UpdateStatus(helmRequest *v1beta1.HelmRequest) (result *v1beta1.HelmRequest, err error) {
	result = &v1beta1.HelmRequest{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("helmrequests").
		Name(helmRequest.Name).
		SubResource("status").
		Body(helmRequest).
		Do().
		Into(result)
	return
}

// Delete takes name of the helmRequest and deletes it. Returns an error if one occurs.
func (c *helmRequests) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("helmrequests").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *helmRequests) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("helmrequests").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched helmRequest.
func (c *helmRequests) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.HelmRequest, err error) {
	result = &v1beta1.HelmRequest{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("helmrequests").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
