// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/endpoint/load_report.proto

package endpoint

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/core"
import _ "github.com/gogo/protobuf/gogoproto"
import duration "github.com/golang/protobuf/ptypes/duration"
import _ "github.com/lyft/protoc-gen-validate/validate"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// These are stats Envoy reports to GLB every so often. Report frequency is
// defined by
// :ref:`LoadStatsResponse.load_reporting_interval<envoy_api_field_load_stats.LoadStatsResponse.load_reporting_interval>`.
// Stats per upstream region/zone and optionally per subzone.
// [#not-implemented-hide:] Not configuration. TBD how to doc proto APIs.
type UpstreamLocalityStats struct {
	// Name of zone, region and optionally endpoint group these metrics were
	// collected from. Zone and region names could be empty if unknown.
	Locality *core.Locality `protobuf:"bytes,1,opt,name=locality,proto3" json:"locality,omitempty"`
	// The total number of requests sent by this Envoy since the last report. A
	// single HTTP or gRPC request or stream is counted as one request. A TCP
	// connection is also treated as one request. There is no explicit
	// total_requests field below for a locality, but it may be inferred from:
	//
	// .. code-block:: none
	//
	//   total_requests = total_successful_requests + total_requests_in_progress +
	//     total_error_requests
	//
	// The total number of requests successfully completed by the endpoints in the
	// locality. These include non-5xx responses for HTTP, where errors
	// originate at the client and the endpoint responded successfully. For gRPC,
	// the grpc-status values are those not covered by total_error_requests below.
	TotalSuccessfulRequests uint64 `protobuf:"varint,2,opt,name=total_successful_requests,json=totalSuccessfulRequests,proto3" json:"total_successful_requests,omitempty"`
	// The total number of unfinished requests
	TotalRequestsInProgress uint64 `protobuf:"varint,3,opt,name=total_requests_in_progress,json=totalRequestsInProgress,proto3" json:"total_requests_in_progress,omitempty"`
	// The total number of requests that failed due to errors at the endpoint.
	// For HTTP these are responses with 5xx status codes and for gRPC the
	// grpc-status values:
	//
	//   - DeadlineExceeded
	//   - Unimplemented
	//   - Internal
	//   - Unavailable
	//   - Unknown
	//   - DataLoss
	TotalErrorRequests uint64 `protobuf:"varint,4,opt,name=total_error_requests,json=totalErrorRequests,proto3" json:"total_error_requests,omitempty"`
	// Stats for multi-dimensional load balancing.
	LoadMetricStats []*EndpointLoadMetricStats `protobuf:"bytes,5,rep,name=load_metric_stats,json=loadMetricStats,proto3" json:"load_metric_stats,omitempty"`
	// [#not-implemented-hide:] The priority of the endpoint group these metrics
	// were collected from.
	Priority             uint32   `protobuf:"varint,6,opt,name=priority,proto3" json:"priority,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpstreamLocalityStats) Reset()         { *m = UpstreamLocalityStats{} }
func (m *UpstreamLocalityStats) String() string { return proto.CompactTextString(m) }
func (*UpstreamLocalityStats) ProtoMessage()    {}
func (*UpstreamLocalityStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_load_report_270597675435f766, []int{0}
}
func (m *UpstreamLocalityStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpstreamLocalityStats.Unmarshal(m, b)
}
func (m *UpstreamLocalityStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpstreamLocalityStats.Marshal(b, m, deterministic)
}
func (dst *UpstreamLocalityStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpstreamLocalityStats.Merge(dst, src)
}
func (m *UpstreamLocalityStats) XXX_Size() int {
	return xxx_messageInfo_UpstreamLocalityStats.Size(m)
}
func (m *UpstreamLocalityStats) XXX_DiscardUnknown() {
	xxx_messageInfo_UpstreamLocalityStats.DiscardUnknown(m)
}

var xxx_messageInfo_UpstreamLocalityStats proto.InternalMessageInfo

func (m *UpstreamLocalityStats) GetLocality() *core.Locality {
	if m != nil {
		return m.Locality
	}
	return nil
}

func (m *UpstreamLocalityStats) GetTotalSuccessfulRequests() uint64 {
	if m != nil {
		return m.TotalSuccessfulRequests
	}
	return 0
}

func (m *UpstreamLocalityStats) GetTotalRequestsInProgress() uint64 {
	if m != nil {
		return m.TotalRequestsInProgress
	}
	return 0
}

func (m *UpstreamLocalityStats) GetTotalErrorRequests() uint64 {
	if m != nil {
		return m.TotalErrorRequests
	}
	return 0
}

func (m *UpstreamLocalityStats) GetLoadMetricStats() []*EndpointLoadMetricStats {
	if m != nil {
		return m.LoadMetricStats
	}
	return nil
}

func (m *UpstreamLocalityStats) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

// [#not-implemented-hide:] Not configuration. TBD how to doc proto APIs.
type EndpointLoadMetricStats struct {
	// Name of the metric; may be empty.
	MetricName string `protobuf:"bytes,1,opt,name=metric_name,json=metricName,proto3" json:"metric_name,omitempty"`
	// Number of calls that finished and included this metric.
	NumRequestsFinishedWithMetric uint64 `protobuf:"varint,2,opt,name=num_requests_finished_with_metric,json=numRequestsFinishedWithMetric,proto3" json:"num_requests_finished_with_metric,omitempty"`
	// Sum of metric values across all calls that finished with this metric for
	// load_reporting_interval.
	TotalMetricValue     float64  `protobuf:"fixed64,3,opt,name=total_metric_value,json=totalMetricValue,proto3" json:"total_metric_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EndpointLoadMetricStats) Reset()         { *m = EndpointLoadMetricStats{} }
func (m *EndpointLoadMetricStats) String() string { return proto.CompactTextString(m) }
func (*EndpointLoadMetricStats) ProtoMessage()    {}
func (*EndpointLoadMetricStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_load_report_270597675435f766, []int{1}
}
func (m *EndpointLoadMetricStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndpointLoadMetricStats.Unmarshal(m, b)
}
func (m *EndpointLoadMetricStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndpointLoadMetricStats.Marshal(b, m, deterministic)
}
func (dst *EndpointLoadMetricStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndpointLoadMetricStats.Merge(dst, src)
}
func (m *EndpointLoadMetricStats) XXX_Size() int {
	return xxx_messageInfo_EndpointLoadMetricStats.Size(m)
}
func (m *EndpointLoadMetricStats) XXX_DiscardUnknown() {
	xxx_messageInfo_EndpointLoadMetricStats.DiscardUnknown(m)
}

var xxx_messageInfo_EndpointLoadMetricStats proto.InternalMessageInfo

func (m *EndpointLoadMetricStats) GetMetricName() string {
	if m != nil {
		return m.MetricName
	}
	return ""
}

func (m *EndpointLoadMetricStats) GetNumRequestsFinishedWithMetric() uint64 {
	if m != nil {
		return m.NumRequestsFinishedWithMetric
	}
	return 0
}

func (m *EndpointLoadMetricStats) GetTotalMetricValue() float64 {
	if m != nil {
		return m.TotalMetricValue
	}
	return 0
}

// Per cluster load stats. Envoy reports these stats a management server in a
// :ref:`LoadStatsRequest<envoy_api_msg_load_stats.LoadStatsRequest>`
// [#not-implemented-hide:] Not configuration. TBD how to doc proto APIs.
type ClusterStats struct {
	// The name of the cluster.
	ClusterName string `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	// Need at least one.
	UpstreamLocalityStats []*UpstreamLocalityStats `protobuf:"bytes,2,rep,name=upstream_locality_stats,json=upstreamLocalityStats,proto3" json:"upstream_locality_stats,omitempty"`
	// Cluster-level stats such as total_successful_requests may be computed by
	// summing upstream_locality_stats. In addition, below there are additional
	// cluster-wide stats. The following total_requests equality holds at the
	// cluster-level:
	//
	// .. code-block:: none
	//
	//   sum_locality(total_successful_requests) + sum_locality(total_requests_in_progress) +
	//     sum_locality(total_error_requests) + total_dropped_requests`
	//
	// The total number of dropped requests. This covers requests
	// deliberately dropped by the drop_overload policy and circuit breaking.
	TotalDroppedRequests uint64 `protobuf:"varint,3,opt,name=total_dropped_requests,json=totalDroppedRequests,proto3" json:"total_dropped_requests,omitempty"`
	// Information about deliberately dropped requests for each category specified
	// in the DropOverload policy.
	DroppedRequests []*ClusterStats_DroppedRequests `protobuf:"bytes,5,rep,name=dropped_requests,json=droppedRequests,proto3" json:"dropped_requests,omitempty"`
	// Period over which the actual load report occurred. This will be guaranteed to include every
	// request reported. Due to system load and delays between the *LoadStatsRequest* sent from Envoy
	// and the *LoadStatsResponse* message sent from the management server, this may be longer than
	// the requested load reporting interval in the *LoadStatsResponse*.
	LoadReportInterval   *duration.Duration `protobuf:"bytes,4,opt,name=load_report_interval,json=loadReportInterval,proto3" json:"load_report_interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ClusterStats) Reset()         { *m = ClusterStats{} }
func (m *ClusterStats) String() string { return proto.CompactTextString(m) }
func (*ClusterStats) ProtoMessage()    {}
func (*ClusterStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_load_report_270597675435f766, []int{2}
}
func (m *ClusterStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterStats.Unmarshal(m, b)
}
func (m *ClusterStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterStats.Marshal(b, m, deterministic)
}
func (dst *ClusterStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterStats.Merge(dst, src)
}
func (m *ClusterStats) XXX_Size() int {
	return xxx_messageInfo_ClusterStats.Size(m)
}
func (m *ClusterStats) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterStats.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterStats proto.InternalMessageInfo

func (m *ClusterStats) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *ClusterStats) GetUpstreamLocalityStats() []*UpstreamLocalityStats {
	if m != nil {
		return m.UpstreamLocalityStats
	}
	return nil
}

func (m *ClusterStats) GetTotalDroppedRequests() uint64 {
	if m != nil {
		return m.TotalDroppedRequests
	}
	return 0
}

func (m *ClusterStats) GetDroppedRequests() []*ClusterStats_DroppedRequests {
	if m != nil {
		return m.DroppedRequests
	}
	return nil
}

func (m *ClusterStats) GetLoadReportInterval() *duration.Duration {
	if m != nil {
		return m.LoadReportInterval
	}
	return nil
}

type ClusterStats_DroppedRequests struct {
	// Identifier for the policy specifying the drop.
	Category string `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	// Total number of deliberately dropped requests for the category.
	DroppedCount         uint64   `protobuf:"varint,2,opt,name=dropped_count,json=droppedCount,proto3" json:"dropped_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterStats_DroppedRequests) Reset()         { *m = ClusterStats_DroppedRequests{} }
func (m *ClusterStats_DroppedRequests) String() string { return proto.CompactTextString(m) }
func (*ClusterStats_DroppedRequests) ProtoMessage()    {}
func (*ClusterStats_DroppedRequests) Descriptor() ([]byte, []int) {
	return fileDescriptor_load_report_270597675435f766, []int{2, 0}
}
func (m *ClusterStats_DroppedRequests) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterStats_DroppedRequests.Unmarshal(m, b)
}
func (m *ClusterStats_DroppedRequests) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterStats_DroppedRequests.Marshal(b, m, deterministic)
}
func (dst *ClusterStats_DroppedRequests) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterStats_DroppedRequests.Merge(dst, src)
}
func (m *ClusterStats_DroppedRequests) XXX_Size() int {
	return xxx_messageInfo_ClusterStats_DroppedRequests.Size(m)
}
func (m *ClusterStats_DroppedRequests) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterStats_DroppedRequests.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterStats_DroppedRequests proto.InternalMessageInfo

func (m *ClusterStats_DroppedRequests) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *ClusterStats_DroppedRequests) GetDroppedCount() uint64 {
	if m != nil {
		return m.DroppedCount
	}
	return 0
}

func init() {
	proto.RegisterType((*UpstreamLocalityStats)(nil), "envoy.api.v2.endpoint.UpstreamLocalityStats")
	proto.RegisterType((*EndpointLoadMetricStats)(nil), "envoy.api.v2.endpoint.EndpointLoadMetricStats")
	proto.RegisterType((*ClusterStats)(nil), "envoy.api.v2.endpoint.ClusterStats")
	proto.RegisterType((*ClusterStats_DroppedRequests)(nil), "envoy.api.v2.endpoint.ClusterStats.DroppedRequests")
}

func init() {
	proto.RegisterFile("envoy/api/v2/endpoint/load_report.proto", fileDescriptor_load_report_270597675435f766)
}

var fileDescriptor_load_report_270597675435f766 = []byte{
	// 612 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x18, 0x55, 0xda, 0x6d, 0xea, 0xdc, 0x4d, 0x1b, 0xd6, 0x46, 0xbb, 0xf2, 0x57, 0x86, 0x10, 0xbd,
	0x98, 0x1c, 0xd4, 0x21, 0x21, 0xc1, 0xdd, 0x7e, 0x10, 0x13, 0x03, 0x21, 0x4f, 0x80, 0x84, 0x04,
	0x91, 0x97, 0x78, 0x99, 0xa5, 0xc4, 0x0e, 0xfe, 0x09, 0xda, 0x6b, 0xf0, 0x28, 0x5c, 0x72, 0xc5,
	0xeb, 0x70, 0x87, 0xc4, 0x0b, 0xa0, 0xd8, 0x4e, 0xd6, 0x76, 0xed, 0x9d, 0xf3, 0x9d, 0xef, 0xd8,
	0xe7, 0x3b, 0x3e, 0x0e, 0x78, 0x42, 0x79, 0x29, 0xae, 0x42, 0x52, 0xb0, 0xb0, 0x1c, 0x87, 0x94,
	0x27, 0x85, 0x60, 0x5c, 0x87, 0x99, 0x20, 0x49, 0x24, 0x69, 0x21, 0xa4, 0x46, 0x85, 0x14, 0x5a,
	0xc0, 0x6d, 0xdb, 0x88, 0x48, 0xc1, 0x50, 0x39, 0x46, 0x75, 0xe3, 0xe0, 0xee, 0x14, 0x3f, 0x16,
	0x92, 0x86, 0xe7, 0x44, 0x51, 0x47, 0x1a, 0xdc, 0x4f, 0x85, 0x48, 0x33, 0x1a, 0xda, 0xaf, 0x73,
	0x73, 0x11, 0x26, 0x46, 0x12, 0xcd, 0x04, 0xf7, 0x78, 0xaf, 0x24, 0x19, 0x4b, 0x88, 0xa6, 0x61,
	0xbd, 0xf0, 0xc0, 0x56, 0x2a, 0x52, 0x61, 0x97, 0x61, 0xb5, 0x72, 0xd5, 0xdd, 0x7f, 0x2d, 0xb0,
	0xfd, 0xa1, 0x50, 0x5a, 0x52, 0x92, 0x9f, 0x8a, 0x98, 0x64, 0x4c, 0x5f, 0x9d, 0x69, 0xa2, 0x15,
	0x7c, 0x0e, 0x3a, 0x99, 0x2f, 0xf4, 0x83, 0x61, 0x30, 0xea, 0x8e, 0xef, 0xa0, 0x29, 0xc1, 0x95,
	0x32, 0x54, 0x73, 0x70, 0xd3, 0x0c, 0x5f, 0x80, 0x1d, 0x2d, 0x34, 0xc9, 0x22, 0x65, 0xe2, 0x98,
	0x2a, 0x75, 0x61, 0xb2, 0x48, 0xd2, 0x6f, 0x86, 0x2a, 0xad, 0xfa, 0xad, 0x61, 0x30, 0x5a, 0xc2,
	0x3d, 0xdb, 0x70, 0xd6, 0xe0, 0xd8, 0xc3, 0xf0, 0x25, 0x18, 0x38, 0x6e, 0x4d, 0x88, 0x18, 0x8f,
	0x0a, 0x29, 0x52, 0x49, 0x95, 0xea, 0xb7, 0x27, 0xc8, 0x35, 0xe5, 0x84, 0xbf, 0xf7, 0x30, 0x7c,
	0x0a, 0xb6, 0x1c, 0x99, 0x4a, 0x29, 0xe4, 0xf5, 0x99, 0x4b, 0x96, 0x06, 0x2d, 0x76, 0x5c, 0x41,
	0xcd, 0x71, 0x9f, 0xc1, 0x2d, 0x7b, 0x2d, 0x39, 0xd5, 0x92, 0xc5, 0x91, 0xaa, 0x06, 0xef, 0x2f,
	0x0f, 0xdb, 0xa3, 0xee, 0x18, 0xa1, 0xb9, 0xb7, 0x83, 0x8e, 0xfd, 0xe2, 0x54, 0x90, 0xe4, 0xad,
	0xa5, 0x59, 0xbb, 0xf0, 0x46, 0x36, 0x5d, 0x80, 0x03, 0xd0, 0x29, 0x24, 0x13, 0xb2, 0xf2, 0x6f,
	0x65, 0x18, 0x8c, 0xd6, 0x71, 0xf3, 0xbd, 0xfb, 0x33, 0x00, 0xbd, 0x05, 0x1b, 0xc1, 0x07, 0xa0,
	0xeb, 0xe5, 0x70, 0x92, 0x53, 0x6b, 0xfd, 0x2a, 0x06, 0xae, 0xf4, 0x8e, 0xe4, 0x14, 0xbe, 0x06,
	0x0f, 0xb9, 0xc9, 0xaf, 0x1d, 0xba, 0x60, 0x9c, 0xa9, 0x4b, 0x9a, 0x44, 0xdf, 0x99, 0xbe, 0xf4,
	0xa3, 0x78, 0x9f, 0xef, 0x71, 0x93, 0xd7, 0xc3, 0xbe, 0xf2, 0x6d, 0x9f, 0x98, 0xbe, 0x74, 0xe7,
	0xc1, 0x3d, 0xe0, 0x4c, 0xa9, 0xe7, 0x2f, 0x49, 0x66, 0xa8, 0x75, 0x39, 0xc0, 0x9b, 0x16, 0x71,
	0x8d, 0x1f, 0xab, 0xfa, 0xee, 0xdf, 0x36, 0x58, 0x3b, 0xcc, 0x8c, 0xd2, 0x54, 0x3a, 0xa5, 0x7b,
	0x60, 0x2d, 0x76, 0xdf, 0x13, 0x52, 0x0f, 0x56, 0x7f, 0xfd, 0xf9, 0xdd, 0x5e, 0x92, 0xad, 0x61,
	0x80, 0xbb, 0x1e, 0xb6, 0xb2, 0x0b, 0xd0, 0x33, 0x3e, 0x68, 0x51, 0x9d, 0x15, 0xef, 0x78, 0xcb,
	0x3a, 0xbe, 0xb7, 0xc0, 0xf1, 0xb9, 0xf1, 0x3c, 0x00, 0xd5, 0x31, 0xcb, 0x3f, 0x82, 0x56, 0x27,
	0xc0, 0xdb, 0x66, 0x6e, 0x82, 0x9f, 0x81, 0xdb, 0x6e, 0xbc, 0x44, 0x8a, 0xa2, 0xa0, 0xc9, 0x75,
	0x22, 0x5c, 0x90, 0x5c, 0x5a, 0x8e, 0x1c, 0xd8, 0x64, 0xe2, 0x2b, 0xd8, 0xbc, 0xd1, 0xef, 0x22,
	0xb1, 0xbf, 0x40, 0xe0, 0xa4, 0x29, 0x68, 0x66, 0x3b, 0xbc, 0x91, 0xcc, 0xec, 0xff, 0x06, 0x6c,
	0x4d, 0xfc, 0x0a, 0x22, 0xc6, 0x35, 0x95, 0x25, 0xc9, 0x6c, 0x4a, 0xbb, 0xe3, 0x1d, 0xe4, 0xde,
	0x37, 0xaa, 0xdf, 0x37, 0x3a, 0xf2, 0xef, 0x1b, 0xc3, 0x8a, 0x86, 0x2d, 0xeb, 0xc4, 0x93, 0x06,
	0x5f, 0xc0, 0xc6, 0xac, 0xfe, 0xc7, 0xa0, 0x13, 0x13, 0x4d, 0x53, 0x21, 0xaf, 0x6e, 0xde, 0x48,
	0x03, 0xc1, 0x47, 0x60, 0xbd, 0x1e, 0x33, 0x16, 0x86, 0x6b, 0x9f, 0x98, 0x35, 0x5f, 0x3c, 0xac,
	0x6a, 0xe7, 0x2b, 0x56, 0xc5, 0xfe, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xae, 0x69, 0xc1, 0x3b,
	0xd3, 0x04, 0x00, 0x00,
}