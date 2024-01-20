<template>
  <div class="wrapper-container">
    <div class="kontainer-admin">
      <h1>Beranda Admin</h1>
      <p>Selamat datang, Admin!</p>
    </div>
  </div>
</template>

<style scoped>
h1 {
  font-weight: bold !important;
  font-size: 100px;
}

p {
  font-weight: medium !important;
  font-size: 30px;
}
.wrapper-container {
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
}
.kontainer-admin {
  display: flex;
  flex-direction: column;
  background-color: aliceblue;
}
</style>
