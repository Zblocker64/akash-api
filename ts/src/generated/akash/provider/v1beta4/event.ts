// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.2.0
//   protoc               unknown
// source: akash/provider/v1beta4/event.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import Long from "long";
import { messageTypeRegistry } from "../../../typeRegistry";

/** EventProviderCreated defines an SDK message for provider created event */
export interface EventProviderCreated {
  $type: "akash.provider.v1beta4.EventProviderCreated";
  owner: string;
}

/** EventProviderUpdated defines an SDK message for provider updated event */
export interface EventProviderUpdated {
  $type: "akash.provider.v1beta4.EventProviderUpdated";
  owner: string;
}

/** EventProviderDeleted defines an SDK message for provider deleted event */
export interface EventProviderDeleted {
  $type: "akash.provider.v1beta4.EventProviderDeleted";
  owner: string;
}

function createBaseEventProviderCreated(): EventProviderCreated {
  return { $type: "akash.provider.v1beta4.EventProviderCreated", owner: "" };
}

export const EventProviderCreated: MessageFns<
  EventProviderCreated,
  "akash.provider.v1beta4.EventProviderCreated"
> = {
  $type: "akash.provider.v1beta4.EventProviderCreated" as const,

  encode(
    message: EventProviderCreated,
    writer: BinaryWriter = new BinaryWriter(),
  ): BinaryWriter {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    return writer;
  },

  decode(
    input: BinaryReader | Uint8Array,
    length?: number,
  ): EventProviderCreated {
    const reader =
      input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEventProviderCreated();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.owner = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): EventProviderCreated {
    return {
      $type: EventProviderCreated.$type,
      owner: isSet(object.owner) ? globalThis.String(object.owner) : "",
    };
  },

  toJSON(message: EventProviderCreated): unknown {
    const obj: any = {};
    if (message.owner !== "") {
      obj.owner = message.owner;
    }
    return obj;
  },

  create(base?: DeepPartial<EventProviderCreated>): EventProviderCreated {
    return EventProviderCreated.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<EventProviderCreated>): EventProviderCreated {
    const message = createBaseEventProviderCreated();
    message.owner = object.owner ?? "";
    return message;
  },
};

messageTypeRegistry.set(EventProviderCreated.$type, EventProviderCreated);

function createBaseEventProviderUpdated(): EventProviderUpdated {
  return { $type: "akash.provider.v1beta4.EventProviderUpdated", owner: "" };
}

export const EventProviderUpdated: MessageFns<
  EventProviderUpdated,
  "akash.provider.v1beta4.EventProviderUpdated"
> = {
  $type: "akash.provider.v1beta4.EventProviderUpdated" as const,

  encode(
    message: EventProviderUpdated,
    writer: BinaryWriter = new BinaryWriter(),
  ): BinaryWriter {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    return writer;
  },

  decode(
    input: BinaryReader | Uint8Array,
    length?: number,
  ): EventProviderUpdated {
    const reader =
      input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEventProviderUpdated();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.owner = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): EventProviderUpdated {
    return {
      $type: EventProviderUpdated.$type,
      owner: isSet(object.owner) ? globalThis.String(object.owner) : "",
    };
  },

  toJSON(message: EventProviderUpdated): unknown {
    const obj: any = {};
    if (message.owner !== "") {
      obj.owner = message.owner;
    }
    return obj;
  },

  create(base?: DeepPartial<EventProviderUpdated>): EventProviderUpdated {
    return EventProviderUpdated.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<EventProviderUpdated>): EventProviderUpdated {
    const message = createBaseEventProviderUpdated();
    message.owner = object.owner ?? "";
    return message;
  },
};

messageTypeRegistry.set(EventProviderUpdated.$type, EventProviderUpdated);

function createBaseEventProviderDeleted(): EventProviderDeleted {
  return { $type: "akash.provider.v1beta4.EventProviderDeleted", owner: "" };
}

export const EventProviderDeleted: MessageFns<
  EventProviderDeleted,
  "akash.provider.v1beta4.EventProviderDeleted"
> = {
  $type: "akash.provider.v1beta4.EventProviderDeleted" as const,

  encode(
    message: EventProviderDeleted,
    writer: BinaryWriter = new BinaryWriter(),
  ): BinaryWriter {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    return writer;
  },

  decode(
    input: BinaryReader | Uint8Array,
    length?: number,
  ): EventProviderDeleted {
    const reader =
      input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEventProviderDeleted();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.owner = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): EventProviderDeleted {
    return {
      $type: EventProviderDeleted.$type,
      owner: isSet(object.owner) ? globalThis.String(object.owner) : "",
    };
  },

  toJSON(message: EventProviderDeleted): unknown {
    const obj: any = {};
    if (message.owner !== "") {
      obj.owner = message.owner;
    }
    return obj;
  },

  create(base?: DeepPartial<EventProviderDeleted>): EventProviderDeleted {
    return EventProviderDeleted.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<EventProviderDeleted>): EventProviderDeleted {
    const message = createBaseEventProviderDeleted();
    message.owner = object.owner ?? "";
    return message;
  },
};

messageTypeRegistry.set(EventProviderDeleted.$type, EventProviderDeleted);

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
