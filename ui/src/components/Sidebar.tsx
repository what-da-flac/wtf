import {useLocation} from "react-router-dom";
import {useMemo} from "react";
import {HiHome} from "react-icons/hi";
import {BiSearch} from "react-icons/bi";
import Box from "./Box.tsx";
import SidebarItem from "./SidebarItem.tsx";
import Library from "./Library.tsx";

interface SidebarProps {
    children: React.ReactNode
}

const Sidebar: React.FC<SidebarProps> = ({
                                             children
                                         }) => {
    const location = useLocation()

    const routes = useMemo(() => [
        {
            icon: HiHome,
            label: "Home",
            active: location.pathname === "/",
            href: "/",
        },
        {
            icon: BiSearch,
            label: "Search",
            active: location.pathname === "/search",
            href: "/search",
        }
    ], [location.pathname])
    return (
        <div className="sidebar-layout">
            <div className="sidebar">
                <Box>
                    <div className="sidebar-items sidebar-section">
                        {routes.map((item) => (
                            <SidebarItem key={item.label} {...item}></SidebarItem>
                        ))}
                    </div>
                </Box>
                <Box className="sidebar-overflow">
                    <Library/>
                </Box>
            </div>
            <main className="main-content">
                {children}
            </main>
        </div>
    )
}

export default Sidebar