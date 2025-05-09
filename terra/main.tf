terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
  }

  cloud {
    organization = "M-Connect"

    workspaces {
      name = "terraform-course"
    }
  }
}

provider "aws" {
  region = "us-east-1"
}

resource "aws_instance" "terra" {
  ami           = "ami-084568db4383264d4"
  instance_type = "t2.micro"
}

