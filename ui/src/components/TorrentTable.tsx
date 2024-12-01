import { Link } from 'react-router-dom';
import { Avatar, Loader, Table } from '@mantine/core';
import { Torrent } from '../models/torrent';

type Params = {
  isTableLoading?: boolean;
  onDelete?: any;
  rows: Torrent[];
};

type RowParams = {
  torrent: Torrent;
};

function Row({ torrent }: RowParams) {
  const percentage = new Intl.NumberFormat('default', {
    style: 'percent',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0,
  });
  return (
    <Table.Tr key={torrent.id} className="table-row-container">
      <Table.Td>
        {`${torrent.created.toLocaleDateString()} ${torrent.created.toLocaleTimeString()}`}
      </Table.Td>
      <Table.Td>
        <Link to={`/torrents/${torrent.id}`}>{torrent.name}</Link>
      </Table.Td>
      <Table.Td style={{ textAlign: 'right' }}>{torrent.total_size}</Table.Td>
      <Avatar src={torrent.user.image} radius="xl" />
      <Table.Td>{torrent.status}</Table.Td>
      <Table.Td>
        {torrent.percent ? percentage.format(torrent.percent) : ''}
      </Table.Td>
      <Table.Td>{torrent.eta}</Table.Td>
    </Table.Tr>
  );
}

function EmptyRow() {
  return (
    <Table.Tr className="table-row-container">
      <Table.Td colSpan={4} className="content-center">
        No data
      </Table.Td>
    </Table.Tr>
  );
}

function RowLoader() {
  return (
    <Table.Tr className="table-row-container">
      <Table.Td>
        <Loader className="table-loader" />
      </Table.Td>
    </Table.Tr>
  );
}

export default function TorrentTable({
  rows,
  onDelete = null,
  isTableLoading,
}: Params) {
  return (
    <Table highlightOnHover withTableBorder className="table-container">
      <Table.Thead className="table-head-container">
        <Table.Tr>
          <Table.Th>Created</Table.Th>
          <Table.Th>Name</Table.Th>
          <Table.Th>Size</Table.Th>
          <Table.Th>Who</Table.Th>
          <Table.Th>Status</Table.Th>
          <Table.Th>Percent</Table.Th>
          <Table.Th>Eta</Table.Th>
          {onDelete && <Table.Th className="content-center">Action</Table.Th>}
        </Table.Tr>
      </Table.Thead>
      <Table.Tbody>
        {isTableLoading ? (
          <RowLoader />
        ) : (
          rows.map((row: Torrent) => <Row key={row.id} torrent={row} />)
        )}
      </Table.Tbody>
    </Table>
  );
}
