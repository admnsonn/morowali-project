<template>
    <div class="centered-form">
      <div class="form-container">
        <form class="form" @submit.prevent="submitForm">
          <p class="form-title">Login</p>
          <div class="input-container">
            <input type="text" placeholder="Enter Username" v-model="formData.username" />
          </div>
          <div class="input-container">
            <input type="password" placeholder="Enter Password" autocomplete="new-password" v-model="formData.password" />
          </div>
          <button type="submit" class="submit">Login</button>
        </form>
      </div>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
 
  
  export default {
    data() {
      return {
        formData: {
          username: '',
          password: ''
        }
      };
    },

    

    methods: {
      submitForm() {
        axios.post('http://localhost:7071/login', {
          username: this.formData.username,
          password: this.formData.password
        })
        .then(response => {
          if (response.data.status === true && response.data.data.length > 0) {
            const userData = response.data.data[0];
            if (userData.token) {
              // Simpan token ke local storage atau cookie
              this.setToken(userData.token);
              this.setRole_id(userData.desa_id);
              this.setDesa_id(userData.role_id);
              this.setRole_pengguna(userData.role_pengguna);
              this.setID_pengguna(userData.id_pengguna);
              localStorage.setItem('isLoggedIn', 'true')
  
              // Redirect ke halaman setelah login

              window.location.reload()
            } else {
              alert('Token not found in response.'); // Pesan jika token tidak ada di respons
            }
          } else {
            alert('Login failed. Invalid credentials.'); // Pesan jika login gagal
          }
        })
        .catch(error => {
          console.error('Error occurred while logging in:', error);
          alert('Login failed! Please try again.'); // Pesan jika terjadi kesalahan saat login
        });
      },
      setToken(token) {
        localStorage.setItem('token', token);
      },
      setRole_id(role_id) {
        localStorage.setItem('role_id', role_id);
      },
      setDesa_id(desa_id) {
        localStorage.setItem('desa_id', desa_id);
      },
      setRole_pengguna(role_pengguna) {
        localStorage.setItem('role_pengguna', role_pengguna);
      },
      setID_pengguna(id_pengguna) {
        localStorage.setItem('id_pengguna', id_pengguna);
      },
    }
  };
  </script>
  
<style scoped>
.form {
    background-color: #fff;
    display: block;
    padding: 1rem;
    max-width: 350px;
    border-radius: 0.5rem;
    box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1),
        0 4px 6px -2px rgba(0, 0, 0, 0.05);
}

.template {
    /* padding: 3%; */
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    /* Memastikan form berada di tengah vertikal di seluruh layar */
}

.centered-form {
    text-align: center;
    margin: 6%;
}

.form-container {
    display: inline-block;
    background-color: #fff;
    /* padding: 1rem; */
    max-width: 350px;
    border-radius: 0.5rem;
    box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1),
        0 4px 6px -2px rgba(0, 0, 0, 0.05);
}

.form-title {
    font-size: 1.25rem;
    line-height: 1.75rem;
    font-weight: 600;
    text-align: center;
    color: #000;
}

.input-container {
    position: relative;
}

.input-container input,
.form button {
    outline: none;
    border: 1px solid #e5e7eb;
    margin: 8px 0;
}

.input-container input {
    background-color: #fff;
    padding: 1rem;
    padding-right: 3rem;
    font-size: 0.875rem;
    line-height: 1.25rem;
    width: 300px;
    border-radius: 0.5rem;
    box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.05);
}

.submit {
    display: block;
    padding-top: 0.75rem;
    padding-bottom: 0.75rem;
    padding-left: 1.25rem;
    padding-right: 1.25rem;
    background-color: #003366;
    color: #ffffff;
    font-size: 0.875rem;
    line-height: 1.25rem;
    font-weight: 500;
    width: 100%;
    border-radius: 0.5rem;
    text-transform: uppercase;
}

.signup-link {
    color: #6b7280;
    font-size: 0.875rem;
    line-height: 1.25rem;
    text-align: center;
}

.signup-link a {
    text-decoration: underline;
}
</style>
