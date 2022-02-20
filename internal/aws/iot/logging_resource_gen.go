// Code generated by generators/resource/main.go; DO NOT EDIT.

package iot

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
	registry.AddResourceTypeFactory("awscc_iot_logging", loggingResourceType)
}

// loggingResourceType returns the Terraform awscc_iot_logging resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::IoT::Logging resource type.
func loggingResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"account_id": {
			// Property: AccountId
			// CloudFormation resource type schema:
			// {
			//   "description": "Your 12-digit account ID (used as the primary identifier for the CloudFormation resource).",
			//   "maxLength": 12,
			//   "minLength": 12,
			//   "pattern": "^[0-9]{12}$",
			//   "type": "string"
			// }
			Description: "Your 12-digit account ID (used as the primary identifier for the CloudFormation resource).",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(12, 12),
				validate.StringMatch(regexp.MustCompile("^[0-9]{12}$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"default_log_level": {
			// Property: DefaultLogLevel
			// CloudFormation resource type schema:
			// {
			//   "description": "The log level to use. Valid values are: ERROR, WARN, INFO, DEBUG, or DISABLED.",
			//   "enum": [
			//     "ERROR",
			//     "WARN",
			//     "INFO",
			//     "DEBUG",
			//     "DISABLED"
			//   ],
			//   "type": "string"
			// }
			Description: "The log level to use. Valid values are: ERROR, WARN, INFO, DEBUG, or DISABLED.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"ERROR",
					"WARN",
					"INFO",
					"DEBUG",
					"DISABLED",
				}),
			},
		},
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the role that allows IoT to write to Cloudwatch logs.",
			//   "maxLength": 2048,
			//   "minLength": 20,
			//   "type": "string"
			// }
			Description: "The ARN of the role that allows IoT to write to Cloudwatch logs.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(20, 2048),
			},
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
		Description: "Logging Options enable you to configure your IoT V2 logging role and default logging level so that you can monitor progress events logs as it passes from your devices through Iot core service.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoT::Logging").WithTerraformTypeName("awscc_iot_logging")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_id":        "AccountId",
		"default_log_level": "DefaultLogLevel",
		"role_arn":          "RoleArn",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
