// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.2.0
//   protoc               unknown
// source: akash/deployment/v1/event.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import Long from "long";
import { messageTypeRegistry } from "../../../typeRegistry";
import { DeploymentID } from "./deployment";
import { GroupID } from "./group";

/** EventDeploymentCreated event is triggered when deployment is created on chain */
export interface EventDeploymentCreated {
  $type: "akash.deployment.v1.EventDeploymentCreated";
  id: DeploymentID | undefined;
  hash: Uint8Array;
}

/** EventDeploymentUpdated is triggered when deployment is updated on chain */
export interface EventDeploymentUpdated {
  $type: "akash.deployment.v1.EventDeploymentUpdated";
  id: DeploymentID | undefined;
  hash: Uint8Array;
}

/** EventDeploymentClosed is triggered when deployment is closed on chain */
export interface EventDeploymentClosed {
  $type: "akash.deployment.v1.EventDeploymentClosed";
  id: DeploymentID | undefined;
}

/** EventGroupStarted is triggered when deployment group is started */
export interface EventGroupStarted {
  $type: "akash.deployment.v1.EventGroupStarted";
  id: GroupID | undefined;
}

/** EventGroupPaused is triggered when deployment group is paused */
export interface EventGroupPaused {
  $type: "akash.deployment.v1.EventGroupPaused";
  id: GroupID | undefined;
}

/** EventGroupClosed is triggered when deployment group is closed */
export interface EventGroupClosed {
  $type: "akash.deployment.v1.EventGroupClosed";
  id: GroupID | undefined;
}

function createBaseEventDeploymentCreated(): EventDeploymentCreated {
  return {
    $type: "akash.deployment.v1.EventDeploymentCreated",
    id: undefined,
    hash: new Uint8Array(0),
  };
}

export const EventDeploymentCreated: MessageFns<
  EventDeploymentCreated,
  "akash.deployment.v1.EventDeploymentCreated"
> = {
  $type: "akash.deployment.v1.EventDeploymentCreated" as const,

  encode(
    message: EventDeploymentCreated,
    writer: BinaryWriter = new BinaryWriter(),
  ): BinaryWriter {
    if (message.id !== undefined) {
      DeploymentID.encode(message.id, writer.uint32(10).fork()).join();
    }
    if (message.hash.length !== 0) {
      writer.uint32(18).bytes(message.hash);
    }
    return writer;
  },

  decode(
    input: BinaryReader | Uint8Array,
    length?: number,
  ): EventDeploymentCreated {
    const reader =
      input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEventDeploymentCreated();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.id = DeploymentID.decode(reader, reader.uint32());
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.hash = reader.bytes();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): EventDeploymentCreated {
    return {
      $type: EventDeploymentCreated.$type,
      id: isSet(object.id) ? DeploymentID.fromJSON(object.id) : undefined,
      hash: isSet(object.hash)
        ? bytesFromBase64(object.hash)
        : new Uint8Array(0),
    };
  },

  toJSON(message: EventDeploymentCreated): unknown {
    const obj: any = {};
    if (message.id !== undefined) {
      obj.id = DeploymentID.toJSON(message.id);
    }
    if (message.hash.length !== 0) {
      obj.hash = base64FromBytes(message.hash);
    }
    return obj;
  },

  create(base?: DeepPartial<EventDeploymentCreated>): EventDeploymentCreated {
    return EventDeploymentCreated.fromPartial(base ?? {});
  },
  fromPartial(
    object: DeepPartial<EventDeploymentCreated>,
  ): EventDeploymentCreated {
    const message = createBaseEventDeploymentCreated();
    message.id =
      object.id !== undefined && object.id !== null
        ? DeploymentID.fromPartial(object.id)
        : undefined;
    message.hash = object.hash ?? new Uint8Array(0);
    return message;
  },
};

messageTypeRegistry.set(EventDeploymentCreated.$type, EventDeploymentCreated);

function createBaseEventDeploymentUpdated(): EventDeploymentUpdated {
  return {
    $type: "akash.deployment.v1.EventDeploymentUpdated",
    id: undefined,
    hash: new Uint8Array(0),
  };
}

export const EventDeploymentUpdated: MessageFns<
  EventDeploymentUpdated,
  "akash.deployment.v1.EventDeploymentUpdated"
> = {
  $type: "akash.deployment.v1.EventDeploymentUpdated" as const,

  encode(
    message: EventDeploymentUpdated,
    writer: BinaryWriter = new BinaryWriter(),
  ): BinaryWriter {
    if (message.id !== undefined) {
      DeploymentID.encode(message.id, writer.uint32(10).fork()).join();
    }
    if (message.hash.length !== 0) {
      writer.uint32(18).bytes(message.hash);
    }
    return writer;
  },

  decode(
    input: BinaryReader | Uint8Array,
    length?: number,
  ): EventDeploymentUpdated {
    const reader =
      input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEventDeploymentUpdated();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.id = DeploymentID.decode(reader, reader.uint32());
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.hash = reader.bytes();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): EventDeploymentUpdated {
    return {
      $type: EventDeploymentUpdated.$type,
      id: isSet(object.id) ? DeploymentID.fromJSON(object.id) : undefined,
      hash: isSet(object.hash)
        ? bytesFromBase64(object.hash)
        : new Uint8Array(0),
    };
  },

  toJSON(message: EventDeploymentUpdated): unknown {
    const obj: any = {};
    if (message.id !== undefined) {
      obj.id = DeploymentID.toJSON(message.id);
    }
    if (message.hash.length !== 0) {
      obj.hash = base64FromBytes(message.hash);
    }
    return obj;
  },

  create(base?: DeepPartial<EventDeploymentUpdated>): EventDeploymentUpdated {
    return EventDeploymentUpdated.fromPartial(base ?? {});
  },
  fromPartial(
    object: DeepPartial<EventDeploymentUpdated>,
  ): EventDeploymentUpdated {
    const message = createBaseEventDeploymentUpdated();
    message.id =
      object.id !== undefined && object.id !== null
        ? DeploymentID.fromPartial(object.id)
        : undefined;
    message.hash = object.hash ?? new Uint8Array(0);
    return message;
  },
};

messageTypeRegistry.set(EventDeploymentUpdated.$type, EventDeploymentUpdated);

function createBaseEventDeploymentClosed(): EventDeploymentClosed {
  return { $type: "akash.deployment.v1.EventDeploymentClosed", id: undefined };
}

export const EventDeploymentClosed: MessageFns<
  EventDeploymentClosed,
  "akash.deployment.v1.EventDeploymentClosed"
> = {
  $type: "akash.deployment.v1.EventDeploymentClosed" as const,

  encode(
    message: EventDeploymentClosed,
    writer: BinaryWriter = new BinaryWriter(),
  ): BinaryWriter {
    if (message.id !== undefined) {
      DeploymentID.encode(message.id, writer.uint32(10).fork()).join();
    }
    return writer;
  },

  decode(
    input: BinaryReader | Uint8Array,
    length?: number,
  ): EventDeploymentClosed {
    const reader =
      input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEventDeploymentClosed();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.id = DeploymentID.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): EventDeploymentClosed {
    return {
      $type: EventDeploymentClosed.$type,
      id: isSet(object.id) ? DeploymentID.fromJSON(object.id) : undefined,
    };
  },

  toJSON(message: EventDeploymentClosed): unknown {
    const obj: any = {};
    if (message.id !== undefined) {
      obj.id = DeploymentID.toJSON(message.id);
    }
    return obj;
  },

  create(base?: DeepPartial<EventDeploymentClosed>): EventDeploymentClosed {
    return EventDeploymentClosed.fromPartial(base ?? {});
  },
  fromPartial(
    object: DeepPartial<EventDeploymentClosed>,
  ): EventDeploymentClosed {
    const message = createBaseEventDeploymentClosed();
    message.id =
      object.id !== undefined && object.id !== null
        ? DeploymentID.fromPartial(object.id)
        : undefined;
    return message;
  },
};

messageTypeRegistry.set(EventDeploymentClosed.$type, EventDeploymentClosed);

function createBaseEventGroupStarted(): EventGroupStarted {
  return { $type: "akash.deployment.v1.EventGroupStarted", id: undefined };
}

export const EventGroupStarted: MessageFns<
  EventGroupStarted,
  "akash.deployment.v1.EventGroupStarted"
> = {
  $type: "akash.deployment.v1.EventGroupStarted" as const,

  encode(
    message: EventGroupStarted,
    writer: BinaryWriter = new BinaryWriter(),
  ): BinaryWriter {
    if (message.id !== undefined) {
      GroupID.encode(message.id, writer.uint32(10).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): EventGroupStarted {
    const reader =
      input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEventGroupStarted();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.id = GroupID.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): EventGroupStarted {
    return {
      $type: EventGroupStarted.$type,
      id: isSet(object.id) ? GroupID.fromJSON(object.id) : undefined,
    };
  },

  toJSON(message: EventGroupStarted): unknown {
    const obj: any = {};
    if (message.id !== undefined) {
      obj.id = GroupID.toJSON(message.id);
    }
    return obj;
  },

  create(base?: DeepPartial<EventGroupStarted>): EventGroupStarted {
    return EventGroupStarted.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<EventGroupStarted>): EventGroupStarted {
    const message = createBaseEventGroupStarted();
    message.id =
      object.id !== undefined && object.id !== null
        ? GroupID.fromPartial(object.id)
        : undefined;
    return message;
  },
};

messageTypeRegistry.set(EventGroupStarted.$type, EventGroupStarted);

function createBaseEventGroupPaused(): EventGroupPaused {
  return { $type: "akash.deployment.v1.EventGroupPaused", id: undefined };
}

export const EventGroupPaused: MessageFns<
  EventGroupPaused,
  "akash.deployment.v1.EventGroupPaused"
> = {
  $type: "akash.deployment.v1.EventGroupPaused" as const,

  encode(
    message: EventGroupPaused,
    writer: BinaryWriter = new BinaryWriter(),
  ): BinaryWriter {
    if (message.id !== undefined) {
      GroupID.encode(message.id, writer.uint32(10).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): EventGroupPaused {
    const reader =
      input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEventGroupPaused();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.id = GroupID.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): EventGroupPaused {
    return {
      $type: EventGroupPaused.$type,
      id: isSet(object.id) ? GroupID.fromJSON(object.id) : undefined,
    };
  },

  toJSON(message: EventGroupPaused): unknown {
    const obj: any = {};
    if (message.id !== undefined) {
      obj.id = GroupID.toJSON(message.id);
    }
    return obj;
  },

  create(base?: DeepPartial<EventGroupPaused>): EventGroupPaused {
    return EventGroupPaused.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<EventGroupPaused>): EventGroupPaused {
    const message = createBaseEventGroupPaused();
    message.id =
      object.id !== undefined && object.id !== null
        ? GroupID.fromPartial(object.id)
        : undefined;
    return message;
  },
};

messageTypeRegistry.set(EventGroupPaused.$type, EventGroupPaused);

function createBaseEventGroupClosed(): EventGroupClosed {
  return { $type: "akash.deployment.v1.EventGroupClosed", id: undefined };
}

export const EventGroupClosed: MessageFns<
  EventGroupClosed,
  "akash.deployment.v1.EventGroupClosed"
> = {
  $type: "akash.deployment.v1.EventGroupClosed" as const,

  encode(
    message: EventGroupClosed,
    writer: BinaryWriter = new BinaryWriter(),
  ): BinaryWriter {
    if (message.id !== undefined) {
      GroupID.encode(message.id, writer.uint32(10).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): EventGroupClosed {
    const reader =
      input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEventGroupClosed();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.id = GroupID.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): EventGroupClosed {
    return {
      $type: EventGroupClosed.$type,
      id: isSet(object.id) ? GroupID.fromJSON(object.id) : undefined,
    };
  },

  toJSON(message: EventGroupClosed): unknown {
    const obj: any = {};
    if (message.id !== undefined) {
      obj.id = GroupID.toJSON(message.id);
    }
    return obj;
  },

  create(base?: DeepPartial<EventGroupClosed>): EventGroupClosed {
    return EventGroupClosed.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<EventGroupClosed>): EventGroupClosed {
    const message = createBaseEventGroupClosed();
    message.id =
      object.id !== undefined && object.id !== null
        ? GroupID.fromPartial(object.id)
        : undefined;
    return message;
  },
};

messageTypeRegistry.set(EventGroupClosed.$type, EventGroupClosed);

function bytesFromBase64(b64: string): Uint8Array {
  if ((globalThis as any).Buffer) {
    return Uint8Array.from(globalThis.Buffer.from(b64, "base64"));
  } else {
    const bin = globalThis.atob(b64);
    const arr = new Uint8Array(bin.length);
    for (let i = 0; i < bin.length; ++i) {
      arr[i] = bin.charCodeAt(i);
    }
    return arr;
  }
}

function base64FromBytes(arr: Uint8Array): string {
  if ((globalThis as any).Buffer) {
    return globalThis.Buffer.from(arr).toString("base64");
  } else {
    const bin: string[] = [];
    arr.forEach((byte) => {
      bin.push(globalThis.String.fromCharCode(byte));
    });
    return globalThis.btoa(bin.join(""));
  }
}

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
