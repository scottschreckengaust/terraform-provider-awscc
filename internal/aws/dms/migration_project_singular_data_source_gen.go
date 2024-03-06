// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package dms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_dms_migration_project", migrationProjectDataSource)
}

// migrationProjectDataSource returns the Terraform awscc_dms_migration_project data source.
// This Terraform data source corresponds to the CloudFormation AWS::DMS::MigrationProject resource.
func migrationProjectDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The optional description of the migration project.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The optional description of the migration project.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InstanceProfileArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The property describes an instance profile arn for the migration project. For read",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"instance_profile_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The property describes an instance profile arn for the migration project. For read",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InstanceProfileIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The property describes an instance profile identifier for the migration project. For create",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"instance_profile_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The property describes an instance profile identifier for the migration project. For create",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InstanceProfileName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The property describes an instance profile name for the migration project. For read",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"instance_profile_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The property describes an instance profile name for the migration project. For read",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MigrationProjectArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The property describes an ARN of the migration project.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"migration_project_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The property describes an ARN of the migration project.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MigrationProjectCreationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The property describes a creating time of the migration project.",
		//	  "maxLength": 40,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"migration_project_creation_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The property describes a creating time of the migration project.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MigrationProjectIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The property describes an identifier for the migration project. It is used for describing/deleting/modifying can be name/arn",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"migration_project_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The property describes an identifier for the migration project. It is used for describing/deleting/modifying can be name/arn",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MigrationProjectName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The property describes a name to identify the migration project.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"migration_project_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The property describes a name to identify the migration project.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SchemaConversionApplicationAttributes
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The property describes schema conversion application attributes for the migration project.",
		//	  "properties": {
		//	    "S3BucketPath": {
		//	      "type": "string"
		//	    },
		//	    "S3BucketRoleArn": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"schema_conversion_application_attributes": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: S3BucketPath
				"s3_bucket_path": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: S3BucketRoleArn
				"s3_bucket_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The property describes schema conversion application attributes for the migration project.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SourceDataProviderDescriptors
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The property describes source data provider descriptors for the migration project.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "It is an object that describes Source and Target DataProviders and credentials for connecting to databases that are used in MigrationProject",
		//	    "properties": {
		//	      "DataProviderArn": {
		//	        "type": "string"
		//	      },
		//	      "DataProviderIdentifier": {
		//	        "type": "string"
		//	      },
		//	      "DataProviderName": {
		//	        "type": "string"
		//	      },
		//	      "SecretsManagerAccessRoleArn": {
		//	        "type": "string"
		//	      },
		//	      "SecretsManagerSecretId": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"source_data_provider_descriptors": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: DataProviderArn
					"data_provider_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: DataProviderIdentifier
					"data_provider_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: DataProviderName
					"data_provider_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: SecretsManagerAccessRoleArn
					"secrets_manager_access_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: SecretsManagerSecretId
					"secrets_manager_secret_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The property describes source data provider descriptors for the migration project.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, , and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, , and -.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, , and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, , and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TargetDataProviderDescriptors
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The property describes target data provider descriptors for the migration project.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "It is an object that describes Source and Target DataProviders and credentials for connecting to databases that are used in MigrationProject",
		//	    "properties": {
		//	      "DataProviderArn": {
		//	        "type": "string"
		//	      },
		//	      "DataProviderIdentifier": {
		//	        "type": "string"
		//	      },
		//	      "DataProviderName": {
		//	        "type": "string"
		//	      },
		//	      "SecretsManagerAccessRoleArn": {
		//	        "type": "string"
		//	      },
		//	      "SecretsManagerSecretId": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"target_data_provider_descriptors": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: DataProviderArn
					"data_provider_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: DataProviderIdentifier
					"data_provider_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: DataProviderName
					"data_provider_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: SecretsManagerAccessRoleArn
					"secrets_manager_access_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: SecretsManagerSecretId
					"secrets_manager_secret_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The property describes target data provider descriptors for the migration project.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TransformationRules
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The property describes transformation rules for the migration project.",
		//	  "type": "string"
		//	}
		"transformation_rules": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The property describes transformation rules for the migration project.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::DMS::MigrationProject",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::DMS::MigrationProject").WithTerraformTypeName("awscc_dms_migration_project")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"data_provider_arn":                        "DataProviderArn",
		"data_provider_identifier":                 "DataProviderIdentifier",
		"data_provider_name":                       "DataProviderName",
		"description":                              "Description",
		"instance_profile_arn":                     "InstanceProfileArn",
		"instance_profile_identifier":              "InstanceProfileIdentifier",
		"instance_profile_name":                    "InstanceProfileName",
		"key":                                      "Key",
		"migration_project_arn":                    "MigrationProjectArn",
		"migration_project_creation_time":          "MigrationProjectCreationTime",
		"migration_project_identifier":             "MigrationProjectIdentifier",
		"migration_project_name":                   "MigrationProjectName",
		"s3_bucket_path":                           "S3BucketPath",
		"s3_bucket_role_arn":                       "S3BucketRoleArn",
		"schema_conversion_application_attributes": "SchemaConversionApplicationAttributes",
		"secrets_manager_access_role_arn":          "SecretsManagerAccessRoleArn",
		"secrets_manager_secret_id":                "SecretsManagerSecretId",
		"source_data_provider_descriptors":         "SourceDataProviderDescriptors",
		"tags":                                     "Tags",
		"target_data_provider_descriptors":         "TargetDataProviderDescriptors",
		"transformation_rules":                     "TransformationRules",
		"value":                                    "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
