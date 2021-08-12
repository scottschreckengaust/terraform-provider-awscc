// Code generated by generators/resource/main.go; DO NOT EDIT.

package workspaces

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
	registry.AddResourceTypeFactory("aws_workspaces_connection_alias", connectionAliasResourceType)
}

// connectionAliasResourceType returns the Terraform aws_workspaces_connection_alias resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::WorkSpaces::ConnectionAlias resource type.
func connectionAliasResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"alias_id": {
			// Property: AliasId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 68,
			//   "minLength": 13,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"associations": {
			// Property: Associations
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "AssociatedAccountId": {
			//         "type": "string"
			//       },
			//       "AssociationStatus": {
			//         "enum": [
			//           "NOT_ASSOCIATED",
			//           "PENDING_ASSOCIATION",
			//           "ASSOCIATED_WITH_OWNER_ACCOUNT",
			//           "ASSOCIATED_WITH_SHARED_ACCOUNT",
			//           "PENDING_DISASSOCIATION"
			//         ],
			//         "type": "string"
			//       },
			//       "ConnectionIdentifier": {
			//         "maxLength": 20,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "ResourceId": {
			//         "maxLength": 1000,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "maxLength": 25,
			//   "minLength": 1,
			//   "type": "array"
			// }
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"associated_account_id": {
						// Property: AssociatedAccountId
						Type:     types.StringType,
						Optional: true,
					},
					"association_status": {
						// Property: AssociationStatus
						Type:     types.StringType,
						Optional: true,
					},
					"connection_identifier": {
						// Property: ConnectionIdentifier
						Type:     types.StringType,
						Optional: true,
					},
					"resource_id": {
						// Property: ResourceId
						Type:     types.StringType,
						Optional: true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"connection_alias_state": {
			// Property: ConnectionAliasState
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "CREATING",
			//     "CREATED",
			//     "DELETING"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"connection_string": {
			// Property: ConnectionString
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			// ConnectionString is a force-new attribute.
		},
		"tags": {
			// Property: Tags
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
			//   "uniqueItems": false
			// }
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
			Computed: true,
			// Tags is a force-new attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource Type definition for AWS::WorkSpaces::ConnectionAlias",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::WorkSpaces::ConnectionAlias").WithTerraformTypeName("aws_workspaces_connection_alias").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_workspaces_connection_alias", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
