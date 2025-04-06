import {Route, Routes} from "react-router-dom";
import Home from "../pages/Home.tsx";
import About from "../pages/About.tsx";
import Sidebar from "./Sidebar.tsx";
import Search from "../pages/Search.tsx";

const Layout = () => {
    return (
        <div className="layout">
            <Sidebar>
                <Routes>
                    <Route path="/" element={<Home/>}/>
                    <Route path="/about" element={<About/>}/>
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