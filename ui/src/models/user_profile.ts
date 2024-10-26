import { FileItem } from './file_item';
import { trimAll } from '../helpers/text_utils';

export enum UserGender {
  Hombre,
  Mujer,
  Otro,
}

export enum FileTypeNames {
  UNKNOWN_FILE_TYPE,
  INE_ID_BACK,
  INE_ID_FRONT,
  CURP_ID,
}

export class UserProfile {
  id: string;
  dob: Date;
  first_name: string;
  last_name: string;
  phone_number: string;
  gender: string;
  user_id: string;
  is_completed: boolean;

  constructor() {
    this.id = '';
    this.dob = new Date();
    this.first_name = '';
    this.last_name = '';
    this.phone_number = '';
    this.user_id = '';
    this.gender = '';
    this.is_completed = false;
  }
}

export function toUserProfile(v: any): UserProfile {
  if (!v) return new UserProfile();
  let gender: string = '';
  switch (v.gender) {
    case 'MALE':
      gender = UserGender[UserGender.Hombre];
      break;
    case 'FEMALE':
      gender = UserGender[UserGender.Mujer];
      break;
    default:
      gender = UserGender[UserGender.Otro];
      break;
  }
  return {
    ...v,
    dob: new Date(v['dob']),
    gender: gender,
  };
}

export function fromUserProfile(u: UserProfile): UserProfile {
  const clone = JSON.parse(JSON.stringify(u));
  switch (clone.gender) {
    case UserGender[UserGender.Hombre]:
      clone.gender = 'MALE';
      break;
    case UserGender[UserGender.Mujer]:
      clone.gender = 'FEMALE';
      break;
    default:
      clone.gender = 'OTHER';
      break;
  }
  return clone;
}

export function userProfileValidation() {
  return {
    phone_number: (value: string) =>
      trimAll(value).length === 0 ? 'Telefono es mandatorio' : null,
  };
}

export type UserProfileFile = {
  id?: string;
  file: FileItem;
  profile?: UserProfile;
  is_valid?: boolean;
  type: FileTypeNames;
};

export function toUserProfileFile(v: any): UserProfileFile {
  return {
    ...v,
  };
}
