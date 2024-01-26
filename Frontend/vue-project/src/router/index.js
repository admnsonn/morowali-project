import { createRouter, createWebHistory } from 'vue-router';

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

const isWarga = () => {
  const isLoggedIn = localStorage.getItem('isLoggedIn');
  const userRole = localStorage.getItem('role_pengguna');
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
  { path: '', component: () => import('../views/beranda.vue') },
  { path: '/profil-desa', component: () => import('../views/profildesa.vue') },
  { path: '/pemerintah-desa', component: () => import('../views/pemerintah.vue') },
  { path: '/data-desa', component: () => import('../views/datadesa.vue') },
  { path: '/wisata', component: () => import('../views/wisata.vue') },
  { path: '/idm', component: () => import('../views/idm.vue') },
  { path: '/informasi', component: () => import('../views/informasi.vue') },
  { path: '/login', component: () => import('../views/login.vue') },
  { path: '/media', component: () => import('../views/media.vue') },
  { path: '/potensi', component: () => import('../views/potensi.vue') },
  { path: '/sejarah', component: () => import('../views/sejarah.vue') },
  { path: '/wilayah', component: () => import('../views/wilayah.vue') },
  { path: '/berita', component: () => import('../views/berita.vue') },
  { path: '/artikel', component: () => import('../views/artikel.vue') },
  { path: '/galeri-foto', component: () => import('../views/galeri.vue') },
  { path: '/umkm', component: () => import('../views/umkm.vue') },
  { path: '/lembaga', component: () => import('../views/lembaga.vue') },
  { path: '/data-pendidikan', component: () => import('../views/datapendidikan.vue') },
  { path: '/contentberita', component: () => import('../views/contentberita.vue') },
  // Buat Admin
  { path: '/beranda-admin', component: BerandaAdmin, meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/warga-desa', component: Wargadesa, meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/berita-desa', component: BeritaDesa, meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/umkm-desa', component: UMKMDesa, meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/tambah-umkm', component: TambahUMKM, meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/tambah-berita', component: TambahBerita, meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/update-berita/:id', component: UpdateBerita, meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/tambah-warga', component: TambahWarga, meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/wisata-desa', component: WisataDesa, meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/potensi-desa', component: PotensiDesa, meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/kreatifitas-desa', component: KreatifitasDesa, meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/detail-warga/:id', component: DetailWarga, meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },

  // Buat Warga
  { path: '/beranda-warga', component: () => import('../views/WargaLogin/berandawarga.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardWarga }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;
