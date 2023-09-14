# data "aws_subnets" "db" {
#   filter {

#     name   = "tag:Name"
#     values = ["cdc-vpc-public*"]
#   }
# }


# data "aws_vpc" "cdc_vpc" {
#   filter {

#     name   = "tag:Name"
#     values = ["cdc-vpc"]
#   }
# }