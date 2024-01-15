declare const _default: import("vue").DefineComponent<{
    animation: {
        type: StringConstructor;
        default: undefined;
    };
    animationSpeed: {
        type: StringConstructor;
        default: undefined;
    };
    fill: {
        type: StringConstructor;
        default: string;
    };
    size: {
        type: (StringConstructor | NumberConstructor)[];
        default: number;
    };
    stroke: {
        type: StringConstructor;
        default: string;
    };
    strokeLinecap: {
        type: StringConstructor;
        default: string;
    };
    strokeLinejoin: {
        type: StringConstructor;
        default: string;
    };
    strokeWidth: {
        type: (StringConstructor | NumberConstructor)[];
        default: number;
    };
    tag: {
        type: StringConstructor;
        default: string;
    };
    type: {
        type: StringConstructor;
        default: string;
        validator(type: string): true;
    };
}, unknown, unknown, {
    isRemSize(): boolean;
}, {}, import("vue").ComponentOptionsMixin, import("vue").ComponentOptionsMixin, Record<string, any>, string, import("vue").VNodeProps & import("vue").AllowedComponentProps & import("vue").ComponentCustomProps, Readonly<import("vue").ExtractPropTypes<{
    animation: {
        type: StringConstructor;
        default: undefined;
    };
    animationSpeed: {
        type: StringConstructor;
        default: undefined;
    };
    fill: {
        type: StringConstructor;
        default: string;
    };
    size: {
        type: (StringConstructor | NumberConstructor)[];
        default: number;
    };
    stroke: {
        type: StringConstructor;
        default: string;
    };
    strokeLinecap: {
        type: StringConstructor;
        default: string;
    };
    strokeLinejoin: {
        type: StringConstructor;
        default: string;
    };
    strokeWidth: {
        type: (StringConstructor | NumberConstructor)[];
        default: number;
    };
    tag: {
        type: StringConstructor;
        default: string;
    };
    type: {
        type: StringConstructor;
        default: string;
        validator(type: string): true;
    };
}>>, {
    fill: string;
    animation: string;
    animationSpeed: string;
    type: string;
    size: string | number;
    stroke: string;
    strokeLinecap: string;
    strokeLinejoin: string;
    strokeWidth: string | number;
    tag: string;
}>;
export default _default;
