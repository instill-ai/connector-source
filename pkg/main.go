package source

import (
	"fmt"
	"sync"

	"github.com/gofrs/uuid"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/instill-ai/connector-source/pkg/instill"
	"github.com/instill-ai/connector/pkg/base"
)

var once sync.Once
var connector base.IConnector

type Connector struct {
	base.BaseConnector
	instillConnector base.IConnector
}

type Connection struct {
}

func Init(logger *zap.Logger) base.IConnector {
	once.Do(func() {

		instillConnector := instill.Init(logger)
		connector = &Connector{
			BaseConnector:    base.BaseConnector{Logger: logger},
			instillConnector: instillConnector,
		}

		// TODO: assert no duplicate uid
		// Note: we preserve the order as yaml
		for _, uid := range instillConnector.ListConnectorDefinitionUids() {
			def, err := instillConnector.GetConnectorDefinitionByUid(uid)
			if err != nil {
				logger.Error(err.Error())
			}
			connector.AddConnectorDefinition(uid, def.GetId(), def)
		}

	})
	return connector
}

func (c *Connector) CreateConnection(defUid uuid.UUID, config *structpb.Struct, logger *zap.Logger) (base.IConnection, error) {
	switch {
	case c.instillConnector.HasUid(defUid):
		return c.instillConnector.CreateConnection(defUid, config, logger)

	default:
		return nil, fmt.Errorf("no sourceConnector uid: %s", defUid)
	}
}
