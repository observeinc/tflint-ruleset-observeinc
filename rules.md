# Observe TFLint Rules Documentation

This page serves to elaborate on the purpose of each rule that is tested.

## Observe Dataset Descriptions

This rule checks that `observe_dataset` resources have a `description` attribute, and that the `description` attribute is not an empty string.

**Valid Dataset**

```hcl
resource "observe_dataset" "test" {
	description = "The description."
}
```

**Invalid Datasets**

```hcl
resource "observe_dataset" "test" {
	description = ""
}
```

```hcl
resource "observe_dataset" "test" {}
```