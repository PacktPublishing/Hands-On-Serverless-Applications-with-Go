terraform {
  backend "s3" {
    bucket = "terraform-state-files"
    key    = "path/to/my/key"
    region = "us-east-1"
  }
}
