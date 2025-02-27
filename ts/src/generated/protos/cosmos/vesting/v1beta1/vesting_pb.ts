// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true"
// @generated from file cosmos/vesting/v1beta1/vesting.proto (package cosmos.vesting.v1beta1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import { file_amino_amino } from "../../../amino/amino_pb";
import { file_gogoproto_gogo } from "../../../gogoproto/gogo_pb";
import type { Coin, CoinJson } from "../../base/v1beta1/coin_pb";
import { file_cosmos_base_v1beta1_coin } from "../../base/v1beta1/coin_pb";
import type { BaseAccount, BaseAccountJson } from "../../auth/v1beta1/auth_pb";
import { file_cosmos_auth_v1beta1_auth } from "../../auth/v1beta1/auth_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file cosmos/vesting/v1beta1/vesting.proto.
 */
export const file_cosmos_vesting_v1beta1_vesting: GenFile = /*@__PURE__*/
  fileDesc("CiRjb3Ntb3MvdmVzdGluZy92MWJldGExL3Zlc3RpbmcucHJvdG8SFmNvc21vcy52ZXN0aW5nLnYxYmV0YTEi0wMKEkJhc2VWZXN0aW5nQWNjb3VudBI8CgxiYXNlX2FjY291bnQYASABKAsyIC5jb3Ntb3MuYXV0aC52MWJldGExLkJhc2VBY2NvdW50QgTQ3h8BEmoKEG9yaWdpbmFsX3Zlc3RpbmcYAiADKAsyGS5jb3Ntb3MuYmFzZS52MWJldGExLkNvaW5CNcjeHwCq3x8oZ2l0aHViLmNvbS9jb3Ntb3MvY29zbW9zLXNkay90eXBlcy5Db2luc6jnsCoBEmgKDmRlbGVnYXRlZF9mcmVlGAMgAygLMhkuY29zbW9zLmJhc2UudjFiZXRhMS5Db2luQjXI3h8Aqt8fKGdpdGh1Yi5jb20vY29zbW9zL2Nvc21vcy1zZGsvdHlwZXMuQ29pbnOo57AqARJrChFkZWxlZ2F0ZWRfdmVzdGluZxgEIAMoCzIZLmNvc21vcy5iYXNlLnYxYmV0YTEuQ29pbkI1yN4fAKrfHyhnaXRodWIuY29tL2Nvc21vcy9jb3Ntb3Mtc2RrL3R5cGVzLkNvaW5zqOewKgESEAoIZW5kX3RpbWUYBSABKAM6KoigHwCYoB8AiuewKh1jb3Ntb3Mtc2RrL0Jhc2VWZXN0aW5nQWNjb3VudCKwAQoYQ29udGludW91c1Zlc3RpbmdBY2NvdW50Ek4KFGJhc2VfdmVzdGluZ19hY2NvdW50GAEgASgLMiouY29zbW9zLnZlc3RpbmcudjFiZXRhMS5CYXNlVmVzdGluZ0FjY291bnRCBNDeHwESEgoKc3RhcnRfdGltZRgCIAEoAzowiKAfAJigHwCK57AqI2Nvc21vcy1zZGsvQ29udGludW91c1Zlc3RpbmdBY2NvdW50IpYBChVEZWxheWVkVmVzdGluZ0FjY291bnQSTgoUYmFzZV92ZXN0aW5nX2FjY291bnQYASABKAsyKi5jb3Ntb3MudmVzdGluZy52MWJldGExLkJhc2VWZXN0aW5nQWNjb3VudEIE0N4fATotiKAfAJigHwCK57AqIGNvc21vcy1zZGsvRGVsYXllZFZlc3RpbmdBY2NvdW50IoABCgZQZXJpb2QSDgoGbGVuZ3RoGAEgASgDEmAKBmFtb3VudBgCIAMoCzIZLmNvc21vcy5iYXNlLnYxYmV0YTEuQ29pbkI1yN4fAKrfHyhnaXRodWIuY29tL2Nvc21vcy9jb3Ntb3Mtc2RrL3R5cGVzLkNvaW5zqOewKgE6BJigHwAi8AEKFlBlcmlvZGljVmVzdGluZ0FjY291bnQSTgoUYmFzZV92ZXN0aW5nX2FjY291bnQYASABKAsyKi5jb3Ntb3MudmVzdGluZy52MWJldGExLkJhc2VWZXN0aW5nQWNjb3VudEIE0N4fARISCgpzdGFydF90aW1lGAIgASgDEkIKD3Zlc3RpbmdfcGVyaW9kcxgDIAMoCzIeLmNvc21vcy52ZXN0aW5nLnYxYmV0YTEuUGVyaW9kQgnI3h8AqOewKgE6LoigHwCYoB8AiuewKiFjb3Ntb3Mtc2RrL1BlcmlvZGljVmVzdGluZ0FjY291bnQimAEKFlBlcm1hbmVudExvY2tlZEFjY291bnQSTgoUYmFzZV92ZXN0aW5nX2FjY291bnQYASABKAsyKi5jb3Ntb3MudmVzdGluZy52MWJldGExLkJhc2VWZXN0aW5nQWNjb3VudEIE0N4fATouiKAfAJigHwCK57AqIWNvc21vcy1zZGsvUGVybWFuZW50TG9ja2VkQWNjb3VudEIzWjFnaXRodWIuY29tL2Nvc21vcy9jb3Ntb3Mtc2RrL3gvYXV0aC92ZXN0aW5nL3R5cGVzYgZwcm90bzM", [file_amino_amino, file_gogoproto_gogo, file_cosmos_base_v1beta1_coin, file_cosmos_auth_v1beta1_auth]);

/**
 * BaseVestingAccount implements the VestingAccount interface. It contains all
 * the necessary fields needed for any vesting account implementation.
 *
 * @generated from message cosmos.vesting.v1beta1.BaseVestingAccount
 */
export type BaseVestingAccount = Message<"cosmos.vesting.v1beta1.BaseVestingAccount"> & {
  /**
   * @generated from field: cosmos.auth.v1beta1.BaseAccount base_account = 1;
   */
  baseAccount?: BaseAccount;

  /**
   * @generated from field: repeated cosmos.base.v1beta1.Coin original_vesting = 2;
   */
  originalVesting: Coin[];

  /**
   * @generated from field: repeated cosmos.base.v1beta1.Coin delegated_free = 3;
   */
  delegatedFree: Coin[];

  /**
   * @generated from field: repeated cosmos.base.v1beta1.Coin delegated_vesting = 4;
   */
  delegatedVesting: Coin[];

  /**
   * Vesting end time, as unix timestamp (in seconds).
   *
   * @generated from field: int64 end_time = 5;
   */
  endTime: bigint;
};

/**
 * BaseVestingAccount implements the VestingAccount interface. It contains all
 * the necessary fields needed for any vesting account implementation.
 *
 * @generated from message cosmos.vesting.v1beta1.BaseVestingAccount
 */
export type BaseVestingAccountJson = {
  /**
   * @generated from field: cosmos.auth.v1beta1.BaseAccount base_account = 1;
   */
  baseAccount?: BaseAccountJson;

  /**
   * @generated from field: repeated cosmos.base.v1beta1.Coin original_vesting = 2;
   */
  originalVesting?: CoinJson[];

  /**
   * @generated from field: repeated cosmos.base.v1beta1.Coin delegated_free = 3;
   */
  delegatedFree?: CoinJson[];

  /**
   * @generated from field: repeated cosmos.base.v1beta1.Coin delegated_vesting = 4;
   */
  delegatedVesting?: CoinJson[];

  /**
   * Vesting end time, as unix timestamp (in seconds).
   *
   * @generated from field: int64 end_time = 5;
   */
  endTime?: string;
};

/**
 * Describes the message cosmos.vesting.v1beta1.BaseVestingAccount.
 * Use `create(BaseVestingAccountSchema)` to create a new message.
 */
export const BaseVestingAccountSchema: GenMessage<BaseVestingAccount, BaseVestingAccountJson> = /*@__PURE__*/
  messageDesc(file_cosmos_vesting_v1beta1_vesting, 0);

/**
 * ContinuousVestingAccount implements the VestingAccount interface. It
 * continuously vests by unlocking coins linearly with respect to time.
 *
 * @generated from message cosmos.vesting.v1beta1.ContinuousVestingAccount
 */
export type ContinuousVestingAccount = Message<"cosmos.vesting.v1beta1.ContinuousVestingAccount"> & {
  /**
   * @generated from field: cosmos.vesting.v1beta1.BaseVestingAccount base_vesting_account = 1;
   */
  baseVestingAccount?: BaseVestingAccount;

  /**
   * Vesting start time, as unix timestamp (in seconds).
   *
   * @generated from field: int64 start_time = 2;
   */
  startTime: bigint;
};

/**
 * ContinuousVestingAccount implements the VestingAccount interface. It
 * continuously vests by unlocking coins linearly with respect to time.
 *
 * @generated from message cosmos.vesting.v1beta1.ContinuousVestingAccount
 */
export type ContinuousVestingAccountJson = {
  /**
   * @generated from field: cosmos.vesting.v1beta1.BaseVestingAccount base_vesting_account = 1;
   */
  baseVestingAccount?: BaseVestingAccountJson;

  /**
   * Vesting start time, as unix timestamp (in seconds).
   *
   * @generated from field: int64 start_time = 2;
   */
  startTime?: string;
};

/**
 * Describes the message cosmos.vesting.v1beta1.ContinuousVestingAccount.
 * Use `create(ContinuousVestingAccountSchema)` to create a new message.
 */
export const ContinuousVestingAccountSchema: GenMessage<ContinuousVestingAccount, ContinuousVestingAccountJson> = /*@__PURE__*/
  messageDesc(file_cosmos_vesting_v1beta1_vesting, 1);

/**
 * DelayedVestingAccount implements the VestingAccount interface. It vests all
 * coins after a specific time, but non prior. In other words, it keeps them
 * locked until a specified time.
 *
 * @generated from message cosmos.vesting.v1beta1.DelayedVestingAccount
 */
export type DelayedVestingAccount = Message<"cosmos.vesting.v1beta1.DelayedVestingAccount"> & {
  /**
   * @generated from field: cosmos.vesting.v1beta1.BaseVestingAccount base_vesting_account = 1;
   */
  baseVestingAccount?: BaseVestingAccount;
};

/**
 * DelayedVestingAccount implements the VestingAccount interface. It vests all
 * coins after a specific time, but non prior. In other words, it keeps them
 * locked until a specified time.
 *
 * @generated from message cosmos.vesting.v1beta1.DelayedVestingAccount
 */
export type DelayedVestingAccountJson = {
  /**
   * @generated from field: cosmos.vesting.v1beta1.BaseVestingAccount base_vesting_account = 1;
   */
  baseVestingAccount?: BaseVestingAccountJson;
};

/**
 * Describes the message cosmos.vesting.v1beta1.DelayedVestingAccount.
 * Use `create(DelayedVestingAccountSchema)` to create a new message.
 */
export const DelayedVestingAccountSchema: GenMessage<DelayedVestingAccount, DelayedVestingAccountJson> = /*@__PURE__*/
  messageDesc(file_cosmos_vesting_v1beta1_vesting, 2);

/**
 * Period defines a length of time and amount of coins that will vest.
 *
 * @generated from message cosmos.vesting.v1beta1.Period
 */
export type Period = Message<"cosmos.vesting.v1beta1.Period"> & {
  /**
   * Period duration in seconds.
   *
   * @generated from field: int64 length = 1;
   */
  length: bigint;

  /**
   * @generated from field: repeated cosmos.base.v1beta1.Coin amount = 2;
   */
  amount: Coin[];
};

/**
 * Period defines a length of time and amount of coins that will vest.
 *
 * @generated from message cosmos.vesting.v1beta1.Period
 */
export type PeriodJson = {
  /**
   * Period duration in seconds.
   *
   * @generated from field: int64 length = 1;
   */
  length?: string;

  /**
   * @generated from field: repeated cosmos.base.v1beta1.Coin amount = 2;
   */
  amount?: CoinJson[];
};

/**
 * Describes the message cosmos.vesting.v1beta1.Period.
 * Use `create(PeriodSchema)` to create a new message.
 */
export const PeriodSchema: GenMessage<Period, PeriodJson> = /*@__PURE__*/
  messageDesc(file_cosmos_vesting_v1beta1_vesting, 3);

/**
 * PeriodicVestingAccount implements the VestingAccount interface. It
 * periodically vests by unlocking coins during each specified period.
 *
 * @generated from message cosmos.vesting.v1beta1.PeriodicVestingAccount
 */
export type PeriodicVestingAccount = Message<"cosmos.vesting.v1beta1.PeriodicVestingAccount"> & {
  /**
   * @generated from field: cosmos.vesting.v1beta1.BaseVestingAccount base_vesting_account = 1;
   */
  baseVestingAccount?: BaseVestingAccount;

  /**
   * @generated from field: int64 start_time = 2;
   */
  startTime: bigint;

  /**
   * @generated from field: repeated cosmos.vesting.v1beta1.Period vesting_periods = 3;
   */
  vestingPeriods: Period[];
};

/**
 * PeriodicVestingAccount implements the VestingAccount interface. It
 * periodically vests by unlocking coins during each specified period.
 *
 * @generated from message cosmos.vesting.v1beta1.PeriodicVestingAccount
 */
export type PeriodicVestingAccountJson = {
  /**
   * @generated from field: cosmos.vesting.v1beta1.BaseVestingAccount base_vesting_account = 1;
   */
  baseVestingAccount?: BaseVestingAccountJson;

  /**
   * @generated from field: int64 start_time = 2;
   */
  startTime?: string;

  /**
   * @generated from field: repeated cosmos.vesting.v1beta1.Period vesting_periods = 3;
   */
  vestingPeriods?: PeriodJson[];
};

/**
 * Describes the message cosmos.vesting.v1beta1.PeriodicVestingAccount.
 * Use `create(PeriodicVestingAccountSchema)` to create a new message.
 */
export const PeriodicVestingAccountSchema: GenMessage<PeriodicVestingAccount, PeriodicVestingAccountJson> = /*@__PURE__*/
  messageDesc(file_cosmos_vesting_v1beta1_vesting, 4);

/**
 * PermanentLockedAccount implements the VestingAccount interface. It does
 * not ever release coins, locking them indefinitely. Coins in this account can
 * still be used for delegating and for governance votes even while locked.
 *
 * Since: cosmos-sdk 0.43
 *
 * @generated from message cosmos.vesting.v1beta1.PermanentLockedAccount
 */
export type PermanentLockedAccount = Message<"cosmos.vesting.v1beta1.PermanentLockedAccount"> & {
  /**
   * @generated from field: cosmos.vesting.v1beta1.BaseVestingAccount base_vesting_account = 1;
   */
  baseVestingAccount?: BaseVestingAccount;
};

/**
 * PermanentLockedAccount implements the VestingAccount interface. It does
 * not ever release coins, locking them indefinitely. Coins in this account can
 * still be used for delegating and for governance votes even while locked.
 *
 * Since: cosmos-sdk 0.43
 *
 * @generated from message cosmos.vesting.v1beta1.PermanentLockedAccount
 */
export type PermanentLockedAccountJson = {
  /**
   * @generated from field: cosmos.vesting.v1beta1.BaseVestingAccount base_vesting_account = 1;
   */
  baseVestingAccount?: BaseVestingAccountJson;
};

/**
 * Describes the message cosmos.vesting.v1beta1.PermanentLockedAccount.
 * Use `create(PermanentLockedAccountSchema)` to create a new message.
 */
export const PermanentLockedAccountSchema: GenMessage<PermanentLockedAccount, PermanentLockedAccountJson> = /*@__PURE__*/
  messageDesc(file_cosmos_vesting_v1beta1_vesting, 5);

