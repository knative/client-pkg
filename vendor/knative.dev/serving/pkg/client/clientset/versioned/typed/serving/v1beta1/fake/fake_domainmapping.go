/*
Copyright 2022 The Knative Authors

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
	gentype "k8s.io/client-go/gentype"
	v1beta1 "knative.dev/serving/pkg/apis/serving/v1beta1"
	servingv1beta1 "knative.dev/serving/pkg/client/clientset/versioned/typed/serving/v1beta1"
)

// fakeDomainMappings implements DomainMappingInterface
type fakeDomainMappings struct {
	*gentype.FakeClientWithList[*v1beta1.DomainMapping, *v1beta1.DomainMappingList]
	Fake *FakeServingV1beta1
}

func newFakeDomainMappings(fake *FakeServingV1beta1, namespace string) servingv1beta1.DomainMappingInterface {
	return &fakeDomainMappings{
		gentype.NewFakeClientWithList[*v1beta1.DomainMapping, *v1beta1.DomainMappingList](
			fake.Fake,
			namespace,
			v1beta1.SchemeGroupVersion.WithResource("domainmappings"),
			v1beta1.SchemeGroupVersion.WithKind("DomainMapping"),
			func() *v1beta1.DomainMapping { return &v1beta1.DomainMapping{} },
			func() *v1beta1.DomainMappingList { return &v1beta1.DomainMappingList{} },
			func(dst, src *v1beta1.DomainMappingList) { dst.ListMeta = src.ListMeta },
			func(list *v1beta1.DomainMappingList) []*v1beta1.DomainMapping {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1beta1.DomainMappingList, items []*v1beta1.DomainMapping) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
