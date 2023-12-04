FROM node:lts-alpine
RUN npm install -g http-server
WORKDIR /app
COPY package*.json ./
RUN npm install
RUN npm install leaflet
RUN npm install chart.js
RUN npm install vue-chartjs
COPY . .
RUN npm run build
EXPOSE 8080
CMD [ "http-server", "dist"]