# Jest settings
```sh
npm install --save-dev jest @testing-library/react @testing-library/jest-dom jest-environment-jsdom
npm install next@canary  
```
## Config something ...
- create jest.config.js
- add swcMinify: false, in next.config.js
- create .vscode
- create .babelrc

## Test
```sh
npm run test
```

## Get started
```sh
npm run dev
```



# dagger settings
- install dagger: https://docs.dagger.io/1200/local-dev
- add /cue.mod
- add next_jest.cue
- modify CI file(/.github/workflow/main.yml)


## dager CI
```sh
dagger do build
```