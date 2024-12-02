import React, { useEffect, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { useForm } from '@mantine/form';
import PreLoader from '../components/PreLoader';
import { Torrent, toTorrent } from '../models/torrent';
import { TorrentFile } from '../models/torrent_file';
import TorrentEditForm from '../components/TorrentEditForm';
import {
  ApiTorrentDownload,
  ApiTorrentLoad,
  ApiTorrentUpdateStatus,
} from '../helpers/api';
import { TorrentFilesTable } from '../components/TorrentFilesTable';

export function TorrentEdit() {
  const id = useParams().id || '';
  const [torrent, setTorrent] = useState<Torrent>(new Torrent());
  const [torrentFiles, setTorrentFiles] = useState<TorrentFile[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const form = useForm<Torrent>({
    initialValues: torrent,
  });
  const navigate = useNavigate();

  useEffect(() => {
    loadTorrent(id);
  }, []);

  async function loadTorrent(id: string) {
    try {
      const data = await ApiTorrentLoad(id);
      setTorrentFiles(data.files);
      const t = toTorrent(data);
      setTorrent(t);
      form.setValues(data);
    } catch (e) {
      console.error(e);
    } finally {
      setIsLoading(false);
    }
  }

  async function downloadTorrent(id: string) {
    await ApiTorrentDownload(id);
    navigate('/torrents');
  }

  async function onReset() {
    await ApiTorrentUpdateStatus(id, 'pending');
    navigate('/torrents');
  }

  return isLoading ? (
    <PreLoader />
  ) : (
    <React.Fragment>
      <h1>{torrent.name}</h1>
      <TorrentEditForm
        form={form}
        torrent={torrent}
        canEdit={torrent.status === 'pending'}
        onDelete={() => alert('TODO: delete')}
        onReset={() => onReset()}
        onSubmit={() => downloadTorrent(id)}
      />
      <hr />
      <h3>Files</h3>
      <TorrentFilesTable rows={torrentFiles} />
    </React.Fragment>
  );
}
