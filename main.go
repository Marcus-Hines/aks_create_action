package main

import (
	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/containerservice"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {

	 	pulumi.Run(func(ctx *pulumi.Context) error {
            fmt.Println("ARM_CLIENT_ID:", os.Getenv("ARM_CLIENT_ID"))
            fmt.Println("ARM_SUBSCRIPTION_ID:", os.Getenv("ARM_SUBSCRIPTION_ID"))

		   // export ARM_CLIENT_ID
           // export ARM_CLIENT_SECRET
           // export ARM_SUBSCRIPTION_ID
           // export ARM_TENANT_ID
           // export STORAGE_ACCOUNT_NAME
           // export STORAGE_CONTAINER_NAME
           // export STORAGE_ACCESS_KEY
           // export PULUMI_ACCESS_TOKEN


			exampleKubernetesCluster, err := containerservice.NewKubernetesCluster(ctx, "mahinescluster", &containerservice.KubernetesClusterArgs{
	 			Location:          pulumi.String("East US"),
				ResourceGroupName: pulumi.String("mahinesrg"),
	 			DnsPrefix:         pulumi.String("exampleaks1"),
				DefaultNodePool: &containerservice.KubernetesClusterDefaultNodePoolArgs{
	 				Name:      pulumi.String("default"),
					NodeCount: pulumi.Int(1),
	 				VmSize:    pulumi.String("Standard_D2_v2"),
				},
	 			Identity: &containerservice.KubernetesClusterIdentityArgs{
	 				Type: pulumi.String("SystemAssigned"),
				},
	 			Tags: pulumi.StringMap{
	 				"Environment": pulumi.String("Production"),
	 			},
	 		})
	 		if err != nil {
	 			return err
	 		}


			ctx.Export("clientCertificate", exampleKubernetesCluster.KubeConfigs.ApplyT(func(kubeConfigs []containerservice.KubernetesClusterKubeConfig) (*string, error) {
	 			return kubeConfigs[0].ClientCertificate, nil
	 		}).(pulumi.StringPtrOutput))
			ctx.Export("kubeConfig", exampleKubernetesCluster.KubeConfigRaw)
	 		return nil
		})
	 }

