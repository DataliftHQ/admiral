// @generated by protoc-gen-es v1.8.0
// @generated from file foo/v1/foo.proto (package admiral.foo.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message admiral.foo.v1.GetFooRequest
 */
export const GetFooRequest = /*@__PURE__*/ proto3.makeMessageType(
  "admiral.foo.v1.GetFooRequest",
  [],
);

/**
 * @generated from message admiral.foo.v1.GetFooResponse
 */
export const GetFooResponse = /*@__PURE__*/ proto3.makeMessageType(
  "admiral.foo.v1.GetFooResponse",
  () => [
    { no: 1, name: "foo", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

