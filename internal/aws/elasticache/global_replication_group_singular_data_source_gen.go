// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_elasticache_global_replication_group", globalReplicationGroupDataSource)
}

// globalReplicationGroupDataSource returns the Terraform awscc_elasticache_global_replication_group data source.
// This Terraform data source corresponds to the CloudFormation AWS::ElastiCache::GlobalReplicationGroup resource.
func globalReplicationGroupDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AutomaticFailoverEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "AutomaticFailoverEnabled",
		//	  "type": "boolean"
		//	}
		"automatic_failover_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "AutomaticFailoverEnabled",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CacheNodeType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The cache node type of the Global Datastore",
		//	  "type": "string"
		//	}
		"cache_node_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The cache node type of the Global Datastore",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CacheParameterGroupName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Cache parameter group name to use for the new engine version. This parameter cannot be modified independently.",
		//	  "type": "string"
		//	}
		"cache_parameter_group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Cache parameter group name to use for the new engine version. This parameter cannot be modified independently.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Engine
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The engine of the Global Datastore.",
		//	  "type": "string"
		//	}
		"engine": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The engine of the Global Datastore.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EngineVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The engine version of the Global Datastore.",
		//	  "type": "string"
		//	}
		"engine_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The engine version of the Global Datastore.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: GlobalNodeGroupCount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates the number of node groups in the Global Datastore.",
		//	  "type": "integer"
		//	}
		"global_node_group_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Indicates the number of node groups in the Global Datastore.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: GlobalReplicationGroupDescription
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The optional description of the Global Datastore",
		//	  "type": "string"
		//	}
		"global_replication_group_description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The optional description of the Global Datastore",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: GlobalReplicationGroupId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the Global Datastore, it is generated by ElastiCache adding a prefix to GlobalReplicationGroupIdSuffix.",
		//	  "type": "string"
		//	}
		"global_replication_group_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the Global Datastore, it is generated by ElastiCache adding a prefix to GlobalReplicationGroupIdSuffix.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: GlobalReplicationGroupIdSuffix
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The suffix name of a Global Datastore. Amazon ElastiCache automatically applies a prefix to the Global Datastore ID when it is created. Each AWS Region has its own prefix. ",
		//	  "type": "string"
		//	}
		"global_replication_group_id_suffix": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The suffix name of a Global Datastore. Amazon ElastiCache automatically applies a prefix to the Global Datastore ID when it is created. Each AWS Region has its own prefix. ",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Members
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The replication groups that comprise the Global Datastore.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "ReplicationGroupId": {
		//	        "description": "Regionally unique identifier for the member i.e. ReplicationGroupId.",
		//	        "type": "string"
		//	      },
		//	      "ReplicationGroupRegion": {
		//	        "description": "The AWS region of the Global Datastore member.",
		//	        "type": "string"
		//	      },
		//	      "Role": {
		//	        "description": "Indicates the role of the member, primary or secondary.",
		//	        "enum": [
		//	          "PRIMARY",
		//	          "SECONDARY"
		//	        ],
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "minItems": 1,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"members": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: ReplicationGroupId
					"replication_group_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Regionally unique identifier for the member i.e. ReplicationGroupId.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: ReplicationGroupRegion
					"replication_group_region": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The AWS region of the Global Datastore member.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Role
					"role": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Indicates the role of the member, primary or secondary.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The replication groups that comprise the Global Datastore.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RegionalConfigurations
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Describes the replication group IDs, the AWS regions where they are stored and the shard configuration for each that comprise the Global Datastore ",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "ReplicationGroupId": {
		//	        "description": "The replication group id of the Global Datastore member.",
		//	        "type": "string"
		//	      },
		//	      "ReplicationGroupRegion": {
		//	        "description": "The AWS region of the Global Datastore member.",
		//	        "type": "string"
		//	      },
		//	      "ReshardingConfigurations": {
		//	        "description": "A list of PreferredAvailabilityZones objects that specifies the configuration of a node group in the resharded cluster. ",
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "NodeGroupId": {
		//	              "description": "Unique identifier for the Node Group. This is either auto-generated by ElastiCache (4-digit id) or a user supplied id.",
		//	              "type": "string"
		//	            },
		//	            "PreferredAvailabilityZones": {
		//	              "description": "A list of preferred availability zones for the nodes of new node groups.",
		//	              "items": {
		//	                "type": "string"
		//	              },
		//	              "type": "array",
		//	              "uniqueItems": false
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "type": "array",
		//	        "uniqueItems": true
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"regional_configurations": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: ReplicationGroupId
					"replication_group_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The replication group id of the Global Datastore member.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: ReplicationGroupRegion
					"replication_group_region": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The AWS region of the Global Datastore member.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: ReshardingConfigurations
					"resharding_configurations": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: NodeGroupId
								"node_group_id": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "Unique identifier for the Node Group. This is either auto-generated by ElastiCache (4-digit id) or a user supplied id.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: PreferredAvailabilityZones
								"preferred_availability_zones": schema.ListAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Description: "A list of preferred availability zones for the nodes of new node groups.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Description: "A list of PreferredAvailabilityZones objects that specifies the configuration of a node group in the resharded cluster. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Describes the replication group IDs, the AWS regions where they are stored and the shard configuration for each that comprise the Global Datastore ",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The status of the Global Datastore",
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The status of the Global Datastore",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::ElastiCache::GlobalReplicationGroup",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ElastiCache::GlobalReplicationGroup").WithTerraformTypeName("awscc_elasticache_global_replication_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"automatic_failover_enabled":           "AutomaticFailoverEnabled",
		"cache_node_type":                      "CacheNodeType",
		"cache_parameter_group_name":           "CacheParameterGroupName",
		"engine":                               "Engine",
		"engine_version":                       "EngineVersion",
		"global_node_group_count":              "GlobalNodeGroupCount",
		"global_replication_group_description": "GlobalReplicationGroupDescription",
		"global_replication_group_id":          "GlobalReplicationGroupId",
		"global_replication_group_id_suffix":   "GlobalReplicationGroupIdSuffix",
		"members":                              "Members",
		"node_group_id":                        "NodeGroupId",
		"preferred_availability_zones":         "PreferredAvailabilityZones",
		"regional_configurations":              "RegionalConfigurations",
		"replication_group_id":                 "ReplicationGroupId",
		"replication_group_region":             "ReplicationGroupRegion",
		"resharding_configurations":            "ReshardingConfigurations",
		"role":                                 "Role",
		"status":                               "Status",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
