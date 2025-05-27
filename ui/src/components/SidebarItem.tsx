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
                                                     icon: Icon, label, href,
                                                 }) => {
    return (
        <Link
            to={href}
            className="box-red">
            <Icon size={26} />
            <p>{label}</p>
        </Link>
    )
}

export default SidebarItem