<template>
    <div class="p-5 bg-utama">
        <h1 class="text-title mb-4 mt-0 pb-10">Pemerintahan Desa Bahomoleo</h1>
        <div class="container-lg mt-5">
            <div class="row">
                <div class="col-md-6">
                    <img :src="imageUrl" alt="Desa Bahomoleo" class="rounded mx-auto d-block custom-ukuran" />
                </div>
                <div class="col-md-6">
                    <p v-for="(paragraph, index) in paragraphs" :key="index" class="penjelasan-title">{{ paragraph }}</p>
                </div>
            </div>
        </div>
    </div>

    <div class="container-lg mt-5">
        <h4 class="text-visi mb-4">Visi</h4>
        <p v-for="(visis, index) in visi_misi" :key="index" class="penjelasan-vismis">{{ visis }}</p>
    </div>
    <br>
    <div>
        <h4 class="text-visi mb-4">Misi</h4>
        <p v-for="(visis, index) in visi_misi" :key="index" class="penjelasan-vismis">{{ visis }}</p>
    </div>

    <br>
    <div class="mt-5 p-5 bg-kedua">
        <h2 class="text-subtitle mb-4 mt-0 pb-10">Program Kerja</h2>
        <div class="container-lg mt-5">
            <p v-for="(paragraph, index) in paragraphs" :key="index" class="penjelasan-subtitle">{{ paragraph }}</p>
        </div>
    </div>

    <div class="mt-5 p-5">
        <h2 class="text-subtitle mb-4 mt-0 pb-10">Peraturan Kerja</h2>
        <div class="container-lg mt-5">
            <p v-for="(paragraph, index) in paragraphs" :key="index" class="penjelasan-subtitle">{{ paragraph }}</p>
        </div>
    </div>

    <div class="container py-5">
        <div class="border rounded row align-items-center pt-3 pb-3 mb-4 with-shadow" :class="{ 'with-shadow': isHovered }"
            @mouseenter="addShadow" @mouseleave="removeShadow">
            <div class="row">
                <div class="col-md-12 order-md-1">
                    <div class="text-center text-md-start">
                        <h2 class="mb-3 d-md-none judul-artikel">
                            {{ title_artikel }}
                        </h2>
                        <div class="d-none d-md-block">
                            <h2 class="mb-2 judul-artikel">
                                {{ title_artikel }}
                            </h2>
                        </div>
                        <br>
                        <h5 class="mb-4 subjudul-artikel">{{ subtitle_artikel }}</h5>
                        <a href="/lembaga">
                            <button class="btn btn-secondary" @click="showDetail(item)">
                                Selengkapnya
                            </button>
                        </a>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            isHovered: false,
            imageUrl: "",
            paragraphs: "",
            visi_misi: "",
            title_artikel: "",
            subtitle_artikel: "",
        };
    },
    methods: {
        addShadow() {
            this.isHovered = true;
        },
        removeShadow() {
            this.isHovered = false;
        },

        fetchData() {
            const apiResponse = {
                data: {
                    imageUrl: "src/assets/kepdes1.png",
                    paragraphs: [
                        "Lorem ipsum dolor sit amet, consectetur adipiscing elit. In felis urna, vehicula ut felis non, venenatis lacinia urna. Nullam eget nibh ante. Phasellus sed mi id turpis luctus congue ut sed neque. Pellentesque quis pretium justo, convallis imperdiet felis. Proin porttitor tristique varius.",
                        "Integer non convallis dui, ac cursus neque. Curabitur facilisis scelerisque lacus quis bibendum. Mauris placerat dolor sem, sit amet vestibulum magna venenatis eu. Aliquam consectetur metus sapien, in gravida enim elementum id. Nulla et purus magna. Vestibulum porta vitae ipsum nec facilisis. Phasellus efficitur luctus enim vitae condimentum. Curabitur commodo purus ac mi luctus, vitae lobortis ante semper. Duis quam orci, rhoncus vitae ullamcorper quis, tincidunt non erat. Maecenas dui ante, eleifend gravida sagittis quis, congue vestibulum enim.",
                    ],
                    visi_misi: ["Lorem ipsum dolor sit amet, consectetur adipiscing elit. In felis urna, vehicula ut felis non, venenatis lacinia urna."],
                    title_artikel: "Lembaga Pemerintahan Desa Bahomoleo",
                    subtitle_artikel: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis sodales consequat dictum.",
                },
            };
            this.imageUrl = apiResponse.data.imageUrl;
            this.paragraphs = apiResponse.data.paragraphs;
            this.visi_misi = apiResponse.data.visi_misi;
            this.title_artikel = apiResponse.data.title_artikel;
            this.subtitle_artikel = apiResponse.data.subtitle_artikel;
        },
    },
    mounted() {
        this.fetchData();
    },
};
</script>

<style scoped>
.container {
    padding-left: 20%;
    padding-right: 20%;
}

.bg-utama {
    background-image: linear-gradient(#003366, white);
}

.text-title {
    color: #fff;
    font-weight: bold;
    text-align: center;
}

.custom-ukuran {
    width: 380px;
    max-width: 100%;
    display: block;
    height: auto;
}

.penjelasan-title {
    font-weight: 500;
    color: black;
    font-family: 'Poppins', sans-serif;
}

.text-visi {
    color: #003366;
    font-weight: bold;
    text-align: center;
}

.penjelasan-vismis {
    text-align: center;
}

.text-subtitle {
    color: #003366;
    font-weight: bold;
    text-align: center;
}

.bg-kedua {
    background-color: #D1D9E0;
}

.penjelasan-subtitle {
    text-align: center;
}

.judul-artikel {
    font-weight: 600;
    color: #003366 !important;
}

.subjudul-artikel {
    font-size: 100%;
}

.btn-secondary {
    background-color: #003366;
    border: none;
    padding: 10px 30px;
    font-family: 'Poppins', sans-serif;
    font-weight: 600;
    font-size: 14px;
    cursor: pointer;
    border-radius: 10px;
}
</style>