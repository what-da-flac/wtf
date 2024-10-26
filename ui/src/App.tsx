import React, { useEffect, useLayoutEffect, useState } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';
import { TokenResponse, useGoogleLogin } from '@react-oauth/google';

// Mantine :
import { useDisclosure } from '@mantine/hooks';
import { AppShell, Burger, Drawer, Group } from '@mantine/core';

// Components :
import Logo from './components/Logo';
import Navbar from './components/Navbar';
import RouteSwitcher from './RouteSwitcher';

// Helpers :
import { parseToken, saveUserProfile } from './helpers/sso_service';

// Models :
import { User, UserProfileResponse } from './models/user';

// CSS :
import '@mantine/core/styles.css';
import '@mantine/dates/styles.css';
import '@mantine/core/styles/global.css';
import instance from './helpers/axios';

function App() {
  const location = useLocation();
  const navigate = useNavigate();
  // user contains the access token information, and expiration
  const [user, setUser] = useState<TokenResponse>();
  // profile contains the parsed information after we verify the access token with Google endpoint
  const [profile, setProfile] = useState<UserProfileResponse>();
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const invitationToken = location.pathname.includes('/invitation/')
    ? location.pathname.split('/invitation/').pop()
    : null;

  useLayoutEffect(() => {
    instance.interceptors.response.use(
      function (response) {
        return response;
      },
      function (error) {
        const status = error?.response?.status;
        switch (status) {
          case 403:
            navigate('/errors/unauthorized');
            break;
          default:
            if (status >= 400) {
              return Promise.reject(error);
            }
        }
      }
    );
  }, []);

  const login = useGoogleLogin({
    onSuccess: codeResponse => {
      setUser(codeResponse);
      saveUserProfile(codeResponse);
      setIsLoading(true);
      setTimeout(() => {
        setIsLoading(false);
      }, 2000);
    },
    onError: error => console.log('Login Failed:', error),
  });

  useEffect(() => {
    parseToken()
      .then(p => setProfile(p))
      .catch(() => {});
  }, [user]);

  const logOut = () => {
    setProfile(undefined);
    localStorage.clear();
    navigate('/');
  };

  function CollapseDesktop() {
    const [opened] = useDisclosure();
    const [mobileOpened] = useDisclosure();
    const [desktopOpened, { toggle: toggleDesktop }] = useDisclosure(true);
    const [Mobileopened, { open, close }] = useDisclosure(false);
    const user = new User();
    user.email = profile?.email;
    user.name = profile?.name;
    user.image = profile?.picture;
    const version = process.env.REACT_APP_TAG_NAME || '';

    return (
      <React.Fragment>
        <AppShell
          padding="md"
          header={{ height: 60 }}
          navbar={{
            width: 250,
            breakpoint: 'sm',
            collapsed: { mobile: !mobileOpened, desktop: !desktopOpened },
          }}
        >
          <AppShell.Header className="flex-verticle-center">
            <Group pl="sm">
              <Burger
                className="web-menu"
                opened={opened}
                onClick={toggleDesktop}
                aria-label="Toggle navigation"
              />
              <Burger
                className="mobile-menu"
                onClick={open}
                aria-label="Toggle navigation"
              />
              <Drawer
                size="300px"
                title={<Logo />}
                onClose={close}
                opened={Mobileopened}
                className="mobile-drawer"
                overlayProps={{ backgroundOpacity: 0.5, blur: 4 }}
              >
                <Navbar
                  login={login}
                  logout={logOut}
                  profile={profile}
                  version={version}
                />
              </Drawer>
              <Logo />
            </Group>
          </AppShell.Header>
          <AppShell.Navbar>
            <Navbar
              isLoading={isLoading}
              login={login}
              logout={logOut}
              profile={profile}
              version={version}
            />
          </AppShell.Navbar>
          <AppShell.Main>
            <RouteSwitcher profile={profile} />
          </AppShell.Main>
        </AppShell>
      </React.Fragment>
    );
  }

  return (
    <React.Fragment>
      <CollapseDesktop />
    </React.Fragment>
  );
}

export default App;
