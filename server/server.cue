package main

import (
    "dagger.io/dagger"

    "universe.dagger.io/docker"
    "universe.dagger.io/go"
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
    } 

    actions: {
        test: go.#Test & {
            source:  client.filesystem."./".read.contents
            package: "./..."
        }

        build: #GOBuild & {
            source: client.filesystem."./".read.contents
        }
    }

}