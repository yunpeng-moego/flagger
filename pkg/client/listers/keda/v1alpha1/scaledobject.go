/*
Copyright 2020 The Flux authors

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/fluxcd/flagger/pkg/apis/keda/v1alpha1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// ScaledObjectLister helps list ScaledObjects.
// All objects returned here must be treated as read-only.
type ScaledObjectLister interface {
	// List lists all ScaledObjects in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ScaledObject, err error)
	// ScaledObjects returns an object that can list and get ScaledObjects.
	ScaledObjects(namespace string) ScaledObjectNamespaceLister
	ScaledObjectListerExpansion
}

// scaledObjectLister implements the ScaledObjectLister interface.
type scaledObjectLister struct {
	listers.ResourceIndexer[*v1alpha1.ScaledObject]
}

// NewScaledObjectLister returns a new ScaledObjectLister.
func NewScaledObjectLister(indexer cache.Indexer) ScaledObjectLister {
	return &scaledObjectLister{listers.New[*v1alpha1.ScaledObject](indexer, v1alpha1.Resource("scaledobject"))}
}

// ScaledObjects returns an object that can list and get ScaledObjects.
func (s *scaledObjectLister) ScaledObjects(namespace string) ScaledObjectNamespaceLister {
	return scaledObjectNamespaceLister{listers.NewNamespaced[*v1alpha1.ScaledObject](s.ResourceIndexer, namespace)}
}

// ScaledObjectNamespaceLister helps list and get ScaledObjects.
// All objects returned here must be treated as read-only.
type ScaledObjectNamespaceLister interface {
	// List lists all ScaledObjects in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ScaledObject, err error)
	// Get retrieves the ScaledObject from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ScaledObject, error)
	ScaledObjectNamespaceListerExpansion
}

// scaledObjectNamespaceLister implements the ScaledObjectNamespaceLister
// interface.
type scaledObjectNamespaceLister struct {
	listers.ResourceIndexer[*v1alpha1.ScaledObject]
}