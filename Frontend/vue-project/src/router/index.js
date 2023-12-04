import { createRouter, createWebHistory } from "vue-router";
import BerandaView from "../views/beranda.vue";
import PemerintahView from "../views/pemerintah.vue";
import ProfilView from "../views/profildesa.vue";
import DatadesaView from "../views/datadesa.vue";
import WisataView from "../views/wisata.vue";
import IdmView from "../views/idm.vue";
import InformasiView from "../views/informasi.vue";
import LoginView from "../views/login.vue";
import MediaView from "../views/media.vue";
import PotensiView from "../views/potensi.vue";
import WilayahView from "../views/wilayah.vue";
import BeritaView from "../views/berita.vue";
import ContentView from "../views/contentberita.vue";
import LembagaView from "../views/lembaga.vue";
import SejarahView from "../views/sejarah.vue";
import GaleriView from "../views/galeri.vue";
import PendidikanView from "../views/datapendidikan.vue";
import BerandaAdmin from "../views/AdminLogin/berandaadmin.vue";
import BerandaWarga from "../views/WargaLogin/berandawarga.vue";
import UmkmView from "../views/umkm.vue";
import ArtikelView from "../views/artikel.vue";


const isAdmin = () => {
  const isLoggedIn = localStorage.getItem('isLoggedIn');
  const userRole = localStorage.getItem('role_pengguna');
  
    // ADIT UBAH 
  return isLoggedIn === 'true' && userRole === 'Admin';
};

const authGuardAdmin = (to, from, next) => {
  if (!isAdmin()) {
    next('/');
  } else {
    next();
  }
};

// Adit  Tambah Kondisi Warga
const isWarga = () => {
  const isLoggedIn = localStorage.getItem('isLoggedIn');
  const userRole = localStorage.getItem('role_pengguna');
  
    // ADIT UBAH 
  return isLoggedIn === 'true' && userRole === 'Warga';
};

const authGuardWarga = (to, from, next) => {
  if (!isWarga()) {
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
  { path: '/wisata', component: WisataView },
  { path: '/idm', component: IdmView },
  { path: '/informasi', component: InformasiView },
  { path: '/login', component: LoginView },
  { path: '/media', component: MediaView },
  { path: '/potensi', component: PotensiView },
  { path: '/sejarah', component: SejarahView },
  { path: '/wilayah', component: WilayahView },
  { path: '/berita', component: BeritaView },
  { path: '/artikel', component: ArtikelView },
  { path: '/galeri-foto', component: GaleriView },
  { path: '/umkm', component: UmkmView },
  { path: '/lembaga', component: LembagaView },
  { path: '/data-pendidikan', component: PendidikanView },
  { path: '/contentberita', component: ContentView },
  // Buat Admin
  { path: '/beranda-admin', component: BerandaAdmin, meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  // Buar Warga
  { path: '/beranda-warga', component: BerandaWarga, meta: { requiresAuth: true }, beforeEnter: authGuardWarga }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;
