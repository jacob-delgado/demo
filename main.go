package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"

	// Assuming the generated clientset is in this package
	// clientset "your/generated/clientset/versioned"
	clientset "github.com/jacob-delgado/demo/pkg/client/clientset/versioned"
	informers "github.com/jacob-delgado/demo/pkg/client/informers/externalversions"
)

var (
	kubeconfig string
)

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
}

func main() {
	ctx := SetupSignalHandler()
	logger := klog.FromContext(ctx)
	flag.Parse()

	var config *rest.Config
	if kubeconfig == "" {
		cfg, err := rest.InClusterConfig()
		if err != nil {
			logger.Error(err, "Error building kubeconfig")
			klog.FlushAndExit(klog.ExitFlushTimeout, 1)
		}
		config = cfg
	} else {
		cfg, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			logger.Error(err, "Error building kubeconfig")
			klog.FlushAndExit(klog.ExitFlushTimeout, 1)
		}
		config = cfg
	}

	kubeClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		logger.Error(err, "Error building kubernetes clientset")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}

	exampleClient, err := clientset.NewForConfig(config)
	if err != nil {
		logger.Error(err, "Error building kubernetes clientset")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}

	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(kubeClient, time.Second*30)
	exampleInformerFactory := informers.NewSharedInformerFactory(exampleClient, time.Second*30)

	controller := NewController(ctx, kubeClient, exampleClient,
		kubeInformerFactory.Apps().V1().Deployments(),
		exampleInformerFactory.Foo().V1().HelloTypes())

	kubeInformerFactory.Start(ctx.Done())
	exampleInformerFactory.Start(ctx.Done())

	if err = controller.Run(ctx, 2); err != nil {
		logger.Error(err, "Error running controller")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}
}

func SetupSignalHandler() context.Context {
	c := make(chan os.Signal, 2)
	ctx, cancel := context.WithCancel(context.Background())
	signal.Notify(c, []os.Signal{os.Interrupt, syscall.SIGTERM}...)
	go func() {
		<-c
		cancel()
		<-c
		os.Exit(1) // second signal. Exit directly.
	}()

	return ctx
}
