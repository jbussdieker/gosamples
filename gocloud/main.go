package main

import "text/template"
import "bytes"
import "os/exec"
import "strings"
import "errors"
import "path/filepath"

import "github.com/hoisie/web"

type Reservation struct {
}

type Tag struct {
}

type Instance struct {
	ID string
	Ami string
	PublicName string
	PrivateName string
	Status string
	KeyPair string
	Size string
	Region string
}

type TemplateContext struct {
	Error string
	Instances map[string]*Instance
}

var Instances map[string]*Instance

func ReadInstances() (instances map[string]*Instance, err error) {
	cmd := exec.Command("ec2-describe-instances")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, errors.New(string(out))
	}
	lines := strings.Split(string(out), "\n")
	instances = map[string]*Instance{}
	for _, line := range lines {
		parts := strings.Split(line, "\t")
		type_name := parts[0]
		switch type_name {
		case "INSTANCE":
			if parts[5] != "terminated" {
				instances[parts[1]] = &Instance {
					ID: parts[1],
					Ami: parts[2],
					PublicName: parts[3],
					PrivateName: parts[4],
					Status: parts[5],
					KeyPair: parts[6],
					Size: parts[8],
					Region: parts[10],
				}
			}
		}
	}
	return
}

func kill_instance(instance_id string) (err error) {
	cmd := exec.Command("ec2-terminate-instances", instance_id)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(out))
	}
	return
}

func new_instance() (err error) {
	cmd := exec.Command("ec2-run-instances", "ami-6da8f128", "-t", "t1.micro")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(out))
	}
	return
}

func serve_template(name string, ctx *TemplateContext) string {
	var buf bytes.Buffer
	t, err := template.ParseFiles(filepath.Join("templates", name))
	if err != nil {
		return "template parse error " + err.Error()
	}
	err = t.ExecuteTemplate(&buf, name, ctx)
	if err != nil {
		return "template exec error " + err.Error()
	}
	return string(buf.Bytes())
}

func instance_list(ctx *web.Context, value string) (render string) {
	tctx := &TemplateContext{
		Instances: Instances,
	}

	if value == "error" {
		tctx.Error = ctx.Params["error"]
	}
	if value == "new" {
		err := new_instance()
		if err != nil {
			ctx.Redirect(302, "/error?error=" + err.Error())
		} else {
			ctx.Redirect(302, "/")
		}
		value = "refresh"
	}
	if value == "terminate" {
		err := kill_instance(ctx.Params["id"])
		if err != nil {
			ctx.Redirect(302, "/error?error=" + err.Error())
		} else {
			ctx.Redirect(302, "/")
		}
		value = "refresh"
	}
	if value == "refresh" {
		instances, err := ReadInstances()
		Instances = instances
		if err != nil {
			println("Error getting instance list")
		}
		ctx.Redirect(302, "/")
	}

	return serve_template("instances", tctx)
}

func main() {
	instances, err := ReadInstances()
	Instances = instances
	if err != nil {
		println("Error getting instance list")
		return
	}

    web.Get("/(.*)", instance_list)
	web.Run("0.0.0.0:8000")
}

