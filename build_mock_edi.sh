java -jar swagger-codegen-cli.jar generate -o ./ -t ./mock-esi-template -l go-server -i https://esi.tech.ccp.is/latest/swagger.json?datasource=tranquility
go get 
go run main.go