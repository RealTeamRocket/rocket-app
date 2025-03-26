import axios, { type AxiosResponse } from 'axios';

const axiosApi = axios.create({
  baseURL: '/api',
  timeout: 10000,
  headers: {'content-type': 'application/json'}
})

export default {
  helloWorld() {
    return axiosApi.get('/hello')
  }
}
