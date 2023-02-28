package rules

import (

	// "github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"fmt"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/terraform/addrs"
	"github.com/terraform-linters/tflint-plugin-sdk/terraform/lang"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// ObserveDatasetOutputRule checks whether ...
type ObserveDatasetOutputRule struct {
	tflint.DefaultRule
}

type declarations struct {
	Variables     map[string]*hclext.Block
	DataResources map[string]*hclext.Block
}

// NewObserveDatasetOutputRule returns a new rule
func NewObserveDatasetOutputRule() *ObserveDatasetOutputRule {
	return &ObserveDatasetOutputRule{}
}

// Name returns the rule name
func (r *ObserveDatasetOutputRule) Name() string {
	return "observe_dataset_output"
}

// Enabled returns whether the rule is enabled by default
func (r *ObserveDatasetOutputRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *ObserveDatasetOutputRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *ObserveDatasetOutputRule) Link() string {
	return ""
}

// Check checks whether observe_dataset has output defined in output.tf
func (r *ObserveDatasetOutputRule) Check(runner tflint.Runner) error {
	decl := &declarations{
		Variables:     map[string]*hclext.Block{},
		DataResources: map[string]*hclext.Block{},
	}

	body, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type:       "variable",
				LabelNames: []string{"name"},
				Body:       &hclext.BodySchema{},
			},
			{
				Type:       "data",
				LabelNames: []string{"type", "name"},
				Body:       &hclext.BodySchema{},
			},
		},
	}, &tflint.GetModuleContentOption{ExpandMode: tflint.ExpandModeNone})
	if err != nil {
		return err
	}

	for _, block := range body.Blocks {
		if block.Type == "variable" {
			decl.Variables[block.Labels[0]] = block
		} else {
			decl.DataResources[fmt.Sprintf("data.%s.%s", block.Labels[0], block.Labels[1])] = block
		}
	}

	path, err := runner.GetModulePath()
	if err != nil {
		return err
	}
	if !path.IsRoot() {
		// This rule does not evaluate child modules.
		return nil
	}

	if err != nil {
		return err
	}
	diags := runner.WalkExpressions(tflint.ExprWalkFunc(func(expr hcl.Expression) hcl.Diagnostics {
		r.checkForRefsInExpr(expr, decl)
		return nil
	}))
	if diags.HasErrors() {
		return diags
	}

	for _, variable := range decl.Variables {
		logger.Warn(fmt.Sprintf("var: %+v", variable))
	}
	for _, data := range decl.DataResources {
		logger.Warn(fmt.Sprintf("var: %+v", data))
	}

	// content, err := runner.GetModuleContent(&hclext.BodySchema{
	// 	Blocks: []hclext.BlockSchema{
	// 		{
	// 			Type:       "output",
	// 			LabelNames: []string{"name"},
	// 			Body: &hclext.BodySchema{
	// 				Attributes: []hclext.AttributeSchema{{Name: "value"}},
	// 			},
	// 		},
	// 	},
	// }, &tflint.GetModuleContentOption{ExpandMode: tflint.ExpandModeNone})
	// if err != nil {
	// 	return err
	// }

	// for _, output := range content.Blocks {
	// 	// attr := output.Body.Attributes["value"]
	// 	logger.Warn(fmt.Sprintf("attr type: %+v\n", output))
	// }

	return nil
}

func (r *ObserveDatasetOutputRule) checkForRefsInExpr(expr hcl.Expression, decl *declarations) {
	for _, ref := range lang.ReferencesInExpr(expr) {
		switch sub := ref.Subject.(type) {
		case addrs.InputVariable:
			delete(decl.Variables, sub.Name)
		case addrs.Resource:
			delete(decl.DataResources, sub.String())
		case addrs.ResourceInstance:
			delete(decl.DataResources, sub.Resource.String())
		}
	}
}
