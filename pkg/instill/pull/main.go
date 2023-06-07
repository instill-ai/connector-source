package pull

import (
	"sync"

	"github.com/gofrs/uuid"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/instill-ai/connector/pkg/base"

	connectorPB "github.com/instill-ai/protogen-go/vdp/connector/v1alpha"
)

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
		connector = &Connector{
			BaseConnector: base.BaseConnector{Logger: logger},
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
	return connectorPB.Connector_STATE_UNSPECIFIED, nil
}
