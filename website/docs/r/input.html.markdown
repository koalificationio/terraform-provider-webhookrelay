---
layout: "webhookrelay"
page_title: "Webhookrelay: webhookrelay_input"
description: |-
  Manage Webhookrelay input.
---

# Resource: webhookrelay_input

Use this resource to manage Webhookrelay input.

## Example Usage

```hcl
resource "webhookrelay_bucket" "foo" {
  name                 = "foo"
  description          = "foo"
  delete_default_input = true
}

resource "webhookrelay_input" "foo" {
  name        = "foo"
  description = "bar"
  bucket_id   = webhookrelay_bucket.foo.id
}
```

With headers:

```hcl
resource "webhookrelay_input" "foo" {
  name        = "foo"
  description = "bar"
  bucket_id   = webhookrelay_bucket.foo.id
  headers = {
    Content-Type = "text/plain; charset=utf-8"
    Foo          = "bar"
  }
}
```

## Argument Reference

* `name` - (Required) Name of an input to create.
* `bucket_id` - (Required) ID of a bucket for an input.

-> **NOTE:** Changing `bucket_id` will recreate input

* `description` - (Optional) Input description.
* `status_code` - (Optional) Specify what should Webhook Relay return to HTTP clients that call this endpoint.
* `response_body` - (Optional) Response body can be any text, JSON or XML payload.
* `headers` - (Optional) A map of headers and their values to be sent as a response to webhook.
* `function_id` - (Optional) ID of the function that will be executed for this input

## Attributes Reference

`id` is set to the ID of the input. It can be used in webhook url like: `https://my.webhookrelay.com/v1/webhooks/<input id>`

## Import

Input can be imported using the `id`, e.g.

```
$ terraform import webhookrelay_input.foo e3cb4587-6e3d-4c64-9b50-e9c4c7ce27aa
```
