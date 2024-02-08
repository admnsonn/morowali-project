<template>
    <div class="kontainer-admin">
        <div class="kontainer">
            <div class="bartipis" />
            <div class="">
                <h1 class="teks-admin">Admin desa Bahomoleo</h1>
                <p class="teks-kabupaten">Kabupaten, Morowali</p>
            </div>
        </div>
    </div>

    <div class="container">
        <div class="row">
            <div class="col">
                <h3 class="title-warga">Tambah Data Berita</h3>
                <p class="subtitle-warga">Management Content dan Layanan Berita</p>
            </div>
        </div>

        <div class="content-warga">
            <p class="header-title">Form Input</p>
        </div>
        <div class="isi-tambahdata">
            <div class="grid-container">

                <div class="field1">
                    <div class="form-group">
                        <label for="formFile" class="form-label">Foto Wisata</label>
                        <input class="form-control" v-on:change="onFileChange" type="file" id="formFile" accept="image/*" />
                    </div>
                    <div class="form-group-foto" v-if="this.model.wisata.foto_wisata">
                        <label for="foto">Preview foto</label>
                        <img :src="`data:image/png;base64,${this.model.wisata.foto_wisata}`" alt="foto wisata" height="300"
                            width="400" class="td-foto" />
                    </div>
                    <div v-else>Tidak ada gambar yang dipilih</div>
                </div>

                <div class="field2">
                    <div class="form-group">
                        <label for="Wisata">Nama Wisata</label>
                        <input type="text" v-model="model.wisata.nama_wisata" class="form-control" id="Wisata"
                            aria-label="Wisata" placeholder="Nama Wisata" />
                    </div>
                </div>

                <div class="field3">
                    <div class="form-group">
                        <label for="Konten">Konten Wisata</label>
                        <QuillEditor toolbar="essential" v-model="model.wisata.konten_wisata" theme="snow"
                            content-type="html" />
                    </div>
                </div>

                <div class="field4">
                    <div class="form-group">
                        <label for="Alamat">Alamat</label>
                        <input type="text" v-model="model.wisata.alamat" class="form-control" id="Alamat"
                            aria-label="alamat" placeholder="alamat" />
                    </div>
                </div>

                <div class="field5">
                    <div class="form-group">
                        <label for="Telp">No Telp</label>
                        <input type="text" v-model="model.wisata.no_telp" class="form-control" id="Telp" aria-label="telp"
                            placeholder="No Telp" />
                    </div>
                </div>

                <div class="field6">
                    <div class="form-group">
                        <label for="Lokasi">Link Lokasi</label>
                        <input type="text" v-model="model.wisata.link_lokasi" class="form-control" id="Lokasi"
                            aria-label="lokasi" placeholder="Link Lokasi" />
                    </div>
                </div>

                <div class="field10">
                    <div class="form-group">
                        <label for="formFile" class="form-label">Kategori</label>
                        <br />
                        <select ref="kategoriSelect" v-model="model.wisata.idkategoriwisata" class="form-control"
                            id="kategori_id" aria-label="category">
                            <option value="" disabled selected>--Pilih Kategori--</option>
                            <option v-for="item in kategoriList" :value="item.id_wisata_kategori">
                                {{ item.wisata_kategori }}
                            </option>
                        </select>
                    </div>
                </div>

                <div class="field13">
                    <button type="button" class="btn btn-success btn-simpan p-2 my-2" @click="addNewData">
                        <div class="nav-link router-link-underline teks-tambah">
                            + Tambah Data
                        </div>
                    </button>
                </div>
            </div>
        </div>
    </div>
</template>
  
<script>
import axios from "axios";
import Swal from "sweetalert2";
import { QuillEditor } from "@vueup/vue-quill";
import "@vueup/vue-quill/dist/vue-quill.snow.css";

export default {
    components: {
        QuillEditor,
    },
    name: "beritaCreate",
    data() {
        return {
            model: {
                wisata: {
                    nama_wisata: "",
                    konten_wisata: "",
                    kategori_wisata: "",
                    no_telp: "",
                    desa_id: 1,
                    link_lokasi: "",
                    alamat: "",
                    foto_wisata: "",
                    idkategoriwisata: 1,
                },
            },
            tableData: [],
            kategoriList: [],
        };
    },
    methods: {
        fetchKategori() {
            axios
                .get("http://localhost:8080/wisata/kategori_wisata")
                .then(({ data }) => {
                    this.kategoriList = data.data;
                })
                .catch((error) => {
                    console.error("Error in Axios GET request:", error);
                });
        },
        async addNewData() {
            const result = await Swal.fire({
                title: "Apakah anda yakin?",
                text: "Cek kembali data yang atkan ditambahkan!",
                icon: "warning",
                showCancelButton: true,
                confirmButtonColor: "#003366",
                cancelButtonColor: "#d33",
                confirmButtonText: "Ya, tambahkan!",
            });
            if (result.isConfirmed) {
                const plainTextDescription = this.model.wisata.konten_wisata;
                axios
                    .post("http://localhost:8080/wisata/tambah_data", {
                        nama_wisata: this.model.wisata.nama_wisata,
                        konten_wisata: plainTextDescription,
                        kategori_wisata: this.model.wisata.kategori_wisata,
                        no_telp: this.model.wisata.no_telp,
                        desa_id: this.model.wisata.desa_id,
                        link_lokasi: this.model.wisata.link_lokasi,
                        alamat: this.model.wisata.alamat,
                        idkategoriwisata: this.model.wisata.idkategoriwisata,
                        foto_wisata: this.model.wisata.foto_wisata,
                    })
                    .then((res) => {
                        this.model.berita = {
                            nama_wisata: "",
                            konten_wisata: "",
                            kategori_wisata: "",
                            no_telp: "",
                            desa_id: 1,
                            link_lokasi: "",
                            alamat: "",
                            foto_wisata: "",
                            idkategoriwisata: 1,
                        };
                        if (res.data.status) {
                            Swal.fire(
                                "Data berhasil ditambahkan.",
                                res.data.message,
                                "success"
                            );
                            this.$router.push('/wisata-desa');
                        } else {
                            Swal.fire("Data gagal ditambahkan.", res.data.message, "error");
                        }
                    })
                    .catch((error) => {
                        console.error(error);
                        // Handle error, e.g., show an error message
                    });
            }
        },
        onFileChange(e) {
            var files = e.target.files || e.dataTransfer.files;
            if (!files.length) return;
            const imageFile = files[0];
            const validTypes = ["image/jpeg", "image/png", "image/gif"]; // Adjust as needed
            if (!validTypes.includes(imageFile.type)) {
                // Display an error message or alert
                alert("Silakan unggah file gambar.");
                return;
            }

            const reader = new FileReader();
            reader.readAsDataURL(files[0]);

            reader.onload = (e) => {
                const imageUrl = e.target.result;
                const img = new Image();

                img.onload = () => {
                    const canvas = document.createElement("canvas");
                    const scaleX = 0.5; // Resize to 50%
                    const scaleY = 0.5;
                    const width = img.width * scaleX;
                    const height = img.height * scaleY;

                    canvas.width = width;
                    canvas.height = height;
                    const ctx = canvas.getContext("2d");
                    ctx.drawImage(img, 0, 0, width, height);

                    canvas.toBlob((resizedBlob) => {
                        const reader = new FileReader();
                        reader.readAsDataURL(resizedBlob);

                        reader.onloadend = () => {
                            const base64String = reader.result.split(",")[1]; // Extract base64-encoded data
                            this.model.wisata.foto_wisata = base64String;
                            // Now you can send this.model.berita to your API
                        };
                    }, "image/jpeg"); // Adjust the image type as needed
                };

                img.src = imageUrl;
            };
        },
    },
    created() {
        this.fetchKategori();
    },
    mounted() {
        this.$refs.kategoriSelect.focus();
    },
};
</script>
  
<style scoped>
h3 {
    font-weight: 700;
}

.bartipis {
    background-color: black;
    height: 100%;
    width: 3px;
    grid-row: span 2;
    top: 0;
    left: 0;
    align-self: center;
}

.container {
    margin-top: 30px;
    margin-bottom: 50px;
    width: calc(100% - 100px);
}

.title-warga {
    font-size: 20px;
}

.subtitle-warga {
    font-size: 15px;
    color: #5e5e5e;
}

.content-warga {
    background-color: #f7f7f7;
    padding-top: 20px;
    padding-bottom: 5px;
    padding-left: 20px;
    border-radius: 3px;
    border: 1px solid #c7c7c7;
}

.header-title {
    font-weight: bold;
    color: #000000;
}

.isi-tambahdata {
    background-color: #ffffff;
    padding-top: 20px;
    padding-bottom: 5px;
    padding-left: 20px;
    border-radius: 3px;
    border: 1px solid #c7c7c7;
}

.grid-container {
    display: grid;
    grid-template-columns: auto auto auto;
    grid-gap: 40px;
    padding: 10px;
    padding-right: 30px;
}

.btn-simpan {
    background-color: #003366;
    color: #fff;
    border: none;
}

.btn-simpan:hover {
    background-color: #003366;
    color: #fff;
    border: none;
}

.teks-tambah {
    font-size: 15px;
    padding-top: 2%;
    padding-bottom: 2%;
    padding-left: 5px;
    padding-right: 5px;
}

.form-group-foto {
    display: flex;
    flex-direction: column;
    margin-top: 10%;
}

.td-foto {
    border-radius: 0.375rem;
}
</style>
  