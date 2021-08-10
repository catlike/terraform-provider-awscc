// Code generated by generators/resource/main.go; DO NOT EDIT.

package route53_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/acctest"
)

func TestAccAWSRoute53DNSSEC_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Route53::DNSSEC", "aws_route53_dnssec", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}