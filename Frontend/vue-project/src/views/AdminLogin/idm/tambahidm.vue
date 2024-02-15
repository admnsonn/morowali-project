<template>
  <div class="wrapper-h1">
    <ul class="list-group list-group-flush">
      <li class="list-group-item card-idm">
        <div>
          <h1>IDM</h1>
          <table class="tabel">
            <thead>
              <tr>
                <th>No.</th>
                <th>Tahun</th>
                <th>Indikator</th>
                <th>Skor</th>
                <th>Keterangan</th>
                <th>Kegiatan</th>
                <th>Nilai</th>
                <th>Pusat</th>
                <th>Provinsi</th>
                <th>Kabupaten</th>
                <th>Desa</th>
                <th>CSR</th>
                <th>Lainnya</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(row, index) in tableDataIDM" :key="index">
                <td>{{ index + 1 }}</td>
                <td>{{ row.tahun }}</td>
                <td>{{ row.indikator }}</td>
                <td>{{ row.skor }}</td>
                <td>{{ row.keterangan }}</td>
                <td>{{ row.kegiatan }}</td>
                <td>{{ row.nilai }}</td>
                <td>{{ row.pusat }}</td>
                <td>{{ row.provinsi }}</td>
                <td>{{ row.kabupaten }}</td>
                <td>{{ row.desa }}</td>
                <td>{{ row.csr }}</td>
                <td>{{ row.lainnya }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </li>
      <li class="list-group-item card-idm">
        <div>
          <h1>IKE</h1>
          <table class="tabel">
            <thead>
              <tr>
                <th>No.</th>
                <th>Tahun</th>
                <th>Indikator</th>
                <th>Skor</th>
                <th>Keterangan</th>
                <th>Kegiatan</th>
                <th>Nilai</th>
                <th>Pusat</th>
                <th>Provinsi</th>
                <th>Kabupaten</th>
                <th>Desa</th>
                <th>CSR</th>
                <th>Lainnya</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(row, index) in tableDataIKE" :key="index">
                <td>{{ index + 1 }}</td>
                <td>{{ row.tahun }}</td>
                <td>{{ row.indikator }}</td>
                <td>{{ row.skor }}</td>
                <td>{{ row.keterangan }}</td>
                <td>{{ row.kegiatan }}</td>
                <td>{{ row.nilai }}</td>
                <td>{{ row.pusat }}</td>
                <td>{{ row.provinsi }}</td>
                <td>{{ row.kabupaten }}</td>
                <td>{{ row.desa }}</td>
                <td>{{ row.csr }}</td>
                <td>{{ row.lainnya }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </li>
      <li class="list-group-item card-idm">
        <div>
          <h1>IKS</h1>
          <table class="tabel">
            <thead>
              <tr>
                <th>No.</th>
                <th>Tahun</th>
                <th>Indikator</th>
                <th>Skor</th>
                <th>Keterangan</th>
                <th>Kegiatan</th>
                <th>Nilai</th>
                <th>Pusat</th>
                <th>Provinsi</th>
                <th>Kabupaten</th>
                <th>Desa</th>
                <th>CSR</th>
                <th>Lainnya</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(row, index) in tableDataIKS" :key="index">
                <td>{{ index + 1 }}</td>
                <td>{{ row.tahun }}</td>
                <td>{{ row.indikator }}</td>
                <td>{{ row.skor }}</td>
                <td>{{ row.keterangan }}</td>
                <td>{{ row.kegiatan }}</td>
                <td>{{ row.nilai }}</td>
                <td>{{ row.pusat }}</td>
                <td>{{ row.provinsi }}</td>
                <td>{{ row.kabupaten }}</td>
                <td>{{ row.desa }}</td>
                <td>{{ row.csr }}</td>
                <td>{{ row.lainnya }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </li>
    </ul>

    <div class="excelbutton-wrapper">
      <div class="add-clear-buttons">
        <div class="wrapper-import-button">
          <label for="fileInput" class="btn btn-primary">
            Impor File Excel
          </label>
          <input
            type="file"
            id="fileInput"
            accept=".xlsx"
            style="display: none"
            :key="fileInputKey"
            :v-model="fileInputValue"
            @change="inputXLSX"
          />
        </div>
        <button @click="clearData" type="button" class="btn btn-danger">
          Clear data
        </button>
      </div>
      <div>
        <button @click="submitData" type="button" class="btn btn-success">
          Submit Data
        </button>
      </div>
    </div>
  </div>
</template>

<script>
import ExcelJS from "exceljs";
import Swal from "sweetalert2";
import axios from "axios";

export default {
  name: "exceljsTest",
  data() {
    return {
      tableDataIDM: [],
      tableDataIKE: [],
      tableDataIKS: [],
      desa_id: localStorage.getItem("desa_id"),
      fileInputKey: 0, // Initialize the key
      fileInputValue: {}, // Initially empty object
    };
  },
  methods: {
    clearData() {
      this.tableDataIDM = [];
      this.tableDataIKE = [];
      this.tableDataIKS = [];
      this.fileInputKey += 1; // Increment the key
      this.fileInputValue = {}; // Empty object to reset file selection
    },

    async submitData() {
      const result = await Swal.fire({
        title: "Apakah anda yakin?",
        text: "Cek kembali data yang atkan ditambahkan!",
        icon: "warning",
        showCancelButton: true,
        confirmButtonColor: "#003366",
        cancelButtonColor: "#d33",
        confirmButtonText: "Ya, tambahkan!",
      });
      if (result.isConfirmed) {
        const payload = {
          desa_id: this.desa_id,
          data: [
            {
              data_idm: this.tableDataIDM,
              data_iks: this.tableDataIKS,
              data_ike: this.tableDataIKE,
            },
          ],
        };
        console.log(payload);

        axios
          .post("http://localhost:8080/idmiksike/input", payload)
          .then((res) => {
            this.tableDataIDM = [];
            this.tableDataIKE = [];
            this.tableDataIKS = [];

            if (res.data.status) {
              Swal.fire(
                "Data berhasil ditambahkan.",
                res.data.message,
                "success"
              );
              this.$router.push("/idm-beranda");
            } else {
              Swal.fire("Data gagal ditambahkan.", res.data.message, "error");
            }
          })
          .catch((error) => {
            console.error(error);
            // Handle error, e.g., show an error message
          });
      } else {
        Swal.fire(
          "Data tidak ditambahkan.",
          "Data tidak jadi ditambahkan",
          "info"
        );
      }
    },

    async inputXLSX(event) {
      const selectedFile = event.target.files[0];
      if (!selectedFile) {
        return;
      }

      const workbook = new ExcelJS.Workbook();

      try {
        await workbook.xlsx.load(selectedFile);

        const worksheetIDM = workbook.getWorksheet("idm");
        const worksheetIKE = workbook.getWorksheet("ike");
        const worksheetIKS = workbook.getWorksheet("iks");

        const tableDataIDM = [];
        const tableDataIKE = [];
        const tableDataIKS = [];

        worksheetIDM.eachRow((row, index) => {
          if (index === 1) {
            return;
          }

          const data = {
            tahun: row.getCell(1).value.toString(),
            indikator: row.getCell(2).value,
            skor: row.getCell(3).value.toString(),
            keterangan: row.getCell(4).value,
            kegiatan: row.getCell(5).value,
            nilai: row.getCell(6).value,
            pusat: row.getCell(7).value,
            provinsi: row.getCell(8).value,
            kabupaten: row.getCell(9).value,
            desa: row.getCell(10).value,
            csr: row.getCell(11).value,
            lainnya: row.getCell(12).value,
          };
          tableDataIDM.push(data);
          console.log(tableDataIDM);
        });

        worksheetIKE.eachRow((row, index) => {
          if (index === 1) {
            return;
          }

          const data = {
            tahun: row.getCell(1).value.toString(),
            indikator: row.getCell(2).value,
            skor: row.getCell(3).value.toString(),
            keterangan: row.getCell(4).value,
            kegiatan: row.getCell(5).value,
            nilai: row.getCell(6).value,
            pusat: row.getCell(7).value,
            provinsi: row.getCell(8).value,
            kabupaten: row.getCell(9).value,
            desa: row.getCell(10).value,
            csr: row.getCell(11).value,
            lainnya: row.getCell(12).value,
          };
          tableDataIKE.push(data);
        });

        worksheetIKS.eachRow((row, index) => {
          if (index === 1) {
            return;
          }

          const data = {
            tahun: row.getCell(1).value.toString(),
            indikator: row.getCell(2).value,
            skor: row.getCell(3).value.toString(),
            keterangan: row.getCell(4).value,
            kegiatan: row.getCell(5).value,
            nilai: row.getCell(6).value,
            pusat: row.getCell(7).value,
            provinsi: row.getCell(8).value,
            kabupaten: row.getCell(9).value,
            desa: row.getCell(10).value,
            csr: row.getCell(11).value,
            lainnya: row.getCell(12).value,
          };
          tableDataIKS.push(data);
        });

        this.tableDataIDM = tableDataIDM;
        this.tableDataIKE = tableDataIKE;
        this.tableDataIKS = tableDataIKS;
      } catch (error) {
        console.error("Error reading Excel file: ", error);
      }
    },
  },
};
</script>

<style>
.wrapper-h1 {
  padding: 100px;
}

.excelbutton-wrapper {
  display: flex;
  width: 100%;
  justify-content: space-between;

  margin-top: 100px;
}

.card-idm {
  height: auto;
  display: flex;
}

.add-clear-buttons {
  display: flex;
  width: 100%;
  flex-direction: column;
  max-width: fit-content;
  gap: 10px;
}
</style>
