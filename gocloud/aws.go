package main

import "errors"
import "os/exec"
import "strings"

type Aws struct {
	certfile string
	keyfile string
	CurrentRegion string
}

type Region struct {
	Name string
	Host string
}

type KeyPair struct {
	Name string
	Fingerprint string
}

type Instance struct {
	ID string
	Ami string
	PublicName string
	PrivateName string
	Status string
	Startable bool
	Stopable bool
	Terminateable bool
	Scriptable bool
	KeyPair string
	Size string
	Region string
}

func NewAws(certfile string, keyfile string, region string) (aws *Aws) {
	aws = &Aws{
		certfile: certfile,
		keyfile: keyfile,
		CurrentRegion: region,
	}
	return
}

func (aws *Aws) runCommand(name string, args ...string) (lines[] string, err error) {
	args = append([]string{"-K", aws.keyfile, "-C", aws.certfile}, args...)
	cmd := exec.Command(name, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return lines, errors.New(string(out))
	}
	lines = strings.Split(string(out), "\n")
	return
}

func (aws *Aws) runRegionalCommand(name string, args ...string) (lines[] string, err error) {
	args = append([]string{"--region", aws.CurrentRegion}, args...)
	return aws.runCommand(name, args...)
}

func (aws *Aws) SetRegion(name string) (err error) {
	aws.CurrentRegion = name
	return
}

func (aws *Aws) DeleteKeyPair(keypair string) (err error) {
	_, err = aws.runRegionalCommand("ec2-delete-keypair", keypair)
	return
}

func (aws *Aws) CreateInstance(ami string, keypair string) (err error) {
	_, err = aws.runRegionalCommand("ec2-run-instances", ami, "-t", "t1.micro", "-k", keypair)
	return
}

func (aws *Aws) StartInstance(instance_id string) (err error) {
	_, err = aws.runRegionalCommand("ec2-start-instances", instance_id)
	return
}

func (aws *Aws) StopInstance(instance_id string) (err error) {
	_, err = aws.runRegionalCommand("ec2-stop-instances", instance_id)
	return
}

func (aws *Aws) TerminateInstance(instance_id string) (err error) {
	_, err = aws.runRegionalCommand("ec2-terminate-instances", instance_id)
	return
}

func (aws *Aws) DescribeRegions() (regions []*Region, err error) {
	lines, err := aws.runCommand("ec2-describe-regions")
	if err != nil {
		return
	}
	for _, line := range lines {
		parts := strings.Split(line, "\t")
		if parts[0] == "REGION" {
			regions = append(regions, &Region{
				Name: parts[1],
				Host: parts[2],
			})
		}
	}
	return regions, nil
}

func (aws *Aws) DescribeKeyPairs() (keypairs []*KeyPair, err error) {
	lines, err := aws.runRegionalCommand("ec2-describe-keypairs")
	if err != nil {
		return
	}
	keypairs = []*KeyPair{}
	for _, line := range lines {
		parts := strings.Split(line, "\t")
		if parts[0] == "KEYPAIR" {
			keypairs = append(keypairs, &KeyPair{
				Name: parts[1],
				Fingerprint: parts[2],
			})
		}
	}
	return keypairs, nil
}

func (aws *Aws) DescribeInstances() (instances []*Instance, err error) {
	lines, err := aws.runRegionalCommand("ec2-describe-instances")
	if err != nil {
		return
	}
	instances = []*Instance{}
	for _, line := range lines {
		parts := strings.Split(line, "\t")
		if parts[0] == "INSTANCE" {
			instance := &Instance {
				ID: parts[1],
				Ami: parts[2],
				PublicName: parts[3],
				PrivateName: parts[4],
				Status: parts[5],
				KeyPair: parts[6],
				Size: parts[8],
				Region: parts[10],
			}

			if parts[5] == "running" {
				instance.Stopable = true
			}
			if parts[5] == "stopped" {
				instance.Startable = true
			}
			if parts[5] != "terminated" {
				instance.Terminateable = true
			}

			instances = append(instances, instance)
		}
	}
	return instances, nil
}

