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

package externalversions

import (
	"fmt"

	v1alpha1 "github.com/artemiscloud/activemq-artemis-operator/pkg/apis/broker/v1alpha1"
	v2alpha1 "github.com/artemiscloud/activemq-artemis-operator/pkg/apis/broker/v2alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=broker.amq.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("activemqartemissecurities"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Broker().V1alpha1().ActiveMQArtemisSecurities().Informer()}, nil
	// Group=broker.amq.io, Version=v2alpha1
	case v2alpha1.SchemeGroupVersion.WithResource("activemqartemises"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Broker().V2alpha1().ActiveMQArtemises().Informer()}, nil
	case v2alpha1.SchemeGroupVersion.WithResource("activemqartemisaddresses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Broker().V2alpha1().ActiveMQArtemisAddresses().Informer()}, nil
	case v2alpha1.SchemeGroupVersion.WithResource("activemqartemisscaledowns"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Broker().V2alpha1().ActiveMQArtemisScaledowns().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
