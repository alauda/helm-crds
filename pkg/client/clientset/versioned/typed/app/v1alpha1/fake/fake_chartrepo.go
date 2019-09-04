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
	v1alpha1 "github.com/alauda/helm-crds/pkg/apis/app/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeChartRepos implements ChartRepoInterface
type FakeChartRepos struct {
	Fake *FakeAppV1alpha1
	ns   string
}

var chartreposResource = schema.GroupVersionResource{Group: "app.alauda.io", Version: "v1alpha1", Resource: "chartrepos"}

var chartreposKind = schema.GroupVersionKind{Group: "app.alauda.io", Version: "v1alpha1", Kind: "ChartRepo"}

// Get takes name of the chartRepo, and returns the corresponding chartRepo object, and an error if there is any.
func (c *FakeChartRepos) Get(name string, options v1.GetOptions) (result *v1alpha1.ChartRepo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(chartreposResource, c.ns, name), &v1alpha1.ChartRepo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ChartRepo), err
}

// List takes label and field selectors, and returns the list of ChartRepos that match those selectors.
func (c *FakeChartRepos) List(opts v1.ListOptions) (result *v1alpha1.ChartRepoList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(chartreposResource, chartreposKind, c.ns, opts), &v1alpha1.ChartRepoList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ChartRepoList{ListMeta: obj.(*v1alpha1.ChartRepoList).ListMeta}
	for _, item := range obj.(*v1alpha1.ChartRepoList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested chartRepos.
func (c *FakeChartRepos) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(chartreposResource, c.ns, opts))

}

// Create takes the representation of a chartRepo and creates it.  Returns the server's representation of the chartRepo, and an error, if there is any.
func (c *FakeChartRepos) Create(chartRepo *v1alpha1.ChartRepo) (result *v1alpha1.ChartRepo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(chartreposResource, c.ns, chartRepo), &v1alpha1.ChartRepo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ChartRepo), err
}

// Update takes the representation of a chartRepo and updates it. Returns the server's representation of the chartRepo, and an error, if there is any.
func (c *FakeChartRepos) Update(chartRepo *v1alpha1.ChartRepo) (result *v1alpha1.ChartRepo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(chartreposResource, c.ns, chartRepo), &v1alpha1.ChartRepo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ChartRepo), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeChartRepos) UpdateStatus(chartRepo *v1alpha1.ChartRepo) (*v1alpha1.ChartRepo, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(chartreposResource, "status", c.ns, chartRepo), &v1alpha1.ChartRepo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ChartRepo), err
}

// Delete takes name of the chartRepo and deletes it. Returns an error if one occurs.
func (c *FakeChartRepos) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(chartreposResource, c.ns, name), &v1alpha1.ChartRepo{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeChartRepos) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(chartreposResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ChartRepoList{})
	return err
}

// Patch applies the patch and returns the patched chartRepo.
func (c *FakeChartRepos) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ChartRepo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(chartreposResource, c.ns, name, pt, data, subresources...), &v1alpha1.ChartRepo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ChartRepo), err
}
