variable "aws_region" {
  description = "AWS region"
  default     = "us-east-1"
}

variable "table_name" {
  description = "DynamoDB table name"
  default     = "movies"
}

variable "functions" {
  type        = "list"
  description = "Lambda functions"
  default     = ["FindAllMovies", "DeleteMovie", "UpdateMovie", "InsertMovie"]
}

variable "methods" {
  type = "map"

  default = {
    FindAllMovies = "GET"
    DeleteMovie   = "DELETE"
    UpdateMovie   = "PUT"
    InsertMovie   = "POST"
  }
}
