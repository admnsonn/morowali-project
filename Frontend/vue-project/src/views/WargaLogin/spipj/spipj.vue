<template>
  <div class="wrapper-h1">
    <h1>SPIPJ</h1>
    <table class="tabel">
      <thead>
        <tr>
          <th>No.</th>
          <th>Kolom</th>
          <th>Kolom</th>
          <th>Kolom</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td>1</td>
          <td>asdasdads</td>
          <td>asdasd</td>
          <td>asdasd</td>
        </tr>
        <tr>
          <td>1</td>
          <td>asdasdads</td>
          <td>asdasd</td>
          <td>asdasd</td>
        </tr>
        <tr>
          <td>1</td>
          <td>asdasdads</td>
          <td>asdasd</td>
          <td>asdasd</td>
        </tr>
        <tr>
          <td>1</td>
          <td>asdasdads</td>
          <td>asdasd</td>
          <td>asdasd</td>
        </tr>
        <tr>
          <td>1</td>
          <td>asdasdads</td>
          <td>asdasd</td>
          <td>asdasd</td>
        </tr>
      </tbody>
    </table>
    <button @click="generateXLSX" type="button" class="btn btn-success">
      Expor File Excel
    </button>
  </div>
</template>

<script>
import ExcelJS from "exceljs";

export default {
  name: "exceljsTest",
  methods: {
    async generateXLSX() {
      const options = {
        useStyles: true,
        useSharedStrings: true,
      };
      const workbook = new ExcelJS.Workbook(options);
      workbook.creator = "aditya patty";
      workbook.lastModifiedBy = "aditya patty";
      workbook.created = new Date(2024, 2, 7);
      workbook.modified = new Date();
      workbook.lastPrinted = new Date(2024, 2, 6);

      const worksheet = workbook.addWorksheet("Sheet 1");

      worksheet.columns = [
        { header: "Id", key: "id", width: 10 },
        { header: "Name", key: "name", width: 32 },
        { header: "D.O.B.", key: "dob", width: 15 },
        { header: "Age", key: "age", width: 10 },
      ];

      for (let index = 0; index < 10; index++) {
        worksheet.addRow([index, "anjas", new Date(), 20 + index]);
      }

      const buffer = await workbook.xlsx.writeBuffer();
      const blob = new Blob([buffer], {
        type: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
      });
      const url = URL.createObjectURL(blob);

      // Create a link and set the URL as the href
      const link = document.createElement("a");
      link.href = url;

      // Set the download attribute of the link to the desired filename
      link.download = "test.xlsx";

      // Append the link to the body
      document.body.appendChild(link);

      // Programmatically click the link to start the download
      link.click();

      // Remove the link from the body
      document.body.removeChild(link);
    },
  },
};
</script>

<style></style>
