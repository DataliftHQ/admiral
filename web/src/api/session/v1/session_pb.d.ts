// @generated by protoc-gen-es v1.8.0
// @generated from file session/v1/session.proto (package admiral.session.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message admiral.session.v1.GetSessionRequest
 */
export declare class GetSessionRequest extends Message<GetSessionRequest> {
  constructor(data?: PartialMessage<GetSessionRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "admiral.session.v1.GetSessionRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetSessionRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetSessionRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetSessionRequest;

  static equals(a: GetSessionRequest | PlainMessage<GetSessionRequest> | undefined, b: GetSessionRequest | PlainMessage<GetSessionRequest> | undefined): boolean;
}

/**
 * @generated from message admiral.session.v1.GetSessionResponse
 */
export declare class GetSessionResponse extends Message<GetSessionResponse> {
  /**
   * @generated from field: bool loggedIn = 1;
   */
  loggedIn: boolean;

  /**
   * @generated from field: string username = 2;
   */
  username: string;

  /**
   * @generated from field: string iss = 3;
   */
  iss: string;

  /**
   * @generated from field: repeated string groups = 4;
   */
  groups: string[];

  constructor(data?: PartialMessage<GetSessionResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "admiral.session.v1.GetSessionResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetSessionResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetSessionResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetSessionResponse;

  static equals(a: GetSessionResponse | PlainMessage<GetSessionResponse> | undefined, b: GetSessionResponse | PlainMessage<GetSessionResponse> | undefined): boolean;
}

