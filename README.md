<p class="alert">作成途中</p>

# Summary
- client
    - nextjs
    - typescript
    - jest
- server
    - go
    - mux
- db
    - postgresql
- hosting
    - client -> vercel
    - server -> Azure Function

## Dager Settings
※ dagger already installed<br>

create cue.mod directory. initialize dagger settings. <br>
if you already have cue.mod directory, not execute `dagger project init`, just ``` dagger project update```
```sh
dagger project init
```
```sh
dagger project update
```
execute CI
```sh
dagger do build
```


## Azure Functions
- trigger directory
- function.json
- host.json
- .funcignore
- local.settings.json
- .gitignore

# Github Secret Settings
- VERCEL_TOKEN
- ORG_ID
- PROJECT_ID
- AZURE_FUNCTIONAPP_PUBLISH_PROFILE