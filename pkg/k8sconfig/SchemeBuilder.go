package k8sconfig

import "k8s.io/apimachinery/pkg/runtime"

const (
	TesterGroup = "extensions.practice.com"
	TesterVersion
)

var SchemeBuilder = &Builder{}

type Builder struct {
	runtime.SchemeBuilder
}

func(b *Builder) AddScheme(scheme *runtime.Scheme) error{
	 return b.AddToScheme(scheme)
}
