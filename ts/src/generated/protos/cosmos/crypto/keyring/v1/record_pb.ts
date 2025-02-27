// Since: cosmos-sdk 0.46

// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true"
// @generated from file cosmos/crypto/keyring/v1/record.proto (package cosmos.crypto.keyring.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import { file_gogoproto_gogo } from "../../../../gogoproto/gogo_pb";
import type { Any, AnyJson } from "@bufbuild/protobuf/wkt";
import { file_google_protobuf_any } from "@bufbuild/protobuf/wkt";
import type { BIP44Params, BIP44ParamsJson } from "../../hd/v1/hd_pb";
import { file_cosmos_crypto_hd_v1_hd } from "../../hd/v1/hd_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file cosmos/crypto/keyring/v1/record.proto.
 */
export const file_cosmos_crypto_keyring_v1_record: GenFile = /*@__PURE__*/
  fileDesc("CiVjb3Ntb3MvY3J5cHRvL2tleXJpbmcvdjEvcmVjb3JkLnByb3RvEhhjb3Ntb3MuY3J5cHRvLmtleXJpbmcudjEirgMKBlJlY29yZBIMCgRuYW1lGAEgASgJEiUKB3B1Yl9rZXkYAiABKAsyFC5nb29nbGUucHJvdG9idWYuQW55EjcKBWxvY2FsGAMgASgLMiYuY29zbW9zLmNyeXB0by5rZXlyaW5nLnYxLlJlY29yZC5Mb2NhbEgAEjkKBmxlZGdlchgEIAEoCzInLmNvc21vcy5jcnlwdG8ua2V5cmluZy52MS5SZWNvcmQuTGVkZ2VySAASNwoFbXVsdGkYBSABKAsyJi5jb3Ntb3MuY3J5cHRvLmtleXJpbmcudjEuUmVjb3JkLk11bHRpSAASOwoHb2ZmbGluZRgGIAEoCzIoLmNvc21vcy5jcnlwdG8ua2V5cmluZy52MS5SZWNvcmQuT2ZmbGluZUgAGi8KBUxvY2FsEiYKCHByaXZfa2V5GAEgASgLMhQuZ29vZ2xlLnByb3RvYnVmLkFueRo4CgZMZWRnZXISLgoEcGF0aBgBIAEoCzIgLmNvc21vcy5jcnlwdG8uaGQudjEuQklQNDRQYXJhbXMaBwoFTXVsdGkaCQoHT2ZmbGluZUIGCgRpdGVtQjVaK2dpdGh1Yi5jb20vY29zbW9zL2Nvc21vcy1zZGsvY3J5cHRvL2tleXJpbmfI4R4AmOMeAGIGcHJvdG8z", [file_gogoproto_gogo, file_google_protobuf_any, file_cosmos_crypto_hd_v1_hd]);

/**
 * Record is used for representing a key in the keyring.
 *
 * @generated from message cosmos.crypto.keyring.v1.Record
 */
export type Record = Message<"cosmos.crypto.keyring.v1.Record"> & {
  /**
   * name represents a name of Record
   *
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * pub_key represents a public key in any format
   *
   * @generated from field: google.protobuf.Any pub_key = 2;
   */
  pubKey?: Any;

  /**
   * Record contains one of the following items
   *
   * @generated from oneof cosmos.crypto.keyring.v1.Record.item
   */
  item: {
    /**
     * local stores the private key locally.
     *
     * @generated from field: cosmos.crypto.keyring.v1.Record.Local local = 3;
     */
    value: Record_Local;
    case: "local";
  } | {
    /**
     * ledger stores the information about a Ledger key.
     *
     * @generated from field: cosmos.crypto.keyring.v1.Record.Ledger ledger = 4;
     */
    value: Record_Ledger;
    case: "ledger";
  } | {
    /**
     * Multi does not store any other information.
     *
     * @generated from field: cosmos.crypto.keyring.v1.Record.Multi multi = 5;
     */
    value: Record_Multi;
    case: "multi";
  } | {
    /**
     * Offline does not store any other information.
     *
     * @generated from field: cosmos.crypto.keyring.v1.Record.Offline offline = 6;
     */
    value: Record_Offline;
    case: "offline";
  } | { case: undefined; value?: undefined };
};

/**
 * Record is used for representing a key in the keyring.
 *
 * @generated from message cosmos.crypto.keyring.v1.Record
 */
export type RecordJson = {
  /**
   * name represents a name of Record
   *
   * @generated from field: string name = 1;
   */
  name?: string;

  /**
   * pub_key represents a public key in any format
   *
   * @generated from field: google.protobuf.Any pub_key = 2;
   */
  pubKey?: AnyJson;

  /**
   * local stores the private key locally.
   *
   * @generated from field: cosmos.crypto.keyring.v1.Record.Local local = 3;
   */
  local?: Record_LocalJson;

  /**
   * ledger stores the information about a Ledger key.
   *
   * @generated from field: cosmos.crypto.keyring.v1.Record.Ledger ledger = 4;
   */
  ledger?: Record_LedgerJson;

  /**
   * Multi does not store any other information.
   *
   * @generated from field: cosmos.crypto.keyring.v1.Record.Multi multi = 5;
   */
  multi?: Record_MultiJson;

  /**
   * Offline does not store any other information.
   *
   * @generated from field: cosmos.crypto.keyring.v1.Record.Offline offline = 6;
   */
  offline?: Record_OfflineJson;
};

/**
 * Describes the message cosmos.crypto.keyring.v1.Record.
 * Use `create(RecordSchema)` to create a new message.
 */
export const RecordSchema: GenMessage<Record, RecordJson> = /*@__PURE__*/
  messageDesc(file_cosmos_crypto_keyring_v1_record, 0);

/**
 * Item is a keyring item stored in a keyring backend.
 * Local item
 *
 * @generated from message cosmos.crypto.keyring.v1.Record.Local
 */
export type Record_Local = Message<"cosmos.crypto.keyring.v1.Record.Local"> & {
  /**
   * @generated from field: google.protobuf.Any priv_key = 1;
   */
  privKey?: Any;
};

/**
 * Item is a keyring item stored in a keyring backend.
 * Local item
 *
 * @generated from message cosmos.crypto.keyring.v1.Record.Local
 */
export type Record_LocalJson = {
  /**
   * @generated from field: google.protobuf.Any priv_key = 1;
   */
  privKey?: AnyJson;
};

/**
 * Describes the message cosmos.crypto.keyring.v1.Record.Local.
 * Use `create(Record_LocalSchema)` to create a new message.
 */
export const Record_LocalSchema: GenMessage<Record_Local, Record_LocalJson> = /*@__PURE__*/
  messageDesc(file_cosmos_crypto_keyring_v1_record, 0, 0);

/**
 * Ledger item
 *
 * @generated from message cosmos.crypto.keyring.v1.Record.Ledger
 */
export type Record_Ledger = Message<"cosmos.crypto.keyring.v1.Record.Ledger"> & {
  /**
   * @generated from field: cosmos.crypto.hd.v1.BIP44Params path = 1;
   */
  path?: BIP44Params;
};

/**
 * Ledger item
 *
 * @generated from message cosmos.crypto.keyring.v1.Record.Ledger
 */
export type Record_LedgerJson = {
  /**
   * @generated from field: cosmos.crypto.hd.v1.BIP44Params path = 1;
   */
  path?: BIP44ParamsJson;
};

/**
 * Describes the message cosmos.crypto.keyring.v1.Record.Ledger.
 * Use `create(Record_LedgerSchema)` to create a new message.
 */
export const Record_LedgerSchema: GenMessage<Record_Ledger, Record_LedgerJson> = /*@__PURE__*/
  messageDesc(file_cosmos_crypto_keyring_v1_record, 0, 1);

/**
 * Multi item
 *
 * @generated from message cosmos.crypto.keyring.v1.Record.Multi
 */
export type Record_Multi = Message<"cosmos.crypto.keyring.v1.Record.Multi"> & {
};

/**
 * Multi item
 *
 * @generated from message cosmos.crypto.keyring.v1.Record.Multi
 */
export type Record_MultiJson = {
};

/**
 * Describes the message cosmos.crypto.keyring.v1.Record.Multi.
 * Use `create(Record_MultiSchema)` to create a new message.
 */
export const Record_MultiSchema: GenMessage<Record_Multi, Record_MultiJson> = /*@__PURE__*/
  messageDesc(file_cosmos_crypto_keyring_v1_record, 0, 2);

/**
 * Offline item
 *
 * @generated from message cosmos.crypto.keyring.v1.Record.Offline
 */
export type Record_Offline = Message<"cosmos.crypto.keyring.v1.Record.Offline"> & {
};

/**
 * Offline item
 *
 * @generated from message cosmos.crypto.keyring.v1.Record.Offline
 */
export type Record_OfflineJson = {
};

/**
 * Describes the message cosmos.crypto.keyring.v1.Record.Offline.
 * Use `create(Record_OfflineSchema)` to create a new message.
 */
export const Record_OfflineSchema: GenMessage<Record_Offline, Record_OfflineJson> = /*@__PURE__*/
  messageDesc(file_cosmos_crypto_keyring_v1_record, 0, 3);

