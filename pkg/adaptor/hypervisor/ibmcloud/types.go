package ibmcloud

type Config struct {
	ApiKey                   string
	IamServiceURL            string
	VpcServiceURL            string
	ProfileName              string
	ZoneName                 string
	ImageID                  string
	PrimarySubnetID          string
	PrimarySecurityGroupID   string
	SecondarySubnetID        string
	SecondarySecurityGroupID string
	KeyID                    string
	VpcID                    string
}
