package main

import "text/template"
import "bytes"
import "path/filepath"
import "os"
import "os/exec"
import "path"
import "errors"
import "strings"

import "github.com/hoisie/web"

type Session struct {
	*Aws
	Error string
	Message string
}

var session *Session

func installRecipe(host string, name string) (console string, err error) {
	wd, _ := os.Getwd()
	cmd := exec.Command(path.Clean(path.Join(wd, "setup.sh")), host, name)
	out, err := cmd.CombinedOutput()
	console = string(out)
	console = strings.Replace(console, "\n", "<br>", -1)
	console = strings.Replace(console, "[0;32m", "<font color=green>", -1)
	console = strings.Replace(console, "[0;36m", "<font color=blue>", -1)
	console = strings.Replace(console, "[1;35m", "<font color=red>", -1)
	console = strings.Replace(console, "[0m", "</font>", -1)

	if err != nil {
		err = errors.New(console)
	}
	return
}

func (session *Session) ServeTemplate(name string) string {
	var buf bytes.Buffer
	t, err := template.ParseGlob(filepath.Join("templates", "*"))
	if err != nil {
		return "template parse error " + err.Error()
	}
	err = t.ExecuteTemplate(&buf, name, session)
	if err != nil {
		return "template exec error " + err.Error()
	}
	return string(buf.Bytes())
}

func serv(ctx *web.Context, value string) (render string) {
	session.Error = ""
	session.Message = ""

	if value == "error" {
		session.Error = ctx.Params["error"]
	}
	if value == "change_region" {
		session.SetRegion(ctx.Params["region"])
	}
	if value == "new" {
		return session.ServeTemplate("new_instance")
	}
	if value == "keypairs" {
		return session.ServeTemplate("keypairs")
	}
	if value == "delete_keypair" {
		err := session.DeleteKeyPair(ctx.Params["name"])
		if err != nil {
			ctx.Redirect(302, "/error?error=" + err.Error())
		} else {
			ctx.Redirect(302, "/keypairs")
		}
	}
	if value == "create" {
		err := session.CreateInstance(ctx.Params["ami"], ctx.Params["keypair"])
		if err != nil {
			ctx.Redirect(302, "/error?error=" + err.Error())
		}
	}
	if value == "terminate" {
		err := session.TerminateInstance(ctx.Params["id"])
		if err != nil {
			ctx.Redirect(302, "/error?error=" + err.Error())
		}
	}
	if value == "start" {
		err := session.StartInstance(ctx.Params["id"])
		if err != nil {
			ctx.Redirect(302, "/error?error=" + err.Error())
		}
	}
	if value == "stop" {
		err := session.StopInstance(ctx.Params["id"])
		if err != nil {
			ctx.Redirect(302, "/error?error=" + err.Error())
		}
	}
	if value == "recipe" {
		console, err := installRecipe(ctx.Params["host"], ctx.Params["name"])
		if err != nil {
			ctx.Redirect(302, "/error?error=" + err.Error())
		} else {
			session.Message = console
		}
	}
	if value == "refresh" {
		_, err := session.DescribeInstances()
		if err != nil {
			ctx.Redirect(302, "/error?error=" + err.Error())
		} else {
			ctx.Redirect(302, "/")
		}
	}

	return session.ServeTemplate("instances")
}

func main() {
	aws := NewAws("/home/jbussdieker/.ec2/cert-62VJ7DAYEMXEINZQDU5UVMOB6IOXOSHC.pem", "/home/jbussdieker/.ec2/pk-62VJ7DAYEMXEINZQDU5UVMOB6IOXOSHC.pem", "us-west-1")
	session = &Session{Aws:aws}

    web.Get("/(.*)", serv)
	web.Run("0.0.0.0:8000")
}

