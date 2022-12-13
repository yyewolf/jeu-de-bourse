import axios from 'axios';

let url = import.meta.env.VITE_JDB_BACKEND;

const axiosConfig = {
    baseURL: url+'/api',
    timeout: 30000,
};

const back = axios.create(axiosConfig);

export default back;