package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"text/template"
)

func createTerraformMain() {
	type TerraformData struct {
		DateTime string
		Profile  string
		Region   string
	}
	tpl, err := template.ParseFiles("templates/main.temptf")
	if err != nil {
		log.Fatalln(err)
		return
	}
	td := TerraformData{
		DateTime: initTime,
		Profile:  awsProfile,
		Region:   awsRegion,
	}

	f, err := os.Create(workdir + "/main.tf")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	err = tpl.Execute(f, td)
	if err != nil {
		log.Fatal(err)
	}
}

func createTerraformAutoScaling() {
	type TerraformData struct {
		DateTime string
	}
	tpl, err := template.ParseFiles("templates/autoscaling.temptf")
	if err != nil {
		log.Fatalln(err)
		return
	}
	td := TerraformData{
		DateTime: initTime,
	}

	f, err := os.Create(workdir + "/autoscaling.tf")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	err = tpl.Execute(f, td)
	if err != nil {
		log.Fatal(err)
	}
}

func createTerraformNetwork() {
	type TerraformData struct {
		DateTime string
	}
	tpl, err := template.ParseFiles("templates/network.temptf")
	if err != nil {
		log.Fatalln(err)
		return
	}
	td := TerraformData{
		DateTime: initTime,
	}

	f, err := os.Create(workdir + "/network.tf")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	err = tpl.Execute(f, td)
	if err != nil {
		log.Fatal(err)
	}
}

func createTerraformVariables() {
	type TerraformData struct {
		DateTime     string
		AMIOwner     string
		AMIPattern   string
		LCNamePrefix string
		InstanceType string
		KeyName      string
		ASMin        int
		ASMax        int
	}

	tpl, err := template.ParseFiles("templates/variables.temptf")
	if err != nil {
		log.Fatalln(err)
		return
	}
	td := TerraformData{
		DateTime:     initTime,
		AMIOwner:     awsAmiOwner,
		AMIPattern:   awsAmiPattern,
		LCNamePrefix: awsLCNamePrefix,
		InstanceType: awsInstanceType,
		KeyName:      awsKeyPairName,
		ASMin:        awsASMinSize,
		ASMax:        awsASMaxSize}

	f, err := os.Create(workdir + "/variables.tf")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	err = tpl.Execute(f, td)
	if err != nil {
		log.Fatal(err)
	}
}

func createTerraformConfigs() {
	createTerraformVariables()
	createTerraformNetwork()
	createTerraformAutoScaling()
	createTerraformMain()
}

func initTerraform() {
	log.Printf("Working directory [%s] initializing for Terraform\n", workdir)
	cmd := exec.Command("terraform", "init")
	cmd.Dir = workdir

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error received while initilazing Terraform\n%s", out)
	}
	log.Printf("Directory successfully initialized\n")
}

func applyTerraformChanges() {
	log.Printf("Terraform is starting to apply changes")
	cmd := exec.Command("terraform", "apply", "-input=false", "-auto-approve")
	cmd.Dir = workdir

	var stdout, stderr []byte
	var errStdout, errStderr error
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	err := cmd.Start()
	if err != nil {
		log.Fatalf("Unable to run Terraform with error: %s\n", err)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		stdout, errStdout = copyAndCapture(os.Stdout, stdoutIn)
		wg.Done()
	}()

	stderr, errStderr = copyAndCapture(os.Stderr, stderrIn)
	wg.Wait()
	_ = cmd.Wait()
	if errStdout != nil || errStderr != nil {
		log.Fatal("failed to capture stdout or stderr\n")
	}
	outStr, errStr := string(stdout), string(stderr)
	_ = outStr
	_ = errStr
	log.Printf("Terraform apply is done")
}

func copyAndCapture(w io.Writer, r io.Reader) ([]byte, error) {
	var out []byte
	buf := make([]byte, 1024, 1024)
	for {
		n, err := r.Read(buf[:])
		if n > 0 {
			d := buf[:n]
			out = append(out, d...)

			if w == os.Stderr {
				_, err := w.Write(d)
				if err != nil {
					return out, err
				}
			} else {
				analyze(d)
			}
		}
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return out, err
		}
	}
}

func analyze(d []byte) {
	line := string(d)

	if strings.Contains(line, "Creation complete") || strings.Contains(line, "lb_dnsname = ") {
		log.Printf("%s", line)
	}
}
