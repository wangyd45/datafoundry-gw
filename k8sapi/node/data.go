package node

import (
	"k8s.io/kubernetes/pkg/api/unversioned"
	kapi "k8s.io/kubernetes/pkg/api/v1"
)

type nodes struct {
	unversioned.TypeMeta `json:",inline"`
	Meta                 metadata          `json:"metadata"`
	Items                []nodeMeta `json:"items"`
}

type metadata struct {
	SelfLink        string `json:"selfLink,omitempty" protobuf:"bytes,4,opt,name=selfLink"`
	ResourceVersion string `json:"resourceVersion,omitempty" protobuf:"bytes,6,opt,name=resourceVersion"`
}

type nodeMeta struct {
	kapi.ObjectMeta `json:"metadata,omitempty"`
}


