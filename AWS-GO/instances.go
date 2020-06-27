package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// GetInstances ...
func GetInstances(ec2svc *ec2.EC2) []string {

	var EC2InstanceIDs []string

	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("instance-state-name"),
				Values: []*string{
					aws.String("running"),
					aws.String("pending"),
				},
			},
		},
	}

	resp, err := ec2svc.DescribeInstances(params)
	if err != nil {
		fmt.Println("there was an error listing instances in", err.Error())
		log.Fatal(err.Error())
	}

	for idx := range resp.Reservations {
		//for idx, res := range resp.Reservations {
		//fmt.Println("  > Reservation Id", *res.ReservationId, " Num Instances: ", len(res.Instances))
		for _, inst := range resp.Reservations[idx].Instances {
			//fmt.Println("    - Instance ID: ", *inst.InstanceId)
			EC2InstanceIDs = append(EC2InstanceIDs, *inst.InstanceId)
		}
	}

	return EC2InstanceIDs
}

// GetInstance ...
func GetInstance(ec2svc *ec2.EC2, instanceID string) *ec2.Instance {
	filters := []*ec2.Filter{
		&ec2.Filter{
			Name: aws.String("instance-id"),
			Values: []*string{
				aws.String(instanceID),
			},
		},
	}

	input := ec2.DescribeInstancesInput{Filters: filters}
	result, err := ec2svc.DescribeInstances(&input)
	if err != nil {
		panic(err.Error())
	}
	// for _, reservation := range result.Reservations {
	// 	for _, instance := range reservation.Instances {
	// 		fmt.Printf("%s\n", *instance.InstanceId)
	// 	}
	// }

	res := result.Reservations[0]
	return res.Instances[0]
}
