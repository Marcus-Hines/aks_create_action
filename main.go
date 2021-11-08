package main

import (
	"fmt"
	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/containerservice"
	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/storage"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"os"
)

const STORAGE_ACCOUNT_NAME = "akscreatesa"
const CLUSTER_NAME = "akscreatecluster"
const RESOURCE_GROUP_NAME = "akscreaterg"


func main() {

	 	pulumi.Run(func(ctx *pulumi.Context) error {
	 		location := getLocation()
	 		resourceGroup := getResourceGroup()

	 		// Create storage account
			storageAccount, err := createStorageAccount(ctx, pulumi.String(location), pulumi.String(resourceGroup))

			if err != nil {
				fmt.Print("error happened during storage account creation")
				return err
			}

	 		// Get storage account key
			//storageAccountKey, err := getStorageAccountKey()

			// Create storage container
			createStorageContainer(ctx, storageAccount.Name)



			k8sCluster, err := containerservice.NewKubernetesCluster(ctx, getClusterName(), &containerservice.KubernetesClusterArgs{
	 			Location:          pulumi.String(location),
				ResourceGroupName: pulumi.String(resourceGroup),
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
	 		return CLUSTER_NAME
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
	 	resourceGroup := os.Getenv("RESOURCE_GROUP_NAME")

	 	if resourceGroup == "" {
	 		return RESOURCE_GROUP_NAME
		}
		return resourceGroup
	 }

	 func getDnsPrefix() string {
	 	return "akscreate"
	 }

	 func createStorageAccount(ctx *pulumi.Context, location pulumi.StringInput, resourceGroup pulumi.StringInput ) (*storage.Account, error) {
		 account, err := storage.NewAccount(ctx, STORAGE_ACCOUNT_NAME, &storage.AccountArgs{
			 ResourceGroupName:      resourceGroup,
			 Location:               location,
			 AccountTier:            pulumi.String("Standard"),
			 AccountReplicationType: pulumi.String("GRS"),
			 Tags: pulumi.StringMap{
				 "environment": pulumi.String("Dev"),
			 },
		 })
		 if err != nil {
			 return nil, err
		 }
		 return account, nil
	 }

	 func getStorageAccountKey (ctx *pulumi.Context, location string, resourceGroup string ) (string, error) {
		 storageAccount, err := storage.LookupAccount(ctx, &storage.LookupAccountArgs{
			 Name:             STORAGE_ACCOUNT_NAME,
			 ResourceGroupName: &resourceGroup,
		 }, nil)

		 if err != nil {
			 return "", err
		 }
		 //ctx.Export("storageAccountTier", storageAccount.AccountTier)
		 return storageAccount.PrimaryAccessKey,nil
	 }

	 func createStorageContainer(ctx *pulumi.Context, accountName pulumi.StringInput) error {
		 _, err := storage.NewContainer(ctx, "mahinescontainer", &storage.ContainerArgs{
			 StorageAccountName:  accountName,
			 ContainerAccessType: pulumi.String("private"),
		 })
		 if err != nil {
			 return err
		 }
		 return nil
	 }


