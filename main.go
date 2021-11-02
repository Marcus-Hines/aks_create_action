package main

import (
	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/containerservice"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"os"
)

func main() {

	 	pulumi.Run(func(ctx *pulumi.Context) error {

			k8sCluster, err := containerservice.NewKubernetesCluster(ctx, getClusterName(), &containerservice.KubernetesClusterArgs{
	 			Location:          pulumi.String(getLocation()),
				ResourceGroupName: pulumi.String(getResourceGroup()),
	 			DnsPrefix:         pulumi.String(getDnsPrefix()),
				DefaultNodePool: &containerservice.KubernetesClusterDefaultNodePoolArgs{
	 				Name:      pulumi.String("default"),
					NodeCount: pulumi.Int(1),
	 				VmSize:    pulumi.String("Standard_D2_v2"),
				},
	 			Identity: &containerservice.KubernetesClusterIdentityArgs{
	 				Type: pulumi.String("SystemAssigned"),
				},
	 			Tags: pulumi.StringMap{
	 				"Environment": pulumi.String("Dev"),
	 			},
	 		})
	 		if err != nil {
	 			return err
	 		}


			ctx.Export("clientCertificate", k8sCluster.KubeConfigs.ApplyT(func(kubeConfigs []containerservice.KubernetesClusterKubeConfig) (*string, error) {
	 			return kubeConfigs[0].ClientCertificate, nil
	 		}).(pulumi.StringPtrOutput))
			ctx.Export("kubeConfig", k8sCluster.KubeConfigRaw)
	 		return nil
		})
	 }

	 func getClusterName() string {
	 	clusterName := os.Getenv("CLUSTER_NAME")

	 	if clusterName == "" {
	 		return "aks_create_cluster"
		}
		return clusterName
	 }

	 func getLocation() string {
	 	location := os.Getenv("REGION")

	 	if location == "" {
	 		return "East US"
		}
		return location
	 }

	 func getResourceGroup() string {
	 	resourceGroup := os.Getenv("RESOURCE_GROUP")

	 	if resourceGroup == "" {
	 		return "aks_create_rg"
		}
		return resourceGroup
	 }

	 func getDnsPrefix() string {
	 	return "akscreate"
	 }
