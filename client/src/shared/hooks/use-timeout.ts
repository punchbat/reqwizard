import React from "react";

type Callback = () => void;

export function useTimeout(callback: Callback, delay: number | null) {
    const savedCallback = React.useRef<Callback | null>();

    React.useEffect(() => {
        savedCallback.current = callback;
    }, [callback]);

    React.useEffect(() => {
        function tick() {
            if (savedCallback.current) {
                savedCallback.current();
            }
        }
        if (delay !== null) {
            setTimeout(tick, delay);
        }
    }, [delay]);
}
