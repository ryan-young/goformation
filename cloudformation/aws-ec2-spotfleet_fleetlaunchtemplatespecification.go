package cloudformation

// AWSEC2SpotFleet_FleetLaunchTemplateSpecification AWS CloudFormation Resource (AWS::EC2::SpotFleet.FleetLaunchTemplateSpecification)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-fleetlaunchtemplatespecification.html
type AWSEC2SpotFleet_FleetLaunchTemplateSpecification struct {

	// LaunchTemplateId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-fleetlaunchtemplatespecification.html#cfn-ec2-spotfleet-fleetlaunchtemplatespecification-launchtemplateid
	LaunchTemplateId string `json:"LaunchTemplateId,omitempty"`

	// LaunchTemplateName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-fleetlaunchtemplatespecification.html#cfn-ec2-spotfleet-fleetlaunchtemplatespecification-launchtemplatename
	LaunchTemplateName string `json:"LaunchTemplateName,omitempty"`

	// Version AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-fleetlaunchtemplatespecification.html#cfn-ec2-spotfleet-fleetlaunchtemplatespecification-version
	Version string `json:"Version,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2SpotFleet_FleetLaunchTemplateSpecification) AWSCloudFormationType() string {
	return "AWS::EC2::SpotFleet.FleetLaunchTemplateSpecification"
}
