// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true"
// @generated from file cosmos/gov/module/v1/module.proto (package cosmos.gov.module.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import { file_cosmos_app_v1alpha1_module } from "../../../app/v1alpha1/module_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file cosmos/gov/module/v1/module.proto.
 */
export const file_cosmos_gov_module_v1_module: GenFile = /*@__PURE__*/
  fileDesc("CiFjb3Ntb3MvZ292L21vZHVsZS92MS9tb2R1bGUucHJvdG8SFGNvc21vcy5nb3YubW9kdWxlLnYxImEKBk1vZHVsZRIYChBtYXhfbWV0YWRhdGFfbGVuGAEgASgEEhEKCWF1dGhvcml0eRgCIAEoCToqusCW2gEkCiJnaXRodWIuY29tL2Nvc21vcy9jb3Ntb3Mtc2RrL3gvZ292YgZwcm90bzM", [file_cosmos_app_v1alpha1_module]);

/**
 * Module is the config object of the gov module.
 *
 * @generated from message cosmos.gov.module.v1.Module
 */
export type Module = Message<"cosmos.gov.module.v1.Module"> & {
  /**
   * max_metadata_len defines the maximum proposal metadata length. 
   * Defaults to 255 if not explicitly set.
   *
   * @generated from field: uint64 max_metadata_len = 1;
   */
  maxMetadataLen: bigint;

  /**
   * authority defines the custom module authority. If not set, defaults to the governance module.
   *
   * @generated from field: string authority = 2;
   */
  authority: string;
};

/**
 * Module is the config object of the gov module.
 *
 * @generated from message cosmos.gov.module.v1.Module
 */
export type ModuleJson = {
  /**
   * max_metadata_len defines the maximum proposal metadata length. 
   * Defaults to 255 if not explicitly set.
   *
   * @generated from field: uint64 max_metadata_len = 1;
   */
  maxMetadataLen?: string;

  /**
   * authority defines the custom module authority. If not set, defaults to the governance module.
   *
   * @generated from field: string authority = 2;
   */
  authority?: string;
};

/**
 * Describes the message cosmos.gov.module.v1.Module.
 * Use `create(ModuleSchema)` to create a new message.
 */
export const ModuleSchema: GenMessage<Module, ModuleJson> = /*@__PURE__*/
  messageDesc(file_cosmos_gov_module_v1_module, 0);

