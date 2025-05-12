import React from "react";

interface BoxProps {
    children: React.ReactNode
    className?:string
}

const Box: React.FC<BoxProps> = ({
                                     children,
                                     className,
                                 }) => {
    return (
        <div className={className}>
            {children}
        </div>
    )
}

export default Box