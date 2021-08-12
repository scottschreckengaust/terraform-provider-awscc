// Code generated by generators/resource/main.go; DO NOT EDIT.

package s3outposts

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
	registry.AddResourceTypeFactory("aws_s3outposts_endpoint", endpointResourceType)
}

// endpointResourceType returns the Terraform aws_s3outposts_endpoint resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::S3Outposts::Endpoint resource type.
func endpointResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"access_type": {
			// Property: AccessType
			// CloudFormation resource type schema:
			// {
			//   "description": "The type of access for the on-premise network connectivity for the Outpost endpoint. To access endpoint from an on-premises network, you must specify the access type and provide the customer owned Ipv4 pool.",
			//   "enum": [
			//     "CustomerOwnedIp",
			//     "Private"
			//   ],
			//   "type": "string"
			// }
			Description: "The type of access for the on-premise network connectivity for the Outpost endpoint. To access endpoint from an on-premises network, you must specify the access type and provide the customer owned Ipv4 pool.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// AccessType is a force-new attribute.
		},
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the endpoint.",
			//   "maxLength": 500,
			//   "minLength": 5,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the endpoint.",
			Type:        types.StringType,
			Computed:    true,
		},
		"cidr_block": {
			// Property: CidrBlock
			// CloudFormation resource type schema:
			// {
			//   "description": "The VPC CIDR committed by this endpoint.",
			//   "maxLength": 20,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The VPC CIDR committed by this endpoint.",
			Type:        types.StringType,
			Computed:    true,
		},
		"creation_time": {
			// Property: CreationTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The date value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ssZ)",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The date value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ssZ)",
			Type:        types.StringType,
			Computed:    true,
		},
		"customer_owned_ipv_4_pool": {
			// Property: CustomerOwnedIpv4Pool
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the customer-owned IPv4 pool for the Endpoint. IP addresses will be allocated from this pool for the endpoint.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The ID of the customer-owned IPv4 pool for the Endpoint. IP addresses will be allocated from this pool for the endpoint.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// CustomerOwnedIpv4Pool is a force-new attribute.
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the endpoint.",
			//   "maxLength": 500,
			//   "minLength": 5,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The ID of the endpoint.",
			Type:        types.StringType,
			Computed:    true,
		},
		"network_interfaces": {
			// Property: NetworkInterfaces
			// CloudFormation resource type schema:
			// {
			//   "description": "The network interfaces of the endpoint.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "The container for the network interface.",
			//     "properties": {
			//       "NetworkInterfaceId": {
			//         "maxLength": 100,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "NetworkInterfaceId"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "The network interfaces of the endpoint.",
			// Ordered set.
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"network_interface_id": {
						// Property: NetworkInterfaceId
						Type:     types.StringType,
						Required: true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"outpost_id": {
			// Property: OutpostId
			// CloudFormation resource type schema:
			// {
			//   "description": "The id of the customer outpost on which the bucket resides.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The id of the customer outpost on which the bucket resides.",
			Type:        types.StringType,
			Required:    true,
			// OutpostId is a force-new attribute.
		},
		"security_group_id": {
			// Property: SecurityGroupId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the security group to use with the endpoint.",
			//   "maxLength": 100,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The ID of the security group to use with the endpoint.",
			Type:        types.StringType,
			Required:    true,
			// SecurityGroupId is a force-new attribute.
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "Available",
			//     "Pending",
			//     "Deleting"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"subnet_id": {
			// Property: SubnetId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the subnet in the selected VPC. The subnet must belong to the Outpost.",
			//   "maxLength": 100,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The ID of the subnet in the selected VPC. The subnet must belong to the Outpost.",
			Type:        types.StringType,
			Required:    true,
			// SubnetId is a force-new attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource Type Definition for AWS::S3Outposts::Endpoint",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::S3Outposts::Endpoint").WithTerraformTypeName("aws_s3outposts_endpoint").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_s3outposts_endpoint", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
