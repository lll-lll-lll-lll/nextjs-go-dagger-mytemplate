# syntax=docker/dockerfile:1
FROM node:14.17.4-alpine
ENV NODE_ENV=dev
WORKDIR /client
RUN apk update


COPY package*.json ./
RUN yarn install  --dev

COPY . ./client
CMD ["yarn", "run", "dev"]