// Code generated by generators/resource/main.go; DO NOT EDIT.

package quicksight

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_quicksight_template", templateResourceType)
}

// templateResourceType returns the Terraform awscc_quicksight_template resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::QuickSight::Template resource type.
func templateResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the template.\u003c/p\u003e",
			//   "type": "string"
			// }
			Description: "<p>The Amazon Resource Name (ARN) of the template.</p>",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"aws_account_id": {
			// Property: AwsAccountId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 12,
			//   "minLength": 12,
			//   "pattern": "^[0-9]{12}$",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(12, 12),
				validate.StringMatch(regexp.MustCompile("^[0-9]{12}$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"created_time": {
			// Property: CreatedTime
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eTime when this was created.\u003c/p\u003e",
			//   "format": "string",
			//   "type": "string"
			// }
			Description: "<p>Time when this was created.</p>",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
			// CreatedTime is a write-only property.
		},
		"last_updated_time": {
			// Property: LastUpdatedTime
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eTime when this was last updated.\u003c/p\u003e",
			//   "format": "string",
			//   "type": "string"
			// }
			Description: "<p>Time when this was last updated.</p>",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
			// LastUpdatedTime is a write-only property.
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eA display name for the template.\u003c/p\u003e",
			//   "maxLength": 2048,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "<p>A display name for the template.</p>",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 2048),
			},
		},
		"permissions": {
			// Property: Permissions
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eA list of resource permissions to be set on the template. \u003c/p\u003e",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "\u003cp\u003ePermission for the resource.\u003c/p\u003e",
			//     "properties": {
			//       "Actions": {
			//         "description": "\u003cp\u003eThe IAM action to grant or revoke permissions on.\u003c/p\u003e",
			//         "items": {
			//           "type": "string"
			//         },
			//         "maxItems": 16,
			//         "minItems": 1,
			//         "type": "array"
			//       },
			//       "Principal": {
			//         "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the principal. This can be one of the\n            following:\u003c/p\u003e\n        \u003cul\u003e\n            \u003cli\u003e\n                \u003cp\u003eThe ARN of an Amazon QuickSight user or group associated with a data source or dataset. (This is common.)\u003c/p\u003e\n            \u003c/li\u003e\n            \u003cli\u003e\n                \u003cp\u003eThe ARN of an Amazon QuickSight user, group, or namespace associated with an analysis, dashboard, template, or theme. (This is common.)\u003c/p\u003e\n            \u003c/li\u003e\n            \u003cli\u003e\n                \u003cp\u003eThe ARN of an AWS account root: This is an IAM ARN rather than a QuickSight\n                    ARN. Use this option only to share resources (templates) across AWS accounts.\n                    (This is less common.) \u003c/p\u003e\n            \u003c/li\u003e\n         \u003c/ul\u003e",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Actions",
			//       "Principal"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 64,
			//   "minItems": 1,
			//   "type": "array"
			// }
			Description: "<p>A list of resource permissions to be set on the template. </p>",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"actions": {
						// Property: Actions
						Description: "<p>The IAM action to grant or revoke permissions on.</p>",
						Type:        types.ListType{ElemType: types.StringType},
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenBetween(1, 16),
						},
					},
					"principal": {
						// Property: Principal
						Description: "<p>The Amazon Resource Name (ARN) of the principal. This can be one of the\n            following:</p>\n        <ul>\n            <li>\n                <p>The ARN of an Amazon QuickSight user or group associated with a data source or dataset. (This is common.)</p>\n            </li>\n            <li>\n                <p>The ARN of an Amazon QuickSight user, group, or namespace associated with an analysis, dashboard, template, or theme. (This is common.)</p>\n            </li>\n            <li>\n                <p>The ARN of an AWS account root: This is an IAM ARN rather than a QuickSight\n                    ARN. Use this option only to share resources (templates) across AWS accounts.\n                    (This is less common.) </p>\n            </li>\n         </ul>",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 256),
						},
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenBetween(1, 64),
			},
		},
		"source_entity": {
			// Property: SourceEntity
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "\u003cp\u003eThe source entity of the template.\u003c/p\u003e",
			//   "properties": {
			//     "SourceAnalysis": {
			//       "additionalProperties": false,
			//       "description": "\u003cp\u003eThe source analysis of the template.\u003c/p\u003e",
			//       "properties": {
			//         "Arn": {
			//           "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the resource.\u003c/p\u003e",
			//           "type": "string"
			//         },
			//         "DataSetReferences": {
			//           "description": "\u003cp\u003eA structure containing information about the dataset references used as placeholders\n            in the template.\u003c/p\u003e",
			//           "items": {
			//             "additionalProperties": false,
			//             "description": "\u003cp\u003eDataset reference.\u003c/p\u003e",
			//             "properties": {
			//               "DataSetArn": {
			//                 "description": "\u003cp\u003eDataset Amazon Resource Name (ARN).\u003c/p\u003e",
			//                 "type": "string"
			//               },
			//               "DataSetPlaceholder": {
			//                 "description": "\u003cp\u003eDataset placeholder.\u003c/p\u003e",
			//                 "pattern": ".*\\S.*",
			//                 "type": "string"
			//               }
			//             },
			//             "required": [
			//               "DataSetArn",
			//               "DataSetPlaceholder"
			//             ],
			//             "type": "object"
			//           },
			//           "minItems": 1,
			//           "type": "array"
			//         }
			//       },
			//       "required": [
			//         "Arn",
			//         "DataSetReferences"
			//       ],
			//       "type": "object"
			//     },
			//     "SourceTemplate": {
			//       "additionalProperties": false,
			//       "description": "\u003cp\u003eThe source template of the template.\u003c/p\u003e",
			//       "properties": {
			//         "Arn": {
			//           "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the resource.\u003c/p\u003e",
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "Arn"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "<p>The source entity of the template.</p>",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"source_analysis": {
						// Property: SourceAnalysis
						Description: "<p>The source analysis of the template.</p>",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"arn": {
									// Property: Arn
									Description: "<p>The Amazon Resource Name (ARN) of the resource.</p>",
									Type:        types.StringType,
									Required:    true,
								},
								"data_set_references": {
									// Property: DataSetReferences
									Description: "<p>A structure containing information about the dataset references used as placeholders\n            in the template.</p>",
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"data_set_arn": {
												// Property: DataSetArn
												Description: "<p>Dataset Amazon Resource Name (ARN).</p>",
												Type:        types.StringType,
												Required:    true,
											},
											"data_set_placeholder": {
												// Property: DataSetPlaceholder
												Description: "<p>Dataset placeholder.</p>",
												Type:        types.StringType,
												Required:    true,
												Validators: []tfsdk.AttributeValidator{
													validate.StringMatch(regexp.MustCompile(".*\\S.*"), ""),
												},
											},
										},
										tfsdk.ListNestedAttributesOptions{},
									),
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.ArrayLenAtLeast(1),
									},
								},
							},
						),
						Optional: true,
					},
					"source_template": {
						// Property: SourceTemplate
						Description: "<p>The source template of the template.</p>",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"arn": {
									// Property: Arn
									Description: "<p>The Amazon Resource Name (ARN) of the resource.</p>",
									Type:        types.StringType,
									Required:    true,
								},
							},
						),
						Optional: true,
					},
				},
			),
			Required: true,
			// SourceEntity is a write-only property.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eContains a map of the key-value pairs for the resource tag or tags assigned to the resource.\u003c/p\u003e",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "\u003cp\u003eThe key or keys of the key-value pairs for the resource tag or tags assigned to the\n            resource.\u003c/p\u003e",
			//     "properties": {
			//       "Key": {
			//         "description": "\u003cp\u003eTag key.\u003c/p\u003e",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "\u003cp\u003eTag value.\u003c/p\u003e",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 200,
			//   "minItems": 1,
			//   "type": "array"
			// }
			Description: "<p>Contains a map of the key-value pairs for the resource tag or tags assigned to the resource.</p>",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "<p>Tag key.</p>",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "<p>Tag value.</p>",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 256),
						},
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenBetween(1, 200),
			},
		},
		"template_id": {
			// Property: TemplateId
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 2048,
			//   "minLength": 1,
			//   "pattern": "[\\w\\-]+",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 2048),
				validate.StringMatch(regexp.MustCompile("[\\w\\-]+"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"version": {
			// Property: Version
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "\u003cp\u003eA version of a template.\u003c/p\u003e",
			//   "properties": {
			//     "CreatedTime": {
			//       "description": "\u003cp\u003eThe time that this template version was created.\u003c/p\u003e",
			//       "format": "string",
			//       "type": "string"
			//     },
			//     "DataSetConfigurations": {
			//       "description": "\u003cp\u003eSchema of the dataset identified by the placeholder. Any dashboard created from this\n            template should be bound to new datasets matching the same schema described through this\n            API operation.\u003c/p\u003e",
			//       "items": {
			//         "additionalProperties": false,
			//         "description": "\u003cp\u003eDataset configuration.\u003c/p\u003e",
			//         "properties": {
			//           "ColumnGroupSchemaList": {
			//             "description": "\u003cp\u003eA structure containing the list of column group schemas.\u003c/p\u003e",
			//             "items": {
			//               "additionalProperties": false,
			//               "description": "\u003cp\u003eThe column group schema.\u003c/p\u003e",
			//               "properties": {
			//                 "ColumnGroupColumnSchemaList": {
			//                   "description": "\u003cp\u003eA structure containing the list of schemas for column group columns.\u003c/p\u003e",
			//                   "items": {
			//                     "additionalProperties": false,
			//                     "description": "\u003cp\u003eA structure describing the name, data type, and geographic role of the columns.\u003c/p\u003e",
			//                     "properties": {
			//                       "Name": {
			//                         "description": "\u003cp\u003eThe name of the column group's column schema.\u003c/p\u003e",
			//                         "type": "string"
			//                       }
			//                     },
			//                     "type": "object"
			//                   },
			//                   "maxItems": 500,
			//                   "minItems": 0,
			//                   "type": "array"
			//                 },
			//                 "Name": {
			//                   "description": "\u003cp\u003eThe name of the column group schema.\u003c/p\u003e",
			//                   "type": "string"
			//                 }
			//               },
			//               "type": "object"
			//             },
			//             "maxItems": 500,
			//             "minItems": 0,
			//             "type": "array"
			//           },
			//           "DataSetSchema": {
			//             "additionalProperties": false,
			//             "description": "\u003cp\u003eDataset schema.\u003c/p\u003e",
			//             "properties": {
			//               "ColumnSchemaList": {
			//                 "description": "\u003cp\u003eA structure containing the list of column schemas.\u003c/p\u003e",
			//                 "items": {
			//                   "additionalProperties": false,
			//                   "description": "\u003cp\u003eThe column schema.\u003c/p\u003e",
			//                   "properties": {
			//                     "DataType": {
			//                       "description": "\u003cp\u003eThe data type of the column schema.\u003c/p\u003e",
			//                       "type": "string"
			//                     },
			//                     "GeographicRole": {
			//                       "description": "\u003cp\u003eThe geographic role of the column schema.\u003c/p\u003e",
			//                       "type": "string"
			//                     },
			//                     "Name": {
			//                       "description": "\u003cp\u003eThe name of the column schema.\u003c/p\u003e",
			//                       "type": "string"
			//                     }
			//                   },
			//                   "type": "object"
			//                 },
			//                 "maxItems": 500,
			//                 "minItems": 0,
			//                 "type": "array"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "Placeholder": {
			//             "description": "\u003cp\u003ePlaceholder.\u003c/p\u003e",
			//             "type": "string"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "maxItems": 30,
			//       "minItems": 0,
			//       "type": "array"
			//     },
			//     "Description": {
			//       "description": "\u003cp\u003eThe description of the template.\u003c/p\u003e",
			//       "maxLength": 512,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "Errors": {
			//       "description": "\u003cp\u003eErrors associated with this template version.\u003c/p\u003e",
			//       "items": {
			//         "additionalProperties": false,
			//         "description": "\u003cp\u003eList of errors that occurred when the template version creation failed.\u003c/p\u003e",
			//         "properties": {
			//           "Message": {
			//             "description": "\u003cp\u003eDescription of the error type.\u003c/p\u003e",
			//             "pattern": ".*\\S.*",
			//             "type": "string"
			//           },
			//           "Type": {
			//             "enum": [
			//               "SOURCE_NOT_FOUND",
			//               "DATA_SET_NOT_FOUND",
			//               "INTERNAL_FAILURE",
			//               "ACCESS_DENIED"
			//             ],
			//             "type": "string"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "minItems": 1,
			//       "type": "array"
			//     },
			//     "Sheets": {
			//       "description": "\u003cp\u003eA list of the associated sheets with the unique identifier and name of each sheet.\u003c/p\u003e",
			//       "items": {
			//         "additionalProperties": false,
			//         "description": "\u003cp\u003eA \u003ci\u003esheet\u003c/i\u003e, which is an object that contains a set of visuals that\n            are viewed together on one page in the Amazon QuickSight console. Every analysis and dashboard\n            contains at least one sheet. Each sheet contains at least one visualization widget, for\n            example a chart, pivot table, or narrative insight. Sheets can be associated with other\n            components, such as controls, filters, and so on.\u003c/p\u003e",
			//         "properties": {
			//           "Name": {
			//             "description": "\u003cp\u003eThe name of a sheet. This name is displayed on the sheet's tab in the QuickSight\n            console.\u003c/p\u003e",
			//             "pattern": ".*\\S.*",
			//             "type": "string"
			//           },
			//           "SheetId": {
			//             "description": "\u003cp\u003eThe unique identifier associated with a sheet.\u003c/p\u003e",
			//             "maxLength": 2048,
			//             "minLength": 1,
			//             "pattern": "[\\w\\-]+",
			//             "type": "string"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "maxItems": 20,
			//       "minItems": 0,
			//       "type": "array"
			//     },
			//     "SourceEntityArn": {
			//       "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of an analysis or template that was used to create this\n            template.\u003c/p\u003e",
			//       "type": "string"
			//     },
			//     "Status": {
			//       "enum": [
			//         "CREATION_IN_PROGRESS",
			//         "CREATION_SUCCESSFUL",
			//         "CREATION_FAILED",
			//         "UPDATE_IN_PROGRESS",
			//         "UPDATE_SUCCESSFUL",
			//         "UPDATE_FAILED",
			//         "DELETED"
			//       ],
			//       "type": "string"
			//     },
			//     "ThemeArn": {
			//       "description": "\u003cp\u003eThe ARN of the theme associated with this version of the template.\u003c/p\u003e",
			//       "type": "string"
			//     },
			//     "VersionNumber": {
			//       "description": "\u003cp\u003eThe version number of the template version.\u003c/p\u003e",
			//       "minimum": 1,
			//       "type": "number"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "<p>A version of a template.</p>",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"created_time": {
						// Property: CreatedTime
						Description: "<p>The time that this template version was created.</p>",
						Type:        types.StringType,
						Optional:    true,
					},
					"data_set_configurations": {
						// Property: DataSetConfigurations
						Description: "<p>Schema of the dataset identified by the placeholder. Any dashboard created from this\n            template should be bound to new datasets matching the same schema described through this\n            API operation.</p>",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"column_group_schema_list": {
									// Property: ColumnGroupSchemaList
									Description: "<p>A structure containing the list of column group schemas.</p>",
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"column_group_column_schema_list": {
												// Property: ColumnGroupColumnSchemaList
												Description: "<p>A structure containing the list of schemas for column group columns.</p>",
												Attributes: tfsdk.ListNestedAttributes(
													map[string]tfsdk.Attribute{
														"name": {
															// Property: Name
															Description: "<p>The name of the column group's column schema.</p>",
															Type:        types.StringType,
															Optional:    true,
														},
													},
													tfsdk.ListNestedAttributesOptions{},
												),
												Optional: true,
												Validators: []tfsdk.AttributeValidator{
													validate.ArrayLenBetween(0, 500),
												},
											},
											"name": {
												// Property: Name
												Description: "<p>The name of the column group schema.</p>",
												Type:        types.StringType,
												Optional:    true,
											},
										},
										tfsdk.ListNestedAttributesOptions{},
									),
									Optional: true,
									Validators: []tfsdk.AttributeValidator{
										validate.ArrayLenBetween(0, 500),
									},
								},
								"data_set_schema": {
									// Property: DataSetSchema
									Description: "<p>Dataset schema.</p>",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"column_schema_list": {
												// Property: ColumnSchemaList
												Description: "<p>A structure containing the list of column schemas.</p>",
												Attributes: tfsdk.ListNestedAttributes(
													map[string]tfsdk.Attribute{
														"data_type": {
															// Property: DataType
															Description: "<p>The data type of the column schema.</p>",
															Type:        types.StringType,
															Optional:    true,
														},
														"geographic_role": {
															// Property: GeographicRole
															Description: "<p>The geographic role of the column schema.</p>",
															Type:        types.StringType,
															Optional:    true,
														},
														"name": {
															// Property: Name
															Description: "<p>The name of the column schema.</p>",
															Type:        types.StringType,
															Optional:    true,
														},
													},
													tfsdk.ListNestedAttributesOptions{},
												),
												Optional: true,
												Validators: []tfsdk.AttributeValidator{
													validate.ArrayLenBetween(0, 500),
												},
											},
										},
									),
									Optional: true,
								},
								"placeholder": {
									// Property: Placeholder
									Description: "<p>Placeholder.</p>",
									Type:        types.StringType,
									Optional:    true,
								},
							},
							tfsdk.ListNestedAttributesOptions{},
						),
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenBetween(0, 30),
						},
					},
					"description": {
						// Property: Description
						Description: "<p>The description of the template.</p>",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 512),
						},
					},
					"errors": {
						// Property: Errors
						Description: "<p>Errors associated with this template version.</p>",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"message": {
									// Property: Message
									Description: "<p>Description of the error type.</p>",
									Type:        types.StringType,
									Optional:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringMatch(regexp.MustCompile(".*\\S.*"), ""),
									},
								},
								"type": {
									// Property: Type
									Type:     types.StringType,
									Optional: true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringInSlice([]string{
											"SOURCE_NOT_FOUND",
											"DATA_SET_NOT_FOUND",
											"INTERNAL_FAILURE",
											"ACCESS_DENIED",
										}),
									},
								},
							},
							tfsdk.ListNestedAttributesOptions{},
						),
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenAtLeast(1),
						},
					},
					"sheets": {
						// Property: Sheets
						Description: "<p>A list of the associated sheets with the unique identifier and name of each sheet.</p>",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"name": {
									// Property: Name
									Description: "<p>The name of a sheet. This name is displayed on the sheet's tab in the QuickSight\n            console.</p>",
									Type:        types.StringType,
									Optional:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringMatch(regexp.MustCompile(".*\\S.*"), ""),
									},
								},
								"sheet_id": {
									// Property: SheetId
									Description: "<p>The unique identifier associated with a sheet.</p>",
									Type:        types.StringType,
									Optional:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 2048),
										validate.StringMatch(regexp.MustCompile("[\\w\\-]+"), ""),
									},
								},
							},
							tfsdk.ListNestedAttributesOptions{},
						),
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenBetween(0, 20),
						},
					},
					"source_entity_arn": {
						// Property: SourceEntityArn
						Description: "<p>The Amazon Resource Name (ARN) of an analysis or template that was used to create this\n            template.</p>",
						Type:        types.StringType,
						Optional:    true,
					},
					"status": {
						// Property: Status
						Type:     types.StringType,
						Optional: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"CREATION_IN_PROGRESS",
								"CREATION_SUCCESSFUL",
								"CREATION_FAILED",
								"UPDATE_IN_PROGRESS",
								"UPDATE_SUCCESSFUL",
								"UPDATE_FAILED",
								"DELETED",
							}),
						},
					},
					"theme_arn": {
						// Property: ThemeArn
						Description: "<p>The ARN of the theme associated with this version of the template.</p>",
						Type:        types.StringType,
						Optional:    true,
					},
					"version_number": {
						// Property: VersionNumber
						Description: "<p>The version number of the template version.</p>",
						Type:        types.Float64Type,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.FloatAtLeast(1.000000),
						},
					},
				},
			),
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
			// Version is a write-only property.
		},
		"version_description": {
			// Property: VersionDescription
			// CloudFormation resource type schema:
			// {
			//   "description": "\u003cp\u003eA description of the current template version being created. This API operation creates the\n\t\t\tfirst version of the template. Every time \u003ccode\u003eUpdateTemplate\u003c/code\u003e is called, a new\n\t\t\tversion is created. Each version of the template maintains a description of the version\n\t\t\tin the \u003ccode\u003eVersionDescription\u003c/code\u003e field.\u003c/p\u003e",
			//   "maxLength": 512,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "<p>A description of the current template version being created. This API operation creates the\n\t\t\tfirst version of the template. Every time <code>UpdateTemplate</code> is called, a new\n\t\t\tversion is created. Each version of the template maintains a description of the version\n\t\t\tin the <code>VersionDescription</code> field.</p>",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 512),
			},
			// VersionDescription is a write-only property.
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			tfsdk.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Definition of the AWS::QuickSight::Template Resource Type.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::QuickSight::Template").WithTerraformTypeName("awscc_quicksight_template")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"actions":                         "Actions",
		"arn":                             "Arn",
		"aws_account_id":                  "AwsAccountId",
		"column_group_column_schema_list": "ColumnGroupColumnSchemaList",
		"column_group_schema_list":        "ColumnGroupSchemaList",
		"column_schema_list":              "ColumnSchemaList",
		"created_time":                    "CreatedTime",
		"data_set_arn":                    "DataSetArn",
		"data_set_configurations":         "DataSetConfigurations",
		"data_set_placeholder":            "DataSetPlaceholder",
		"data_set_references":             "DataSetReferences",
		"data_set_schema":                 "DataSetSchema",
		"data_type":                       "DataType",
		"description":                     "Description",
		"errors":                          "Errors",
		"geographic_role":                 "GeographicRole",
		"key":                             "Key",
		"last_updated_time":               "LastUpdatedTime",
		"message":                         "Message",
		"name":                            "Name",
		"permissions":                     "Permissions",
		"placeholder":                     "Placeholder",
		"principal":                       "Principal",
		"sheet_id":                        "SheetId",
		"sheets":                          "Sheets",
		"source_analysis":                 "SourceAnalysis",
		"source_entity":                   "SourceEntity",
		"source_entity_arn":               "SourceEntityArn",
		"source_template":                 "SourceTemplate",
		"status":                          "Status",
		"tags":                            "Tags",
		"template_id":                     "TemplateId",
		"theme_arn":                       "ThemeArn",
		"type":                            "Type",
		"value":                           "Value",
		"version":                         "Version",
		"version_description":             "VersionDescription",
		"version_number":                  "VersionNumber",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/VersionDescription",
		"/properties/SourceEntity",
		"/properties/CreatedTime",
		"/properties/Version",
		"/properties/LastUpdatedTime",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
