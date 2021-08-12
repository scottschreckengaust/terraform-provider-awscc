// Code generated by generators/resource/main.go; DO NOT EDIT.

package imagebuilder

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/generic"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("aws_imagebuilder_infrastructure_configuration", infrastructureConfigurationResourceType)
}

// infrastructureConfigurationResourceType returns the Terraform aws_imagebuilder_infrastructure_configuration resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ImageBuilder::InfrastructureConfiguration resource type.
func infrastructureConfigurationResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the infrastructure configuration.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the infrastructure configuration.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "The description of the infrastructure configuration.",
			//   "type": "string"
			// }
			Description: "The description of the infrastructure configuration.",
			Type:        types.StringType,
			Optional:    true,
		},
		"instance_profile_name": {
			// Property: InstanceProfileName
			// CloudFormation resource type schema:
			// {
			//   "description": "The instance profile of the infrastructure configuration.",
			//   "type": "string"
			// }
			Description: "The instance profile of the infrastructure configuration.",
			Type:        types.StringType,
			Required:    true,
		},
		"instance_types": {
			// Property: InstanceTypes
			// CloudFormation resource type schema:
			// {
			//   "description": "The instance types of the infrastructure configuration.",
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "The instance types of the infrastructure configuration.",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
		},
		"key_pair": {
			// Property: KeyPair
			// CloudFormation resource type schema:
			// {
			//   "description": "The EC2 key pair of the infrastructure configuration..",
			//   "type": "string"
			// }
			Description: "The EC2 key pair of the infrastructure configuration..",
			Type:        types.StringType,
			Optional:    true,
		},
		"logging": {
			// Property: Logging
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The logging configuration of the infrastructure configuration.",
			//   "properties": {
			//     "S3Logs": {
			//       "additionalProperties": false,
			//       "description": "The S3 path in which to store the logs.",
			//       "properties": {
			//         "S3BucketName": {
			//           "description": "S3BucketName",
			//           "type": "string"
			//         },
			//         "S3KeyPrefix": {
			//           "description": "S3KeyPrefix",
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The logging configuration of the infrastructure configuration.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"s3_logs": {
						// Property: S3Logs
						Description: "The S3 path in which to store the logs.",
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"s3_bucket_name": {
									// Property: S3BucketName
									Description: "S3BucketName",
									Type:        types.StringType,
									Optional:    true,
								},
								"s3_key_prefix": {
									// Property: S3KeyPrefix
									Description: "S3KeyPrefix",
									Type:        types.StringType,
									Optional:    true,
								},
							},
						),
						Optional: true,
					},
				},
			),
			Optional: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the infrastructure configuration.",
			//   "type": "string"
			// }
			Description: "The name of the infrastructure configuration.",
			Type:        types.StringType,
			Required:    true,
			// Name is a force-new attribute.
		},
		"resource_tags": {
			// Property: ResourceTags
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The tags attached to the resource created by Image Builder.",
			//   "patternProperties": {
			//     "": {
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The tags attached to the resource created by Image Builder.",
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Optional: true,
		},
		"security_group_ids": {
			// Property: SecurityGroupIds
			// CloudFormation resource type schema:
			// {
			//   "description": "The security group IDs of the infrastructure configuration.",
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "The security group IDs of the infrastructure configuration.",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
		},
		"sns_topic_arn": {
			// Property: SnsTopicArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The SNS Topic Amazon Resource Name (ARN) of the infrastructure configuration.",
			//   "type": "string"
			// }
			Description: "The SNS Topic Amazon Resource Name (ARN) of the infrastructure configuration.",
			Type:        types.StringType,
			Optional:    true,
		},
		"subnet_id": {
			// Property: SubnetId
			// CloudFormation resource type schema:
			// {
			//   "description": "The subnet ID of the infrastructure configuration.",
			//   "type": "string"
			// }
			Description: "The subnet ID of the infrastructure configuration.",
			Type:        types.StringType,
			Optional:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The tags associated with the component.",
			//   "patternProperties": {
			//     "": {
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The tags associated with the component.",
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Optional: true,
		},
		"terminate_instance_on_failure": {
			// Property: TerminateInstanceOnFailure
			// CloudFormation resource type schema:
			// {
			//   "description": "The terminate instance on failure configuration of the infrastructure configuration.",
			//   "type": "boolean"
			// }
			Description: "The terminate instance on failure configuration of the infrastructure configuration.",
			Type:        types.BoolType,
			Optional:    true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource schema for AWS::ImageBuilder::InfrastructureConfiguration",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ImageBuilder::InfrastructureConfiguration").WithTerraformTypeName("aws_imagebuilder_infrastructure_configuration").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_imagebuilder_infrastructure_configuration", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
