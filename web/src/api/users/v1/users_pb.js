// @generated by protoc-gen-es v1.8.0
// @generated from file users/v1/users.proto (package admiral.users.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message admiral.users.v1.GetMeRequest
 */
export const GetMeRequest = /*@__PURE__*/ proto3.makeMessageType(
  "admiral.users.v1.GetMeRequest",
  [],
);

/**
 * @generated from message admiral.users.v1.GetMeResponse
 */
export const GetMeResponse = /*@__PURE__*/ proto3.makeMessageType(
  "admiral.users.v1.GetMeResponse",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "email", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "given_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "family_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message admiral.users.v1.GetUserRequest
 */
export const GetUserRequest = /*@__PURE__*/ proto3.makeMessageType(
  "admiral.users.v1.GetUserRequest",
  [],
);

/**
 * @generated from message admiral.users.v1.GetUserResponse
 */
export const GetUserResponse = /*@__PURE__*/ proto3.makeMessageType(
  "admiral.users.v1.GetUserResponse",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "email", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "given_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "family_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

