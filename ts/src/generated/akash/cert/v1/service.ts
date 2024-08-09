/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import {
    MsgCreateCertificate,
    MsgCreateCertificateResponse,
    MsgRevokeCertificate,
    MsgRevokeCertificateResponse,
} from "./msg";

export const protobufPackage = "akash.cert.v1";

/** Msg defines the provider Msg service */
export interface Msg {
    /** CreateCertificate defines a method to create new certificate given proper inputs. */
    CreateCertificate(
        request: MsgCreateCertificate,
    ): Promise<MsgCreateCertificateResponse>;
    /** RevokeCertificate defines a method to revoke the certificate */
    RevokeCertificate(
        request: MsgRevokeCertificate,
    ): Promise<MsgRevokeCertificateResponse>;
}

export const MsgServiceName = "akash.cert.v1.Msg";
export class MsgClientImpl implements Msg {
    private readonly rpc: Rpc;
    private readonly service: string;
    constructor(rpc: Rpc, opts?: { service?: string }) {
        this.service = opts?.service || MsgServiceName;
        this.rpc = rpc;
        this.CreateCertificate = this.CreateCertificate.bind(this);
        this.RevokeCertificate = this.RevokeCertificate.bind(this);
    }
    CreateCertificate(
        request: MsgCreateCertificate,
    ): Promise<MsgCreateCertificateResponse> {
        const data = MsgCreateCertificate.encode(request).finish();
        const promise = this.rpc.request(
            this.service,
            "CreateCertificate",
            data,
        );
        return promise.then((data) =>
            MsgCreateCertificateResponse.decode(_m0.Reader.create(data)),
        );
    }

    RevokeCertificate(
        request: MsgRevokeCertificate,
    ): Promise<MsgRevokeCertificateResponse> {
        const data = MsgRevokeCertificate.encode(request).finish();
        const promise = this.rpc.request(
            this.service,
            "RevokeCertificate",
            data,
        );
        return promise.then((data) =>
            MsgRevokeCertificateResponse.decode(_m0.Reader.create(data)),
        );
    }
}

interface Rpc {
    request(
        service: string,
        method: string,
        data: Uint8Array,
    ): Promise<Uint8Array>;
}
