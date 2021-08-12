// Code generated by generators/resource/main.go; DO NOT EDIT.

package s3objectlambda

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/generic"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/registry"
	providertypes "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/types"
)

func init() {
	registry.AddResourceTypeFactory("aws_s3objectlambda_access_point", accessPointResourceType)
}

// accessPointResourceType returns the Terraform aws_s3objectlambda_access_point resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::S3ObjectLambda::AccessPoint resource type.
func accessPointResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"creation_date": {
			// Property: CreationDate
			// CloudFormation resource type schema:
			// {
			//   "description": "The date and time when the Object lambda Access Point was created.",
			//   "type": "string"
			// }
			Description: "The date and time when the Object lambda Access Point was created.",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name you want to assign to this Object lambda Access Point.",
			//   "maxLength": 45,
			//   "minLength": 3,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name you want to assign to this Object lambda Access Point.",
			Type:        types.StringType,
			Required:    true,
			// Name is a force-new attribute.
		},
		"object_lambda_configuration": {
			// Property: ObjectLambdaConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Configuration to be applied to this Object lambda Access Point. It specifies Supporting Access Point, Transformation Configurations. Customers can also set if they like to enable Cloudwatch metrics for accesses to this Object lambda Access Point. Default setting for Cloudwatch metrics is disable.",
			//   "properties": {
			//     "AllowedFeatures": {
			//       "insertionOrder": false,
			//       "items": {
			//         "type": "string"
			//       },
			//       "type": "array",
			//       "uniqueItems": true
			//     },
			//     "CloudWatchMetricsEnabled": {
			//       "type": "boolean"
			//     },
			//     "SupportingAccessPoint": {
			//       "maxLength": 2048,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "TransformationConfigurations": {
			//       "insertionOrder": false,
			//       "items": {
			//         "additionalProperties": false,
			//         "description": "Configuration to define what content transformation will be applied on which S3 Action.",
			//         "properties": {
			//           "Actions": {
			//             "insertionOrder": false,
			//             "items": {
			//               "type": "string"
			//             },
			//             "type": "array",
			//             "uniqueItems": true
			//           },
			//           "ContentTransformation": {
			//             "type": "object"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "type": "array",
			//       "uniqueItems": true
			//     }
			//   },
			//   "required": [
			//     "SupportingAccessPoint",
			//     "TransformationConfigurations"
			//   ],
			//   "type": "object"
			// }
			Description: "Configuration to be applied to this Object lambda Access Point. It specifies Supporting Access Point, Transformation Configurations. Customers can also set if they like to enable Cloudwatch metrics for accesses to this Object lambda Access Point. Default setting for Cloudwatch metrics is disable.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"allowed_features": {
						// Property: AllowedFeatures
						Type:     providertypes.SetType{ElemType: types.StringType},
						Optional: true,
					},
					"cloud_watch_metrics_enabled": {
						// Property: CloudWatchMetricsEnabled
						Type:     types.BoolType,
						Optional: true,
					},
					"supporting_access_point": {
						// Property: SupportingAccessPoint
						Type:     types.StringType,
						Required: true,
					},
					"transformation_configurations": {
						// Property: TransformationConfigurations
						Attributes: providertypes.SetNestedAttributes(
							map[string]schema.Attribute{
								"actions": {
									// Property: Actions
									Type:     providertypes.SetType{ElemType: types.StringType},
									Optional: true,
								},
								"content_transformation": {
									// Property: ContentTransformation
									Type:     types.MapType{ElemType: types.StringType},
									Optional: true,
								},
							},
							providertypes.SetNestedAttributesOptions{},
						),
						Required: true,
					},
				},
			),
			Optional: true,
		},
		"policy_status": {
			// Property: PolicyStatus
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "IsPublic": {
			//       "description": "Specifies whether the Object lambda Access Point Policy is Public or not. Object lambda Access Points are private by default.",
			//       "type": "boolean"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"is_public": {
						// Property: IsPublic
						Description: "Specifies whether the Object lambda Access Point Policy is Public or not. Object lambda Access Points are private by default.",
						Type:        types.BoolType,
						Optional:    true,
					},
				},
			),
			Computed: true,
		},
		"public_access_block_configuration": {
			// Property: PublicAccessBlockConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The Public Access Block Configuration is used to block policies that would allow public access to this Object lambda Access Point. All public access to Object lambda Access Points are blocked by default, and any policy that would give public access to them will be also blocked. This behavior cannot be changed for Object lambda Access Points.",
			//   "properties": {
			//     "BlockPublicAcls": {
			//       "description": "Specifies whether Amazon S3 should block public access control lists (ACLs) to this object lambda access point. Setting this element to TRUE causes the following behavior:\n- PUT Bucket acl and PUT Object acl calls fail if the specified ACL is public.\n - PUT Object calls fail if the request includes a public ACL.\n. - PUT Bucket calls fail if the request includes a public ACL.\nEnabling this setting doesn't affect existing policies or ACLs.",
			//       "type": "boolean"
			//     },
			//     "BlockPublicPolicy": {
			//       "description": "Specifies whether Amazon S3 should block public bucket policies for buckets in this account. Setting this element to TRUE causes Amazon S3 to reject calls to PUT Bucket policy if the specified bucket policy allows public access. Enabling this setting doesn't affect existing bucket policies.",
			//       "type": "boolean"
			//     },
			//     "IgnorePublicAcls": {
			//       "description": "Specifies whether Amazon S3 should ignore public ACLs for buckets in this account. Setting this element to TRUE causes Amazon S3 to ignore all public ACLs on buckets in this account and any objects that they contain. Enabling this setting doesn't affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set.",
			//       "type": "boolean"
			//     },
			//     "RestrictPublicBuckets": {
			//       "description": "Specifies whether Amazon S3 should restrict public bucket policies for this bucket. Setting this element to TRUE restricts access to this bucket to only AWS services and authorized users within this account if the bucket has a public policy.\nEnabling this setting doesn't affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked.",
			//       "type": "boolean"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The Public Access Block Configuration is used to block policies that would allow public access to this Object lambda Access Point. All public access to Object lambda Access Points are blocked by default, and any policy that would give public access to them will be also blocked. This behavior cannot be changed for Object lambda Access Points.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"block_public_acls": {
						// Property: BlockPublicAcls
						Description: "Specifies whether Amazon S3 should block public access control lists (ACLs) to this object lambda access point. Setting this element to TRUE causes the following behavior:\n- PUT Bucket acl and PUT Object acl calls fail if the specified ACL is public.\n - PUT Object calls fail if the request includes a public ACL.\n. - PUT Bucket calls fail if the request includes a public ACL.\nEnabling this setting doesn't affect existing policies or ACLs.",
						Type:        types.BoolType,
						Optional:    true,
					},
					"block_public_policy": {
						// Property: BlockPublicPolicy
						Description: "Specifies whether Amazon S3 should block public bucket policies for buckets in this account. Setting this element to TRUE causes Amazon S3 to reject calls to PUT Bucket policy if the specified bucket policy allows public access. Enabling this setting doesn't affect existing bucket policies.",
						Type:        types.BoolType,
						Optional:    true,
					},
					"ignore_public_acls": {
						// Property: IgnorePublicAcls
						Description: "Specifies whether Amazon S3 should ignore public ACLs for buckets in this account. Setting this element to TRUE causes Amazon S3 to ignore all public ACLs on buckets in this account and any objects that they contain. Enabling this setting doesn't affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set.",
						Type:        types.BoolType,
						Optional:    true,
					},
					"restrict_public_buckets": {
						// Property: RestrictPublicBuckets
						Description: "Specifies whether Amazon S3 should restrict public bucket policies for this bucket. Setting this element to TRUE restricts access to this bucket to only AWS services and authorized users within this account if the bucket has a public policy.\nEnabling this setting doesn't affect previously stored bucket policies, except that public and cross-account access within any public bucket policy, including non-public delegation to specific accounts, is blocked.",
						Type:        types.BoolType,
						Optional:    true,
					},
				},
			),
			Computed: true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "The AWS::S3ObjectLambda::AccessPoint resource is an Amazon S3ObjectLambda resource type that you can use to add computation to S3 actions",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::S3ObjectLambda::AccessPoint").WithTerraformTypeName("aws_s3objectlambda_access_point").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_s3objectlambda_access_point", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
