/*
Copyright © 2021-2022 Nikita Ivanovski info@slnt-opp.xyz

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
package settings

import (
	"context"
	"fmt"

	redis "github.com/go-redis/redis/v8"
	"github.com/slntopp/nocloud/pkg/nocloud"
	pb "github.com/slntopp/nocloud/pkg/settings/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"

	strcase "github.com/stoewer/go-strcase"

	"go.uber.org/zap"
)

const KEYS_PREFIX = "_settings"

var KEY_NS_PATTERN = fmt.Sprintf("%s:*", KEYS_PREFIX)

type SettingsServiceServer struct {
	pb.UnimplementedSettingsServiceServer

	log *zap.Logger
	rdb *redis.Client
}

func NewSettingsServer(log *zap.Logger, rdb *redis.Client) *SettingsServiceServer {
	return &SettingsServiceServer{
		log: log.Named("SettingsServer"), rdb: rdb,
	}
}

func (s *SettingsServiceServer) Get(ctx context.Context, req *pb.GetRequest) (*structpb.Struct, error) {
	result := make(map[string]interface{})

	for _, key := range req.GetKeys() {
		r := s.rdb.HGet(ctx, strcase.LowerCamelCase(key), "value")
		result[key], _ = r.Result()
	}

	res, err := structpb.NewStruct(result)
	if err != nil {
		s.log.Error("Error serializing map to Struct", zap.Error(err))
		return nil, status.Error(codes.Internal, "Error serializing map to Struct")
	}
	return res, nil
}

func (s *SettingsServiceServer) Put(ctx context.Context, req *pb.PutRequest) (*pb.PutResponse, error) {
	access := ctx.Value(nocloud.NoCloudRootAccess).(int32)
	if access < 3 {
		return nil, status.Error(codes.PermissionDenied, "Not enough access rights")
	}

	key := fmt.Sprintf("%s:%s", KEYS_PREFIX, strcase.LowerCamelCase(req.GetKey()))
	r := s.rdb.HSet(ctx, key, "value", req.GetValue(), 
		"desc", req.GetDescription(), "pub", req.GetPublic())
	_, err := r.Result()
	if err != nil {
		s.log.Error("Error allocating keys in Redis", zap.String("key", key), zap.Error(err))
		return nil, status.Error(codes.Internal, "Error allocating keys in Redis")
	}

	go func() {
		s.rdb.Publish(ctx, key, req.GetValue())
	}()

	return &pb.PutResponse{Key: req.GetKey()}, nil
}

func (s *SettingsServiceServer) Sub(req *pb.SubRequest, stream pb.SettingsService_SubServer) error {
	sub := s.rdb.Subscribe(context.TODO(), strcase.LowerCamelCase(req.GetKey()))
	defer sub.Close()
	for {
		select {
		case msg := <- sub.Channel():
			stream.Send(&pb.SubRequest{
				Key: msg.Channel,
				Value: msg.Payload,
			})
		case <- stream.Context().Done():
			return nil
		}
	}
}

func (s *SettingsServiceServer) Keys(ctx context.Context, _ *pb.KeysRequest) (*pb.KeysResponse, error) {
	r := s.rdb.Keys(ctx, KEY_NS_PATTERN)
	keys, err := r.Result()
	if err != nil {
		s.log.Error("Error getting keys from Redis", zap.Error(err))
		return nil, status.Error(codes.Internal, "Error getting keys from Redis")
	}
	
	result := make([]*pb.KeysResponse_Key, len(keys))
	for i, key := range keys {
		r := s.rdb.HGetAll(ctx, key)
		data, err := r.Result()
		if err != nil {
			result[i] = &pb.KeysResponse_Key{Key: strcase.KebabCase(key), Description: "Unresolved"}
			continue
		}
		result[i] = &pb.KeysResponse_Key{
			Key: strcase.KebabCase(key),
			Description: data["desc"],
			Public: data["pub"] == "1",
		}
	}

	return &pb.KeysResponse{Pool: result}, nil
}