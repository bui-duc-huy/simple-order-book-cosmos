/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Coin } from "../cosmos/base/v1beta1/coin";

export const protobufPackage = "orderbook.dex";

export interface Order {
  id: number;
  sourceToken: Coin | undefined;
  destToken: Coin | undefined;
  creator: string;
  state: string;
  buyer: string;
}

const baseOrder: object = { id: 0, creator: "", state: "", buyer: "" };

export const Order = {
  encode(message: Order, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.sourceToken !== undefined) {
      Coin.encode(message.sourceToken, writer.uint32(18).fork()).ldelim();
    }
    if (message.destToken !== undefined) {
      Coin.encode(message.destToken, writer.uint32(26).fork()).ldelim();
    }
    if (message.creator !== "") {
      writer.uint32(34).string(message.creator);
    }
    if (message.state !== "") {
      writer.uint32(42).string(message.state);
    }
    if (message.buyer !== "") {
      writer.uint32(50).string(message.buyer);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Order {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseOrder } as Order;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.sourceToken = Coin.decode(reader, reader.uint32());
          break;
        case 3:
          message.destToken = Coin.decode(reader, reader.uint32());
          break;
        case 4:
          message.creator = reader.string();
          break;
        case 5:
          message.state = reader.string();
          break;
        case 6:
          message.buyer = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Order {
    const message = { ...baseOrder } as Order;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.sourceToken !== undefined && object.sourceToken !== null) {
      message.sourceToken = Coin.fromJSON(object.sourceToken);
    } else {
      message.sourceToken = undefined;
    }
    if (object.destToken !== undefined && object.destToken !== null) {
      message.destToken = Coin.fromJSON(object.destToken);
    } else {
      message.destToken = undefined;
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.state !== undefined && object.state !== null) {
      message.state = String(object.state);
    } else {
      message.state = "";
    }
    if (object.buyer !== undefined && object.buyer !== null) {
      message.buyer = String(object.buyer);
    } else {
      message.buyer = "";
    }
    return message;
  },

  toJSON(message: Order): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.sourceToken !== undefined &&
      (obj.sourceToken = message.sourceToken
        ? Coin.toJSON(message.sourceToken)
        : undefined);
    message.destToken !== undefined &&
      (obj.destToken = message.destToken
        ? Coin.toJSON(message.destToken)
        : undefined);
    message.creator !== undefined && (obj.creator = message.creator);
    message.state !== undefined && (obj.state = message.state);
    message.buyer !== undefined && (obj.buyer = message.buyer);
    return obj;
  },

  fromPartial(object: DeepPartial<Order>): Order {
    const message = { ...baseOrder } as Order;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.sourceToken !== undefined && object.sourceToken !== null) {
      message.sourceToken = Coin.fromPartial(object.sourceToken);
    } else {
      message.sourceToken = undefined;
    }
    if (object.destToken !== undefined && object.destToken !== null) {
      message.destToken = Coin.fromPartial(object.destToken);
    } else {
      message.destToken = undefined;
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.state !== undefined && object.state !== null) {
      message.state = object.state;
    } else {
      message.state = "";
    }
    if (object.buyer !== undefined && object.buyer !== null) {
      message.buyer = object.buyer;
    } else {
      message.buyer = "";
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
