import { IconType } from "react-icons"
import React from "react";
import {Link} from "react-router-dom";

interface SidebarItemProps {
    icon: IconType
    label: string
    active?:boolean
    href: string
}

const SidebarItem: React.FC<SidebarItemProps> = ({
                                                     icon: Icon, label, active, href,
                                                 }) => {
    return (
        <Link
            to={href}
            className="flex flex-row h-auto
            items-center w-full gap-x-4 font-medium 
            cursor-pointer hover:text-white transition 
            text-neutral-400 py-1"
        >
            <Icon size={26} />
            <p className="truncate w-full">{label}</p>
        </Link>
    )
}

export default SidebarItem