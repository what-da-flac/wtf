import { toDate } from '../helpers/text_utils';
import { FileTypeNames } from './user_profile';

export class FileItem {
  id: string;
  created: Date;
  reference_id: string;
  key: string;
  name: string;
  size: number;
  type: string;
  content_type: string;

  constructor() {
    this.id = '';
    this.created = new Date();
    this.reference_id = '';
    this.key = '';
    this.name = '';
    this.size = 0;
    this.type = '';
    this.content_type = '';
  }
}

export function toFileItem(v: any): FileItem {
  if (!v) return new FileItem();
  return {
    ...v,
    last_login: toDate(new Date(v['created'])),
  };
}

export type FileType = {
  description: string;
  name: FileTypeNames;
};
