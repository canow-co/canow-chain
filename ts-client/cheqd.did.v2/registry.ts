import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreateDidDoc } from "./types/cheqd/did/v2/tx";
import { MsgUpdateDidDoc } from "./types/cheqd/did/v2/tx";
import { MsgDeactivateDidDoc } from "./types/cheqd/did/v2/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/cheqd.did.v2.MsgCreateDidDoc", MsgCreateDidDoc],
    ["/cheqd.did.v2.MsgUpdateDidDoc", MsgUpdateDidDoc],
    ["/cheqd.did.v2.MsgDeactivateDidDoc", MsgDeactivateDidDoc],
    
];

export { msgTypes }