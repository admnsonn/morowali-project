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