import React from 'react';

// Mantine :
import { Table } from '@mantine/core';

// Models :
import { Role } from '../models/role';

type Params = {
  rows: Role[];
};

type RowParams = {
  role: Role;
};

function Row({ role }: RowParams) {
  return (
    <Table.Tr key={role.id} className="table-row-container">
      <Table.Td>{role.name}</Table.Td>
      <Table.Td>{role.description}</Table.Td>
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

export function RolesTable({ rows }: Params) {
  return (
    <Table highlightOnHover withTableBorder className="table-container">
      <Table.Thead className="table-head-container">
        <Table.Tr>
          <Table.Th>Name</Table.Th>
          <Table.Th>Description</Table.Th>
        </Table.Tr>
      </Table.Thead>
      <Table.Tbody>
        {rows.length > 0 ? (
          <React.Fragment>
            {rows.map((role: Role) => (
              <Row key={role.id} role={role} />
            ))}
          </React.Fragment>
        ) : (
          <EmptyTable />
        )}
      </Table.Tbody>
    </Table>
  );
}
