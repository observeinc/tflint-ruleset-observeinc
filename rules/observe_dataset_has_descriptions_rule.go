package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// ObserveDatasetHasDescriptionsRule checks whether ...
type ObserveDatasetHasDescriptionsRule struct {
	tflint.DefaultRule
}

// NewObserveDatasetHasDescriptionsRule returns a new rule
func NewObserveDatasetHasDescriptionsRule() *ObserveDatasetHasDescriptionsRule {
	return &ObserveDatasetHasDescriptionsRule{}
}

// Name returns the rule name
func (r *ObserveDatasetHasDescriptionsRule) Name() string {
	return "observe_dataset_has_descriptions_rule"
}

// Enabled returns whether the rule is enabled by default
func (r *ObserveDatasetHasDescriptionsRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *ObserveDatasetHasDescriptionsRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *ObserveDatasetHasDescriptionsRule) Link() string {
	return ""
}

// Check checks whether observe_dataset has description attribute
func (r *ObserveDatasetHasDescriptionsRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("observe_dataset", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "description"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		_, exists := resource.Body.Attributes["description"]
		if !exists {
			return runner.EmitIssue(
				r,
				"dataset does not have a description",
				resource.DefRange,
			)
		}
	}

	return nil
}
