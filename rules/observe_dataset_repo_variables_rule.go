package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// ObserveDatasetRepoVariablesRule checks whether ...
type ObserveDatasetRepoVariablesRule struct {
	tflint.DefaultRule
}

// NewObserveDatasetRepoVariablesRule returns a new rule
func NewObserveDatasetRepoVariablesRule() *ObserveDatasetRepoVariablesRule {
	return &ObserveDatasetRepoVariablesRule{}
}

// Name returns the rule name
func (r *ObserveDatasetRepoVariablesRule) Name() string {
	return "observe_dataset_repo_variables_rule"
}

// Enabled returns whether the rule is enabled by default
func (r *ObserveDatasetRepoVariablesRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *ObserveDatasetRepoVariablesRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *ObserveDatasetRepoVariablesRule) Link() string {
	return "https://github.com/observeinc/tflint-ruleset-observeinc/blob/dani/no-dataset-descriptions/rules.md#observe-dataset-repo-variables"
}

// Check checks whether observe_dataset has repo variable attributes
func (r *ObserveDatasetRepoVariablesRule) Check(runner tflint.Runner) error {
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
