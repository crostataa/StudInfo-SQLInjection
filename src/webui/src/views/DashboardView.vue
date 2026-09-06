<template>
  <div class="dashboard">
    <header class="dashboard-header">
      <h2>Benvenuta/o, {{ studente.nome }} {{ studente.cognome }}</h2>
      <p class="matricola">Matricola: {{ studente.matricola }}</p>
    </header>

    <!-- Riepilogo Statistiche -->
    <div class="stats-grid">
      <div class="stat-card">
        <h3>Media Ponderata</h3>
        <div class="stat-value">{{ statistiche.media }}</div>
      </div>
      <div class="stat-card">
        <h3>CFU Acquisiti</h3>
        <div class="stat-value">{{ statistiche.cfu }} / 120</div>
      </div>
      <div class="stat-card">
        <h3>Esami Sostenuti</h3>
        <div class="stat-value">{{ statistiche.esamiSostenuti }}</div>
      </div>
    </div>

    <!-- Prossimi impegni -->
    <div class="upcoming-section">
      <h3>I tuoi prossimi appelli</h3>
      <table class="data-table" v-if="prossimiAppelli.length > 0">
        <thead>
        <tr>
          <th>Materia</th>
          <th>Data</th>
          <th>Ora</th>
          <th>Aula</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="appello in prossimiAppelli" :key="appello.id">
          <td><strong>{{ appello.materia }}</strong></td>
          <td>{{ appello.data }}</td>
          <td>{{ appello.ora }}</td>
          <td>{{ appello.aula }}</td>
        </tr>
        </tbody>
      </table>
      <p v-else class="no-data">Non hai appelli prenotati a breve.</p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'DashboardView',
  data() {
    return {
      // DATI FINTI PER IL MOCKUP (Nella Fase 3 arriveranno dal database MySQL)
      studente: {
        nome: 'Mario',
        cognome: 'Rossi',
        matricola: '123456'
      },
      statistiche: {
        media: '27.5',
        cfu: 85,
        esamiSostenuti: 10
      },
      prossimiAppelli: [
        { id: 1, materia: 'Sicurezza Informatica', data: '15/10/2026', ora: '09:00', aula: 'Aula Magna' },
        { id: 2, materia: 'Basi di Dati', data: '22/10/2026', ora: '14:30', aula: 'Laboratorio 3' }
      ]
    }
  }
}
</script>

<style scoped>
.dashboard {
  display: flex;
  flex-direction: column;
  gap: 30px;
}

.dashboard-header h2 {
  margin-bottom: 5px;
  color: #812936;
}

.matricola {
  color: #718096;
  font-size: 1.1rem;
  margin-top: 0;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 20px;
}

.stat-card {
  background: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.05);
  text-align: center;
  border-top: 4px solid #812936;
}

.stat-card h3 {
  font-size: 1rem;
  color: #4a5568;
  margin-bottom: 10px;
  margin-top: 0;
}

.stat-value {
  font-size: 2rem;
  font-weight: bold;
  color: #812936;
}

.upcoming-section h3 {
  color: #812936;
  border-bottom: 2px solid #e2e8f0;
  padding-bottom: 10px;
}

.data-table {
  width: 100%;
  border-collapse: collapse;
  background: white;
  box-shadow: 0 1px 3px rgba(0,0,0,0.1);
  border-radius: 8px;
  overflow: hidden;
}

.data-table th, .data-table td {
  padding: 15px;
  text-align: left;
  border-bottom: 1px solid #e2e8f0;
}

.data-table th {
  background-color: #f7fafc;
  font-weight: 600;
  color: #812936;
}

.data-table tr:last-child td {
  border-bottom: none;
}

.no-data {
  color: #718096;
  font-style: italic;
}
</style>