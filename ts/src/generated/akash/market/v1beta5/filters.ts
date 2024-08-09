/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import Long = require("long");

export const protobufPackage = "akash.market.v1beta5";

/** BidFilters defines flags for bid list filter */
export interface BidFilters {
    owner: string;
    dseq: number;
    gseq: number;
    oseq: number;
    provider: string;
    state: string;
}

/** OrderFilters defines flags for order list filter */
export interface OrderFilters {
    owner: string;
    dseq: number;
    gseq: number;
    oseq: number;
    state: string;
}

function createBaseBidFilters(): BidFilters {
    return { owner: "", dseq: 0, gseq: 0, oseq: 0, provider: "", state: "" };
}

export const BidFilters = {
    encode(
        message: BidFilters,
        writer: _m0.Writer = _m0.Writer.create(),
    ): _m0.Writer {
        if (message.owner !== "") {
            writer.uint32(10).string(message.owner);
        }
        if (message.dseq !== 0) {
            writer.uint32(16).uint64(message.dseq);
        }
        if (message.gseq !== 0) {
            writer.uint32(24).uint32(message.gseq);
        }
        if (message.oseq !== 0) {
            writer.uint32(32).uint32(message.oseq);
        }
        if (message.provider !== "") {
            writer.uint32(42).string(message.provider);
        }
        if (message.state !== "") {
            writer.uint32(50).string(message.state);
        }
        return writer;
    },

    decode(input: _m0.Reader | Uint8Array, length?: number): BidFilters {
        const reader =
            input instanceof _m0.Reader ? input : _m0.Reader.create(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseBidFilters();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    if (tag !== 10) {
                        break;
                    }

                    message.owner = reader.string();
                    continue;
                case 2:
                    if (tag !== 16) {
                        break;
                    }

                    message.dseq = longToNumber(reader.uint64() as Long);
                    continue;
                case 3:
                    if (tag !== 24) {
                        break;
                    }

                    message.gseq = reader.uint32();
                    continue;
                case 4:
                    if (tag !== 32) {
                        break;
                    }

                    message.oseq = reader.uint32();
                    continue;
                case 5:
                    if (tag !== 42) {
                        break;
                    }

                    message.provider = reader.string();
                    continue;
                case 6:
                    if (tag !== 50) {
                        break;
                    }

                    message.state = reader.string();
                    continue;
            }
            if ((tag & 7) === 4 || tag === 0) {
                break;
            }
            reader.skipType(tag & 7);
        }
        return message;
    },

    fromJSON(object: any): BidFilters {
        return {
            owner: isSet(object.owner) ? globalThis.String(object.owner) : "",
            dseq: isSet(object.dseq) ? globalThis.Number(object.dseq) : 0,
            gseq: isSet(object.gseq) ? globalThis.Number(object.gseq) : 0,
            oseq: isSet(object.oseq) ? globalThis.Number(object.oseq) : 0,
            provider: isSet(object.provider)
                ? globalThis.String(object.provider)
                : "",
            state: isSet(object.state) ? globalThis.String(object.state) : "",
        };
    },

    toJSON(message: BidFilters): unknown {
        const obj: any = {};
        if (message.owner !== "") {
            obj.owner = message.owner;
        }
        if (message.dseq !== 0) {
            obj.dseq = Math.round(message.dseq);
        }
        if (message.gseq !== 0) {
            obj.gseq = Math.round(message.gseq);
        }
        if (message.oseq !== 0) {
            obj.oseq = Math.round(message.oseq);
        }
        if (message.provider !== "") {
            obj.provider = message.provider;
        }
        if (message.state !== "") {
            obj.state = message.state;
        }
        return obj;
    },

    create<I extends Exact<DeepPartial<BidFilters>, I>>(base?: I): BidFilters {
        return BidFilters.fromPartial(base ?? ({} as any));
    },
    fromPartial<I extends Exact<DeepPartial<BidFilters>, I>>(
        object: I,
    ): BidFilters {
        const message = createBaseBidFilters();
        message.owner = object.owner ?? "";
        message.dseq = object.dseq ?? 0;
        message.gseq = object.gseq ?? 0;
        message.oseq = object.oseq ?? 0;
        message.provider = object.provider ?? "";
        message.state = object.state ?? "";
        return message;
    },
};

function createBaseOrderFilters(): OrderFilters {
    return { owner: "", dseq: 0, gseq: 0, oseq: 0, state: "" };
}

export const OrderFilters = {
    encode(
        message: OrderFilters,
        writer: _m0.Writer = _m0.Writer.create(),
    ): _m0.Writer {
        if (message.owner !== "") {
            writer.uint32(10).string(message.owner);
        }
        if (message.dseq !== 0) {
            writer.uint32(16).uint64(message.dseq);
        }
        if (message.gseq !== 0) {
            writer.uint32(24).uint32(message.gseq);
        }
        if (message.oseq !== 0) {
            writer.uint32(32).uint32(message.oseq);
        }
        if (message.state !== "") {
            writer.uint32(42).string(message.state);
        }
        return writer;
    },

    decode(input: _m0.Reader | Uint8Array, length?: number): OrderFilters {
        const reader =
            input instanceof _m0.Reader ? input : _m0.Reader.create(input);
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = createBaseOrderFilters();
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    if (tag !== 10) {
                        break;
                    }

                    message.owner = reader.string();
                    continue;
                case 2:
                    if (tag !== 16) {
                        break;
                    }

                    message.dseq = longToNumber(reader.uint64() as Long);
                    continue;
                case 3:
                    if (tag !== 24) {
                        break;
                    }

                    message.gseq = reader.uint32();
                    continue;
                case 4:
                    if (tag !== 32) {
                        break;
                    }

                    message.oseq = reader.uint32();
                    continue;
                case 5:
                    if (tag !== 42) {
                        break;
                    }

                    message.state = reader.string();
                    continue;
            }
            if ((tag & 7) === 4 || tag === 0) {
                break;
            }
            reader.skipType(tag & 7);
        }
        return message;
    },

    fromJSON(object: any): OrderFilters {
        return {
            owner: isSet(object.owner) ? globalThis.String(object.owner) : "",
            dseq: isSet(object.dseq) ? globalThis.Number(object.dseq) : 0,
            gseq: isSet(object.gseq) ? globalThis.Number(object.gseq) : 0,
            oseq: isSet(object.oseq) ? globalThis.Number(object.oseq) : 0,
            state: isSet(object.state) ? globalThis.String(object.state) : "",
        };
    },

    toJSON(message: OrderFilters): unknown {
        const obj: any = {};
        if (message.owner !== "") {
            obj.owner = message.owner;
        }
        if (message.dseq !== 0) {
            obj.dseq = Math.round(message.dseq);
        }
        if (message.gseq !== 0) {
            obj.gseq = Math.round(message.gseq);
        }
        if (message.oseq !== 0) {
            obj.oseq = Math.round(message.oseq);
        }
        if (message.state !== "") {
            obj.state = message.state;
        }
        return obj;
    },

    create<I extends Exact<DeepPartial<OrderFilters>, I>>(
        base?: I,
    ): OrderFilters {
        return OrderFilters.fromPartial(base ?? ({} as any));
    },
    fromPartial<I extends Exact<DeepPartial<OrderFilters>, I>>(
        object: I,
    ): OrderFilters {
        const message = createBaseOrderFilters();
        message.owner = object.owner ?? "";
        message.dseq = object.dseq ?? 0;
        message.gseq = object.gseq ?? 0;
        message.oseq = object.oseq ?? 0;
        message.state = object.state ?? "";
        return message;
    },
};

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
    : T extends globalThis.Array<infer U>
      ? globalThis.Array<DeepPartial<U>>
      : T extends ReadonlyArray<infer U>
        ? ReadonlyArray<DeepPartial<U>>
        : T extends {}
          ? { [K in keyof T]?: DeepPartial<T[K]> }
          : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin
    ? P
    : P & { [K in keyof P]: Exact<P[K], I[K]> } & {
          [K in Exclude<keyof I, KeysOfUnion<P>>]: never;
      };

function longToNumber(long: Long): number {
    if (long.gt(globalThis.Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error(
            "Value is larger than Number.MAX_SAFE_INTEGER",
        );
    }
    return long.toNumber();
}

if (_m0.util.Long !== Long) {
    _m0.util.Long = Long as any;
    _m0.configure();
}

function isSet(value: any): boolean {
    return value !== null && value !== undefined;
}
