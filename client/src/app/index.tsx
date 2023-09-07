import { Routing } from "../pages/index";
import { StoresLayout, ThemeLayout } from "./layouts";

import "./index.scss";

function App() {
    return (
        <StoresLayout>
            {/* <AppLayout> */}
            <ThemeLayout>
                <Routing />
            </ThemeLayout>
            {/* </AppLayout> */}
        </StoresLayout>
    );
}

export default App;
