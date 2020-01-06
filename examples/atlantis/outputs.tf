output "webhookrelay_key" {
  value = webhookrelay_token.atlantis_prod.key
}

output "webhookrelay_secret" {
  value     = webhookrelay_token.atlantis_prod.secret
  sensitive = true
}
