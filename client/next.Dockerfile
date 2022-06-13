# syntax=docker/dockerfile:1
FROM node:14.17.4-alpine
ENV NODE_ENV=production
WORKDIR /usr/src/app
RUN apk update


COPY package*.json ./
RUN yarn install  --production

COPY . .
CMD ["next", "start"]