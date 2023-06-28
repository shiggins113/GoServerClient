<template>
  <div>
    <!-- Table to display the certificate data -->
    <table>
      <thead>
        <tr>
          <!-- Table headers -->
          <th class="table-header">Certificate ID</th>
          <th class="table-header">Certificate Details</th>
          <th class="table-header">Valid From</th>
          <th class="table-header">Valid Until</th>
          <th></th>
        </tr>
      </thead>
      <tbody>
        <!-- Table rows to display certificate data dynamically -->
        <tr v-for="item in tableData" :key="item.id">
          <td>{{ item.certID }}</td>
          <td>{{ item.cert }}</td>
          <td>{{ item.beforeDate }}</td>
          <td>{{ item.afterDate }}</td>
          <td>
            <!-- Delete button to remove a certificate -->
            <button class="delete-button" @click="confirmDelete(item.certID)">
              Delete
            </button>
          </td>
        </tr>
        <!-- Additional table rows can be added dynamically -->
      </tbody>
    </table>
  </div>
</template>

<style>
/* Style the table and its elements as per your requirements */

/* Set the width of the table to 450% of its container */
table {
  width: 450%;
  margin-top: 50px;
  margin-left: -950px; /* Adjust the left margin as needed */
  border-collapse: collapse;
}

/* Style the table headers */
.table-header {
  text-align: left;
}

/* Style the delete button */
.delete-button {
  background-color: red;
  color: white;
  padding: 5px 10px;
  border: none;
  cursor: pointer;
}

/* Style the table cells (both header and data cells) */
th,
td {
  border: 1px solid #ddd;
  padding: 8px;
  text-align: left;
}

/* Style the table header row */
th {
  background-color: #f2f2f2;
}

/* Center the table within its container */
.table-container {
  display: flex;
  justify-content: center;
}
</style>

<script>
import axios from 'axios';

export default {
  name: 'TableComponent',
  data() {
    return {
      tableData: [], // Array to hold the certificate data to be displayed in the table
    };
  },
  mounted() {
    // Call the fetchTableData method when the component is mounted
    this.fetchTableData();
  },
  methods: {
    // Method to confirm the deletion of a certificate
    confirmDelete(id) {
      const confirmMessage = 'Are you sure you want to delete this item?';
      if (window.confirm(confirmMessage)) {
        // If confirmed, call the deleteItem method
        this.deleteItem(id);
      } else {
        // User canceled the delete operation, do nothing or perform any additional action if needed
      }
    },
    // Method to fetch the certificate data from the server
    fetchTableData() {
      axios
        .get('http://localhost:8081/certData')
        .then(response => {
          this.tableData = response.data; // Assign the fetched data to the tableData array
        })
        .catch(error => {
          console.error('Error fetching table data:', error);
        });
    },
    // Method to delete a certificate
    deleteItem(id) {
      console.log('Deleting item with id:', id);
      axios
        .delete('http://localhost:8081/certDelete', {
          params: {
            certID: id
          }
        })
        // eslint-disable-next-line no-unused-vars
        .then(response => {
          // If the deletion is successful, update the tableData array
          this.tableData = this.tableData.filter(item => item.certID !== id);
        })
        .catch(error => {
          console.error('Error deleting item:', error);
        });
    },
  },
};
</script>