import { useEffect, useState } from 'react';
import { IconArrowNarrowLeft, IconArrowNarrowRight } from '@tabler/icons-react';
import { Search } from 'tabler-icons-react';

// Mantine :
import { Button, TextInput } from '@mantine/core';

// Components :
import PreLoader from '../components/PreLoader';
import { UserTable } from '../components/UserTable';

// Helpers :
import { isLoggedIn } from '../helpers/sso_service';
import { PostUserList, UserListParams } from '../helpers/api';

// Models :
import { User } from '../models/user';

export function UsersList() {
  const [match, setMatch] = useState('');
  const [offset, setOffset] = useState(0);
  const [rows, setRows] = useState<User[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const [isTableLoading, setIsTableLoading] = useState<boolean>(false);
  const limit = 10;

  useEffect(() => {
    loadData();
  }, [offset]);

  async function loadData() {
    try {
      setIsTableLoading(true);
      const params: UserListParams = {
        email_match: match,
      };
      const data = await PostUserList(params);
      setRows(data);
      setIsTableLoading(false);
    } catch (e) {
      console.error(e);
      setIsTableLoading(false);
    } finally {
      setIsLoading(false);
      setIsTableLoading(false);
    }
  }

  const handleSubmit = async (e: any) => {
    e.preventDefault();
    await loadData();
  };

  const shouldBackButtonDisabled = offset <= 0 && match.length === 0;
  if (!isLoggedIn()) return null;

  return isLoading ? (
    <PreLoader />
  ) : (
    <div className="user-list-container">
      <h1>User List</h1>
      <form className="table-search" onSubmit={e => handleSubmit(e)}>
        <TextInput
          className="search-input"
          value={match}
          rightSection={
            <Search
              size={18}
              strokeWidth={2}
              color="#fff"
              cursor="pointer"
              onClick={e => handleSubmit(e)}
            />
          }
          label="Email Search"
          placeholder="john.doe@example.com"
          description="Type name or email to look up users, then hit enter"
          onChange={e => setMatch(e.target.value)}
        />
      </form>
      <UserTable rows={rows} isTableLoading={isTableLoading} />
      <div className="action-buttons">
        <Button
          className="back-button"
          size="md"
          type="button"
          variant="outline"
          disabled={shouldBackButtonDisabled}
          onClick={() => {
            setOffset(offset - limit);
            setMatch('');
          }}
        >
          <IconArrowNarrowLeft className="icon" />
          Previous
        </Button>
        <Button
          className="next-button"
          size="md"
          type="button"
          variant="outline"
          disabled={rows.length < limit}
          onClick={async () => {
            setOffset(offset + limit);
          }}
        >
          Next <IconArrowNarrowRight className="icon" />
        </Button>
      </div>
    </div>
  );
}
