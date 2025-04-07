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
        <div className="flex h-full">
            <div className="hidden md:flex flex-col gap-y-2 bg-black h-full w-[300px] p-2">
                <Box>
                    <div className="flex flex-col gap-y-4 px-5 py-4">
                        {routes.map((item) => (
                            <SidebarItem key={item.label} {...item}></SidebarItem>
                        ))}
                    </div>
                </Box>
                <Box className="overflow-y-auto h-full">
                    <Library/>
                </Box>
            </div>
            <main className="h-full flex-1/6 overflow-y-auto py-2">
                {children}
            </main>
        </div>
    )
}

export default Sidebar