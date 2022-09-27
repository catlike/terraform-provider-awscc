// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package supportapp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_supportapp_account_alias", accountAliasDataSourceType)
}

// accountAliasDataSourceType returns the Terraform awscc_supportapp_account_alias data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::SupportApp::AccountAlias resource type.
func accountAliasDataSourceType(ctx context.Context) (provider.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"account_alias": {
			// Property: AccountAlias
			// CloudFormation resource type schema:
			// {
			//   "description": "An account alias associated with a customer's account.",
			//   "maxLength": 30,
			//   "minLength": 1,
			//   "pattern": "^[\\w\\- ]+$",
			//   "type": "string"
			// }
			Description: "An account alias associated with a customer's account.",
			Type:        types.StringType,
			Computed:    true,
		},
		"account_alias_resource_id": {
			// Property: AccountAliasResourceId
			// CloudFormation resource type schema:
			// {
			//   "description": "Unique identifier representing an alias tied to an account",
			//   "maxLength": 29,
			//   "minLength": 29,
			//   "pattern": "^[\\w\\- ]+$",
			//   "type": "string"
			// }
			Description: "Unique identifier representing an alias tied to an account",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::SupportApp::AccountAlias",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SupportApp::AccountAlias").WithTerraformTypeName("awscc_supportapp_account_alias")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_alias":             "AccountAlias",
		"account_alias_resource_id": "AccountAliasResourceId",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}