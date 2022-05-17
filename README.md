# Create tempalte nextjs
```sh
npx create-next-app@latest --ts next_jest_sample
```
# Jest settings
```sh
npm install --save-dev jest @testing-library/react @testing-library/jest-dom jest-environment-jsdom
npm install next@canary  
```
# Config something ...
- create jest.config.js
- add swcMinify: false, in next.config.js
- add .vscode
# Get started
```
npm run dev
```