package main

import (
	"os"
	"path/filepath"
	"strconv"

	istionetv1beta1 "service-chains/crds/kubernetes/networking/v1beta1"

	appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/apps/v1"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	currentNamespace, _, _ := loadKubeConfig().Namespace()

	pulumi.Run(func(ctx *pulumi.Context) error {

		/* Services */

		svcNames := []pulumi.StringOutput{}
		for i := 1; i <= 5; i++ {
			name := "service-" + strconv.Itoa(i)
			labels := pulumi.StringMap{
				"app": pulumi.String(name),
			}

			svcDefinition := getServiceArgs(pulumi.String(name), labels)

			service, err := corev1.NewService(
				ctx,
				name,
				svcDefinition,
			)
			if err != nil {
				return err
			}
			svcNames = append(svcNames, service.Metadata.Elem().Name().Elem())
		}

		/* Deployments */

		// TODO doc base.Spec.(appsv1.DeploymentSpecArgs).Template.Spec.Containers[0].Args = nil // append(base.Spec.Template.Spec.Containers[0].Args, pulumi.Sprintf("-p=http://%s", svcNames[i+1-1]))
		for i := 1; i <= 5; i++ {
			baseName := "service-" + strconv.Itoa(i)
			baseLabels := pulumi.StringMap{
				"app": pulumi.String(baseName),
			}

			containerExtraArgs := pulumi.StringArray{}
			if i < 5 {
				containerExtraArgs = append(containerExtraArgs, pulumi.Sprintf("-p=http://%s", svcNames[i+1-1]))
			}

			/* Primary */

			primaryLabels := cloneMap(baseLabels)
			primaryLabels["version"] = pulumi.String("v1")
			primary := getDeploymentArgs(primaryLabels, containerExtraArgs)

			_, err := appsv1.NewDeployment(
				ctx,
				baseName+"-primary",
				primary,
			)
			if err != nil {
				return err
			}

			/* Canary */

			canaryLabels := baseLabels
			canaryLabels["version"] = pulumi.String("v2")
			canary := getDeploymentArgs(canaryLabels, containerExtraArgs)

			_, err = appsv1.NewDeployment(
				ctx,
				baseName+"-beta",
				canary,
			)
			if err != nil {
				return err
			}
		}

		/* Ingress */

		gateway, err := istionetv1beta1.NewGateway(
			ctx,
			"service-chain",
			&istionetv1beta1.GatewayArgs{
				Spec: istionetv1beta1.GatewaySpecArgs{
					Selector: pulumi.StringMap{
						"istio": pulumi.String("ingressgateway"),
					},
					Servers: istionetv1beta1.GatewaySpecServersArray{
						istionetv1beta1.GatewaySpecServersArgs{
							Port: istionetv1beta1.GatewaySpecServersPortArgs{
								Number:   pulumi.Int(80),
								Name:     pulumi.String("http"),
								Protocol: pulumi.String("HTTP"),
							},
							Hosts: pulumi.StringArray{pulumi.String("*")},
						},
					},
				},
			},
		)
		if err != nil {
			return err
		}

		_, err = istionetv1beta1.NewVirtualService(
			ctx,
			"service-chain",
			&istionetv1beta1.VirtualServiceArgs{
				Spec: istionetv1beta1.VirtualServiceSpecArgs{
					Gateways: pulumi.StringArray{gateway.Metadata.Elem().Name().Elem()},
					Hosts:    pulumi.StringArray{pulumi.String("*")},
					Http: istionetv1beta1.VirtualServiceSpecHttpArray{
						istionetv1beta1.VirtualServiceSpecHttpArgs{
							Match: istionetv1beta1.VirtualServiceSpecHttpMatchArray{
								istionetv1beta1.VirtualServiceSpecHttpMatchArgs{
									Uri: pulumi.StringMap{
										"Prefix": pulumi.String("/"),
									},
								},
							},
							Route: istionetv1beta1.VirtualServiceSpecHttpRouteArray{
								istionetv1beta1.VirtualServiceSpecHttpRouteArgs{
									Destination: istionetv1beta1.VirtualServiceSpecHttpRouteDestinationArgs{
										Host: pulumi.Sprintf("%s.%s.svc.cluster.local", svcNames[0], currentNamespace),
										Port: istionetv1beta1.VirtualServiceSpecHttpRouteDestinationPortArgs{
											Number: pulumi.Int(80),
										},
									},
								},
							},
						},
					},
				},
			},
		)
		if err != nil {
			return err
		}

		/* Fin */

		return nil
	})

}

// the operator should have a CLI mode (see zack notes, to notion journal)
// - make cli-code first, operator 2nd (package as lib)
// - checkout what controller-runtime is like the CRDs - if it's simple use that, if it's a nightmare use rust+kopium

func getDeploymentArgs(labels pulumi.StringMap, containerExtraArgs pulumi.StringArray) *appsv1.DeploymentArgs {
	return &appsv1.DeploymentArgs{
		Spec: appsv1.DeploymentSpecArgs{
			Selector: &metav1.LabelSelectorArgs{
				MatchLabels: labels,
			},
			Replicas: pulumi.Int(1),
			Template: &corev1.PodTemplateSpecArgs{
				Metadata: &metav1.ObjectMetaArgs{
					Labels: labels,
				},
				Spec: &corev1.PodSpecArgs{
					Containers: corev1.ContainerArray{
						corev1.ContainerArgs{
							Name:            pulumi.String("http-log"),
							Image:           pulumi.String("docker.io/mtinside/http-log:0.7.9"),
							ImagePullPolicy: pulumi.String("Always"),
							Args: append(
								pulumi.StringArray{
									pulumi.String("-o=pretty"),
									pulumi.String("-M"), // full headers
									pulumi.String("-r"), // response summary
								},
								containerExtraArgs...,
							),
						},
					},
				},
			},
		},
	}
}

func getServiceArgs(name pulumi.String, labels pulumi.StringMap) *corev1.ServiceArgs {
	return &corev1.ServiceArgs{
		Metadata: &metav1.ObjectMetaArgs{
			Name: name,
		},
		Spec: &corev1.ServiceSpecArgs{
			Type:     pulumi.String("ClusterIP"),
			Selector: labels,
			Ports: &corev1.ServicePortArray{
				&corev1.ServicePortArgs{
					Protocol:   pulumi.String("TCP"),
					Port:       pulumi.Int(80),
					TargetPort: pulumi.Int(8080),
				},
			},
		},
	}
}

func loadKubeConfig() clientcmd.ClientConfig {
	// Nasty
	kubeConfigPath := filepath.Join(homedir.HomeDir(), ".kube", "config")
	kubeConfigBytes, err := os.ReadFile(kubeConfigPath)
	if err != nil {
		panic(err)
	}
	kubeConfig, err := clientcmd.NewClientConfigFromBytes(kubeConfigBytes)
	if err != nil {
		panic(err)
	}
	return kubeConfig
}

// GoLang, it's 2022
func cloneMap[K comparable, V any](src map[K]V) map[K]V {
	clone := make(map[K]V, len(src))
	for k, v := range src {
		clone[k] = v
	}
	return clone
}
