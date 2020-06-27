package main

import "gopkg.in/alecthomas/kingpin.v2"

func main() {
	var (
		argSecretsFile = kingpin.Flag("openstacksecrets", "AWS Profile to use, can be created in .aws/credentials file").String()
	)
	kingpin.HelpFlag.Short('h')
}
