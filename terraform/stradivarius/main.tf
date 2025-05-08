resource "aws_security_group" "stradivarius_sg" {
  name = "stradivarius_sg"
  description = "security group for this instance. Allow SSH and inbound traffic based on process port"

  ingress {
    from_port = 22
    to_port = 22
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port = var.port_http
    to_port = var.port_http
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port = 0
    to_port = 0
    protocol = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_instance" "stradivarius_instance" {  
  ami = var.ami
  instance_type = var.instance_type
  key_name = var.key_name

  security_groups = [aws_security_group.stradivarius_sg.name]

  user_data = templatefile("${path.module}/user_data.tftpl", {
    url_db = var.url_db
    token = var.token
    port_http = var.port_http
  })

  tags = {
    Name = "Stradivarius Process App"
  }
}