// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: mq.proto

package mq_pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	SeaweedMessaging_FindBrokerLeader_FullMethodName           = "/messaging_pb.SeaweedMessaging/FindBrokerLeader"
	SeaweedMessaging_PublisherToPubBalancer_FullMethodName     = "/messaging_pb.SeaweedMessaging/PublisherToPubBalancer"
	SeaweedMessaging_BalanceTopics_FullMethodName              = "/messaging_pb.SeaweedMessaging/BalanceTopics"
	SeaweedMessaging_ListTopics_FullMethodName                 = "/messaging_pb.SeaweedMessaging/ListTopics"
	SeaweedMessaging_ConfigureTopic_FullMethodName             = "/messaging_pb.SeaweedMessaging/ConfigureTopic"
	SeaweedMessaging_LookupTopicBrokers_FullMethodName         = "/messaging_pb.SeaweedMessaging/LookupTopicBrokers"
	SeaweedMessaging_AssignTopicPartitions_FullMethodName      = "/messaging_pb.SeaweedMessaging/AssignTopicPartitions"
	SeaweedMessaging_ClosePublishers_FullMethodName            = "/messaging_pb.SeaweedMessaging/ClosePublishers"
	SeaweedMessaging_CloseSubscribers_FullMethodName           = "/messaging_pb.SeaweedMessaging/CloseSubscribers"
	SeaweedMessaging_SubscriberToSubCoordinator_FullMethodName = "/messaging_pb.SeaweedMessaging/SubscriberToSubCoordinator"
	SeaweedMessaging_Publish_FullMethodName                    = "/messaging_pb.SeaweedMessaging/Publish"
	SeaweedMessaging_Subscribe_FullMethodName                  = "/messaging_pb.SeaweedMessaging/Subscribe"
)

// SeaweedMessagingClient is the client API for SeaweedMessaging service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SeaweedMessagingClient interface {
	// control plane
	FindBrokerLeader(ctx context.Context, in *FindBrokerLeaderRequest, opts ...grpc.CallOption) (*FindBrokerLeaderResponse, error)
	// control plane for balancer
	PublisherToPubBalancer(ctx context.Context, opts ...grpc.CallOption) (SeaweedMessaging_PublisherToPubBalancerClient, error)
	BalanceTopics(ctx context.Context, in *BalanceTopicsRequest, opts ...grpc.CallOption) (*BalanceTopicsResponse, error)
	// control plane for topic partitions
	ListTopics(ctx context.Context, in *ListTopicsRequest, opts ...grpc.CallOption) (*ListTopicsResponse, error)
	ConfigureTopic(ctx context.Context, in *ConfigureTopicRequest, opts ...grpc.CallOption) (*ConfigureTopicResponse, error)
	LookupTopicBrokers(ctx context.Context, in *LookupTopicBrokersRequest, opts ...grpc.CallOption) (*LookupTopicBrokersResponse, error)
	// invoked by the balancer, running on each broker
	AssignTopicPartitions(ctx context.Context, in *AssignTopicPartitionsRequest, opts ...grpc.CallOption) (*AssignTopicPartitionsResponse, error)
	ClosePublishers(ctx context.Context, in *ClosePublishersRequest, opts ...grpc.CallOption) (*ClosePublishersResponse, error)
	CloseSubscribers(ctx context.Context, in *CloseSubscribersRequest, opts ...grpc.CallOption) (*CloseSubscribersResponse, error)
	// subscriber connects to broker balancer, which coordinates with the subscribers
	SubscriberToSubCoordinator(ctx context.Context, opts ...grpc.CallOption) (SeaweedMessaging_SubscriberToSubCoordinatorClient, error)
	// data plane for each topic partition
	Publish(ctx context.Context, opts ...grpc.CallOption) (SeaweedMessaging_PublishClient, error)
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (SeaweedMessaging_SubscribeClient, error)
}

type seaweedMessagingClient struct {
	cc grpc.ClientConnInterface
}

func NewSeaweedMessagingClient(cc grpc.ClientConnInterface) SeaweedMessagingClient {
	return &seaweedMessagingClient{cc}
}

func (c *seaweedMessagingClient) FindBrokerLeader(ctx context.Context, in *FindBrokerLeaderRequest, opts ...grpc.CallOption) (*FindBrokerLeaderResponse, error) {
	out := new(FindBrokerLeaderResponse)
	err := c.cc.Invoke(ctx, SeaweedMessaging_FindBrokerLeader_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedMessagingClient) PublisherToPubBalancer(ctx context.Context, opts ...grpc.CallOption) (SeaweedMessaging_PublisherToPubBalancerClient, error) {
	stream, err := c.cc.NewStream(ctx, &SeaweedMessaging_ServiceDesc.Streams[0], SeaweedMessaging_PublisherToPubBalancer_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &seaweedMessagingPublisherToPubBalancerClient{stream}
	return x, nil
}

type SeaweedMessaging_PublisherToPubBalancerClient interface {
	Send(*PublisherToPubBalancerRequest) error
	Recv() (*PublisherToPubBalancerResponse, error)
	grpc.ClientStream
}

type seaweedMessagingPublisherToPubBalancerClient struct {
	grpc.ClientStream
}

func (x *seaweedMessagingPublisherToPubBalancerClient) Send(m *PublisherToPubBalancerRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *seaweedMessagingPublisherToPubBalancerClient) Recv() (*PublisherToPubBalancerResponse, error) {
	m := new(PublisherToPubBalancerResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *seaweedMessagingClient) BalanceTopics(ctx context.Context, in *BalanceTopicsRequest, opts ...grpc.CallOption) (*BalanceTopicsResponse, error) {
	out := new(BalanceTopicsResponse)
	err := c.cc.Invoke(ctx, SeaweedMessaging_BalanceTopics_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedMessagingClient) ListTopics(ctx context.Context, in *ListTopicsRequest, opts ...grpc.CallOption) (*ListTopicsResponse, error) {
	out := new(ListTopicsResponse)
	err := c.cc.Invoke(ctx, SeaweedMessaging_ListTopics_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedMessagingClient) ConfigureTopic(ctx context.Context, in *ConfigureTopicRequest, opts ...grpc.CallOption) (*ConfigureTopicResponse, error) {
	out := new(ConfigureTopicResponse)
	err := c.cc.Invoke(ctx, SeaweedMessaging_ConfigureTopic_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedMessagingClient) LookupTopicBrokers(ctx context.Context, in *LookupTopicBrokersRequest, opts ...grpc.CallOption) (*LookupTopicBrokersResponse, error) {
	out := new(LookupTopicBrokersResponse)
	err := c.cc.Invoke(ctx, SeaweedMessaging_LookupTopicBrokers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedMessagingClient) AssignTopicPartitions(ctx context.Context, in *AssignTopicPartitionsRequest, opts ...grpc.CallOption) (*AssignTopicPartitionsResponse, error) {
	out := new(AssignTopicPartitionsResponse)
	err := c.cc.Invoke(ctx, SeaweedMessaging_AssignTopicPartitions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedMessagingClient) ClosePublishers(ctx context.Context, in *ClosePublishersRequest, opts ...grpc.CallOption) (*ClosePublishersResponse, error) {
	out := new(ClosePublishersResponse)
	err := c.cc.Invoke(ctx, SeaweedMessaging_ClosePublishers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedMessagingClient) CloseSubscribers(ctx context.Context, in *CloseSubscribersRequest, opts ...grpc.CallOption) (*CloseSubscribersResponse, error) {
	out := new(CloseSubscribersResponse)
	err := c.cc.Invoke(ctx, SeaweedMessaging_CloseSubscribers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedMessagingClient) SubscriberToSubCoordinator(ctx context.Context, opts ...grpc.CallOption) (SeaweedMessaging_SubscriberToSubCoordinatorClient, error) {
	stream, err := c.cc.NewStream(ctx, &SeaweedMessaging_ServiceDesc.Streams[1], SeaweedMessaging_SubscriberToSubCoordinator_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &seaweedMessagingSubscriberToSubCoordinatorClient{stream}
	return x, nil
}

type SeaweedMessaging_SubscriberToSubCoordinatorClient interface {
	Send(*SubscriberToSubCoordinatorRequest) error
	Recv() (*SubscriberToSubCoordinatorResponse, error)
	grpc.ClientStream
}

type seaweedMessagingSubscriberToSubCoordinatorClient struct {
	grpc.ClientStream
}

func (x *seaweedMessagingSubscriberToSubCoordinatorClient) Send(m *SubscriberToSubCoordinatorRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *seaweedMessagingSubscriberToSubCoordinatorClient) Recv() (*SubscriberToSubCoordinatorResponse, error) {
	m := new(SubscriberToSubCoordinatorResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *seaweedMessagingClient) Publish(ctx context.Context, opts ...grpc.CallOption) (SeaweedMessaging_PublishClient, error) {
	stream, err := c.cc.NewStream(ctx, &SeaweedMessaging_ServiceDesc.Streams[2], SeaweedMessaging_Publish_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &seaweedMessagingPublishClient{stream}
	return x, nil
}

type SeaweedMessaging_PublishClient interface {
	Send(*PublishRequest) error
	Recv() (*PublishResponse, error)
	grpc.ClientStream
}

type seaweedMessagingPublishClient struct {
	grpc.ClientStream
}

func (x *seaweedMessagingPublishClient) Send(m *PublishRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *seaweedMessagingPublishClient) Recv() (*PublishResponse, error) {
	m := new(PublishResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *seaweedMessagingClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (SeaweedMessaging_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &SeaweedMessaging_ServiceDesc.Streams[3], SeaweedMessaging_Subscribe_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &seaweedMessagingSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SeaweedMessaging_SubscribeClient interface {
	Recv() (*SubscribeResponse, error)
	grpc.ClientStream
}

type seaweedMessagingSubscribeClient struct {
	grpc.ClientStream
}

func (x *seaweedMessagingSubscribeClient) Recv() (*SubscribeResponse, error) {
	m := new(SubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SeaweedMessagingServer is the server API for SeaweedMessaging service.
// All implementations must embed UnimplementedSeaweedMessagingServer
// for forward compatibility
type SeaweedMessagingServer interface {
	// control plane
	FindBrokerLeader(context.Context, *FindBrokerLeaderRequest) (*FindBrokerLeaderResponse, error)
	// control plane for balancer
	PublisherToPubBalancer(SeaweedMessaging_PublisherToPubBalancerServer) error
	BalanceTopics(context.Context, *BalanceTopicsRequest) (*BalanceTopicsResponse, error)
	// control plane for topic partitions
	ListTopics(context.Context, *ListTopicsRequest) (*ListTopicsResponse, error)
	ConfigureTopic(context.Context, *ConfigureTopicRequest) (*ConfigureTopicResponse, error)
	LookupTopicBrokers(context.Context, *LookupTopicBrokersRequest) (*LookupTopicBrokersResponse, error)
	// invoked by the balancer, running on each broker
	AssignTopicPartitions(context.Context, *AssignTopicPartitionsRequest) (*AssignTopicPartitionsResponse, error)
	ClosePublishers(context.Context, *ClosePublishersRequest) (*ClosePublishersResponse, error)
	CloseSubscribers(context.Context, *CloseSubscribersRequest) (*CloseSubscribersResponse, error)
	// subscriber connects to broker balancer, which coordinates with the subscribers
	SubscriberToSubCoordinator(SeaweedMessaging_SubscriberToSubCoordinatorServer) error
	// data plane for each topic partition
	Publish(SeaweedMessaging_PublishServer) error
	Subscribe(*SubscribeRequest, SeaweedMessaging_SubscribeServer) error
	mustEmbedUnimplementedSeaweedMessagingServer()
}

// UnimplementedSeaweedMessagingServer must be embedded to have forward compatible implementations.
type UnimplementedSeaweedMessagingServer struct {
}

func (UnimplementedSeaweedMessagingServer) FindBrokerLeader(context.Context, *FindBrokerLeaderRequest) (*FindBrokerLeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindBrokerLeader not implemented")
}
func (UnimplementedSeaweedMessagingServer) PublisherToPubBalancer(SeaweedMessaging_PublisherToPubBalancerServer) error {
	return status.Errorf(codes.Unimplemented, "method PublisherToPubBalancer not implemented")
}
func (UnimplementedSeaweedMessagingServer) BalanceTopics(context.Context, *BalanceTopicsRequest) (*BalanceTopicsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BalanceTopics not implemented")
}
func (UnimplementedSeaweedMessagingServer) ListTopics(context.Context, *ListTopicsRequest) (*ListTopicsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTopics not implemented")
}
func (UnimplementedSeaweedMessagingServer) ConfigureTopic(context.Context, *ConfigureTopicRequest) (*ConfigureTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigureTopic not implemented")
}
func (UnimplementedSeaweedMessagingServer) LookupTopicBrokers(context.Context, *LookupTopicBrokersRequest) (*LookupTopicBrokersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookupTopicBrokers not implemented")
}
func (UnimplementedSeaweedMessagingServer) AssignTopicPartitions(context.Context, *AssignTopicPartitionsRequest) (*AssignTopicPartitionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignTopicPartitions not implemented")
}
func (UnimplementedSeaweedMessagingServer) ClosePublishers(context.Context, *ClosePublishersRequest) (*ClosePublishersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClosePublishers not implemented")
}
func (UnimplementedSeaweedMessagingServer) CloseSubscribers(context.Context, *CloseSubscribersRequest) (*CloseSubscribersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseSubscribers not implemented")
}
func (UnimplementedSeaweedMessagingServer) SubscriberToSubCoordinator(SeaweedMessaging_SubscriberToSubCoordinatorServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscriberToSubCoordinator not implemented")
}
func (UnimplementedSeaweedMessagingServer) Publish(SeaweedMessaging_PublishServer) error {
	return status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedSeaweedMessagingServer) Subscribe(*SubscribeRequest, SeaweedMessaging_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedSeaweedMessagingServer) mustEmbedUnimplementedSeaweedMessagingServer() {}

// UnsafeSeaweedMessagingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SeaweedMessagingServer will
// result in compilation errors.
type UnsafeSeaweedMessagingServer interface {
	mustEmbedUnimplementedSeaweedMessagingServer()
}

func RegisterSeaweedMessagingServer(s grpc.ServiceRegistrar, srv SeaweedMessagingServer) {
	s.RegisterService(&SeaweedMessaging_ServiceDesc, srv)
}

func _SeaweedMessaging_FindBrokerLeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindBrokerLeaderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedMessagingServer).FindBrokerLeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SeaweedMessaging_FindBrokerLeader_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedMessagingServer).FindBrokerLeader(ctx, req.(*FindBrokerLeaderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedMessaging_PublisherToPubBalancer_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SeaweedMessagingServer).PublisherToPubBalancer(&seaweedMessagingPublisherToPubBalancerServer{stream})
}

type SeaweedMessaging_PublisherToPubBalancerServer interface {
	Send(*PublisherToPubBalancerResponse) error
	Recv() (*PublisherToPubBalancerRequest, error)
	grpc.ServerStream
}

type seaweedMessagingPublisherToPubBalancerServer struct {
	grpc.ServerStream
}

func (x *seaweedMessagingPublisherToPubBalancerServer) Send(m *PublisherToPubBalancerResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *seaweedMessagingPublisherToPubBalancerServer) Recv() (*PublisherToPubBalancerRequest, error) {
	m := new(PublisherToPubBalancerRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SeaweedMessaging_BalanceTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceTopicsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedMessagingServer).BalanceTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SeaweedMessaging_BalanceTopics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedMessagingServer).BalanceTopics(ctx, req.(*BalanceTopicsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedMessaging_ListTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTopicsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedMessagingServer).ListTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SeaweedMessaging_ListTopics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedMessagingServer).ListTopics(ctx, req.(*ListTopicsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedMessaging_ConfigureTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedMessagingServer).ConfigureTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SeaweedMessaging_ConfigureTopic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedMessagingServer).ConfigureTopic(ctx, req.(*ConfigureTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedMessaging_LookupTopicBrokers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupTopicBrokersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedMessagingServer).LookupTopicBrokers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SeaweedMessaging_LookupTopicBrokers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedMessagingServer).LookupTopicBrokers(ctx, req.(*LookupTopicBrokersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedMessaging_AssignTopicPartitions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignTopicPartitionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedMessagingServer).AssignTopicPartitions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SeaweedMessaging_AssignTopicPartitions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedMessagingServer).AssignTopicPartitions(ctx, req.(*AssignTopicPartitionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedMessaging_ClosePublishers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClosePublishersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedMessagingServer).ClosePublishers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SeaweedMessaging_ClosePublishers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedMessagingServer).ClosePublishers(ctx, req.(*ClosePublishersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedMessaging_CloseSubscribers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseSubscribersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedMessagingServer).CloseSubscribers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SeaweedMessaging_CloseSubscribers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedMessagingServer).CloseSubscribers(ctx, req.(*CloseSubscribersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedMessaging_SubscriberToSubCoordinator_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SeaweedMessagingServer).SubscriberToSubCoordinator(&seaweedMessagingSubscriberToSubCoordinatorServer{stream})
}

type SeaweedMessaging_SubscriberToSubCoordinatorServer interface {
	Send(*SubscriberToSubCoordinatorResponse) error
	Recv() (*SubscriberToSubCoordinatorRequest, error)
	grpc.ServerStream
}

type seaweedMessagingSubscriberToSubCoordinatorServer struct {
	grpc.ServerStream
}

func (x *seaweedMessagingSubscriberToSubCoordinatorServer) Send(m *SubscriberToSubCoordinatorResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *seaweedMessagingSubscriberToSubCoordinatorServer) Recv() (*SubscriberToSubCoordinatorRequest, error) {
	m := new(SubscriberToSubCoordinatorRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SeaweedMessaging_Publish_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SeaweedMessagingServer).Publish(&seaweedMessagingPublishServer{stream})
}

type SeaweedMessaging_PublishServer interface {
	Send(*PublishResponse) error
	Recv() (*PublishRequest, error)
	grpc.ServerStream
}

type seaweedMessagingPublishServer struct {
	grpc.ServerStream
}

func (x *seaweedMessagingPublishServer) Send(m *PublishResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *seaweedMessagingPublishServer) Recv() (*PublishRequest, error) {
	m := new(PublishRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SeaweedMessaging_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SeaweedMessagingServer).Subscribe(m, &seaweedMessagingSubscribeServer{stream})
}

type SeaweedMessaging_SubscribeServer interface {
	Send(*SubscribeResponse) error
	grpc.ServerStream
}

type seaweedMessagingSubscribeServer struct {
	grpc.ServerStream
}

func (x *seaweedMessagingSubscribeServer) Send(m *SubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

// SeaweedMessaging_ServiceDesc is the grpc.ServiceDesc for SeaweedMessaging service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SeaweedMessaging_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "messaging_pb.SeaweedMessaging",
	HandlerType: (*SeaweedMessagingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindBrokerLeader",
			Handler:    _SeaweedMessaging_FindBrokerLeader_Handler,
		},
		{
			MethodName: "BalanceTopics",
			Handler:    _SeaweedMessaging_BalanceTopics_Handler,
		},
		{
			MethodName: "ListTopics",
			Handler:    _SeaweedMessaging_ListTopics_Handler,
		},
		{
			MethodName: "ConfigureTopic",
			Handler:    _SeaweedMessaging_ConfigureTopic_Handler,
		},
		{
			MethodName: "LookupTopicBrokers",
			Handler:    _SeaweedMessaging_LookupTopicBrokers_Handler,
		},
		{
			MethodName: "AssignTopicPartitions",
			Handler:    _SeaweedMessaging_AssignTopicPartitions_Handler,
		},
		{
			MethodName: "ClosePublishers",
			Handler:    _SeaweedMessaging_ClosePublishers_Handler,
		},
		{
			MethodName: "CloseSubscribers",
			Handler:    _SeaweedMessaging_CloseSubscribers_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PublisherToPubBalancer",
			Handler:       _SeaweedMessaging_PublisherToPubBalancer_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SubscriberToSubCoordinator",
			Handler:       _SeaweedMessaging_SubscriberToSubCoordinator_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Publish",
			Handler:       _SeaweedMessaging_Publish_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Subscribe",
			Handler:       _SeaweedMessaging_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "mq.proto",
}
