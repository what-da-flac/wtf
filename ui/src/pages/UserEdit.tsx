import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { IconCornerDownLeft, IconCornerDownRight } from '@tabler/icons-react';

// Mantine :
import { Grid } from '@mantine/core';
import { useForm } from '@mantine/form';

// Components :
import PreLoader from '../components/PreLoader';
import { notifyErrResponse } from '../components/Errors';
import { UserEditForm } from '../components/UserEditForm';
import { RolesSimpleTable } from '../components/RolesSimpleTable';

// Models :
import { Role } from '../models/role';
import { User } from '../models/user';

// Helpers :
import {
  UserLoad,
  RoleList,
  RolesInUser,
  RoleAddUser,
  RoleRemoveUser,
} from '../helpers/api';

export function UserEdit() {
  const id = useParams().id || '';
  const [user, setUser] = useState<User>(new User());
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const [selectedRoles, setSelectedRoles] = useState<Role[]>([]);
  const [unselectedRoles, setUnselectedRoles] = useState<Role[]>([]);
  const [addRoleLoading, setAddRoleLoading] = useState<boolean>(false);
  const [removeRoleLoading, setRemoveRoleLoading] = useState<boolean>(false);
  const form = useForm<User>({
    initialValues: user,
  });

  useEffect(() => {
    loadUser(id);
    loadRoles();
  }, []);

  async function loadRoles() {
    try {
      const selected = await RolesInUser(id);
      const allRoles = await RoleList();
      const unselected = allRoles.filter(
        (x: Role) => !selected.find((r: Role) => r.id === x.id)
      );
      setSelectedRoles(selected);
      setUnselectedRoles(unselected);
    } catch (err) {
      await notifyErrResponse(err);
    } finally {
      setIsLoading(false);
    }
  }

  async function loadUser(id: string) {
    try {
      const data = await UserLoad(id);
      setUser(data);
      form.setValues(data);
    } catch (e) {
      console.error(e);
    } finally {
      setIsLoading(false);
    }
  }

  async function addRole(r: Role) {
    try {
      setAddRoleLoading(true);
      await RoleAddUser(r.id, user.id);
      await loadRoles();
      setAddRoleLoading(false);
    } catch (e) {
      notifyErrResponse(e);
      setAddRoleLoading(false);
    }
  }

  async function removeRole(r: Role) {
    try {
      setRemoveRoleLoading(true);
      await RoleRemoveUser(r.id, user.id);
      await loadRoles();
      setRemoveRoleLoading(false);
    } catch (e) {
      notifyErrResponse(e);
      setRemoveRoleLoading(false);
    }
  }

  return isLoading ? (
    <PreLoader />
  ) : (
    <React.Fragment>
      <h1>{user.name}</h1>
      <UserEditForm form={form} user={user} />
      <h1>Role Membership</h1>
      <Grid className="form-wrapper roles-table">
        <Grid.Col span={6}>
          <React.Fragment>
            <h4>Available</h4>
            <RolesSimpleTable
              icon={
                <IconCornerDownRight
                  cursor="pointer"
                  color="var(--success-color)"
                />
              }
              loading={addRoleLoading}
              rows={unselectedRoles}
              onClick={addRole}
            />
          </React.Fragment>
        </Grid.Col>
        <Grid.Col span={6}>
          <React.Fragment>
            <h4>Selected</h4>
            <RolesSimpleTable
              icon={
                <IconCornerDownLeft
                  cursor="pointer"
                  color="var(--danger-color)"
                />
              }
              loading={removeRoleLoading}
              rows={selectedRoles}
              onClick={removeRole}
            />
          </React.Fragment>
        </Grid.Col>
      </Grid>
    </React.Fragment>
  );
}
