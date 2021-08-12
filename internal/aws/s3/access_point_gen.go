// Code generated by generators/resource/main.go; DO NOT EDIT.

package s3

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
	registry.AddResourceTypeFactory("aws_s3_access_point", accessPointResourceType)
}

// accessPointResourceType returns the Terraform aws_s3_access_point resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::S3::AccessPoint resource type.
func accessPointResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"alias": {
			// Property: Alias
			// CloudFormation resource type schema:
			// {
			//   "description": "The alias of this Access Point. This alias can be used for compatibility purposes with other AWS services and third-party applications.",
			//   "maxLength": 63,
			//   "minLength": 3,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The alias of this Access Point. This alias can be used for compatibility purposes with other AWS services and third-party applications.",
			Type:        types.StringType,
			Computed:    true,
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "the Amazon Resource Name (ARN) of the specified accesspoint.",
			//   "type": "string"
			// }
			Description: "the Amazon Resource Name (ARN) of the specified accesspoint.",
			Type:        types.StringType,
			Computed:    true,
		},
		"bucket": {
			// Property: Bucket
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of the bucket that you want to associate this Access Point with.",
			//   "maxLength": 255,
			//   "minLength": 3,
			//   "type": "string"
			// }
			Description: "The name of the bucket that you want to associate this Access Point with.",
			Type:        types.StringType,
			Required:    true,
			// Bucket is a force-new attribute.
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "The name you want to assign to this Access Point. If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the access point name.",
			//   "maxLength": 50,
			//   "minLength": 3,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name you want to assign to this Access Point. If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the access point name.",
			Type:        types.StringType,
			Computed:    true,
			// Name is a force-new attribute.
		},
		"network_origin": {
			// Property: NetworkOrigin
			// CloudFormation resource type schema:
			// {
			//   "description": "Indicates whether this Access Point allows access from the public Internet. If VpcConfiguration is specified for this Access Point, then NetworkOrigin is VPC, and the Access Point doesn't allow access from the public Internet. Otherwise, NetworkOrigin is Internet, and the Access Point allows access from the public Internet, subject to the Access Point and bucket access policies.",
			//   "enum": [
			//     "Internet",
			//     "VPC"
			//   ],
			//   "type": "string"
			// }
			Description: "Indicates whether this Access Point allows access from the public Internet. If VpcConfiguration is specified for this Access Point, then NetworkOrigin is VPC, and the Access Point doesn't allow access from the public Internet. Otherwise, NetworkOrigin is Internet, and the Access Point allows access from the public Internet, subject to the Access Point and bucket access policies.",
			Type:        types.StringType,
			Computed:    true,
		},
		"policy": {
			// Property: Policy
			// CloudFormation resource type schema:
			// {
			//   "description": "The Access Point Policy you want to apply to this access point.",
			//   "type": "object"
			// }
			Description: "The Access Point Policy you want to apply to this access point.",
			Type:        types.MapType{ElemType: types.StringType},
			Optional:    true,
		},
		"policy_status": {
			// Property: PolicyStatus
			// CloudFormation resource type schema:
			// {
			//   "properties": {
			//     "IsPublic": {
			//       "description": "Specifies whether the policy is public or not.",
			//       "enum": [
			//         "true",
			//         "false"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"is_public": {
						// Property: IsPublic
						Description: "Specifies whether the policy is public or not.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"public_access_block_configuration": {
			// Property: PublicAccessBlockConfiguration
			// CloudFormation resource type schema:
			// {
			//   "properties": {
			//     "BlockPublicAcls": {
			//       "description": "Specifies whether Amazon S3 should block public access control lists (ACLs) for buckets in this account. Setting this element to TRUE causes the following behavior:\n- PUT Bucket acl and PUT Object acl calls fail if the specified ACL is public.\n - PUT Object calls fail if the request includes a public ACL.\n. - PUT Bucket calls fail if the request includes a public ACL.\nEnabling this setting doesn't affect existing policies or ACLs.",
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
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"block_public_acls": {
						// Property: BlockPublicAcls
						Description: "Specifies whether Amazon S3 should block public access control lists (ACLs) for buckets in this account. Setting this element to TRUE causes the following behavior:\n- PUT Bucket acl and PUT Object acl calls fail if the specified ACL is public.\n - PUT Object calls fail if the request includes a public ACL.\n. - PUT Bucket calls fail if the request includes a public ACL.\nEnabling this setting doesn't affect existing policies or ACLs.",
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
			Optional: true,
			Computed: true,
			// PublicAccessBlockConfiguration is a force-new attribute.
		},
		"vpc_configuration": {
			// Property: VpcConfiguration
			// CloudFormation resource type schema:
			// {
			//   "description": "The Virtual Private Cloud (VPC) configuration for a bucket access point.",
			//   "properties": {
			//     "VpcId": {
			//       "description": "If this field is specified, this access point will only allow connections from the specified VPC ID.",
			//       "maxLength": 1024,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "The Virtual Private Cloud (VPC) configuration for a bucket access point.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"vpc_id": {
						// Property: VpcId
						Description: "If this field is specified, this access point will only allow connections from the specified VPC ID.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Optional: true,
			Computed: true,
			// VpcConfiguration is a force-new attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "The AWS::S3::AccessPoint resource is an Amazon S3 resource type that you can use to access buckets.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::S3::AccessPoint").WithTerraformTypeName("aws_s3_access_point").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_s3_access_point", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
