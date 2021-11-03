// Code generated by generators/resource/main.go; DO NOT EDIT.

package s3outposts_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSS3OutpostsBucket_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::S3Outposts::Bucket", "awscc_s3outposts_bucket", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
