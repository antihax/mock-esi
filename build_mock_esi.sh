#!/bin/bash
rm -rf ./latest/*
rm -rf ./dev/*
rm -rf ./legacy/*

set -e
java -jar swagger-codegen-cli.jar generate -DpackageName=esiLatest -o ./latest -t ./mock-esi-template -l go-server -i https://esi.tech.ccp.is/_latest/swagger.json?datasource=tranquility
java -jar swagger-codegen-cli.jar generate -DpackageName=esiDev -o ./dev -t ./mock-esi-template -l go-server -i https://esi.tech.ccp.is/_dev/swagger.json?datasource=tranquility
java -jar swagger-codegen-cli.jar generate -DpackageName=esiLegacy -o ./legacy -t ./mock-esi-template -l go-server -i https://esi.tech.ccp.is/_legacy/swagger.json?datasource=tranquility

go test ./...