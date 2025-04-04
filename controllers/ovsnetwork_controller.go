/*
Copyright 2024.

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

package controllers

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	sriovnetworkv1 "github.com/k8snetworkplumbingwg/sriov-network-operator/api/v1"
)

// OVSNetworkReconciler reconciles a OVSNetwork object
type OVSNetworkReconciler struct {
	client.Client
	Scheme            *runtime.Scheme
	genericReconciler *genericNetworkReconciler
}

//+kubebuilder:rbac:groups=sriovnetwork.openshift.io,resources=ovsnetworks,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=sriovnetwork.openshift.io,resources=ovsnetworks/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=sriovnetwork.openshift.io,resources=ovsnetworks/finalizers,verbs=update

// Reconcile loop for OVSNetwork CRs
func (r *OVSNetworkReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	return r.genericReconciler.Reconcile(ctx, req)
}

// return name of the controller
func (r *OVSNetworkReconciler) Name() string {
	return "OVSNetwork"
}

// return empty instance of the OVSNetwork CR
func (r *OVSNetworkReconciler) GetObject() NetworkCRInstance {
	return &sriovnetworkv1.OVSNetwork{}
}

// return empty list of the OVSNetwork CRs
func (r *OVSNetworkReconciler) GetObjectList() client.ObjectList {
	return &sriovnetworkv1.OVSNetworkList{}
}

// SetupWithManager sets up the controller with the Manager.
func (r *OVSNetworkReconciler) SetupWithManager(mgr ctrl.Manager) error {
	r.genericReconciler = newGenericNetworkReconciler(r.Client, r.Scheme, r)
	return r.genericReconciler.SetupWithManager(mgr)
}
