// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true"
// @generated from file akash/staking/v1beta3/paramsmsg.proto (package akash.staking.v1beta3, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import { file_gogoproto_gogo } from "../../../gogoproto/gogo_pb";
import { file_amino_amino } from "../../../amino/amino_pb";
import { file_cosmos_msg_v1_msg } from "../../../cosmos/msg/v1/msg_pb";
import { file_cosmos_proto_cosmos } from "../../../cosmos_proto/cosmos_pb";
import type { Params, ParamsJson } from "./params_pb";
import { file_akash_staking_v1beta3_params } from "./params_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file akash/staking/v1beta3/paramsmsg.proto.
 */
export const file_akash_staking_v1beta3_paramsmsg: GenFile = /*@__PURE__*/
  fileDesc("CiVha2FzaC9zdGFraW5nL3YxYmV0YTMvcGFyYW1zbXNnLnByb3RvEhVha2FzaC5zdGFraW5nLnYxYmV0YTMirAEKD01zZ1VwZGF0ZVBhcmFtcxIrCglhdXRob3JpdHkYASABKAlCGNK0LRRjb3Ntb3MuQWRkcmVzc1N0cmluZxI4CgZwYXJhbXMYAiABKAsyHS5ha2FzaC5zdGFraW5nLnYxYmV0YTMuUGFyYW1zQgnI3h8AqOewKgE6MoLnsCoJYXV0aG9yaXR5iuewKh9ha2FzaC94L3N0YWtpbmcvTXNnVXBkYXRlUGFyYW1zIhkKF01zZ1VwZGF0ZVBhcmFtc1Jlc3BvbnNlQiVaI3BrZy5ha3QuZGV2L2dvL25vZGUvc3Rha2luZy92MWJldGEzYgZwcm90bzM", [file_gogoproto_gogo, file_amino_amino, file_cosmos_msg_v1_msg, file_cosmos_proto_cosmos, file_akash_staking_v1beta3_params]);

/**
 * MsgUpdateParams is the Msg/UpdateParams request type.
 *
 * Since: akash v1.0.0
 *
 * @generated from message akash.staking.v1beta3.MsgUpdateParams
 */
export type MsgUpdateParams = Message<"akash.staking.v1beta3.MsgUpdateParams"> & {
  /**
   * authority is the address of the governance account.
   *
   * @generated from field: string authority = 1;
   */
  authority: string;

  /**
   * params defines the x/deployment parameters to update.
   *
   * NOTE: All parameters must be supplied.
   *
   * @generated from field: akash.staking.v1beta3.Params params = 2;
   */
  params?: Params;
};

/**
 * MsgUpdateParams is the Msg/UpdateParams request type.
 *
 * Since: akash v1.0.0
 *
 * @generated from message akash.staking.v1beta3.MsgUpdateParams
 */
export type MsgUpdateParamsJson = {
  /**
   * authority is the address of the governance account.
   *
   * @generated from field: string authority = 1;
   */
  authority?: string;

  /**
   * params defines the x/deployment parameters to update.
   *
   * NOTE: All parameters must be supplied.
   *
   * @generated from field: akash.staking.v1beta3.Params params = 2;
   */
  params?: ParamsJson;
};

/**
 * Describes the message akash.staking.v1beta3.MsgUpdateParams.
 * Use `create(MsgUpdateParamsSchema)` to create a new message.
 */
export const MsgUpdateParamsSchema: GenMessage<MsgUpdateParams, MsgUpdateParamsJson> = /*@__PURE__*/
  messageDesc(file_akash_staking_v1beta3_paramsmsg, 0);

/**
 * MsgUpdateParamsResponse defines the response structure for executing a
 * MsgUpdateParams message.
 *
 * Since: akash v1.0.0
 *
 * @generated from message akash.staking.v1beta3.MsgUpdateParamsResponse
 */
export type MsgUpdateParamsResponse = Message<"akash.staking.v1beta3.MsgUpdateParamsResponse"> & {
};

/**
 * MsgUpdateParamsResponse defines the response structure for executing a
 * MsgUpdateParams message.
 *
 * Since: akash v1.0.0
 *
 * @generated from message akash.staking.v1beta3.MsgUpdateParamsResponse
 */
export type MsgUpdateParamsResponseJson = {
};

/**
 * Describes the message akash.staking.v1beta3.MsgUpdateParamsResponse.
 * Use `create(MsgUpdateParamsResponseSchema)` to create a new message.
 */
export const MsgUpdateParamsResponseSchema: GenMessage<MsgUpdateParamsResponse, MsgUpdateParamsResponseJson> = /*@__PURE__*/
  messageDesc(file_akash_staking_v1beta3_paramsmsg, 1);

