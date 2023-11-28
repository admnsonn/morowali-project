<template>
    <footer className="footer bg-transparent mt-5">
      <div className="container">
        <div className="row">
          <div className="col-lg-4 mb-4">
            <img src="../assets/logo.png" alt="Logo Desa" height="50" className="me-3" />
            <div>
              <span className="fw-bold nama-desa text-white">Desa Bahomoleo</span>
              <span className="sub-text text-white-50">Kab. Morowali</span>
            </div>
          </div>
          <div className="col-lg-8 mb-4">
            <ul className="nav justify-content-end">
              <li className="nav-item">
                <h5 className="nav-link text-white tulisan">Alamat Kantor</h5>
              </li>
              <li className="nav-item">
                <h5 className="nav-link text-white tulisan">Profil Desa</h5>
              </li>
              <!-- Tambahkan menu lainnya sesuai kebutuhan -->
            </ul>
          </div>
        </div>
      </div>
    </footer>
  </template>
  
  <script>
  export default {
    // Tambahkan logika atau data yang diperlukan di sini jika diperlukan
  };
  </script>
  
  <style scoped>
  /* Gaya untuk footer */
  .footer {
    padding: 50px 0;
    background-color: transparent; /* Sesuaikan dengan preferensi warna latar belakang Anda */
    box-shadow: 0 -2px 4px rgba(0, 0, 0, 0.1); /* Ubah nilai sesuai preferensi Anda */
  }
  
  /* Gaya untuk elemen-elemen dalam footer */
  .nama-desa {
    font-weight: bold;
  }
  
  .sub-text {
    font-size: 0.7em;
    opacity: 0.7;
    display: block;
    margin-top: -10px;
  }
  
  .jarak {
    margin-right: 400px; /* Jika diperlukan */
  }
  
  .tulisan {
    font-weight: bold;
  }
  
  .text-white {
    color: #fff !important;
  }
  
  .text-white-50 {
    color: rgba(255, 255, 255, 0.5) !important;
  }
  </style>
  