package main

import (
    "dagger.io/dagger"

    "universe.dagger.io/docker"
    "universe.dagger.io/go"
    "universe.dagger.io/netlify"
)

#GOBuild: docker.#Dockerfile  & {
    dockerfile: { 
        path: "go.Dockerfile"
    }
}

dagger.#Plan & {
    client: {
        filesystem: "./": read: contents: dagger.#FS
        network: "unix:///var/run/docker.sock": connect: dagger.#Socket
        env: {
            NETLIFY_TOKEN: dagger.#Secret
            NETLIFY_SITE_NAME: string
        }
    } 

    actions: {
        test: go.#Test & {
            source:  client.filesystem."./".read.contents
            package: "./..."
        }

        build: #GOBuild & {
            source: client.filesystem."./".read.contents
        }

        // 開発用のnetlifyへデプロイ
        deploy: netlify.#Deploy & {
            contents: build.output.rootfs
            site: "ec-server"
            token: client.env.NETLIFY_TOKEN
            create: false
        }
    }
}