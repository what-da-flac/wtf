import {
  IconBellQuestion,
  IconDoorEnter,
  IconHome2,
  IconLock,
  IconNotes,
  IconUser,
} from '@tabler/icons-react';

// Models :
import { RoleNames } from '../models/role';
import { MenuGroup } from '../models/menuGroup';

// returns the default menu groups for the whole application
// these items appear in the sidebar menu
export function MenuGroups(): MenuGroup[] {
  return [
    {
      label: 'Cinito',
      icon: IconHome2,
      initiallyOpened: false,
      items: [
        {
          label: 'About',
          link: '/',
        },
      ],
    },
    {
      label: 'Movies',
      icon: IconNotes,
      initiallyOpened: true,
      items: [
        {
          label: 'List',
          link: '/movies',
          roles: [RoleNames.Users],
        },
      ],
    },
    {
      label: 'Incoming',
      initiallyOpened: true,
      icon: IconDoorEnter,
      items: [
        {
          label: 'Torrents',
          link: '/torrents',
          roles: [RoleNames.Administrators],
        },
      ],
    },
    {
      label: 'User',
      icon: IconUser,
      initiallyOpened: false,
      items: [
        {
          label: 'Profile',
          link: '/user/profile',
          roles: [RoleNames.Administrators],
        },
        // {
        //   label: 'Recompensas',
        //   link: '/user/rewards',
        //   roles: [RoleNames.Users],
        // },
      ],
    },
    {
      label: 'User Management',
      icon: IconLock,
      initiallyOpened: true,
      items: [
        {
          label: 'Roles',
          link: '/roles',
          roles: [RoleNames.Administrators],
        },
        {
          label: 'Users',
          link: '/users',
          roles: [RoleNames.Administrators],
        },
      ],
    },
  ];
}
