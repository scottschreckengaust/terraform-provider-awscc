// Code generated by generators/resource/main.go; DO NOT EDIT.

package gamelift

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
	registry.AddResourceTypeFactory("aws_gamelift_fleet", fleetResourceType)
}

// fleetResourceType returns the Terraform aws_gamelift_fleet resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::GameLift::Fleet resource type.
func fleetResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"build_id": {
			// Property: BuildId
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A unique identifier for a build to be deployed on the new fleet. If you are deploying the fleet with a custom game build, you must specify this property. The build must have been successfully uploaded to Amazon GameLift and be in a READY status. This fleet setting cannot be changed once the fleet is created.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "A unique identifier for a build to be deployed on the new fleet. If you are deploying the fleet with a custom game build, you must specify this property. The build must have been successfully uploaded to Amazon GameLift and be in a READY status. This fleet setting cannot be changed once the fleet is created.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// BuildId is a force-new attribute.
		},
		"certificate_configuration": {
			// Property: CertificateConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Information about the use of a TLS/SSL certificate for a fleet. TLS certificate generation is enabled at the fleet level, with one certificate generated for the fleet. When this feature is enabled, the certificate can be retrieved using the GameLift Server SDK call GetInstanceCertificate. All instances in a fleet share the same certificate.",
			//   "properties": {
			//     "CertificateType": {
			//       "additionalProperties": false,
			//       "enum": [
			//         "DISABLED",
			//         "GENERATED"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "CertificateType"
			//   ],
			//   "type": "object"
			// }
			Description: "Information about the use of a TLS/SSL certificate for a fleet. TLS certificate generation is enabled at the fleet level, with one certificate generated for the fleet. When this feature is enabled, the certificate can be retrieved using the GameLift Server SDK call GetInstanceCertificate. All instances in a fleet share the same certificate.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"certificate_type": {
						// Property: CertificateType
						Type:     types.StringType,
						Required: true,
					},
				},
			),
			Optional: true,
			Computed: true,
			// CertificateConfiguration is a force-new attribute.
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A human-readable description of a fleet.",
			//   "maxLength": 1024,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "A human-readable description of a fleet.",
			Type:        types.StringType,
			Optional:    true,
		},
		"desired_ec2_instances": {
			// Property: DesiredEC2Instances
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "[DEPRECATED] The number of EC2 instances that you want this fleet to host. When creating a new fleet, GameLift automatically sets this value to \"1\" and initiates a single instance. Once the fleet is active, update this value to trigger GameLift to add or remove instances from the fleet.",
			//   "type": "integer"
			// }
			Description: "[DEPRECATED] The number of EC2 instances that you want this fleet to host. When creating a new fleet, GameLift automatically sets this value to \"1\" and initiates a single instance. Once the fleet is active, update this value to trigger GameLift to add or remove instances from the fleet.",
			Type:        types.NumberType,
			Optional:    true,
		},
		"ec2_inbound_permissions": {
			// Property: EC2InboundPermissions
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A range of IP addresses and port settings that allow inbound traffic to connect to server processes on an Amazon GameLift server.",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A range of IP addresses and port settings that allow inbound traffic to connect to server processes on an Amazon GameLift hosting resource. New game sessions that are started on the fleet are assigned an IP address/port number combination, which must fall into the fleet's allowed ranges. For fleets created with a custom game server, the ranges reflect the server's game session assignments. For Realtime Servers fleets, Amazon GameLift automatically opens two port ranges, one for TCP messaging and one for UDP, for use by the Realtime servers.",
			//     "properties": {
			//       "FromPort": {
			//         "additionalProperties": false,
			//         "description": "A starting value for a range of allowed port numbers.",
			//         "type": "integer"
			//       },
			//       "IpRange": {
			//         "additionalProperties": false,
			//         "description": "A range of allowed IP addresses. This value must be expressed in CIDR notation. Example: \"000.000.000.000/[subnet mask]\" or optionally the shortened version \"0.0.0.0/[subnet mask]\".",
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Protocol": {
			//         "additionalProperties": false,
			//         "description": "The network communication protocol used by the fleet.",
			//         "enum": [
			//           "TCP",
			//           "UDP"
			//         ],
			//         "type": "string"
			//       },
			//       "ToPort": {
			//         "additionalProperties": false,
			//         "description": "An ending value for a range of allowed port numbers. Port numbers are end-inclusive. This value must be higher than FromPort.",
			//         "type": "integer"
			//       }
			//     },
			//     "required": [
			//       "FromPort",
			//       "IpRange",
			//       "Protocol",
			//       "ToPort"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "type": "array"
			// }
			Description: "A range of IP addresses and port settings that allow inbound traffic to connect to server processes on an Amazon GameLift server.",
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"from_port": {
						// Property: FromPort
						Description: "A starting value for a range of allowed port numbers.",
						Type:        types.NumberType,
						Required:    true,
					},
					"ip_range": {
						// Property: IpRange
						Description: "A range of allowed IP addresses. This value must be expressed in CIDR notation. Example: \"000.000.000.000/[subnet mask]\" or optionally the shortened version \"0.0.0.0/[subnet mask]\".",
						Type:        types.StringType,
						Required:    true,
					},
					"protocol": {
						// Property: Protocol
						Description: "The network communication protocol used by the fleet.",
						Type:        types.StringType,
						Required:    true,
					},
					"to_port": {
						// Property: ToPort
						Description: "An ending value for a range of allowed port numbers. Port numbers are end-inclusive. This value must be higher than FromPort.",
						Type:        types.NumberType,
						Required:    true,
					},
				},
				schema.ListNestedAttributesOptions{
					MaxItems: 50,
				},
			),
			Optional: true,
		},
		"ec2_instance_type": {
			// Property: EC2InstanceType
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The name of an EC2 instance type that is supported in Amazon GameLift. A fleet instance type determines the computing resources of each instance in the fleet, including CPU, memory, storage, and networking capacity. Amazon GameLift supports the following EC2 instance types. See Amazon EC2 Instance Types for detailed descriptions.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of an EC2 instance type that is supported in Amazon GameLift. A fleet instance type determines the computing resources of each instance in the fleet, including CPU, memory, storage, and networking capacity. Amazon GameLift supports the following EC2 instance types. See Amazon EC2 Instance Types for detailed descriptions.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// EC2InstanceType is a force-new attribute.
		},
		"fleet_id": {
			// Property: FleetId
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Unique fleet ID",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Unique fleet ID",
			Type:        types.StringType,
			Computed:    true,
		},
		"fleet_type": {
			// Property: FleetType
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Indicates whether to use On-Demand instances or Spot instances for this fleet. If empty, the default is ON_DEMAND. Both categories of instances use identical hardware and configurations based on the instance type selected for this fleet.",
			//   "enum": [
			//     "ON_DEMAND",
			//     "SPOT"
			//   ],
			//   "type": "string"
			// }
			Description: "Indicates whether to use On-Demand instances or Spot instances for this fleet. If empty, the default is ON_DEMAND. Both categories of instances use identical hardware and configurations based on the instance type selected for this fleet.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// FleetType is a force-new attribute.
		},
		"instance_role_arn": {
			// Property: InstanceRoleARN
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A unique identifier for an AWS IAM role that manages access to your AWS services. With an instance role ARN set, any application that runs on an instance in this fleet can assume the role, including install scripts, server processes, and daemons (background processes). Create a role or look up a role's ARN from the IAM dashboard in the AWS Management Console.",
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "A unique identifier for an AWS IAM role that manages access to your AWS services. With an instance role ARN set, any application that runs on an instance in this fleet can assume the role, including install scripts, server processes, and daemons (background processes). Create a role or look up a role's ARN from the IAM dashboard in the AWS Management Console.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// InstanceRoleARN is a force-new attribute.
		},
		"locations": {
			// Property: Locations
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A remote location where a multi-location fleet can deploy EC2 instances for game hosting.",
			//     "properties": {
			//       "Location": {
			//         "additionalProperties": false,
			//         "maxLength": 64,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "LocationCapacity": {
			//         "additionalProperties": false,
			//         "description": "Current resource capacity settings in a specified fleet or location. The location value might refer to a fleet's remote location or its home Region.",
			//         "properties": {
			//           "DesiredEC2Instances": {
			//             "additionalProperties": false,
			//             "description": "The number of EC2 instances you want to maintain in the specified fleet location. This value must fall between the minimum and maximum size limits.",
			//             "type": "integer"
			//           },
			//           "MaxSize": {
			//             "additionalProperties": false,
			//             "description": "The maximum value that is allowed for the fleet's instance count for a location. When creating a new fleet, GameLift automatically sets this value to \"1\". Once the fleet is active, you can change this value.",
			//             "type": "integer"
			//           },
			//           "MinSize": {
			//             "additionalProperties": false,
			//             "description": "The minimum value allowed for the fleet's instance count for a location. When creating a new fleet, GameLift automatically sets this value to \"0\". After the fleet is active, you can change this value.",
			//             "type": "integer"
			//           }
			//         },
			//         "required": [
			//           "DesiredEC2Instances",
			//           "MinSize",
			//           "MaxSize"
			//         ],
			//         "type": "object"
			//       }
			//     },
			//     "required": [
			//       "Location"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 100,
			//   "minItems": 1,
			//   "type": "array"
			// }
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"location": {
						// Property: Location
						Type:     types.StringType,
						Required: true,
					},
					"location_capacity": {
						// Property: LocationCapacity
						Description: "Current resource capacity settings in a specified fleet or location. The location value might refer to a fleet's remote location or its home Region.",
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"desired_ec2_instances": {
									// Property: DesiredEC2Instances
									Description: "The number of EC2 instances you want to maintain in the specified fleet location. This value must fall between the minimum and maximum size limits.",
									Type:        types.NumberType,
									Required:    true,
								},
								"max_size": {
									// Property: MaxSize
									Description: "The maximum value that is allowed for the fleet's instance count for a location. When creating a new fleet, GameLift automatically sets this value to \"1\". Once the fleet is active, you can change this value.",
									Type:        types.NumberType,
									Required:    true,
								},
								"min_size": {
									// Property: MinSize
									Description: "The minimum value allowed for the fleet's instance count for a location. When creating a new fleet, GameLift automatically sets this value to \"0\". After the fleet is active, you can change this value.",
									Type:        types.NumberType,
									Required:    true,
								},
							},
						),
						Optional: true,
					},
				},
				schema.ListNestedAttributesOptions{
					MinItems: 1,
					MaxItems: 100,
				},
			),
			Optional: true,
		},
		"log_paths": {
			// Property: LogPaths
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "This parameter is no longer used. When hosting a custom game build, specify where Amazon GameLift should store log files using the Amazon GameLift server API call ProcessReady()",
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "This parameter is no longer used. When hosting a custom game build, specify where Amazon GameLift should store log files using the Amazon GameLift server API call ProcessReady()",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
			Computed:    true,
			// LogPaths is a force-new attribute.
		},
		"max_size": {
			// Property: MaxSize
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "[DEPRECATED] The maximum value that is allowed for the fleet's instance count. When creating a new fleet, GameLift automatically sets this value to \"1\". Once the fleet is active, you can change this value.",
			//   "type": "integer"
			// }
			Description: "[DEPRECATED] The maximum value that is allowed for the fleet's instance count. When creating a new fleet, GameLift automatically sets this value to \"1\". Once the fleet is active, you can change this value.",
			Type:        types.NumberType,
			Optional:    true,
		},
		"metric_groups": {
			// Property: MetricGroups
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The name of an Amazon CloudWatch metric group. A metric group aggregates the metrics for all fleets in the group. Specify a string containing the metric group name. You can use an existing name or use a new name to create a new metric group. Currently, this parameter can have only one string.",
			//   "items": {
			//     "type": "string"
			//   },
			//   "maxItems": 1,
			//   "type": "array"
			// }
			Description: "The name of an Amazon CloudWatch metric group. A metric group aggregates the metrics for all fleets in the group. Specify a string containing the metric group name. You can use an existing name or use a new name to create a new metric group. Currently, this parameter can have only one string.",
			Type:        types.ListType{ElemType: types.StringType},
			Optional:    true,
		},
		"min_size": {
			// Property: MinSize
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "[DEPRECATED] The minimum value allowed for the fleet's instance count. When creating a new fleet, GameLift automatically sets this value to \"0\". After the fleet is active, you can change this value.",
			//   "type": "integer"
			// }
			Description: "[DEPRECATED] The minimum value allowed for the fleet's instance count. When creating a new fleet, GameLift automatically sets this value to \"0\". After the fleet is active, you can change this value.",
			Type:        types.NumberType,
			Optional:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A descriptive label that is associated with a fleet. Fleet names do not need to be unique.",
			//   "maxLength": 1024,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "A descriptive label that is associated with a fleet. Fleet names do not need to be unique.",
			Type:        types.StringType,
			Optional:    true,
		},
		"new_game_session_protection_policy": {
			// Property: NewGameSessionProtectionPolicy
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A game session protection policy to apply to all game sessions hosted on instances in this fleet. When protected, active game sessions cannot be terminated during a scale-down event. If this parameter is not set, instances in this fleet default to no protection. You can change a fleet's protection policy to affect future game sessions on the fleet. You can also set protection for individual game sessions.",
			//   "enum": [
			//     "FullProtection",
			//     "NoProtection"
			//   ],
			//   "type": "string"
			// }
			Description: "A game session protection policy to apply to all game sessions hosted on instances in this fleet. When protected, active game sessions cannot be terminated during a scale-down event. If this parameter is not set, instances in this fleet default to no protection. You can change a fleet's protection policy to affect future game sessions on the fleet. You can also set protection for individual game sessions.",
			Type:        types.StringType,
			Optional:    true,
		},
		"peer_vpc_aws_account_id": {
			// Property: PeerVpcAwsAccountId
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A unique identifier for the AWS account with the VPC that you want to peer your Amazon GameLift fleet with. You can find your account ID in the AWS Management Console under account settings.",
			//   "maxLength": 1024,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "A unique identifier for the AWS account with the VPC that you want to peer your Amazon GameLift fleet with. You can find your account ID in the AWS Management Console under account settings.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// PeerVpcAwsAccountId is a force-new attribute.
		},
		"peer_vpc_id": {
			// Property: PeerVpcId
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A unique identifier for a VPC with resources to be accessed by your Amazon GameLift fleet. The VPC must be in the same Region as your fleet. To look up a VPC ID, use the VPC Dashboard in the AWS Management Console.",
			//   "maxLength": 1024,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "A unique identifier for a VPC with resources to be accessed by your Amazon GameLift fleet. The VPC must be in the same Region as your fleet. To look up a VPC ID, use the VPC Dashboard in the AWS Management Console.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// PeerVpcId is a force-new attribute.
		},
		"resource_creation_limit_policy": {
			// Property: ResourceCreationLimitPolicy
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A policy that limits the number of game sessions a player can create on the same fleet. This optional policy gives game owners control over how players can consume available game server resources. A resource creation policy makes the following statement: \"An individual player can create a maximum number of new game sessions within a specified time period\".\n\nThe policy is evaluated when a player tries to create a new game session. For example, assume you have a policy of 10 new game sessions and a time period of 60 minutes. On receiving a CreateGameSession request, Amazon GameLift checks that the player (identified by CreatorId) has created fewer than 10 game sessions in the past 60 minutes.",
			//   "properties": {
			//     "NewGameSessionsPerCreator": {
			//       "additionalProperties": false,
			//       "description": "The maximum number of game sessions that an individual can create during the policy period.",
			//       "type": "integer"
			//     },
			//     "PolicyPeriodInMinutes": {
			//       "additionalProperties": false,
			//       "description": "The time span used in evaluating the resource creation limit policy.",
			//       "type": "integer"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "A policy that limits the number of game sessions a player can create on the same fleet. This optional policy gives game owners control over how players can consume available game server resources. A resource creation policy makes the following statement: \"An individual player can create a maximum number of new game sessions within a specified time period\".\n\nThe policy is evaluated when a player tries to create a new game session. For example, assume you have a policy of 10 new game sessions and a time period of 60 minutes. On receiving a CreateGameSession request, Amazon GameLift checks that the player (identified by CreatorId) has created fewer than 10 game sessions in the past 60 minutes.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"new_game_sessions_per_creator": {
						// Property: NewGameSessionsPerCreator
						Description: "The maximum number of game sessions that an individual can create during the policy period.",
						Type:        types.NumberType,
						Optional:    true,
					},
					"policy_period_in_minutes": {
						// Property: PolicyPeriodInMinutes
						Description: "The time span used in evaluating the resource creation limit policy.",
						Type:        types.NumberType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"runtime_configuration": {
			// Property: RuntimeConfiguration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A collection of server process configurations that describe the processes to run on each instance in a fleet. All fleets must have a runtime configuration. Each instance in the fleet maintains server processes as specified in the runtime configuration, launching new ones as existing processes end. Each instance regularly checks for an updated runtime configuration makes adjustments as called for.\n\nThe runtime configuration enables the instances in a fleet to run multiple processes simultaneously. Potential scenarios are as follows: (1) Run multiple processes of a single game server executable to maximize usage of your hosting resources. (2) Run one or more processes of different executables, such as your game server and a metrics tracking program. (3) Run multiple processes of a single game server but with different launch parameters, for example to run one process on each instance in debug mode.\n\nAn Amazon GameLift instance is limited to 50 processes running simultaneously. A runtime configuration must specify fewer than this limit. To calculate the total number of processes specified in a runtime configuration, add the values of the ConcurrentExecutions parameter for each ServerProcess object in the runtime configuration.",
			//   "properties": {
			//     "GameSessionActivationTimeoutSeconds": {
			//       "additionalProperties": false,
			//       "description": "The maximum amount of time (in seconds) that a game session can remain in status ACTIVATING. If the game session is not active before the timeout, activation is terminated and the game session status is changed to TERMINATED.",
			//       "type": "integer"
			//     },
			//     "MaxConcurrentGameSessionActivations": {
			//       "additionalProperties": false,
			//       "description": "The maximum number of game sessions with status ACTIVATING to allow on an instance simultaneously. This setting limits the amount of instance resources that can be used for new game activations at any one time.",
			//       "type": "integer"
			//     },
			//     "ServerProcesses": {
			//       "additionalProperties": false,
			//       "description": "A collection of server process configurations that describe which server processes to run on each instance in a fleet.",
			//       "items": {
			//         "additionalProperties": false,
			//         "description": "A set of instructions for launching server processes on each instance in a fleet. Each instruction set identifies the location of the server executable, optional launch parameters, and the number of server processes with this configuration to maintain concurrently on the instance. Server process configurations make up a fleet's RuntimeConfiguration.",
			//         "properties": {
			//           "ConcurrentExecutions": {
			//             "additionalProperties": false,
			//             "description": "The number of server processes that use this configuration to run concurrently on an instance.",
			//             "type": "integer"
			//           },
			//           "LaunchPath": {
			//             "additionalProperties": false,
			//             "description": "The location of the server executable in a custom game build or the name of the Realtime script file that contains the Init() function. Game builds and Realtime scripts are installed on instances at the root:\n\nWindows (for custom game builds only): C:\\game. Example: \"C:\\game\\MyGame\\server.exe\"\n\nLinux: /local/game. Examples: \"/local/game/MyGame/server.exe\" or \"/local/game/MyRealtimeScript.js\"",
			//             "maxLength": 1024,
			//             "minLength": 1,
			//             "pattern": "",
			//             "type": "string"
			//           },
			//           "Parameters": {
			//             "additionalProperties": false,
			//             "description": "An optional list of parameters to pass to the server executable or Realtime script on launch.",
			//             "maxLength": 1024,
			//             "minLength": 1,
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "ConcurrentExecutions",
			//           "LaunchPath"
			//         ],
			//         "type": "object"
			//       },
			//       "maxItems": 50,
			//       "type": "array"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "A collection of server process configurations that describe the processes to run on each instance in a fleet. All fleets must have a runtime configuration. Each instance in the fleet maintains server processes as specified in the runtime configuration, launching new ones as existing processes end. Each instance regularly checks for an updated runtime configuration makes adjustments as called for.\n\nThe runtime configuration enables the instances in a fleet to run multiple processes simultaneously. Potential scenarios are as follows: (1) Run multiple processes of a single game server executable to maximize usage of your hosting resources. (2) Run one or more processes of different executables, such as your game server and a metrics tracking program. (3) Run multiple processes of a single game server but with different launch parameters, for example to run one process on each instance in debug mode.\n\nAn Amazon GameLift instance is limited to 50 processes running simultaneously. A runtime configuration must specify fewer than this limit. To calculate the total number of processes specified in a runtime configuration, add the values of the ConcurrentExecutions parameter for each ServerProcess object in the runtime configuration.",
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"game_session_activation_timeout_seconds": {
						// Property: GameSessionActivationTimeoutSeconds
						Description: "The maximum amount of time (in seconds) that a game session can remain in status ACTIVATING. If the game session is not active before the timeout, activation is terminated and the game session status is changed to TERMINATED.",
						Type:        types.NumberType,
						Optional:    true,
					},
					"max_concurrent_game_session_activations": {
						// Property: MaxConcurrentGameSessionActivations
						Description: "The maximum number of game sessions with status ACTIVATING to allow on an instance simultaneously. This setting limits the amount of instance resources that can be used for new game activations at any one time.",
						Type:        types.NumberType,
						Optional:    true,
					},
					"server_processes": {
						// Property: ServerProcesses
						Description: "A collection of server process configurations that describe which server processes to run on each instance in a fleet.",
						Attributes: schema.ListNestedAttributes(
							map[string]schema.Attribute{
								"concurrent_executions": {
									// Property: ConcurrentExecutions
									Description: "The number of server processes that use this configuration to run concurrently on an instance.",
									Type:        types.NumberType,
									Required:    true,
								},
								"launch_path": {
									// Property: LaunchPath
									Description: "The location of the server executable in a custom game build or the name of the Realtime script file that contains the Init() function. Game builds and Realtime scripts are installed on instances at the root:\n\nWindows (for custom game builds only): C:\\game. Example: \"C:\\game\\MyGame\\server.exe\"\n\nLinux: /local/game. Examples: \"/local/game/MyGame/server.exe\" or \"/local/game/MyRealtimeScript.js\"",
									Type:        types.StringType,
									Required:    true,
								},
								"parameters": {
									// Property: Parameters
									Description: "An optional list of parameters to pass to the server executable or Realtime script on launch.",
									Type:        types.StringType,
									Optional:    true,
								},
							},
							schema.ListNestedAttributesOptions{
								MaxItems: 50,
							},
						),
						Optional: true,
					},
				},
			),
			Optional: true,
		},
		"script_id": {
			// Property: ScriptId
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A unique identifier for a Realtime script to be deployed on a new Realtime Servers fleet. The script must have been successfully uploaded to Amazon GameLift. This fleet setting cannot be changed once the fleet is created.\n\nNote: It is not currently possible to use the !Ref command to reference a script created with a CloudFormation template for the fleet property ScriptId. Instead, use Fn::GetAtt Script.Arn or Fn::GetAtt Script.Id to retrieve either of these properties as input for ScriptId. Alternatively, enter a ScriptId string manually.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "A unique identifier for a Realtime script to be deployed on a new Realtime Servers fleet. The script must have been successfully uploaded to Amazon GameLift. This fleet setting cannot be changed once the fleet is created.\n\nNote: It is not currently possible to use the !Ref command to reference a script created with a CloudFormation template for the fleet property ScriptId. Instead, use Fn::GetAtt Script.Arn or Fn::GetAtt Script.Id to retrieve either of these properties as input for ScriptId. Alternatively, enter a ScriptId string manually.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// ScriptId is a force-new attribute.
		},
		"server_launch_parameters": {
			// Property: ServerLaunchParameters
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "This parameter is no longer used but is retained for backward compatibility. Instead, specify server launch parameters in the RuntimeConfiguration parameter. A request must specify either a runtime configuration or values for both ServerLaunchParameters and ServerLaunchPath.",
			//   "maxLength": 1024,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "This parameter is no longer used but is retained for backward compatibility. Instead, specify server launch parameters in the RuntimeConfiguration parameter. A request must specify either a runtime configuration or values for both ServerLaunchParameters and ServerLaunchPath.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// ServerLaunchParameters is a force-new attribute.
		},
		"server_launch_path": {
			// Property: ServerLaunchPath
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "This parameter is no longer used. Instead, specify a server launch path using the RuntimeConfiguration parameter. Requests that specify a server launch path and launch parameters instead of a runtime configuration will continue to work.",
			//   "maxLength": 1024,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "This parameter is no longer used. Instead, specify a server launch path using the RuntimeConfiguration parameter. Requests that specify a server launch path and launch parameters instead of a runtime configuration will continue to work.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// ServerLaunchPath is a force-new attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "The AWS::GameLift::Fleet resource creates an Amazon GameLift (GameLift) fleet to host game servers.  A fleet is a set of EC2 instances, each of which can host multiple game sessions.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::GameLift::Fleet").WithTerraformTypeName("aws_gamelift_fleet").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_gamelift_fleet", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
