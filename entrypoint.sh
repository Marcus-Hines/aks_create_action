#!/bin/sh -l

export ARM_CLIENT_ID=$INPUT_ARM_CLIENT_ID
export ARM_CLIENT_SECRET=$INPUT_ARM_CLIENT_SECRET
export ARM_SUBSCRIPTION_ID=$INPUT_ARM_SUBSCRIPTION_ID
export ARM_TENANT_ID=$INPUT_ARM_TENANT_ID
export STORAGE_ACCOUNT_NAME=$INPUT_STORAGE_ACCOUNT_NAME
export STORAGE_CONTAINER_NAME=$INPUT_STORAGE_CONTAINER_NAME
export PULUMI_ACCESS_TOKEN=$INPUT_PULUMI_ACCESS_TOKEN
export RESOURCE_GROUP_NAME=$INPUT_RESOURCE_GROUP_NAME
export REGION=$INPUT_REGION


export CLUSTER_NAME=$INPUT_CLUSTER_NAME
export CREATE_ACR=$INPUT_CREATE_ACR

## Use Pulumi based on cluster size variable | /pulumi/bin/pulumi
cd /action/$INPUT_CLUSTER_SIZE

echo "*******************"
echo "Running init"
echo "*******************"


/pulumi/bin/pulumi stack select dev --create
/pulumi/bin/pulumi config set azure:clientId ${ARM_CLIENT_ID}
/pulumi/bin/pulumi config set azure:clientSecret ${ARM_CLIENT_SECRET} --secret
/pulumi/bin/pulumi config set azure:tenantId ${ARM_TENANT_ID}
/pulumi/bin/pulumi config set azure:subscriptionId ${ARM_SUBSCRIPTION_ID}


if [ $INPUT_ACTION_TYPE = "destroy" ]; then
    echo "*******************"
    echo "Running destroy"
    echo "*******************"
    pulumi destroy -force
else
    echo "*******************"
    echo "Running apply"
    echo "*******************"
    /pulumi/bin/pulumi up --yes
fi


