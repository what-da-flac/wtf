import express, { Request, Response, NextFunction } from 'express';
import { createProxyMiddleware } from 'http-proxy-middleware';

const app = express();

// middleware example: log all requests
app.use((req: Request, res: Response, next: NextFunction) => {
  console.log(`[${req.method}] ${req.url}`);
  next();
});

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
  
const PORT = process.env.PORT || 3000;

app.listen(PORT, () => {
  console.log(`Gateway listening on port ${PORT}`);
});
