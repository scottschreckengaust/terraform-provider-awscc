# Get current AWS account ID
data "aws_caller_identity" "current" {}

# Get current AWS region
data "aws_region" "current" {}

# Enable Detective Organization Admin
resource "awscc_detective_organization_admin" "admin" {
  account_id = data.aws_caller_identity.current.account_id
}