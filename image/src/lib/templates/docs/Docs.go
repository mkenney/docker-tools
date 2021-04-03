package docs

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"lib/config"
	"os"
	"text/template"
	"time"

	"github.com/golang/glog"
)

/*
compileHelpTemplate will compile a help document from the specified template
and Execute it using the provided templateData storing the result in the
provided buffer
*/
func compileHelpTemplate(buffer *bytes.Buffer, templateName string, templateData map[string]string) (reterr error) {

	tmpl := getTemplate("_head")
    reterr = tmpl.Execute(buffer, templateData)
    if nil == reterr {
		tmpl = getTemplate(templateName)
		reterr = tmpl.Execute(buffer, templateData)
		if nil == reterr {
			stats, reterr := os.Stat("/go/bin/docker-tools")
			if nil == reterr {
				templateData["footBUILD_ID"]   = fmt.Sprintf("%s", os.Getenv("GIT_COMMIT"))
				templateData["footBUILD_DATE"] = fmt.Sprintf("%s", stats.ModTime().Format(time.RFC3339))
				tmpl = getTemplate("_foot")
				reterr = tmpl.Execute(buffer, templateData)
			}
		}
	}

    return
}

/*
compileUsageTemplate will compile a help document from the specified template
and Execute it using the provided templateData storing the result in the
provided buffer
*/
func compileUsageTemplate(buffer *bytes.Buffer, templateName string, templateData map[string]string) (reterr error) {
	tmpl := getTemplate(templateName)
	reterr = tmpl.Execute(buffer, templateData)
    return
}

/*
getTemplate will return the requested template file as a parsed Template pointer
*/
func getTemplate(tmplName string) (retval *template.Template) {
    file, err := ioutil.ReadFile(config.DockerToolsTemplateDir+"/docs/"+tmplName+".tmpl")
	if nil != err {
		glog.Fatalf("Could not read template file '%s': %s", tmplName, err)
	}

	retval, err = template.New("css").Parse(string(file))
	if nil != err {
		glog.Fatalf("Could not parse template file '%s': %s", tmplName, err)
	}

    return
}
