import * as React from 'react';
import { Route, Routes } from 'react-router-dom';

// Pages :
import Root from './pages/Root';
import Error404 from './pages/Error404';
import { RoleEdit } from './pages/RoleEdit';
import { UserEdit } from './pages/UserEdit';
import { UsersList } from './pages/UserList';
import { RolesList } from './pages/RolesList';
import MovieList from './pages/MovieList';
import ProfileEdit from './pages/UserProfileEdit';

import Error401 from './pages/Error401';
import Error403 from './pages/Error403';

// Models :
import { RoleNames } from './models/role';
import { UserProfileResponse } from './models/user';

// Helpers :
import { roleNameToEnum } from './helpers/roles';
import TorrentsList from './pages/TorrentsList';
import TorrentBulkForm from './pages/TorrentBulkForm';
import { TorrentEdit } from './pages/TorrentEdit';

type params = {
  profile: UserProfileResponse | undefined;
};

type RouteRole = {
  route: React.ReactElement;
  roles: RoleNames[];
};

// checks if the provided rule has role access defined in the provided profile
function routeHasAccess(
  route: RouteRole,
  profile: UserProfileResponse | undefined
): boolean {
  // no role declared in the route, means anyone can access it
  if (route.roles?.length === 0) return true;

  // profile may come undefined
  if (!profile) return false;

  // loop over role accesses
  for (let i = 0; i < route.roles.length; i++) {
    for (let j = 0; j < profile.roles?.length; j++) {
      const role: string = profile.roles[j];
      const roleName: RoleNames = roleNameToEnum(role);
      if (route.roles[i] === roleName) {
        return true;
      }
    }
  }
  return false;
}

// returns an array of routes, which have specific role access
function routesWithRoles(): RouteRole[] {
  return [
    {
      route: <Route path="/movies" element={<MovieList />} />,
      roles: [RoleNames.Users],
    },
    {
      route: <Route path="/torrents" element={<TorrentsList />} />,
      roles: [RoleNames.Administrators],
    },
    {
      route: <Route path="/torrents/new" element={<TorrentBulkForm />} />,
      roles: [RoleNames.Administrators],
    },
    {
      route: <Route path="/torrents/:id" element={<TorrentEdit />} />,
      roles: [RoleNames.Administrators],
    },
    {
      route: <Route path="/user/profile" element={<ProfileEdit />} />,
      roles: [RoleNames.Administrators],
    },
    {
      route: <Route path="/roles/:id" element={<RoleEdit />} />,
      roles: [RoleNames.Administrators],
    },
    {
      route: <Route path="/roles" element={<RolesList />} />,
      roles: [RoleNames.Administrators],
    },
    {
      route: <Route path="/users" element={<UsersList />} />,
      roles: [RoleNames.Administrators],
    },
    {
      route: <Route path="/users/:id" element={<UserEdit />} />,
      roles: [RoleNames.Administrators],
    },
  ];
}

export default function RouteSwitcher({ profile }: params) {
  //  allow to access routes based on role membership
  const routes = routesWithRoles();
  return (
    <Routes>
      <Route path="/" element={<Root />} />

      {routes
        .filter(route => routeHasAccess(route, profile))
        .map((route, index) => (
          <Route
            key={route.route.props.path.toString()}
            {...route.route.props}
          />
        ))}
      <Route path="/errors/unauthenticated" element={<Error401 />} />
      <Route path="/errors/unauthorized" element={<Error403 />} />
      <Route path="/*" element={<Error404 />} />
    </Routes>
  );
}
