import { User } from './user';
import { trimAll, truncateTime } from '../helpers/text_utils';

export class Campaign {
  id: string;
  name: string;
  start_date: Date;
  end_date: Date;
  user: User;
  image_url: string;
  description: string;
  status: string;

  constructor() {
    this.id = '';
    this.name = '';
    this.description = '';
    this.user = new User();
    this.start_date = new Date();
    this.end_date = new Date();
    this.status = '';
    this.image_url = '';
  }
}

export function notTrunCampaign(v: any): Campaign {
  if (!v) return new Campaign();
  return {
    ...v,
    start_date: new Date(v['start_date']),
    end_date: new Date(v['end_date']),
    image: v['image_url'],
  };
}

export function toCampaign(v: any): Campaign {
  if (!v) return new Campaign();
  return {
    ...v,
    start_date: truncateTime(new Date(v['start_date'])),
    end_date: truncateTime(new Date(v['end_date'])),
    image: v['image_url'],
  };
}

export function campaignValidation() {
  return {
    name: (value: string) =>
      trimAll(value).length === 0 ? 'Nombre es mandatorio' : null,
    start_date: (value: Date) =>
      !value ? 'Fecha de inicio es mandatorio' : null,
    end_date: (value: Date) =>
      !value ? 'Fecha de termino es mandatorio' : null,
    image_url: (value: string) => (!value ? 'Imagen es mandatorio' : null),
  };
}

export function cleanCampaign(c: Campaign): Campaign {
  c.start_date = truncateTime(c.start_date);
  c.end_date = truncateTime(c.end_date);
  return c;
}
