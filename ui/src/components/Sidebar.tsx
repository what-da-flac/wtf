import React, {useMemo} from "react";
import {HiHome} from "react-icons/hi";
import {BiSearch} from "react-icons/bi";
import {useLocation} from "react-router-dom";

interface SidebarProps {
    children: React.ReactNode;
}

const Sidebar: React.FC<SidebarProps> = ({
                                             children
                                         }) => {
    const location = useLocation();
    const routes = useMemo(() => [
        {
            icon: HiHome,
            label: "Home",
            active: location.pathname !== "/search",
            href: "/",
        },
        {
            icon: BiSearch,
            label: "Search",
            active: location.pathname === "/search",
            href: "/search",
        }
    ], [location.pathname]);
    return (
        <>
            <div>Sidebar</div>
            {children}
        </>
    )
}

export default Sidebar;