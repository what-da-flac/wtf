import { FormEvent, useEffect, useState } from 'react';
import { IconPlus } from '@tabler/icons-react';
import { Button, Select } from '@mantine/core';
import PreLoader from '../components/PreLoader';
import { Torrent } from '../models/torrent';
import TorrentTable from '../components/TorrentTable';
import { ApiTorrentList, ApiTorrentStatuses } from '../helpers/api';
import { Link } from 'react-router-dom';

type StatusData = {
  label: string;
  value: string;
};

export default function TorrentsList() {
  const [match, setMatch] = useState('');
  const [offset, setOffset] = useState(0);
  const [rows, setRows] = useState<Torrent[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const [isTableLoading, setIsTableLoading] = useState<boolean>(false);
  const [statusData, setStatusData] = useState<StatusData[]>([]);
  const [status, setStatus] = useState<string>('');
  const limit = 100
  useEffect(() => {
    loadData(status);
  }, []);

  async function loadData(s: string) {
    try {
      setIsTableLoading(true);
      const st = !s ? null : s;
      const data = await ApiTorrentList({ limit, status: st ,sort_field: 'created',sort_direction: 'asc',});
      setRows(data);
      let stData = await ApiTorrentStatuses();
      stData.sort();
      const items: StatusData[] = [
        {
          label: 'All',
          value: '',
        },
      ];
      stData.forEach((x: string) =>
        items.push({
          label: x,
          value: x,
        })
      );
      setStatusData(items);
      setIsTableLoading(false);
    } catch (e) {
      console.error(e);
      setIsTableLoading(false);
    } finally {
      setIsLoading(false);
      setIsTableLoading(false);
    }
  }

  function handleSubmit(e: FormEvent) {
    e.preventDefault();
  }

  return isLoading ? (
    <PreLoader />
  ) : (
    <div className="user-list-container">
      <h1>Torrents</h1>
      <form className="table-search" onSubmit={e => handleSubmit(e)}>
        <Select
          className="select-input"
          label="Status Filter"
          placeholder="select any status"
          data={statusData}
          onChange={(e: any) => loadData(e)}
        />
      </form>
      <TorrentTable rows={rows} isTableLoading={isTableLoading} />
      <div className="action-buttons">
        <Link to="/torrents/new">
          <Button
            className="link-button"
            size="md"
            type="button"
            variant="outline"
            disabled={false}
          >
            Add Torrent <IconPlus className="icon" />
          </Button>
        </Link>
      </div>
    </div>
  );
}
