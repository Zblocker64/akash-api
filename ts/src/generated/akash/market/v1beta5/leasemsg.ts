// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.2.0
//   protoc               unknown
// source: akash/market/v1beta5/leasemsg.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import Long from "long";
import { messageTypeRegistry } from "../../../typeRegistry";
import { BidID } from "../v1/bid";
import { LeaseID } from "../v1/lease";

/** MsgCreateLease is sent to create a lease */
export interface MsgCreateLease {
  $type: "akash.market.v1beta5.MsgCreateLease";
  bidId: BidID | undefined;
}

/** MsgCreateLeaseResponse is the response from creating a lease */
export interface MsgCreateLeaseResponse {
  $type: "akash.market.v1beta5.MsgCreateLeaseResponse";
}

/** MsgWithdrawLease defines an SDK message for withdrawing lease funds */
export interface MsgWithdrawLease {
  $type: "akash.market.v1beta5.MsgWithdrawLease";
  bidId: LeaseID | undefined;
}

/** MsgWithdrawLeaseResponse defines the Msg/WithdrawLease response type. */
export interface MsgWithdrawLeaseResponse {
  $type: "akash.market.v1beta5.MsgWithdrawLeaseResponse";
}

/** MsgCloseLease defines an SDK message for closing order */
export interface MsgCloseLease {
  $type: "akash.market.v1beta5.MsgCloseLease";
  leaseId: LeaseID | undefined;
}

/** MsgCloseLeaseResponse defines the Msg/CloseLease response type. */
export interface MsgCloseLeaseResponse {
  $type: "akash.market.v1beta5.MsgCloseLeaseResponse";
}

function createBaseMsgCreateLease(): MsgCreateLease {
  return { $type: "akash.market.v1beta5.MsgCreateLease", bidId: undefined };
}

export const MsgCreateLease: MessageFns<
  MsgCreateLease,
  "akash.market.v1beta5.MsgCreateLease"
> = {
  $type: "akash.market.v1beta5.MsgCreateLease" as const,

  encode(
    message: MsgCreateLease,
    writer: BinaryWriter = new BinaryWriter(),
  ): BinaryWriter {
    if (message.bidId !== undefined) {
      BidID.encode(message.bidId, writer.uint32(10).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): MsgCreateLease {
    const reader =
      input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateLease();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.bidId = BidID.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MsgCreateLease {
    return {
      $type: MsgCreateLease.$type,
      bidId: isSet(object.bidId) ? BidID.fromJSON(object.bidId) : undefined,
    };
  },

  toJSON(message: MsgCreateLease): unknown {
    const obj: any = {};
    if (message.bidId !== undefined) {
      obj.bidId = BidID.toJSON(message.bidId);
    }
    return obj;
  },

  create(base?: DeepPartial<MsgCreateLease>): MsgCreateLease {
    return MsgCreateLease.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<MsgCreateLease>): MsgCreateLease {
    const message = createBaseMsgCreateLease();
    message.bidId =
      object.bidId !== undefined && object.bidId !== null
        ? BidID.fromPartial(object.bidId)
        : undefined;
    return message;
  },
};

messageTypeRegistry.set(MsgCreateLease.$type, MsgCreateLease);

function createBaseMsgCreateLeaseResponse(): MsgCreateLeaseResponse {
  return { $type: "akash.market.v1beta5.MsgCreateLeaseResponse" };
}

export const MsgCreateLeaseResponse: MessageFns<
  MsgCreateLeaseResponse,
  "akash.market.v1beta5.MsgCreateLeaseResponse"
> = {
  $type: "akash.market.v1beta5.MsgCreateLeaseResponse" as const,

  encode(
    _: MsgCreateLeaseResponse,
    writer: BinaryWriter = new BinaryWriter(),
  ): BinaryWriter {
    return writer;
  },

  decode(
    input: BinaryReader | Uint8Array,
    length?: number,
  ): MsgCreateLeaseResponse {
    const reader =
      input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateLeaseResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): MsgCreateLeaseResponse {
    return { $type: MsgCreateLeaseResponse.$type };
  },

  toJSON(_: MsgCreateLeaseResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create(base?: DeepPartial<MsgCreateLeaseResponse>): MsgCreateLeaseResponse {
    return MsgCreateLeaseResponse.fromPartial(base ?? {});
  },
  fromPartial(_: DeepPartial<MsgCreateLeaseResponse>): MsgCreateLeaseResponse {
    const message = createBaseMsgCreateLeaseResponse();
    return message;
  },
};

messageTypeRegistry.set(MsgCreateLeaseResponse.$type, MsgCreateLeaseResponse);

function createBaseMsgWithdrawLease(): MsgWithdrawLease {
  return { $type: "akash.market.v1beta5.MsgWithdrawLease", bidId: undefined };
}

export const MsgWithdrawLease: MessageFns<
  MsgWithdrawLease,
  "akash.market.v1beta5.MsgWithdrawLease"
> = {
  $type: "akash.market.v1beta5.MsgWithdrawLease" as const,

  encode(
    message: MsgWithdrawLease,
    writer: BinaryWriter = new BinaryWriter(),
  ): BinaryWriter {
    if (message.bidId !== undefined) {
      LeaseID.encode(message.bidId, writer.uint32(10).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): MsgWithdrawLease {
    const reader =
      input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgWithdrawLease();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.bidId = LeaseID.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MsgWithdrawLease {
    return {
      $type: MsgWithdrawLease.$type,
      bidId: isSet(object.bidId) ? LeaseID.fromJSON(object.bidId) : undefined,
    };
  },

  toJSON(message: MsgWithdrawLease): unknown {
    const obj: any = {};
    if (message.bidId !== undefined) {
      obj.bidId = LeaseID.toJSON(message.bidId);
    }
    return obj;
  },

  create(base?: DeepPartial<MsgWithdrawLease>): MsgWithdrawLease {
    return MsgWithdrawLease.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<MsgWithdrawLease>): MsgWithdrawLease {
    const message = createBaseMsgWithdrawLease();
    message.bidId =
      object.bidId !== undefined && object.bidId !== null
        ? LeaseID.fromPartial(object.bidId)
        : undefined;
    return message;
  },
};

messageTypeRegistry.set(MsgWithdrawLease.$type, MsgWithdrawLease);

function createBaseMsgWithdrawLeaseResponse(): MsgWithdrawLeaseResponse {
  return { $type: "akash.market.v1beta5.MsgWithdrawLeaseResponse" };
}

export const MsgWithdrawLeaseResponse: MessageFns<
  MsgWithdrawLeaseResponse,
  "akash.market.v1beta5.MsgWithdrawLeaseResponse"
> = {
  $type: "akash.market.v1beta5.MsgWithdrawLeaseResponse" as const,

  encode(
    _: MsgWithdrawLeaseResponse,
    writer: BinaryWriter = new BinaryWriter(),
  ): BinaryWriter {
    return writer;
  },

  decode(
    input: BinaryReader | Uint8Array,
    length?: number,
  ): MsgWithdrawLeaseResponse {
    const reader =
      input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgWithdrawLeaseResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): MsgWithdrawLeaseResponse {
    return { $type: MsgWithdrawLeaseResponse.$type };
  },

  toJSON(_: MsgWithdrawLeaseResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create(
    base?: DeepPartial<MsgWithdrawLeaseResponse>,
  ): MsgWithdrawLeaseResponse {
    return MsgWithdrawLeaseResponse.fromPartial(base ?? {});
  },
  fromPartial(
    _: DeepPartial<MsgWithdrawLeaseResponse>,
  ): MsgWithdrawLeaseResponse {
    const message = createBaseMsgWithdrawLeaseResponse();
    return message;
  },
};

messageTypeRegistry.set(
  MsgWithdrawLeaseResponse.$type,
  MsgWithdrawLeaseResponse,
);

function createBaseMsgCloseLease(): MsgCloseLease {
  return { $type: "akash.market.v1beta5.MsgCloseLease", leaseId: undefined };
}

export const MsgCloseLease: MessageFns<
  MsgCloseLease,
  "akash.market.v1beta5.MsgCloseLease"
> = {
  $type: "akash.market.v1beta5.MsgCloseLease" as const,

  encode(
    message: MsgCloseLease,
    writer: BinaryWriter = new BinaryWriter(),
  ): BinaryWriter {
    if (message.leaseId !== undefined) {
      LeaseID.encode(message.leaseId, writer.uint32(10).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): MsgCloseLease {
    const reader =
      input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCloseLease();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.leaseId = LeaseID.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MsgCloseLease {
    return {
      $type: MsgCloseLease.$type,
      leaseId: isSet(object.leaseId)
        ? LeaseID.fromJSON(object.leaseId)
        : undefined,
    };
  },

  toJSON(message: MsgCloseLease): unknown {
    const obj: any = {};
    if (message.leaseId !== undefined) {
      obj.leaseId = LeaseID.toJSON(message.leaseId);
    }
    return obj;
  },

  create(base?: DeepPartial<MsgCloseLease>): MsgCloseLease {
    return MsgCloseLease.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<MsgCloseLease>): MsgCloseLease {
    const message = createBaseMsgCloseLease();
    message.leaseId =
      object.leaseId !== undefined && object.leaseId !== null
        ? LeaseID.fromPartial(object.leaseId)
        : undefined;
    return message;
  },
};

messageTypeRegistry.set(MsgCloseLease.$type, MsgCloseLease);

function createBaseMsgCloseLeaseResponse(): MsgCloseLeaseResponse {
  return { $type: "akash.market.v1beta5.MsgCloseLeaseResponse" };
}

export const MsgCloseLeaseResponse: MessageFns<
  MsgCloseLeaseResponse,
  "akash.market.v1beta5.MsgCloseLeaseResponse"
> = {
  $type: "akash.market.v1beta5.MsgCloseLeaseResponse" as const,

  encode(
    _: MsgCloseLeaseResponse,
    writer: BinaryWriter = new BinaryWriter(),
  ): BinaryWriter {
    return writer;
  },

  decode(
    input: BinaryReader | Uint8Array,
    length?: number,
  ): MsgCloseLeaseResponse {
    const reader =
      input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCloseLeaseResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): MsgCloseLeaseResponse {
    return { $type: MsgCloseLeaseResponse.$type };
  },

  toJSON(_: MsgCloseLeaseResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create(base?: DeepPartial<MsgCloseLeaseResponse>): MsgCloseLeaseResponse {
    return MsgCloseLeaseResponse.fromPartial(base ?? {});
  },
  fromPartial(_: DeepPartial<MsgCloseLeaseResponse>): MsgCloseLeaseResponse {
    const message = createBaseMsgCloseLeaseResponse();
    return message;
  },
};

messageTypeRegistry.set(MsgCloseLeaseResponse.$type, MsgCloseLeaseResponse);

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
