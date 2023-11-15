import { createRouter, createWebHistory } from "vue-router";
import BerandaView from "../views/beranda.vue"
import PemerintahView from "../views/pemerintah.vue"
import ProfilView from "../views/profildesa.vue"

const routes = [
    { path: '', component: BerandaView },
    { path: '/profil-desa', component: ProfilView },
    { path: '/pemerintah-desa', component: PemerintahView },
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router