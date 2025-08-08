import React, {ChangeEvent, useState} from "react";
import FileList from "./FileList.tsx";
import {postFile} from '../lib/api.ts'

const FileUpload: React.FC = () => {
    const [files, setFiles] = useState<File[]>([]);
    const [uploading, setUploading] = useState<boolean>(false);
    const [uploadResults, setUploadResults] = useState<string[]>([]);

    const handleFileChange = (e: ChangeEvent<HTMLInputElement>) => {
        if (e.target.files) {
            const selectedFiles = Array.from(e.target.files);
            setFiles(prevFiles => {
                // Combine existing files with new files, avoiding duplicates
                const allFiles = [...prevFiles, ...selectedFiles];
                const uniqueFiles = allFiles.filter((file, index, self) => 
                    index === self.findIndex(f => f.name === file.name)
                );
                return uniqueFiles;
            });
            setUploadResults([]);
        }
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
            const formData = new FormData();
            formData.append("file", file);

            try {
                await postFile(formData);
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
                <div>
                    <FileList files={files} onRemove={onRemove}/>
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
