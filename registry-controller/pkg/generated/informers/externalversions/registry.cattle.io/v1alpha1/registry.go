/*
Copyright 2019 Wrangler Sample Controller Authors

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

// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	registrycattleiov1alpha1 "github.com/rancher/registry-controller/pkg/apis/registry.cattle.io/v1alpha1"
	versioned "github.com/rancher/registry-controller/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/rancher/registry-controller/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/rancher/registry-controller/pkg/generated/listers/registry.cattle.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// RegistryInformer provides access to a shared informer and lister for
// Registries.
type RegistryInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.RegistryLister
}

type registryInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewRegistryInformer constructs a new informer for Registry type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewRegistryInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredRegistryInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredRegistryInformer constructs a new informer for Registry type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredRegistryInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ManagementV1alpha1().Registries().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ManagementV1alpha1().Registries().Watch(context.TODO(), options)
			},
		},
		&registrycattleiov1alpha1.Registry{},
		resyncPeriod,
		indexers,
	)
}

func (f *registryInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredRegistryInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *registryInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&registrycattleiov1alpha1.Registry{}, f.defaultInformer)
}

func (f *registryInformer) Lister() v1alpha1.RegistryLister {
	return v1alpha1.NewRegistryLister(f.Informer().GetIndexer())
}