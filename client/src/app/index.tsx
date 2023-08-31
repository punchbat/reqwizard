import { Routing } from "../pages/index";
import { StoresLayout, ThemeLayout } from "./layouts";

import "./index.scss";

function App() {
    return (
        <StoresLayout>
            <ThemeLayout>
                <Routing />
            </ThemeLayout>
        </StoresLayout>
    );
}

export default App;
