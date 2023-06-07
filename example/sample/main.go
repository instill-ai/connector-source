package main

import (
	"fmt"

	connectorDestination "github.com/instill-ai/connector-source/pkg"
	"go.uber.org/zap"
)

func main() {

	logger, _ := zap.NewDevelopment()
	connector := connectorDestination.Init(logger)
	// Connection, err := connector.CreateConnection("70d8664a-d512-4517-a5e8-5d4da81756a7", request.Config{})
	// fmt.Println(err)
	for k, v := range connector.GetConnectorDefinitionMap() {
		fmt.Printf("%s %s\n", k, v)
	}
	// Connection.Execute(nil)

}
