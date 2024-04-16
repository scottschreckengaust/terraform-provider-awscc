// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package detective

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_detective_member_invitation", memberInvitationResource)
}

// memberInvitationResource returns the Terraform awscc_detective_member_invitation resource.
// This Terraform resource corresponds to the CloudFormation AWS::Detective::MemberInvitation resource.
func memberInvitationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DisableEmailNotification
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": false,
		//	  "description": "When set to true, invitation emails are not sent to the member accounts. Member accounts must still accept the invitation before they are added to the behavior graph. Updating this field has no effect.",
		//	  "type": "boolean"
		//	}
		"disable_email_notification": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "When set to true, invitation emails are not sent to the member accounts. Member accounts must still accept the invitation before they are added to the behavior graph. Updating this field has no effect.",
			Optional:    true,
			Computed:    true,
			Default:     booldefault.StaticBool(false),
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// DisableEmailNotification is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: GraphArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the graph to which the member account will be invited",
		//	  "pattern": "arn:aws(-[\\w]+)*:detective:(([a-z]+-)+[0-9]+):[0-9]{12}:graph:[0-9a-f]{32}",
		//	  "type": "string"
		//	}
		"graph_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the graph to which the member account will be invited",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("arn:aws(-[\\w]+)*:detective:(([a-z]+-)+[0-9]+):[0-9]{12}:graph:[0-9a-f]{32}"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MemberEmailAddress
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The root email address for the account to be invited, for validation. Updating this field has no effect.",
		//	  "pattern": ".*@.*",
		//	  "type": "string"
		//	}
		"member_email_address": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The root email address for the account to be invited, for validation. Updating this field has no effect.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile(".*@.*"), ""),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: MemberId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The AWS account ID to be invited to join the graph as a member",
		//	  "pattern": "[0-9]{12}",
		//	  "type": "string"
		//	}
		"member_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The AWS account ID to be invited to join the graph as a member",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("[0-9]{12}"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Message
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A message to be included in the email invitation sent to the invited account. Updating this field has no effect.",
		//	  "maxLength": 1000,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"message": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A message to be included in the email invitation sent to the invited account. Updating this field has no effect.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 1000),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// Message is a write-only property.
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	// Corresponds to CloudFormation primaryIdentifier.
	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Resource schema for AWS::Detective::MemberInvitation",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Detective::MemberInvitation").WithTerraformTypeName("awscc_detective_member_invitation")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"disable_email_notification": "DisableEmailNotification",
		"graph_arn":                  "GraphArn",
		"member_email_address":       "MemberEmailAddress",
		"member_id":                  "MemberId",
		"message":                    "Message",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Message",
		"/properties/DisableEmailNotification",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
