[![.github/workflows/test.yml](https://github.com/gambtho/aks_create_action/actions/workflows/test.yml/badge.svg)](https://github.com/gambtho/aks_create_action/actions/workflows/test.yml)

# AKS Cluster Creation action

This action creates an Azure Kubernetes Service Cluster using Pulumi

## Setup

Making use of this action requires an Azure Service Principal and a resource group.

These can be created using the setup.sh script in this repo

```
./setup.sh -g <<resource group name>> -s <<subscription id>> -r <<region>>
```

The output from this command should look like this and matches the variables that need to be passed to the action

```
RESOURCE_GROUP_NAME: newGroup
ARM_CLIENT_ID: ******
ARM_CLIENT_SECRET: ******
ARM_SUBSCRIPTION_ID: ******
ARM_TENANT_ID: ******
```


## Inputs

* `RESOURCE_GROUP_NAME` ***required***
* `ARM_CLIENT_ID` ***required***
* `ARM_CLIENT_SECRET` ***required***
* `ARM_SUBSCRIPTION_ID` ***required***
* `ARM_TENANT_ID` ***required***
  
* `STORAGE_ACCOUNT_NAME` ***optional***
* `STORAGE_CONTAINER_NAME` ***optional***
* `CLUSTER_NAME` ***optional***
* `CLUSTER_SIZE` ***optional*** - dev (default) or test
* `ACTION_TYPE` ***optional*** - create (default) or delete
* `CREATE_ACR` ***optional*** - true or false (default)

## Example usage
```
uses: actions/aks_create_action@v1
with:
  CLUSTER_NAME: testCluster
  RESOURCE_GROUP_NAME: newGroup
  STORAGE_ACCOUNT_NAME: newgroup27941
  STORAGE_CONTAINER_NAME: testclustertstate
  STORAGE_ACCESS_KEY: ******
  ARM_CLIENT_ID: ******
  ARM_CLIENT_SECRET: ******
  ARM_SUBSCRIPTION_ID: ******
  ARM_TENANT_ID: ******
  CLUSTER_SIZE: dev # optional
  CREATE_ACR: false # optional
```

## References
An example of this action being used can be viewed here: https://github.com/Marcus-Hines/workflow-dummy
* https://docs.microsoft.com/en-us/azure/aks/kubernetes-action
* https://www.pulumi.com/docs/
* https://github.com/Azure/actions-workflow-samples/tree/master/Kubernetes
