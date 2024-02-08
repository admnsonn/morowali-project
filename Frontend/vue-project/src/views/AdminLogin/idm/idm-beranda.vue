<template>
  <div class="wrapper-h1">
    <div class="card" style="width: 100%">
      <ul class="list-group list-group-flush">
        <li class="list-group-item">Cras justo odio</li>
        <li class="list-group-item">Dapibus ac facilisis in</li>
        <li class="list-group-item">Vestibulum at eros</li>
      </ul>
    </div>
  </div>
</template>

<script>
import ExcelJS from "exceljs";

export default {
  name: "exceljsTest",
  data() {
    return {
      tableDataIDM: [],
      tableDataIKE: [],
      tableDataIKS: [],
    };
  },
  methods: {
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
            tahun: row.getCell(1).value,
            indikator: row.getCell(2).value,
            skor: row.getCell(3).value,
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
        });

        worksheetIKE.eachRow((row, index) => {
          if (index === 1) {
            return;
          }

          const data = {
            tahun: row.getCell(1).value,
            indikator: row.getCell(2).value,
            skor: row.getCell(3).value,
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
            tahun: row.getCell(1).value,
            indikator: row.getCell(2).value,
            skor: row.getCell(3).value,
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
</style>
