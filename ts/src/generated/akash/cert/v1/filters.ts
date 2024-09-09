// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.2.0
//   protoc               unknown
// source: akash/cert/v1/filters.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import Long from "long";
import { messageTypeRegistry } from "../../../typeRegistry";

/** CertificateFilter defines filters used to filter certificates */
export interface CertificateFilter {
  $type: "akash.cert.v1.CertificateFilter";
  owner: string;
  serial: string;
  state: string;
}

function createBaseCertificateFilter(): CertificateFilter {
  return {
    $type: "akash.cert.v1.CertificateFilter",
    owner: "",
    serial: "",
    state: "",
  };
}

export const CertificateFilter: MessageFns<
  CertificateFilter,
  "akash.cert.v1.CertificateFilter"
> = {
  $type: "akash.cert.v1.CertificateFilter" as const,

  encode(
    message: CertificateFilter,
    writer: BinaryWriter = new BinaryWriter(),
  ): BinaryWriter {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.serial !== "") {
      writer.uint32(18).string(message.serial);
    }
    if (message.state !== "") {
      writer.uint32(26).string(message.state);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): CertificateFilter {
    const reader =
      input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCertificateFilter();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.owner = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.serial = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.state = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CertificateFilter {
    return {
      $type: CertificateFilter.$type,
      owner: isSet(object.owner) ? globalThis.String(object.owner) : "",
      serial: isSet(object.serial) ? globalThis.String(object.serial) : "",
      state: isSet(object.state) ? globalThis.String(object.state) : "",
    };
  },

  toJSON(message: CertificateFilter): unknown {
    const obj: any = {};
    if (message.owner !== "") {
      obj.owner = message.owner;
    }
    if (message.serial !== "") {
      obj.serial = message.serial;
    }
    if (message.state !== "") {
      obj.state = message.state;
    }
    return obj;
  },

  create(base?: DeepPartial<CertificateFilter>): CertificateFilter {
    return CertificateFilter.fromPartial(base ?? {});
  },
  fromPartial(object: DeepPartial<CertificateFilter>): CertificateFilter {
    const message = createBaseCertificateFilter();
    message.owner = object.owner ?? "";
    message.serial = object.serial ?? "";
    message.state = object.state ?? "";
    return message;
  },
};

messageTypeRegistry.set(CertificateFilter.$type, CertificateFilter);

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
