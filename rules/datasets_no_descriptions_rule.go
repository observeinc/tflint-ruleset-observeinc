package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// DatasetsNoDescriptionsRule checks whether ...
type DatasetsNoDescriptionsRule struct {
	tflint.DefaultRule
}

// NewDatasetsNoDescriptionsRule returns a new rule
func NewDatasetsNoDescriptionsRule() *DatasetsNoDescriptionsRule {
	return &DatasetsNoDescriptionsRule{}
}

// Name returns the rule name
func (r *DatasetsNoDescriptionsRule) Name() string {
	return "datasets_no_descriptions_rule"
}

// Enabled returns whether the rule is enabled by default
func (r *DatasetsNoDescriptionsRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *DatasetsNoDescriptionsRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *DatasetsNoDescriptionsRule) Link() string {
	return ""
}

// Check checks whether observe_dataset has description attribute
func (r *DatasetsNoDescriptionsRule) Check(runner tflint.Runner) error {
	_, err := runner.GetResourceContent("observe_dataset", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "description"},
		},
	}, nil)
	fmt.Print("Test")
	if err != nil {
		return err
	}

	return nil
}
