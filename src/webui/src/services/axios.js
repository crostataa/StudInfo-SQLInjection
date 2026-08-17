import axios from "src/webui/src/services/axios";

const instance = axios.create({
	baseURL: __API_URL__,
	timeout: 1000 * 5
});

export default instance;
