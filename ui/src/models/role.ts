import { trimAll } from '../helpers/text_utils';

export enum RoleNames {
  Unknown,
  Administrators,
  CampaignOwners,
  Users,
}

export class Role {
  id: string;
  name: string;
  description: string;

  constructor() {
    this.id = '';
    this.name = '';
    this.description = '';
  }
}

export function roleValidation() {
  return {
    name: (value: string) =>
      trimAll(value).length === 0 ? 'Name is mandatory' : null,
  };
}
