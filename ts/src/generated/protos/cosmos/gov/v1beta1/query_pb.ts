// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true"
// @generated from file cosmos/gov/v1beta1/query.proto (package cosmos.gov.v1beta1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import type { PageRequest, PageRequestJson, PageResponse, PageResponseJson } from "../../base/query/v1beta1/pagination_pb";
import { file_cosmos_base_query_v1beta1_pagination } from "../../base/query/v1beta1/pagination_pb";
import { file_gogoproto_gogo } from "../../../gogoproto/gogo_pb";
import { file_google_api_annotations } from "../../../google/api/annotations_pb";
import type { Deposit, DepositJson, DepositParams, DepositParamsJson, Proposal, ProposalJson, ProposalStatus, ProposalStatusJson, TallyParams, TallyParamsJson, TallyResult, TallyResultJson, Vote, VoteJson, VotingParams, VotingParamsJson } from "./gov_pb";
import { file_cosmos_gov_v1beta1_gov } from "./gov_pb";
import { file_cosmos_proto_cosmos } from "../../../cosmos_proto/cosmos_pb";
import { file_amino_amino } from "../../../amino/amino_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file cosmos/gov/v1beta1/query.proto.
 */
export const file_cosmos_gov_v1beta1_query: GenFile = /*@__PURE__*/
  fileDesc("Ch5jb3Ntb3MvZ292L3YxYmV0YTEvcXVlcnkucHJvdG8SEmNvc21vcy5nb3YudjFiZXRhMSIrChRRdWVyeVByb3Bvc2FsUmVxdWVzdBITCgtwcm9wb3NhbF9pZBgBIAEoBCJSChVRdWVyeVByb3Bvc2FsUmVzcG9uc2USOQoIcHJvcG9zYWwYASABKAsyHC5jb3Ntb3MuZ292LnYxYmV0YTEuUHJvcG9zYWxCCcjeHwCo57AqASLwAQoVUXVlcnlQcm9wb3NhbHNSZXF1ZXN0EjsKD3Byb3Bvc2FsX3N0YXR1cxgBIAEoDjIiLmNvc21vcy5nb3YudjFiZXRhMS5Qcm9wb3NhbFN0YXR1cxInCgV2b3RlchgCIAEoCUIY0rQtFGNvc21vcy5BZGRyZXNzU3RyaW5nEisKCWRlcG9zaXRvchgDIAEoCUIY0rQtFGNvc21vcy5BZGRyZXNzU3RyaW5nEjoKCnBhZ2luYXRpb24YBCABKAsyJi5jb3Ntb3MuYmFzZS5xdWVyeS52MWJldGExLlBhZ2VSZXF1ZXN0OgiIoB8A6KAfACKRAQoWUXVlcnlQcm9wb3NhbHNSZXNwb25zZRI6Cglwcm9wb3NhbHMYASADKAsyHC5jb3Ntb3MuZ292LnYxYmV0YTEuUHJvcG9zYWxCCcjeHwCo57AqARI7CgpwYWdpbmF0aW9uGAIgASgLMicuY29zbW9zLmJhc2UucXVlcnkudjFiZXRhMS5QYWdlUmVzcG9uc2UiWgoQUXVlcnlWb3RlUmVxdWVzdBITCgtwcm9wb3NhbF9pZBgBIAEoBBInCgV2b3RlchgCIAEoCUIY0rQtFGNvc21vcy5BZGRyZXNzU3RyaW5nOgiIoB8A6KAfACJGChFRdWVyeVZvdGVSZXNwb25zZRIxCgR2b3RlGAEgASgLMhguY29zbW9zLmdvdi52MWJldGExLlZvdGVCCcjeHwCo57AqASJkChFRdWVyeVZvdGVzUmVxdWVzdBITCgtwcm9wb3NhbF9pZBgBIAEoBBI6CgpwYWdpbmF0aW9uGAIgASgLMiYuY29zbW9zLmJhc2UucXVlcnkudjFiZXRhMS5QYWdlUmVxdWVzdCKFAQoSUXVlcnlWb3Rlc1Jlc3BvbnNlEjIKBXZvdGVzGAEgAygLMhguY29zbW9zLmdvdi52MWJldGExLlZvdGVCCcjeHwCo57AqARI7CgpwYWdpbmF0aW9uGAIgASgLMicuY29zbW9zLmJhc2UucXVlcnkudjFiZXRhMS5QYWdlUmVzcG9uc2UiKQoSUXVlcnlQYXJhbXNSZXF1ZXN0EhMKC3BhcmFtc190eXBlGAEgASgJIuEBChNRdWVyeVBhcmFtc1Jlc3BvbnNlEkIKDXZvdGluZ19wYXJhbXMYASABKAsyIC5jb3Ntb3MuZ292LnYxYmV0YTEuVm90aW5nUGFyYW1zQgnI3h8AqOewKgESRAoOZGVwb3NpdF9wYXJhbXMYAiABKAsyIS5jb3Ntb3MuZ292LnYxYmV0YTEuRGVwb3NpdFBhcmFtc0IJyN4fAKjnsCoBEkAKDHRhbGx5X3BhcmFtcxgDIAEoCzIfLmNvc21vcy5nb3YudjFiZXRhMS5UYWxseVBhcmFtc0IJyN4fAKjnsCoBImEKE1F1ZXJ5RGVwb3NpdFJlcXVlc3QSEwoLcHJvcG9zYWxfaWQYASABKAQSKwoJZGVwb3NpdG9yGAIgASgJQhjStC0UY29zbW9zLkFkZHJlc3NTdHJpbmc6CIigHwDooB8AIk8KFFF1ZXJ5RGVwb3NpdFJlc3BvbnNlEjcKB2RlcG9zaXQYASABKAsyGy5jb3Ntb3MuZ292LnYxYmV0YTEuRGVwb3NpdEIJyN4fAKjnsCoBImcKFFF1ZXJ5RGVwb3NpdHNSZXF1ZXN0EhMKC3Byb3Bvc2FsX2lkGAEgASgEEjoKCnBhZ2luYXRpb24YAiABKAsyJi5jb3Ntb3MuYmFzZS5xdWVyeS52MWJldGExLlBhZ2VSZXF1ZXN0Io4BChVRdWVyeURlcG9zaXRzUmVzcG9uc2USOAoIZGVwb3NpdHMYASADKAsyGy5jb3Ntb3MuZ292LnYxYmV0YTEuRGVwb3NpdEIJyN4fAKjnsCoBEjsKCnBhZ2luYXRpb24YAiABKAsyJy5jb3Ntb3MuYmFzZS5xdWVyeS52MWJldGExLlBhZ2VSZXNwb25zZSIuChdRdWVyeVRhbGx5UmVzdWx0UmVxdWVzdBITCgtwcm9wb3NhbF9pZBgBIAEoBCJVChhRdWVyeVRhbGx5UmVzdWx0UmVzcG9uc2USOQoFdGFsbHkYASABKAsyHy5jb3Ntb3MuZ292LnYxYmV0YTEuVGFsbHlSZXN1bHRCCcjeHwCo57AqATLUCQoFUXVlcnkSlAEKCFByb3Bvc2FsEiguY29zbW9zLmdvdi52MWJldGExLlF1ZXJ5UHJvcG9zYWxSZXF1ZXN0GikuY29zbW9zLmdvdi52MWJldGExLlF1ZXJ5UHJvcG9zYWxSZXNwb25zZSIzgtPkkwItEisvY29zbW9zL2dvdi92MWJldGExL3Byb3Bvc2Fscy97cHJvcG9zYWxfaWR9EokBCglQcm9wb3NhbHMSKS5jb3Ntb3MuZ292LnYxYmV0YTEuUXVlcnlQcm9wb3NhbHNSZXF1ZXN0GiouY29zbW9zLmdvdi52MWJldGExLlF1ZXJ5UHJvcG9zYWxzUmVzcG9uc2UiJYLT5JMCHxIdL2Nvc21vcy9nb3YvdjFiZXRhMS9wcm9wb3NhbHMSlgEKBFZvdGUSJC5jb3Ntb3MuZ292LnYxYmV0YTEuUXVlcnlWb3RlUmVxdWVzdBolLmNvc21vcy5nb3YudjFiZXRhMS5RdWVyeVZvdGVSZXNwb25zZSJBgtPkkwI7EjkvY29zbW9zL2dvdi92MWJldGExL3Byb3Bvc2Fscy97cHJvcG9zYWxfaWR9L3ZvdGVzL3t2b3Rlcn0SkQEKBVZvdGVzEiUuY29zbW9zLmdvdi52MWJldGExLlF1ZXJ5Vm90ZXNSZXF1ZXN0GiYuY29zbW9zLmdvdi52MWJldGExLlF1ZXJ5Vm90ZXNSZXNwb25zZSI5gtPkkwIzEjEvY29zbW9zL2dvdi92MWJldGExL3Byb3Bvc2Fscy97cHJvcG9zYWxfaWR9L3ZvdGVzEosBCgZQYXJhbXMSJi5jb3Ntb3MuZ292LnYxYmV0YTEuUXVlcnlQYXJhbXNSZXF1ZXN0GicuY29zbW9zLmdvdi52MWJldGExLlF1ZXJ5UGFyYW1zUmVzcG9uc2UiMILT5JMCKhIoL2Nvc21vcy9nb3YvdjFiZXRhMS9wYXJhbXMve3BhcmFtc190eXBlfRKmAQoHRGVwb3NpdBInLmNvc21vcy5nb3YudjFiZXRhMS5RdWVyeURlcG9zaXRSZXF1ZXN0GiguY29zbW9zLmdvdi52MWJldGExLlF1ZXJ5RGVwb3NpdFJlc3BvbnNlIkiC0+STAkISQC9jb3Ntb3MvZ292L3YxYmV0YTEvcHJvcG9zYWxzL3twcm9wb3NhbF9pZH0vZGVwb3NpdHMve2RlcG9zaXRvcn0SnQEKCERlcG9zaXRzEiguY29zbW9zLmdvdi52MWJldGExLlF1ZXJ5RGVwb3NpdHNSZXF1ZXN0GikuY29zbW9zLmdvdi52MWJldGExLlF1ZXJ5RGVwb3NpdHNSZXNwb25zZSI8gtPkkwI2EjQvY29zbW9zL2dvdi92MWJldGExL3Byb3Bvc2Fscy97cHJvcG9zYWxfaWR9L2RlcG9zaXRzEqMBCgtUYWxseVJlc3VsdBIrLmNvc21vcy5nb3YudjFiZXRhMS5RdWVyeVRhbGx5UmVzdWx0UmVxdWVzdBosLmNvc21vcy5nb3YudjFiZXRhMS5RdWVyeVRhbGx5UmVzdWx0UmVzcG9uc2UiOYLT5JMCMxIxL2Nvc21vcy9nb3YvdjFiZXRhMS9wcm9wb3NhbHMve3Byb3Bvc2FsX2lkfS90YWxseUIyWjBnaXRodWIuY29tL2Nvc21vcy9jb3Ntb3Mtc2RrL3gvZ292L3R5cGVzL3YxYmV0YTFiBnByb3RvMw", [file_cosmos_base_query_v1beta1_pagination, file_gogoproto_gogo, file_google_api_annotations, file_cosmos_gov_v1beta1_gov, file_cosmos_proto_cosmos, file_amino_amino]);

/**
 * QueryProposalRequest is the request type for the Query/Proposal RPC method.
 *
 * @generated from message cosmos.gov.v1beta1.QueryProposalRequest
 */
export type QueryProposalRequest = Message<"cosmos.gov.v1beta1.QueryProposalRequest"> & {
  /**
   * proposal_id defines the unique id of the proposal.
   *
   * @generated from field: uint64 proposal_id = 1;
   */
  proposalId: bigint;
};

/**
 * QueryProposalRequest is the request type for the Query/Proposal RPC method.
 *
 * @generated from message cosmos.gov.v1beta1.QueryProposalRequest
 */
export type QueryProposalRequestJson = {
  /**
   * proposal_id defines the unique id of the proposal.
   *
   * @generated from field: uint64 proposal_id = 1;
   */
  proposalId?: string;
};

/**
 * Describes the message cosmos.gov.v1beta1.QueryProposalRequest.
 * Use `create(QueryProposalRequestSchema)` to create a new message.
 */
export const QueryProposalRequestSchema: GenMessage<QueryProposalRequest, QueryProposalRequestJson> = /*@__PURE__*/
  messageDesc(file_cosmos_gov_v1beta1_query, 0);

/**
 * QueryProposalResponse is the response type for the Query/Proposal RPC method.
 *
 * @generated from message cosmos.gov.v1beta1.QueryProposalResponse
 */
export type QueryProposalResponse = Message<"cosmos.gov.v1beta1.QueryProposalResponse"> & {
  /**
   * @generated from field: cosmos.gov.v1beta1.Proposal proposal = 1;
   */
  proposal?: Proposal;
};

/**
 * QueryProposalResponse is the response type for the Query/Proposal RPC method.
 *
 * @generated from message cosmos.gov.v1beta1.QueryProposalResponse
 */
export type QueryProposalResponseJson = {
  /**
   * @generated from field: cosmos.gov.v1beta1.Proposal proposal = 1;
   */
  proposal?: ProposalJson;
};

/**
 * Describes the message cosmos.gov.v1beta1.QueryProposalResponse.
 * Use `create(QueryProposalResponseSchema)` to create a new message.
 */
export const QueryProposalResponseSchema: GenMessage<QueryProposalResponse, QueryProposalResponseJson> = /*@__PURE__*/
  messageDesc(file_cosmos_gov_v1beta1_query, 1);

/**
 * QueryProposalsRequest is the request type for the Query/Proposals RPC method.
 *
 * @generated from message cosmos.gov.v1beta1.QueryProposalsRequest
 */
export type QueryProposalsRequest = Message<"cosmos.gov.v1beta1.QueryProposalsRequest"> & {
  /**
   * proposal_status defines the status of the proposals.
   *
   * @generated from field: cosmos.gov.v1beta1.ProposalStatus proposal_status = 1;
   */
  proposalStatus: ProposalStatus;

  /**
   * voter defines the voter address for the proposals.
   *
   * @generated from field: string voter = 2;
   */
  voter: string;

  /**
   * depositor defines the deposit addresses from the proposals.
   *
   * @generated from field: string depositor = 3;
   */
  depositor: string;

  /**
   * pagination defines an optional pagination for the request.
   *
   * @generated from field: cosmos.base.query.v1beta1.PageRequest pagination = 4;
   */
  pagination?: PageRequest;
};

/**
 * QueryProposalsRequest is the request type for the Query/Proposals RPC method.
 *
 * @generated from message cosmos.gov.v1beta1.QueryProposalsRequest
 */
export type QueryProposalsRequestJson = {
  /**
   * proposal_status defines the status of the proposals.
   *
   * @generated from field: cosmos.gov.v1beta1.ProposalStatus proposal_status = 1;
   */
  proposalStatus?: ProposalStatusJson;

  /**
   * voter defines the voter address for the proposals.
   *
   * @generated from field: string voter = 2;
   */
  voter?: string;

  /**
   * depositor defines the deposit addresses from the proposals.
   *
   * @generated from field: string depositor = 3;
   */
  depositor?: string;

  /**
   * pagination defines an optional pagination for the request.
   *
   * @generated from field: cosmos.base.query.v1beta1.PageRequest pagination = 4;
   */
  pagination?: PageRequestJson;
};

/**
 * Describes the message cosmos.gov.v1beta1.QueryProposalsRequest.
 * Use `create(QueryProposalsRequestSchema)` to create a new message.
 */
export const QueryProposalsRequestSchema: GenMessage<QueryProposalsRequest, QueryProposalsRequestJson> = /*@__PURE__*/
  messageDesc(file_cosmos_gov_v1beta1_query, 2);

/**
 * QueryProposalsResponse is the response type for the Query/Proposals RPC
 * method.
 *
 * @generated from message cosmos.gov.v1beta1.QueryProposalsResponse
 */
export type QueryProposalsResponse = Message<"cosmos.gov.v1beta1.QueryProposalsResponse"> & {
  /**
   * proposals defines all the requested governance proposals.
   *
   * @generated from field: repeated cosmos.gov.v1beta1.Proposal proposals = 1;
   */
  proposals: Proposal[];

  /**
   * pagination defines the pagination in the response.
   *
   * @generated from field: cosmos.base.query.v1beta1.PageResponse pagination = 2;
   */
  pagination?: PageResponse;
};

/**
 * QueryProposalsResponse is the response type for the Query/Proposals RPC
 * method.
 *
 * @generated from message cosmos.gov.v1beta1.QueryProposalsResponse
 */
export type QueryProposalsResponseJson = {
  /**
   * proposals defines all the requested governance proposals.
   *
   * @generated from field: repeated cosmos.gov.v1beta1.Proposal proposals = 1;
   */
  proposals?: ProposalJson[];

  /**
   * pagination defines the pagination in the response.
   *
   * @generated from field: cosmos.base.query.v1beta1.PageResponse pagination = 2;
   */
  pagination?: PageResponseJson;
};

/**
 * Describes the message cosmos.gov.v1beta1.QueryProposalsResponse.
 * Use `create(QueryProposalsResponseSchema)` to create a new message.
 */
export const QueryProposalsResponseSchema: GenMessage<QueryProposalsResponse, QueryProposalsResponseJson> = /*@__PURE__*/
  messageDesc(file_cosmos_gov_v1beta1_query, 3);

/**
 * QueryVoteRequest is the request type for the Query/Vote RPC method.
 *
 * @generated from message cosmos.gov.v1beta1.QueryVoteRequest
 */
export type QueryVoteRequest = Message<"cosmos.gov.v1beta1.QueryVoteRequest"> & {
  /**
   * proposal_id defines the unique id of the proposal.
   *
   * @generated from field: uint64 proposal_id = 1;
   */
  proposalId: bigint;

  /**
   * voter defines the voter address for the proposals.
   *
   * @generated from field: string voter = 2;
   */
  voter: string;
};

/**
 * QueryVoteRequest is the request type for the Query/Vote RPC method.
 *
 * @generated from message cosmos.gov.v1beta1.QueryVoteRequest
 */
export type QueryVoteRequestJson = {
  /**
   * proposal_id defines the unique id of the proposal.
   *
   * @generated from field: uint64 proposal_id = 1;
   */
  proposalId?: string;

  /**
   * voter defines the voter address for the proposals.
   *
   * @generated from field: string voter = 2;
   */
  voter?: string;
};

/**
 * Describes the message cosmos.gov.v1beta1.QueryVoteRequest.
 * Use `create(QueryVoteRequestSchema)` to create a new message.
 */
export const QueryVoteRequestSchema: GenMessage<QueryVoteRequest, QueryVoteRequestJson> = /*@__PURE__*/
  messageDesc(file_cosmos_gov_v1beta1_query, 4);

/**
 * QueryVoteResponse is the response type for the Query/Vote RPC method.
 *
 * @generated from message cosmos.gov.v1beta1.QueryVoteResponse
 */
export type QueryVoteResponse = Message<"cosmos.gov.v1beta1.QueryVoteResponse"> & {
  /**
   * vote defines the queried vote.
   *
   * @generated from field: cosmos.gov.v1beta1.Vote vote = 1;
   */
  vote?: Vote;
};

/**
 * QueryVoteResponse is the response type for the Query/Vote RPC method.
 *
 * @generated from message cosmos.gov.v1beta1.QueryVoteResponse
 */
export type QueryVoteResponseJson = {
  /**
   * vote defines the queried vote.
   *
   * @generated from field: cosmos.gov.v1beta1.Vote vote = 1;
   */
  vote?: VoteJson;
};

/**
 * Describes the message cosmos.gov.v1beta1.QueryVoteResponse.
 * Use `create(QueryVoteResponseSchema)` to create a new message.
 */
export const QueryVoteResponseSchema: GenMessage<QueryVoteResponse, QueryVoteResponseJson> = /*@__PURE__*/
  messageDesc(file_cosmos_gov_v1beta1_query, 5);

/**
 * QueryVotesRequest is the request type for the Query/Votes RPC method.
 *
 * @generated from message cosmos.gov.v1beta1.QueryVotesRequest
 */
export type QueryVotesRequest = Message<"cosmos.gov.v1beta1.QueryVotesRequest"> & {
  /**
   * proposal_id defines the unique id of the proposal.
   *
   * @generated from field: uint64 proposal_id = 1;
   */
  proposalId: bigint;

  /**
   * pagination defines an optional pagination for the request.
   *
   * @generated from field: cosmos.base.query.v1beta1.PageRequest pagination = 2;
   */
  pagination?: PageRequest;
};

/**
 * QueryVotesRequest is the request type for the Query/Votes RPC method.
 *
 * @generated from message cosmos.gov.v1beta1.QueryVotesRequest
 */
export type QueryVotesRequestJson = {
  /**
   * proposal_id defines the unique id of the proposal.
   *
   * @generated from field: uint64 proposal_id = 1;
   */
  proposalId?: string;

  /**
   * pagination defines an optional pagination for the request.
   *
   * @generated from field: cosmos.base.query.v1beta1.PageRequest pagination = 2;
   */
  pagination?: PageRequestJson;
};

/**
 * Describes the message cosmos.gov.v1beta1.QueryVotesRequest.
 * Use `create(QueryVotesRequestSchema)` to create a new message.
 */
export const QueryVotesRequestSchema: GenMessage<QueryVotesRequest, QueryVotesRequestJson> = /*@__PURE__*/
  messageDesc(file_cosmos_gov_v1beta1_query, 6);

/**
 * QueryVotesResponse is the response type for the Query/Votes RPC method.
 *
 * @generated from message cosmos.gov.v1beta1.QueryVotesResponse
 */
export type QueryVotesResponse = Message<"cosmos.gov.v1beta1.QueryVotesResponse"> & {
  /**
   * votes defines the queried votes.
   *
   * @generated from field: repeated cosmos.gov.v1beta1.Vote votes = 1;
   */
  votes: Vote[];

  /**
   * pagination defines the pagination in the response.
   *
   * @generated from field: cosmos.base.query.v1beta1.PageResponse pagination = 2;
   */
  pagination?: PageResponse;
};

/**
 * QueryVotesResponse is the response type for the Query/Votes RPC method.
 *
 * @generated from message cosmos.gov.v1beta1.QueryVotesResponse
 */
export type QueryVotesResponseJson = {
  /**
   * votes defines the queried votes.
   *
   * @generated from field: repeated cosmos.gov.v1beta1.Vote votes = 1;
   */
  votes?: VoteJson[];

  /**
   * pagination defines the pagination in the response.
   *
   * @generated from field: cosmos.base.query.v1beta1.PageResponse pagination = 2;
   */
  pagination?: PageResponseJson;
};

/**
 * Describes the message cosmos.gov.v1beta1.QueryVotesResponse.
 * Use `create(QueryVotesResponseSchema)` to create a new message.
 */
export const QueryVotesResponseSchema: GenMessage<QueryVotesResponse, QueryVotesResponseJson> = /*@__PURE__*/
  messageDesc(file_cosmos_gov_v1beta1_query, 7);

/**
 * QueryParamsRequest is the request type for the Query/Params RPC method.
 *
 * @generated from message cosmos.gov.v1beta1.QueryParamsRequest
 */
export type QueryParamsRequest = Message<"cosmos.gov.v1beta1.QueryParamsRequest"> & {
  /**
   * params_type defines which parameters to query for, can be one of "voting",
   * "tallying" or "deposit".
   *
   * @generated from field: string params_type = 1;
   */
  paramsType: string;
};

/**
 * QueryParamsRequest is the request type for the Query/Params RPC method.
 *
 * @generated from message cosmos.gov.v1beta1.QueryParamsRequest
 */
export type QueryParamsRequestJson = {
  /**
   * params_type defines which parameters to query for, can be one of "voting",
   * "tallying" or "deposit".
   *
   * @generated from field: string params_type = 1;
   */
  paramsType?: string;
};

/**
 * Describes the message cosmos.gov.v1beta1.QueryParamsRequest.
 * Use `create(QueryParamsRequestSchema)` to create a new message.
 */
export const QueryParamsRequestSchema: GenMessage<QueryParamsRequest, QueryParamsRequestJson> = /*@__PURE__*/
  messageDesc(file_cosmos_gov_v1beta1_query, 8);

/**
 * QueryParamsResponse is the response type for the Query/Params RPC method.
 *
 * @generated from message cosmos.gov.v1beta1.QueryParamsResponse
 */
export type QueryParamsResponse = Message<"cosmos.gov.v1beta1.QueryParamsResponse"> & {
  /**
   * voting_params defines the parameters related to voting.
   *
   * @generated from field: cosmos.gov.v1beta1.VotingParams voting_params = 1;
   */
  votingParams?: VotingParams;

  /**
   * deposit_params defines the parameters related to deposit.
   *
   * @generated from field: cosmos.gov.v1beta1.DepositParams deposit_params = 2;
   */
  depositParams?: DepositParams;

  /**
   * tally_params defines the parameters related to tally.
   *
   * @generated from field: cosmos.gov.v1beta1.TallyParams tally_params = 3;
   */
  tallyParams?: TallyParams;
};

/**
 * QueryParamsResponse is the response type for the Query/Params RPC method.
 *
 * @generated from message cosmos.gov.v1beta1.QueryParamsResponse
 */
export type QueryParamsResponseJson = {
  /**
   * voting_params defines the parameters related to voting.
   *
   * @generated from field: cosmos.gov.v1beta1.VotingParams voting_params = 1;
   */
  votingParams?: VotingParamsJson;

  /**
   * deposit_params defines the parameters related to deposit.
   *
   * @generated from field: cosmos.gov.v1beta1.DepositParams deposit_params = 2;
   */
  depositParams?: DepositParamsJson;

  /**
   * tally_params defines the parameters related to tally.
   *
   * @generated from field: cosmos.gov.v1beta1.TallyParams tally_params = 3;
   */
  tallyParams?: TallyParamsJson;
};

/**
 * Describes the message cosmos.gov.v1beta1.QueryParamsResponse.
 * Use `create(QueryParamsResponseSchema)` to create a new message.
 */
export const QueryParamsResponseSchema: GenMessage<QueryParamsResponse, QueryParamsResponseJson> = /*@__PURE__*/
  messageDesc(file_cosmos_gov_v1beta1_query, 9);

/**
 * QueryDepositRequest is the request type for the Query/Deposit RPC method.
 *
 * @generated from message cosmos.gov.v1beta1.QueryDepositRequest
 */
export type QueryDepositRequest = Message<"cosmos.gov.v1beta1.QueryDepositRequest"> & {
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
};

/**
 * QueryDepositRequest is the request type for the Query/Deposit RPC method.
 *
 * @generated from message cosmos.gov.v1beta1.QueryDepositRequest
 */
export type QueryDepositRequestJson = {
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
};

/**
 * Describes the message cosmos.gov.v1beta1.QueryDepositRequest.
 * Use `create(QueryDepositRequestSchema)` to create a new message.
 */
export const QueryDepositRequestSchema: GenMessage<QueryDepositRequest, QueryDepositRequestJson> = /*@__PURE__*/
  messageDesc(file_cosmos_gov_v1beta1_query, 10);

/**
 * QueryDepositResponse is the response type for the Query/Deposit RPC method.
 *
 * @generated from message cosmos.gov.v1beta1.QueryDepositResponse
 */
export type QueryDepositResponse = Message<"cosmos.gov.v1beta1.QueryDepositResponse"> & {
  /**
   * deposit defines the requested deposit.
   *
   * @generated from field: cosmos.gov.v1beta1.Deposit deposit = 1;
   */
  deposit?: Deposit;
};

/**
 * QueryDepositResponse is the response type for the Query/Deposit RPC method.
 *
 * @generated from message cosmos.gov.v1beta1.QueryDepositResponse
 */
export type QueryDepositResponseJson = {
  /**
   * deposit defines the requested deposit.
   *
   * @generated from field: cosmos.gov.v1beta1.Deposit deposit = 1;
   */
  deposit?: DepositJson;
};

/**
 * Describes the message cosmos.gov.v1beta1.QueryDepositResponse.
 * Use `create(QueryDepositResponseSchema)` to create a new message.
 */
export const QueryDepositResponseSchema: GenMessage<QueryDepositResponse, QueryDepositResponseJson> = /*@__PURE__*/
  messageDesc(file_cosmos_gov_v1beta1_query, 11);

/**
 * QueryDepositsRequest is the request type for the Query/Deposits RPC method.
 *
 * @generated from message cosmos.gov.v1beta1.QueryDepositsRequest
 */
export type QueryDepositsRequest = Message<"cosmos.gov.v1beta1.QueryDepositsRequest"> & {
  /**
   * proposal_id defines the unique id of the proposal.
   *
   * @generated from field: uint64 proposal_id = 1;
   */
  proposalId: bigint;

  /**
   * pagination defines an optional pagination for the request.
   *
   * @generated from field: cosmos.base.query.v1beta1.PageRequest pagination = 2;
   */
  pagination?: PageRequest;
};

/**
 * QueryDepositsRequest is the request type for the Query/Deposits RPC method.
 *
 * @generated from message cosmos.gov.v1beta1.QueryDepositsRequest
 */
export type QueryDepositsRequestJson = {
  /**
   * proposal_id defines the unique id of the proposal.
   *
   * @generated from field: uint64 proposal_id = 1;
   */
  proposalId?: string;

  /**
   * pagination defines an optional pagination for the request.
   *
   * @generated from field: cosmos.base.query.v1beta1.PageRequest pagination = 2;
   */
  pagination?: PageRequestJson;
};

/**
 * Describes the message cosmos.gov.v1beta1.QueryDepositsRequest.
 * Use `create(QueryDepositsRequestSchema)` to create a new message.
 */
export const QueryDepositsRequestSchema: GenMessage<QueryDepositsRequest, QueryDepositsRequestJson> = /*@__PURE__*/
  messageDesc(file_cosmos_gov_v1beta1_query, 12);

/**
 * QueryDepositsResponse is the response type for the Query/Deposits RPC method.
 *
 * @generated from message cosmos.gov.v1beta1.QueryDepositsResponse
 */
export type QueryDepositsResponse = Message<"cosmos.gov.v1beta1.QueryDepositsResponse"> & {
  /**
   * deposits defines the requested deposits.
   *
   * @generated from field: repeated cosmos.gov.v1beta1.Deposit deposits = 1;
   */
  deposits: Deposit[];

  /**
   * pagination defines the pagination in the response.
   *
   * @generated from field: cosmos.base.query.v1beta1.PageResponse pagination = 2;
   */
  pagination?: PageResponse;
};

/**
 * QueryDepositsResponse is the response type for the Query/Deposits RPC method.
 *
 * @generated from message cosmos.gov.v1beta1.QueryDepositsResponse
 */
export type QueryDepositsResponseJson = {
  /**
   * deposits defines the requested deposits.
   *
   * @generated from field: repeated cosmos.gov.v1beta1.Deposit deposits = 1;
   */
  deposits?: DepositJson[];

  /**
   * pagination defines the pagination in the response.
   *
   * @generated from field: cosmos.base.query.v1beta1.PageResponse pagination = 2;
   */
  pagination?: PageResponseJson;
};

/**
 * Describes the message cosmos.gov.v1beta1.QueryDepositsResponse.
 * Use `create(QueryDepositsResponseSchema)` to create a new message.
 */
export const QueryDepositsResponseSchema: GenMessage<QueryDepositsResponse, QueryDepositsResponseJson> = /*@__PURE__*/
  messageDesc(file_cosmos_gov_v1beta1_query, 13);

/**
 * QueryTallyResultRequest is the request type for the Query/Tally RPC method.
 *
 * @generated from message cosmos.gov.v1beta1.QueryTallyResultRequest
 */
export type QueryTallyResultRequest = Message<"cosmos.gov.v1beta1.QueryTallyResultRequest"> & {
  /**
   * proposal_id defines the unique id of the proposal.
   *
   * @generated from field: uint64 proposal_id = 1;
   */
  proposalId: bigint;
};

/**
 * QueryTallyResultRequest is the request type for the Query/Tally RPC method.
 *
 * @generated from message cosmos.gov.v1beta1.QueryTallyResultRequest
 */
export type QueryTallyResultRequestJson = {
  /**
   * proposal_id defines the unique id of the proposal.
   *
   * @generated from field: uint64 proposal_id = 1;
   */
  proposalId?: string;
};

/**
 * Describes the message cosmos.gov.v1beta1.QueryTallyResultRequest.
 * Use `create(QueryTallyResultRequestSchema)` to create a new message.
 */
export const QueryTallyResultRequestSchema: GenMessage<QueryTallyResultRequest, QueryTallyResultRequestJson> = /*@__PURE__*/
  messageDesc(file_cosmos_gov_v1beta1_query, 14);

/**
 * QueryTallyResultResponse is the response type for the Query/Tally RPC method.
 *
 * @generated from message cosmos.gov.v1beta1.QueryTallyResultResponse
 */
export type QueryTallyResultResponse = Message<"cosmos.gov.v1beta1.QueryTallyResultResponse"> & {
  /**
   * tally defines the requested tally.
   *
   * @generated from field: cosmos.gov.v1beta1.TallyResult tally = 1;
   */
  tally?: TallyResult;
};

/**
 * QueryTallyResultResponse is the response type for the Query/Tally RPC method.
 *
 * @generated from message cosmos.gov.v1beta1.QueryTallyResultResponse
 */
export type QueryTallyResultResponseJson = {
  /**
   * tally defines the requested tally.
   *
   * @generated from field: cosmos.gov.v1beta1.TallyResult tally = 1;
   */
  tally?: TallyResultJson;
};

/**
 * Describes the message cosmos.gov.v1beta1.QueryTallyResultResponse.
 * Use `create(QueryTallyResultResponseSchema)` to create a new message.
 */
export const QueryTallyResultResponseSchema: GenMessage<QueryTallyResultResponse, QueryTallyResultResponseJson> = /*@__PURE__*/
  messageDesc(file_cosmos_gov_v1beta1_query, 15);

/**
 * Query defines the gRPC querier service for gov module
 *
 * @generated from service cosmos.gov.v1beta1.Query
 */
export const Query: GenService<{
  /**
   * Proposal queries proposal details based on ProposalID.
   *
   * @generated from rpc cosmos.gov.v1beta1.Query.Proposal
   */
  proposal: {
    methodKind: "unary";
    input: typeof QueryProposalRequestSchema;
    output: typeof QueryProposalResponseSchema;
  },
  /**
   * Proposals queries all proposals based on given status.
   *
   * @generated from rpc cosmos.gov.v1beta1.Query.Proposals
   */
  proposals: {
    methodKind: "unary";
    input: typeof QueryProposalsRequestSchema;
    output: typeof QueryProposalsResponseSchema;
  },
  /**
   * Vote queries voted information based on proposalID, voterAddr.
   *
   * @generated from rpc cosmos.gov.v1beta1.Query.Vote
   */
  vote: {
    methodKind: "unary";
    input: typeof QueryVoteRequestSchema;
    output: typeof QueryVoteResponseSchema;
  },
  /**
   * Votes queries votes of a given proposal.
   *
   * @generated from rpc cosmos.gov.v1beta1.Query.Votes
   */
  votes: {
    methodKind: "unary";
    input: typeof QueryVotesRequestSchema;
    output: typeof QueryVotesResponseSchema;
  },
  /**
   * Params queries all parameters of the gov module.
   *
   * @generated from rpc cosmos.gov.v1beta1.Query.Params
   */
  params: {
    methodKind: "unary";
    input: typeof QueryParamsRequestSchema;
    output: typeof QueryParamsResponseSchema;
  },
  /**
   * Deposit queries single deposit information based proposalID, depositAddr.
   *
   * @generated from rpc cosmos.gov.v1beta1.Query.Deposit
   */
  deposit: {
    methodKind: "unary";
    input: typeof QueryDepositRequestSchema;
    output: typeof QueryDepositResponseSchema;
  },
  /**
   * Deposits queries all deposits of a single proposal.
   *
   * @generated from rpc cosmos.gov.v1beta1.Query.Deposits
   */
  deposits: {
    methodKind: "unary";
    input: typeof QueryDepositsRequestSchema;
    output: typeof QueryDepositsResponseSchema;
  },
  /**
   * TallyResult queries the tally of a proposal vote.
   *
   * @generated from rpc cosmos.gov.v1beta1.Query.TallyResult
   */
  tallyResult: {
    methodKind: "unary";
    input: typeof QueryTallyResultRequestSchema;
    output: typeof QueryTallyResultResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_cosmos_gov_v1beta1_query, 0);

