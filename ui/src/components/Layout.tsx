import {Route, Routes} from "react-router-dom";
import Home from "../pages/Home.tsx";
import Sidebar from "./Sidebar.tsx";
import Search from "../pages/Search.tsx";

const Layout = () => {
        return (
            <div className="layout">
                <Sidebar>
                    <div>
                        I am another section within main layout
                        <hr/>
                    </div>
                    <Routes>
                        <Route path="/" element={<Home/>}/>
                        <Route path="/search" element={<Search/>}/>
                    </Routes>
                </Sidebar>
                <div>
                </div>
            </div>
        );
    }
;

export default Layout;