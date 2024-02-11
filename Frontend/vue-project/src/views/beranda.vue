<template>
  <div class="gambar">
    <img src="../assets/img/desa.png" alt="Gambar Berita" class="img-fluid gambar" width="2000" />
    <h1 class="text_tt text-start">{{ title }}</h1>
    <h4 class="text_st text-start">{{ subtitle }}</h4>
  </div>

  <div class="container py-5">
    <div v-for="(potensi, id_potensi) in potensiData" :key="id_potensi"
      class="border rounded row align-items-center pt-3 pb-3 mb-4 with-shadow" :class="{ 'with-shadow': isHovered }"
      @mouseenter="addShadow" @mouseleave="removeShadow">
      <div class="container mt-4">
        <div class="row">
          <div class="col-12 d-md-none mb-4">
            <img src="../../src/assets/img/Artikel.png" alt="Latest Image" class="img-fluid" />
          </div>

          <div class="col-md-6 order-md-2 d-none d-md-block">
            <img src="../../src/assets/img/Artikel.png" alt="Gambar Berita" class="img-fluid" />
          </div>

          <div class="col-md-6 order-md-1">
            <div class="text-center text-md-start">
              <h2 class="mb-3 d-md-none warna-judul-artikel">
                {{ potensi.judul_potensi }}
              </h2>
              <div class="d-none d-md-block">
                <h2 class="mb-2 warna-judul-artikel">
                  {{ potensi.judul_potensi }}
                </h2>
                <br />
              </div>
              <p class="mb-4 sub-berita">{{ potensi.deskripsi }}</p>
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

  <div v-for="(idm, index) in idmData" :key="index" class="container-lg mt-2">
    <h2 class="text-idm mb-4">Index Desa Membangun</h2>
    <p class="text-sidm mb-4">Lihat total index desa</p>
    <img src="../assets/img/Idm.png" alt="Gambar Berita" class="img-fluid gambar" width="2000" />
    <p class="text-ssidm mt-3">{{ idm.teks }}</p>
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
      <div v-for="(kepala_desa, id_kepala_desa) in kepalaDesaData" :key="id_kepala_desa" class="col-md-4 mb-4">
        <img :src="`data:image/png;base64,${kepala_desa.foto_kepala_desa}`" :alt="kepala_desa.name"
          class="img-fluid foto-kepdes" />
        <div class="text-center mt-2">
          <p class="font-weight-bold nama-kepdes">{{ kepala_desa.nama }}</p>
          <p class="periode-kepdes">{{ kepala_desa.jabatan_dimulai }} - {{ kepala_desa.jabatan_berakhir }}</p>
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
      <div v-for="(berita, id_berita) in beritaData" :key="id_berita"
        class="border rounded row align-items-center pt-3 pb-3 mb-4 with-shadow" :class="{ 'with-shadow': isHovered }"
        @mouseenter="addShadow" @mouseleave="removeShadow">
        <div class="container mt-4">
          <div class="row">
            <div class="col-md-6">
              <img src="/src/assets/img/Artikel.png" alt="Latest Image" class="img-fluid" />
            </div>

            <div class="col-md-6 order-md-1">
              <div class="text-center text-md-start">
                <h2 class="mb-3 d-md-none warna-judul-artikel">
                  {{ berita.judul }}
                </h2>
                <div class="d-none d-md-block">
                  <h2 class="mb-2 warna-judul-artikel">
                    {{ berita.judul }}
                  </h2>
                  <br />
                </div>
              </div>
              <h5 class="mb-4 sub-berita">{{ berita.sub_judul }}</h5>
              <p class="mb-4 sub-berita">{{ berita.deskripsi }}</p>
              <a href="/berita">
                <button class="btn btn-secondary" @click="showDetail(berita)">
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
        <div v-for="(galeri, index) in galeriData" :key="index" class="col-4 mb-4 ml-2 border rounded with-shadow"
          :class="{ 'with-shadow': isHovered }" @mouseenter="addShadow" @mouseleave="removeShadow">
          <div class="mb-4 mt-2 gambar-orang">
            <img :src="`data:image/png;base64,${galeri.nama}`" :alt="galeri.foto" class="img-fluid" />
          </div>
          <p class="mb-2 mt-2 text-center galeri-description">
            {{ galeri.foto }}
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      imageUrl: "",
      id_desa: "",
      beritaData: [],
      galeriData: [],
      kepalaDesaData: [],
      potensiData: [],
      idmData: [],
      // 
      title: "",
      subtitle: "",
      isHovered: false,
      currentHost: "",
    };
  },

  methods: {
    cekDesa() {
      this.currentHost = window.location.host;
      console.log("Ini adalah nilainya :", this.currentHost);

      axios
        .post("http://localhost:8080/link_url", {
          cari_link_url: this.currentHost,
        })
        .then((response) => {
          if (response.data.status === true) {
            const IDDESA = response.data.id_desa;
            localStorage.setItem("id_desa", IDDESA);
          }
        })
        .catch((error) => {
          console.error(
            "Terjadi Kesalahan saat Cek Ketersediaan ID Desa:",
            error
          );
          alert("Desa Tidak Ditemukan !");
        });
    },
    fetchData() {
      const apiResponse = {
        data: {
          imageUrl: "/../../src/assets/img/desa.png",
          title: "Kenali Lebih dalam Desa Kami!",
          subtitle: "Desa yang berada di Kab Morowali ",
          idmData: [
            {
              imageUrl: "/../../src/assets/img/Idm.png",
              teks: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer nec dui at dui tempor convallis id id lacus. In quam sem, euismod in elit sit amet, aliquet mollis magna. Etiam a viverra tortor. In malesuada faucibus mi. In laoreet sapien vitae felis suscipit, quis dignissim elit vestibulum. Sed convallis posuere metus eget volutpat. Duis odio mi, pellentesque et erat in, gravida malesuada nibh. Donec in fringilla orci, ut dapibus justo. Integer elementum non dui id convallis. Suspendisse elementum lectus non ullamcorper laoreet. In consequat dapibus sapien eu volutpat.",
            },
          ],
        },
      };
      this.imageUrl = apiResponse.data.imageUrl;
      this.title = apiResponse.data.title;
      this.subtitle = apiResponse.data.subtitle;
      this.idmData = apiResponse.data.idmData;
    },
    async fetchPotensi() {
      try {
        const response = await axios.get(`http://localhost:8080/potensi_desa/list/${this.id_desa}`, {});
        if (response.data.status) {
          this.potensiData = response.data.data
            .sort((a, b) => b.id_potensi - a.id_potensi)
            .slice(0, 1);
        } else {
          console.log("Data Kosong atau Terjadi Kesalahan")
        }
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    },
    async fetchIdm() {
      try {
        const response = await axios.get(`http://localhost:8080/idm/${this.id_desa}`);
        if (response.data.status) {
          this.idmData = response.data.data;
        } else {
          console.log("Data Kosong atau Terjadi Kesalahan")
        }
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    },
    async fetchKepalaDesa() {
      try {
        const response = await axios.get(`http://localhost:8080/pemerintah/kepdes/${this.id_desa}`);
        if (response.data.status) {
          this.kepalaDesaData = response.data.data
            .map((item, index) => ({ ...item, index }))
            .sort((a, b) => b.index - a.index)
            .slice(0, 3);
        } else {
          console.log("Data Kosong atau Terjadi Kesalahan")
        }
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    },
    async fetchBerita() {
      try {
        const response = await axios.post("http://localhost:8080/berita/list", { id_desa: this.id_desa });
        if (response.data.status) {
          this.beritaData = response.data.data
            .sort((a, b) => b.id_berita - a.id_berita)
            .slice(0, 2);
        } else {
          console.log("Data Kosong atau Terjadi Kesalahan")
        }
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    },
    async fetchGaleri() {
      try {
        const response = await axios.get(`http://localhost:8080/galeri/${this.id_desa}`);
        if (response.data.status) {
          this.galeriData = response.data.data
            .map((item, index) => ({ ...item, index }))
            .sort((a, b) => b.index - a.index)
            .slice(0, 3);
        } else {
          console.log("Data Kosong atau Terjadi Kesalahan")
        }
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    },
    addShadow() {
      this.isHovered = true;
    },
    removeShadow() {
      this.isHovered = false;
    },
  },

  created() {
    this.id_desa = localStorage.getItem("id_desa");
    this.cekDesa();
    this.fetchData();
    this.fetchPotensi();
    this.fetchIdm();
    this.fetchKepalaDesa();
    this.fetchBerita();
    this.fetchGaleri();
  },
};
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
