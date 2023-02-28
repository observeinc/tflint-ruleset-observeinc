package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_ObserveDatasetHasDescriptions(t *testing.T) {
	tests := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "issue found",
			Content: `
resource "observe_dataset" "aws_ecs" {}`,
			Expected: helper.Issues{
				{
					Rule:    NewObserveDatasetHasDescriptionsRule(),
					Message: "dataset does not have a description",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 37},
					},
				},
			},
		},
		{
			Name: "issue not found",
			Content: `
resource "observe_dataset" "aws_ecs" {
	description = "Amazon ECS is a fully managed container orchestration service that makes it easy for you to deploy, manage, and scale containerized applications."
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewObserveDatasetHasDescriptionsRule()

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{"resource.tf": test.Content})

			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error occurred: %s", err)
			}

			helper.AssertIssues(t, test.Expected, runner.Issues)
		})
	}
}
