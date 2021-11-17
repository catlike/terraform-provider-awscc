// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package redshift_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSRedshiftScheduledActionDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Redshift::ScheduledAction", "awscc_redshift_scheduled_action", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSRedshiftScheduledActionDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Redshift::ScheduledAction", "awscc_redshift_scheduled_action", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}