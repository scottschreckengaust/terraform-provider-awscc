---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_cloudfront_key_group Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::CloudFront::KeyGroup
---

# awscc_cloudfront_key_group (Data Source)

Data Source schema for AWS::CloudFront::KeyGroup



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `key_group_config` (Attributes) The key group configuration. (see [below for nested schema](#nestedatt--key_group_config))
- `key_group_id` (String)
- `last_modified_time` (String)

<a id="nestedatt--key_group_config"></a>
### Nested Schema for `key_group_config`

Read-Only:

- `comment` (String) A comment to describe the key group. The comment cannot be longer than 128 characters.
- `items` (List of String) A list of the identifiers of the public keys in the key group.
- `name` (String) A name to identify the key group.
