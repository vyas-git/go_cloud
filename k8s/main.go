package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type K8s_Service struct {
	Client *kubernetes.Clientset
}

func main() {
	k8Service := NewK8sService()

	labels, replicas, err := k8Service.deploy()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Labels : %s , Replicas %d", labels, replicas)

}
func NewK8sService() *K8s_Service {
	k8Service := K8s_Service{}
	client, err := k8Service.getK8sClient()

	if err != nil {
		log.Fatal(err)
	}

	k8Service.Client = client

	return &k8Service
}
func (ks *K8s_Service) getK8sClient() (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(homedir.HomeDir(), ".kube", "config"))
	if err != nil {
		return nil, err
	}

	//create the client from config

	client, err := kubernetes.NewForConfig(config)

	if err != nil {
		return nil, err
	}
	return client, nil
}

func (ks *K8s_Service) deploy() (map[string]string, int32, error) {
	ctx := context.Background()
	//deployementFile := v1.Deployment{}

	bytes, err := ioutil.ReadFile("deploy.yml")
	if err != nil {
		return nil, 0, fmt.Errorf("Unable to read file", err)
	}
	obj, _, err := scheme.Codecs.UniversalDeserializer().Decode(bytes, nil, nil)
	if err != nil {
		return nil, 0, fmt.Errorf("Decode error: %s", err)
	}
	deploymentRespone, err := ks.Client.AppsV1().Deployments("default").Create(ctx, obj.(*v1.Deployment), metav1.CreateOptions{})

	if err != nil {
		return nil, 0, err
	}

	return deploymentRespone.Spec.Template.Labels, *deploymentRespone.Spec.Replicas, nil
}
