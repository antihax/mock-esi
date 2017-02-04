#!/bin/bash
rm ./latest/*
rm ./dev/*
rm ./legacy/*

java -jar swagger-codegen-cli.jar generate -DpackageName=esiLatest -o ./latest -t ./mock-esi-template -l go-server -i https://esi.tech.ccp.is/latest/swagger.json?datasource=tranquility
java -jar swagger-codegen-cli.jar generate -DpackageName=esiDev -o ./dev -t ./mock-esi-template -l go-server -i https://esi.tech.ccp.is/dev/swagger.json?datasource=tranquility
java -jar swagger-codegen-cli.jar generate -DpackageName=esiLegacy -o ./legacy -t ./mock-esi-template -l go-server -i https://esi.tech.ccp.is/legacy/swagger.json?datasource=tranquility

COUNTER=1
while [  $COUNTER -lt 5 ]; do
    rm ./v$COUNTER/*
    java -jar swagger-codegen-cli.jar generate -DpackageName=esiV$COUNTER -o ./v$COUNTER -t ./mock-esi-template -l go-server -i https://esi.tech.ccp.is/v$COUNTER/swagger.json?datasource=tranquility
    let COUNTER=COUNTER+1 
done

go get 
go run main.go