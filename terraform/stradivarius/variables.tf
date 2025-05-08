variable "ami" {
  type = string
  description = "ami-id for launch this instance"
}

variable "instance_type" {
  type = string
  description = "specifications for this instance"
}

variable "key_name" {
  type = string
  description = "key_pair(login). Already must be created from AWS. Include its key_name"
}

variable "port_http" {
  type = number
  description = "port to launch process for this instance"
}

variable "url_db" {
  type = string
  sensitive = true
  description = "Url of your DB in CockroachDB"
}

variable "token" {
  type = string
  sensitive = true
  description = "Token for external API to populate data stock. Please set it with -var"
}