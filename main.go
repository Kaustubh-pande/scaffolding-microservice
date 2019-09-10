package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)
var (
	path = flag.String("path", ".", "Path for creating project ")
	name = flag.String("name","Default-microservice","Name of the microservice")
	protoFileName = flag.String("proto", "proto", "proto file name")
	providerUsername =flag.String("username", "github.com/username", "Path for provider/Username")
)

func main()  {
	flag.Parse()
	microserviceName := *path+"/"+*name
	Createmicroservice(microserviceName)
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	TemplateFile := dir+"/Template/"
	folders := []string{"api","bin","cmd","cmd/grpc","cmd/rest","config","middleware","model","repository","transport","usecase","usecase/handler"}
	//Creating folder & files
	for _,folder := range folders {
		basestring:= microserviceName+"/"+folder+"/"
		Createfolder(folder, microserviceName)
		switch folder {
		case "api":
				Createfile(basestring + *protoFileName + ".proto")
				ReadWrite(TemplateFile+"proto.txt", basestring+*protoFileName+".proto")
				Createfile(basestring + *protoFileName + ".yaml")
				ReadWrite(TemplateFile+"yaml.txt", basestring+*protoFileName+".yaml")
		case "config":
				Createfile(basestring + "config.go")
				ReadWrite(TemplateFile+"config.txt", basestring+"config.go")
		case "middleware":
				Createfile(basestring + "logger.go")
				ReadWrite(TemplateFile+"middleware.txt", basestring+"logger.go")
		case "model":
				Createfile(basestring + "datastore.go")
				ReadWrite(TemplateFile+"model.txt", basestring+"datastore.go")
		case "repository":
				Createfile(basestring + *protoFileName + ".repository.go")
				ReadWrite(TemplateFile+"repository.txt", basestring+*protoFileName+".repository.go")
		case "transport":
				Createfile(basestring + "grpc.go")
				ReadWrite(TemplateFile+"grpc.txt", basestring+"grpc.go")
				Createfile(basestring + "rest.go")
				ReadWrite(TemplateFile+"rest.txt", basestring+"rest.go")
		case "usecase/handler":
				Createfile(basestring + *protoFileName + ".service.go")
				ReadWrite(TemplateFile+"service.txt", basestring+*protoFileName+".service.go")
		case "cmd","cmd/grpc","cmd/rest":
				Createfile(basestring +"main.go")
				ReadWrite(TemplateFile+folder+".txt",basestring +"main.go")

		}
	}
	//file on root
	//config.yaml
	Createfile(microserviceName+"/"+"config.yaml")
	ReadWrite( TemplateFile+"configyaml.txt",microserviceName+"/"+"config.yaml")
	//makefile
	Createfile(microserviceName+"/"+"Makefile")
	ReadWrite( TemplateFile+"makefile.txt",microserviceName+"/"+"Makefile")
	//mod
	Createfile(microserviceName+"/"+"go.mod")
	ReadWrite( TemplateFile+"mod.txt",microserviceName+"/"+"go.mod")
	fmt.Println("Microservice Structure Created")
}
func Createmicroservice(microservicename string){
	err := os.Mkdir(microservicename,0700)
	if err !=nil{
		fmt.Println("err while creating microservice dir",err)
	}
}
func Createfolder(folder string, microserviceName string){
	err := os.Mkdir(microserviceName+"/"+ folder,0700)
	if err !=nil{
		fmt.Println("err while creating folder",err)
	}

}
func Createfile(file string){
	createFile,err := os.Create(file)
	if err !=nil{
		fmt.Println("err while creating file",err)
	}
	defer createFile.Close()
}
func ReadWrite(sourceFile string,destinationFile string){
	templateContent, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		fmt.Println("Error while Reading Template ", destinationFile)
		fmt.Println(err)
		return
	}
	customTemplate := strings.Replace(string(templateContent),"{name}",*protoFileName,-1)
	customTemplate = strings.Replace(string(customTemplate),"{Name}",strings.Title(*protoFileName),-1)
	customTemplate = strings.Replace(string(customTemplate),"{provider-username}",*providerUsername,-1)
	customTemplate = strings.Replace(string(customTemplate),"{microservices}",*name,-1)
	err = ioutil.WriteFile(destinationFile, []byte(customTemplate), 0777)
	if err != nil {
		fmt.Println("Error while writing file ", destinationFile)
		fmt.Println(err)
		return
	}
}