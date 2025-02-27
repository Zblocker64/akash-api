// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true"
// @generated from file cosmos/slashing/v1beta1/slashing.proto (package cosmos.slashing.v1beta1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import { file_gogoproto_gogo } from "../../../gogoproto/gogo_pb";
import type { Duration, DurationJson, Timestamp, TimestampJson } from "@bufbuild/protobuf/wkt";
import { file_google_protobuf_duration, file_google_protobuf_timestamp } from "@bufbuild/protobuf/wkt";
import { file_cosmos_proto_cosmos } from "../../../cosmos_proto/cosmos_pb";
import { file_amino_amino } from "../../../amino/amino_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file cosmos/slashing/v1beta1/slashing.proto.
 */
export const file_cosmos_slashing_v1beta1_slashing: GenFile = /*@__PURE__*/
  fileDesc("CiZjb3Ntb3Mvc2xhc2hpbmcvdjFiZXRhMS9zbGFzaGluZy5wcm90bxIXY29zbW9zLnNsYXNoaW5nLnYxYmV0YTEi6wEKFFZhbGlkYXRvclNpZ25pbmdJbmZvEikKB2FkZHJlc3MYASABKAlCGNK0LRRjb3Ntb3MuQWRkcmVzc1N0cmluZxIUCgxzdGFydF9oZWlnaHQYAiABKAMSFAoMaW5kZXhfb2Zmc2V0GAMgASgDEj8KDGphaWxlZF91bnRpbBgEIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXBCDcjeHwCQ3x8BqOewKgESEgoKdG9tYnN0b25lZBgFIAEoCBIdChVtaXNzZWRfYmxvY2tzX2NvdW50ZXIYBiABKAM6CJigHwDooB8BIpYDCgZQYXJhbXMSHAoUc2lnbmVkX2Jsb2Nrc193aW5kb3cYASABKAMSUgoVbWluX3NpZ25lZF9wZXJfd2luZG93GAIgASgMQjPI3h8A2t4fJmdpdGh1Yi5jb20vY29zbW9zL2Nvc21vcy1zZGsvdHlwZXMuRGVjqOewKgESSAoWZG93bnRpbWVfamFpbF9kdXJhdGlvbhgDIAEoCzIZLmdvb2dsZS5wcm90b2J1Zi5EdXJhdGlvbkINyN4fAJjfHwGo57AqARJXChpzbGFzaF9mcmFjdGlvbl9kb3VibGVfc2lnbhgEIAEoDEIzyN4fANreHyZnaXRodWIuY29tL2Nvc21vcy9jb3Ntb3Mtc2RrL3R5cGVzLkRlY6jnsCoBElQKF3NsYXNoX2ZyYWN0aW9uX2Rvd250aW1lGAUgASgMQjPI3h8A2t4fJmdpdGh1Yi5jb20vY29zbW9zL2Nvc21vcy1zZGsvdHlwZXMuRGVjqOewKgE6IYrnsCocY29zbW9zLXNkay94L3NsYXNoaW5nL1BhcmFtc0IzWi1naXRodWIuY29tL2Nvc21vcy9jb3Ntb3Mtc2RrL3gvc2xhc2hpbmcvdHlwZXOo4h4BYgZwcm90bzM", [file_gogoproto_gogo, file_google_protobuf_duration, file_google_protobuf_timestamp, file_cosmos_proto_cosmos, file_amino_amino]);

/**
 * ValidatorSigningInfo defines a validator's signing info for monitoring their
 * liveness activity.
 *
 * @generated from message cosmos.slashing.v1beta1.ValidatorSigningInfo
 */
export type ValidatorSigningInfo = Message<"cosmos.slashing.v1beta1.ValidatorSigningInfo"> & {
  /**
   * @generated from field: string address = 1;
   */
  address: string;

  /**
   * Height at which validator was first a candidate OR was unjailed
   *
   * @generated from field: int64 start_height = 2;
   */
  startHeight: bigint;

  /**
   * Index which is incremented each time the validator was a bonded
   * in a block and may have signed a precommit or not. This in conjunction with the
   * `SignedBlocksWindow` param determines the index in the `MissedBlocksBitArray`.
   *
   * @generated from field: int64 index_offset = 3;
   */
  indexOffset: bigint;

  /**
   * Timestamp until which the validator is jailed due to liveness downtime.
   *
   * @generated from field: google.protobuf.Timestamp jailed_until = 4;
   */
  jailedUntil?: Timestamp;

  /**
   * Whether or not a validator has been tombstoned (killed out of validator set). It is set
   * once the validator commits an equivocation or for any other configured misbehiavor.
   *
   * @generated from field: bool tombstoned = 5;
   */
  tombstoned: boolean;

  /**
   * A counter kept to avoid unnecessary array reads.
   * Note that `Sum(MissedBlocksBitArray)` always equals `MissedBlocksCounter`.
   *
   * @generated from field: int64 missed_blocks_counter = 6;
   */
  missedBlocksCounter: bigint;
};

/**
 * ValidatorSigningInfo defines a validator's signing info for monitoring their
 * liveness activity.
 *
 * @generated from message cosmos.slashing.v1beta1.ValidatorSigningInfo
 */
export type ValidatorSigningInfoJson = {
  /**
   * @generated from field: string address = 1;
   */
  address?: string;

  /**
   * Height at which validator was first a candidate OR was unjailed
   *
   * @generated from field: int64 start_height = 2;
   */
  startHeight?: string;

  /**
   * Index which is incremented each time the validator was a bonded
   * in a block and may have signed a precommit or not. This in conjunction with the
   * `SignedBlocksWindow` param determines the index in the `MissedBlocksBitArray`.
   *
   * @generated from field: int64 index_offset = 3;
   */
  indexOffset?: string;

  /**
   * Timestamp until which the validator is jailed due to liveness downtime.
   *
   * @generated from field: google.protobuf.Timestamp jailed_until = 4;
   */
  jailedUntil?: TimestampJson;

  /**
   * Whether or not a validator has been tombstoned (killed out of validator set). It is set
   * once the validator commits an equivocation or for any other configured misbehiavor.
   *
   * @generated from field: bool tombstoned = 5;
   */
  tombstoned?: boolean;

  /**
   * A counter kept to avoid unnecessary array reads.
   * Note that `Sum(MissedBlocksBitArray)` always equals `MissedBlocksCounter`.
   *
   * @generated from field: int64 missed_blocks_counter = 6;
   */
  missedBlocksCounter?: string;
};

/**
 * Describes the message cosmos.slashing.v1beta1.ValidatorSigningInfo.
 * Use `create(ValidatorSigningInfoSchema)` to create a new message.
 */
export const ValidatorSigningInfoSchema: GenMessage<ValidatorSigningInfo, ValidatorSigningInfoJson> = /*@__PURE__*/
  messageDesc(file_cosmos_slashing_v1beta1_slashing, 0);

/**
 * Params represents the parameters used for by the slashing module.
 *
 * @generated from message cosmos.slashing.v1beta1.Params
 */
export type Params = Message<"cosmos.slashing.v1beta1.Params"> & {
  /**
   * @generated from field: int64 signed_blocks_window = 1;
   */
  signedBlocksWindow: bigint;

  /**
   * @generated from field: bytes min_signed_per_window = 2;
   */
  minSignedPerWindow: Uint8Array;

  /**
   * @generated from field: google.protobuf.Duration downtime_jail_duration = 3;
   */
  downtimeJailDuration?: Duration;

  /**
   * @generated from field: bytes slash_fraction_double_sign = 4;
   */
  slashFractionDoubleSign: Uint8Array;

  /**
   * @generated from field: bytes slash_fraction_downtime = 5;
   */
  slashFractionDowntime: Uint8Array;
};

/**
 * Params represents the parameters used for by the slashing module.
 *
 * @generated from message cosmos.slashing.v1beta1.Params
 */
export type ParamsJson = {
  /**
   * @generated from field: int64 signed_blocks_window = 1;
   */
  signedBlocksWindow?: string;

  /**
   * @generated from field: bytes min_signed_per_window = 2;
   */
  minSignedPerWindow?: string;

  /**
   * @generated from field: google.protobuf.Duration downtime_jail_duration = 3;
   */
  downtimeJailDuration?: DurationJson;

  /**
   * @generated from field: bytes slash_fraction_double_sign = 4;
   */
  slashFractionDoubleSign?: string;

  /**
   * @generated from field: bytes slash_fraction_downtime = 5;
   */
  slashFractionDowntime?: string;
};

/**
 * Describes the message cosmos.slashing.v1beta1.Params.
 * Use `create(ParamsSchema)` to create a new message.
 */
export const ParamsSchema: GenMessage<Params, ParamsJson> = /*@__PURE__*/
  messageDesc(file_cosmos_slashing_v1beta1_slashing, 1);

