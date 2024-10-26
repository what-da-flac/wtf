import axios from 'axios';
import { googleLogout, TokenResponse } from '@react-oauth/google';
import instance from './axios';

// userProfileKey stores the access token and expiration from Google OAuth.
const userProfileKey = '_profile';

function loadUserProfile(): TokenResponse | undefined {
  const data = localStorage.getItem(userProfileKey);
  if (!data) return;
  return JSON.parse(data);
}

function saveUserProfile(data: TokenResponse) {
  localStorage.setItem(userProfileKey, JSON.stringify(data));
}

function clearUserProfile() {
  localStorage.clear();
}

async function parseToken() {
  const localProfile = loadUserProfile();
  if (!localProfile) throw new Error('no token found');
  const path = 'https://www.googleapis.com/oauth2/v1/userinfo';
  const res = await axios.get(
    `${path}?access_token=${localProfile.access_token}`,
    {
      headers: {
        Authorization: `Bearer ${localProfile.access_token}`,
        Accept: 'application/json',
      },
    }
  );
  // obtain google parsed data
  const { data } = res;
  // register login for this user in the API
  const user = {
    ...data,
    image: data.picture,
  };
  const loginRes = await instance.post(`/v1/users/login`, user);
  user.roles = loginRes.data?.roles;
  return user;
}

function logout() {
  googleLogout();
  clearUserProfile();
}

function isLoggedIn(): boolean {
  const profile = loadUserProfile();
  return !!profile;
}

export {
  clearUserProfile,
  isLoggedIn,
  loadUserProfile,
  logout,
  parseToken,
  saveUserProfile,
};
