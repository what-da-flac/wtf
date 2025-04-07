import express, { Request, Response, NextFunction } from 'express';
import { createProxyMiddleware } from 'http-proxy-middleware';
import multer from "multer";
import axios from "axios"
import FormData from "form-data";
import fs from "fs";
import path from "path";
import cors from "cors"

const app = express();

// enable cors
app.use(cors())

// Temp disk storage
const upload = multer({ dest: "temp_uploads/" });

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
    target: 'http://localhost:8080',
    changeOrigin: true,
    pathRewrite: {
      '^/system': '',
    },
  })
);

// example internal route handled by Node.js
app.get('/hello', (req: Request, res: Response) => {
  res.json({ id: 1, first_name: 'John', last_name: "Doe" });
});

app.post(
  "/api/files",
  upload.single("file"),
  async (req: Request, res: Response, next: NextFunction): Promise<void> => {
    const fileReq = req as MulterRequest;

    if (!fileReq.file) {
      res.status(400).json({ error: "No file uploaded" });
      return;
    }

    const filePath = path.join(__dirname, fileReq.file.path);

    const form = new FormData();
    form.append("file", fs.createReadStream(filePath), fileReq.file.originalname);

    try {
      const response = await axios.post("https://example.com/api/receive", form, {
        headers: form.getHeaders(),
      });

      fs.unlink(filePath, () => {});

      res.status(response.status).json({
        message: "File forwarded successfully",
        remoteResponse: response.data,
      });
    } catch (err) {
      fs.unlink(filePath, () => {});
      const error = err as Error; // cast `unknown` to `Error`
      console.error("Forwarding failed:", error.message);
      res.status(500).json({ error: "Failed to forward file" });
    }
  }
);
const PORT = process.env.PORT || 3000;

app.listen(PORT, () => {
  console.log(`Gateway listening on port ${PORT}`);
});
