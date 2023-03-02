# TFLint Observeinc Plugin Ruleset
[![Build Status](https://github.com/terraform-linters/tflint-ruleset-template/workflows/build/badge.svg?branch=main)](https://github.com/terraform-linters/tflint-ruleset-template/actions)

This is a custom ruleset for Observe's Terraform provider. For documentation on writing custom rules, see [writing plugins](https://github.com/terraform-linters/tflint/blob/master/docs/developer-guide/plugins.md).

## Requirements

- TFLint v0.40+
- Go v1.19

## Rules

|Name|Description|Severity|Enabled|Link|
| --- | --- | --- | --- | --- |
|observe_dataset_description_rule|Checks that `observe_dataset` resources have a non-empty description attribute|WARNING|✔||

## Building the plugin

Clone the repository locally and run the following command:

```
$ make
```

You can easily install the built plugin with the following:

```
$ make install
```

You can run the built plugin like the following:

```
$ cat << EOS > .tflint.hcl
plugin "template" {
  enabled = true
}
EOS
$ tflint
```
