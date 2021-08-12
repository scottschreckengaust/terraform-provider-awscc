// Code generated by generators/resource/main.go; DO NOT EDIT.

package ssm

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
	registry.AddResourceTypeFactory("aws_ssm_resource_data_sync", resourceDataSyncResourceType)
}

// resourceDataSyncResourceType returns the Terraform aws_ssm_resource_data_sync resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::SSM::ResourceDataSync resource type.
func resourceDataSyncResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"bucket_name": {
			// Property: BucketName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 2048,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			// BucketName is a force-new attribute.
		},
		"bucket_prefix": {
			// Property: BucketPrefix
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 64,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			// BucketPrefix is a force-new attribute.
		},
		"bucket_region": {
			// Property: BucketRegion
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			// BucketRegion is a force-new attribute.
		},
		"kms_key_arn": {
			// Property: KMSKeyArn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 512,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			// KMSKeyArn is a force-new attribute.
		},
		"s3_destination": {
			// Property: S3Destination
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "BucketName": {
			//       "maxLength": 2048,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "BucketPrefix": {
			//       "maxLength": 256,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "BucketRegion": {
			//       "maxLength": 64,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "KMSKeyArn": {
			//       "maxLength": 512,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "SyncFormat": {
			//       "maxLength": 1024,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "BucketName",
			//     "BucketRegion",
			//     "SyncFormat"
			//   ],
			//   "type": "object"
			// }
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"bucket_name": {
						// Property: BucketName
						Type:     types.StringType,
						Required: true,
					},
					"bucket_prefix": {
						// Property: BucketPrefix
						Type:     types.StringType,
						Optional: true,
					},
					"bucket_region": {
						// Property: BucketRegion
						Type:     types.StringType,
						Required: true,
					},
					"kms_key_arn": {
						// Property: KMSKeyArn
						Type:     types.StringType,
						Optional: true,
					},
					"sync_format": {
						// Property: SyncFormat
						Type:     types.StringType,
						Required: true,
					},
				},
			),
			Optional: true,
			Computed: true,
			// S3Destination is a force-new attribute.
		},
		"sync_format": {
			// Property: SyncFormat
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1024,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			// SyncFormat is a force-new attribute.
		},
		"sync_name": {
			// Property: SyncName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			// SyncName is a force-new attribute.
		},
		"sync_source": {
			// Property: SyncSource
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "AwsOrganizationsSource": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "OrganizationSourceType": {
			//           "maxLength": 64,
			//           "minLength": 1,
			//           "type": "string"
			//         },
			//         "OrganizationalUnits": {
			//           "items": {
			//             "type": "string"
			//           },
			//           "type": "array",
			//           "uniqueItems": false
			//         }
			//       },
			//       "required": [
			//         "OrganizationSourceType"
			//       ],
			//       "type": "object"
			//     },
			//     "IncludeFutureRegions": {
			//       "type": "boolean"
			//     },
			//     "SourceRegions": {
			//       "items": {
			//         "type": "string"
			//       },
			//       "type": "array",
			//       "uniqueItems": false
			//     },
			//     "SourceType": {
			//       "maxLength": 64,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "SourceType",
			//     "SourceRegions"
			//   ],
			//   "type": "object"
			// }
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"aws_organizations_source": {
						// Property: AwsOrganizationsSource
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"organization_source_type": {
									// Property: OrganizationSourceType
									Type:     types.StringType,
									Required: true,
								},
								"organizational_units": {
									// Property: OrganizationalUnits
									Type:     types.ListType{ElemType: types.StringType},
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"include_future_regions": {
						// Property: IncludeFutureRegions
						Type:     types.BoolType,
						Optional: true,
					},
					"source_regions": {
						// Property: SourceRegions
						Type:     types.ListType{ElemType: types.StringType},
						Required: true,
					},
					"source_type": {
						// Property: SourceType
						Type:     types.StringType,
						Required: true,
					},
				},
			),
			Optional: true,
		},
		"sync_type": {
			// Property: SyncType
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 64,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			// SyncType is a force-new attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource Type definition for AWS::SSM::ResourceDataSync",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SSM::ResourceDataSync").WithTerraformTypeName("aws_ssm_resource_data_sync").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_ssm_resource_data_sync", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
