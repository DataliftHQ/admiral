// @generated by protoc-gen-es v1.8.0
// @generated from file users/v1/users.proto (package admiral.users.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message admiral.users.v1.GetMeRequest
 */
export declare class GetMeRequest extends Message<GetMeRequest> {
  constructor(data?: PartialMessage<GetMeRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "admiral.users.v1.GetMeRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetMeRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetMeRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetMeRequest;

  static equals(a: GetMeRequest | PlainMessage<GetMeRequest> | undefined, b: GetMeRequest | PlainMessage<GetMeRequest> | undefined): boolean;
}

/**
 * @generated from message admiral.users.v1.GetMeResponse
 */
export declare class GetMeResponse extends Message<GetMeResponse> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: string email = 2;
   */
  email: string;

  /**
   * @generated from field: string given_name = 3;
   */
  givenName: string;

  /**
   * @generated from field: string family_name = 4;
   */
  familyName: string;

  constructor(data?: PartialMessage<GetMeResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "admiral.users.v1.GetMeResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetMeResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetMeResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetMeResponse;

  static equals(a: GetMeResponse | PlainMessage<GetMeResponse> | undefined, b: GetMeResponse | PlainMessage<GetMeResponse> | undefined): boolean;
}

/**
 * @generated from message admiral.users.v1.GetUserRequest
 */
export declare class GetUserRequest extends Message<GetUserRequest> {
  constructor(data?: PartialMessage<GetUserRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "admiral.users.v1.GetUserRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetUserRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetUserRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetUserRequest;

  static equals(a: GetUserRequest | PlainMessage<GetUserRequest> | undefined, b: GetUserRequest | PlainMessage<GetUserRequest> | undefined): boolean;
}

/**
 * @generated from message admiral.users.v1.GetUserResponse
 */
export declare class GetUserResponse extends Message<GetUserResponse> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: string email = 2;
   */
  email: string;

  /**
   * @generated from field: string given_name = 3;
   */
  givenName: string;

  /**
   * @generated from field: string family_name = 4;
   */
  familyName: string;

  constructor(data?: PartialMessage<GetUserResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "admiral.users.v1.GetUserResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetUserResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetUserResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetUserResponse;

  static equals(a: GetUserResponse | PlainMessage<GetUserResponse> | undefined, b: GetUserResponse | PlainMessage<GetUserResponse> | undefined): boolean;
}

