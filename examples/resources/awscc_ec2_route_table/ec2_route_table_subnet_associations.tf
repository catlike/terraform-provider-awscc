resource "awscc_ec2_vpc" "vpc" {
  cidr_block       = "10.0.0.0/16"
  instance_tenancy = "default"
  tags = [{
    key   = "Name"
    value = "demovpc"
  }]
}

resource "awscc_ec2_route_table" "custom_route_table" {
  vpc_id = awscc_ec2_vpc.vpc.id
  tags = [{
    key   = "Name"
    value = "Custom Route Table"
  }]
}

resource "awscc_ec2_subnet" "subnet1" {
  vpc_id = awscc_ec2_vpc.vpc.id

  cidr_block        = "10.0.101.0/24"
  availability_zone = "us-east-1a"

  tags = [{
    key   = "Name"
    value = "Demo Subnet"
  }]
}

resource "awscc_ec2_subnet_route_table_association" "subnet_route_table_association" {
  route_table_id = awscc_ec2_route_table.custom_route_table.id
  subnet_id      = awscc_ec2_subnet.subnet1.id
}
