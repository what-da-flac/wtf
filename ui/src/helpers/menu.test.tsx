import { UserProfileResponse } from '../models/user';
import { glueMenus, menuItemHasRoleMembership } from './menu';
import { RoleNames } from '../models/role';
import { MenuGroup } from '../models/menuGroup';
import { MenuItem } from '../models/menuItem';
import React from 'react';
import { IconGauge } from '@tabler/icons-react';

function mockedProfile(): UserProfileResponse {
  return {
    name: 'test',
    email: 'me@example.com',
    roles: [],
    picture: 'whatever',
    id: '1',
  };
}

function mockedIcon(): React.FC<any> {
  return IconGauge;
}

test('no profile no auth membership', () => {
  const item: MenuItem = {
    label: 'Inicio',
    link: '/',
  };
  const got = menuItemHasRoleMembership(item, undefined);
  expect(got).toEqual(true);
});

test('membership matches profile', () => {
  const item: MenuItem = {
    label: 'Matches',
    link: '/',
    roles: [RoleNames.Unknown],
  };
  const profile = mockedProfile();
  profile.roles = ['users'];
  const got = menuItemHasRoleMembership(item, profile);
  expect(got).toEqual(false);
});

test('home page no auth', () => {
  const menuGroups: MenuGroup[] = [
    {
      label: 'Home',
      icon: mockedIcon(),
      initiallyOpened: false,
      items: [
        {
          label: 'Inicio',
          link: '/',
        },
      ],
    },
  ];
  const profile = mockedProfile();
  profile.roles = [];
  const result = glueMenus(menuGroups, profile);
  const expected = [
    {
      label: 'Home',
      icon: mockedIcon(),
      initiallyOpened: false,
      items: [
        {
          label: 'Inicio',
          link: '/',
        },
      ],
    },
  ];
  expect(result).toEqual(expected);
});

test('home page auth', () => {
  const menus: MenuGroup[] = [
    {
      label: 'Home',
      icon: mockedIcon(),
      initiallyOpened: false,
      items: [
        {
          label: 'Inicio',
          link: '/',
          roles: [RoleNames.Users],
        },
        {
          label: 'Another',
          link: '/',
          roles: [RoleNames.Administrators],
        },
      ],
    },
  ];
  const profile = mockedProfile();
  profile.roles = ['users'];
  const result = glueMenus(menus, profile);
  const expected = [
    {
      label: 'Home',
      icon: mockedIcon(),
      initiallyOpened: false,
      items: [
        {
          label: 'Inicio',
          link: '/',
          roles: [RoleNames.Users],
        },
      ],
    },
  ];
  expect(result).toEqual(expected);
});
