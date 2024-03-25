// @generated by protoc-gen-es v1.8.0
// @generated from file audit/v1/audit.proto (package admiral.audit.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { Any, Duration, proto3, Timestamp } from "@bufbuild/protobuf";
import { ActionType } from "../../common/v1/schema_pb.js";
import { Status } from "../../google/rpc/status_pb.js";

/**
 * @generated from message admiral.audit.v1.TimeRange
 */
export const TimeRange = /*@__PURE__*/ proto3.makeMessageType(
  "admiral.audit.v1.TimeRange",
  () => [
    { no: 1, name: "start_time", kind: "message", T: Timestamp },
    { no: 2, name: "end_time", kind: "message", T: Timestamp },
  ],
);

/**
 * @generated from message admiral.audit.v1.GetEventsRequest
 */
export const GetEventsRequest = /*@__PURE__*/ proto3.makeMessageType(
  "admiral.audit.v1.GetEventsRequest",
  () => [
    { no: 1, name: "range", kind: "message", T: TimeRange, oneof: "window" },
    { no: 2, name: "since", kind: "message", T: Duration, oneof: "window" },
    { no: 3, name: "page_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "limit", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
  ],
);

/**
 * @generated from message admiral.audit.v1.Resource
 */
export const Resource = /*@__PURE__*/ proto3.makeMessageType(
  "admiral.audit.v1.Resource",
  () => [
    { no: 1, name: "type_url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message admiral.audit.v1.RequestMetadata
 */
export const RequestMetadata = /*@__PURE__*/ proto3.makeMessageType(
  "admiral.audit.v1.RequestMetadata",
  () => [
    { no: 1, name: "body", kind: "message", T: Any },
  ],
);

/**
 * @generated from message admiral.audit.v1.ResponseMetadata
 */
export const ResponseMetadata = /*@__PURE__*/ proto3.makeMessageType(
  "admiral.audit.v1.ResponseMetadata",
  () => [
    { no: 1, name: "body", kind: "message", T: Any },
  ],
);

/**
 * @generated from message admiral.audit.v1.RequestEvent
 */
export const RequestEvent = /*@__PURE__*/ proto3.makeMessageType(
  "admiral.audit.v1.RequestEvent",
  () => [
    { no: 1, name: "username", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "service_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "method_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "type", kind: "enum", T: proto3.getEnumType(ActionType) },
    { no: 5, name: "status", kind: "message", T: Status },
    { no: 6, name: "resources", kind: "message", T: Resource, repeated: true },
    { no: 7, name: "request_metadata", kind: "message", T: RequestMetadata },
    { no: 8, name: "response_metadata", kind: "message", T: ResponseMetadata },
  ],
);

/**
 * @generated from message admiral.audit.v1.Event
 */
export const Event = /*@__PURE__*/ proto3.makeMessageType(
  "admiral.audit.v1.Event",
  () => [
    { no: 1, name: "occurred_at", kind: "message", T: Timestamp },
    { no: 2, name: "event", kind: "message", T: RequestEvent, oneof: "event_type" },
    { no: 3, name: "id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ],
);

/**
 * @generated from message admiral.audit.v1.GetEventsResponse
 */
export const GetEventsResponse = /*@__PURE__*/ proto3.makeMessageType(
  "admiral.audit.v1.GetEventsResponse",
  () => [
    { no: 1, name: "events", kind: "message", T: Event, repeated: true },
    { no: 2, name: "next_page_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message admiral.audit.v1.GetEventRequest
 */
export const GetEventRequest = /*@__PURE__*/ proto3.makeMessageType(
  "admiral.audit.v1.GetEventRequest",
  () => [
    { no: 1, name: "event_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ],
);

/**
 * @generated from message admiral.audit.v1.GetEventResponse
 */
export const GetEventResponse = /*@__PURE__*/ proto3.makeMessageType(
  "admiral.audit.v1.GetEventResponse",
  () => [
    { no: 1, name: "event", kind: "message", T: Event },
  ],
);

