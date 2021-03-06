package main

import (
	sdk "github.com/gaia-pipeline/gosdk"
)

var jobs = sdk.Jobs{
	sdk.Job{
		Handler:     CreateUser,
		Title:       "Create DB User",
		Description: "Creates a database user with least privileged permissions.",
		Args: sdk.Arguments{
			sdk.Argument{
				Description: "Username for the database schema.",
				Type:        sdk.TextFieldInp,
				Key:         "username",
			},
		},
	},
	sdk.Job{
		Handler:     MigrateDB,
		Title:       "DB Migration",
		Description: "Imports newest test data dump and migrates to newest version.",
		DependsOn:   []string{"Create DB User"},
		Args: sdk.Arguments{
			sdk.Argument{
				Type: sdk.VaultInp,
				Key:  "dbpassword",
			},
		},
	},
	sdk.Job{
		Handler:     CreateNamespace,
		Title:       "Create K8S Namespace",
		Description: "Creates a new Kubernetes namespace for the new test environment.",
		DependsOn:   []string{"DB Migration"},
		Args: sdk.Arguments{
			sdk.Argument{
				Description: "Enter the name for the namespace",
				Type:        sdk.TextFieldInp,
				Key:         "namespace",
			},
		},
	},
	sdk.Job{
		Handler:     CreateDeployment,
		Title:       "Create K8S Deployment",
		Description: "Creates a new Kubernetes deployment for the new test environment.",
		DependsOn:   []string{"Create K8S namespace"},
	},
	sdk.Job{
		Handler:     CreateService,
		Title:       "Create K8S Service",
		Description: "Creates a new Kubernetes service for the new test environment.",
		DependsOn:   []string{"Create K8S namespace"},
	},
	sdk.Job{
		Handler:     CreateIngress,
		Title:       "Create K8S Ingress",
		Description: "Creates a new Kubernetes ingress for the new test environment.",
		DependsOn:   []string{"Create K8S namespace"},
	},
	sdk.Job{
		Handler:     Cleanup,
		Title:       "Clean up",
		Description: "Removes all temporary files.",
		DependsOn:   []string{"Create K8S Deployment", "Create K8S Service", "Create K8S Ingress"},
	},
}
