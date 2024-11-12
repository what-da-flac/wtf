import React from 'react';
import ReactDOM from 'react-dom/client';
import { BrowserRouter } from 'react-router-dom';
import { GoogleOAuthProvider } from '@react-oauth/google';

import reportWebVitals from './reportWebVitals';

// Mantine :
import { ModalsProvider } from '@mantine/modals';
import { createTheme, MantineProvider } from '@mantine/core';

// Component :
import App from './App';

// CSS :
import './styles/main.scss';

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);

console.log('starting web app');

// read google credentials from environment
console.log('checking environment variables');
const envVarNames = [
  'REACT_APP_BASE_API_URL',
  'REACT_APP_GOOGLE_CLIENT_ID',
  'REACT_APP_GOOGLE_CLIENT_SECRET',
  'REACT_APP_TAG_NAME',
  'REACT_APP_GOOGLE_API_KEY',
];
const errors = [];
envVarNames.forEach(name => {
  if (!process.env[name]) {
    errors.push(`missing env var: ${name}`);
  }
});
if (errors.length > 0) {
  throw new Error(errors.join('\n'));
}
const clientId = process.env.REACT_APP_GOOGLE_CLIENT_ID;

// display git info
const gitTag = process.env.REACT_APP_TAG_NAME;

console.table({
  tag: gitTag,
});

const theme = createTheme({
  cursorType: 'pointer',
});

root.render(
  <GoogleOAuthProvider clientId={clientId}>
    <MantineProvider theme={theme} defaultColorScheme="dark">
      <ModalsProvider>
        <BrowserRouter>
          <App />
        </BrowserRouter>
      </ModalsProvider>
    </MantineProvider>
  </GoogleOAuthProvider>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
