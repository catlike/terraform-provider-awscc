// Code generated by generators/resource/main.go; DO NOT EDIT.

package iotsitewise

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("awscc_iotsitewise_project", projectResourceType)
}

// projectResourceType returns the Terraform awscc_iotsitewise_project resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::IoTSiteWise::Project resource type.
func projectResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"portal_id": {
			// Property: PortalId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the portal in which to create the project.",
			//   "type": "string"
			// }
			Description: "The ID of the portal in which to create the project.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"project_arn": {
			// Property: ProjectArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the project.",
			//   "type": "string"
			// }
			Description: "The ARN of the project.",
			Type:        types.StringType,
			Computed:    true,
		},
		"project_description": {
			// Property: ProjectDescription
			// CloudFormation resource type schema:
			// {
			//   "description": "A description for the project.",
			//   "type": "string"
			// }
			Description: "A description for the project.",
			Type:        types.StringType,
			Optional:    true,
		},
		"project_id": {
			// Property: ProjectId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the project.",
			//   "type": "string"
			// }
			Description: "The ID of the project.",
			Type:        types.StringType,
			Computed:    true,
		},
		"project_name": {
			// Property: ProjectName
			// CloudFormation resource type schema:
			// {
			//   "description": "A friendly name for the project.",
			//   "type": "string"
			// }
			Description: "A friendly name for the project.",
			Type:        types.StringType,
			Required:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of key-value pairs that contain metadata for the project.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "To add or update tag, provide both key and value. To delete tag, provide only tag key to be deleted",
			//     "properties": {
			//       "Key": {
			//         "type": "string"
			//       },
			//       "Value": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "A list of key-value pairs that contain metadata for the project.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				Multiset(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS::IoTSiteWise::Project",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTSiteWise::Project").WithTerraformTypeName("awscc_iotsitewise_project")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"key":                 "Key",
		"portal_id":           "PortalId",
		"project_arn":         "ProjectArn",
		"project_description": "ProjectDescription",
		"project_id":          "ProjectId",
		"project_name":        "ProjectName",
		"tags":                "Tags",
		"value":               "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
