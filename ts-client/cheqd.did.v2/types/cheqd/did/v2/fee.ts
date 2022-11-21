/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Coin } from "../../../cosmos/base/v1beta1/coin";

export const protobufPackage = "cheqd.did.v2";

/** FeeParams defines the parameters for the `did` module fixed fee. */
export interface FeeParams {
  /** Tx types define the fixed fee each for the `did` module. */
  txTypes: { [key: string]: Coin };
  burnFactor: string;
}

export interface FeeParams_TxTypesEntry {
  key: string;
  value: Coin | undefined;
}

function createBaseFeeParams(): FeeParams {
  return { txTypes: {}, burnFactor: "" };
}

export const FeeParams = {
  encode(message: FeeParams, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    Object.entries(message.txTypes).forEach(([key, value]) => {
      FeeParams_TxTypesEntry.encode({ key: key as any, value }, writer.uint32(10).fork()).ldelim();
    });
    if (message.burnFactor !== "") {
      writer.uint32(18).string(message.burnFactor);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FeeParams {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFeeParams();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          const entry1 = FeeParams_TxTypesEntry.decode(reader, reader.uint32());
          if (entry1.value !== undefined) {
            message.txTypes[entry1.key] = entry1.value;
          }
          break;
        case 2:
          message.burnFactor = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FeeParams {
    return {
      txTypes: isObject(object.txTypes)
        ? Object.entries(object.txTypes).reduce<{ [key: string]: Coin }>((acc, [key, value]) => {
          acc[key] = Coin.fromJSON(value);
          return acc;
        }, {})
        : {},
      burnFactor: isSet(object.burnFactor) ? String(object.burnFactor) : "",
    };
  },

  toJSON(message: FeeParams): unknown {
    const obj: any = {};
    obj.txTypes = {};
    if (message.txTypes) {
      Object.entries(message.txTypes).forEach(([k, v]) => {
        obj.txTypes[k] = Coin.toJSON(v);
      });
    }
    message.burnFactor !== undefined && (obj.burnFactor = message.burnFactor);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<FeeParams>, I>>(object: I): FeeParams {
    const message = createBaseFeeParams();
    message.txTypes = Object.entries(object.txTypes ?? {}).reduce<{ [key: string]: Coin }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = Coin.fromPartial(value);
      }
      return acc;
    }, {});
    message.burnFactor = object.burnFactor ?? "";
    return message;
  },
};

function createBaseFeeParams_TxTypesEntry(): FeeParams_TxTypesEntry {
  return { key: "", value: undefined };
}

export const FeeParams_TxTypesEntry = {
  encode(message: FeeParams_TxTypesEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== undefined) {
      Coin.encode(message.value, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FeeParams_TxTypesEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFeeParams_TxTypesEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = Coin.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): FeeParams_TxTypesEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? Coin.fromJSON(object.value) : undefined,
    };
  },

  toJSON(message: FeeParams_TxTypesEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value ? Coin.toJSON(message.value) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<FeeParams_TxTypesEntry>, I>>(object: I): FeeParams_TxTypesEntry {
    const message = createBaseFeeParams_TxTypesEntry();
    message.key = object.key ?? "";
    message.value = (object.value !== undefined && object.value !== null) ? Coin.fromPartial(object.value) : undefined;
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

function isObject(value: any): boolean {
  return typeof value === "object" && value !== null;
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
