package main

import "text/template"
import "bytes"
import "path/filepath"
import "os"
import "errors"

import "github.com/hoisie/web"

type SessionInstance struct {
	*Instance
	*Apt
	HasNginx bool
	HasApache2 bool
	HasMySQL bool
}

type Session struct {
	*Aws
	regions []*Region
	instances map[string][]*Instance
	keypairs map[string][]*KeyPair
	*SessionInstance
	Error string
	Message string
}

var session *Session

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

func (instance *SessionInstance) ReadInfo() {
	instance.HasNginx = true
}

func (session *Session) GetInstance(id string) *SessionInstance {
	for _, instance := range session.Instances() {
		if instance.ID == id {
			session_instance := &SessionInstance{
				Instance: instance,
				Apt: NewApt(instance.PublicName),
			}
			session_instance.ReadInfo()
			return session_instance
		}
	}
	return nil
}

func instance(ctx *web.Context, value string) (render string) {
	if value == "new" {
		return session.ServeTemplate("new_instance")
	}
	if value == "recipe" {
		var err error
		if ctx.Params["action"] == "install" {
			err = session.Apt.Install(ctx.Params["name"])
		} else if ctx.Params["action"] == "remove" {
			err = session.Apt.Remove(ctx.Params["name"])
		} else {
			err = errors.New("Invalid recipe action")
		}
		if err != nil {
			session.Error = err.Error()
		}
		ctx.Redirect(302, "/instance/?id=" + ctx.Params["id"])
		return
	}
	session.SessionInstance = session.GetInstance(ctx.Params["id"])
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

