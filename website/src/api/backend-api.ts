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
  },

  getUserStatistics(id?: string): Promise<AxiosResponse> {
    const data = id ? { id } : {};
    const config = { withCredentials: true };
    return protectedAxiosApi.post('/user/statistics', data, config);
  },
  getActivityFeed(): Promise<AxiosResponse> {
    return protectedAxiosApi.get('/activites', { withCredentials: true });
  },
  getUser(): Promise<AxiosResponse> {
    return protectedAxiosApi.get('/user', { withCredentials: true });
  },
  getChatHistory(): Promise<AxiosResponse> {
    return protectedAxiosApi.get('/chat/history', { withCredentials: true });
  },
  getChallenges(): Promise<AxiosResponse> {
  return protectedAxiosApi.get('/challenges/new', { withCredentials: true });
},
}
