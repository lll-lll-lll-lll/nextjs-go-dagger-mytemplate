## push and Github Actions succeed



## dager CI
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