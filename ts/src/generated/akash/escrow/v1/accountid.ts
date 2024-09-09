// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.2.0
//   protoc               unknown
// source: akash/escrow/v1/accountid.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import Long from "long";
import { messageTypeRegistry } from "../../../typeRegistry";

/** AccountID is the account identifier */
export interface AccountID {
  $type: "akash.escrow.v1.AccountID";
  scope: string;
  xid: string;
}

function createBaseAccountID(): AccountID {
  return { $type: "akash.escrow.v1.AccountID", scope: "", xid: "" };
}

export const AccountID: MessageFns<AccountID, "akash.escrow.v1.AccountID"> = {
  $type: "akash.escrow.v1.AccountID" as const,

  encode(
    message: AccountID,
    writer: BinaryWriter = new BinaryWriter(),
  ): BinaryWriter {
    if (message.scope !== "") {
      writer.uint32(10).string(message.scope);
    }
    if (message.xid !== "") {
      writer.uint32(18).string(message.xid);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): AccountID {
    const reader =
      input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAccountID();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.scope = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.xid = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): AccountID {
    return {
      $type: AccountID.$type,
      scope: isSet(object.scope) ? globalThis.String(object.scope) : "",
      xid: isSet(object.xid) ? globalThis.String(object.xid) : "",
    };
  },

  toJSON(message: AccountID): unknown {
    const obj: any = {};
    if (message.scope !== "") {
      obj.scope = message.scope;
    }
    if (message.xid !== "") {
      obj.xid = message.xid;
    }
    return obj;
  },

  create(base?: DeepPartial<AccountID>): AccountID {
    return AccountID.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<AccountID>): AccountID {
    const message = createBaseAccountID();
    message.scope = object.scope ?? "";
    message.xid = object.xid ?? "";
    return message;
  },
};

messageTypeRegistry.set(AccountID.$type, AccountID);

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
