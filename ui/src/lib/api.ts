import axios from "axios";
import environment from "./environment.ts";

export const postFile = async (formData: FormData) => {
    return axios.post(`${environment.apiURL}/v1/files`, formData, {
        headers: {
            "Content-Type": "multipart/form-data",
        },
    })
}

