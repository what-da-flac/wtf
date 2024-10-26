export class TorrentFile {
  id: string;
  file_name: string;
  file_size: string;

  constructor() {
    this.id = '';
    this.file_name = '';
    this.file_size = '';
  }
}

export function toTorrentFile(v: any): TorrentFile {
  if (!v) return new TorrentFile();
  return {
    ...v,
  };
}
