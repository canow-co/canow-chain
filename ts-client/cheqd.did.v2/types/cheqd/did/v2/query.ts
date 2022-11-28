/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { DidDocWithMetadata } from "./diddoc";

export const protobufPackage = "cheqd.did.v2";

export interface QueryGetDidDocRequest {
  id: string;
}

export interface QueryGetDidDocResponse {
  value: DidDocWithMetadata | undefined;
}

function createBaseQueryGetDidDocRequest(): QueryGetDidDocRequest {
  return { id: "" };
}

export const QueryGetDidDocRequest = {
  encode(message: QueryGetDidDocRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetDidDocRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetDidDocRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetDidDocRequest {
    return { id: isSet(object.id) ? String(object.id) : "" };
  },

  toJSON(message: QueryGetDidDocRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetDidDocRequest>, I>>(object: I): QueryGetDidDocRequest {
    const message = createBaseQueryGetDidDocRequest();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseQueryGetDidDocResponse(): QueryGetDidDocResponse {
  return { value: undefined };
}

export const QueryGetDidDocResponse = {
  encode(message: QueryGetDidDocResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.value !== undefined) {
      DidDocWithMetadata.encode(message.value, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetDidDocResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetDidDocResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.value = DidDocWithMetadata.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetDidDocResponse {
    return { value: isSet(object.value) ? DidDocWithMetadata.fromJSON(object.value) : undefined };
  },

  toJSON(message: QueryGetDidDocResponse): unknown {
    const obj: any = {};
    message.value !== undefined && (obj.value = message.value ? DidDocWithMetadata.toJSON(message.value) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetDidDocResponse>, I>>(object: I): QueryGetDidDocResponse {
    const message = createBaseQueryGetDidDocResponse();
    message.value = (object.value !== undefined && object.value !== null)
      ? DidDocWithMetadata.fromPartial(object.value)
      : undefined;
    return message;
  },
};

export interface Query {
  DidDoc(request: QueryGetDidDocRequest): Promise<QueryGetDidDocResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.DidDoc = this.DidDoc.bind(this);
  }
  DidDoc(request: QueryGetDidDocRequest): Promise<QueryGetDidDocResponse> {
    const data = QueryGetDidDocRequest.encode(request).finish();
    const promise = this.rpc.request("cheqd.did.v2.Query", "DidDoc", data);
    return promise.then((data) => QueryGetDidDocResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
