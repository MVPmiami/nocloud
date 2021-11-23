/*
Copyright © 2021 Nikita Ivanovski info@slnt-opp.xyz

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package graph

import (
	"context"
	"errors"

	"github.com/arangodb/go-driver"
	pb "github.com/slntopp/nocloud/pkg/services/proto"
	"go.uber.org/zap"
)

const (
	SERVICES_COL = "Services"
	NS2SERV = NAMESPACES_COL + "2" + SERVICES_COL
)

type Service struct {
	*pb.Service
	driver.DocumentMeta
}

type ServicesController struct {
	log *zap.Logger

	col driver.Collection // Services Collection
	ig_ctrl InstancesGroupsController

	db driver.Database
}

func NewServicesController(log *zap.Logger, db driver.Database) ServicesController {
	col, _ := db.Collection(nil, SERVICES_COL)
	return ServicesController{log: log, col: col, ig_ctrl: NewInstancesGroupsController(log, db), db:db}
}

// Create Service and underlaying entities and store in DB
func (ctrl *ServicesController) Create(ctx context.Context, service *pb.Service) (*Service, error) {
	ctrl.log.Debug("Creating Service", zap.Any("service", service))
	for _, ig := range service.GetInstancesGroups() {
		err := ctrl.ig_ctrl.Create(ctx, ig)
		if err != nil {
			return nil, err
		}
	}
	meta, err := ctrl.col.CreateDocument(ctx, service)
	if err != nil {
		ctrl.log.Debug("Error creating document", zap.Error(err))
		return nil, errors.New("Error creating document")
	}
	service.Uuid = meta.ID.Key()
	return &Service{service, meta}, nil
}

// Join Service into Namespace
func (ctrl *ServicesController) Join(ctx context.Context, service *Service, namespace *Namespace, access int32, role string) (error) {
	ctrl.log.Debug("Joining service to namespace")
	edge, _ := ctrl.db.Collection(ctx, NS2SERV)
	_, err := edge.CreateDocument(ctx, Access{
		From: namespace.ID,
		To: service.ID,
		Level: access,
		Role: role,
	})
	return err
}