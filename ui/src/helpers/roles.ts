import { RoleNames } from '../models/role';

export function roleNameToEnum(role: string): RoleNames {
  switch (role) {
    case 'administrators':
      return RoleNames.Administrators;
    case 'campaign_owners':
      return RoleNames.CampaignOwners;
    case 'users':
      return RoleNames.Users;
    default:
      return RoleNames.Unknown;
  }
}
