FROM node:lts-alpine as build-stage
WORKDIR /app
COPY package*.json /app/
RUN npm install \
    npm install leaflet \
    npm install chart.js \
    npm install vue-chart.js \
    npm run build
COPY . /app/

# production stage
FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]