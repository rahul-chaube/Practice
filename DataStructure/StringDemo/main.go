package main

// import (
// 	"context"
// 	"fmt"
// 	"os"

// 	" e-sdk-for-go/services/resources/mgmt/2019-05-01/resources"
// 	"github.com/Azure/go-autorest/autorest/azure/auth"
// )

// func main() {
// 	// Create a new authorizer from environment variables
// 	authorizer, err := auth.NewAuthorizerFromEnvironment()
// 	if err != nil {
// 		fmt.Println("Failed to get OAuth config or token from environment: " + err.Error())
// 		return
// 	}

// 	// Create a resources client
// 	resourcesClient := resources.NewGroupsClient(os.Getenv("AZURE_SUBSCRIPTION_ID"))
// 	resourcesClient.Authorizer = authorizer

// 	// List all resource groups in the subscription
// 	groups, err := resourcesClient.ListComplete(context.Background(), "", nil)
// 	if err != nil {
// 		fmt.Printf("Failed to list resource groups: %v\n", err)
// 		return
// 	}

// 	for groups.NotDone() {
// 		group := groups.Value()

// 		// List resources within each resource group
// 		resourcesList, err := resourcesClient.ListComplete(context.Background(), *group.Name, nil)
// 		if err != nil {
// 			fmt.Printf("Failed to list resources in resource group %s: %v\n", *group.Name, err)
// 			return
// 		}

// 		for resourcesList.NotDone() {
// 			         resource := resourcesList.Value()

// 			// Check if the resource has a managed identity
// 			if resource.Identity != nil {
// 				fmt.Printf("Resource: %s in Resource Group: %s has Managed Identity\n", *resource.Name, *group.Name)
// 				// Print or store any relevant information about the managed identity
// 				fmt.Printf("Managed Identity Type: %s\n", *resource.Identity.Type)
// 				fmt.Printf("Managed Identity Principal ID: %s\n", *resource.Identity.PrincipalID)
// 				fmt.Printf("Managed Identity Tenant ID: %s\n", *resource.Identity.TenantID)
// 			}

// 			resourcesList.NextWithContext(context.Background())
// 		}

// 		groups.NextWithContext(context.Background())
// 	}
// }
