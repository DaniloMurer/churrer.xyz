FROM node:22.3.0-slim as base
EXPOSE 3000
ENV NODE_ENV=production

WORKDIR /src

# get my bad boy yarn for building
RUN npm install yarn

FROM base as build
EXPOSE 3000
COPY . .
RUN yarn install --production=false
RUN yarn run build
FROM base

ENV PORT=$PORT

COPY --from=build /src/.output /src/.output

CMD [ "node", ".output/server/index.mjs" ]