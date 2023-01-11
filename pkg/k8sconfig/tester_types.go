package k8sconfig

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// 自定义资源对象

type TesterSpec struct {
	Version string `json:"version,omitempty"`
}

type Tester struct {
	metav1.TypeMeta `json:",inline"`
 	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec TesterSpec `json:"spec,omitempty"`
}


type TesterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items []Tester `json:"items,omitempty"`
}

func init() {
	// 自定义crd资源注册到Scheme
	SchemeBuilder.Register(func(scheme *runtime.Scheme) error {
		gv := schema.GroupVersion{
			Group: TesterGroup,
			Version: TesterVersion,
		}
		scheme.AddKnownTypes(gv, &Tester{}, &TesterList{})
		metav1.AddToGroupVersion(scheme, gv)
		return nil
	})
}
