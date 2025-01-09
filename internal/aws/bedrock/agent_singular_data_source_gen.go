// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package bedrock

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_bedrock_agent", agentDataSource)
}

// agentDataSource returns the Terraform awscc_bedrock_agent data source.
// This Terraform data source corresponds to the CloudFormation AWS::Bedrock::Agent resource.
func agentDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ActionGroups
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "List of ActionGroups",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Contains the information of an Agent Action Group",
		//	    "properties": {
		//	      "ActionGroupExecutor": {
		//	        "description": "Type of Executors for an Action Group",
		//	        "properties": {
		//	          "CustomControl": {
		//	            "description": "Custom control of action execution",
		//	            "enum": [
		//	              "RETURN_CONTROL"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "Lambda": {
		//	            "description": "ARN of a Lambda.",
		//	            "maxLength": 2048,
		//	            "pattern": "^arn:(aws[a-zA-Z-]*)?:lambda:[a-z]{2}(-gov)?-[a-z]+-\\d{1}:\\d{12}:function:[a-zA-Z0-9-_\\.]+(:(\\$LATEST|[a-zA-Z0-9-_]+))?$",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "ActionGroupName": {
		//	        "description": "Name of the action group",
		//	        "pattern": "^([0-9a-zA-Z][_-]?){1,100}$",
		//	        "type": "string"
		//	      },
		//	      "ActionGroupState": {
		//	        "description": "State of the action group",
		//	        "enum": [
		//	          "ENABLED",
		//	          "DISABLED"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "ApiSchema": {
		//	        "description": "Contains information about the API Schema for the Action Group",
		//	        "properties": {
		//	          "Payload": {
		//	            "description": "String OpenAPI Payload",
		//	            "type": "string"
		//	          },
		//	          "S3": {
		//	            "additionalProperties": false,
		//	            "description": "The identifier for the S3 resource.",
		//	            "properties": {
		//	              "S3BucketName": {
		//	                "description": "A bucket in S3.",
		//	                "maxLength": 63,
		//	                "minLength": 3,
		//	                "pattern": "^[a-z0-9][\\.\\-a-z0-9]{1,61}[a-z0-9]$",
		//	                "type": "string"
		//	              },
		//	              "S3ObjectKey": {
		//	                "description": "A object key in S3.",
		//	                "maxLength": 1024,
		//	                "minLength": 1,
		//	                "pattern": "^[\\.\\-\\!\\*\\_\\'\\(\\)a-zA-Z0-9][\\.\\-\\!\\*\\_\\'\\(\\)\\/a-zA-Z0-9]*$",
		//	                "type": "string"
		//	              }
		//	            },
		//	            "type": "object"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "Description": {
		//	        "description": "Description of action group",
		//	        "maxLength": 200,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "FunctionSchema": {
		//	        "additionalProperties": false,
		//	        "description": "Schema of Functions",
		//	        "properties": {
		//	          "Functions": {
		//	            "description": "List of Function definitions",
		//	            "insertionOrder": false,
		//	            "items": {
		//	              "additionalProperties": false,
		//	              "description": "Function definition",
		//	              "properties": {
		//	                "Description": {
		//	                  "description": "Description of function",
		//	                  "maxLength": 1200,
		//	                  "minLength": 1,
		//	                  "type": "string"
		//	                },
		//	                "Name": {
		//	                  "description": "Name for a resource.",
		//	                  "pattern": "^([0-9a-zA-Z][_-]?){1,100}$",
		//	                  "type": "string"
		//	                },
		//	                "Parameters": {
		//	                  "additionalProperties": false,
		//	                  "description": "A map of parameter name and detail",
		//	                  "patternProperties": {
		//	                    "": {
		//	                      "additionalProperties": false,
		//	                      "description": "Parameter detail",
		//	                      "properties": {
		//	                        "Description": {
		//	                          "description": "Description of function parameter.",
		//	                          "maxLength": 500,
		//	                          "minLength": 1,
		//	                          "type": "string"
		//	                        },
		//	                        "Required": {
		//	                          "description": "Information about if a parameter is required for function call. Default to false.",
		//	                          "type": "boolean"
		//	                        },
		//	                        "Type": {
		//	                          "description": "Parameter Type",
		//	                          "enum": [
		//	                            "string",
		//	                            "number",
		//	                            "integer",
		//	                            "boolean",
		//	                            "array"
		//	                          ],
		//	                          "type": "string"
		//	                        }
		//	                      },
		//	                      "required": [
		//	                        "Type"
		//	                      ],
		//	                      "type": "object"
		//	                    }
		//	                  },
		//	                  "type": "object"
		//	                },
		//	                "RequireConfirmation": {
		//	                  "description": "ENUM to check if action requires user confirmation",
		//	                  "enum": [
		//	                    "ENABLED",
		//	                    "DISABLED"
		//	                  ],
		//	                  "type": "string"
		//	                }
		//	              },
		//	              "required": [
		//	                "Name"
		//	              ],
		//	              "type": "object"
		//	            },
		//	            "type": "array"
		//	          }
		//	        },
		//	        "required": [
		//	          "Functions"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "ParentActionGroupSignature": {
		//	        "description": "Action Group Signature for a BuiltIn Action",
		//	        "enum": [
		//	          "AMAZON.UserInput",
		//	          "AMAZON.CodeInterpreter"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "SkipResourceInUseCheckOnDelete": {
		//	        "default": false,
		//	        "description": "Specifies whether to allow deleting action group while it is in use.",
		//	        "type": "boolean"
		//	      }
		//	    },
		//	    "required": [
		//	      "ActionGroupName"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"action_groups": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: ActionGroupExecutor
					"action_group_executor": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: CustomControl
							"custom_control": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Custom control of action execution",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Lambda
							"lambda": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "ARN of a Lambda.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "Type of Executors for an Action Group",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: ActionGroupName
					"action_group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Name of the action group",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: ActionGroupState
					"action_group_state": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "State of the action group",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: ApiSchema
					"api_schema": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Payload
							"payload": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "String OpenAPI Payload",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: S3
							"s3": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: S3BucketName
									"s3_bucket_name": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "A bucket in S3.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: S3ObjectKey
									"s3_object_key": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "A object key in S3.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Description: "The identifier for the S3 resource.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "Contains information about the API Schema for the Action Group",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Description
					"description": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Description of action group",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: FunctionSchema
					"function_schema": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Functions
							"functions": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
								NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: Description
										"description": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "Description of function",
											Computed:    true,
										}, /*END ATTRIBUTE*/
										// Property: Name
										"name": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "Name for a resource.",
											Computed:    true,
										}, /*END ATTRIBUTE*/
										// Property: Parameters
										"parameters":              // Pattern: ""
										schema.MapNestedAttribute{ /*START ATTRIBUTE*/
											NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
												Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
													// Property: Description
													"description": schema.StringAttribute{ /*START ATTRIBUTE*/
														Description: "Description of function parameter.",
														Computed:    true,
													}, /*END ATTRIBUTE*/
													// Property: Required
													"required": schema.BoolAttribute{ /*START ATTRIBUTE*/
														Description: "Information about if a parameter is required for function call. Default to false.",
														Computed:    true,
													}, /*END ATTRIBUTE*/
													// Property: Type
													"type": schema.StringAttribute{ /*START ATTRIBUTE*/
														Description: "Parameter Type",
														Computed:    true,
													}, /*END ATTRIBUTE*/
												}, /*END SCHEMA*/
											}, /*END NESTED OBJECT*/
											Description: "A map of parameter name and detail",
											Computed:    true,
										}, /*END ATTRIBUTE*/
										// Property: RequireConfirmation
										"require_confirmation": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "ENUM to check if action requires user confirmation",
											Computed:    true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
								}, /*END NESTED OBJECT*/
								Description: "List of Function definitions",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "Schema of Functions",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: ParentActionGroupSignature
					"parent_action_group_signature": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Action Group Signature for a BuiltIn Action",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: SkipResourceInUseCheckOnDelete
					"skip_resource_in_use_check_on_delete": schema.BoolAttribute{ /*START ATTRIBUTE*/
						Description: "Specifies whether to allow deleting action group while it is in use.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "List of ActionGroups",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AgentArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Arn representation of the Agent.",
		//	  "maxLength": 2048,
		//	  "pattern": "^arn:aws(|-cn|-us-gov):bedrock:[a-z0-9-]{1,20}:[0-9]{12}:agent/[0-9a-zA-Z]{10}$",
		//	  "type": "string"
		//	}
		"agent_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Arn representation of the Agent.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AgentId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Identifier for a resource.",
		//	  "pattern": "^[0-9a-zA-Z]{10}$",
		//	  "type": "string"
		//	}
		"agent_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Identifier for a resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AgentName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name for a resource.",
		//	  "pattern": "^([0-9a-zA-Z][_-]?){1,100}$",
		//	  "type": "string"
		//	}
		"agent_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name for a resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AgentResourceRoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ARN of a IAM role.",
		//	  "maxLength": 2048,
		//	  "type": "string"
		//	}
		"agent_resource_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ARN of a IAM role.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AgentStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Schema Type for Action APIs.",
		//	  "enum": [
		//	    "CREATING",
		//	    "PREPARING",
		//	    "PREPARED",
		//	    "NOT_PREPARED",
		//	    "DELETING",
		//	    "FAILED",
		//	    "VERSIONING",
		//	    "UPDATING"
		//	  ],
		//	  "type": "string"
		//	}
		"agent_status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Schema Type for Action APIs.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AgentVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Draft Agent Version.",
		//	  "maxLength": 5,
		//	  "minLength": 5,
		//	  "pattern": "^DRAFT$",
		//	  "type": "string"
		//	}
		"agent_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Draft Agent Version.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AutoPrepare
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": false,
		//	  "description": "Specifies whether to automatically prepare after creating or updating the agent.",
		//	  "type": "boolean"
		//	}
		"auto_prepare": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies whether to automatically prepare after creating or updating the agent.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Time Stamp.",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  timetypes.RFC3339Type{},
			Description: "Time Stamp.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CustomerEncryptionKeyArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A KMS key ARN",
		//	  "maxLength": 2048,
		//	  "minLength": 1,
		//	  "pattern": "^arn:aws(|-cn|-us-gov):kms:[a-zA-Z0-9-]*:[0-9]{12}:key/[a-zA-Z0-9-]{36}$",
		//	  "type": "string"
		//	}
		"customer_encryption_key_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A KMS key ARN",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Description of the Resource.",
		//	  "maxLength": 200,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Description of the Resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: FailureReasons
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Failure Reasons for Error.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "description": "Failure Reason for Error.",
		//	    "maxLength": 2048,
		//	    "type": "string"
		//	  },
		//	  "maxItems": 2048,
		//	  "type": "array"
		//	}
		"failure_reasons": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "Failure Reasons for Error.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: FoundationModel
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ARN or name of a Bedrock model.",
		//	  "maxLength": 2048,
		//	  "minLength": 1,
		//	  "pattern": "^arn:aws(-[^:]+)?:bedrock:[a-z0-9-]{1,20}:(([0-9]{12}:custom-model/[a-z0-9-]{1,63}[.]{1}[a-z0-9-]{1,63}(([:][a-z0-9-]{1,63}){0,2})?/[a-z0-9]{12})|(:foundation-model/([a-z0-9-]{1,63}[.]{1}[a-z0-9-]{1,63}([.]?[a-z0-9-]{1,63})([:][a-z0-9-]{1,63}){0,2}))|([0-9]{12}:(inference-profile|application-inference-profile)/[a-zA-Z0-9-:.]+))|(([a-z0-9-]{1,63}[.]{1}[a-z0-9-]{1,63}([.]?[a-z0-9-]{1,63})([:][a-z0-9-]{1,63}){0,2}))|(([0-9a-zA-Z][_-]?)+)$",
		//	  "type": "string"
		//	}
		"foundation_model": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ARN or name of a Bedrock model.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: GuardrailConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Configuration for a guardrail.",
		//	  "properties": {
		//	    "GuardrailIdentifier": {
		//	      "description": "Identifier for the guardrail, could be the id or the arn",
		//	      "maxLength": 2048,
		//	      "pattern": "^(([a-z0-9]+)|(arn:aws(-[^:]+)?:bedrock:[a-z0-9-]{1,20}:[0-9]{12}:guardrail/[a-z0-9]+))$",
		//	      "type": "string"
		//	    },
		//	    "GuardrailVersion": {
		//	      "description": "Version of the guardrail",
		//	      "pattern": "^(([0-9]{1,8})|(DRAFT))$",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"guardrail_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: GuardrailIdentifier
				"guardrail_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Identifier for the guardrail, could be the id or the arn",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: GuardrailVersion
				"guardrail_version": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Version of the guardrail",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Configuration for a guardrail.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IdleSessionTTLInSeconds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Max Session Time.",
		//	  "maximum": 3600,
		//	  "minimum": 60,
		//	  "type": "number"
		//	}
		"idle_session_ttl_in_seconds": schema.Float64Attribute{ /*START ATTRIBUTE*/
			Description: "Max Session Time.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Instruction
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Instruction for the agent.",
		//	  "minLength": 40,
		//	  "type": "string"
		//	}
		"instruction": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Instruction for the agent.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: KnowledgeBases
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "List of Agent Knowledge Bases",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Agent Knowledge Base",
		//	    "properties": {
		//	      "Description": {
		//	        "description": "Description of the Resource.",
		//	        "maxLength": 200,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "KnowledgeBaseId": {
		//	        "description": "Identifier for a resource.",
		//	        "pattern": "^[0-9a-zA-Z]{10}$",
		//	        "type": "string"
		//	      },
		//	      "KnowledgeBaseState": {
		//	        "description": "State of the knowledge base; whether it is enabled or disabled",
		//	        "enum": [
		//	          "ENABLED",
		//	          "DISABLED"
		//	        ],
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "KnowledgeBaseId",
		//	      "Description"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"knowledge_bases": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Description
					"description": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Description of the Resource.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: KnowledgeBaseId
					"knowledge_base_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Identifier for a resource.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: KnowledgeBaseState
					"knowledge_base_state": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "State of the knowledge base; whether it is enabled or disabled",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "List of Agent Knowledge Bases",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PreparedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Time Stamp.",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"prepared_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  timetypes.RFC3339Type{},
			Description: "Time Stamp.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PromptOverrideConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Configuration for prompt override.",
		//	  "properties": {
		//	    "OverrideLambda": {
		//	      "description": "ARN of a Lambda.",
		//	      "maxLength": 2048,
		//	      "pattern": "^arn:(aws[a-zA-Z-]*)?:lambda:[a-z]{2}(-gov)?-[a-z]+-\\d{1}:\\d{12}:function:[a-zA-Z0-9-_\\.]+(:(\\$LATEST|[a-zA-Z0-9-_]+))?$",
		//	      "type": "string"
		//	    },
		//	    "PromptConfigurations": {
		//	      "description": "List of BasePromptConfiguration",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "BasePromptConfiguration per Prompt Type.",
		//	        "properties": {
		//	          "BasePromptTemplate": {
		//	            "description": "Base Prompt Template.",
		//	            "maxLength": 100000,
		//	            "minLength": 1,
		//	            "type": "string"
		//	          },
		//	          "InferenceConfiguration": {
		//	            "additionalProperties": false,
		//	            "description": "Configuration for inference in prompt configuration",
		//	            "properties": {
		//	              "MaximumLength": {
		//	                "description": "Maximum length of output",
		//	                "maximum": 4096,
		//	                "minimum": 0,
		//	                "type": "number"
		//	              },
		//	              "StopSequences": {
		//	                "description": "List of stop sequences",
		//	                "insertionOrder": false,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "maxItems": 4,
		//	                "minItems": 0,
		//	                "type": "array"
		//	              },
		//	              "Temperature": {
		//	                "description": "Controls randomness, higher values increase diversity",
		//	                "maximum": 1,
		//	                "minimum": 0,
		//	                "type": "number"
		//	              },
		//	              "TopK": {
		//	                "description": "Sample from the k most likely next tokens",
		//	                "maximum": 500,
		//	                "minimum": 0,
		//	                "type": "number"
		//	              },
		//	              "TopP": {
		//	                "description": "Cumulative probability cutoff for token selection",
		//	                "maximum": 1,
		//	                "minimum": 0,
		//	                "type": "number"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "ParserMode": {
		//	            "description": "Creation Mode for Prompt Configuration.",
		//	            "enum": [
		//	              "DEFAULT",
		//	              "OVERRIDDEN"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "PromptCreationMode": {
		//	            "description": "Creation Mode for Prompt Configuration.",
		//	            "enum": [
		//	              "DEFAULT",
		//	              "OVERRIDDEN"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "PromptState": {
		//	            "description": "Prompt State.",
		//	            "enum": [
		//	              "ENABLED",
		//	              "DISABLED"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "PromptType": {
		//	            "description": "Prompt Type.",
		//	            "enum": [
		//	              "PRE_PROCESSING",
		//	              "ORCHESTRATION",
		//	              "POST_PROCESSING",
		//	              "KNOWLEDGE_BASE_RESPONSE_GENERATION"
		//	            ],
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "maxItems": 10,
		//	      "type": "array"
		//	    }
		//	  },
		//	  "required": [
		//	    "PromptConfigurations"
		//	  ],
		//	  "type": "object"
		//	}
		"prompt_override_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: OverrideLambda
				"override_lambda": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "ARN of a Lambda.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: PromptConfigurations
				"prompt_configurations": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: BasePromptTemplate
							"base_prompt_template": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Base Prompt Template.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: InferenceConfiguration
							"inference_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: MaximumLength
									"maximum_length": schema.Float64Attribute{ /*START ATTRIBUTE*/
										Description: "Maximum length of output",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: StopSequences
									"stop_sequences": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Description: "List of stop sequences",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: Temperature
									"temperature": schema.Float64Attribute{ /*START ATTRIBUTE*/
										Description: "Controls randomness, higher values increase diversity",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: TopK
									"top_k": schema.Float64Attribute{ /*START ATTRIBUTE*/
										Description: "Sample from the k most likely next tokens",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: TopP
									"top_p": schema.Float64Attribute{ /*START ATTRIBUTE*/
										Description: "Cumulative probability cutoff for token selection",
										Computed:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Description: "Configuration for inference in prompt configuration",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: ParserMode
							"parser_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Creation Mode for Prompt Configuration.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: PromptCreationMode
							"prompt_creation_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Creation Mode for Prompt Configuration.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: PromptState
							"prompt_state": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Prompt State.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: PromptType
							"prompt_type": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Prompt Type.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "List of BasePromptConfiguration",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Configuration for prompt override.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RecommendedActions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The recommended actions users can take to resolve an error in failureReasons.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "description": "The recommended action users can take to resolve an error in failureReasons.",
		//	    "maxLength": 2048,
		//	    "type": "string"
		//	  },
		//	  "maxItems": 2048,
		//	  "type": "array"
		//	}
		"recommended_actions": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The recommended actions users can take to resolve an error in failureReasons.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SkipResourceInUseCheckOnDelete
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": false,
		//	  "description": "Specifies whether to allow deleting agent while it is in use.",
		//	  "type": "boolean"
		//	}
		"skip_resource_in_use_check_on_delete": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies whether to allow deleting agent while it is in use.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "A map of tag keys and values",
		//	  "patternProperties": {
		//	    "": {
		//	      "description": "Value of a tag",
		//	      "maxLength": 256,
		//	      "minLength": 0,
		//	      "pattern": "^[a-zA-Z0-9\\s._:/=+@-]*$",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A map of tag keys and values",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TestAliasTags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "A map of tag keys and values",
		//	  "patternProperties": {
		//	    "": {
		//	      "description": "Value of a tag",
		//	      "maxLength": 256,
		//	      "minLength": 0,
		//	      "pattern": "^[a-zA-Z0-9\\s._:/=+@-]*$",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"test_alias_tags":   // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A map of tag keys and values",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: UpdatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Time Stamp.",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"updated_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  timetypes.RFC3339Type{},
			Description: "Time Stamp.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Bedrock::Agent",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Bedrock::Agent").WithTerraformTypeName("awscc_bedrock_agent")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"action_group_executor":                "ActionGroupExecutor",
		"action_group_name":                    "ActionGroupName",
		"action_group_state":                   "ActionGroupState",
		"action_groups":                        "ActionGroups",
		"agent_arn":                            "AgentArn",
		"agent_id":                             "AgentId",
		"agent_name":                           "AgentName",
		"agent_resource_role_arn":              "AgentResourceRoleArn",
		"agent_status":                         "AgentStatus",
		"agent_version":                        "AgentVersion",
		"api_schema":                           "ApiSchema",
		"auto_prepare":                         "AutoPrepare",
		"base_prompt_template":                 "BasePromptTemplate",
		"created_at":                           "CreatedAt",
		"custom_control":                       "CustomControl",
		"customer_encryption_key_arn":          "CustomerEncryptionKeyArn",
		"description":                          "Description",
		"failure_reasons":                      "FailureReasons",
		"foundation_model":                     "FoundationModel",
		"function_schema":                      "FunctionSchema",
		"functions":                            "Functions",
		"guardrail_configuration":              "GuardrailConfiguration",
		"guardrail_identifier":                 "GuardrailIdentifier",
		"guardrail_version":                    "GuardrailVersion",
		"idle_session_ttl_in_seconds":          "IdleSessionTTLInSeconds",
		"inference_configuration":              "InferenceConfiguration",
		"instruction":                          "Instruction",
		"knowledge_base_id":                    "KnowledgeBaseId",
		"knowledge_base_state":                 "KnowledgeBaseState",
		"knowledge_bases":                      "KnowledgeBases",
		"lambda":                               "Lambda",
		"maximum_length":                       "MaximumLength",
		"name":                                 "Name",
		"override_lambda":                      "OverrideLambda",
		"parameters":                           "Parameters",
		"parent_action_group_signature":        "ParentActionGroupSignature",
		"parser_mode":                          "ParserMode",
		"payload":                              "Payload",
		"prepared_at":                          "PreparedAt",
		"prompt_configurations":                "PromptConfigurations",
		"prompt_creation_mode":                 "PromptCreationMode",
		"prompt_override_configuration":        "PromptOverrideConfiguration",
		"prompt_state":                         "PromptState",
		"prompt_type":                          "PromptType",
		"recommended_actions":                  "RecommendedActions",
		"require_confirmation":                 "RequireConfirmation",
		"required":                             "Required",
		"s3":                                   "S3",
		"s3_bucket_name":                       "S3BucketName",
		"s3_object_key":                        "S3ObjectKey",
		"skip_resource_in_use_check_on_delete": "SkipResourceInUseCheckOnDelete",
		"stop_sequences":                       "StopSequences",
		"tags":                                 "Tags",
		"temperature":                          "Temperature",
		"test_alias_tags":                      "TestAliasTags",
		"top_k":                                "TopK",
		"top_p":                                "TopP",
		"type":                                 "Type",
		"updated_at":                           "UpdatedAt",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
