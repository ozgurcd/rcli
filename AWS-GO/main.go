package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	sess, err := session.NewSession(
		&aws.Config{
			Region:      aws.String("us-east-1"),
			Credentials: credentials.NewSharedCredentials("", "aws-imagizer"),
		},
	)
	if nil != err {
		panic(err)
	}

	ec2svc := ec2.New(sess)

	instances := GetInstances(ec2svc)
	for _, instanceID := range instances {
		fmt.Printf("InstanceID: [%s]\n", instanceID)
		instance := GetInstance(ec2svc, instanceID)
		fmt.Printf("Name: %s\n", *instance.KeyName)
		fmt.Printf("ImageId: %s\n", *instance.ImageId)
		fmt.Printf("Type: %s\n", *instance.InstanceType)
		fmt.Printf("Launch Time: %v\n", *instance.LaunchTime)
		fmt.Printf("Private DNS Name: %s\n", *instance.PrivateDnsName)
		fmt.Printf("Private IP Address: %s\n", *instance.PrivateIpAddress)
		fmt.Printf("Security Groups:\n")
		secGroups := instance.SecurityGroups
		for _, sec := range secGroups {
			fmt.Printf(">> %s\n", *sec.GroupName)
		}
		//fmt.Printf("Name: %+v\n", *instance.SecurityGroups)
	}
}
