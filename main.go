/*
Copyright 2016 The Kubernetes Authors.

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

// Note: the example only works with the code within the same release/branch.
package main

import (
	"context"
	"fmt"
	"os"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	for {
		sc, err := clientset.StorageV1().StorageClasses().List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			fmt.Printf("Unable to list Storage classes: %v", err.Error())
		} else {
			fmt.Printf("Found  %d storage classes: \n", len(sc.Items))
			for _, storageClass := range sc.Items {
				fmt.Println("Annotations:")
				for name, value := range storageClass.ObjectMeta.Annotations {
					fmt.Printf("%v:%v \n", name, value)
					if (name == "storageclass.kubernetes.io/is-default-class") && (value == "true") {
						fmt.Println("Deafult storage class found, exiting")
						os.Exit(0)
					}
				}
			}
		}

		fmt.Println("No default storage class found trying again later")
		time.Sleep(5 * time.Second)
	}
}
