// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package cloudtrail

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_cloudtrail_resource_policy", resourcePolicyDataSource)
}

// resourcePolicyDataSource returns the Terraform awscc_cloudtrail_resource_policy data source.
// This Terraform data source corresponds to the CloudFormation AWS::CloudTrail::ResourcePolicy resource.
func resourcePolicyDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ResourceArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the AWS CloudTrail resource to which the policy applies.",
		//	  "type": "string"
		//	}
		"resource_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the AWS CloudTrail resource to which the policy applies.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResourcePolicy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A policy document containing permissions to add to the specified resource. In IAM, you must provide policy documents in JSON format. However, in CloudFormation you can provide the policy in JSON or YAML format because CloudFormation converts YAML to JSON before submitting it to IAM.",
		//	  "type": "string"
		//	}
		"resource_policy": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A policy document containing permissions to add to the specified resource. In IAM, you must provide policy documents in JSON format. However, in CloudFormation you can provide the policy in JSON or YAML format because CloudFormation converts YAML to JSON before submitting it to IAM.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::CloudTrail::ResourcePolicy",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudTrail::ResourcePolicy").WithTerraformTypeName("awscc_cloudtrail_resource_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"resource_arn":    "ResourceArn",
		"resource_policy": "ResourcePolicy",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
