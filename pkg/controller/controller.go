package controller

import (
	"context"
	"fmt"
	"golanglearning/new_project/crd_practice/pkg/type_object"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

//@Controller
type CrdPracticeController struct {
	client.Client
	// 如果需要用到
	Scheme *runtime.Scheme
}

func NewCrdPracticeController() *CrdPracticeController {
	return &CrdPracticeController{}
}

func (r *CrdPracticeController) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {

	obj := &type_object.Tester{}
	err := r.Get(ctx, req.NamespacedName, obj)
	if err != nil {
		return reconcile.Result{}, err
	}
	fmt.Println(obj)


	return reconcile.Result{}, nil
}


func (r *CrdPracticeController) InjectClient(c client.Client) error {
	r.Client = c
	return nil
}


