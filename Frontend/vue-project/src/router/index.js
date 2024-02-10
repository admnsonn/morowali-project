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
  { path: '/beranda-admin', component: () => import('../views/AdminLogin/beranda/berandaadmin.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/warga-desa', component: () => import('../views/AdminLogin/warga/wargadesa.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/tambah-warga', component: () => import('../views/AdminLogin/warga/tambahwarga.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/detail-warga/:id', component: () => import('../views/AdminLogin/warga/detailwarga.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/update-warga/:id', component: () => import('../views/AdminLogin/warga/updatewarga.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/berita-desa', component: import('../views/AdminLogin/berita/beritadesa.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/tambah-berita', component: import('../views/AdminLogin/berita/tambahberita.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/update-berita/:id', component: import('../views/AdminLogin/berita/updateberita.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/detail-berita/:id', component: import('../views/AdminLogin/berita/detailberita.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/umkm-desa', component: import('../views/AdminLogin/umkm/umkmdesa.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/tambah-umkm', component: import('../views/AdminLogin/umkm/tambahumkm.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/update-umkm/:id', component: import('../views/AdminLogin/umkm/updateumkm.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/detail-umkm/:id', component: import('../views/AdminLogin/umkm/detailumkm.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  // Admin: Wisata
  { path: '/wisata-desa', component: import('../views/AdminLogin/wisata/wisatadesa.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/detail-wisata/:id', component: import('../views/AdminLogin/wisata/detailwisata.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/update-wisata/:id', component: import('../views/AdminLogin/wisata/updatewisata.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/tambah-wisata', component: import('../views/AdminLogin/wisata/tambahwisata.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  // Admin: Potensi
  { path: '/potensi-desa', component: import('../views/AdminLogin/potensi/potensidesa.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/detail-potensi/:id', component: import('../views/AdminLogin/potensi/detailpotensi.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/update-potensi/:id', component: import('../views/AdminLogin/potensi/updatepotensi.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/tambah-potensi', component: import('../views/AdminLogin/potensi/tambahpotensi.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  // Admin: Wilayah
  { path: '/wilayah-desa', component: import('../views/AdminLogin/wilayah/wilayahdesa.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/update-wilayah/:id', component: import('../views/AdminLogin/wilayah/updatewilayah.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  // 
  { path: '/kreatifitas-desa', component: import('../views/AdminLogin/kreatifitas/kreatifitasdesa.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/tambah-kreatifitas', component: import('../views/AdminLogin/kreatifitas/tambahkreatifitas.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/update-kreatifitas/:id', component: import('../views/AdminLogin/kreatifitas/updatekreatifitas.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/detail-kreatifitas/:id', component: import('../views/AdminLogin/kreatifitas/detailkreatifitas.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  // Admin: Dropdown Pemerintahan
  { path: '/pemerintahan/visi-misi', component: import('../views/AdminLogin/pemerintah/visimisi.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/pemerintahan/update-visi-misi/:id', component: import('../views/AdminLogin/pemerintah/updatevismis.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/pemerintahan/kepala-desa', component: import('../views/AdminLogin/pemerintah/kepdes.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/pemerintahan/tambah-kepala-desa', component: import('../views/AdminLogin/pemerintah/tambahkepdes.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/pemerintahan/sambutan-kades', component: import('../views/AdminLogin/pemerintah/sambutan.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/pemerintahan/update-sambutan-kades/:id', component: import('../views/AdminLogin/pemerintah/updatesambutan.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/pemerintahan/sejarah-desa', component: import('../views/AdminLogin/pemerintah/sejarah.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/pemerintahan/update-sejarah-desa/:id', component: import('../views/AdminLogin/pemerintah/updatesejarah.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/pemerintahan/program-kerja', component: import('../views/AdminLogin/pemerintah/proker.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/pemerintahan/update-program-kerja/:id', component: import('../views/AdminLogin/pemerintah/updateproker.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/pemerintahan/peraturan', component: import('../views/AdminLogin/pemerintah/peraturan.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/pemerintahan/update-peraturan/:id', component: import('../views/AdminLogin/pemerintah/updateperaturan.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/pemerintahan/struktur-organisasi', component: import('../views/AdminLogin/pemerintah/struktur.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/pemerintahan/pegawai', component: import('../views/AdminLogin/pemerintah/pegawai.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  // Dropdown Informasi
  { path: '/informasi/produk-hukum', component: import('../views/AdminLogin/informasi/produk_hukum.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },

  // Admin: IDM
  { path: '/idm-beranda', component: import('../views/AdminLogin/idm/idm-beranda.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },
  { path: '/idm-management', component: import('../views/AdminLogin/idm/idm-management.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardAdmin },

  // Buat Warga
  { path: '/beranda-warga', component: () => import('../views/WargaLogin/beranda/berandawarga.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardWarga },
  { path: '/skbi', component: () => import('../views/WargaLogin/skbi/skbi.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardWarga },
  { path: '/skm-skck', component: () => import('../views/WargaLogin/skm-skck/skm-skck.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardWarga },
  { path: '/skp', component: () => import('../views/WargaLogin/skp/skp.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardWarga },
  { path: '/skpd', component: () => import('../views/WargaLogin/skpd/skpd.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardWarga },
  { path: '/skrn', component: () => import('../views/WargaLogin/skrn/skrn.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardWarga },
  { path: '/sktu', component: () => import('../views/WargaLogin/sktu/sktu.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardWarga },
  { path: '/sktm', component: () => import('../views/WargaLogin/sktm/sktm.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardWarga },
  { path: '/sku', component: () => import('../views/WargaLogin/sku/sku.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardWarga },
  { path: '/skbm', component: () => import('../views/WargaLogin/skbm/skbm.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardWarga },
  { path: '/skbm', component: () => import('../views/WargaLogin/skbm/skbm.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardWarga },
  { path: '/skd', component: () => import('../views/WargaLogin/skd/skd.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardWarga },
  { path: '/sksm', component: () => import('../views/WargaLogin/sksm/sksm.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardWarga },
  { path: '/skktm', component: () => import('../views/WargaLogin/skktm/skktm.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardWarga },
  { path: '/skkm', component: () => import('../views/WargaLogin/skkm/skkm.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardWarga },
  { path: '/skik', component: () => import('../views/WargaLogin/skik/skik.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardWarga },
  { path: '/spipj', component: () => import('../views/WargaLogin/spipj/spipj.vue'), meta: { requiresAuth: true }, beforeEnter: authGuardWarga },

];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;
