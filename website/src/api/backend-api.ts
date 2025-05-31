import axios, { type AxiosResponse } from 'axios';

const publicAxiosApi = axios.create({
  baseURL: '/api/v1',
  timeout: 10000,
  headers: {'content-type': 'application/json'}
})

const protectedAxiosApi = axios.create({
  baseURL: '/api/v1/protected',
  timeout: 10000,
  headers: {'content-type': 'application/json'}
})

export default {
  login(email: string, password: string): Promise<AxiosResponse> {
    return publicAxiosApi.post('/login', { email, password })
  },
  register(username: string, email: string, password: string): Promise<AxiosResponse> {
    return publicAxiosApi.post('/register', { username, email, password })
  },
  checkAuthenticated(): Promise<AxiosResponse> {
    return protectedAxiosApi.get('/', { withCredentials: true })
  },
  logout(): Promise<AxiosResponse> {
    console.log("Logging out")
    return publicAxiosApi.post('/logout', {}, { withCredentials: true })
  }
}
