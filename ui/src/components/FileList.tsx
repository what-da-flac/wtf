import React from "react";
import {TbX} from "react-icons/tb";
import {formatSize} from "../lib/format.ts";

interface BoxProps {
    files: File[]
    onRemove: (file: File) => void
}

const FileList: React.FC<BoxProps> = ({
                                          onRemove,
                                          files,
                                      }) => {
    return (
        <table className="styled-table">
            <thead>
            <tr>
                <th>Name</th>
                <th>Size</th>
                <th>Type</th>
                <th></th>
            </tr>
            </thead>
            <tbody>
            {files.map((file, index) => (
                <tr key={index}>
                    <td>{file.name}</td>
                    <td className="center">{formatSize(file.size)}</td>
                    <td className="center">{file.type || "—"}</td>
                    <td onClick={() => onRemove(file)} className="center">
                        <button
                        >
                            <TbX/>
                        </button>
                    </td>
                </tr>
            ))}
            </tbody>
        </table>
    )
}

export default FileList