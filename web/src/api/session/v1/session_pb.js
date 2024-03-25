// @generated by protoc-gen-es v1.8.0
// @generated from file session/v1/session.proto (package admiral.session.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message admiral.session.v1.GetSessionRequest
 */
export const GetSessionRequest = /*@__PURE__*/ proto3.makeMessageType(
  "admiral.session.v1.GetSessionRequest",
  [],
);

/**
 * @generated from message admiral.session.v1.GetSessionResponse
 */
export const GetSessionResponse = /*@__PURE__*/ proto3.makeMessageType(
  "admiral.session.v1.GetSessionResponse",
  () => [
    { no: 1, name: "loggedIn", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 2, name: "username", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "iss", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "groups", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);
