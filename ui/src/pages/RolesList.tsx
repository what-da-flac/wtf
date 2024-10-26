import React, { useEffect, useState } from 'react';

// Components :
import PreLoader from '../components/PreLoader';
import { RolesTable } from '../components/RolesTable';

// Helpers :
import { RoleList } from '../helpers/api';
import { isLoggedIn } from '../helpers/sso_service';

// Models :
import { Role } from '../models/role';

export function RolesList() {
  const [rows, setRows] = useState<Role[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(true);

  useEffect(() => {
    async function loadData() {
      try {
        const data = await RoleList();
        setRows(data);
      } catch (e) {
        console.error(e);
      } finally {
        setIsLoading(false);
      }
    }
    loadData();
  }, []);

  if (!isLoggedIn()) return null;

  return isLoading ? (
    <PreLoader />
  ) : (
    <React.Fragment>
      <h1>Role List</h1>
      <RolesTable rows={rows} />
    </React.Fragment>
  );
}
