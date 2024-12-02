import instance from './axios';

// Models :
import { Role } from '../models/role';
import { toUser } from '../models/user';
import {
  toUserProfile,
  toUserProfileFile,
  UserProfile,
  UserProfileFile,
} from '../models/user_profile';
import { toMovie } from '../models/movie';
import { toTorrent } from '../models/torrent';

const defaultMockedTimeout = 100;

function mockRequest(payload: any): Promise<any> {
  return new Promise((resolve, reject) => {
    // Simulate network delay
    setTimeout(() => {
      resolve(payload);
    }, defaultMockedTimeout);
  });
}

async function ApiMovieList() {
  const data = [
    {
      id: '1',
      title: 'Alien Covenant',
      image_url:
        'https://resizing.flixster.com/FqkI1Wx0fwZGwm67MH0DiRfFZQE=/206x305/v2/https://resizing.flixster.com/v7kKyi8CdrLtNDAtsW_MErYm-sk=/ems.cHJkLWVtcy1hc3NldHMvbW92aWVzLzA2MmY2NzJhLWQ1MmUtNDUyOS05OGZiLWU4ZDBmYTE3YjQ5Yy53ZWJw',
      description:
        'Bound for a remote planet on the far side of the galaxy, members (Katherine Waterston, Billy Crudup) of the colony ship Covenant discover what they think to be an uncharted paradise. While there, they meet David (Michael Fassbender), the synthetic survivor of the doomed Prometheus expedition. The mysterious world soon turns dark and dangerous when a hostile alien life-form forces the crew into a deadly fight for survival.',
    },
    {
      id: '2',
      title: 'Scarface',
      image_url:
        'https://resizing.flixster.com/yFsJgH0__B4GOIA_fAADXeaK9nI=/206x305/v2/https://resizing.flixster.com/xqpbVMOd3KhO_ryiETThKCwstxk=/ems.cHJkLWVtcy1hc3NldHMvbW92aWVzLzUwNzhhMWVhLWQ0YmItNDRiMi04MzE1LTc4ZjBiYWVhOWRmZi5qcGc=',
      description:
        'After getting a green card in exchange for assassinating a Cuban government official, Tony Montana (Al Pacino) stakes a claim on the drug trade in Miami. Viciously murdering anyone who stands in his way, Tony eventually becomes the biggest drug lord in the state, controlling nearly all the cocaine that comes through Miami. But increased pressure from the police, wars with Colombian drug cartels and his own drug-fueled paranoia serve to fuel the flames of his eventual downfall.',
    },
    {
      id: '3',
      title: 'Secret Admirer',
      image_url:
        'https://resizing.flixster.com/QPG9HzIw9PDX9YFVE1Tr4Lrl6c0=/206x305/v2/https://resizing.flixster.com/-XZAfHZM39UwaGJIFWKAE8fS0ak=/v3/t/assets/p8776_p_v8_aa.jpg',
      description:
        "Michael Ryan (C. Thomas Howell) has been in love with Deborah Ann (Kelly Preston) for as long as he can remember, not realizing that his female best friend, Toni Williams (Lori Loughlin), is secretly pining for him. He leaves Deborah Ann an anonymous love letter and soon the two begin a relationship -- while a jealous Toni is forced to watch from afar. Meanwhile, the letter is found by Deborah Ann's parents and sparks suspicion and accusations of philandering among them and their neighbors.",
    },
    {
      id: '4',
      title: '12 Monkeys',
      image_url:
        'https://resizing.flixster.com/oCn_sq0co1EkH17SIbrwmRVQ4I8=/206x305/v2/https://resizing.flixster.com/-XZAfHZM39UwaGJIFWKAE8fS0ak=/v3/t/assets/p17517_p_v8_au.jpg',
      description:
        "Traveling back in time isn't simple, as James Cole (Bruce Willis) learns the hard way. Imprisoned in the 2030s, James is recruited for a mission that will send him back to the 1990s. Once there, he's supposed to gather information about a nascent plague that's about to exterminate the vast majority of the world's population. But, aside from the manic Jeffrey (Brad Pitt), he gets little in the way of cooperation, not least from medical gatekeepers like Dr. Kathryn Railly (Madeleine Stowe).",
    },
    {
      id: '5',
      title: 'Taste the Blood of Dracula',
      image_url:
        'https://resizing.flixster.com/08H2zUCnp4uEqw43791O5DJECGU=/206x305/v2/https://resizing.flixster.com/-XZAfHZM39UwaGJIFWKAE8fS0ak=/v3/t/assets/p7842_p_v8_aa.jpg',
      description:
        "Victorian thrill-seekers kill the vampire's (Christopher Lee) helper; he seduces their daughters to get even.",
    },
    {
      id: '7',
      title: 'Back to the Future',
      image_url:
        'https://resizing.flixster.com/Cw_w2trVGZKIvCoR9SCRe2CL3ro=/206x305/v2/https://resizing.flixster.com/-XZAfHZM39UwaGJIFWKAE8fS0ak=/v3/t/assets/p8717_p_v8_ac.jpg',
      description:
        "In this 1980s sci-fi classic, small-town California teen Marty McFly (Michael J. Fox) is thrown back into the '50s when an experiment by his eccentric scientist friend Doc Brown (Christopher Lloyd) goes awry. Traveling through time in a modified DeLorean car, Marty encounters young versions of his parents (Crispin Glover, Lea Thompson), and must make sure that they fall in love or he'll cease to exist. Even more dauntingly, Marty has to return to his own time and save the life of Doc Brown.",
    },
  ];
  const res = await mockRequest(data);
  return (res || []).map((x: any) => toMovie(x));
}

type TorrentListParams = {
  limit: number;
  offset?: number;
  status?: string;
  sort_field: string;
  sort_direction: string;
};

async function ApiTorrentList(params: TorrentListParams) {
  const res = await instance.get(`/v1/torrents`, { params });
  return (res.data || []).map((x: any) => toTorrent(x));
}

async function ApiTorrentLoad(id: string) {
  const res = await instance.get(`/v1/torrents/${id}`);
  return res.data;
}

async function ApiTorrentPost(data: any) {
  const res = await instance.post(`/v1/torrents/magnets`, data);
  return res.data;
}

async function ApiTorrentDownload(id: string) {
  const res = await instance.post(`/v1/torrents/${id}/download`, {});
  return res.data;
}

async function ApiTorrentStatuses() {
  const res = await instance.get(`/v1/torrents/statuses`);
  return res.data;
}

async function ApiTorrentUpdateStatus(id: string, status: string) {
  const res = await instance.put(`/v1/torrents/${id}/status/${status}`);
  return res.data;
}

async function ApiTorrentDelete(id: string) {
  const res = await instance.delete(`/v1/torrents/${id}`);
  return res.data;
}

/////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////

/////////////////////////////////////////////////////////////
// Role
/////////////////////////////////////////////////////////////

async function RoleList() {
  const res = await instance.get(`/v1/roles`);
  return (res.data || []).map((x: any) => x as Role);
}

async function RolePut(role: Role) {
  const res = await instance.put(`/v1/roles/${role.id}`, role);
  return res.data as Role;
}

async function RoleGet(id: string) {
  const res = await instance.get(`/v1/roles/${id}`);
  return res.data as Role;
}

async function RoleAddUser(roleId: string, userId: string) {
  const res = await instance.put(`/v1/roles/${roleId}/users/${userId}`);
  return res.data;
}

async function RoleRemoveUser(roleId: string, userId: string) {
  const res = await instance.delete(`/v1/roles/${roleId}/users/${userId}`);
  return res.data;
}

async function RolesInUser(userId: string) {
  const res = await instance.get(`/v1/users/${userId}/roles`);
  return (res.data || []).map((x: any) => x as Role);
}

async function UsersInRole(roleId: string) {
  const res = await instance.get(`/v1/roles/${roleId}/users`);
  return (res.data || []).map((x: any) => toUser(x));
}

/////////////////////////////////////////////////////////////
// User Profile
/////////////////////////////////////////////////////////////

async function UserProfilePost(profile: UserProfile) {
  const res = await instance.post(`/v1/user-profile`, profile);
  return toUserProfile(res.data);
}

async function UserProfileLoad() {
  const res = await instance.get(`/v1/user-profile`);
  return toUserProfile(res.data);
}

async function UserProfileFilesLoad() {
  const res = await instance.get(`/v1/user-profile/files`);
  return (res.data || []).map((x: any) => toUserProfileFile(x));
}

async function UserProfileFileSave(file: UserProfileFile) {
  const res = await instance.put(`/v1/user-profile/files`, file);
  return toUserProfileFile(res.data);
}

/////////////////////////////////////////////////////////////
// User
/////////////////////////////////////////////////////////////

export type UserListParams = {
  offset?: number;
  limit?: number;
  emails?: string[];
  email_match: string;
};

async function PostUserList(params: UserListParams) {
  const res = await instance.post(`/v1/user-list`, params);
  return (res.data || []).map((x: any) => toUser(x));
}

async function UserWhoAmi() {
  const res = await instance.get(`/v1/users/whoami`);
  return toUser(res.data);
}

async function UserLoad(id: string) {
  const res = await instance.get(`/v1/users/${id}`);
  return toUser(res.data);
}

export {
  ApiMovieList,
  ApiTorrentDelete,
  ApiTorrentDownload,
  ApiTorrentList,
  ApiTorrentLoad,
  ApiTorrentPost,
  ApiTorrentStatuses,
  ApiTorrentUpdateStatus,
  RoleList,
  RolePut,
  RoleGet,
  RoleAddUser,
  RoleRemoveUser,
  RolesInUser,
  UsersInRole,
  UserProfilePost,
  UserProfileLoad,
  UserProfileFilesLoad,
  UserProfileFileSave,
  UserWhoAmi,
  PostUserList,
  UserLoad,
};
