// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package datapipeline

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_datapipeline_pipeline", pipelineDataSource)
}

// pipelineDataSource returns the Terraform awscc_datapipeline_pipeline data source.
// This Terraform data source corresponds to the CloudFormation AWS::DataPipeline::Pipeline resource.
func pipelineDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"activate": {
			// Property: Activate
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Indicates whether to validate and start the pipeline or stop an active pipeline. By default, the value is set to true.",
			//	  "type": "boolean"
			//	}
			Description: "Indicates whether to validate and start the pipeline or stop an active pipeline. By default, the value is set to true.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "A description of the pipeline.",
			//	  "type": "string"
			//	}
			Description: "A description of the pipeline.",
			Type:        types.StringType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of the pipeline.",
			//	  "type": "string"
			//	}
			Description: "The name of the pipeline.",
			Type:        types.StringType,
			Computed:    true,
		},
		"parameter_objects": {
			// Property: ParameterObjects
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The parameter objects used with the pipeline.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "properties": {
			//	      "Attributes": {
			//	        "description": "The attributes of the parameter object.",
			//	        "insertionOrder": false,
			//	        "items": {
			//	          "additionalProperties": false,
			//	          "properties": {
			//	            "Key": {
			//	              "description": "The field identifier.",
			//	              "type": "string"
			//	            },
			//	            "StringValue": {
			//	              "description": "The field value, expressed as a String.",
			//	              "type": "string"
			//	            }
			//	          },
			//	          "required": [
			//	            "Key",
			//	            "StringValue"
			//	          ],
			//	          "type": "object"
			//	        },
			//	        "type": "array",
			//	        "uniqueItems": false
			//	      },
			//	      "Id": {
			//	        "description": "The ID of the parameter object.",
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "Attributes",
			//	      "Id"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "type": "array",
			//	  "uniqueItems": false
			//	}
			Description: "The parameter objects used with the pipeline.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"attributes": {
						// Property: Attributes
						Description: "The attributes of the parameter object.",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"key": {
									// Property: Key
									Description: "The field identifier.",
									Type:        types.StringType,
									Computed:    true,
								},
								"string_value": {
									// Property: StringValue
									Description: "The field value, expressed as a String.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"id": {
						// Property: Id
						Description: "The ID of the parameter object.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"parameter_values": {
			// Property: ParameterValues
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The parameter values used with the pipeline.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "properties": {
			//	      "Id": {
			//	        "description": "The ID of the parameter value.",
			//	        "type": "string"
			//	      },
			//	      "StringValue": {
			//	        "description": "The field value, expressed as a String.",
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "Id",
			//	      "StringValue"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "type": "array",
			//	  "uniqueItems": false
			//	}
			Description: "The parameter values used with the pipeline.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"id": {
						// Property: Id
						Description: "The ID of the parameter value.",
						Type:        types.StringType,
						Computed:    true,
					},
					"string_value": {
						// Property: StringValue
						Description: "The field value, expressed as a String.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"pipeline_id": {
			// Property: PipelineId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"pipeline_objects": {
			// Property: PipelineObjects
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The objects that define the pipeline. These objects overwrite the existing pipeline definition. Not all objects, fields, and values can be updated. For information about restrictions, see Editing Your Pipeline in the AWS Data Pipeline Developer Guide.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "properties": {
			//	      "Fields": {
			//	        "description": "Key-value pairs that define the properties of the object.",
			//	        "insertionOrder": false,
			//	        "items": {
			//	          "additionalProperties": false,
			//	          "properties": {
			//	            "Key": {
			//	              "description": "Specifies the name of a field for a particular object. To view valid values for a particular field, see Pipeline Object Reference in the AWS Data Pipeline Developer Guide.",
			//	              "type": "string"
			//	            },
			//	            "RefValue": {
			//	              "description": "A field value that you specify as an identifier of another object in the same pipeline definition.",
			//	              "type": "string"
			//	            },
			//	            "StringValue": {
			//	              "description": "A field value that you specify as a string. To view valid values for a particular field, see Pipeline Object Reference in the AWS Data Pipeline Developer Guide.",
			//	              "type": "string"
			//	            }
			//	          },
			//	          "required": [
			//	            "Key"
			//	          ],
			//	          "type": "object"
			//	        },
			//	        "type": "array",
			//	        "uniqueItems": false
			//	      },
			//	      "Id": {
			//	        "description": "The ID of the object.",
			//	        "type": "string"
			//	      },
			//	      "Name": {
			//	        "description": "The name of the object.",
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "Fields",
			//	      "Id",
			//	      "Name"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "type": "array",
			//	  "uniqueItems": false
			//	}
			Description: "The objects that define the pipeline. These objects overwrite the existing pipeline definition. Not all objects, fields, and values can be updated. For information about restrictions, see Editing Your Pipeline in the AWS Data Pipeline Developer Guide.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"fields": {
						// Property: Fields
						Description: "Key-value pairs that define the properties of the object.",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"key": {
									// Property: Key
									Description: "Specifies the name of a field for a particular object. To view valid values for a particular field, see Pipeline Object Reference in the AWS Data Pipeline Developer Guide.",
									Type:        types.StringType,
									Computed:    true,
								},
								"ref_value": {
									// Property: RefValue
									Description: "A field value that you specify as an identifier of another object in the same pipeline definition.",
									Type:        types.StringType,
									Computed:    true,
								},
								"string_value": {
									// Property: StringValue
									Description: "A field value that you specify as a string. To view valid values for a particular field, see Pipeline Object Reference in the AWS Data Pipeline Developer Guide.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"id": {
						// Property: Id
						Description: "The ID of the object.",
						Type:        types.StringType,
						Computed:    true,
					},
					"name": {
						// Property: Name
						Description: "The name of the object.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"pipeline_tags": {
			// Property: PipelineTags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "A list of arbitrary tags (key-value pairs) to associate with the pipeline, which you can use to control permissions. For more information, see Controlling Access to Pipelines and Resources in the AWS Data Pipeline Developer Guide.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "properties": {
			//	      "Key": {
			//	        "description": "The key name of a tag.",
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "description": "The value to associate with the key name.",
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "Key",
			//	      "Value"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "type": "array",
			//	  "uniqueItems": false
			//	}
			Description: "A list of arbitrary tags (key-value pairs) to associate with the pipeline, which you can use to control permissions. For more information, see Controlling Access to Pipelines and Resources in the AWS Data Pipeline Developer Guide.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of a tag.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value to associate with the key name.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::DataPipeline::Pipeline",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataPipeline::Pipeline").WithTerraformTypeName("awscc_datapipeline_pipeline")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"activate":          "Activate",
		"attributes":        "Attributes",
		"description":       "Description",
		"fields":            "Fields",
		"id":                "Id",
		"key":               "Key",
		"name":              "Name",
		"parameter_objects": "ParameterObjects",
		"parameter_values":  "ParameterValues",
		"pipeline_id":       "PipelineId",
		"pipeline_objects":  "PipelineObjects",
		"pipeline_tags":     "PipelineTags",
		"ref_value":         "RefValue",
		"string_value":      "StringValue",
		"value":             "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}