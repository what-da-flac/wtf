FROM node:22-alpine AS builder
ARG SERVICE_NAME=""

WORKDIR /app

COPY services/${SERVICE_NAME}/package*.json ./
RUN npm install

COPY services/${SERVICE_NAME}/tsconfig.json ./
COPY services/${SERVICE_NAME}/src ./src

RUN npm run build

FROM node:22-alpine
ARG SERVICE_NAME=""

WORKDIR /app

COPY --from=builder /app/dist ./dist
COPY services/${SERVICE_NAME}/package*.json ./

RUN npm install --only=production

EXPOSE 3000

CMD ["node", "dist/index.js"]
