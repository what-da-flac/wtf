import { Link } from 'react-router-dom';

// Mantine :
import { Button, Loader, Table } from '@mantine/core';

// Models :
import { User } from '../models/user';

type Params = {
  isTableLoading?: boolean;
  onDelete?: any;
  rows: User[];
};

type RowParams = {
  onDelete: any;
  user: User;
};

function Row({ user, onDelete }: RowParams) {
  return (
    <Table.Tr key={user.id} className="table-row-container">
      <Table.Td>
        <Link to={`/users/${user.id}`}>{user.email}</Link>
      </Table.Td>
      <Table.Td>
        {user.image && <img alt="avatar" src={user.image} width={'10%'} />}
      </Table.Td>
      <Table.Td>{user.name}</Table.Td>
      <Table.Td>
        {`${user.last_login?.toLocaleDateString()} ${user.last_login?.toLocaleTimeString()}`}
      </Table.Td>
      {onDelete && (
        <Table.Td className="content-center">
          <Button
            type="button"
            variant="outline"
            onClick={() => onDelete(user)}
          >
            Quitar
          </Button>
        </Table.Td>
      )}
    </Table.Tr>
  );
}
function EmptyRow() {
  return (
    <Table.Tr className="table-row-container">
      <Table.Td></Table.Td>
      <Table.Td></Table.Td>
      <Table.Td className="content-center">No data</Table.Td>
      <Table.Td></Table.Td>
    </Table.Tr>
  );
}
function RowLoader() {
  return (
    <Table.Tr className="table-row-container">
      <Table.Td></Table.Td>
      <Table.Td></Table.Td>
      <Table.Td>
        <Loader className="table-loader" />
      </Table.Td>
      <Table.Td></Table.Td>
    </Table.Tr>
  );
}
export function UserTable({ rows, onDelete = null, isTableLoading }: Params) {
  return (
    <Table highlightOnHover withTableBorder className="table-container">
      <Table.Thead className="table-head-container">
        <Table.Tr>
          <Table.Th>Email</Table.Th>
          <Table.Th></Table.Th>
          <Table.Th>Name</Table.Th>
          <Table.Th>Last Login</Table.Th>
          {onDelete && <Table.Th className="content-center">Action</Table.Th>}
        </Table.Tr>
      </Table.Thead>
      <Table.Tbody>
        {isTableLoading ? (
          <RowLoader />
        ) : (
          <>
            {rows.length > 0 ? (
              rows.map((user: User) => (
                <Row key={user.id} user={user} onDelete={onDelete} />
              ))
            ) : (
              <EmptyRow />
            )}
          </>
        )}
      </Table.Tbody>
    </Table>
  );
}
