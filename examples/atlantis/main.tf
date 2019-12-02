resource "webhookrelay_bucket" "atlantis_prod" {
  name                 = "atlantis-prod"
  description          = "Prod atlantis webhook bucket"
  ephemeral_webhooks   = true
  delete_default_input = true
}

resource "webhookrelay_input" "atlantis_prod" {
  name        = "atlantis-prod"
  description = "Prod atlantis webhook"
  bucket_id   = webhookrelay_bucket.atlantis_prod.id
}

resource "webhookrelay_output" "atlantis_local" {
  name        = "atlantis-local"
  destination = "http://127.0.0.1:4141/events"
  internal    = true
  bucket_id   = webhookrelay_bucket.atlantis_prod.id
}

resource "webhookrelay_token" "atlantis_prod" {
  description = "atlantis-prod"
  api_access  = "disabled"
  scopes {
    buckets = [webhookrelay_bucket.atlantis_prod.name]
  }
}

resource "gitlab_project_hook" "atlantis" {
  project                 = var.gitlab_project
  url                     = "https://my.webhookrelay.com/v1/webhooks/${webhookrelay_input.atlantis_prod.id}"
  token                   = var.atlantis_secret
  enable_ssl_verification = true

  merge_requests_events = true
  push_events           = true
  note_events           = true
}
