/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	"flag"
	"os"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	"github.com/spf13/cobra"
	resumesv1 "github.com/thomasscothamilton/thomasscothamilton/apis/resumes/v1"
	resumescontrollers "github.com/thomasscothamilton/thomasscothamilton/controllers/resumes"
	//+kubebuilder:scaffold:imports
)

type Manager struct {
	MetricsAddr          string
	EnableLEaderElection bool
	ProbeAddr            string
}

// managerCmd represents the manager command
var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
	manager  = Manager{
		EnableLEaderElection: false,
		MetricsAddr:          ":8080",
		ProbeAddr:            ":8081",
	}
	managerCmd = &cobra.Command{
		Use:   "manager",
		Short: "start the manager",
		Run: func(cmd *cobra.Command, args []string) {
			opts := zap.Options{
				Development: true,
			}
			opts.BindFlags(flag.CommandLine)
			flag.Parse()

			ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))

			mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
				Scheme:                 scheme,
				MetricsBindAddress:     manager.MetricsAddr,
				Port:                   9443,
				HealthProbeBindAddress: manager.ProbeAddr,
				LeaderElection:         manager.EnableLEaderElection,
				LeaderElectionID:       "99b73599.thomasscothamilton.com",
				// LeaderElectionReleaseOnCancel defines if the leader should step down voluntarily
				// when the Manager ends. This requires the binary to immediately end when the
				// Manager is stopped, otherwise, this setting is unsafe. Setting this significantly
				// speeds up voluntary leader transitions as the new leader don't have to wait
				// LeaseDuration time first.
				//
				// In the default scaffold provided, the program ends immediately after
				// the manager stops, so would be fine to enable this option. However,
				// if you are doing or is intended to do any operation such as perform cleanups
				// after the manager stops then its usage might be unsafe.
				// LeaderElectionReleaseOnCancel: true,
			})
			if err != nil {
				setupLog.Error(err, "unable to start manager")
				os.Exit(1)
			}

			if err = (&resumescontrollers.ResumeReconciler{
				Client: mgr.GetClient(),
				Scheme: mgr.GetScheme(),
			}).SetupWithManager(mgr); err != nil {
				setupLog.Error(err, "unable to create controller", "controller", "Resume")
				os.Exit(1)
			}
			//+kubebuilder:scaffold:builder

			if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
				setupLog.Error(err, "unable to set up health check")
				os.Exit(1)
			}
			if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
				setupLog.Error(err, "unable to set up ready check")
				os.Exit(1)
			}

			setupLog.Info("starting manager")
			if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
				setupLog.Error(err, "problem running manager")
				os.Exit(1)
			}
		},
	}
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	utilruntime.Must(resumesv1.AddToScheme(scheme))
	//+kubebuilder:scaffold:scheme

	rootCmd.AddCommand(managerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// managerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// flag.StringVar(&metricsAddr, "metrics-bind-address", ":8080", "The address the metric endpoint binds to.")
	// flag.StringVar(&probeAddr, "health-probe-bind-address", ":8081", "The address the probe endpoint binds to.")
	// flag.BoolVar(&enableLeaderElection, "leader-elect", false,
	// 	"Enable leader election for controller manager. "+
	// 		"Enabling this will ensure there is only one active controller manager.")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	managerCmd.Flags().StringVarP(&manager.MetricsAddr, "metrics-bind-address", "", manager.MetricsAddr, "The address the metric endpoint binds to.")
	managerCmd.Flags().StringVarP(&manager.ProbeAddr, "health-probe-bind-address", "", manager.ProbeAddr, "The address the probe endpoint binds to.")
	managerCmd.Flags().BoolVarP(&manager.EnableLEaderElection, "leader-elect", "", manager.EnableLEaderElection,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")

}
