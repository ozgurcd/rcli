# RCLI

RCLI is a command line tool, written in Go language to create an AWS autoscalable installation with an Application Load Balancer. It produces a load balancer address to reach the deployed instances.

## Installation

### Prerequisites

RCLI uses the "Terraform" to manage AWS, thus relies on a working Terraform installation. Also, since it is written in Golang, it requires Go compiler to be able to compile it from source code. Lastly, Terraform relies on a working awscli installation (with .aws/credentials file).

### Terraform

RCLI generates Terraform tf files, then calls terraform to apply them. In order to do it, it creates a working directory. This directory name can be set as
a command line argument. However if it is not set, the default value Rcliworkdir will be used.

After the first run Terraform will create state files in the same directory.

NOTE: By default, RCli uses Amazon Linux 2 AMI for the instances, thus Deployment Template user data is set according to it. If you wish to use a different AMI with a different Linux distribution, you should update the user data section accordingly in templates/autoscaling.temptf.


## Build from Source Code

First please modify Makefile to be sure if your Operating System and Architecture is correctly set. In order to compile against Linux TARGET_OS should be set to linux

Then, issue the command "make", it should compile the source code and produce a binary named rcli

## Usage

RCLI has several command line options, the can be seen by issuing a rcli -h command.

Flags:
  -h, --help                     Show context-sensitive help (also try --help-long and --help-man).
      --instanceType="t2.micro"  Instance type to be created
      --amiPattern="amzn2-ami-hvm-2.*-x86_64*"
                                 Amazon Machine Image Pattern to be used
      --amiOwner="amazon"        Amazon Machine Image Owner to be used
      --region="us-west-1"       AWS region to be used
      --tags="Created by RCli"   Tags to use for AWS Resources
      --workdir="Rcliworkdir"    An empty directory for creating terraform scripts
      --profile=PROFILE          AWS Profile to use, can be created in .aws/credentials file
      --keypairname=KEYPAIRNAME  AWS SSH Keypair name to use
      --namePrefix="rcli-"       Launch Configuration Instance Name Prefix
      --autoScaleMin=2           Min number of backend instances
      --autoScaleMax=5           Max number of backend instances

      Example:

      ./rcli --keypairname=rclitest --profile=rclitest


## Authors

      Ozgur Demir <ozgurcd@gmail.com>
