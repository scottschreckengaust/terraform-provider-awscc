---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_workspacesweb_identity_provider Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::WorkSpacesWeb::IdentityProvider
---

# awscc_workspacesweb_identity_provider (Data Source)

Data Source schema for AWS::WorkSpacesWeb::IdentityProvider



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `identity_provider_arn` (String)
- `identity_provider_details` (Map of String)
- `identity_provider_name` (String)
- `identity_provider_type` (String)
- `portal_arn` (String)
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String)
- `value` (String)
