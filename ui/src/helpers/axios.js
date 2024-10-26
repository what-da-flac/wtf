import axios from 'axios';
import { isLoggedIn, loadUserProfile } from './sso_service';

const baseURL = process.env.REACT_APP_BASE_API_URL;

const instance = axios.create({
  baseURL,
});

instance.interceptors.request.use(config => {
  if (isLoggedIn()) {
    const profile = loadUserProfile();
    config.headers.Authorization = `Bearer ${profile.access_token}`;
    return Promise.resolve(config);
  }
  return Promise.resolve(config);
});

export default instance;
