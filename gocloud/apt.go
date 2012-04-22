package main

import "os/exec"
import "errors"

type Apt struct {
	host string
}

func NewApt(host string) *Apt {
	return &Apt{host}
}

func (apt *Apt) Status(pkgname string) bool {
	cmd := exec.Command("ssh", "-i", "master.pem", "ubuntu@" + apt.host, "sudo", "dpkg", "-s", pkgname)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}
	return true	
}

func (apt *Apt) Install(pkgname string) error {
	cmd := exec.Command("ssh", "-i", "master.pem", "ubuntu@" + apt.host, "sudo", "apt-get", "install", "-y", pkgname)
	out, err := cmd.CombinedOutput()
	println(string(out))
	if err != nil {
		return errors.New(err.Error() + "\n" + string(out))
	}
	return nil	
}

func (apt *Apt) Remove(pkgname string) error {
	cmd := exec.Command("ssh", "-i", "master.pem", "ubuntu@" + apt.host, "sudo", "apt-get", "remove", "-y", pkgname)
	out, err := cmd.CombinedOutput()
	println(string(out))
	if err != nil {
		return errors.New(err.Error() + "\n" + string(out))
	}
	return nil	
}
