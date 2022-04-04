/*
Copyright The KCP Authors.

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

	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"

	v1alpha1 "github.com/kcp-dev/kcp/pkg/apis/apiresource/v1alpha1"
	apisv1alpha1 "github.com/kcp-dev/kcp/pkg/apis/apis/v1alpha1"
	tenancyv1alpha1 "github.com/kcp-dev/kcp/pkg/apis/tenancy/v1alpha1"
	v1beta1 "github.com/kcp-dev/kcp/pkg/apis/tenancy/v1beta1"
	workloadv1alpha1 "github.com/kcp-dev/kcp/pkg/apis/workload/v1alpha1"
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
	// Group=apiresource.kcp.dev, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("apiresourceimports"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apiresource().V1alpha1().APIResourceImports().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("negotiatedapiresources"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apiresource().V1alpha1().NegotiatedAPIResources().Informer()}, nil

		// Group=apis.kcp.dev, Version=v1alpha1
	case apisv1alpha1.SchemeGroupVersion.WithResource("apibindings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apis().V1alpha1().APIBindings().Informer()}, nil
	case apisv1alpha1.SchemeGroupVersion.WithResource("apiexports"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apis().V1alpha1().APIExports().Informer()}, nil
	case apisv1alpha1.SchemeGroupVersion.WithResource("apiresourceschemas"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apis().V1alpha1().APIResourceSchemas().Informer()}, nil

		// Group=tenancy.kcp.dev, Version=v1alpha1
	case tenancyv1alpha1.SchemeGroupVersion.WithResource("clusterworkspaces"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Tenancy().V1alpha1().ClusterWorkspaces().Informer()}, nil
	case tenancyv1alpha1.SchemeGroupVersion.WithResource("clusterworkspaceshards"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Tenancy().V1alpha1().ClusterWorkspaceShards().Informer()}, nil
	case tenancyv1alpha1.SchemeGroupVersion.WithResource("clusterworkspacetypes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Tenancy().V1alpha1().ClusterWorkspaceTypes().Informer()}, nil

		// Group=tenancy.kcp.dev, Version=v1beta1
	case v1beta1.SchemeGroupVersion.WithResource("workspaces"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Tenancy().V1beta1().Workspaces().Informer()}, nil

		// Group=workload.kcp.dev, Version=v1alpha1
	case workloadv1alpha1.SchemeGroupVersion.WithResource("workloadclusters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Workload().V1alpha1().WorkloadClusters().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
