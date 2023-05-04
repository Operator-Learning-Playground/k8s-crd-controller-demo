package controller

import (
	"context"
	"fmt"
	"golanglearning/new_project/crd_practice/pkg/type_object"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// CrdPracticeController 自定义控制器
type CrdPracticeController struct {
	client.Client
	// 如果需要用到，可以加入
	Scheme *runtime.Scheme
}

func NewCrdPracticeController() *CrdPracticeController {
	return &CrdPracticeController{}
}

// Reconcile 控制器核心方法，当informer接收到add update delete事件时，会进入Reconcile方法
func (r *CrdPracticeController) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {

	obj := &type_object.Tester{}
	err := r.Get(ctx, req.NamespacedName, obj)
	if err != nil {
		return reconcile.Result{}, err
	}
	fmt.Println(obj)

	return reconcile.Result{}, nil
}

// InjectClient 注入client
func (r *CrdPracticeController) InjectClient(c client.Client) error {
	r.Client = c
	return nil
}
