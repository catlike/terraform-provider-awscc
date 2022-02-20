// Code generated by generators/resource/main.go; DO NOT EDIT.

package route53

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
	registry.AddResourceTypeFactory("awscc_route53_key_signing_key", keySigningKeyResourceType)
}

// keySigningKeyResourceType returns the Terraform awscc_route53_key_signing_key resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Route53::KeySigningKey resource type.
func keySigningKeyResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"hosted_zone_id": {
			// Property: HostedZoneId
			// CloudFormation resource type schema:
			// {
			//   "description": "The unique string (ID) used to identify a hosted zone.",
			//   "pattern": "^[A-Z0-9]{1,32}$",
			//   "type": "string"
			// }
			Description: "The unique string (ID) used to identify a hosted zone.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringMatch(regexp.MustCompile("^[A-Z0-9]{1,32}$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"key_management_service_arn": {
			// Property: KeyManagementServiceArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon resource name (ARN) for a customer managed key (CMK) in AWS Key Management Service (KMS). The KeyManagementServiceArn must be unique for each key signing key (KSK) in a single hosted zone.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The Amazon resource name (ARN) for a customer managed key (CMK) in AWS Key Management Service (KMS). The KeyManagementServiceArn must be unique for each key signing key (KSK) in a single hosted zone.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 256),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "An alphanumeric string used to identify a key signing key (KSK). Name must be unique for each key signing key in the same hosted zone.",
			//   "pattern": "^[a-zA-Z0-9_]{3,128}$",
			//   "type": "string"
			// }
			Description: "An alphanumeric string used to identify a key signing key (KSK). Name must be unique for each key signing key in the same hosted zone.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringMatch(regexp.MustCompile("^[a-zA-Z0-9_]{3,128}$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "description": "A string specifying the initial status of the key signing key (KSK). You can set the value to ACTIVE or INACTIVE.",
			//   "enum": [
			//     "ACTIVE",
			//     "INACTIVE"
			//   ],
			//   "type": "string"
			// }
			Description: "A string specifying the initial status of the key signing key (KSK). You can set the value to ACTIVE or INACTIVE.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"ACTIVE",
					"INACTIVE",
				}),
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
		Description: "Represents a key signing key (KSK) associated with a hosted zone. You can only have two KSKs per hosted zone.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53::KeySigningKey").WithTerraformTypeName("awscc_route53_key_signing_key")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"hosted_zone_id":             "HostedZoneId",
		"key_management_service_arn": "KeyManagementServiceArn",
		"name":                       "Name",
		"status":                     "Status",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
