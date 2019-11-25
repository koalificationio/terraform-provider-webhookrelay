output "webhookrelay_key" {
  value = webhookrelay_tokan.atlantis_prod.secret.key
}

output "webhookrelay_secret" {
  value = webhookrelay_tokan.atlantis_prod.secret
  sensitive = true
}
