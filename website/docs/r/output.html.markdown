---
layout: "webhookrelay"
page_title: "Webhookrelay: webhookrelay_output"
description: |-
  Manage Webhookrelay output.
---

# Resource: webhookrelay_output

Use this resource to manage Webhookrelay output.

## Example Usage

```hcl
resource "webhookrelay_bucket" "foo" {
  name        = "foo"
  description = "foo"
}

resource "webhookrelay_output" "foo" {
  name        = "foo"
  description = "bar"
  destination = "https://example.com:8080"
  internal    = false
  bucket_id   = webhookrelay_bucket.foo.id
}
```

Output with rules for [blog post][3]:

```hcl
resource "webhookrelay_output" "foo" {
  name        = "foo"
  destination = "https://example.com:8080"
  bucket_id   = webhookrelay_bucket.foo.id
  
  rules = jsonencode({
    # we need to match first match rule and any of rules in or directive
    and = [
      {
        match = {
          parameter = {
            source = "header"
            name   = "X-Hub-Signature"
          }
          type = "payload-hash-sha1"
          secret = "very-secret"
        }
      },
      {
        or = [
          {
            match = {
              parameter = {
                source = "payload"
                name   = "ref"
              }
              type = "contains"
              substring = "refs/tags"
            }
          },
          {
            match = {
              parameter = {
                source = "payload"
                name   = "ref_type"
              }
              type = "value"
              value = "tag"
            }
          }
        }
      }
    ]
  })
}
```

## Argument Reference

* `name` - (Required) Name of an output to create.
* `bucket_id` - (Required) bucket id to attach output to.

-> **NOTE:** Changing `bucket_id` will recreate output

* `description` - (Optional) description of an output.
* `destination` - (Required) output destination.
* `internal` - (Optional) set to `true` if output is on local network. Defaults to `false`.
* `tls_verification` - (Optional) TLS verification for public endpoints.
* `rules` - (Optional) Configuration for rules-based webhook filtering & routing. This is a JSON formatted string. See [WebhookRelay documentation][1] on how to configure rules
  See [Rules](#rules) below for details about JSON fields.
* `function_id` - (Optional) ID of the function that will be executed for this output

## Attributes Reference

`id` is set to the ID of the output.

### Rules

. The `rules` mapping has following fields:

* `or` - (Optional) array of rules, any of which can be matched. Supports any number of `or`, `and` 
  and `match`. See [Match](#mathc) below for details.
* `and` - (Optional) array of rules, all of which must match. Supports any number of `or`, `and` 
  and `match`. See [Match](#mathc) below for details.
  
### Match

Match mapping has following fields:

* `parameter` - (Required) Payload parameter that rule will be matching. It has two fields:  
  - `source` - one of `header`, `payload`, `query`.
  - `name` - name of parameter from source to match.
* `type` - (Required) Rule type. Can be one of `value`, `contains`, `does-not-contain`,
  `regex`, `payload-hash-sha1`, `payload-hash-sha256`.

Depending on type you will need to set matching values:

* `regex` - (Optional) required when `regex` type is used.
* `secret` - (Optional) required when `payload-hash-sha1` or `payload-hash-sha256` types are used.
* `substring` - (Optional) required when `contains` or `does-not-contain` types are used.
* `value` - (Optional) required when `value` type is used.

## Import

Output can be imported using the `id`, e.g.

```
$ terraform import webhookrelay_output.foo e3cb4587-6e3d-4c64-9b50-e9c4c7ce27aa
```

[1]: https://webhookrelay.com/v1/guide/webhook-forwarding.html#Request-matching-rules
[2]: https://godoc.org/github.com/koalificationio/go-webhookrelay/pkg/openapi/models#MatchRule
[3]: https://webhookrelay.com/blog/2019/04/02/webhook-rule-based-filters/
