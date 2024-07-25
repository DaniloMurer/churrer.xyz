FROM node:22.3.0-slim as base
EXPOSE 3000
ENV NODE_ENV=production

WORKDIR /src


FROM base as build
EXPOSE 3000
COPY . .
RUN npm install --production=false
RUN npm run build
FROM base

ENV PORT=$PORT

COPY --from=build /src/.output /src/.output

CMD [ "node", ".output/server/index.mjs" ]