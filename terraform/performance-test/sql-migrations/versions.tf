# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

terraform {

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.7.0"
    }

    sql = {
      source  = "paultyng/sql"
      version = "0.5.0"
    }
  }

  required_version = "~> 1.3"
}
