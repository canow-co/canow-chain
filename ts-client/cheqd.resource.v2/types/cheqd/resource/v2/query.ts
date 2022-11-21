/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Metadata, ResourceWithMetadata } from "./resource";

export const protobufPackage = "cheqd.resource.v2";

export interface QueryGetResourceRequest {
  collectionId: string;
  id: string;
}

export interface QueryGetResourceResponse {
  resource: ResourceWithMetadata | undefined;
}

export interface QueryGetResourceMetadataRequest {
  collectionId: string;
  id: string;
}

export interface QueryGetResourceMetadataResponse {
  resource: Metadata | undefined;
}

export interface QueryGetCollectionResourcesRequest {
  collectionId: string;
}

export interface QueryGetCollectionResourcesResponse {
  resources: Metadata[];
}

function createBaseQueryGetResourceRequest(): QueryGetResourceRequest {
  return { collectionId: "", id: "" };
}

export const QueryGetResourceRequest = {
  encode(message: QueryGetResourceRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.collectionId !== "") {
      writer.uint32(10).string(message.collectionId);
    }
    if (message.id !== "") {
      writer.uint32(18).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetResourceRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetResourceRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.collectionId = reader.string();
          break;
        case 2:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetResourceRequest {
    return {
      collectionId: isSet(object.collectionId) ? String(object.collectionId) : "",
      id: isSet(object.id) ? String(object.id) : "",
    };
  },

  toJSON(message: QueryGetResourceRequest): unknown {
    const obj: any = {};
    message.collectionId !== undefined && (obj.collectionId = message.collectionId);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetResourceRequest>, I>>(object: I): QueryGetResourceRequest {
    const message = createBaseQueryGetResourceRequest();
    message.collectionId = object.collectionId ?? "";
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseQueryGetResourceResponse(): QueryGetResourceResponse {
  return { resource: undefined };
}

export const QueryGetResourceResponse = {
  encode(message: QueryGetResourceResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.resource !== undefined) {
      ResourceWithMetadata.encode(message.resource, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetResourceResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetResourceResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.resource = ResourceWithMetadata.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetResourceResponse {
    return { resource: isSet(object.resource) ? ResourceWithMetadata.fromJSON(object.resource) : undefined };
  },

  toJSON(message: QueryGetResourceResponse): unknown {
    const obj: any = {};
    message.resource !== undefined
      && (obj.resource = message.resource ? ResourceWithMetadata.toJSON(message.resource) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetResourceResponse>, I>>(object: I): QueryGetResourceResponse {
    const message = createBaseQueryGetResourceResponse();
    message.resource = (object.resource !== undefined && object.resource !== null)
      ? ResourceWithMetadata.fromPartial(object.resource)
      : undefined;
    return message;
  },
};

function createBaseQueryGetResourceMetadataRequest(): QueryGetResourceMetadataRequest {
  return { collectionId: "", id: "" };
}

export const QueryGetResourceMetadataRequest = {
  encode(message: QueryGetResourceMetadataRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.collectionId !== "") {
      writer.uint32(10).string(message.collectionId);
    }
    if (message.id !== "") {
      writer.uint32(18).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetResourceMetadataRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetResourceMetadataRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.collectionId = reader.string();
          break;
        case 2:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetResourceMetadataRequest {
    return {
      collectionId: isSet(object.collectionId) ? String(object.collectionId) : "",
      id: isSet(object.id) ? String(object.id) : "",
    };
  },

  toJSON(message: QueryGetResourceMetadataRequest): unknown {
    const obj: any = {};
    message.collectionId !== undefined && (obj.collectionId = message.collectionId);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetResourceMetadataRequest>, I>>(
    object: I,
  ): QueryGetResourceMetadataRequest {
    const message = createBaseQueryGetResourceMetadataRequest();
    message.collectionId = object.collectionId ?? "";
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseQueryGetResourceMetadataResponse(): QueryGetResourceMetadataResponse {
  return { resource: undefined };
}

export const QueryGetResourceMetadataResponse = {
  encode(message: QueryGetResourceMetadataResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.resource !== undefined) {
      Metadata.encode(message.resource, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetResourceMetadataResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetResourceMetadataResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.resource = Metadata.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetResourceMetadataResponse {
    return { resource: isSet(object.resource) ? Metadata.fromJSON(object.resource) : undefined };
  },

  toJSON(message: QueryGetResourceMetadataResponse): unknown {
    const obj: any = {};
    message.resource !== undefined && (obj.resource = message.resource ? Metadata.toJSON(message.resource) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetResourceMetadataResponse>, I>>(
    object: I,
  ): QueryGetResourceMetadataResponse {
    const message = createBaseQueryGetResourceMetadataResponse();
    message.resource = (object.resource !== undefined && object.resource !== null)
      ? Metadata.fromPartial(object.resource)
      : undefined;
    return message;
  },
};

function createBaseQueryGetCollectionResourcesRequest(): QueryGetCollectionResourcesRequest {
  return { collectionId: "" };
}

export const QueryGetCollectionResourcesRequest = {
  encode(message: QueryGetCollectionResourcesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.collectionId !== "") {
      writer.uint32(10).string(message.collectionId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetCollectionResourcesRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetCollectionResourcesRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.collectionId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetCollectionResourcesRequest {
    return { collectionId: isSet(object.collectionId) ? String(object.collectionId) : "" };
  },

  toJSON(message: QueryGetCollectionResourcesRequest): unknown {
    const obj: any = {};
    message.collectionId !== undefined && (obj.collectionId = message.collectionId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetCollectionResourcesRequest>, I>>(
    object: I,
  ): QueryGetCollectionResourcesRequest {
    const message = createBaseQueryGetCollectionResourcesRequest();
    message.collectionId = object.collectionId ?? "";
    return message;
  },
};

function createBaseQueryGetCollectionResourcesResponse(): QueryGetCollectionResourcesResponse {
  return { resources: [] };
}

export const QueryGetCollectionResourcesResponse = {
  encode(message: QueryGetCollectionResourcesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.resources) {
      Metadata.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetCollectionResourcesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetCollectionResourcesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.resources.push(Metadata.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetCollectionResourcesResponse {
    return {
      resources: Array.isArray(object?.resources) ? object.resources.map((e: any) => Metadata.fromJSON(e)) : [],
    };
  },

  toJSON(message: QueryGetCollectionResourcesResponse): unknown {
    const obj: any = {};
    if (message.resources) {
      obj.resources = message.resources.map((e) => e ? Metadata.toJSON(e) : undefined);
    } else {
      obj.resources = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetCollectionResourcesResponse>, I>>(
    object: I,
  ): QueryGetCollectionResourcesResponse {
    const message = createBaseQueryGetCollectionResourcesResponse();
    message.resources = object.resources?.map((e) => Metadata.fromPartial(e)) || [];
    return message;
  },
};

export interface Query {
  Resource(request: QueryGetResourceRequest): Promise<QueryGetResourceResponse>;
  ResourceMetadata(request: QueryGetResourceMetadataRequest): Promise<QueryGetResourceMetadataResponse>;
  CollectionResources(request: QueryGetCollectionResourcesRequest): Promise<QueryGetCollectionResourcesResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Resource = this.Resource.bind(this);
    this.ResourceMetadata = this.ResourceMetadata.bind(this);
    this.CollectionResources = this.CollectionResources.bind(this);
  }
  Resource(request: QueryGetResourceRequest): Promise<QueryGetResourceResponse> {
    const data = QueryGetResourceRequest.encode(request).finish();
    const promise = this.rpc.request("cheqd.resource.v2.Query", "Resource", data);
    return promise.then((data) => QueryGetResourceResponse.decode(new _m0.Reader(data)));
  }

  ResourceMetadata(request: QueryGetResourceMetadataRequest): Promise<QueryGetResourceMetadataResponse> {
    const data = QueryGetResourceMetadataRequest.encode(request).finish();
    const promise = this.rpc.request("cheqd.resource.v2.Query", "ResourceMetadata", data);
    return promise.then((data) => QueryGetResourceMetadataResponse.decode(new _m0.Reader(data)));
  }

  CollectionResources(request: QueryGetCollectionResourcesRequest): Promise<QueryGetCollectionResourcesResponse> {
    const data = QueryGetCollectionResourcesRequest.encode(request).finish();
    const promise = this.rpc.request("cheqd.resource.v2.Query", "CollectionResources", data);
    return promise.then((data) => QueryGetCollectionResourcesResponse.decode(new _m0.Reader(data)));
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
