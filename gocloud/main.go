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

type SessionInstance struct {
	*Instance
	HasNginx bool
	HasApache2 bool
	HasMySQL bool
}

type Session struct {
	*Aws
	regions []*Region
	instances map[string][]*Instance
	keypairs map[string][]*KeyPair
	SessionInstance
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
	console = strings.Replace(console, "[0;33m", "<font color=orange>", -1)
	console = strings.Replace(console, "[0;36m", "<font color=blue>", -1)
	console = strings.Replace(console, "[1;35m", "<font color=red>", -1)
	console = strings.Replace(console, "[0m", "</font>", -1)

	if err != nil {
		err = errors.New(console)
	}
	return
}

func (session *Session) Instances() []*Instance {
	if session.instances[session.CurrentRegion] == nil {
		println("instances not found for", session.CurrentRegion)
		instances, _ := session.DescribeInstances()
		session.instances[session.CurrentRegion] = instances
	}
	return session.instances[session.CurrentRegion]
}

func (session *Session) Regions() []*Region {
	if session.regions == nil {
		regions, _ := session.DescribeRegions()
		session.regions = regions
	}
	return session.regions
}

func (session *Session) KeyPairs() []*KeyPair {
	if session.keypairs[session.CurrentRegion] == nil {
		println("Key pair not found for", session.CurrentRegion)
		keypairs, _ := session.DescribeKeyPairs()
		session.keypairs[session.CurrentRegion] = keypairs
	}
	return session.keypairs[session.CurrentRegion]
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
	session.Error = ""
	session.Message = ""
	return string(buf.Bytes())
}

func (session *Session) GetInstance(id string) *Instance {
	for _, instance := range session.Instances() {
		if instance.ID == id {
			return instance
		}
	}
	return nil
}

func instance(ctx *web.Context, value string) (render string) {
	if value == "new" {
		return session.ServeTemplate("new_instance")
	}
	if value == "recipe" {
		console, err := installRecipe(ctx.Params["host"], ctx.Params["name"])
		if err != nil {
			session.Error = err.Error()
		} else {
			session.Message = console
			// TODO: This does not scale :(
			if ctx.Params["name"] == "nginx" {
				session.HasNginx = true
			}
			if ctx.Params["name"] == "nginx-remove" {
				session.HasNginx = false
			}
			if ctx.Params["name"] == "apache2" {
				session.HasApache2 = true
			}
			if ctx.Params["name"] == "apache2-remove" {
				session.HasApache2 = false
			}
			if ctx.Params["name"] == "mysql" {
				session.HasMySQL = true
			}
			if ctx.Params["name"] == "mysql-remove" {
				session.HasMySQL = false
			}
		}
		ctx.Redirect(302, "/instance/?id=" + ctx.Params["id"])
		return
	}
	session.Instance = session.GetInstance(ctx.Params["id"])
	return session.ServeTemplate("instance")
}

func instances(ctx *web.Context, value string) (render string) {
	if value == "refresh" {
		session.instances[session.CurrentRegion] = nil
		ctx.Redirect(302, "/instances/")
		return
	}
	if value == "create" {
		err := session.CreateInstance(ctx.Params["ami"], ctx.Params["keypair"])
		if err != nil {
			session.Error = err.Error()
		}
	}
	if value == "start" {
		err := session.StartInstance(ctx.Params["id"])
		if err != nil {
			session.Error = err.Error()
		}
	}
	if value == "stop" {
		err := session.StopInstance(ctx.Params["id"])
		if err != nil {
			session.Error = err.Error()
		}
	}
	if value == "terminate" {
		err := session.TerminateInstance(ctx.Params["id"])
		if err != nil {
			session.Error = err.Error()
		}
	}
	return session.ServeTemplate("instances")
}

func keypairs(ctx *web.Context, value string) (render string) {
	if value == "refresh" {
		session.keypairs[session.CurrentRegion] = nil
		ctx.Redirect(302, "/keypairs/")
		return
	}
	if value == "delete" {
		err := session.DeleteKeyPair(ctx.Params["name"])
		if err != nil {
			session.Error = err.Error()
		}
		ctx.Redirect(302, "/keypairs/")
		return
	}
	return session.ServeTemplate("keypairs")
}

func region(ctx *web.Context, value string) (render string) {
	if value == "set_region" {
		session.SetRegion(ctx.Params["region"])
		ctx.Redirect(302, "/instances/")
	}
	// TODO: Redirect for now...
	ctx.Redirect(302, "/instances/")
	return
}

func sitemain(ctx *web.Context) (render string) {
	ctx.Redirect(302, "/instances/")
	return
}

func main() {
	certfile := os.ExpandEnv("$EC2_CERT")
	keyfile := os.ExpandEnv("$EC2_PRIVATE_KEY")
	aws := NewAws(certfile, keyfile, "us-west-1")
	session = &Session{
		Aws:aws,
		keypairs: map[string][]*KeyPair{},
		instances: map[string][]*Instance{},
	}
    web.Get("/instance/(.*)", instance)
    web.Get("/instances/(.*)", instances)
    web.Get("/keypairs/(.*)", keypairs)
    web.Get("/region/(.*)", region)
    web.Get("/", sitemain)
	web.Run("0.0.0.0:8000")
}

