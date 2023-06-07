package request

import (
	"sync"

	_ "embed"

	"github.com/gofrs/uuid"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/instill-ai/connector/pkg/base"
	"github.com/instill-ai/connector/pkg/configLoader"

	connectorPB "github.com/instill-ai/protogen-go/vdp/connector/v1alpha"
)

// Note: this is a dummy connector

//go:embed config/seed/source_definitions.yaml
var sourceDefinitionsYaml []byte

//go:embed config/seed/source_specs.yaml
var sourceSpecsYaml []byte

var once sync.Once
var connector base.IConnector

type Connector struct {
	base.BaseConnector
}

type Connection struct {
	base.BaseConnection
	config *structpb.Struct
}

func Init(logger *zap.Logger) base.IConnector {
	once.Do(func() {
		connDefs := []*connectorPB.SourceConnectorDefinition{}

		loader := configLoader.InitJSONSchema(logger)
		loader.Load(sourceDefinitionsYaml, sourceSpecsYaml, &connDefs)

		connector = &Connector{
			BaseConnector: base.BaseConnector{Logger: logger},
		}
		for idx := range connDefs {
			connector.AddConnectorDefinition(uuid.FromStringOrNil(connDefs[idx].GetUid()), connDefs[idx].GetId(), connDefs[idx])
		}
	})
	return connector
}

func (c *Connector) CreateConnection(defUid uuid.UUID, config *structpb.Struct, logger *zap.Logger) (base.IConnection, error) {
	return &Connection{
		BaseConnection: base.BaseConnection{Logger: logger},
		config:         config,
	}, nil
}

func (con *Connection) Execute(input interface{}) (interface{}, error) {
	return input, nil
}

func (con *Connection) Test() (connectorPB.Connector_State, error) {
	// Always connected
	return connectorPB.Connector_STATE_CONNECTED, nil
}
