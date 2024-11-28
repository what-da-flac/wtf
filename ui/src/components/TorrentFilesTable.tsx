import React from 'react';

// Mantine :
import { Table } from '@mantine/core';

// Models :
import { Role } from '../models/role';
import { Torrent } from '../models/torrent';
import { TorrentFile } from '../models/torrent_file';

type Params = {
  rows: TorrentFile[];
};

type RowParams = {
  file: TorrentFile;
};

function Row({ file }: RowParams) {
  return (
    <Table.Tr key={file.id} className="table-row-container">
      <Table.Td>{file.file_name}</Table.Td>
      <Table.Td>{file.file_size}</Table.Td>
    </Table.Tr>
  );
}

function EmptyTable() {
  return (
    <Table.Tr>
      <Table.Td></Table.Td>
      <Table.Td>No data</Table.Td>
    </Table.Tr>
  );
}

export function TorrentFilesTable({ rows }: Params) {
  return (
    <Table highlightOnHover withTableBorder className="table-container">
      <Table.Thead className="table-head-container">
        <Table.Tr>
          <Table.Th>Name</Table.Th>
          <Table.Th>Size</Table.Th>
        </Table.Tr>
      </Table.Thead>
      <Table.Tbody>
        {rows && rows.length > 0 ? (
          <React.Fragment>
            {rows.map((f: TorrentFile) => (
              <Row key={f.id} file={f} />
            ))}
          </React.Fragment>
        ) : (
          <EmptyTable />
        )}
      </Table.Tbody>
    </Table>
  );
}
