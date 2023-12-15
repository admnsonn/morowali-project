# vue-project

This template should help get you started developing with Vue 3 in Vite.

## Recommended IDE Setup

[VSCode](https://code.visualstudio.com/) + [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) (and disable Vetur) + [TypeScript Vue Plugin (Volar)](https://marketplace.visualstudio.com/items?itemName=Vue.vscode-typescript-vue-plugin).

## Customize configuration

See [Vite Configuration Reference](https://vitejs.dev/config/).

## Project Setup

```sh
npm install
npm install leaflet
npm install chart.js
npm install vue-chartjs
```

### Compile and Hot-Reload for Development

```sh
npm run dev
```

### Compile and Minify for Production

```sh
npm run build
```

# Langkah run Frontend di aws
* Clone dulu, lalu...
```shell
$ cd morowali-project/Frontend/vue-project
```
* Install dependencies
```shell
$ npm install
$ npm install leaflet
$ npm install chart.js
$ npm install vue-chartjs
```
* Pake pm2 buat run di background...
```shell
$ pm2 start npm --name "vue-project" -- run dev --prefix /home/ec2-user/morowali-project/Frontend/vue-project
```