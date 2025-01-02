# Get current AWS region
data "aws_region" "current" {}

# Get current AWS account ID
data "aws_caller_identity" "current" {}

# IoT Wireless Network Analyzer Configuration
resource "awscc_iotwireless_network_analyzer_configuration" "example" {
  name        = "example-network-analyzer"
  description = "Example IoT Wireless Network Analyzer Configuration"

  trace_content = {
    log_level                  = "INFO"
    wireless_device_frame_info = "ENABLED"
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}