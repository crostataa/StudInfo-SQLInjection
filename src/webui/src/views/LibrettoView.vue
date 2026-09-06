<template>
  <div class="libretto-container">
    <h2>Libretto Universitario</h2>
    <p class="subtitle">Riepilogo della tua carriera accademica.</p>

    <table class="data-table">
      <thead>
      <tr>
        <th>Insegnamento</th>
        <th>Data Registrazione</th>
        <th>Voto</th>
        <th>CFU</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="esame in esami" :key="esame.id">
        <td><strong>{{ esame.materia }}</strong></td>
        <td>{{ esame.data }}</td>
        <td><span class="voto">{{ esame.voto }}</span></td>
        <td>{{ esame.cfu }}</td>
      </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  name: 'LibrettoView',
  data() {
    return {
      esami: [
        { id: 1, materia: 'Architettura dei Calcolatori', data: '12/02/2025', voto: '28', cfu: 9 },
        { id: 2, materia: 'Programmazione I', data: '20/06/2025', voto: '30 L', cfu: 12 }
      ]
    }
  }
}
</script>

<style scoped>
.libretto-container { display: flex; flex-direction: column; gap: 20px; }
h2 { color: #812936; margin-bottom: 0; }
.subtitle { color: #718096; margin-top: 5px; }
.data-table { width: 100%; border-collapse: collapse; background: white; box-shadow: 0 1px 3px rgba(0,0,0,0.1); border-radius: 8px; overflow: hidden; }
.data-table th, .data-table td { padding: 15px; text-align: left; border-bottom: 1px solid #e2e8f0; }
.data-table th { background-color: #f7fafc; font-weight: 600; color: #4a5568; }
.voto { font-weight: bold; color: #b02a37; font-size: 1.1rem; }
</style>