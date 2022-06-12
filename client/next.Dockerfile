# syntax=docker/dockerfile:1
FROM node:14.17.4-alpine
ENV NODE_ENV=production
WORKDIR /usr/src/app

RUN apt-get -y update 


COPY package*.json ./
RUN yanr install  --production

COPY . .
CMD ["next", "start"]