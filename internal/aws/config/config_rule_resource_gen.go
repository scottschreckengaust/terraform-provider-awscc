// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_config_config_rule", configRuleResource)
}

// configRuleResource returns the Terraform awscc_config_config_rule resource.
// This Terraform resource corresponds to the CloudFormation AWS::Config::ConfigRule resource.
func configRuleResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Compliance
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Indicates whether an AWS resource or CC rule is compliant and provides the number of contributors that affect the compliance.",
		//	  "properties": {
		//	    "Type": {
		//	      "description": "Compliance type determined by the Config rule",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"compliance": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Type
				"type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Compliance type determined by the Config rule",
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Indicates whether an AWS resource or CC rule is compliant and provides the number of contributors that affect the compliance.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ConfigRuleId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"config_rule_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ConfigRuleName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A name for the CC rule. If you don't specify a name, CFN generates a unique physical ID and uses that ID for the rule name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).",
		//	  "type": "string"
		//	}
		"config_rule_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A name for the CC rule. If you don't specify a name, CFN generates a unique physical ID and uses that ID for the rule name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description that you provide for the CC rule.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description that you provide for the CC rule.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EvaluationModes
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The modes the CC rule can be evaluated in. The valid values are distinct objects. By default, the value is Detective evaluation mode only.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Mode of evaluation of AWS Config rule",
		//	    "properties": {
		//	      "Mode": {
		//	        "description": "The mode of an evaluation. The valid values are Detective or Proactive.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"evaluation_modes": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Mode
					"mode": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The mode of an evaluation. The valid values are Detective or Proactive.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The modes the CC rule can be evaluated in. The valid values are distinct objects. By default, the value is Detective evaluation mode only.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: InputParameters
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A string, in JSON format, that is passed to the CC rule Lambda function.",
		//	  "type": "string"
		//	}
		"input_parameters": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A string, in JSON format, that is passed to the CC rule Lambda function.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MaximumExecutionFrequency
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The maximum frequency with which CC runs evaluations for a rule. You can specify a value for ``MaximumExecutionFrequency`` when:\n  +  You are using an AWS managed rule that is triggered at a periodic frequency.\n  +  Your custom rule is triggered when CC delivers the configuration snapshot. For more information, see [ConfigSnapshotDeliveryProperties](https://docs.aws.amazon.com/config/latest/APIReference/API_ConfigSnapshotDeliveryProperties.html).\n  \n  By default, rules with a periodic trigger are evaluated every 24 hours. To change the frequency, specify a valid value for the ``MaximumExecutionFrequency`` parameter.",
		//	  "type": "string"
		//	}
		"maximum_execution_frequency": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The maximum frequency with which CC runs evaluations for a rule. You can specify a value for ``MaximumExecutionFrequency`` when:\n  +  You are using an AWS managed rule that is triggered at a periodic frequency.\n  +  Your custom rule is triggered when CC delivers the configuration snapshot. For more information, see [ConfigSnapshotDeliveryProperties](https://docs.aws.amazon.com/config/latest/APIReference/API_ConfigSnapshotDeliveryProperties.html).\n  \n  By default, rules with a periodic trigger are evaluated every 24 hours. To change the frequency, specify a valid value for the ``MaximumExecutionFrequency`` parameter.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Scope
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Defines which resources can trigger an evaluation for the rule. The scope can include one or more resource types, a combination of one resource type and one resource ID, or a combination of a tag key and value. Specify a scope to constrain the resources that can trigger an evaluation for the rule. If you do not specify a scope, evaluations are triggered when any resource in the recording group changes.\n  The scope can be empty.",
		//	  "properties": {
		//	    "ComplianceResourceId": {
		//	      "description": "The ID of the only AWS resource that you want to trigger an evaluation for the rule. If you specify a resource ID, you must specify one resource type for ``ComplianceResourceTypes``.",
		//	      "type": "string"
		//	    },
		//	    "ComplianceResourceTypes": {
		//	      "description": "The resource types of only those AWS resources that you want to trigger an evaluation for the rule. You can only specify one type if you also specify a resource ID for ``ComplianceResourceId``.",
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "TagKey": {
		//	      "description": "The tag key that is applied to only those AWS resources that you want to trigger an evaluation for the rule.",
		//	      "type": "string"
		//	    },
		//	    "TagValue": {
		//	      "description": "The tag value applied to only those AWS resources that you want to trigger an evaluation for the rule. If you specify a value for ``TagValue``, you must also specify a value for ``TagKey``.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"scope": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ComplianceResourceId
				"compliance_resource_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ID of the only AWS resource that you want to trigger an evaluation for the rule. If you specify a resource ID, you must specify one resource type for ``ComplianceResourceTypes``.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ComplianceResourceTypes
				"compliance_resource_types": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "The resource types of only those AWS resources that you want to trigger an evaluation for the rule. You can only specify one type if you also specify a resource ID for ``ComplianceResourceId``.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.UniqueValues(),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: TagKey
				"tag_key": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The tag key that is applied to only those AWS resources that you want to trigger an evaluation for the rule.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: TagValue
				"tag_value": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The tag value applied to only those AWS resources that you want to trigger an evaluation for the rule. If you specify a value for ``TagValue``, you must also specify a value for ``TagKey``.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Defines which resources can trigger an evaluation for the rule. The scope can include one or more resource types, a combination of one resource type and one resource ID, or a combination of a tag key and value. Specify a scope to constrain the resources that can trigger an evaluation for the rule. If you do not specify a scope, evaluations are triggered when any resource in the recording group changes.\n  The scope can be empty.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Source
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Provides the rule owner (```` for managed rules, ``CUSTOM_POLICY`` for Custom Policy rules, and ``CUSTOM_LAMBDA`` for Custom Lambda rules), the rule identifier, and the notifications that cause the function to evaluate your AWS resources.",
		//	  "properties": {
		//	    "CustomPolicyDetails": {
		//	      "additionalProperties": false,
		//	      "description": "Provides the runtime system, policy definition, and whether debug logging is enabled. Required when owner is set to ``CUSTOM_POLICY``.",
		//	      "properties": {
		//	        "EnableDebugLogDelivery": {
		//	          "description": "The boolean expression for enabling debug logging for your CC Custom Policy rule. The default value is ``false``.",
		//	          "type": "boolean"
		//	        },
		//	        "PolicyRuntime": {
		//	          "description": "The runtime system for your CC Custom Policy rule. Guard is a policy-as-code language that allows you to write policies that are enforced by CC Custom Policy rules. For more information about Guard, see the [Guard GitHub Repository](https://docs.aws.amazon.com/https://github.com/aws-cloudformation/cloudformation-guard).",
		//	          "type": "string"
		//	        },
		//	        "PolicyText": {
		//	          "description": "The policy definition containing the logic for your CC Custom Policy rule.",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "Owner": {
		//	      "description": "Indicates whether AWS or the customer owns and manages the CC rule.\n  CC Managed Rules are predefined rules owned by AWS. For more information, see [Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_use-managed-rules.html) in the *developer guide*.\n  CC Custom Rules are rules that you can develop either with Guard (``CUSTOM_POLICY``) or LAMlong (``CUSTOM_LAMBDA``). For more information, see [Custom Rules](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_develop-rules.html) in the *developer guide*.",
		//	      "type": "string"
		//	    },
		//	    "SourceDetails": {
		//	      "description": "Provides the source and the message types that cause CC to evaluate your AWS resources against a rule. It also provides the frequency with which you want CC to run evaluations for the rule if the trigger type is periodic.\n If the owner is set to ``CUSTOM_POLICY``, the only acceptable values for the CC rule trigger message type are ``ConfigurationItemChangeNotification`` and ``OversizedConfigurationItemChangeNotification``.",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "Source and message type that can trigger the rule",
		//	        "properties": {
		//	          "EventSource": {
		//	            "description": "The source of the event, such as an AWS service, that triggers CC to evaluate your AWS resources.",
		//	            "type": "string"
		//	          },
		//	          "MaximumExecutionFrequency": {
		//	            "description": "The frequency at which you want CC to run evaluations for a custom rule with a periodic trigger. If you specify a value for ``MaximumExecutionFrequency``, then ``MessageType`` must use the ``ScheduledNotification`` value.\n  By default, rules with a periodic trigger are evaluated every 24 hours. To change the frequency, specify a valid value for the ``MaximumExecutionFrequency`` parameter.\n Based on the valid value you choose, CC runs evaluations once for each valid value. For example, if you choose ``Three_Hours``, CC runs evaluations once every three hours. In this case, ``Three_Hours`` is the frequency of this rule.",
		//	            "type": "string"
		//	          },
		//	          "MessageType": {
		//	            "description": "The type of notification that triggers CC to run an evaluation for a rule. You can specify the following notification types:\n  +   ``ConfigurationItemChangeNotification`` - Triggers an evaluation when CC delivers a configuration item as a result of a resource change.\n  +   ``OversizedConfigurationItemChangeNotification`` - Triggers an evaluation when CC delivers an oversized configuration item. CC may generate this notification type when a resource changes and the notification exceeds the maximum size allowed by Amazon SNS.\n  +   ``ScheduledNotification`` - Triggers a periodic evaluation at the frequency specified for ``MaximumExecutionFrequency``.\n  +   ``ConfigurationSnapshotDeliveryCompleted`` - Triggers a periodic evaluation when CC delivers a configuration snapshot.\n  \n If you want your custom rule to be triggered by configuration changes, specify two SourceDetail objects, one for ``ConfigurationItemChangeNotification`` and one for ``OversizedConfigurationItemChangeNotification``.",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "EventSource",
		//	          "MessageType"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "SourceIdentifier": {
		//	      "description": "For CC Managed rules, a predefined identifier from a list. For example, ``IAM_PASSWORD_POLICY`` is a managed rule. To reference a managed rule, see [List of Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html).\n For CC Custom Lambda rules, the identifier is the Amazon Resource Name (ARN) of the rule's LAMlong function, such as ``arn:aws:lambda:us-east-2:123456789012:function:custom_rule_name``.\n For CC Custom Policy rules, this field will be ignored.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Owner"
		//	  ],
		//	  "type": "object"
		//	}
		"source": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: CustomPolicyDetails
				"custom_policy_details": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: EnableDebugLogDelivery
						"enable_debug_log_delivery": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "The boolean expression for enabling debug logging for your CC Custom Policy rule. The default value is ``false``.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
								boolplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: PolicyRuntime
						"policy_runtime": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The runtime system for your CC Custom Policy rule. Guard is a policy-as-code language that allows you to write policies that are enforced by CC Custom Policy rules. For more information about Guard, see the [Guard GitHub Repository](https://docs.aws.amazon.com/https://github.com/aws-cloudformation/cloudformation-guard).",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: PolicyText
						"policy_text": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The policy definition containing the logic for your CC Custom Policy rule.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
							// PolicyText is a write-only property.
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Provides the runtime system, policy definition, and whether debug logging is enabled. Required when owner is set to ``CUSTOM_POLICY``.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Owner
				"owner": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Indicates whether AWS or the customer owns and manages the CC rule.\n  CC Managed Rules are predefined rules owned by AWS. For more information, see [Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_use-managed-rules.html) in the *developer guide*.\n  CC Custom Rules are rules that you can develop either with Guard (``CUSTOM_POLICY``) or LAMlong (``CUSTOM_LAMBDA``). For more information, see [Custom Rules](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_develop-rules.html) in the *developer guide*.",
					Required:    true,
				}, /*END ATTRIBUTE*/
				// Property: SourceDetails
				"source_details": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: EventSource
							"event_source": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The source of the event, such as an AWS service, that triggers CC to evaluate your AWS resources.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									fwvalidators.NotNullString(),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: MaximumExecutionFrequency
							"maximum_execution_frequency": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The frequency at which you want CC to run evaluations for a custom rule with a periodic trigger. If you specify a value for ``MaximumExecutionFrequency``, then ``MessageType`` must use the ``ScheduledNotification`` value.\n  By default, rules with a periodic trigger are evaluated every 24 hours. To change the frequency, specify a valid value for the ``MaximumExecutionFrequency`` parameter.\n Based on the valid value you choose, CC runs evaluations once for each valid value. For example, if you choose ``Three_Hours``, CC runs evaluations once every three hours. In this case, ``Three_Hours`` is the frequency of this rule.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: MessageType
							"message_type": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The type of notification that triggers CC to run an evaluation for a rule. You can specify the following notification types:\n  +   ``ConfigurationItemChangeNotification`` - Triggers an evaluation when CC delivers a configuration item as a result of a resource change.\n  +   ``OversizedConfigurationItemChangeNotification`` - Triggers an evaluation when CC delivers an oversized configuration item. CC may generate this notification type when a resource changes and the notification exceeds the maximum size allowed by Amazon SNS.\n  +   ``ScheduledNotification`` - Triggers a periodic evaluation at the frequency specified for ``MaximumExecutionFrequency``.\n  +   ``ConfigurationSnapshotDeliveryCompleted`` - Triggers a periodic evaluation when CC delivers a configuration snapshot.\n  \n If you want your custom rule to be triggered by configuration changes, specify two SourceDetail objects, one for ``ConfigurationItemChangeNotification`` and one for ``OversizedConfigurationItemChangeNotification``.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									fwvalidators.NotNullString(),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "Provides the source and the message types that cause CC to evaluate your AWS resources against a rule. It also provides the frequency with which you want CC to run evaluations for the rule if the trigger type is periodic.\n If the owner is set to ``CUSTOM_POLICY``, the only acceptable values for the CC rule trigger message type are ``ConfigurationItemChangeNotification`` and ``OversizedConfigurationItemChangeNotification``.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.UniqueValues(),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: SourceIdentifier
				"source_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "For CC Managed rules, a predefined identifier from a list. For example, ``IAM_PASSWORD_POLICY`` is a managed rule. To reference a managed rule, see [List of Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html).\n For CC Custom Lambda rules, the identifier is the Amazon Resource Name (ARN) of the rule's LAMlong function, such as ``arn:aws:lambda:us-east-2:123456789012:function:custom_rule_name``.\n For CC Custom Policy rules, this field will be ignored.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Provides the rule owner (```` for managed rules, ``CUSTOM_POLICY`` for Custom Policy rules, and ``CUSTOM_LAMBDA`` for Custom Lambda rules), the rule identifier, and the notifications that cause the function to evaluate your AWS resources.",
			Required:    true,
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
		Description: "You must first create and start the CC configuration recorder in order to create CC managed rules with CFNlong. For more information, see [Managing the Configuration Recorder](https://docs.aws.amazon.com/config/latest/developerguide/stop-start-recorder.html).\n  Adds or updates an CC rule to evaluate if your AWS resources comply with your desired configurations. For information on how many CC rules you can have per account, see [Service Limits](https://docs.aws.amazon.com/config/latest/developerguide/configlimits.html) in the *Developer Guide*.\n There are two types of rules: *Managed Rules* and *Custom Rules*. You can use the ``ConfigRule`` resource to create both CC Managed Rules and CC Custom Rules.\n  CC Managed Rules are predefined, customizable rules created by CC. For a list of managed rules, see [List of Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html). If you are adding an CC managed rule, you must specify the rule's identifier for the ``SourceIdentifier`` key.\n  CC Custom Rules are rules that you create from scratch. There are two ways to create CC custom rules: with Lambda functions ([Developer Guide](https://docs.aws.amazon.com/config/latest/developerguide/gettingstarted-concepts.html#gettingstarted-concepts-function)) and with CFNGUARDshort ([Guard GitHub Repository](https://docs.aws.amazon.com/https://github.com/aws-cloudformation/cloudformation-guard)), a policy-as-code language. CC custom rules created with LAMlong are called *Custom Lambda Rules* and CC custom rules created with CFNGUARDshort are called *Custom Policy Rules*.\n If you are adding a new CC Custom LAM rule, you first need to create an LAMlong function that the rule invokes to evaluate your resources. When you use the ``ConfigRule`` resource to add a Custom LAM rule to CC, you must specify the Amazon Resource Name (ARN) that LAMlong assigns to the function. You specify the ARN in the ``SourceIdentifier`` key. This key is part of the ``Source`` object, which is part of the ``ConfigRule`` object. \n For any new CC rule that you add, specify the ``ConfigRuleName`` in the ``ConfigRule`` object. Do not specify the ``ConfigRuleArn`` or the ``ConfigRuleId``. These values are generated by CC for new rules.\n If you are updating a rule that you added previously, you can specify the rule by ``ConfigRuleName``, ``ConfigRuleId``, or ``ConfigRuleArn`` in the ``ConfigRule`` data type that you use in this request.\n For more information about developing and using CC rules, see [Evaluating Resources with Rules](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config.html) in the *Developer Guide*.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Config::ConfigRule").WithTerraformTypeName("awscc_config_config_rule")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                         "Arn",
		"compliance":                  "Compliance",
		"compliance_resource_id":      "ComplianceResourceId",
		"compliance_resource_types":   "ComplianceResourceTypes",
		"config_rule_id":              "ConfigRuleId",
		"config_rule_name":            "ConfigRuleName",
		"custom_policy_details":       "CustomPolicyDetails",
		"description":                 "Description",
		"enable_debug_log_delivery":   "EnableDebugLogDelivery",
		"evaluation_modes":            "EvaluationModes",
		"event_source":                "EventSource",
		"input_parameters":            "InputParameters",
		"maximum_execution_frequency": "MaximumExecutionFrequency",
		"message_type":                "MessageType",
		"mode":                        "Mode",
		"owner":                       "Owner",
		"policy_runtime":              "PolicyRuntime",
		"policy_text":                 "PolicyText",
		"scope":                       "Scope",
		"source":                      "Source",
		"source_details":              "SourceDetails",
		"source_identifier":           "SourceIdentifier",
		"tag_key":                     "TagKey",
		"tag_value":                   "TagValue",
		"type":                        "Type",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Source/CustomPolicyDetails/PolicyText",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
