import {Route, Routes} from "react-router-dom";
import Home from "../pages/Home.tsx";
import Sidebar from "./Sidebar.tsx";
import Search from "../pages/Search.tsx";
import LibraryNew from "../pages/LibraryNew.tsx";

const Layout = () => {
        return (
            <div>
                <Sidebar>
                    <div className="box-green">
                        <Routes>
                            <Route path="/" element={<Home/>}/>
                            <Route path="/library/new" element={<LibraryNew/>}/>
                            <Route path="/search" element={<Search/>}/>
                        </Routes>
                    </div>
                </Sidebar>
                <div>
                </div>
            </div>
        );
    }
;

export default Layout;