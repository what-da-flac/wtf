import { RoleNames } from './role';

export type MenuItem = {
  label: string;
  link: string;
  roles?: RoleNames[];
};
