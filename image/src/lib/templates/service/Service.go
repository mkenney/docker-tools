package service

import (
	"flag"
    "io/ioutil"
    "lib/config"
	"text/template"

	"github.com/golang/glog"
)

/*
Template is a reference to this module's template instance
*/
var Template *template.Template

func init() {
	file, err := ioutil.ReadFile(config.DefaultTemplateDir + "/service/Service.tmpl")
    if nil != err {
        glog.Fatalf("Could not read template file '%s'", err)
    }

    tmpl, err := template.New("css").Parse(string(file))
    if nil != err {
        glog.Fatalf("Could not parse template file '%s'", err)
    }
    Template = tmpl
}

func init() {
	flag.Parse() // Required for glog
}
