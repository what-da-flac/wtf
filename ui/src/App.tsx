import Layout from "./components/Layout.tsx";
import {BrowserRouter} from "react-router-dom";


function App() {

    return (
        <>
            <BrowserRouter>
                <Layout/>
            </BrowserRouter>
        </>
    )
}

export default App
