import { toDate } from '../helpers/text_utils';

export class User {
  email?: string;
  image?: string | undefined;
  name?: string;
  id: string;
  last_login: Date;

  constructor() {
    this.email = '';
    this.image = '';
    this.name = '';
    this.id = '';
    this.last_login = new Date();
  }
}

export function toUser(v: any): User {
  if (!v) return new User();
  return {
    ...v,
    last_login: toDate(new Date(v['last_login'])),
  };
}

export type UserProfileResponse = {
  id: string;
  picture: string;
  name: string;
  email: string;
  roles: string[];
};
