<template>
  <div class="prenotazioni-container">
    <h2>Le tue Prenotazioni</h2>
    <p class="subtitle">Esami a cui sei attualmente iscritta/o.</p>

    <div class="results-section">
      <table class="data-table">
        <thead>
        <tr>
          <th>Insegnamento</th>
          <th>Data e Ora</th>
          <th>Aula</th>
          <th>Stato</th>
          <th>Azione</th>
        </tr>
        </thead>
        <tbody>
        <tr>
          <td><strong>Sicurezza Informatica</strong></td>
          <td>15/10/2026 - 09:00</td>
          <td>Aula Magna</td>
          <td><span class="status-badge">Confermato</span></td>
          <td><button class="btn-cancel" @click="cancella">Annulla</button></td>
        </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
export default {
  name: 'PrenotazioniView',
  methods: {
    cancella() {
      alert("Mockup: Prenotazione annullata!");
    }
  }
}
</script>

<style scoped>
.prenotazioni-container { display: flex; flex-direction: column; gap: 20px; }
h2 { color: #812936; margin-bottom: 0; }
.subtitle { color: #718096; margin-top: 5px; }
.data-table { width: 100%; border-collapse: collapse; background: white; box-shadow: 0 1px 3px rgba(0,0,0,0.1); border-radius: 8px; overflow: hidden; }
.data-table th, .data-table td { padding: 15px; text-align: left; border-bottom: 1px solid #e2e8f0; }
.data-table th { background-color: #f7fafc; font-weight: 600; color: #4a5568; }
.status-badge { background-color: #ebf8f2; color: #38a169; padding: 4px 8px; border-radius: 12px; font-size: 0.85rem; font-weight: bold; }
.btn-cancel { background-color: #fc8181; color: white; border: none; padding: 8px 15px; border-radius: 4px; font-weight: bold; cursor: pointer; }
.btn-cancel:hover { background-color: #e53e3e; }
</style>