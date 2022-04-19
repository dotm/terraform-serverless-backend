# Learn Terraform - Lambda functions and API Gateway

AWS Lambda functions and API gateway are often used to create serverlesss
applications.

Follow along with this [tutorial on HashiCorp
Learn](https://learn.hashicorp.com/tutorials/terraform/lambda-api-gateway?in=terraform/aws).

## TODO

- add CRUD to dynamodb
- refactor main.tf
- add stages (prod, staging, qa) (don't forget to change api gateway tf config)
- add tests

## Go Init

- go mod init module/name
- go get github.com/aws/aws-lambda-go

## Testing

- curl "$(terraform output -raw base_url)/hello?Name=Terraform"