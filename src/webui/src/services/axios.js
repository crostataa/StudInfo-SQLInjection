import axios from 'axios';

// Creiamo un'istanza di Axios configurata per puntare al nostro server Go
const api = axios.create({
	baseURL: 'http://localhost:3000', // L'indirizzo esatto del nostro backend!
	timeout: 10000,
});

export default api;