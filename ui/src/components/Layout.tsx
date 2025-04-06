import {Route, Routes} from "react-router-dom";
import Home from "../pages/Home.tsx";
import About from "../pages/About.tsx";
import Sidebar from "./Sidebar.tsx";

const Layout = () => {
    return (
        <div className="layout">
            <Sidebar/>
            <div>
                <Routes>
                    <Route path="/" element={<Home/>}/>
                    <Route path="/about" element={<About/>}/>
                </Routes>
            </div>
        </div>
    );
};

export default Layout;