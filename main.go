package main

import (
	"log"
	"os"
	"time"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	workdir         string
	awsInstanceType string
	awsAmi          string
	awsAmiPattern   string
	awsAmiOwner     string
	awsRegion       string
	awsTags         string
	awsProfile      string
	awsKeyPairName  string
	awsLCNamePrefix string
	awsASMinSize    int
	awsASMaxSize    int
	initTime        string
)

func main() {
	var (
		argInstanceType = kingpin.Flag("instanceType", "Instance type to be created").Default("t2.micro").String()
		argAmiPattern   = kingpin.Flag("amiPattern", "Amazon Machine Image Pattern to be used").Default("amzn2-ami-hvm-2.*-x86_64*").String()
		argAmiOwner     = kingpin.Flag("amiOwner", "Amazon Machine Image Owner to be used").Default("amazon").String()
		argRegion       = kingpin.Flag("region", "AWS region to be used").Default("us-west-1").String()
		argTags         = kingpin.Flag("tags", "Tags to use for AWS Resources").Default("Created by RCli").String()
		argWorkdir      = kingpin.Flag("workdir", "An empty directory for creating terraform scripts").Default("Rcliworkdir").String()
		argAwsProfile   = kingpin.Flag("profile", "AWS Profile to use, can be created in .aws/credentials file").String()
		argAwsKeyName   = kingpin.Flag("keypairname", "AWS SSH Keypair name to use").String()
		argLCNamePrefix = kingpin.Flag("namePrefix", "Launch Configuration Instance Name Prefix").Default("rcli-").String()
		argAsMinSize    = kingpin.Flag("autoScaleMin", "Min number of backend instances").Default("2").Int()
		argAsMaxSize    = kingpin.Flag("autoScaleMax", "Max number of backend instances").Default("5").Int()
	)

	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	workdir = *argWorkdir
	awsInstanceType = *argInstanceType
	awsAmiOwner = *argAmiOwner
	awsAmiPattern = *argAmiPattern
	awsRegion = *argRegion
	awsTags = *argTags
	awsASMinSize = *argAsMinSize
	awsASMaxSize = *argAsMaxSize
	awsProfile = *argAwsProfile
	awsKeyPairName = *argAwsKeyName
	awsLCNamePrefix = *argLCNamePrefix

	if awsKeyPairName == "" {
		log.Fatal("Keypair name is a mandatory field, please provide it")
	}

	os.RemoveAll(workdir)
	os.Mkdir(workdir, 0755)
	initTime = time.Now().String()

	createTerraformConfigs()
	initTerraform()
	applyTerraformChanges()
}
