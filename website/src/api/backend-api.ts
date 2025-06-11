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
  getUserImage(id?: string): Promise<AxiosResponse> {
    const config = { withCredentials: true };
    const data = id ? { user_id: id } : {};
    return protectedAxiosApi.post('/user/image', data, config);
  },
  getUser(username: string): Promise<AxiosResponse> {
    return protectedAxiosApi.get(`/user/${username}`, { withCredentials: true });
  },
  getChatHistory(): Promise<AxiosResponse> {
    return protectedAxiosApi.get('/chat/history', { withCredentials: true });
  },
  getPastRuns(): Promise<AxiosResponse> {
    return protectedAxiosApi.get('/runs', { withCredentials: true });
  },
  deletePastRun(id: string): Promise<AxiosResponse> {
    return protectedAxiosApi.delete(`/runs/${id}`, { withCredentials: true });
  },
  savePlannedRun(route: string, name: string, distance: number): Promise<AxiosResponse> {
    return protectedAxiosApi.post(
      '/runs/plan',
      { route, name, distance },
      { withCredentials: true }
    );
  },
  getPlannedRuns(): Promise<AxiosResponse> {
    return protectedAxiosApi.get('/runs/plan', { withCredentials: true });
  },
  deletePlannedRun(id: string): Promise<AxiosResponse> {
    return protectedAxiosApi.delete(`/runs/plan/${id}`, { withCredentials: true });
  },
}
