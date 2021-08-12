// Code generated by generators/resource/main.go; DO NOT EDIT.

package groundstation

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
	registry.AddResourceTypeFactory("aws_groundstation_dataflow_endpoint_group", dataflowEndpointGroupResourceType)
}

// dataflowEndpointGroupResourceType returns the Terraform aws_groundstation_dataflow_endpoint_group resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::GroundStation::DataflowEndpointGroup resource type.
func dataflowEndpointGroupResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
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
		"endpoint_details": {
			// Property: EndpointDetails
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Endpoint": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "Address": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "Name": {
			//                 "type": "string"
			//               },
			//               "Port": {
			//                 "type": "integer"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "Mtu": {
			//             "type": "integer"
			//           },
			//           "Name": {
			//             "pattern": "",
			//             "type": "string"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "SecurityDetails": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "RoleArn": {
			//             "type": "string"
			//           },
			//           "SecurityGroupIds": {
			//             "items": {
			//               "type": "string"
			//             },
			//             "type": "array"
			//           },
			//           "SubnetIds": {
			//             "items": {
			//               "type": "string"
			//             },
			//             "type": "array"
			//           }
			//         },
			//         "type": "object"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "minItems": 1,
			//   "type": "array"
			// }
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"endpoint": {
						// Property: Endpoint
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"address": {
									// Property: Address
									Attributes: schema.SingleNestedAttributes(
										map[string]schema.Attribute{
											"name": {
												// Property: Name
												Type:     types.StringType,
												Optional: true,
											},
											"port": {
												// Property: Port
												Type:     types.NumberType,
												Optional: true,
											},
										},
									),
									Optional: true,
								},
								"mtu": {
									// Property: Mtu
									Type:     types.NumberType,
									Optional: true,
								},
								"name": {
									// Property: Name
									Type:     types.StringType,
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"security_details": {
						// Property: SecurityDetails
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"role_arn": {
									// Property: RoleArn
									Type:     types.StringType,
									Optional: true,
								},
								"security_group_ids": {
									// Property: SecurityGroupIds
									Type:     types.ListType{ElemType: types.StringType},
									Optional: true,
								},
								"subnet_ids": {
									// Property: SubnetIds
									Type:     types.ListType{ElemType: types.StringType},
									Optional: true,
								},
							},
						),
						Optional: true,
					},
				},
				schema.ListNestedAttributesOptions{
					MinItems: 1,
				},
			),
			Required: true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Optional: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Optional: true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
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
		Description: "AWS Ground Station DataflowEndpointGroup schema for CloudFormation",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::GroundStation::DataflowEndpointGroup").WithTerraformTypeName("aws_groundstation_dataflow_endpoint_group").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_groundstation_dataflow_endpoint_group", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
