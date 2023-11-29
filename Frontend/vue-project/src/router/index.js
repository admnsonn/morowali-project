import { createRouter, createWebHistory } from "vue-router";
<<<<<<< Updated upstream
import BerandaView from "../views/beranda.vue"
import PemerintahView from "../views/pemerintah.vue"
import ProfilView from "../views/profildesa.vue"
import DatadesaView from "../views/datadesa.vue"
import DestinasiView from "../views/destinasi.vue"
import IdmView from "../views/idm.vue"
import InformasiView from "../views/informasi.vue"
import LoginView from "../views/login.vue"
import MediaView from "../views/media.vue"
import PotensiView from "../views/potensi.vue"
import WilayahView from "../views/wilayah.vue"
import BeritaView from "../views/berita.vue"
import SejarahView from "../views/sejarah.vue"
import ContentView from "../views/contentberita.vue"

const routes = [
    { path: '', component: BerandaView },
    { path: '/profil-desa', component: ProfilView },
    { path: '/pemerintah-desa', component: PemerintahView },
    { path: '/data-desa', component: DatadesaView },
    { path: '/destinasi', component: DestinasiView },
    { path: '/idm', component: IdmView },
    { path: '/informasi', component: InformasiView },
    { path: '/login', component: LoginView },
    { path: '/media', component: MediaView },
    { path: '/potensi-desa', component: PotensiView },
    { path: '/sejarah-desa', component: PotensiView },
    { path: '/wilayah', component: WilayahView },
    { path: '/berita', component: BeritaView },
    { path: '/sejarah', component: SejarahView },
    { path: '/contentberita', component: ContentView },
]
=======
import BerandaView from "../views/beranda.vue";
import PemerintahView from "../views/pemerintah.vue";
import ProfilView from "../views/profildesa.vue";
import DatadesaView from "../views/datadesa.vue";
import DestinasiView from "../views/destinasi.vue";
import IdmView from "../views/idm.vue";
import InformasiView from "../views/informasi.vue";
import LoginView from "../views/login.vue";
import MediaView from "../views/media.vue";
import PotensiView from "../views/potensi.vue";
import WilayahView from "../views/wilayah.vue";
import BeritaView from "../views/berita.vue";
import SejarahView from "../views/sejarah.vue";
import BerandaAdmin from "../views/AdminLogin/berandaadmin.vue";

const isAdmin = () => {
  const isLoggedIn = localStorage.getItem('isLoggedIn');
  const userRole = localStorage.getItem('role_pengguna');
  
  
  return isLoggedIn === 'true' && userRole === 'Admin';
};

const authGuardAdmin = (to, from, next) => {
  if (!isAdmin()) {
    next('/');
  } else {
    next();
  }
};

const routes = [
  { path: '', component: BerandaView },
  { path: '/profil-desa', component: ProfilView },
  { path: '/pemerintah-desa', component: PemerintahView },
  { path: '/data-desa', component: DatadesaView },
  { path: '/destinasi', component: DestinasiView },
  { path: '/idm', component: IdmView },
  { path: '/informasi', component: InformasiView },
  { path: '/login', component: LoginView },
  { path: '/media', component: MediaView },
  { path: '/potensi-desa', component: PotensiView },
  { path: '/sejarah-desa', component: SejarahView },
  { path: '/wilayah', component: WilayahView },
  { path: '/berita', component: BeritaView },
  { path: '/sejarah', component: SejarahView },
  { path: '/beranda-admin', component: BerandaAdmin, meta: { requiresAuth: true }, beforeEnter: authGuardAdmin }
];
>>>>>>> Stashed changes

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;
