package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// ObserveDatasetDescriptionRule checks whether ...
type ObserveDatasetDescriptionRule struct {
	tflint.DefaultRule
}

// NewObserveDatasetDescriptionRule returns a new rule
func NewObserveDatasetDescriptionRule() *ObserveDatasetDescriptionRule {
	return &ObserveDatasetDescriptionRule{}
}

// Name returns the rule name
func (r *ObserveDatasetDescriptionRule) Name() string {
	return "observe_dataset_description_rule"
}

// Enabled returns whether the rule is enabled by default
func (r *ObserveDatasetDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *ObserveDatasetDescriptionRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *ObserveDatasetDescriptionRule) Link() string {
	return "https://github.com/observeinc/tflint-ruleset-observeinc/blob/dani/no-dataset-descriptions/rules.md#observe-dataset-description"
}

// Check checks whether observe_dataset has description attribute
func (r *ObserveDatasetDescriptionRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent("observe_dataset", &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: "description"},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attr, exists := resource.Body.Attributes["description"]
		if !exists {
			return runner.EmitIssue(
				r,
				"Dataset needs a description",
				resource.DefRange,
			)
		}

		var description string
		err := runner.EvaluateExpr(attr.Expr, &description, nil)

		err = runner.EnsureNoError(err, func() error {
			if description == "" {
				return runner.EmitIssue(
					r,
					"Dataset description should not be empty",
					attr.Expr.Range(),
				)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
