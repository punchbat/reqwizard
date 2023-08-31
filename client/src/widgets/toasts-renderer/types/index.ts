import { MessageType, MessagePositions, MessageProp } from "../store";

interface DefaultParamsProp {
    message: MessageProp;
    options: { type: MessageType; duration: number; position: MessagePositions };
}

export type { DefaultParamsProp };
