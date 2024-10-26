import React from 'react';
import { IconLogout } from '@tabler/icons-react';

// Mantine :
import {
  Box,
  Code,
  Group,
  rem,
  ScrollArea,
  Skeleton,
  ThemeIcon,
} from '@mantine/core';

// Components :
import Logo from './Logo';
import { UserButton } from './UserButton';
import { LinksGroup } from './NavbarLinksGroup';

// Models :
import { UserProfileResponse } from '../models/user';

// Helpers :
import { glueMenus } from '../helpers/menu';
import { MenuGroups } from '../helpers/menu_groups';

// CSS :
import classes from './NavbarNested.module.css';

type NavbarParams = {
  version: string;
  profile: UserProfileResponse | undefined;
  login: any;
  logout: any;
  isLoading?: boolean;
};

// NavbarMain is the real navbar that appears on the left pane of the app.
export default function NavbarMain({
  isLoading,
  profile,
  version,
  login,
  logout,
}: NavbarParams) {
  const menuGroups = glueMenus(MenuGroups(), profile);

  function Loader() {
    return (
      <div style={{ padding: '10px' }}>
        {Array(15)
          .fill(0)
          .map((_, index) => (
            <Skeleton key={index} h={28} mt="md" animate={true} />
          ))}
      </div>
    );
  }

  return isLoading ? (
    <Loader />
  ) : (
    <React.Fragment>
      <nav className={classes.navbar}>
        <div className={classes.header}>
          <Group justify="space-between">
            <Logo style={{ width: rem(120) }} />
            <Code fw={700}>{version}</Code>
          </Group>
        </div>
        <ScrollArea className={classes.links}>
          <div className={classes.linksInner}>
            {menuGroups.map(menuGroup => (
              <LinksGroup menuGroup={menuGroup} key={menuGroup.label} />
            ))}
          </div>
        </ScrollArea>
      </nav>
      {profile?.email ? (
        <React.Fragment>
          <div onClick={logout} className={classes.logout}>
            <ThemeIcon variant="light" size={30}>
              <IconLogout style={{ width: rem(18), height: rem(18) }} />
            </ThemeIcon>
            <Box fw={500}>Logout</Box>
          </div>
          <div className={classes.footer}>
            <UserButton user={profile} />
          </div>
        </React.Fragment>
      ) : (
        <div className={classes.footer} onClick={login}>
          <Box fw={500}>Login</Box>
        </div>
      )}
    </React.Fragment>
  );
}
