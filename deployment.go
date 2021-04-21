package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	//appsv1 "k8s.io/api/apps/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	//"k8s.io/client-go/util/retry"
	//
)

func main() {
	var namespace *string
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	namespace = flag.String("namespace", "default", "Specify namespace. Default is default")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	flag.Parse()

	deploymentsClient := clientset.AppsV1().Deployments(*namespace)

	// List Deployments
	fmt.Printf("Listing deployments in namespace %q:\n", *namespace)
	list, err := deploymentsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Deployment                Image                  Last Update Time")
	for _, d := range list.Items {
		//		fmt.Printf(" * %s (%d replicas)\n", d.Name, *d.Spec.Replicas)
		fmt.Println(d.Name, "    ", *&d.Spec.Template.Spec.Containers[0].Image, "    ", *&d.Status.Conditions[1].LastUpdateTime)
	}

}

func int32Ptr(i int32) *int32 { return &i }
