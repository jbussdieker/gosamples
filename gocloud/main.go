package main

import "text/template"
import "bytes"
import "path/filepath"

import "github.com/hoisie/web"

type Session struct {
	*Aws
	Error string
}

var session *Session

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

func instance_list(ctx *web.Context, value string) (render string) {
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

    web.Get("/(.*)", instance_list)
	web.Run("0.0.0.0:8000")
}

