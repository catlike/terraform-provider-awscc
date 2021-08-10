// Code generated by generators/resource/main.go; DO NOT EDIT.

package sagemaker

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
	registry.AddResourceTypeFactory("aws_sagemaker_feature_group", featureGroupResourceType)
}

// featureGroupResourceType returns the Terraform aws_sagemaker_feature_group resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::SageMaker::FeatureGroup resource type.
func featureGroupResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Description about the FeatureGroup.",
			     "maxLength": 128,
			     "type": "string"
			   }
			*/
			Description: "Description about the FeatureGroup.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// Description is a force-new attribute.
		},
		"event_time_feature_name": {
			// Property: EventTimeFeatureName
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The Event Time Feature Name.",
			     "maxLength": 64,
			     "minLength": 1,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The Event Time Feature Name.",
			Type:        types.StringType,
			Required:    true,
			// EventTimeFeatureName is a force-new attribute.
		},
		"feature_definitions": {
			// Property: FeatureDefinitions
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "An Array of Feature Definition",
			     "items": {
			       "additionalProperties": false,
			       "properties": {
			         "FeatureName": {
			           "maxLength": 64,
			           "minLength": 1,
			           "pattern": "",
			           "type": "string"
			         },
			         "FeatureType": {
			           "enum": [
			             "Integral",
			             "Fractional",
			             "String"
			           ],
			           "type": "string"
			         }
			       },
			       "$ref": "#/definitions/FeatureDefinition",
			       "required": [
			         "FeatureName",
			         "FeatureType"
			       ],
			       "type": "object"
			     },
			     "maxItems": 2500,
			     "minItems": 1,
			     "type": "array",
			     "uniqueItems": false
			   }
			*/
			Description: "An Array of Feature Definition",
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"feature_name": {
						// Property: FeatureName
						// CloudFormation resource type schema:
						/*
						   {
						     "maxLength": 64,
						     "minLength": 1,
						     "pattern": "",
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Required: true,
					},
					"feature_type": {
						// Property: FeatureType
						// CloudFormation resource type schema:
						/*
						   {
						     "enum": [
						       "Integral",
						       "Fractional",
						       "String"
						     ],
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Required: true,
					},
				},
				schema.ListNestedAttributesOptions{
					MinItems: 1,
					MaxItems: 2500,
				},
			),
			Required: true,
			// FeatureDefinitions is a force-new attribute.
		},
		"feature_group_name": {
			// Property: FeatureGroupName
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The Name of the FeatureGroup.",
			     "maxLength": 64,
			     "minLength": 1,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The Name of the FeatureGroup.",
			Type:        types.StringType,
			Required:    true,
			// FeatureGroupName is a force-new attribute.
		},
		"offline_store_config": {
			// Property: OfflineStoreConfig
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "properties": {
			       "DataCatalogConfig": {
			         "additionalProperties": false,
			         "properties": {
			           "Catalog": {
			             "maxLength": 255,
			             "minLength": 1,
			             "pattern": "",
			             "type": "string"
			           },
			           "Database": {
			             "maxLength": 255,
			             "minLength": 1,
			             "pattern": "",
			             "type": "string"
			           },
			           "TableName": {
			             "maxLength": 255,
			             "minLength": 1,
			             "pattern": "",
			             "type": "string"
			           }
			         },
			         "$ref": "#/definitions/DataCatalogConfig",
			         "required": [
			           "TableName",
			           "Catalog",
			           "Database"
			         ],
			         "type": "object"
			       },
			       "DisableGlueTableCreation": {
			         "type": "boolean"
			       },
			       "S3StorageConfig": {
			         "additionalProperties": false,
			         "properties": {
			           "KmsKeyId": {
			             "maxLength": 2048,
			             "$ref": "#/definitions/KmsKeyId",
			             "type": "string"
			           },
			           "S3Uri": {
			             "maxLength": 1024,
			             "pattern": "",
			             "type": "string"
			           }
			         },
			         "$ref": "#/definitions/S3StorageConfig",
			         "required": [
			           "S3Uri"
			         ],
			         "type": "object"
			       }
			     },
			     "required": [
			       "S3StorageConfig"
			     ],
			     "type": "object"
			   }
			*/
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"data_catalog_config": {
						// Property: DataCatalogConfig
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "properties": {
						       "Catalog": {
						         "maxLength": 255,
						         "minLength": 1,
						         "pattern": "",
						         "type": "string"
						       },
						       "Database": {
						         "maxLength": 255,
						         "minLength": 1,
						         "pattern": "",
						         "type": "string"
						       },
						       "TableName": {
						         "maxLength": 255,
						         "minLength": 1,
						         "pattern": "",
						         "type": "string"
						       }
						     },
						     "$ref": "#/definitions/DataCatalogConfig",
						     "required": [
						       "TableName",
						       "Catalog",
						       "Database"
						     ],
						     "type": "object"
						   }
						*/
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"catalog": {
									// Property: Catalog
									// CloudFormation resource type schema:
									/*
									   {
									     "maxLength": 255,
									     "minLength": 1,
									     "pattern": "",
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Required: true,
								},
								"database": {
									// Property: Database
									// CloudFormation resource type schema:
									/*
									   {
									     "maxLength": 255,
									     "minLength": 1,
									     "pattern": "",
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Required: true,
								},
								"table_name": {
									// Property: TableName
									// CloudFormation resource type schema:
									/*
									   {
									     "maxLength": 255,
									     "minLength": 1,
									     "pattern": "",
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Required: true,
								},
							},
						),
						Optional: true,
					},
					"disable_glue_table_creation": {
						// Property: DisableGlueTableCreation
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "boolean"
						   }
						*/
						Type:     types.BoolType,
						Optional: true,
					},
					"s3_storage_config": {
						// Property: S3StorageConfig
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "properties": {
						       "KmsKeyId": {
						         "maxLength": 2048,
						         "$ref": "#/definitions/KmsKeyId",
						         "type": "string"
						       },
						       "S3Uri": {
						         "maxLength": 1024,
						         "pattern": "",
						         "type": "string"
						       }
						     },
						     "$ref": "#/definitions/S3StorageConfig",
						     "required": [
						       "S3Uri"
						     ],
						     "type": "object"
						   }
						*/
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"kms_key_id": {
									// Property: KmsKeyId
									// CloudFormation resource type schema:
									/*
									   {
									     "maxLength": 2048,
									     "$ref": "#/definitions/KmsKeyId",
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Optional: true,
								},
								"s3_uri": {
									// Property: S3Uri
									// CloudFormation resource type schema:
									/*
									   {
									     "maxLength": 1024,
									     "pattern": "",
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Required: true,
								},
							},
						),
						Required: true,
					},
				},
			),
			Optional: true,
			Computed: true,
			// OfflineStoreConfig is a force-new attribute.
		},
		"online_store_config": {
			// Property: OnlineStoreConfig
			// CloudFormation resource type schema:
			/*
			   {
			     "additionalProperties": false,
			     "properties": {
			       "EnableOnlineStore": {
			         "type": "boolean"
			       },
			       "SecurityConfig": {
			         "additionalProperties": false,
			         "properties": {
			           "KmsKeyId": {
			             "maxLength": 2048,
			             "$ref": "#/definitions/KmsKeyId",
			             "type": "string"
			           }
			         },
			         "$ref": "#/definitions/OnlineStoreSecurityConfig",
			         "type": "object"
			       }
			     },
			     "type": "object"
			   }
			*/
			Attributes: schema.SingleNestedAttributes(
				map[string]schema.Attribute{
					"enable_online_store": {
						// Property: EnableOnlineStore
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "boolean"
						   }
						*/
						Type:     types.BoolType,
						Optional: true,
					},
					"security_config": {
						// Property: SecurityConfig
						// CloudFormation resource type schema:
						/*
						   {
						     "additionalProperties": false,
						     "properties": {
						       "KmsKeyId": {
						         "maxLength": 2048,
						         "$ref": "#/definitions/KmsKeyId",
						         "type": "string"
						       }
						     },
						     "$ref": "#/definitions/OnlineStoreSecurityConfig",
						     "type": "object"
						   }
						*/
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"kms_key_id": {
									// Property: KmsKeyId
									// CloudFormation resource type schema:
									/*
									   {
									     "maxLength": 2048,
									     "$ref": "#/definitions/KmsKeyId",
									     "type": "string"
									   }
									*/
									Type:     types.StringType,
									Optional: true,
								},
							},
						),
						Optional: true,
					},
				},
			),
			Optional: true,
			Computed: true,
			// OnlineStoreConfig is a force-new attribute.
		},
		"record_identifier_feature_name": {
			// Property: RecordIdentifierFeatureName
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "The Record Identifier Feature Name.",
			     "maxLength": 64,
			     "minLength": 1,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "The Record Identifier Feature Name.",
			Type:        types.StringType,
			Required:    true,
			// RecordIdentifierFeatureName is a force-new attribute.
		},
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "Role Arn",
			     "maxLength": 2048,
			     "minLength": 20,
			     "pattern": "",
			     "type": "string"
			   }
			*/
			Description: "Role Arn",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// RoleArn is a force-new attribute.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			/*
			   {
			     "description": "An array of key-value pair to apply to this resource.",
			     "items": {
			       "additionalProperties": false,
			       "description": "A key-value pair to associate with a resource.",
			       "properties": {
			         "Key": {
			           "type": "string"
			         },
			         "Value": {
			           "type": "string"
			         }
			       },
			       "$ref": "#/definitions/Tag",
			       "required": [
			         "Value",
			         "Key"
			       ],
			       "type": "object"
			     },
			     "maxItems": 50,
			     "type": "array",
			     "uniqueItems": false
			   }
			*/
			Description: "An array of key-value pair to apply to this resource.",
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"key": {
						// Property: Key
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						// CloudFormation resource type schema:
						/*
						   {
						     "type": "string"
						   }
						*/
						Type:     types.StringType,
						Required: true,
					},
				},
				schema.ListNestedAttributesOptions{
					MaxItems: 50,
				},
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
		Description: "Resource Type definition for AWS::SageMaker::FeatureGroup",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SageMaker::FeatureGroup").WithTerraformTypeName("aws_sagemaker_feature_group").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_sagemaker_feature_group", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}