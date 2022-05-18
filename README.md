# Create tempalte nextjs
```sh
npx create-next-app@latest --ts next_jest_sample
```

# Jest settings
```sh
npm install --save-dev jest @testing-library/react @testing-library/jest-dom jest-environment-jsdom
npm install next@canary  
```
## Config something ...
- create jest.config.js
- add swcMinify: false in next.config.js
- add .vscode
- add .babelrc
- add CI file

## push and Github Actions succeed



# dagger settings
- install dagger: https://docs.dagger.io/1200/local-dev
- add next_jest.cue
- modify CI file(/.github/workflow/main.yml)


## dager CI
※ dagger installed<br>

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