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

// ChartReposGetter has a method to return a ChartRepoInterface.
// A group's client should implement this interface.
type ChartReposGetter interface {
	ChartRepos(namespace string) ChartRepoInterface
}

// ChartRepoInterface has methods to work with ChartRepo resources.
type ChartRepoInterface interface {
	Create(*v1beta1.ChartRepo) (*v1beta1.ChartRepo, error)
	Update(*v1beta1.ChartRepo) (*v1beta1.ChartRepo, error)
	UpdateStatus(*v1beta1.ChartRepo) (*v1beta1.ChartRepo, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.ChartRepo, error)
	List(opts v1.ListOptions) (*v1beta1.ChartRepoList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ChartRepo, err error)
	ChartRepoExpansion
}

// chartRepos implements ChartRepoInterface
type chartRepos struct {
	client rest.Interface
	ns     string
}

// newChartRepos returns a ChartRepos
func newChartRepos(c *AppV1beta1Client, namespace string) *chartRepos {
	return &chartRepos{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the chartRepo, and returns the corresponding chartRepo object, and an error if there is any.
func (c *chartRepos) Get(name string, options v1.GetOptions) (result *v1beta1.ChartRepo, err error) {
	result = &v1beta1.ChartRepo{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("chartrepos").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ChartRepos that match those selectors.
func (c *chartRepos) List(opts v1.ListOptions) (result *v1beta1.ChartRepoList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.ChartRepoList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("chartrepos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested chartRepos.
func (c *chartRepos) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("chartrepos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a chartRepo and creates it.  Returns the server's representation of the chartRepo, and an error, if there is any.
func (c *chartRepos) Create(chartRepo *v1beta1.ChartRepo) (result *v1beta1.ChartRepo, err error) {
	result = &v1beta1.ChartRepo{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("chartrepos").
		Body(chartRepo).
		Do().
		Into(result)
	return
}

// Update takes the representation of a chartRepo and updates it. Returns the server's representation of the chartRepo, and an error, if there is any.
func (c *chartRepos) Update(chartRepo *v1beta1.ChartRepo) (result *v1beta1.ChartRepo, err error) {
	result = &v1beta1.ChartRepo{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("chartrepos").
		Name(chartRepo.Name).
		Body(chartRepo).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *chartRepos) UpdateStatus(chartRepo *v1beta1.ChartRepo) (result *v1beta1.ChartRepo, err error) {
	result = &v1beta1.ChartRepo{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("chartrepos").
		Name(chartRepo.Name).
		SubResource("status").
		Body(chartRepo).
		Do().
		Into(result)
	return
}

// Delete takes name of the chartRepo and deletes it. Returns an error if one occurs.
func (c *chartRepos) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("chartrepos").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *chartRepos) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("chartrepos").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched chartRepo.
func (c *chartRepos) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ChartRepo, err error) {
	result = &v1beta1.ChartRepo{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("chartrepos").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
