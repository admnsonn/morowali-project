FROM node:lts-alpine as build-stage



COPY . .
RUN npm install
RUN npm install leaflet
RUN npm install chart.js
RUN npm install vue-chartjs



CMD [ "npm", "run", "dev" ]