<template>
  <div class="gambar">
    <img src="../assets/img/desa.png" alt="Gambar Berita" class="img-fluid absolute-image" width="2000" />
    <h1 class="text_tt text-start">{{ title }}</h1>
    <h4 class="text_st text-start">{{ subtitle }}</h4>
  </div>

  <div v-for="(idm, index) in idmData" :key="index" class="container-lg wrapper-idm">
    <h2 class="text-idm mb-4">Index Desa Membangun</h2>
    <p class="text-sidm mb-4">Lihat total index desa</p>
    <img src="../assets/img/Idm.png" alt="Gambar Berita" class="img-fluid" width="2000" />
    <p class="text-ssidm mt-3">{{ idm.teks }}</p>
    <a href="/idm">
      <button class="btn1 btn btn-secondary mt-4" @click="showDetail(item)">
        Selengkapnya
      </button>
    </a>
  </div>

  <div class="container-lg mt-5 py-5 tengahin">
    <h2 class="text-idm mb-4">Struktur Organisasi</h2>
    <p class="text-sidm mb-4">Lihat total struktur organisasi</p>
    <div class="row">
      <div v-for="(kepala_desa, id_kepala_desa) in kepalaDesaData" :key="id_kepala_desa" class="container-kepdes">
        <img :src="`data:image/png;base64,${kepala_desa.foto_kepala_desa}`" :alt="kepala_desa.name"
          class="img-fluid foto-kepdes" />
        <div class="text-center mt-2">
          <p class="font-weight-bold nama-kepdes">{{ kepala_desa.nama }}</p>
          <p class="periode-kepdes">
            {{ kepala_desa.jabatan_dimulai }} -
            {{ kepala_desa.jabatan_berakhir }}
          </p>
        </div>
      </div>
    </div>
    <a href="/sejarah">
      <button class="btn1 btn btn-secondary mt-2" @click="showDetail(item)">
        Selengkapnya
      </button>
    </a>
  </div>

  <div class="container-lg mt-5 paddingx-disini">
    <h2 class="text-idm mb-4">Arsip Berita dan Artikel</h2>
    <p class="text-sidm mb-4">Berita terbaru dari desa Bahomoleo</p>
    
    <div class="berita-container">
      <div v-for="(berita, id_berita) in beritaData" :key="id_berita" class="card-berita with-shadow"
        :class="{ 'with-shadow': isHovered }" @mouseenter="addShadow" @mouseleave="removeShadow">
        <div class="card-kiri">
          <img :src="`data:image/png;base64,${berita.foto_berita}`" alt="Latest Image" class="img-fluid foto-berita" />
        </div>

        <div class="card-kanan">
          <div class="align-ini">
            <div>
              <h2 class="mb-2 warna-judul-artikel">
                {{ berita.judul }}
              </h2>
            </div>
            <h5 class="sub-berita">{{ berita.sub_judul }}</h5>
            <p class="mb-4 sub-berita" v-html="berita.deskripsi" />
          </div>
          <a href="/berita" class="anchor-berita">
            <button class="btn btn-secondary button-berita" @click="showDetail(berita)">
              Selengkapnya
            </button>
          </a>
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
    <div class="container mt-">
      <div class="container-galeri">
        <div v-for="(galeri, index) in galeriData" :key="index"
          class="col-4 mb-4 ml-2 border rounded min-lebar with-shadow" :class="{ 'with-shadow': isHovered }"
          @mouseenter="addShadow" @mouseleave="removeShadow">
          <div class="mb-4 gambar-orang">
            <img :src="`data:image/png;base64,${galeri.nama}`" :alt="galeri.foto" class="img-fluid foto-galeri" />
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
        const response = await axios.get(
          `http://localhost:8080/potensi_desa/list/${this.id_desa}`,
          {}
        );
        if (response.data.status) {
          this.potensiData = response.data.data
            .sort((a, b) => b.id_potensi - a.id_potensi)
            .slice(0, 1);
        } else {
          console.log("Data Kosong atau Terjadi Kesalahan");
        }
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    },
    async fetchIdm() {
      try {
        const response = await axios.get(
          `http://localhost:8080/idm/${this.id_desa}`
        );
        if (response.data.status) {
          this.idmData = response.data.data;
        } else {
          console.log("Data Kosong atau Terjadi Kesalahan");
        }
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    },
    async fetchKepalaDesa() {
      try {
        const response = await axios.get(
          `http://localhost:8080/pemerintah/kepdes/${this.id_desa}`
        );
        if (response.data.status) {
          this.kepalaDesaData = response.data.data
            .map((item, index) => ({ ...item, index }))
            .sort((a, b) => b.index - a.index)
            .slice(0, 3);
        } else {
          console.log("Data Kosong atau Terjadi Kesalahan");
        }
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    },
    async fetchBerita() {
      try {
        const response = await axios.post("http://localhost:8080/berita/list", {
          id_desa: this.id_desa,
        });
        if (response.data.status) {
          this.beritaData = response.data.data
            .sort((a, b) => b.id_berita - a.id_berita)
            .slice(0, 2);
        } else {
          console.log("Data Kosong atau Terjadi Kesalahan");
        }
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    },
    async fetchGaleri() {
      try {
        const response = await axios.get(
          `http://localhost:8080/galeri/${this.id_desa}`
        );
        if (response.data.status) {
          this.galeriData = response.data.data
            .map((item, index) => ({ ...item, index }))
            .sort((a, b) => b.index - a.index)
            .slice(0, 3);
        } else {
          console.log("Data Kosong atau Terjadi Kesalahan");
        }
      } catch (error) {
        console.error("Error fetching data:", error);
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
.absolute-image {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  object-fit: cover;
  object-position: center;
  z-index: 0;
}

.gambar {
  display: flex;
  flex-direction: column;
  text-align: start;
  justify-content: center;
  margin: auto;
  position: relative;
  min-height: 200px;
  height: 400px;
  width: 100%;
}

.text_tt {
  z-index: 1;
  color: aliceblue;
  font-weight: bold;
  margin-left: 20px;
}

.text_st {
  z-index: 1;
  color: aliceblue;
  font-weight: bold;
  margin-left: 20px;
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
  text-align: end;
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

.row {
  width: 100%;
  display: flex;
  flex-direction: row;
  justify-content: space-evenly;
}

.container-kepdes {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  width: fit-content;
}

.foto-kepdes {
  width: 300px;
  height: 400px;
}

.berita-container {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 40px;
}

.card-berita {
  display: flex;
  flex-direction: row;
  justify-content: center;
  gap: 10px;
  border: solid;
  border-width: 1px;
  border-color: #c4c4c4;
  border-radius: 10px;
  min-height: 300px;
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  padding-left: 40px;
  padding-right: 40px;
}

.wrapper-atas {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

.tengahin {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

.foto-galeri {
  width: 100%;
  height: 300px;
  object-fit: cover;
}

.min-lebar {
  min-width: 200px;
}

.container-galeri {
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  gap: 20px;
}

.card-kanan {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: end;
  gap: 20px;
  width: 100%;
}

.card-kiri {
  border-radius: 0%;
}

.align-ini {
  text-align: end;
}

.wrapper-idm {
  margin-top: 60px;
}

@media (max-width: 425px) {
  /* For small-sized screens: */

  .berita-container {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    gap: 50px;
  }

  .text_tt {
    z-index: 1;
    color: aliceblue;
    font-weight: bold;
    margin-left: 20px;
    font-size: 1rem;
  }

  .text_st {
    z-index: 1;
    color: aliceblue;
    font-weight: bold;
    margin-left: 20px;
    font-size: 1rem;
  }

  .text-ssidm {
    text-align: center;
    font-size: small;
  }

  .paddingx-disini {
    padding-left: 20px;
    padding-right: 20px;
  }

  .foto-berita {
    width: 100%;
    height: 300px;
    object-fit: cover;
    border-radius: 5px;
    margin-bottom: 8px;
  }

  .warna-judul-artikel {
    font-weight: bold;
    font-size: large;
    color: #003366 !important;
    text-align: start;
  }

  .sub-berita {
    font-weight: bold;
    font-size: small;
  }

  .container-galeri {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    gap: 10px;
  }

  .text-galeri {
    color: black;
    text-align: center;
    text-decoration: dashed;
  }

  .card-berita {
    border: solid;
    border-width: 1px;
    border-color: #c4c4c4;
    border-radius: 10px;
    min-height: 300px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    gap: 20px;
    padding-left: 0px;
    padding-right: 0px;
  }

  .button-berita {
    max-width: fit-content;
    padding: 10px;
  }

  .anchor-berita {
    display: flex;
    justify-content: end;
  }

  .card-kanan {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    align-items: start;
    gap: 20px;
    width: 100%;
    height: 100%;
    padding: 20px;
  }

  .align-ini {
    text-align: start;
  }
}
</style>
