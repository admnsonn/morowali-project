<template>
  <div class="gambar">
    <img
      :src="imageUrl"
      alt="Gambar Berita"
      class="img-fluid gambar"
      width="2000"
    />
    <h1 class="text_tt text-start">{{ title }}</h1>
    <h4 class="text_st text-start">{{ subtitle }}</h4>
  </div>

  <div class="container py-5">
    <div
      v-for="(contentsatu, index) in penjelasan"
      :key="index"
      class="border rounded row align-items-center pt-3 pb-3 mb-4 with-shadow"
      :class="{ 'with-shadow': isHovered }"
      @mouseenter="addShadow"
      @mouseleave="removeShadow"
    >
      <div class="container mt-4">
        <div class="row">
          <div class="col-12 d-md-none mb-4">
            <img
              src="/src/assets/img/Artikel.png"
              alt="Latest Image"
              class="img-fluid"
            />
          </div>

          <div class="col-md-6 order-md-2 d-none d-md-block">
            <img
              src="/src/assets/img//Artikel.png"
              alt="Gambar Berita"
              class="img-fluid"
            />
          </div>

          <div class="col-md-6 order-md-1">
            <div class="text-center text-md-start">
              <h2 class="mb-3 d-md-none warna-judul-artikel">
                {{ contentsatu.judul }}
              </h2>
              <div class="d-none d-md-block">
                <h2 class="mb-2 warna-judul-artikel">
                  {{ contentsatu.judul }}
                </h2>
                <br />
              </div>
              <p class="mb-4 sub-berita">{{ contentsatu.teks }}</p>
              <a href="/sejarah">
                <button class="btn btn-secondary" @click="showDetail(item)">
                  Selengkapnya
                </button>
              </a>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <div
    v-for="(contentdua, index) in idm"
    :key="index"
    class="container-lg mt-2"
  >
    <h2 class="text-idm mb-4">Index Desa Membangun</h2>
    <p class="text-sidm mb-4">Lihat total index desa</p>
    <img
      v-bind:src="contentdua.imageUrl"
      alt="Gambar Berita"
      class="img-fluid gambar"
      width="2000"
    />
    <p class="text-ssidm mt-3">{{ contentdua.teks }}</p>
    <a href="/idm">
      <button class="btn1 btn btn-secondary mt-4" @click="showDetail(item)">
        Selengkapnya
      </button>
    </a>
  </div>

  <div class="container-lg mt-5 py-5">
    <h2 class="text-idm mb-4">Struktur Organisasi</h2>
    <p class="text-sidm mb-4">Lihat total struktur organisasi</p>
    <div class="row">
      <div
        v-for="(contentketiga, index) in struktur"
        :key="index"
        class="col-md-4 mb-4"
      >
        <img
          v-bind:src="contentketiga.imageUrl"
          :alt="contentketiga.name"
          class="img-fluid foto-kepdes"
        />
        <div class="text-center mt-2">
          <p class="font-weight-bold nama-kepdes">{{ contentketiga.name }}</p>
          <p class="periode-kepdes">{{ contentketiga.periode }}</p>
        </div>
      </div>
    </div>
    <a href="/sejarah">
      <button class="btn1 btn btn-secondary mt-2" @click="showDetail(item)">
        Selengkapnya
      </button>
    </a>
  </div>

  <div class="container-lg mt-5">
    <h2 class="text-idm mb-4">Arsip Berita dan Artikel</h2>
    <p class="text-sidm mb-4">Berita terbaru dari desa Bahomoleo</p>
    <div class="container py-2">
      <div
        v-for="(contentkeempat, index) in berita"
        :key="index"
        class="border rounded row align-items-center pt-3 pb-3 mb-4 with-shadow"
        :class="{ 'with-shadow': isHovered }"
        @mouseenter="addShadow"
        @mouseleave="removeShadow"
      >
        <div class="container mt-4">
          <div class="row">
            <div class="col-md-6">
              <img
                src="/src/assets/img/Artikel.png"
                alt="Latest Image"
                class="img-fluid"
              />
            </div>

            <div class="col-md-6 order-md-1">
              <div class="text-center text-md-start">
                <h2 class="mb-3 d-md-none warna-judul-artikel">
                  {{ contentkeempat.judul }}
                </h2>
                <div class="d-none d-md-block">
                  <h2 class="mb-2 warna-judul-artikel">
                    {{ contentkeempat.judul }}
                  </h2>
                  <br />
                </div>
              </div>
              <h5 class="mb-4 sub-berita">{{ contentkeempat.subtitle }}</h5>
              <h5 class="mb-4 item-berita">{{ contentkeempat.tanggal }}</h5>
              <p class="mb-4 sub-berita">{{ contentkeempat.teks }}</p>
              <a href="/berita">
                <button class="btn btn-secondary" @click="showDetail(item)">
                  Selengkapnya
                </button>
              </a>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>

  <div class="container-lg mt-5">
    <h2 class="text-idm mb-4">Galeri Foto</h2>
    <p class="text-sidm mb-4">Lihat galeri foto desa</p>
    <a href="/galeri-foto">
      <p class="text-galeri mb-4">Lihat semua</p>
    </a>
    <div class="container mt-4">
      <div class="row justify-content-between">
        <div
          v-for="(contentkelima, index) in galeri"
          :key="index"
          class="col-4 mb-4 ml-2 border rounded with-shadow"
          :class="{ 'with-shadow': isHovered }"
          @mouseenter="addShadow"
          @mouseleave="removeShadow"
        >
          <div class="mb-4 mt-2 gambar-orang">
            <img
              v-bind:src="contentkelima.imageUrl"
              :alt="contentkelima.name"
              class="img-fluid"
            />
          </div>
          <p class="mb-2 mt-2 text-center galeri-description">
            {{ contentkelima.nama }}
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      imageUrl: "",
      title: "",
      subtitle: "",
      isHovered: false,
      currentHost: "" 
    };
  },
  methods: {
    CekDesa(){
      this.currentHost = window.location.host;
      console.log("Ini adalah nilainya :", this.currentHost);

      axios.post('http://localhost:7071/link_url', {
        cari_link_url: this.currentHost
      })
      .then(response => {
        if (response.data.status === true) {
          const IDDESA = response.data.id_desa;
          localStorage.setItem('id_desa', IDDESA);
        }
      })
      .catch(error => {
        console.error('Terjadi Kesalahan saat Cek Ketersediaan ID Desa:', error);
        alert('Desa Tidak Ditemukan !');
      });
    },

    fetchData() {
      const apiResponse = {
        data: {
          imageUrl: "src/assets/img/desa.png",
          title: "Kenali Lebih dalam Desa Kami!",
          subtitle: "Desa yang berada di Kab Morowali ",
          penjelasan: [
            { judul: 'Desa Bahomoleo', teks: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis sodales consequat dictum. Praesent fermentum blandit ipsum et ultricies. Nunc ultrices neque ac velit aliquet, in iaculis nisi pellentesque.' }
          ],
          idm: [
            { imageUrl: 'src/assets/img/Idm.png', teks: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer nec dui at dui tempor convallis id id lacus. In quam sem, euismod in elit sit amet, aliquet mollis magna. Etiam a viverra tortor. In malesuada faucibus mi. In laoreet sapien vitae felis suscipit, quis dignissim elit vestibulum. Sed convallis posuere metus eget volutpat. Duis odio mi, pellentesque et erat in, gravida malesuada nibh. Donec in fringilla orci, ut dapibus justo. Integer elementum non dui id convallis. Suspendisse elementum lectus non ullamcorper laoreet. In consequat dapibus sapien eu volutpat.' }
          ],
          struktur: [
            { imageUrl: 'src/assets/img/kepdes1.png', name: 'Nama Kepala Desa 1', periode: 'Tahun 20XX-20YY' },
            { imageUrl: 'src/assets/img/kepdes1.png', name: 'Nama Kepala Desa 2', periode: 'Tahun 20YY-20ZZ' },
            { imageUrl: 'src/assets/img/kepdes1.png', name: 'Nama Kepala Desa 2', periode: 'Tahun 20YY-20ZZ' },
          ],
          berita: [
            { judul: 'Warga Morowali Curhat Rumah Rusak-Anak Ispa gegara Debu Batu Bara PT IMIP', subtitle: 'Deskripsi tambahan untuk berita', teks: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis sodales consequat dictum. Praesent fermentum blandit ipsum et ultricies. Nunc ultrices neque ac velit aliquet, in iaculis nisi pellentesque.', tanggal: '08-09-2019' }
          ],
          galeri: [
            {
              imageUrl: "src/assets/img/Artikel.png",
              nama: "Pemancingan",
            },
            {
              imageUrl: "src/assets/img/Artikel.png",
              nama: "Pemancingan",
            },
            {
              imageUrl: "src/assets/img/Artikel.png",
              nama: "Pemancingan",
            },
          ]
        },
      };
      this.imageUrl = apiResponse.data.imageUrl;
      this.title = apiResponse.data.title;
      this.subtitle = apiResponse.data.subtitle;
      this.penjelasan = apiResponse.data.penjelasan;
      this.idm = apiResponse.data.idm;
      this.struktur = apiResponse.data.struktur;
      this.berita = apiResponse.data.berita;
      this.galeri = apiResponse.data.galeri;
    },
  },
  addShadow() {
    this.isHovered = true;
  },
  removeShadow() {
    this.isHovered = false;
  },
  mounted() {
    this.fetchData();
    this.CekDesa();
  }


}
</script>

<style scoped>
.gambar {
  text-align: center;
  margin: auto;
}

.text_tt {
  color: aliceblue;
  position: absolute;
  left: 20%;
  top: 40%;
  transform: translate(-40%, -50%);
  font-weight: bold;
}

.text_st {
  color: aliceblue;
  position: absolute;
  margin-left: 15%;
  top: 50%;
  transform: translate(-41%, -50%);
  font-weight: bold;
}

.btn1 {
  margin-left: 40%;
  display: grid;
  margin: auto;
}

.btn-secondary {
  background-color: #003366;
  border: none;
  padding: 10px 30px;
  font-family: "Poppins", sans-serif;
  font-weight: 600;
  font-size: 14px;
  cursor: pointer;
  border-radius: 10px;
}

.warna-judul-artikel {
  font-weight: bold;
  color: #003366 !important;
}

.with-shadow {
  transition: box-shadow 0.3s ease-in-out;
  box-shadow: none;
}

.with-shadow:hover,
.with-shadow.active {
  box-shadow: 0 0 10px 0 rgba(0, 51, 102, 0.5);
}

.text-idm {
  color: #003366;
  font-weight: bold;
  text-align: center;
}

.text-sidm {
  color: black;
  font-size: 20px;
  font-weight: bold;
  text-align: center;
}

.text-galeri {
  color: black;
  text-align: end;
  text-decoration: dashed;
}

.text-ssidm {
  text-align: center;
}

.nama-kepdes {
  text-align: center;
  font-size: 20px;
  font-weight: 600;
  color: black;
  font-family: "Poppins", sans-serif;
}

.periode-kepdes {
  text-align: center;
  font-size: 20px;
  font-weight: 400;
  color: black;
  font-family: "Poppins", sans-serif;
}

.sub-berita {
  font-weight: bold;
}

.item-berita {
  color: #003366;
  font-size: 18px;
  font-weight: bold;
}

.galeri-description {
  font-weight: bold;
}

@media only screen and (max-width: 768px) {
  /* For mobile phones: */
  [class*="col-"] {
    width: 100%;
  }
}

@media only screen and (max-width: 1200px) {
  /* For medium-sized screens: */
  .text_tt {
    left: 10%;
    top: 40%;
    transform: translate(-10%, -50%);
  }

  .text_st {
    margin-left: 10%;
    top: 50%;
    transform: translate(-11%, -50%);
  }
}

@media only screen and (max-width: 992px) {
  /* For small-sized screens: */
  .text_tt {
    left: 5%;
    top: 40%;
    transform: translate(-5%, -50%);
  }

  .text_st {
    margin-left: 5%;
    top: 50%;
    transform: translate(-6%, -50%);
  }
}

@media only screen and (max-width: 768px) {
  /* For mobile phones: */
  .text_tt {
    position: absolute;
    left: 40%;
    font-size: 27px;
    top: 11%;
    transform: translate(-50%, 0);
    text-align: center;
    margin: 0;
  }

  .text_st {
    position: absolute;
    font-size: 15px;
    left: 40%;
    top: 19%;
    transform: translate(-50%, 0);
    text-align: center;
    margin: 0;
  }
}
</style>
