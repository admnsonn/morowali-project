<template>
    <div>
        <div class="container lg mt-3 content-desa">
            <h1 class="mb-3 mt-4 judul-berita">{{ title }}</h1>
            <img :src="imageurl" alt="Desa Bahomoleo" class="img-fluid" />
            <h2 class="mb-3 mt-4 subtitle-berita">{{ subtitle }}</h2>
            <div class="row">
                <div class="col-md-2">
                    <p class="kategori-berita">Kategori: {{ kategori }}</p>
                </div>
                <div class="col-md-2">
                    <p class="date-berita">Tanggal: {{ date }}</p>
                </div>
            </div>
            <p v-for="(paragraph, index) in paragraphs" :Key="index" class="penjelasan-berita">{{ paragraph }}</p>
        </div>
    </div>
</template>

<script>
export default {
    data() {
        return {
            title: "",
            imageurl: "",
            subtitle: "",
            kategori: "",
            date: "",
            paragraphs: "",
        };
    },
    methods: {
        fetchData() {
            const apiResponse = {
                data: {
                    title: "Warga Morowali Curhat Rumah Rusak-Anak Ispa gegara Debu Batu Bara PT IMIP",
                    imageurl: "src/assets/berita-terkini.png",
                    subtitle: "PENYERAHAN BANTUAN ALKES KE RS SARKIES AISYIAH KUDUS.",
                    kategori: "berita",
                    date: "08-09-2019",
                    paragraphs:
                        ["Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus leo diam, luctus vitae ipsum eget, tempor finibus nisi. Suspendisse in enim non metus convallis convallis ut ut lectus. Maecenas dui neque, pretium non magna id, pretium pellentesque libero. Curabitur ut sapien erat. Proin eu ligula odio. Vivamus convallis quis lacus sed egestas. Curabitur sit amet quam quis diam vehicula vestibulum ut ac nisl. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Cras hendrerit pulvinar leo, sed faucibus nisl tristique eu. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Etiam tristique cursus nisi at pulvinar. Integer faucibus lacus in lorem rutrum scelerisque vitae eu nunc. Aenean tempor vulputate lorem, eu fermentum tortor volutpat sed. Cras mollis tempor volutpat. Praesent vel fringilla tellus. Curabitur et dictum felis. Ut sed dolor accumsan nisl tristique tincidunt. Donec vitae efficitur tortor. Cras convallis metus sed nisi placerat, eu pellentesque turpis rhoncus. Nullam elementum dolor vel erat dignissim ornare. Sed porttitor, justo quis luctus aliquet, mauris urna venenatis dolor, a consequat mauris dolor vel diam.  Donec tincidunt lacus mattis risus ultricies elementum. Quisque feugiat dictum commodo. Suspendisse interdum tellus vitae libero aliquam feugiat. Morbi euismod sapien molestie mi efficitur, at volutpat purus ullamcorper."]
                },
            };
            this.title = apiResponse.data.title;
            this.imageurl = apiResponse.data.imageurl;
            this.subtitle = apiResponse.data.subtitle;
            this.kategori = apiResponse.data.kategori;
            this.date = apiResponse.data.date;
            this.paragraphs = apiResponse.data.paragraphs;
        },
    },
    mounted() {
        this.fetchData();
    },
};
</script>

<style scoped>
.judul-berita {
    font-weight: bold;
}

.subtitle-berita {
    font-weight: bold;
}

.kategori-berita {
    font-style: normal;
    color: #003366;
    font-weight: bold;
}

.date-berita {
    font-style: normal;
    color: #003366;
    font-weight: bold;
    margin-right: 20px;
}

.penjelasan-berita {
    display: flex;
    text-align: justify;
    font-weight: bold;
}
</style>