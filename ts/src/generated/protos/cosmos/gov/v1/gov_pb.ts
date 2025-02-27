// Since: cosmos-sdk 0.46

// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true"
// @generated from file cosmos/gov/v1/gov.proto (package cosmos.gov.v1, syntax proto3)
/* eslint-disable */

import type { GenEnum, GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { enumDesc, fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import type { Coin, CoinJson } from "../../base/v1beta1/coin_pb";
import { file_cosmos_base_v1beta1_coin } from "../../base/v1beta1/coin_pb";
import { file_gogoproto_gogo } from "../../../gogoproto/gogo_pb";
import type { Any, AnyJson, Duration, DurationJson, Timestamp, TimestampJson } from "@bufbuild/protobuf/wkt";
import { file_google_protobuf_any, file_google_protobuf_duration, file_google_protobuf_timestamp } from "@bufbuild/protobuf/wkt";
import { file_cosmos_proto_cosmos } from "../../../cosmos_proto/cosmos_pb";
import { file_amino_amino } from "../../../amino/amino_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file cosmos/gov/v1/gov.proto.
 */
export const file_cosmos_gov_v1_gov: GenFile = /*@__PURE__*/
  fileDesc("Chdjb3Ntb3MvZ292L3YxL2dvdi5wcm90bxINY29zbW9zLmdvdi52MSJfChJXZWlnaHRlZFZvdGVPcHRpb24SKQoGb3B0aW9uGAEgASgOMhkuY29zbW9zLmdvdi52MS5Wb3RlT3B0aW9uEh4KBndlaWdodBgCIAEoCUIO0rQtCmNvc21vcy5EZWMigQEKB0RlcG9zaXQSEwoLcHJvcG9zYWxfaWQYASABKAQSKwoJZGVwb3NpdG9yGAIgASgJQhjStC0UY29zbW9zLkFkZHJlc3NTdHJpbmcSNAoGYW1vdW50GAMgAygLMhkuY29zbW9zLmJhc2UudjFiZXRhMS5Db2luQgnI3h8AqOewKgEivgQKCFByb3Bvc2FsEgoKAmlkGAEgASgEEiYKCG1lc3NhZ2VzGAIgAygLMhQuZ29vZ2xlLnByb3RvYnVmLkFueRItCgZzdGF0dXMYAyABKA4yHS5jb3Ntb3MuZ292LnYxLlByb3Bvc2FsU3RhdHVzEjYKEmZpbmFsX3RhbGx5X3Jlc3VsdBgEIAEoCzIaLmNvc21vcy5nb3YudjEuVGFsbHlSZXN1bHQSNQoLc3VibWl0X3RpbWUYBSABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wQgSQ3x8BEjoKEGRlcG9zaXRfZW5kX3RpbWUYBiABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wQgSQ3x8BEjsKDXRvdGFsX2RlcG9zaXQYByADKAsyGS5jb3Ntb3MuYmFzZS52MWJldGExLkNvaW5CCcjeHwCo57AqARI7ChF2b3Rpbmdfc3RhcnRfdGltZRgIIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXBCBJDfHwESOQoPdm90aW5nX2VuZF90aW1lGAkgASgLMhouZ29vZ2xlLnByb3RvYnVmLlRpbWVzdGFtcEIEkN8fARIQCghtZXRhZGF0YRgKIAEoCRINCgV0aXRsZRgLIAEoCRIPCgdzdW1tYXJ5GAwgASgJEioKCHByb3Bvc2VyGA0gASgJQhjStC0UY29zbW9zLkFkZHJlc3NTdHJpbmcSEQoJZXhwZWRpdGVkGA4gASgIIqUBCgtUYWxseVJlc3VsdBIhCgl5ZXNfY291bnQYASABKAlCDtK0LQpjb3Ntb3MuSW50EiUKDWFic3RhaW5fY291bnQYAiABKAlCDtK0LQpjb3Ntb3MuSW50EiAKCG5vX2NvdW50GAMgASgJQg7StC0KY29zbW9zLkludBIqChJub193aXRoX3ZldG9fY291bnQYBCABKAlCDtK0LQpjb3Ntb3MuSW50IpABCgRWb3RlEhMKC3Byb3Bvc2FsX2lkGAEgASgEEicKBXZvdGVyGAIgASgJQhjStC0UY29zbW9zLkFkZHJlc3NTdHJpbmcSMgoHb3B0aW9ucxgEIAMoCzIhLmNvc21vcy5nb3YudjEuV2VpZ2h0ZWRWb3RlT3B0aW9uEhAKCG1ldGFkYXRhGAUgASgJSgQIAxAEIrsBCg1EZXBvc2l0UGFyYW1zEk0KC21pbl9kZXBvc2l0GAEgAygLMhkuY29zbW9zLmJhc2UudjFiZXRhMS5Db2luQh3I3h8A6t4fFW1pbl9kZXBvc2l0LG9taXRlbXB0eRJbChJtYXhfZGVwb3NpdF9wZXJpb2QYAiABKAsyGS5nb29nbGUucHJvdG9idWYuRHVyYXRpb25CJOreHxxtYXhfZGVwb3NpdF9wZXJpb2Qsb21pdGVtcHR5mN8fASJGCgxWb3RpbmdQYXJhbXMSNgoNdm90aW5nX3BlcmlvZBgBIAEoCzIZLmdvb2dsZS5wcm90b2J1Zi5EdXJhdGlvbkIEmN8fASJ4CgtUYWxseVBhcmFtcxIeCgZxdW9ydW0YASABKAlCDtK0LQpjb3Ntb3MuRGVjEiEKCXRocmVzaG9sZBgCIAEoCUIO0rQtCmNvc21vcy5EZWMSJgoOdmV0b190aHJlc2hvbGQYAyABKAlCDtK0LQpjb3Ntb3MuRGVjIo4FCgZQYXJhbXMSOQoLbWluX2RlcG9zaXQYASADKAsyGS5jb3Ntb3MuYmFzZS52MWJldGExLkNvaW5CCcjeHwCo57AqARI7ChJtYXhfZGVwb3NpdF9wZXJpb2QYAiABKAsyGS5nb29nbGUucHJvdG9idWYuRHVyYXRpb25CBJjfHwESNgoNdm90aW5nX3BlcmlvZBgDIAEoCzIZLmdvb2dsZS5wcm90b2J1Zi5EdXJhdGlvbkIEmN8fARIeCgZxdW9ydW0YBCABKAlCDtK0LQpjb3Ntb3MuRGVjEiEKCXRocmVzaG9sZBgFIAEoCUIO0rQtCmNvc21vcy5EZWMSJgoOdmV0b190aHJlc2hvbGQYBiABKAlCDtK0LQpjb3Ntb3MuRGVjEjEKGW1pbl9pbml0aWFsX2RlcG9zaXRfcmF0aW8YByABKAlCDtK0LQpjb3Ntb3MuRGVjEkAKF2V4cGVkaXRlZF92b3RpbmdfcGVyaW9kGAogASgLMhkuZ29vZ2xlLnByb3RvYnVmLkR1cmF0aW9uQgSY3x8BEisKE2V4cGVkaXRlZF90aHJlc2hvbGQYCyABKAlCDtK0LQpjb3Ntb3MuRGVjEkMKFWV4cGVkaXRlZF9taW5fZGVwb3NpdBgMIAMoCzIZLmNvc21vcy5iYXNlLnYxYmV0YTEuQ29pbkIJyN4fAKjnsCoBEhgKEGJ1cm5fdm90ZV9xdW9ydW0YDSABKAgSJQodYnVybl9wcm9wb3NhbF9kZXBvc2l0X3ByZXZvdGUYDiABKAgSFgoOYnVybl92b3RlX3ZldG8YDyABKAgSKQoRbWluX2RlcG9zaXRfcmF0aW8YECABKAlCDtK0LQpjb3Ntb3MuRGVjKokBCgpWb3RlT3B0aW9uEhsKF1ZPVEVfT1BUSU9OX1VOU1BFQ0lGSUVEEAASEwoPVk9URV9PUFRJT05fWUVTEAESFwoTVk9URV9PUFRJT05fQUJTVEFJThACEhIKDlZPVEVfT1BUSU9OX05PEAMSHAoYVk9URV9PUFRJT05fTk9fV0lUSF9WRVRPEAQqzgEKDlByb3Bvc2FsU3RhdHVzEh8KG1BST1BPU0FMX1NUQVRVU19VTlNQRUNJRklFRBAAEiIKHlBST1BPU0FMX1NUQVRVU19ERVBPU0lUX1BFUklPRBABEiEKHVBST1BPU0FMX1NUQVRVU19WT1RJTkdfUEVSSU9EEAISGgoWUFJPUE9TQUxfU1RBVFVTX1BBU1NFRBADEhwKGFBST1BPU0FMX1NUQVRVU19SRUpFQ1RFRBAEEhoKFlBST1BPU0FMX1NUQVRVU19GQUlMRUQQBUItWitnaXRodWIuY29tL2Nvc21vcy9jb3Ntb3Mtc2RrL3gvZ292L3R5cGVzL3YxYgZwcm90bzM", [file_cosmos_base_v1beta1_coin, file_gogoproto_gogo, file_google_protobuf_timestamp, file_google_protobuf_any, file_google_protobuf_duration, file_cosmos_proto_cosmos, file_amino_amino]);

/**
 * WeightedVoteOption defines a unit of vote for vote split.
 *
 * @generated from message cosmos.gov.v1.WeightedVoteOption
 */
export type WeightedVoteOption = Message<"cosmos.gov.v1.WeightedVoteOption"> & {
  /**
   * option defines the valid vote options, it must not contain duplicate vote options.
   *
   * @generated from field: cosmos.gov.v1.VoteOption option = 1;
   */
  option: VoteOption;

  /**
   * weight is the vote weight associated with the vote option.
   *
   * @generated from field: string weight = 2;
   */
  weight: string;
};

/**
 * WeightedVoteOption defines a unit of vote for vote split.
 *
 * @generated from message cosmos.gov.v1.WeightedVoteOption
 */
export type WeightedVoteOptionJson = {
  /**
   * option defines the valid vote options, it must not contain duplicate vote options.
   *
   * @generated from field: cosmos.gov.v1.VoteOption option = 1;
   */
  option?: VoteOptionJson;

  /**
   * weight is the vote weight associated with the vote option.
   *
   * @generated from field: string weight = 2;
   */
  weight?: string;
};

/**
 * Describes the message cosmos.gov.v1.WeightedVoteOption.
 * Use `create(WeightedVoteOptionSchema)` to create a new message.
 */
export const WeightedVoteOptionSchema: GenMessage<WeightedVoteOption, WeightedVoteOptionJson> = /*@__PURE__*/
  messageDesc(file_cosmos_gov_v1_gov, 0);

/**
 * Deposit defines an amount deposited by an account address to an active
 * proposal.
 *
 * @generated from message cosmos.gov.v1.Deposit
 */
export type Deposit = Message<"cosmos.gov.v1.Deposit"> & {
  /**
   * proposal_id defines the unique id of the proposal.
   *
   * @generated from field: uint64 proposal_id = 1;
   */
  proposalId: bigint;

  /**
   * depositor defines the deposit addresses from the proposals.
   *
   * @generated from field: string depositor = 2;
   */
  depositor: string;

  /**
   * amount to be deposited by depositor.
   *
   * @generated from field: repeated cosmos.base.v1beta1.Coin amount = 3;
   */
  amount: Coin[];
};

/**
 * Deposit defines an amount deposited by an account address to an active
 * proposal.
 *
 * @generated from message cosmos.gov.v1.Deposit
 */
export type DepositJson = {
  /**
   * proposal_id defines the unique id of the proposal.
   *
   * @generated from field: uint64 proposal_id = 1;
   */
  proposalId?: string;

  /**
   * depositor defines the deposit addresses from the proposals.
   *
   * @generated from field: string depositor = 2;
   */
  depositor?: string;

  /**
   * amount to be deposited by depositor.
   *
   * @generated from field: repeated cosmos.base.v1beta1.Coin amount = 3;
   */
  amount?: CoinJson[];
};

/**
 * Describes the message cosmos.gov.v1.Deposit.
 * Use `create(DepositSchema)` to create a new message.
 */
export const DepositSchema: GenMessage<Deposit, DepositJson> = /*@__PURE__*/
  messageDesc(file_cosmos_gov_v1_gov, 1);

/**
 * Proposal defines the core field members of a governance proposal.
 *
 * @generated from message cosmos.gov.v1.Proposal
 */
export type Proposal = Message<"cosmos.gov.v1.Proposal"> & {
  /**
   * id defines the unique id of the proposal.
   *
   * @generated from field: uint64 id = 1;
   */
  id: bigint;

  /**
   * messages are the arbitrary messages to be executed if the proposal passes.
   *
   * @generated from field: repeated google.protobuf.Any messages = 2;
   */
  messages: Any[];

  /**
   * status defines the proposal status.
   *
   * @generated from field: cosmos.gov.v1.ProposalStatus status = 3;
   */
  status: ProposalStatus;

  /**
   * final_tally_result is the final tally result of the proposal. When
   * querying a proposal via gRPC, this field is not populated until the
   * proposal's voting period has ended.
   *
   * @generated from field: cosmos.gov.v1.TallyResult final_tally_result = 4;
   */
  finalTallyResult?: TallyResult;

  /**
   * submit_time is the time of proposal submission.
   *
   * @generated from field: google.protobuf.Timestamp submit_time = 5;
   */
  submitTime?: Timestamp;

  /**
   * deposit_end_time is the end time for deposition.
   *
   * @generated from field: google.protobuf.Timestamp deposit_end_time = 6;
   */
  depositEndTime?: Timestamp;

  /**
   * total_deposit is the total deposit on the proposal.
   *
   * @generated from field: repeated cosmos.base.v1beta1.Coin total_deposit = 7;
   */
  totalDeposit: Coin[];

  /**
   * voting_start_time is the starting time to vote on a proposal.
   *
   * @generated from field: google.protobuf.Timestamp voting_start_time = 8;
   */
  votingStartTime?: Timestamp;

  /**
   * voting_end_time is the end time of voting on a proposal.
   *
   * @generated from field: google.protobuf.Timestamp voting_end_time = 9;
   */
  votingEndTime?: Timestamp;

  /**
   * metadata is any arbitrary metadata attached to the proposal.
   *
   * @generated from field: string metadata = 10;
   */
  metadata: string;

  /**
   * title is the title of the proposal
   *
   * Since: cosmos-sdk 0.47
   *
   * @generated from field: string title = 11;
   */
  title: string;

  /**
   * summary is a short summary of the proposal
   *
   * Since: cosmos-sdk 0.47
   *
   * @generated from field: string summary = 12;
   */
  summary: string;

  /**
   * proposer is the address of the proposal sumbitter
   *
   * Since: cosmos-sdk 0.47
   *
   * @generated from field: string proposer = 13;
   */
  proposer: string;

  /**
   * expedited defines if the proposal is expedited
   *
   * NOTE: backported from v50 (https://github.com/cosmos/cosmos-sdk/pull/14720)
   *
   * @generated from field: bool expedited = 14;
   */
  expedited: boolean;
};

/**
 * Proposal defines the core field members of a governance proposal.
 *
 * @generated from message cosmos.gov.v1.Proposal
 */
export type ProposalJson = {
  /**
   * id defines the unique id of the proposal.
   *
   * @generated from field: uint64 id = 1;
   */
  id?: string;

  /**
   * messages are the arbitrary messages to be executed if the proposal passes.
   *
   * @generated from field: repeated google.protobuf.Any messages = 2;
   */
  messages?: AnyJson[];

  /**
   * status defines the proposal status.
   *
   * @generated from field: cosmos.gov.v1.ProposalStatus status = 3;
   */
  status?: ProposalStatusJson;

  /**
   * final_tally_result is the final tally result of the proposal. When
   * querying a proposal via gRPC, this field is not populated until the
   * proposal's voting period has ended.
   *
   * @generated from field: cosmos.gov.v1.TallyResult final_tally_result = 4;
   */
  finalTallyResult?: TallyResultJson;

  /**
   * submit_time is the time of proposal submission.
   *
   * @generated from field: google.protobuf.Timestamp submit_time = 5;
   */
  submitTime?: TimestampJson;

  /**
   * deposit_end_time is the end time for deposition.
   *
   * @generated from field: google.protobuf.Timestamp deposit_end_time = 6;
   */
  depositEndTime?: TimestampJson;

  /**
   * total_deposit is the total deposit on the proposal.
   *
   * @generated from field: repeated cosmos.base.v1beta1.Coin total_deposit = 7;
   */
  totalDeposit?: CoinJson[];

  /**
   * voting_start_time is the starting time to vote on a proposal.
   *
   * @generated from field: google.protobuf.Timestamp voting_start_time = 8;
   */
  votingStartTime?: TimestampJson;

  /**
   * voting_end_time is the end time of voting on a proposal.
   *
   * @generated from field: google.protobuf.Timestamp voting_end_time = 9;
   */
  votingEndTime?: TimestampJson;

  /**
   * metadata is any arbitrary metadata attached to the proposal.
   *
   * @generated from field: string metadata = 10;
   */
  metadata?: string;

  /**
   * title is the title of the proposal
   *
   * Since: cosmos-sdk 0.47
   *
   * @generated from field: string title = 11;
   */
  title?: string;

  /**
   * summary is a short summary of the proposal
   *
   * Since: cosmos-sdk 0.47
   *
   * @generated from field: string summary = 12;
   */
  summary?: string;

  /**
   * proposer is the address of the proposal sumbitter
   *
   * Since: cosmos-sdk 0.47
   *
   * @generated from field: string proposer = 13;
   */
  proposer?: string;

  /**
   * expedited defines if the proposal is expedited
   *
   * NOTE: backported from v50 (https://github.com/cosmos/cosmos-sdk/pull/14720)
   *
   * @generated from field: bool expedited = 14;
   */
  expedited?: boolean;
};

/**
 * Describes the message cosmos.gov.v1.Proposal.
 * Use `create(ProposalSchema)` to create a new message.
 */
export const ProposalSchema: GenMessage<Proposal, ProposalJson> = /*@__PURE__*/
  messageDesc(file_cosmos_gov_v1_gov, 2);

/**
 * TallyResult defines a standard tally for a governance proposal.
 *
 * @generated from message cosmos.gov.v1.TallyResult
 */
export type TallyResult = Message<"cosmos.gov.v1.TallyResult"> & {
  /**
   * yes_count is the number of yes votes on a proposal.
   *
   * @generated from field: string yes_count = 1;
   */
  yesCount: string;

  /**
   * abstain_count is the number of abstain votes on a proposal.
   *
   * @generated from field: string abstain_count = 2;
   */
  abstainCount: string;

  /**
   * no_count is the number of no votes on a proposal.
   *
   * @generated from field: string no_count = 3;
   */
  noCount: string;

  /**
   * no_with_veto_count is the number of no with veto votes on a proposal.
   *
   * @generated from field: string no_with_veto_count = 4;
   */
  noWithVetoCount: string;
};

/**
 * TallyResult defines a standard tally for a governance proposal.
 *
 * @generated from message cosmos.gov.v1.TallyResult
 */
export type TallyResultJson = {
  /**
   * yes_count is the number of yes votes on a proposal.
   *
   * @generated from field: string yes_count = 1;
   */
  yesCount?: string;

  /**
   * abstain_count is the number of abstain votes on a proposal.
   *
   * @generated from field: string abstain_count = 2;
   */
  abstainCount?: string;

  /**
   * no_count is the number of no votes on a proposal.
   *
   * @generated from field: string no_count = 3;
   */
  noCount?: string;

  /**
   * no_with_veto_count is the number of no with veto votes on a proposal.
   *
   * @generated from field: string no_with_veto_count = 4;
   */
  noWithVetoCount?: string;
};

/**
 * Describes the message cosmos.gov.v1.TallyResult.
 * Use `create(TallyResultSchema)` to create a new message.
 */
export const TallyResultSchema: GenMessage<TallyResult, TallyResultJson> = /*@__PURE__*/
  messageDesc(file_cosmos_gov_v1_gov, 3);

/**
 * Vote defines a vote on a governance proposal.
 * A Vote consists of a proposal ID, the voter, and the vote option.
 *
 * @generated from message cosmos.gov.v1.Vote
 */
export type Vote = Message<"cosmos.gov.v1.Vote"> & {
  /**
   * proposal_id defines the unique id of the proposal.
   *
   * @generated from field: uint64 proposal_id = 1;
   */
  proposalId: bigint;

  /**
   * voter is the voter address of the proposal.
   *
   * @generated from field: string voter = 2;
   */
  voter: string;

  /**
   * options is the weighted vote options.
   *
   * @generated from field: repeated cosmos.gov.v1.WeightedVoteOption options = 4;
   */
  options: WeightedVoteOption[];

  /**
   * metadata is any  arbitrary metadata to attached to the vote.
   *
   * @generated from field: string metadata = 5;
   */
  metadata: string;
};

/**
 * Vote defines a vote on a governance proposal.
 * A Vote consists of a proposal ID, the voter, and the vote option.
 *
 * @generated from message cosmos.gov.v1.Vote
 */
export type VoteJson = {
  /**
   * proposal_id defines the unique id of the proposal.
   *
   * @generated from field: uint64 proposal_id = 1;
   */
  proposalId?: string;

  /**
   * voter is the voter address of the proposal.
   *
   * @generated from field: string voter = 2;
   */
  voter?: string;

  /**
   * options is the weighted vote options.
   *
   * @generated from field: repeated cosmos.gov.v1.WeightedVoteOption options = 4;
   */
  options?: WeightedVoteOptionJson[];

  /**
   * metadata is any  arbitrary metadata to attached to the vote.
   *
   * @generated from field: string metadata = 5;
   */
  metadata?: string;
};

/**
 * Describes the message cosmos.gov.v1.Vote.
 * Use `create(VoteSchema)` to create a new message.
 */
export const VoteSchema: GenMessage<Vote, VoteJson> = /*@__PURE__*/
  messageDesc(file_cosmos_gov_v1_gov, 4);

/**
 * DepositParams defines the params for deposits on governance proposals.
 *
 * @generated from message cosmos.gov.v1.DepositParams
 */
export type DepositParams = Message<"cosmos.gov.v1.DepositParams"> & {
  /**
   * Minimum deposit for a proposal to enter voting period.
   *
   * @generated from field: repeated cosmos.base.v1beta1.Coin min_deposit = 1;
   */
  minDeposit: Coin[];

  /**
   * Maximum period for Atom holders to deposit on a proposal. Initial value: 2
   * months.
   *
   * @generated from field: google.protobuf.Duration max_deposit_period = 2;
   */
  maxDepositPeriod?: Duration;
};

/**
 * DepositParams defines the params for deposits on governance proposals.
 *
 * @generated from message cosmos.gov.v1.DepositParams
 */
export type DepositParamsJson = {
  /**
   * Minimum deposit for a proposal to enter voting period.
   *
   * @generated from field: repeated cosmos.base.v1beta1.Coin min_deposit = 1;
   */
  minDeposit?: CoinJson[];

  /**
   * Maximum period for Atom holders to deposit on a proposal. Initial value: 2
   * months.
   *
   * @generated from field: google.protobuf.Duration max_deposit_period = 2;
   */
  maxDepositPeriod?: DurationJson;
};

/**
 * Describes the message cosmos.gov.v1.DepositParams.
 * Use `create(DepositParamsSchema)` to create a new message.
 */
export const DepositParamsSchema: GenMessage<DepositParams, DepositParamsJson> = /*@__PURE__*/
  messageDesc(file_cosmos_gov_v1_gov, 5);

/**
 * VotingParams defines the params for voting on governance proposals.
 *
 * @generated from message cosmos.gov.v1.VotingParams
 */
export type VotingParams = Message<"cosmos.gov.v1.VotingParams"> & {
  /**
   * Duration of the voting period.
   *
   * @generated from field: google.protobuf.Duration voting_period = 1;
   */
  votingPeriod?: Duration;
};

/**
 * VotingParams defines the params for voting on governance proposals.
 *
 * @generated from message cosmos.gov.v1.VotingParams
 */
export type VotingParamsJson = {
  /**
   * Duration of the voting period.
   *
   * @generated from field: google.protobuf.Duration voting_period = 1;
   */
  votingPeriod?: DurationJson;
};

/**
 * Describes the message cosmos.gov.v1.VotingParams.
 * Use `create(VotingParamsSchema)` to create a new message.
 */
export const VotingParamsSchema: GenMessage<VotingParams, VotingParamsJson> = /*@__PURE__*/
  messageDesc(file_cosmos_gov_v1_gov, 6);

/**
 * TallyParams defines the params for tallying votes on governance proposals.
 *
 * @generated from message cosmos.gov.v1.TallyParams
 */
export type TallyParams = Message<"cosmos.gov.v1.TallyParams"> & {
  /**
   * Minimum percentage of total stake needed to vote for a result to be
   * considered valid.
   *
   * @generated from field: string quorum = 1;
   */
  quorum: string;

  /**
   * Minimum proportion of Yes votes for proposal to pass. Default value: 0.5.
   *
   * @generated from field: string threshold = 2;
   */
  threshold: string;

  /**
   * Minimum value of Veto votes to Total votes ratio for proposal to be
   * vetoed. Default value: 1/3.
   *
   * @generated from field: string veto_threshold = 3;
   */
  vetoThreshold: string;
};

/**
 * TallyParams defines the params for tallying votes on governance proposals.
 *
 * @generated from message cosmos.gov.v1.TallyParams
 */
export type TallyParamsJson = {
  /**
   * Minimum percentage of total stake needed to vote for a result to be
   * considered valid.
   *
   * @generated from field: string quorum = 1;
   */
  quorum?: string;

  /**
   * Minimum proportion of Yes votes for proposal to pass. Default value: 0.5.
   *
   * @generated from field: string threshold = 2;
   */
  threshold?: string;

  /**
   * Minimum value of Veto votes to Total votes ratio for proposal to be
   * vetoed. Default value: 1/3.
   *
   * @generated from field: string veto_threshold = 3;
   */
  vetoThreshold?: string;
};

/**
 * Describes the message cosmos.gov.v1.TallyParams.
 * Use `create(TallyParamsSchema)` to create a new message.
 */
export const TallyParamsSchema: GenMessage<TallyParams, TallyParamsJson> = /*@__PURE__*/
  messageDesc(file_cosmos_gov_v1_gov, 7);

/**
 * Params defines the parameters for the x/gov module.
 *
 * Since: cosmos-sdk 0.47
 *
 * @generated from message cosmos.gov.v1.Params
 */
export type Params = Message<"cosmos.gov.v1.Params"> & {
  /**
   * Minimum deposit for a proposal to enter voting period.
   *
   * @generated from field: repeated cosmos.base.v1beta1.Coin min_deposit = 1;
   */
  minDeposit: Coin[];

  /**
   * Maximum period for Atom holders to deposit on a proposal. Initial value: 2
   * months.
   *
   * @generated from field: google.protobuf.Duration max_deposit_period = 2;
   */
  maxDepositPeriod?: Duration;

  /**
   * Duration of the voting period.
   *
   * @generated from field: google.protobuf.Duration voting_period = 3;
   */
  votingPeriod?: Duration;

  /**
   *  Minimum percentage of total stake needed to vote for a result to be
   *  considered valid.
   *
   * @generated from field: string quorum = 4;
   */
  quorum: string;

  /**
   *  Minimum proportion of Yes votes for proposal to pass. Default value: 0.5.
   *
   * @generated from field: string threshold = 5;
   */
  threshold: string;

  /**
   *  Minimum value of Veto votes to Total votes ratio for proposal to be
   *  vetoed. Default value: 1/3.
   *
   * @generated from field: string veto_threshold = 6;
   */
  vetoThreshold: string;

  /**
   *  The ratio representing the proportion of the deposit value that must be paid at proposal submission.
   *
   * @generated from field: string min_initial_deposit_ratio = 7;
   */
  minInitialDepositRatio: string;

  /**
   * Duration of the voting period of an expedited proposal.
   *
   * NOTE: backported from v50 (https://github.com/cosmos/cosmos-sdk/pull/14720)
   *
   * @generated from field: google.protobuf.Duration expedited_voting_period = 10;
   */
  expeditedVotingPeriod?: Duration;

  /**
   * Minimum proportion of Yes votes for proposal to pass. Default value: 0.67.
   *
   * NOTE: backported from v50 (https://github.com/cosmos/cosmos-sdk/pull/14720)
   *
   * @generated from field: string expedited_threshold = 11;
   */
  expeditedThreshold: string;

  /**
   *  Minimum expedited deposit for a proposal to enter voting period.
   *
   * NOTE: backported from v50 (https://github.com/cosmos/cosmos-sdk/pull/14720)
   *
   * @generated from field: repeated cosmos.base.v1beta1.Coin expedited_min_deposit = 12;
   */
  expeditedMinDeposit: Coin[];

  /**
   * burn deposits if a proposal does not meet quorum
   *
   * @generated from field: bool burn_vote_quorum = 13;
   */
  burnVoteQuorum: boolean;

  /**
   * burn deposits if the proposal does not enter voting period
   *
   * @generated from field: bool burn_proposal_deposit_prevote = 14;
   */
  burnProposalDepositPrevote: boolean;

  /**
   * burn deposits if quorum with vote type no_veto is met
   *
   * @generated from field: bool burn_vote_veto = 15;
   */
  burnVoteVeto: boolean;

  /**
   * The ratio representing the proportion of the deposit value minimum that must be met when making a deposit.
   * Default value: 0.01. Meaning that for a chain with a min_deposit of 100stake, a deposit of 1stake would be
   * required.
   *
   * Since: cosmos-sdk 0.50
   * NOTE: backported from v50 (https://github.com/cosmos/cosmos-sdk/pull/18146)
   *
   * @generated from field: string min_deposit_ratio = 16;
   */
  minDepositRatio: string;
};

/**
 * Params defines the parameters for the x/gov module.
 *
 * Since: cosmos-sdk 0.47
 *
 * @generated from message cosmos.gov.v1.Params
 */
export type ParamsJson = {
  /**
   * Minimum deposit for a proposal to enter voting period.
   *
   * @generated from field: repeated cosmos.base.v1beta1.Coin min_deposit = 1;
   */
  minDeposit?: CoinJson[];

  /**
   * Maximum period for Atom holders to deposit on a proposal. Initial value: 2
   * months.
   *
   * @generated from field: google.protobuf.Duration max_deposit_period = 2;
   */
  maxDepositPeriod?: DurationJson;

  /**
   * Duration of the voting period.
   *
   * @generated from field: google.protobuf.Duration voting_period = 3;
   */
  votingPeriod?: DurationJson;

  /**
   *  Minimum percentage of total stake needed to vote for a result to be
   *  considered valid.
   *
   * @generated from field: string quorum = 4;
   */
  quorum?: string;

  /**
   *  Minimum proportion of Yes votes for proposal to pass. Default value: 0.5.
   *
   * @generated from field: string threshold = 5;
   */
  threshold?: string;

  /**
   *  Minimum value of Veto votes to Total votes ratio for proposal to be
   *  vetoed. Default value: 1/3.
   *
   * @generated from field: string veto_threshold = 6;
   */
  vetoThreshold?: string;

  /**
   *  The ratio representing the proportion of the deposit value that must be paid at proposal submission.
   *
   * @generated from field: string min_initial_deposit_ratio = 7;
   */
  minInitialDepositRatio?: string;

  /**
   * Duration of the voting period of an expedited proposal.
   *
   * NOTE: backported from v50 (https://github.com/cosmos/cosmos-sdk/pull/14720)
   *
   * @generated from field: google.protobuf.Duration expedited_voting_period = 10;
   */
  expeditedVotingPeriod?: DurationJson;

  /**
   * Minimum proportion of Yes votes for proposal to pass. Default value: 0.67.
   *
   * NOTE: backported from v50 (https://github.com/cosmos/cosmos-sdk/pull/14720)
   *
   * @generated from field: string expedited_threshold = 11;
   */
  expeditedThreshold?: string;

  /**
   *  Minimum expedited deposit for a proposal to enter voting period.
   *
   * NOTE: backported from v50 (https://github.com/cosmos/cosmos-sdk/pull/14720)
   *
   * @generated from field: repeated cosmos.base.v1beta1.Coin expedited_min_deposit = 12;
   */
  expeditedMinDeposit?: CoinJson[];

  /**
   * burn deposits if a proposal does not meet quorum
   *
   * @generated from field: bool burn_vote_quorum = 13;
   */
  burnVoteQuorum?: boolean;

  /**
   * burn deposits if the proposal does not enter voting period
   *
   * @generated from field: bool burn_proposal_deposit_prevote = 14;
   */
  burnProposalDepositPrevote?: boolean;

  /**
   * burn deposits if quorum with vote type no_veto is met
   *
   * @generated from field: bool burn_vote_veto = 15;
   */
  burnVoteVeto?: boolean;

  /**
   * The ratio representing the proportion of the deposit value minimum that must be met when making a deposit.
   * Default value: 0.01. Meaning that for a chain with a min_deposit of 100stake, a deposit of 1stake would be
   * required.
   *
   * Since: cosmos-sdk 0.50
   * NOTE: backported from v50 (https://github.com/cosmos/cosmos-sdk/pull/18146)
   *
   * @generated from field: string min_deposit_ratio = 16;
   */
  minDepositRatio?: string;
};

/**
 * Describes the message cosmos.gov.v1.Params.
 * Use `create(ParamsSchema)` to create a new message.
 */
export const ParamsSchema: GenMessage<Params, ParamsJson> = /*@__PURE__*/
  messageDesc(file_cosmos_gov_v1_gov, 8);

/**
 * VoteOption enumerates the valid vote options for a given governance proposal.
 *
 * @generated from enum cosmos.gov.v1.VoteOption
 */
export enum VoteOption {
  /**
   * VOTE_OPTION_UNSPECIFIED defines a no-op vote option.
   *
   * @generated from enum value: VOTE_OPTION_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * VOTE_OPTION_YES defines a yes vote option.
   *
   * @generated from enum value: VOTE_OPTION_YES = 1;
   */
  YES = 1,

  /**
   * VOTE_OPTION_ABSTAIN defines an abstain vote option.
   *
   * @generated from enum value: VOTE_OPTION_ABSTAIN = 2;
   */
  ABSTAIN = 2,

  /**
   * VOTE_OPTION_NO defines a no vote option.
   *
   * @generated from enum value: VOTE_OPTION_NO = 3;
   */
  NO = 3,

  /**
   * VOTE_OPTION_NO_WITH_VETO defines a no with veto vote option.
   *
   * @generated from enum value: VOTE_OPTION_NO_WITH_VETO = 4;
   */
  NO_WITH_VETO = 4,
}

/**
 * VoteOption enumerates the valid vote options for a given governance proposal.
 *
 * @generated from enum cosmos.gov.v1.VoteOption
 */
export type VoteOptionJson = "VOTE_OPTION_UNSPECIFIED" | "VOTE_OPTION_YES" | "VOTE_OPTION_ABSTAIN" | "VOTE_OPTION_NO" | "VOTE_OPTION_NO_WITH_VETO";

/**
 * Describes the enum cosmos.gov.v1.VoteOption.
 */
export const VoteOptionSchema: GenEnum<VoteOption, VoteOptionJson> = /*@__PURE__*/
  enumDesc(file_cosmos_gov_v1_gov, 0);

/**
 * ProposalStatus enumerates the valid statuses of a proposal.
 *
 * @generated from enum cosmos.gov.v1.ProposalStatus
 */
export enum ProposalStatus {
  /**
   * PROPOSAL_STATUS_UNSPECIFIED defines the default proposal status.
   *
   * @generated from enum value: PROPOSAL_STATUS_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * PROPOSAL_STATUS_DEPOSIT_PERIOD defines a proposal status during the deposit
   * period.
   *
   * @generated from enum value: PROPOSAL_STATUS_DEPOSIT_PERIOD = 1;
   */
  DEPOSIT_PERIOD = 1,

  /**
   * PROPOSAL_STATUS_VOTING_PERIOD defines a proposal status during the voting
   * period.
   *
   * @generated from enum value: PROPOSAL_STATUS_VOTING_PERIOD = 2;
   */
  VOTING_PERIOD = 2,

  /**
   * PROPOSAL_STATUS_PASSED defines a proposal status of a proposal that has
   * passed.
   *
   * @generated from enum value: PROPOSAL_STATUS_PASSED = 3;
   */
  PASSED = 3,

  /**
   * PROPOSAL_STATUS_REJECTED defines a proposal status of a proposal that has
   * been rejected.
   *
   * @generated from enum value: PROPOSAL_STATUS_REJECTED = 4;
   */
  REJECTED = 4,

  /**
   * PROPOSAL_STATUS_FAILED defines a proposal status of a proposal that has
   * failed.
   *
   * @generated from enum value: PROPOSAL_STATUS_FAILED = 5;
   */
  FAILED = 5,
}

/**
 * ProposalStatus enumerates the valid statuses of a proposal.
 *
 * @generated from enum cosmos.gov.v1.ProposalStatus
 */
export type ProposalStatusJson = "PROPOSAL_STATUS_UNSPECIFIED" | "PROPOSAL_STATUS_DEPOSIT_PERIOD" | "PROPOSAL_STATUS_VOTING_PERIOD" | "PROPOSAL_STATUS_PASSED" | "PROPOSAL_STATUS_REJECTED" | "PROPOSAL_STATUS_FAILED";

/**
 * Describes the enum cosmos.gov.v1.ProposalStatus.
 */
export const ProposalStatusSchema: GenEnum<ProposalStatus, ProposalStatusJson> = /*@__PURE__*/
  enumDesc(file_cosmos_gov_v1_gov, 1);

