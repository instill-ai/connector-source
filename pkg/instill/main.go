package instill

import (
	"fmt"
	"sync"

	"github.com/gofrs/uuid"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/instill-ai/connector-source/pkg/instill/request"
	"github.com/instill-ai/connector/pkg/base"
)

var once sync.Once
var connector base.IConnector

type Connector struct {
	base.BaseConnector
	instillRequestConnector base.IConnector
	// instillPullConnector    base.IConnector
}

func Init(logger *zap.Logger) base.IConnector {
	once.Do(func() {

		instillRequestConnector := request.Init(logger)
		// instillPullConnector := pull.Init(logger)

		connector = &Connector{
			BaseConnector:           base.BaseConnector{Logger: logger},
			instillRequestConnector: instillRequestConnector,
			// instillPullConnector:    instillPullConnector,
		}

		// TODO: assert no duplicate uid
		// Note: we preserve the order as yaml
		for _, uid := range instillRequestConnector.ListConnectorDefinitionUids() {
			def, err := instillRequestConnector.GetConnectorDefinitionByUid(uid)
			if err != nil {
				logger.Error(err.Error())
			}
			connector.AddConnectorDefinition(uid, def.GetId(), def)
		}
		// for _, uid := range instillPullConnector.ListConnectorDefinitionUids() {
		// 	def, err := instillPullConnector.GetConnectorDefinitionByUid(uid)
		// 	if err != nil {
		// 		logger.Error(err.Error())
		// 	}
		// 	connector.AddConnectorDefinition(uid, def.GetId(), def)
		// }
	})
	return connector
}

func (c *Connector) CreateConnection(defUid uuid.UUID, config *structpb.Struct, logger *zap.Logger) (base.IConnection, error) {
	switch {
	case c.instillRequestConnector.HasUid(defUid):
		return c.instillRequestConnector.CreateConnection(defUid, config, logger)
	// case c.instillPullConnector.HasUid(defUid):
	// 	return c.instillPullConnector.CreateConnection(defUid, config, logger)
	default:
		return nil, fmt.Errorf("no sourceConnector uid: %s", defUid)
	}
}
