import { useState } from 'react';
// Mantine :
import { Loader, Table } from '@mantine/core';

// Models :
import { Role } from '../models/role';

// Helpers :
import { capitalizeAllWords } from '../helpers/text_utils';

type Params = {
  rows: Role[];
  onClick: any;
  icon: any;
  loading: boolean;
};

type RowParams = {
  role: Role;
  onClick: any;
  icon: any;
  loading: boolean;
};

function Row({ icon, loading, role, onClick }: RowParams) {
  const [clickedRole, setClickedRole] = useState(null);

  const handleClick = (clickedRole: any) => {
    setClickedRole(clickedRole);
    onClick(clickedRole);
  };

  return (
    <Table.Tr key={role.id} className="table-row-container">
      <Table.Td>{capitalizeAllWords(role.name)}</Table.Td>
      <Table.Td className="content-center" onClick={() => handleClick(role)}>
        <div className="flex-column">
          <div>
            {loading && role.id === clickedRole?.id && <Loader size="sm" />}
          </div>
          {icon}
        </div>
      </Table.Td>
    </Table.Tr>
  );
}

export function RolesSimpleTable({ icon, loading, rows, onClick }: Params) {
  return (
    <Table highlightOnHover withTableBorder className="table-container">
      <Table.Thead className="table-head-container">
        <Table.Tr>
          <Table.Th>Role Name</Table.Th>
          <Table.Th></Table.Th>
        </Table.Tr>
      </Table.Thead>
      <Table.Tbody>
        {rows.map((role: Role) => (
          <Row
            loading={loading}
            key={role.id}
            role={role}
            onClick={onClick}
            icon={icon}
          />
        ))}
      </Table.Tbody>
    </Table>
  );
}
