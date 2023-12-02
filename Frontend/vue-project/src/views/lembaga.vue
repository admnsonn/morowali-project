<template>
    <div class="p-5 text-center bg-hero mt-0 pb-10">
        <h1 class="mb-3 tulisan_judul">Lembaga Pemerintahan</h1>
        <a class="btn btn-primary" href="" role="button">Pemerintahan Daerah</a>
    </div>

    <div class="container-lg mt-5">
        <h2 class="text-sub">Regulasi terkait Lembaga Kemasyarakatan Desa (LKD) :</h2>
        <p v-for="(penjelasan, index) in paragraph" :key="index" class="text-penjelasan mt-4">{{ penjelasan }}</p>
    </div>
</template>

<script>
export default {
    data() {
        return {
            paragraph: "",
        };
    },
    methods: {
        fetchData() {
            const apiResponse = {
                data: {
                    paragraph:
                        ["Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla convallis dapibus ex, ut aliquam turpis scelerisque ut. Morbi non cursus elit. Maecenas at nunc congue, aliquet tellus vel, rutrum enim. In purus sem, cursus quis eros ac, lobortis tempus metus. Phasellus id ipsum cursus, sodales diam sollicitudin, consectetur tellus. Maecenas accumsan aliquet orci ut volutpat. Suspendisse dignissim rutrum elit aliquet suscipit. Nam arcu ante, consequat et magna nec, mollis auctor nibh. Duis dictum elit quis arcu varius sollicitudin. Nunc et placerat erat, at vulputate elit.",
                            "Donec tincidunt arcu vel consectetur lacinia. Cras sed dolor quam. Integer facilisis nisi et luctus dignissim. Curabitur at arcu sed leo faucibus gravida. Quisque pellentesque condimentum ipsum ut accumsan. Nulla quis maximus sapien, in vestibulum orci. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi pharetra dignissim ex. Curabitur eget viverra ex. Phasellus at augue ac enim porta pulvinar. Maecenas id orci in lectus eleifend interdum. Donec pharetra consectetur ante eu pharetra.",
                            "Etiam vel hendrerit magna. Nunc finibus eget mi nec pulvinar. Etiam at ligula convallis, rhoncus justo at, ultrices tortor. Aliquam porttitor quam ac massa tempor mattis. Phasellus ultrices vulputate rhoncus. Cras at justo sapien. Praesent gravida nibh tortor, et scelerisque urna sodales at. Duis mollis nulla id ex sodales tristique. Fusce ut efficitur neque, a mollis odio. Vestibulum scelerisque magna et eros feugiat ornare. Sed euismod vitae diam molestie faucibus. Proin facilisis mauris ac malesuada tincidunt. Maecenas nec tincidunt justo.",
                            "Phasellus vehicula laoreet mauris, at facilisis lacus sodales ut. Etiam ultricies mi lectus, sit amet semper sapien euismod in. Morbi mollis eget erat quis fringilla. Maecenas porta enim non ipsum molestie fringilla. Aliquam at tortor sit amet mauris lobortis scelerisque. Donec sodales congue ex, eu fringilla lacus. Sed condimentum sem quis augue euismod auctor. Maecenas nec pretium tellus. Vivamus molestie dapibus sapien, sed gravida nunc euismod eget. Cras ullamcorper vehicula turpis sed porta. Ut vehicula vehicula ultrices."]
                },
            };
            this.paragraph = apiResponse.data.paragraph;
        },
    },
    mounted() {
        this.fetchData();
    }
}
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

.text-sub {
    text-align: center;
    color: #003366;
    font-weight: bold;
}

.text-penjelasan {
    color: black;
}
</style>