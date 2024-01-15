// import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import "bootstrap/dist/css/bootstrap.min.css"
import "bootstrap"
import 'bootstrap-icons/font/bootstrap-icons.css';
import { OhVueIcon, addIcons } from "oh-vue-icons";
import { AiAcademiaSquare } from "oh-vue-icons/icons";

addIcons(AiAcademiaSquare);

const app = createApp(App);

// Add the OhVueIcon component to the app
app.component("v-icon", OhVueIcon);

// Use the router and mount the app
app.use(router).mount('#app');
