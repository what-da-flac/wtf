import { toDate } from '../helpers/text_utils';
import { toUser, User } from './user';

export class Torrent {
  id: string;
  created: Date;
  updated: Date;
  magnet_link: string;
  status: string;
  user: User;
  filename: string;
  name: string;
  total_size: string;
  percent: number;
  eta: string;

  constructor() {
    this.id = '';
    this.created = new Date();
    this.updated = new Date();
    this.magnet_link = '';
    this.status = 'DRAFT';
    this.user = new User();
    this.filename = '';
    this.name = 'New torrent';
    this.total_size = '';
    this.percent = 0;
    this.eta = '';
  }
}

export function toTorrent(v: any): Torrent {
  if (!v) return new Torrent();
  return {
    ...v,
    created: toDate(new Date(v['created'])),
    updated: toDate(new Date(v['updated'])),
    user: toUser(v['user']),
  };
}

export function torrentValidation() {
  return {
    title: (value: string) =>
      value.length === 0 ? 'Title is mandatory' : null,
    magnet_link: (value: string) =>
      value.length === 0 ? 'Magnet Link is mandatory' : null,
  };
}
