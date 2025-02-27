// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true"
// @generated from file cosmos/capability/v1beta1/capability.proto (package cosmos.capability.v1beta1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import { file_gogoproto_gogo } from "../../../gogoproto/gogo_pb";
import { file_amino_amino } from "../../../amino/amino_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file cosmos/capability/v1beta1/capability.proto.
 */
export const file_cosmos_capability_v1beta1_capability: GenFile = /*@__PURE__*/
  fileDesc("Cipjb3Ntb3MvY2FwYWJpbGl0eS92MWJldGExL2NhcGFiaWxpdHkucHJvdG8SGWNvc21vcy5jYXBhYmlsaXR5LnYxYmV0YTEiIQoKQ2FwYWJpbGl0eRINCgVpbmRleBgBIAEoBDoEmKAfACIvCgVPd25lchIOCgZtb2R1bGUYASABKAkSDAoEbmFtZRgCIAEoCToIiKAfAJigHwAiTwoQQ2FwYWJpbGl0eU93bmVycxI7CgZvd25lcnMYASADKAsyIC5jb3Ntb3MuY2FwYWJpbGl0eS52MWJldGExLk93bmVyQgnI3h8AqOewKgFCMVovZ2l0aHViLmNvbS9jb3Ntb3MvY29zbW9zLXNkay94L2NhcGFiaWxpdHkvdHlwZXNiBnByb3RvMw", [file_gogoproto_gogo, file_amino_amino]);

/**
 * Capability defines an implementation of an object capability. The index
 * provided to a Capability must be globally unique.
 *
 * @generated from message cosmos.capability.v1beta1.Capability
 */
export type Capability = Message<"cosmos.capability.v1beta1.Capability"> & {
  /**
   * @generated from field: uint64 index = 1;
   */
  index: bigint;
};

/**
 * Capability defines an implementation of an object capability. The index
 * provided to a Capability must be globally unique.
 *
 * @generated from message cosmos.capability.v1beta1.Capability
 */
export type CapabilityJson = {
  /**
   * @generated from field: uint64 index = 1;
   */
  index?: string;
};

/**
 * Describes the message cosmos.capability.v1beta1.Capability.
 * Use `create(CapabilitySchema)` to create a new message.
 */
export const CapabilitySchema: GenMessage<Capability, CapabilityJson> = /*@__PURE__*/
  messageDesc(file_cosmos_capability_v1beta1_capability, 0);

/**
 * Owner defines a single capability owner. An owner is defined by the name of
 * capability and the module name.
 *
 * @generated from message cosmos.capability.v1beta1.Owner
 */
export type Owner = Message<"cosmos.capability.v1beta1.Owner"> & {
  /**
   * @generated from field: string module = 1;
   */
  module: string;

  /**
   * @generated from field: string name = 2;
   */
  name: string;
};

/**
 * Owner defines a single capability owner. An owner is defined by the name of
 * capability and the module name.
 *
 * @generated from message cosmos.capability.v1beta1.Owner
 */
export type OwnerJson = {
  /**
   * @generated from field: string module = 1;
   */
  module?: string;

  /**
   * @generated from field: string name = 2;
   */
  name?: string;
};

/**
 * Describes the message cosmos.capability.v1beta1.Owner.
 * Use `create(OwnerSchema)` to create a new message.
 */
export const OwnerSchema: GenMessage<Owner, OwnerJson> = /*@__PURE__*/
  messageDesc(file_cosmos_capability_v1beta1_capability, 1);

/**
 * CapabilityOwners defines a set of owners of a single Capability. The set of
 * owners must be unique.
 *
 * @generated from message cosmos.capability.v1beta1.CapabilityOwners
 */
export type CapabilityOwners = Message<"cosmos.capability.v1beta1.CapabilityOwners"> & {
  /**
   * @generated from field: repeated cosmos.capability.v1beta1.Owner owners = 1;
   */
  owners: Owner[];
};

/**
 * CapabilityOwners defines a set of owners of a single Capability. The set of
 * owners must be unique.
 *
 * @generated from message cosmos.capability.v1beta1.CapabilityOwners
 */
export type CapabilityOwnersJson = {
  /**
   * @generated from field: repeated cosmos.capability.v1beta1.Owner owners = 1;
   */
  owners?: OwnerJson[];
};

/**
 * Describes the message cosmos.capability.v1beta1.CapabilityOwners.
 * Use `create(CapabilityOwnersSchema)` to create a new message.
 */
export const CapabilityOwnersSchema: GenMessage<CapabilityOwners, CapabilityOwnersJson> = /*@__PURE__*/
  messageDesc(file_cosmos_capability_v1beta1_capability, 2);

