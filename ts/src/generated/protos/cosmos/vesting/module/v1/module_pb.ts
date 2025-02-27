// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true"
// @generated from file cosmos/vesting/module/v1/module.proto (package cosmos.vesting.module.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import { file_cosmos_app_v1alpha1_module } from "../../../app/v1alpha1/module_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file cosmos/vesting/module/v1/module.proto.
 */
export const file_cosmos_vesting_module_v1_module: GenFile = /*@__PURE__*/
  fileDesc("CiVjb3Ntb3MvdmVzdGluZy9tb2R1bGUvdjEvbW9kdWxlLnByb3RvEhhjb3Ntb3MudmVzdGluZy5tb2R1bGUudjEiPQoGTW9kdWxlOjO6wJbaAS0KK2dpdGh1Yi5jb20vY29zbW9zL2Nvc21vcy1zZGsveC9hdXRoL3Zlc3RpbmdiBnByb3RvMw", [file_cosmos_app_v1alpha1_module]);

/**
 * Module is the config object of the vesting module.
 *
 * @generated from message cosmos.vesting.module.v1.Module
 */
export type Module = Message<"cosmos.vesting.module.v1.Module"> & {
};

/**
 * Module is the config object of the vesting module.
 *
 * @generated from message cosmos.vesting.module.v1.Module
 */
export type ModuleJson = {
};

/**
 * Describes the message cosmos.vesting.module.v1.Module.
 * Use `create(ModuleSchema)` to create a new message.
 */
export const ModuleSchema: GenMessage<Module, ModuleJson> = /*@__PURE__*/
  messageDesc(file_cosmos_vesting_module_v1_module, 0);

