// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ec2_vpn_connection", vPNConnectionDataSource)
}

// vPNConnectionDataSource returns the Terraform awscc_ec2_vpn_connection data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::VPNConnection resource.
func vPNConnectionDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CustomerGatewayId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the customer gateway at your end of the VPN connection.",
		//	  "type": "string"
		//	}
		"customer_gateway_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the customer gateway at your end of the VPN connection.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EnableAcceleration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "boolean"
		//	}
		"enable_acceleration": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: StaticRoutesOnly
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether the VPN connection uses static routes only. Static routes must be used for devices that don't support BGP.\n If you are creating a VPN connection for a device that does not support Border Gateway Protocol (BGP), you must specify ``true``.",
		//	  "type": "boolean"
		//	}
		"static_routes_only": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether the VPN connection uses static routes only. Static routes must be used for devices that don't support BGP.\n If you are creating a VPN connection for a device that does not support Border Gateway Protocol (BGP), you must specify ``true``.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Any tags assigned to the VPN connection.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Specifies a tag. For more information, see [Add tags to a resource](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html#cloudformation-add-tag-specifications).",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The tag key.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The tag value.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The tag key.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The tag value.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Any tags assigned to the VPN connection.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TransitGatewayId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the transit gateway associated with the VPN connection.\n You must specify either ``TransitGatewayId`` or ``VpnGatewayId``, but not both.",
		//	  "type": "string"
		//	}
		"transit_gateway_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the transit gateway associated with the VPN connection.\n You must specify either ``TransitGatewayId`` or ``VpnGatewayId``, but not both.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of VPN connection.",
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of VPN connection.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VpnConnectionId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"vpn_connection_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VpnGatewayId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the virtual private gateway at the AWS side of the VPN connection.\n You must specify either ``TransitGatewayId`` or ``VpnGatewayId``, but not both.",
		//	  "type": "string"
		//	}
		"vpn_gateway_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the virtual private gateway at the AWS side of the VPN connection.\n You must specify either ``TransitGatewayId`` or ``VpnGatewayId``, but not both.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VpnTunnelOptionsSpecifications
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tunnel options for the VPN connection.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The tunnel options for a single VPN tunnel.",
		//	    "properties": {
		//	      "PreSharedKey": {
		//	        "description": "The pre-shared key (PSK) to establish initial authentication between the virtual private gateway and customer gateway.\n Constraints: Allowed characters are alphanumeric characters, periods (.), and underscores (_). Must be between 8 and 64 characters in length and cannot start with zero (0).",
		//	        "type": "string"
		//	      },
		//	      "TunnelInsideCidr": {
		//	        "description": "The range of inside IP addresses for the tunnel. Any specified CIDR blocks must be unique across all VPN connections that use the same virtual private gateway. \n Constraints: A size /30 CIDR block from the ``169.254.0.0/16`` range. The following CIDR blocks are reserved and cannot be used:\n  +   ``169.254.0.0/30`` \n  +   ``169.254.1.0/30`` \n  +   ``169.254.2.0/30`` \n  +   ``169.254.3.0/30`` \n  +   ``169.254.4.0/30`` \n  +   ``169.254.5.0/30`` \n  +   ``169.254.169.252/30``",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"vpn_tunnel_options_specifications": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: PreSharedKey
					"pre_shared_key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The pre-shared key (PSK) to establish initial authentication between the virtual private gateway and customer gateway.\n Constraints: Allowed characters are alphanumeric characters, periods (.), and underscores (_). Must be between 8 and 64 characters in length and cannot start with zero (0).",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: TunnelInsideCidr
					"tunnel_inside_cidr": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The range of inside IP addresses for the tunnel. Any specified CIDR blocks must be unique across all VPN connections that use the same virtual private gateway. \n Constraints: A size /30 CIDR block from the ``169.254.0.0/16`` range. The following CIDR blocks are reserved and cannot be used:\n  +   ``169.254.0.0/30`` \n  +   ``169.254.1.0/30`` \n  +   ``169.254.2.0/30`` \n  +   ``169.254.3.0/30`` \n  +   ``169.254.4.0/30`` \n  +   ``169.254.5.0/30`` \n  +   ``169.254.169.252/30``",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The tunnel options for the VPN connection.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EC2::VPNConnection",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::VPNConnection").WithTerraformTypeName("awscc_ec2_vpn_connection")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"customer_gateway_id":               "CustomerGatewayId",
		"enable_acceleration":               "EnableAcceleration",
		"key":                               "Key",
		"pre_shared_key":                    "PreSharedKey",
		"static_routes_only":                "StaticRoutesOnly",
		"tags":                              "Tags",
		"transit_gateway_id":                "TransitGatewayId",
		"tunnel_inside_cidr":                "TunnelInsideCidr",
		"type":                              "Type",
		"value":                             "Value",
		"vpn_connection_id":                 "VpnConnectionId",
		"vpn_gateway_id":                    "VpnGatewayId",
		"vpn_tunnel_options_specifications": "VpnTunnelOptionsSpecifications",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
