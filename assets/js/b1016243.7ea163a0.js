"use strict";(self.webpackChunkstencil=self.webpackChunkstencil||[]).push([[958],{7208:function(e,t,a){a.r(t),a.d(t,{frontMatter:function(){return o},contentTitle:function(){return s},metadata:function(){return p},toc:function(){return c},default:function(){return m}});var r=a(7462),n=a(3366),i=(a(7294),a(3905)),l=["components"],o={},s="Stencil server",p={unversionedId:"server/overview",id:"server/overview",isDocsHomePage:!1,title:"Stencil server",description:"Stencil is dynamic protobuf schema registry. It provides REST interface for storing and retrieving protobuf file descriptors.",source:"@site/docs/server/overview.md",sourceDirName:"server",slug:"/server/overview",permalink:"/stencil/docs/server/overview",editUrl:"https://github.com/odpf/stencil/edit/master/docs/docs/server/overview.md",tags:[],version:"current",frontMatter:{},sidebar:"docsSidebar",previous:{title:"Quick start guide",permalink:"/stencil/docs/guides/quick_start"},next:{title:"Backward compatability rules",permalink:"/stencil/docs/server/rules"}},c=[{value:"Features",id:"features",children:[]},{value:"Requirements",id:"requirements",children:[]},{value:"Installation",id:"installation",children:[{value:"Configuring environment Variables",id:"configuring-environment-variables",children:[]}]},{value:"Reference",id:"reference",children:[]},{value:"Quick start API usage examples",id:"quick-start-api-usage-examples",children:[]}],d={toc:c};function m(e){var t=e.components,a=(0,n.Z)(e,l);return(0,i.kt)("wrapper",(0,r.Z)({},d,a,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("h1",{id:"stencil-server"},"Stencil server"),(0,i.kt)("p",null,"Stencil is dynamic protobuf schema registry. It provides REST interface for storing and retrieving protobuf file descriptors."),(0,i.kt)("h2",{id:"features"},"Features"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"stores versioned history of proto descriptor file on specified namespace and name"),(0,i.kt)("li",{parentName:"ul"},"enforce backward compatability check on upload by default"),(0,i.kt)("li",{parentName:"ul"},"ability to skip some of the backward compatability checks while upload"),(0,i.kt)("li",{parentName:"ul"},"ability to download fully contained proto descriptor file for specified proto message ",(0,i.kt)("a",{parentName:"li",href:"https://pkg.go.dev/google.golang.org/protobuf@v1.27.1/reflect/protoreflect#FullName"},"fullName")),(0,i.kt)("li",{parentName:"ul"},"provides metadata API to retrieve latest version number given a name and namespace")),(0,i.kt)("h2",{id:"requirements"},"Requirements"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"postgres 11")),(0,i.kt)("h2",{id:"installation"},"Installation"),(0,i.kt)("p",null,"Run the following commands to run from docker image"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},"$ docker pull odpf/stencil\n")),(0,i.kt)("p",null,"Run the following commands to compile from source"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},"$ git clone git@github.com:odpf/stencil.git\n$ cd stencil/server\n$ go build -o stencil\n$ ./stencil # specify envs before executing this command\n")),(0,i.kt)("h3",{id:"configuring-environment-variables"},"Configuring environment Variables"),(0,i.kt)("p",null,"To run the stencil server, you will need to add the following environment variables"),(0,i.kt)("table",null,(0,i.kt)("thead",{parentName:"table"},(0,i.kt)("tr",{parentName:"thead"},(0,i.kt)("th",{parentName:"tr",align:"left"},"ENV"),(0,i.kt)("th",{parentName:"tr",align:"left"},"Description"))),(0,i.kt)("tbody",{parentName:"table"},(0,i.kt)("tr",{parentName:"tbody"},(0,i.kt)("td",{parentName:"tr",align:"left"},(0,i.kt)("inlineCode",{parentName:"td"},"PORT")),(0,i.kt)("td",{parentName:"tr",align:"left"},"port number default to ",(0,i.kt)("inlineCode",{parentName:"td"},"8080"))),(0,i.kt)("tr",{parentName:"tbody"},(0,i.kt)("td",{parentName:"tr",align:"left"},(0,i.kt)("inlineCode",{parentName:"td"},"TIMEOUT")),(0,i.kt)("td",{parentName:"tr",align:"left"},"graceful time to wait before shutting down the server. Takes ",(0,i.kt)("inlineCode",{parentName:"td"},"time.Duration")," format. Eg: ",(0,i.kt)("inlineCode",{parentName:"td"},"30s")," or ",(0,i.kt)("inlineCode",{parentName:"td"},"20m"))),(0,i.kt)("tr",{parentName:"tbody"},(0,i.kt)("td",{parentName:"tr",align:"left"},(0,i.kt)("inlineCode",{parentName:"td"},"DB_CONNECTIONSTRING")),(0,i.kt)("td",{parentName:"tr",align:"left"},"postgres db connection ",(0,i.kt)("a",{parentName:"td",href:"https://www.postgresql.org/docs/11/libpq-connect.html#LIBPQ-CONNSTRING"},"url"),". Eg: ",(0,i.kt)("inlineCode",{parentName:"td"},"postgres://postgres@localhost:5432/db_name"))),(0,i.kt)("tr",{parentName:"tbody"},(0,i.kt)("td",{parentName:"tr",align:"left"},(0,i.kt)("inlineCode",{parentName:"td"},"NEWRELIC_ENABLED")),(0,i.kt)("td",{parentName:"tr",align:"left"},"boolean to enable newrelic")),(0,i.kt)("tr",{parentName:"tbody"},(0,i.kt)("td",{parentName:"tr",align:"left"},(0,i.kt)("inlineCode",{parentName:"td"},"NEWRELIC_APPNAME")),(0,i.kt)("td",{parentName:"tr",align:"left"},"appname")),(0,i.kt)("tr",{parentName:"tbody"},(0,i.kt)("td",{parentName:"tr",align:"left"},(0,i.kt)("inlineCode",{parentName:"td"},"NEWRELIC_LICENSE")),(0,i.kt)("td",{parentName:"tr",align:"left"},"License key for newrelic")))),(0,i.kt)("h2",{id:"reference"},"Reference"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"/stencil/docs/server/api"},"API")),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"/stencil/docs/server/rules"},"Rules"))),(0,i.kt)("h2",{id:"quick-start-api-usage-examples"},"Quick start API usage examples"),(0,i.kt)("p",null,"The following assumes you have Stencil server up and running at port 8080 and ",(0,i.kt)("inlineCode",{parentName:"p"},"protoc")," is installed."),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},'$ mkdir example\n$ cd example\n# create example proto file. You can add as many proto files as you want.\n$ echo "syntax=\\"proto3\\";\\npackage stencil;\\nmessage One {\\n  int32 field_one = 1;\\n}" > 1.proto\n\n# create descriptor file\n$ protoc --descriptor_set_out=./file.desc --include_imports ./1.proto\n\n# upload descriptor file to server with name as `example` under `quickstart` namespace\n$ curl -X POST http://localhost:8080/v1/namespaces/quickstart/descriptors -F "file=@./file.desc" -F "version=0.0.1" -F "name=example" -F "latest=true" -H "Content-Type: multipart/form-data"\n\n# get list of descriptors available in a namespace\n$ curl -X GET http://localhost:8080/v1/namespaces/quickstart/descriptors\n\n# get list of versions available for particular descriptor\n$ curl -X GET http://localhost:8080/v1/namespaces/quickstart/descriptors/example/versions\n\n# download specific version of particular desciptor\n$ curl -X GET http://localhost:8080/v1/namespaces/quickstart/descriptors/example/versions/0.0.1\n\n# download latest version of particular descriptor\n$ curl -X GET http://localhost:8080/v1/namespaces/quickstart/descriptors/example/versions/latest\n\n# get latest version number of particular descriptor\n$ curl -X GET http://localhost:8080/v1/namespaces/quickstart/metadata/example\n\n# modify latest version number of particular descriptor\n$ curl -X POST \'http://localhost:8080/v1/namespaces/quickstart/metadata\' -H \'Content-Type: application/json\' --data-raw \'{"name": "example","version": "0.0.1"}\'\n')))}m.isMDXComponent=!0}}]);