#!/bin/sh -l

export ARM_CLIENT_ID=$INPUT_ARM_CLIENT_ID
export ARM_CLIENT_SECRET=$INPUT_ARM_CLIENT_SECRET
export ARM_SUBSCRIPTION_ID=$INPUT_ARM_SUBSCRIPTION_ID
export ARM_TENANT_ID=$INPUT_ARM_TENANT_ID
export STORAGE_ACCOUNT_NAME=$INPUT_STORAGE_ACCOUNT_NAME
export STORAGE_CONTAINER_NAME=$INPUT_STORAGE_CONTAINER_NAME
export STORAGE_ACCESS_KEY=$INPUT_STORAGE_ACCESS_KEY
export PULUMI_ACCESS_TOKEN=$INPUT_PULUMI_ACCESS_TOKEN

export TF_VAR_resource_group_name=$INPUT_RESOURCE_GROUP_NAME
export TF_VAR_cluster_name=$INPUT_CLUSTER_NAME
export TF_VAR_create_acr=$INPUT_CREATE_ACR
export TF_IN_AUTOMATION=true

## Use Pulumi based on cluster size variable
cd /action/$INPUT_CLUSTER_SIZE



go version



echo "*******************"
echo "Running init"
echo "*******************"

pulumi stack select dev --create
pulumi config set azure:clientId ${ARM_CLIENT_ID}
pulumi config set azure:clientSecret ${ARM_CLIENT_SECRET} --secret
pulumi config set azure:tenantId ${ARM_TENANT_ID}
pulumi config set azure:subscriptionId ${ARM_SUBSCRIPTION_ID}


if [ $INPUT_ACTION_TYPE = "destroy" ]; then
    echo "*******************"
    echo "Running destroy"
    echo "*******************"
    pulumi destroy -force
else
    echo "*******************"
    echo "Running apply"
    echo "*******************"
    pulumi up --yes
fi


