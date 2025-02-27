// @generated by protoc-gen-es v2.2.3 with parameter "target=ts,json_types=true"
// @generated from file akash/manifest/v2beta3/serviceexpose.proto (package akash.manifest.v2beta3, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import { file_gogoproto_gogo } from "../../../gogoproto/gogo_pb";
import type { ServiceExposeHTTPOptions, ServiceExposeHTTPOptionsJson } from "./httpoptions_pb";
import { file_akash_manifest_v2beta3_httpoptions } from "./httpoptions_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file akash/manifest/v2beta3/serviceexpose.proto.
 */
export const file_akash_manifest_v2beta3_serviceexpose: GenFile = /*@__PURE__*/
  fileDesc("Cipha2FzaC9tYW5pZmVzdC92MmJldGEzL3NlcnZpY2VleHBvc2UucHJvdG8SFmFrYXNoLm1hbmlmZXN0LnYyYmV0YTMiyQQKDVNlcnZpY2VFeHBvc2USJQoEcG9ydBgBIAEoDUIX6t4fBHBvcnTy3h8LeWFtbDoicG9ydCISPgoNZXh0ZXJuYWxfcG9ydBgCIAEoDUIn6t4fDGV4dGVybmFsUG9ydPLeHxN5YW1sOiJleHRlcm5hbFBvcnQiEjsKBXByb3RvGAMgASgJQizq3h8FcHJvdG/y3h8MeWFtbDoicHJvdG8i+t4fD1NlcnZpY2VQcm90b2NvbBIuCgdzZXJ2aWNlGAQgASgJQh3q3h8Hc2VydmljZfLeHw55YW1sOiJzZXJ2aWNlIhIrCgZnbG9iYWwYBSABKAhCG+reHwZnbG9iYWzy3h8NeWFtbDoiZ2xvYmFsIhIoCgVob3N0cxgGIAMoCUIZ6t4fBWhvc3Rz8t4fDHlhbWw6Imhvc3RzIhKAAQoMaHR0cF9vcHRpb25zGAcgASgLMjAuYWthc2gubWFuaWZlc3QudjJiZXRhMy5TZXJ2aWNlRXhwb3NlSFRUUE9wdGlvbnNCOMjeHwDi3h8LSFRUUE9wdGlvbnPq3h8LaHR0cE9wdGlvbnPy3h8SeWFtbDoiaHR0cE9wdGlvbnMiEiUKAmlwGAggASgJQhni3h8CSVDq3h8CaXDy3h8JeWFtbDoiaXAiEl0KGGVuZHBvaW50X3NlcXVlbmNlX251bWJlchgJIAEoDUI76t4fFmVuZHBvaW50U2VxdWVuY2VOdW1iZXLy3h8deWFtbDoiZW5kcG9pbnRTZXF1ZW5jZU51bWJlciI6BIigHwBCKVofcGtnLmFrdC5kZXYvZ28vbWFuaWZlc3QvdjJiZXRhM9jhHgCA4h4BYgZwcm90bzM", [file_gogoproto_gogo, file_akash_manifest_v2beta3_httpoptions]);

/**
 * ServiceExpose stores exposed ports and hosts details
 *
 * @generated from message akash.manifest.v2beta3.ServiceExpose
 */
export type ServiceExpose = Message<"akash.manifest.v2beta3.ServiceExpose"> & {
  /**
   * port on the container
   *
   * @generated from field: uint32 port = 1;
   */
  port: number;

  /**
   * port on the service definition
   *
   * @generated from field: uint32 external_port = 2;
   */
  externalPort: number;

  /**
   * @generated from field: string proto = 3;
   */
  proto: string;

  /**
   * @generated from field: string service = 4;
   */
  service: string;

  /**
   * @generated from field: bool global = 5;
   */
  global: boolean;

  /**
   * @generated from field: repeated string hosts = 6;
   */
  hosts: string[];

  /**
   * @generated from field: akash.manifest.v2beta3.ServiceExposeHTTPOptions http_options = 7;
   */
  httpOptions?: ServiceExposeHTTPOptions;

  /**
   * The name of the IP address associated with this, if any
   *
   * @generated from field: string ip = 8;
   */
  ip: string;

  /**
   * The sequence number of the associated endpoint in the on-chain data
   *
   * @generated from field: uint32 endpoint_sequence_number = 9;
   */
  endpointSequenceNumber: number;
};

/**
 * ServiceExpose stores exposed ports and hosts details
 *
 * @generated from message akash.manifest.v2beta3.ServiceExpose
 */
export type ServiceExposeJson = {
  /**
   * port on the container
   *
   * @generated from field: uint32 port = 1;
   */
  port?: number;

  /**
   * port on the service definition
   *
   * @generated from field: uint32 external_port = 2;
   */
  externalPort?: number;

  /**
   * @generated from field: string proto = 3;
   */
  proto?: string;

  /**
   * @generated from field: string service = 4;
   */
  service?: string;

  /**
   * @generated from field: bool global = 5;
   */
  global?: boolean;

  /**
   * @generated from field: repeated string hosts = 6;
   */
  hosts?: string[];

  /**
   * @generated from field: akash.manifest.v2beta3.ServiceExposeHTTPOptions http_options = 7;
   */
  httpOptions?: ServiceExposeHTTPOptionsJson;

  /**
   * The name of the IP address associated with this, if any
   *
   * @generated from field: string ip = 8;
   */
  ip?: string;

  /**
   * The sequence number of the associated endpoint in the on-chain data
   *
   * @generated from field: uint32 endpoint_sequence_number = 9;
   */
  endpointSequenceNumber?: number;
};

/**
 * Describes the message akash.manifest.v2beta3.ServiceExpose.
 * Use `create(ServiceExposeSchema)` to create a new message.
 */
export const ServiceExposeSchema: GenMessage<ServiceExpose, ServiceExposeJson> = /*@__PURE__*/
  messageDesc(file_akash_manifest_v2beta3_serviceexpose, 0);

