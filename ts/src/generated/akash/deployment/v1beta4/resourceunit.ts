// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.2.0
//   protoc               unknown
// source: akash/deployment/v1beta4/resourceunit.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import Long from "long";
import { DecCoin } from "../../../cosmos/base/v1beta1/coin";
import { messageTypeRegistry } from "../../../typeRegistry";
import { Resources } from "../../base/resources/v1beta4/resources";

/** ResourceUnit extends Resources and adds Count along with the Price */
export interface ResourceUnit {
  $type: "akash.deployment.v1beta4.ResourceUnit";
  resource: Resources | undefined;
  count: number;
  price: DecCoin | undefined;
}

function createBaseResourceUnit(): ResourceUnit {
  return {
    $type: "akash.deployment.v1beta4.ResourceUnit",
    resource: undefined,
    count: 0,
    price: undefined,
  };
}

export const ResourceUnit: MessageFns<
  ResourceUnit,
  "akash.deployment.v1beta4.ResourceUnit"
> = {
  $type: "akash.deployment.v1beta4.ResourceUnit" as const,

  encode(
    message: ResourceUnit,
    writer: BinaryWriter = new BinaryWriter(),
  ): BinaryWriter {
    if (message.resource !== undefined) {
      Resources.encode(message.resource, writer.uint32(10).fork()).join();
    }
    if (message.count !== 0) {
      writer.uint32(16).uint32(message.count);
    }
    if (message.price !== undefined) {
      DecCoin.encode(message.price, writer.uint32(26).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): ResourceUnit {
    const reader =
      input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseResourceUnit();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.resource = Resources.decode(reader, reader.uint32());
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.count = reader.uint32();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.price = DecCoin.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ResourceUnit {
    return {
      $type: ResourceUnit.$type,
      resource: isSet(object.resource)
        ? Resources.fromJSON(object.resource)
        : undefined,
      count: isSet(object.count) ? globalThis.Number(object.count) : 0,
      price: isSet(object.price) ? DecCoin.fromJSON(object.price) : undefined,
    };
  },

  toJSON(message: ResourceUnit): unknown {
    const obj: any = {};
    if (message.resource !== undefined) {
      obj.resource = Resources.toJSON(message.resource);
    }
    if (message.count !== 0) {
      obj.count = Math.round(message.count);
    }
    if (message.price !== undefined) {
      obj.price = DecCoin.toJSON(message.price);
    }
    return obj;
  },

  create(base?: DeepPartial<ResourceUnit>): ResourceUnit {
    return ResourceUnit.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<ResourceUnit>): ResourceUnit {
    const message = createBaseResourceUnit();
    message.resource =
      object.resource !== undefined && object.resource !== null
        ? Resources.fromPartial(object.resource)
        : undefined;
    message.count = object.count ?? 0;
    message.price =
      object.price !== undefined && object.price !== null
        ? DecCoin.fromPartial(object.price)
        : undefined;
    return message;
  },
};

messageTypeRegistry.set(ResourceUnit.$type, ResourceUnit);

type Builtin =
  | Date
  | Function
  | Uint8Array
  | string
  | number
  | boolean
  | undefined;

type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Long
    ? string | number | Long
    : T extends globalThis.Array<infer U>
      ? globalThis.Array<DeepPartial<U>>
      : T extends ReadonlyArray<infer U>
        ? ReadonlyArray<DeepPartial<U>>
        : T extends {}
          ? { [K in Exclude<keyof T, "$type">]?: DeepPartial<T[K]> }
          : Partial<T>;

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

interface MessageFns<T, V extends string> {
  readonly $type: V;
  encode(message: T, writer?: BinaryWriter): BinaryWriter;
  decode(input: BinaryReader | Uint8Array, length?: number): T;
  fromJSON(object: any): T;
  toJSON(message: T): unknown;
  create(base?: DeepPartial<T>): T;
  fromPartial(object: DeepPartial<T>): T;
}
