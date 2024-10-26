import { roleNameToEnum } from './roles';

// Models :
import { MenuItem } from '../models/menuItem';
import { UserProfileResponse } from '../models/user';
import { MenuGroup } from '../models/menuGroup';

// returns true if any of the profile roles is allowed in the MenuItem
export function menuItemHasRoleMembership(
  item: MenuItem,
  profile: UserProfileResponse | undefined
): boolean {
  // no roles defined, means no authentication
  if (!item.roles || item.roles?.length === 0) return true;
  // no profile email, cannot have any role membership
  if (!profile || !profile?.email) return false;
  // compute the role membership on each item, first match returns true
  for (let i = 0; i < item.roles?.length; i++) {
    for (let j = 0; j < profile?.roles?.length; j++) {
      const role = roleNameToEnum(profile?.roles[j]);
      if (role === item.roles[i]) {
        return true;
      }
    }
  }
  return false;
}

// creates a copy of the menu group, but filtering only those which have permissions
export function filterMenuGroup(
  group: MenuGroup,
  profile: UserProfileResponse | undefined
): MenuGroup | null {
  const items = group.items.filter(x => menuItemHasRoleMembership(x, profile));
  if (items.length === 0) return null;
  return {
    icon: group.icon,
    initiallyOpened: group.initiallyOpened,
    label: group.label,
    items: items,
  };
}

// parses an array of MenuGroup vs the profile
export function glueMenus(
  groups: MenuGroup[],
  profile: UserProfileResponse | undefined
): MenuGroup[] {
  let result: MenuGroup[] = [];
  groups.forEach(group => {
    const filtered = filterMenuGroup(group, profile);
    if (!filtered) return;
    result.push(filtered);
  });
  return result.filter(x => x.items.length > 0);
}
