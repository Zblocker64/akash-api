/* eslint-disable */
import { messageTypeRegistry } from "../../../typeRegistry";
import Long from "long";
import { Coin } from "../../../cosmos/base/v1beta1/coin";
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "akash.deployment.v1beta1";

/**
 * DepositDeploymentAuthorization allows the grantee to deposit up to spend_limit coins from
 * the granter's account for a deployment.
 */
export interface DepositDeploymentAuthorization {
    $type: "akash.deployment.v1beta1.DepositDeploymentAuthorization";
    /**
     * SpendLimit is the amount the grantee is authorized to spend from the granter's account for
     * the purpose of deployment.
     */
    spendLimit: Coin | undefined;
}

function createBaseDepositDeploymentAuthorization(): DepositDeploymentAuthorization {
    return {
        $type: "akash.deployment.v1beta1.DepositDeploymentAuthorization",
        spendLimit: undefined,
    };
}

export const DepositDeploymentAuthorization = {
    $type: "akash.deployment.v1beta1.DepositDeploymentAuthorization" as const,

    encode(
        message: DepositDeploymentAuthorization,
        writer: _m0.Writer = _m0.Writer.create(),
    ): _m0.Writer {
        if (message.spendLimit !== undefined) {
            Coin.encode(message.spendLimit, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },

    decode(
        input: _m0.Reader | Uint8Array,
        length?: number,
    ): DepositDeploymentAuthorization {
        const reader =
            input instanceof _m0.Reader ? input : new _m0.Reader(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseDepositDeploymentAuthorization();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.spendLimit = Coin.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },

    fromJSON(object: any): DepositDeploymentAuthorization {
        return {
            $type: DepositDeploymentAuthorization.$type,
            spendLimit: isSet(object.spendLimit)
                ? Coin.fromJSON(object.spendLimit)
                : undefined,
        };
    },

    toJSON(message: DepositDeploymentAuthorization): unknown {
        const obj: any = {};
        message.spendLimit !== undefined &&
            (obj.spendLimit = message.spendLimit
                ? Coin.toJSON(message.spendLimit)
                : undefined);
        return obj;
    },

    fromPartial<
        I extends Exact<DeepPartial<DepositDeploymentAuthorization>, I>,
    >(object: I): DepositDeploymentAuthorization {
        const message = createBaseDepositDeploymentAuthorization();
        message.spendLimit =
            object.spendLimit !== undefined && object.spendLimit !== null
                ? Coin.fromPartial(object.spendLimit)
                : undefined;
        return message;
    },
};

messageTypeRegistry.set(
    DepositDeploymentAuthorization.$type,
    DepositDeploymentAuthorization,
);

type Builtin =
    | Date
    | Function
    | Uint8Array
    | string
    | number
    | boolean
    | undefined;

export type DeepPartial<T> = T extends Builtin
    ? T
    : T extends Long
      ? string | number | Long
      : T extends Array<infer U>
        ? Array<DeepPartial<U>>
        : T extends ReadonlyArray<infer U>
          ? ReadonlyArray<DeepPartial<U>>
          : T extends {}
            ? { [K in Exclude<keyof T, "$type">]?: DeepPartial<T[K]> }
            : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin
    ? P
    : P & { [K in keyof P]: Exact<P[K], I[K]> } & Record<
              Exclude<keyof I, KeysOfUnion<P> | "$type">,
              never
          >;

if (_m0.util.Long !== Long) {
    _m0.util.Long = Long as any;
    _m0.configure();
}

function isSet(value: any): boolean {
    return value !== null && value !== undefined;
}
