terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = var.primary_region
}

module "stradivarius_app" {
  source = "./stradivarius"
  ami = var.ami
  instance_type = var.instance_type
  key_name = var.key_name
  port_http = var.port_http
  url_db = var.url_db
  token = var.token
}