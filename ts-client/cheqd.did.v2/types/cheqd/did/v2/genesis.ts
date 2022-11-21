/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { DidDocWithMetadata } from "./diddoc";
import { FeeParams } from "./fee";

export const protobufPackage = "cheqd.did.v2";

export interface GenesisState {
  didNamespace: string;
  didDocs: DidDocWithMetadata[];
  feeParams: FeeParams | undefined;
}

function createBaseGenesisState(): GenesisState {
  return { didNamespace: "", didDocs: [], feeParams: undefined };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.didNamespace !== "") {
      writer.uint32(10).string(message.didNamespace);
    }
    for (const v of message.didDocs) {
      DidDocWithMetadata.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.feeParams !== undefined) {
      FeeParams.encode(message.feeParams, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.didNamespace = reader.string();
          break;
        case 2:
          message.didDocs.push(DidDocWithMetadata.decode(reader, reader.uint32()));
          break;
        case 3:
          message.feeParams = FeeParams.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    return {
      didNamespace: isSet(object.didNamespace) ? String(object.didNamespace) : "",
      didDocs: Array.isArray(object?.didDocs) ? object.didDocs.map((e: any) => DidDocWithMetadata.fromJSON(e)) : [],
      feeParams: isSet(object.feeParams) ? FeeParams.fromJSON(object.feeParams) : undefined,
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.didNamespace !== undefined && (obj.didNamespace = message.didNamespace);
    if (message.didDocs) {
      obj.didDocs = message.didDocs.map((e) => e ? DidDocWithMetadata.toJSON(e) : undefined);
    } else {
      obj.didDocs = [];
    }
    message.feeParams !== undefined
      && (obj.feeParams = message.feeParams ? FeeParams.toJSON(message.feeParams) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.didNamespace = object.didNamespace ?? "";
    message.didDocs = object.didDocs?.map((e) => DidDocWithMetadata.fromPartial(e)) || [];
    message.feeParams = (object.feeParams !== undefined && object.feeParams !== null)
      ? FeeParams.fromPartial(object.feeParams)
      : undefined;
    return message;
  },
};

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
