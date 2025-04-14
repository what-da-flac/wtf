import express, {NextFunction, Request, Response} from 'express';
import {createProxyMiddleware} from 'http-proxy-middleware';
import multer from "multer";
import axios from "axios"
import FormData from "form-data";
import cors from "cors"

const app = express();

// read environment
const env = {
    gatewayUrl : process.env.GATEWAY_URL,
    port: process.env.PORT || 3000,
    urlPrefix: process.env.API_URL_PREFIX,
}

console.log(`environment: ${JSON.stringify(env)}`)

// enable cors
app.use(cors())

// Use memoryStorage to avoid writing files to disk
const upload = multer({storage: multer.memoryStorage()});

// middleware example: log all requests
app.use((req: Request, res: Response, next: NextFunction) => {
    console.log(`[${req.method}] ${req.url}`);
    next();
});

// Extend the Express Request type to include `file`
interface MulterRequest extends Request {
    file: Express.Multer.File;
}

// proxy that rewrites url removing "/system" prefix
// and routes requests to Golang service
app.use(
    '/system',
    createProxyMiddleware({
        target: env.gatewayUrl,
        changeOrigin: true,
        pathRewrite: {
            '^/system': '',
        },
    })
);

app.post(
    "/api/files",
    upload.single("file"),
    async (req: Request, res: Response): Promise<void> => {
        const fileReq = req as MulterRequest;

        if (!fileReq.file) {
            res.status(400).json({error: "No file uploaded"});
            return;
        }

        try {
            const form = new FormData();
            console.log(JSON.stringify({
                mime_type: fileReq.file.mimetype,
                original_name: fileReq.file.originalname,
            }))
            form.append("file", fileReq.file.buffer, {
                filename: fileReq.file.originalname,
                contentType: fileReq.file.mimetype,
            })
            const uri = `${env.gatewayUrl}${env.urlPrefix}/v1/files`
            console.log(`uri: ${uri}`)
            const response = await axios.post(uri, form, {
                headers: {
                    ...form.getHeaders(),
                },
                maxContentLength: Infinity,
                maxBodyLength: Infinity,
            });
            res.status(response.status).json({
                message: "File forwarded successfully",
                remoteResponse: response.data,
            });
        } catch (err) {
            const error = err as Error; // cast `unknown` to `Error`
            console.error("Forwarding failed:", error.message);
            res.status(500).json({error: "Failed to forward file"});
        }
    }
);

app.listen(env.port, () => {
    console.log(`Gateway listening on port ${env.port}`);
});
