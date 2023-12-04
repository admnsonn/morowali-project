<template>
    <div class="p-5 text-center bg-hero mt-0 pb-10">
        <h1 class="mb-3 tulisan_judul">Galeri Foto</h1>
        <a class="btn btn-primary" role="button">Media</a>
    </div>

    <div class="container mt-4">
        <div class="row justify-content-between">
            <div v-for="(head, index) in galeri" :key="index" class="col-4 mb-4 ml-2 border rounded with-shadow "
                :class="{ 'with-shadow': isHovered }" @mouseenter="addShadow" @mouseleave="removeShadow">
                <div class="mb-4 mt-2 gambar-orang">
                    <img src="src/assets/img//Artikel.png" alt="Latest Image" class="img-fluid" />
                </div>
                <p class=" mb-2 mt-2 text-center galeri-description">{{ head.nama }}</p>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            galeri: [],
        };
    },
    methods: {
        addShadow() {
            this.isHovered = true;
        },
        removeShadow() {
            this.isHovered = false;
        },

        fetchdata() {
            const apiResponse = {
                data:
                    [
                        {
                            imageUrl: "src/assets/orang.png",
                            nama: "Pemancingan",
                        },
                        {
                            imageUrl: "src/assets/orang.png",
                            nama: "Pemancingan",
                        },
                        {
                            imageUrl: "src/assets/orang.png",
                            nama: "Pemancingan",
                        },
                        {
                            imageUrl: "src/assets/orang.png",
                            nama: "Pemancingan",
                        },
                        {
                            imageUrl: "src/assets/orang.png",
                            nama: "Pemancingan",
                        },
                        {
                            imageUrl: "src/assets/orang.png",
                            nama: "Pemancingan",
                        },
                    ]
            };
            this.galeri = apiResponse.data;
        },
    },
    mounted() {
        this.fetchdata();
    },
};
</script>

<style scoped>
.bg-hero {
    background-color: #003366;
}

.tulisan_judul {
    color: #fff;
    font-weight: bold;
    text-shadow: 2px 2px #252525;
}

.btn-primary {
    background-color: white;
    color: #003366;
    font-weight: bold;
    border: none;
    padding: 10px 30px;
    font-family: 'Poppins', sans-serif;
    font-weight: 700;
    font-size: 14px;
    cursor: pointer;
    border-radius: 30px;
}

.ml-2 {
    margin-left: -2%;
}

.galeri-description {
    font-weight: bold;
}

.with-shadow {
  transition: box-shadow 0.3s ease-in-out;
  box-shadow: none;
}

.with-shadow:hover,
.with-shadow.active {
  box-shadow: 0 0 10px 0 rgba(0, 51, 102, 0.5);
}
</style>