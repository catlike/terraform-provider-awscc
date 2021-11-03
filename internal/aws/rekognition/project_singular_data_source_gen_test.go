// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package rekognition_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSRekognitionProjectDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Rekognition::Project", "awscc_rekognition_project", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSRekognitionProjectDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Rekognition::Project", "awscc_rekognition_project", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
