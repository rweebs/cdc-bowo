provider "aws" {
  region = "us-east-1"
}

provider "sql" {
  alias = "blue"
  url   = "postgres://${local.primary_username}:${local.primary_password}@${local.primary_endpoint_address}/${local.primary_database}?sslmode=allow"
}

provider "sql" {
  alias = "green"
  url   = "postgres://${local.secondary_username}:${local.secondary_password}@${local.secondary_endpoint_address}/${local.secondary_database}?sslmode=allow"
}
