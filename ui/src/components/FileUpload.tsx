import React, {ChangeEvent, useState} from "react";
import {TbX} from "react-icons/tb";
import axios from "axios";
import environment from "../lib/environment.ts"

const FileUpload: React.FC = () => {
    const [files, setFiles] = useState<File[]>([]);
    const [uploading, setUploading] = useState<boolean>(false);
    const [uploadResults, setUploadResults] = useState<string[]>([]);

    const handleFileChange = (e: ChangeEvent<HTMLInputElement>) => {
        if (e.target.files) {
            const selectedFiles = Array.from(e.target.files);
            setFiles(selectedFiles);
            setUploadResults([]);
        }
    };

    const formatSize = (bytes: number) => {
        const sizes = ["Bytes", "KB", "MB", "GB"];
        if (bytes === 0) return "0 Byte";
        const i = Math.floor(Math.log(bytes) / Math.log(1024));
        return parseFloat((bytes / Math.pow(1024, i)).toFixed(2)) + " " + sizes[i];
    };

    const onRemove = (file: File) => {
        let _files = Array.from(files);
        _files = _files.filter((f: File) => f.name !== file.name);
        setFiles(_files);
    }

    const uploadFiles = async () => {
        setUploading(true);
        const results: string[] = [];

        for (const file of files) {
            if (!["audio/flac", "audio/mpeg", "audio/x-m4a"].includes(file.type)) {
                results.push(`❌ ${file.name}: unsupported file type`);
                continue;
            }
            // TODO: consolidate all api calls into another file
            const formData = new FormData();
            formData.append("file", file);

            try {
                await axios.post(`${environment.apiURL}/v1/files`, formData, {
                    headers: {
                        "Content-Type": "multipart/form-data",
                    },
                });
                results.push(`✅ ${file.name}`);
            } catch (error) {
                console.error(error);
                results.push(`❌ ${file.name}: upload failed`);
            }
        }
        setFiles([])
        setUploadResults(results);
        setUploading(false);
    };
    return (
        <div>
            <div className="upload-button-wrapper">
                <button className="upload-button">Select Files</button>
                <input
                    id="file_input"
                    type="file"
                    multiple
                    accept=".mp3,.flac,.m4a,audio/m4a,audio/mp3,audio/flac"
                    onChange={handleFileChange}
                />
            </div>
            <p className="mt-1 text-sm text-gray-500">You can select multiple files.</p>

            <div className="mt-6">
                <div>
                    {files.length > 0 && (
                        <button
                            onClick={uploadFiles}
                            disabled={uploading}
                            className="button"
                        >
                            {uploading ? "Uploading..." : "Upload All"}
                        </button>
                    )}
                </div>
                <h3>File Details:</h3>
                <div>
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
                                <td  className="center">{formatSize(file.size)}</td>
                                <td  className="center">{file.type || "—"}</td>
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
                    {uploadResults.length > 0 && (
                        <div>
                            {uploadResults.map((msg, idx) => (
                                <div key={idx}>
                                    {msg}
                                </div>
                            ))}
                        </div>
                    )}
                </div>
            </div>
        </div>
    );
};

export default FileUpload;
