<template>
    <div class="p-5 bg-utama">
        <h1 class="text-title mb-4 mt-0 pb-10">Pemerintahan Desa Bahomoleo</h1>
        <div class="container-lg mt-5 sejarah-desa">
            <div class="row">
                <div class="col-md-6">
                    <img :src="imageUrl" alt="Desa Lamberea" class="rounded mx-auto d-block custom-ukuran" />
                </div>
                <div class="col-md-6">
                    <p v-for="(paragraph, index) in paragraphs" :key="index" class="sejarah-paragraph">{{ paragraph }}</p>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            imageUrl: "",
            paragraphs: "",
        };
    },
    methods: {
        fetchData() {
            const apiResponse = {
                data: {
                    imageUrl: "src/assets/kepdes1.png",
                    paragraphs: [
                        "Lorem ipsum dolor sit amet, consectetur adipiscing elit. In felis urna, vehicula ut felis non, venenatis lacinia urna. Nullam eget nibh ante. Phasellus sed mi id turpis luctus congue ut sed neque. Pellentesque quis pretium justo, convallis imperdiet felis. Proin porttitor tristique varius.",
                        "Integer non convallis dui, ac cursus neque. Curabitur facilisis scelerisque lacus quis bibendum. Mauris placerat dolor sem, sit amet vestibulum magna venenatis eu. Aliquam consectetur metus sapien, in gravida enim elementum id. Nulla et purus magna. Vestibulum porta vitae ipsum nec facilisis. Phasellus efficitur luctus enim vitae condimentum. Curabitur commodo purus ac mi luctus, vitae lobortis ante semper. Duis quam orci, rhoncus vitae ullamcorper quis, tincidunt non erat. Maecenas dui ante, eleifend gravida sagittis quis, congue vestibulum enim.",
                    ],
                },
            };
            this.imageUrl = apiResponse.data.imageUrl;
            this.paragraphs = apiResponse.data.paragraphs;
        },
    },
    mounted() {
        this.fetchData();
    },
};
</script>

<style>
.bg-utama {
    background-image: linear-gradient(#003366, white);

}

.text-title {
    color: #fff;
    text-align: center;
}

.custom-ukuran {
    position:relative;
    display: flex;
    width: 380px;
}
</style>