/*
Copyright 2023 rakshitgondwal.

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

package controller

import (
	"context"
	"log"

	clustersv1 "github.com/rakshitgondwal/clusterify.git/api/v1"
	"github.com/rakshitgondwal/clusterify.git/pkg/civo"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// CivoClusterReconciler reconciles a CivoCluster object
type CivoClusterReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=clusters.core.clusterify.io,resources=civoclusters,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=clusters.core.clusterify.io,resources=civoclusters/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=clusters.core.clusterify.io,resources=civoclusters/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the CivoCluster object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.15.0/pkg/reconcile
func (r *CivoClusterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log.Printf("reconcile called")
	cluster := &clustersv1.CivoCluster{}

	if err := r.Get(ctx, req.NamespacedName, cluster); err != nil {
		return ctrl.Result{}, err
	}

	clusterConfig := civo.CreateClusterConfig(cluster.Spec.ClusterName, cluster.Spec.Region)
	k8sCluster, err := civo.CreateCluster(clusterConfig, cluster.Spec.APIKey)
	if err != nil {
		return ctrl.Result{}, err
	}

	if k8sCluster.Ready {
		cluster.Status.RunningStatus = "running"
	} else {
		cluster.Status.RunningStatus = "not running"
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CivoClusterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&clustersv1.CivoCluster{}).
		Complete(r)
}
