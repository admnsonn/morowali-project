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
                <h3 class="title-warga">Ubah Data Visi-Misi</h3>
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
                        <label for="Peraturan">Peraturan</label>
                        <QuillEditor toolbar="essential" v-model:content="this.tableData[0].peraturan" theme="snow"
                            content-type="html" />
                    </div>
                </div>

                <div class="wrapper-button">
                    <button type="button" class="btn btn-success btn-simpan p-2 my-2 button-styling" @click="updateData">
                        <div class="nav-link router-link-underline teks-tambah">
                            > Ganti Data
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
            tableData: [],
            id_desa: "1",
        };
    },

    methods: {
        fetchData() {
            axios
                .get(`http://localhost:8080/pemerintah/peraturan/1`)
                .then(({ data }) => {
                    this.tableData = data.data;
                    console.log(this.tableData);
                    this.filteredData = this.tableData; // Initialize filteredData
                })
                .catch((error) => {
                    console.error("Error in Axios POST request:", error);
                });
        },
        async updateData() {
            const result = await Swal.fire({
                title: "Apakah anda yakin?",
                text: "Cek kembali data yang akan diubah!",
                icon: "warning",
                showCancelButton: true,
                confirmButtonColor: "#003366",
                cancelButtonColor: "#d33",
                confirmButtonText: "Ya, ubah!",
            });
            if (result.isConfirmed) {
                axios
                    .put("http://localhost:8080/pemerintah/peraturan/setting", {
                        peraturan: this.tableData[0].peraturan,
                        id_desa: this.id_desa,
                    })
                    .then((res) => {
                        if (res.data.status) {
                            Swal.fire("Data berhasil diubah.", res.data.message, "success");
                            this.$router.push('/pemerintahan/peraturan');
                        } else {
                            Swal.fire("Data gagal diubah.", res.data.message, "error");
                        }
                    })
                    .catch((error) => {
                        console.error(error);
                        Swal.fire({
                            title: "Error",
                            text: "Terjadi kesalahan saat mengubah data.",
                            icon: "error",
                        });
                    });
            }
        },
    },
    created() {
        this.fetchData(); // Get original data
        console.log(this.tableData[0]); //
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
}

.td-foto {
    border-radius: 0.375rem;
}

.wrapper-button {
    display: flex;
    justify-content: flex-end;
    flex-direction: column;
}
</style>
  