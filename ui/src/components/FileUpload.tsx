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
        // setFiles([])
        setUploadResults(results);
        setUploading(false);
    };
    return (
        <div className="max-w-xl mx-auto mt-10 p-4 border border-gray-300 rounded-lg shadow-sm bg-white">
            <label
                htmlFor="file_input"
                className="block text-sm font-medium text-gray-700 mb-2"
            >
                Upload files
            </label>
            <input
                id="file_input"
                type="file"
                multiple
                accept=".mp3,.flac,.m4a,audio/m4a,audio/mp3,audio/flac"
                onChange={handleFileChange}
                className="block w-full text-sm text-gray-900 border border-gray-300 rounded-lg cursor-pointer bg-gray-50 focus:outline-none"
            />
            <p className="mt-1 text-sm text-gray-500">You can select multiple images or files.</p>

            {files.length > 0 && (
                <div className="mt-6">
                    <div className="flex justify-between items-center mb-2">
                        <button
                            onClick={uploadFiles}
                            disabled={uploading}
                            className="bg-blue-600 hover:bg-blue-700 text-white text-sm px-4 py-2 rounded disabled:opacity-50"
                        >
                            {uploading ? "Uploading..." : "Upload All"}
                        </button>
                    </div>
                    <h3 className="text-sm font-medium text-gray-700 mb-2">File Details:</h3>
                    <div className="overflow-auto">
                        <table className="min-w-full text-sm text-left border border-gray-200 rounded">
                            <thead className="bg-gray-100 text-gray-700">
                            <tr>
                                <th className="px-4 py-2">Name</th>
                                <th className="px-4 py-2">Size</th>
                                <th className="px-4 py-2">Type</th>
                                <th></th>
                            </tr>
                            </thead>
                            <tbody className="divide-y divide-gray-200">
                            {files.map((file, index) => (
                                <tr key={index} className="text-gray-700">
                                    <td className="px-4 py-2 truncate">{file.name}</td>
                                    <td className="px-4 py-2">{formatSize(file.size)}</td>
                                    <td className="px-4 py-2">{file.type || "—"}</td>
                                    <td onClick={() => onRemove(file)}>
                                        <button
                                            className="bg-red-500 hover:bg-red-700 text-white text-xs px-4 py-2 rounded cursor-pointer"
                                        >
                                            <TbX/>
                                        </button>
                                    </td>
                                </tr>
                            ))}
                            </tbody>
                        </table>
                        {uploadResults.length > 0 && (
                            <div className="mt-4 space-y-1">
                                {uploadResults.map((msg, idx) => (
                                    <div key={idx} className="text-sm text-gray-700">
                                        {msg}
                                    </div>
                                ))}
                            </div>
                        )}
                    </div>
                </div>
            )}
        </div>
    );
};

export default FileUpload;
