---
description: |-
  Manage Webhookrelay bucket.
---

# Resource: webhookrelay_bucket

Use this resource to manage Webhookrelay bucket.

## Example Usage

```hcl
resource "webhookrelay_bucket" "foo" {
  name                 = "foo"
  description          = "bar"
}
```

## Argument Reference

* `name` - (Required) Name of a bucket to create.
* `description` - (Optional) description of a bucket.
* `auth` - (Optional) Enable authentication fo bucket. See [Auth](#auth) below for details.
* `ephemeral_webhooks` - (Optional) when enabled - request body, headers and query are not recorded in the database.
* `websocket_streaming` - (Optional) allows agents to subscribe via WebSocket protocol to this bucket.
* `delete_default_input` - (Optional) delete default input that is added upon bucket creation. You can create new inputs using [`resource_webhookrelay_input`][1] resource.


## Attributes Reference

`id` is set to the ID of the bucket. In addition, the following attributes are exported:

* `default_input` - configuration of default input. See [Inputs](#inputs) below for details.
* `input` - configuration of other inputs available in bucket. See [Inputs](#inputs) below for details.

### Auth

The `auth` field allows following attributes:

* `type` - (Optional) Type of authentication. Can be `basic` or `token`.
* `username` - (Optional) required when `basic` type is used.
* `password` - (Optional) required when `basic` type is used.
* `token` - (Optional) required when `token` type is used.

### Inputs

The `input` mapping provides following attributes:

* `name`
* `description`
* `id` - ID of an input. It can be used in webhook url like: `https://my.webhookrelay.com/v1/webhooks/<input id>`

## Import

Bucket can be imported using the `id`, e.g.

```
$ terraform import webhookrelay_bucket.foo e3cb4587-6e3d-4c64-9b50-e9c4c7ce27aa
```

[1]: /providers/koalificationio/webhookrelay/latest/docs/resources/input
