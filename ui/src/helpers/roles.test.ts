import { RoleNames } from '../models/role';
import { roleNameToEnum } from './roles';

type testPayload = {
  name: string;
  want: RoleNames;
  role: string;
};

test('something', () => {
  const cases: testPayload[] = [
    {
      name: 'unknown',
      want: RoleNames.Unknown,
      role: '',
    },
    {
      name: 'administrators',
      want: RoleNames.Administrators,
      role: 'administrators',
    },
    {
      name: 'campaign_owners',
      want: RoleNames.CampaignOwners,
      role: 'campaign_owners',
    },
    {
      name: 'users',
      want: RoleNames.Users,
      role: 'users',
    },
  ];
  cases.forEach(tc => {
    const got = roleNameToEnum(tc.role);
    expect(got).toEqual(tc.want);
  });
});
