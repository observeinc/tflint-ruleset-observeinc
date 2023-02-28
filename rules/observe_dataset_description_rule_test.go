package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_ObserveDatasetDescription(t *testing.T) {
	tests := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "no description attribute",
			Content: `
resource "observe_dataset" "test" {}`,
			Expected: helper.Issues{
				{
					Rule:    NewObserveDatasetDescriptionRule(),
					Message: "dataset needs a description",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 34},
					},
				},
			},
		},
		{
			Name: "empty description attribute",
			Content: `
resource "observe_dataset" "test" {
	description = ""
}`,
			Expected: helper.Issues{
				{
					Rule:    NewObserveDatasetDescriptionRule(),
					Message: "dataset description should not be empty",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 16},
						End:      hcl.Pos{Line: 3, Column: 18},
					},
				},
			},
		},
		{
			Name: "issue not found",
			Content: `
resource "observe_dataset" "test" {
	description = "The description."
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewObserveDatasetDescriptionRule()

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
