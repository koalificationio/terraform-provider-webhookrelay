---
layout: "webhookrelay"
page_title: "webhookrelay: webhookrelay_channel"
sidebar_current: "docs-aws-datasource-channel"
description: |-
  Get information on webhookrelay channel.
---

# Data Source: webhookrelay_channel

Use this data source to get information on an existing webhookrelay channel.

## Example Usage

```hcl
data "webhookrelay_channel" "alerts" {
  name = "alerts"
}
```

## Argument Reference

* `name` - (Required) Indicate a name of a channel to find.


## Attributes Reference

`id` is set to the ID of the found channel. In addition, the following attributes are exported:

* `is_general` - will be true if this channel is the `general` channel that includes all regular team members.
* `is_archived` - will be true if the channel is archived.
* `is_shared` - means the channel is in some way shared between multiple workspaces.
* `is_private` - means the conversation is privileged between two or more members.


