[![.github/workflows/test.yml](https://github.com/gambtho/aks_create_action/actions/workflows/test.yml/badge.svg)](https://github.com/gambtho/aks_create_action/actions/workflows/test.yml)

# AKS Cluster Creation action

This action creates an Azure Kubernetes Service Cluster using Pulumi

## Setup

Making use of this action requires an Azure Service Principal and a resource group.

These can be created using the setup.sh script in this repo

```
./setup.sh -c <<cluster name> -g <<resource group name>> -s <<subscription id>> -r <<region>>
```

The output from this command should look like this and matches the variables that need to be passed to the action

```
CLUSTER_NAME: testCluster
RESOURCE_GROUP_NAME: newGroup
STORAGE_ACCOUNT_NAME: newgroup27941
STORAGE_CONTAINER_NAME: testclustertstate
STORAGE_ACCESS_KEY: ******
ARM_CLIENT_ID: ******
ARM_CLIENT_SECRET: ******
ARM_SUBSCRIPTION_ID: ******
ARM_TENANT_ID: ******
```


## Inputs

* `CLUSTER_NAME` ***required***
* `RESOURCE_GROUP_NAME` ***required***
* `STORAGE_ACCOUNT_NAME` ***required***
* `STORAGE_CONTAINER_NAME` ***required***
* `STORAGE_ACCESS_KEY` ***required***
* `ARM_CLIENT_ID` ***required***
* `ARM_CLIENT_SECRET` ***required***
* `ARM_SUBSCRIPTION_ID` ***required***
* `ARM_TENANT_ID` ***required***
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
  ACTION_TYPE: create # optional
  CLUSTER_SIZE: dev # optional
  CREATE_ACR: false # optional
```

Full deployment workflow showing this action in use - https://github.com/gambtho/go_echo

## References

* https://docs.microsoft.com/en-us/azure/aks/kubernetes-action
* https://wahlnetwork.com/2020/05/12/continuous-integration-with-github-actions-and-terraform/
* https://github.com/Azure/actions-workflow-samples/tree/master/Kubernetes