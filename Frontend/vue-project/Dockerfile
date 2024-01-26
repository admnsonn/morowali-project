FROM node:lts-alpine as build-stage
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm install
RUN npm install leaflet
RUN npm install chart.js
RUN npm install vue-chartjs
RUN npm run build


EXPOSE 5173

CMD ["npm run dev"]
