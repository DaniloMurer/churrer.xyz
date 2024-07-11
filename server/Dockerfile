FROM node:22.3.0 as builder

WORKDIR /app/server

COPY . .

RUN npm install
RUN npm run build

FROM node:22.3.0-alpine
COPY --from=builder /app/server ./

CMD ["npm", "run", "start:prod"]