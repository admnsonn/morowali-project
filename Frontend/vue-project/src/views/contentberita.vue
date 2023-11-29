<template>
    <div class="container mt-5">
        <h1 class="text-black mb-4 judul-text">ssss{{ item.number }}</h1>
    </div>
</template>

<script>

export default {
    data() {
        return {

            methods: {
                fetchData() {
                    // Simulasi pengambilan data dari API (dalam hal ini, data respons API)
                    const apiResponse = {
                        data: [
                            {
                                number: 123.123,
                                description: "Hektare Wilayah",
                            },
                            {
                                number: 56.565,
                                description: "Jumlah Warga Desa",
                            },
                            {
                                number: 89.789,
                                description: "Luas Perkebunan",
                            },
                            {
                                number: 123.123,
                                description: "Hektare Wilayah",
                            },
                            {
                                number: 56.565,
                                description: "Jumlah Warga Desa",
                            },
                            {
                                number: 89.789,
                                description: "Luas Perkebunan",
                            },
                            {
                                number: 123.123,
                                description: "Hektare Wilayah",
                            },
                            {
                                number: 56.565,
                                description: "Jumlah Warga Desa",
                            },
                            {
                                number: 89.789,
                                description: "Luas Perkebunan",
                            },
                        ],
                    }
                },
            }
        }
    }
}
</script>