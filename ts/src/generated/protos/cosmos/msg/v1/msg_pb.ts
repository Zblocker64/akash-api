// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true"
// @generated from file cosmos/msg/v1/msg.proto (package cosmos.msg.v1, syntax proto3)
/* eslint-disable */

import type { GenExtension, GenFile } from "@bufbuild/protobuf/codegenv1";
import { extDesc, fileDesc } from "@bufbuild/protobuf/codegenv1";
import type { MessageOptions, ServiceOptions } from "@bufbuild/protobuf/wkt";
import { file_google_protobuf_descriptor } from "@bufbuild/protobuf/wkt";

/**
 * Describes the file cosmos/msg/v1/msg.proto.
 */
export const file_cosmos_msg_v1_msg: GenFile = /*@__PURE__*/
  fileDesc("Chdjb3Ntb3MvbXNnL3YxL21zZy5wcm90bxINY29zbW9zLm1zZy52MTo8CgdzZXJ2aWNlEh8uZ29vZ2xlLnByb3RvYnVmLlNlcnZpY2VPcHRpb25zGPCMpgUgASgIUgdzZXJ2aWNlOjoKBnNpZ25lchIfLmdvb2dsZS5wcm90b2J1Zi5NZXNzYWdlT3B0aW9ucxjwjKYFIAMoCVIGc2lnbmVyQi9aLWdpdGh1Yi5jb20vY29zbW9zL2Nvc21vcy1zZGsvdHlwZXMvbXNnc2VydmljZWIGcHJvdG8z", [file_google_protobuf_descriptor]);

/**
 * service indicates that the service is a Msg service and that requests
 * must be transported via blockchain transactions rather than gRPC.
 * Tooling can use this annotation to distinguish between Msg services and
 * other types of services via reflection.
 *
 * @generated from extension: bool service = 11110000;
 */
export const service: GenExtension<ServiceOptions, boolean> = /*@__PURE__*/
  extDesc(file_cosmos_msg_v1_msg, 0);

/**
 * signer must be used in cosmos messages in order
 * to signal to external clients which fields in a
 * given cosmos message must be filled with signer
 * information (address).
 * The field must be the protobuf name of the message
 * field extended with this MessageOption.
 * The field must either be of string kind, or of message
 * kind in case the signer information is contained within
 * a message inside the cosmos message.
 *
 * @generated from extension: repeated string signer = 11110000;
 */
export const signer: GenExtension<MessageOptions, string[]> = /*@__PURE__*/
  extDesc(file_cosmos_msg_v1_msg, 1);

