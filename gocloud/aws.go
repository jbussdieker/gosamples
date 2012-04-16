package main

import "errors"
import "os/exec"
import "strings"

type Aws struct {
	Regions []*Region
	Instances []*Instance
	KeyPairs []*KeyPair

	certificateFile string
	privateKeyFile string
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

func NewAws(cert string, privkey string, region string) (aws *Aws) {
	aws = &Aws{
		certificateFile: cert,
		privateKeyFile: privkey,
		CurrentRegion: region,
	}

	regions, err := aws.DescribeRegions()
	if err != nil {
		return nil
	}
	aws.Regions = regions

	keypairs, err := aws.DescribeKeyPairs()
	if err != nil {
		return nil
	}
	aws.KeyPairs = keypairs

	instances, err := aws.DescribeInstances()
	if err != nil {
		return nil
	}
	aws.Instances = instances
	return
}

func (aws *Aws) DescribeKeyPairs() (keypairs []*KeyPair, err error) {
	cmd := exec.Command("ec2-describe-keypairs", "-K", aws.privateKeyFile, "-C", aws.certificateFile, "--region", aws.CurrentRegion)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, errors.New(string(out))
	}
	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		parts := strings.Split(line, "\t")
		if parts[0] == "KEYPAIR" {
			keypairs = append(keypairs, &KeyPair{
				Name: parts[1],
				Fingerprint: parts[2],
			})
		}
	}
	aws.KeyPairs = keypairs
	return keypairs, nil
}

func (aws *Aws) CreateInstance(ami string, keypair string) (err error) {
	cmd := exec.Command("ec2-run-instances", "-K", aws.privateKeyFile, "-C", aws.certificateFile, ami, "-t", "t1.micro", "-k", keypair, "--region", aws.CurrentRegion)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(out))
	}
	aws.DescribeInstances()
	return
}

func (aws *Aws) TerminateInstance(instance_id string) (err error) {
	cmd := exec.Command("ec2-terminate-instances", "-K", aws.privateKeyFile, "-C", aws.certificateFile, instance_id, "--region", aws.CurrentRegion)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(out))
	}
	aws.DescribeInstances()
	return
}

func (aws *Aws) StartInstance(instance_id string) (err error) {
	cmd := exec.Command("ec2-start-instances", "-K", aws.privateKeyFile, "-C", aws.certificateFile, instance_id, "--region", aws.CurrentRegion)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(out))
	}
	aws.DescribeInstances()
	return
}

func (aws *Aws) StopInstance(instance_id string) (err error) {
	cmd := exec.Command("ec2-stop-instances", "-K", aws.privateKeyFile, "-C", aws.certificateFile, instance_id, "--region", aws.CurrentRegion)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(out))
	}
	aws.DescribeInstances()
	return
}

func (aws *Aws) DeleteKeyPair(keypair string) (err error) {
	cmd := exec.Command("ec2-delete-keypair", "-K", aws.privateKeyFile, "-C", aws.certificateFile, keypair, "--region", aws.CurrentRegion)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(out))
	}
	aws.DescribeKeyPairs()
	return
}

func (aws *Aws) SetRegion(name string) (err error) {
	aws.CurrentRegion = name
	_, err = aws.DescribeKeyPairs()
	if err != nil {
		return
	}
	_, err = aws.DescribeInstances()
	return
}

func (aws *Aws) DescribeRegions() (regions []*Region, err error) {
	cmd := exec.Command("ec2-describe-regions", "-K", aws.privateKeyFile, "-C", aws.certificateFile)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, errors.New(string(out))
	}
	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		parts := strings.Split(line, "\t")
		if parts[0] == "REGION" {
			regions = append(regions, &Region{
				Name: parts[1],
				Host: parts[2],
			})
		}
	}
	aws.Regions = regions
	return regions, nil
}

func (aws *Aws) DescribeInstances() (instances []*Instance, err error) {
	cmd := exec.Command("ec2-describe-instances", "-K", aws.privateKeyFile, "-C", aws.certificateFile, "--region", aws.CurrentRegion)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, errors.New(string(out))
	}
	lines := strings.Split(string(out), "\n")
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
	aws.Instances = instances
	return instances, nil
}

