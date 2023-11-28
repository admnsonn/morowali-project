
<template>
  <div classNameName="menu">
    <Navbar />
  </div>

  <div className="main">
    <RouterView />
  </div>

  <div className="kaki">
    <footerComponent />
  </div>
  
</template>

<script>
import Navbar from './components/Navbar.vue'
import footerComponent from './components/Footer.vue'

export default {
  components: {
    Navbar,
    footerComponent
  }
}
</script>

<style>
body {
  margin: 0;
}
</style>

<style scoped>
.menu {
  background-color: #003366;
  line-height: 2;
  padding: 10px;
}

.kaki {
  margin-bottom: 0;
  left: 0;
  width: 100%;
  background-color: #003366;
  line-height: 2;
  padding: 10px;
}

</style>