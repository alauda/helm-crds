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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	appv1 "github.com/alauda/helm-crds/pkg/apis/app/v1"
	versioned "github.com/alauda/helm-crds/pkg/client/clientset/versioned"
	internalinterfaces "github.com/alauda/helm-crds/pkg/client/informers/externalversions/internalinterfaces"
	listerV1 "github.com/alauda/helm-crds/pkg/client/listers/app/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// HelmRequestInformer provides access to a shared informer and lister for
// HelmRequests.
type HelmRequestInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() listerV1.HelmRequestLister
}

type helmRequestInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewHelmRequestInformer constructs a new informer for HelmRequest type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewHelmRequestInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredHelmRequestInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredHelmRequestInformer constructs a new informer for HelmRequest type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredHelmRequestInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppV1().HelmRequests(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppV1().HelmRequests(namespace).Watch(options)
			},
		},
		&appv1.HelmRequest{},
		resyncPeriod,
		indexers,
	)
}

func (f *helmRequestInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredHelmRequestInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *helmRequestInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&appv1.HelmRequest{}, f.defaultInformer)
}

func (f *helmRequestInformer) Lister() listerV1.HelmRequestLister {
	return listerV1.NewHelmRequestLister(f.Informer().GetIndexer())
}