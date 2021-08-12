// Code generated by generators/resource/main.go; DO NOT EDIT.

package efs

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
	registry.AddResourceTypeFactory("aws_efs_file_system", fileSystemResourceType)
}

// fileSystemResourceType returns the Terraform aws_efs_file_system resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::EFS::FileSystem resource type.
func fileSystemResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"availability_zone_name": {
			// Property: AvailabilityZoneName
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			// AvailabilityZoneName is a force-new attribute.
		},
		"backup_policy": {
			// Property: BackupPolicy
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "Status": {
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Status"
			//   ],
			//   "type": "object"
			// }
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"status": {
						// Property: Status
						Type:     types.StringType,
						Required: true,
					},
				},
			),
			Optional: true,
		},
		"bypass_policy_lockout_safety_check": {
			// Property: BypassPolicyLockoutSafetyCheck
			// CloudFormation resource type schema:
			// {
			//   "description": "Whether to bypass the FileSystemPolicy lockout safety check. The policy lockout safety check determines whether the policy in the request will prevent the principal making the request to be locked out from making future PutFileSystemPolicy requests on the file system. Set BypassPolicyLockoutSafetyCheck to True only when you intend to prevent the principal that is making the request from making a subsequent PutFileSystemPolicy request on the file system. Defaults to false",
			//   "type": "boolean"
			// }
			Description: "Whether to bypass the FileSystemPolicy lockout safety check. The policy lockout safety check determines whether the policy in the request will prevent the principal making the request to be locked out from making future PutFileSystemPolicy requests on the file system. Set BypassPolicyLockoutSafetyCheck to True only when you intend to prevent the principal that is making the request from making a subsequent PutFileSystemPolicy request on the file system. Defaults to false",
			Type:        types.BoolType,
			Optional:    true,
			// BypassPolicyLockoutSafetyCheck is a write-only attribute.
		},
		"encrypted": {
			// Property: Encrypted
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Optional: true,
			Computed: true,
			// Encrypted is a force-new attribute.
		},
		"file_system_id": {
			// Property: FileSystemId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"file_system_policy": {
			// Property: FileSystemPolicy
			// CloudFormation resource type schema:
			// {
			//   "type": "object"
			// }
			Type:     types.MapType{ElemType: types.StringType},
			Optional: true,
		},
		"file_system_tags": {
			// Property: FileSystemTags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "type": "string"
			//       },
			//       "Value": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			// Ordered set.
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"kms_key_id": {
			// Property: KmsKeyId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			// KmsKeyId is a force-new attribute.
		},
		"lifecycle_policies": {
			// Property: LifecyclePolicies
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "TransitionToIA": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "TransitionToIA"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			// Ordered set.
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"transition_to_ia": {
						// Property: TransitionToIA
						Type:     types.StringType,
						Required: true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"performance_mode": {
			// Property: PerformanceMode
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			// PerformanceMode is a force-new attribute.
		},
		"provisioned_throughput_in_mibps": {
			// Property: ProvisionedThroughputInMibps
			// CloudFormation resource type schema:
			// {
			//   "type": "number"
			// }
			Type:     types.NumberType,
			Optional: true,
		},
		"throughput_mode": {
			// Property: ThroughputMode
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource Type definition for AWS::EFS::FileSystem",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::EFS::FileSystem").WithTerraformTypeName("aws_efs_file_system").WithTerraformSchema(schema)

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/BypassPolicyLockoutSafetyCheck",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_efs_file_system", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
