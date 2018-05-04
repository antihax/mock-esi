#!/bin/bash
go get github.com/sqs/goreturns
set -e
curl -s https://esi.evetech.net/versions/ | grep -Po '[a-z0-9]+' | {
    imports=()

    while read -r version ; do
        echo "Processing $version"
        rm -rf ./$version/*

        imports+=("_ \"github.com/antihax/mock-esi/$version/go\"\n")
   
        java -jar swagger-codegen-cli.jar generate -DpackageName=esi$version \
            -o ./$version -t ./mock-esi-template -l go-server \
            -i https://esi.evetech.net/$version/swagger.json?datasource=tranquility
    done

    imp=${imports[@]}

    sed "s#IMPORTS#$imp#" < ./mock-esi-template/main.go.template > ./main.go
}

goreturns -w .
gofmt -s -w .

go test ./...

