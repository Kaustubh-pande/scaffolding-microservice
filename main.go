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
	providerUsername =flag.String("username", "bitbucket.org/rxsense", "Path for provider/Username")
)

func main()  {
	flag.Parse()
	microserviceName := *path+"/"+*name
	Createmicroservice(microserviceName)
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	folders := []string{"api","bin","cmd","cmd/client","cmd/grpc","cmd/rest","config","middleware","model","repository","server","usecase"}
	nestedFolders := []string{"cmd/client","cmd/grpc","cmd/rest"}
	//Creating folder & files
	for _,folder := range folders {
		basestring:= microserviceName+"/"+folder+"/"
		Createfolder(folder, microserviceName)
		switch folder {
		case "api":
				Createfile(basestring + *protoFileName + ".proto")
				ReadWrite(dir+"/Template/proto.txt", basestring+*protoFileName+".proto")
				Createfile(basestring + *protoFileName + ".yaml")
				ReadWrite(dir+"/Template/yaml.txt", basestring+*protoFileName+".yaml")
		case "config":
				Createfile(basestring + "config.go")
				ReadWrite(dir+"/Template/config.txt", basestring+"config.go")
		case "middleware":
				Createfile(basestring + "logger.go")
				ReadWrite(dir+"/Template/logger.txt", basestring+"logger.go")
		case "model":
				Createfile(basestring + "datastore.go")
				ReadWrite(dir+"/Template/datastore.txt", basestring+"datastore.go")
		case "repository":
				Createfile(basestring + *protoFileName + ".repository.go")
				ReadWrite(dir+"/Template/repository.txt", basestring+*protoFileName+".repository.go")
		case "server":
				Createfile(basestring + "grpc.go")
				ReadWrite(dir+"/Template/grpc.txt", basestring+"grpc.go")
				Createfile(basestring + "rest.go")
				ReadWrite(dir+"/Template/rest.txt", basestring+"rest.go")
		case "usecase":
				Createfile(basestring + *protoFileName + ".service.go")
				ReadWrite(dir+"/Template/service.txt", basestring+*protoFileName+".service.go")

		}
	}
	//Nested folders
	for _,folder := range nestedFolders{
		Createfile(microserviceName+"/"+folder+"/"+"main.go")
		ReadWrite( dir+"/Template/"+folder+".txt",microserviceName+"/"+folder+"/"+"main.go")
	}
	//file on root
	//config.yaml
	Createfile(microserviceName+"/"+"config.yaml")
	ReadWrite( dir+"/Template/configyaml.txt",microserviceName+"/"+"config.yaml")
	//makefile
	Createfile(microserviceName+"/"+"Makefile")
	ReadWrite( dir+"/Template/makefile.txt",microserviceName+"/"+"Makefile")
	//mod
	Createfile(microserviceName+"/"+"go.mod")
	ReadWrite( dir+"/Template/mod.txt",microserviceName+"/"+"go.mod")
	//cmd:=exec.Command("go mod init "+*path+"/"+*name)
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