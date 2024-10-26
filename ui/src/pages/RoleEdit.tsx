import React, { useEffect, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';

// Mantine :
import { useForm } from '@mantine/form';

// Components :
import { UserTable } from '../components/UserTable';
import { notifyErrResponse } from '../components/Errors';
import { RoleEditForm } from '../components/RoleEditForm';

// Helpers :
import { RoleGet, RolePut, RoleRemoveUser, UsersInRole } from '../helpers/api';

// Models :
import { User } from '../models/user';
import { Role, roleValidation } from '../models/role';

export function RoleEdit() {
  const id = useParams().id || '';
  const navigate = useNavigate();
  const [role, setRole] = useState<Role>(new Role());
  const [users, setUsers] = useState<User[]>([]);

  const form = useForm<Role>({
    initialValues: role,
    validate: roleValidation(),
  });

  async function loadUsers() {
    try {
      const data = await UsersInRole(id);
      setUsers(data);
    } catch (err) {
      await notifyErrResponse(err);
    }
  }
  async function loadRole(id: string) {
    try {
      const data = await RoleGet(id);
      setRole(data);
      form.setValues(data);
    } catch (e) {
      console.error(e);
    }
  }
  useEffect(() => {
    loadRole(id);
    loadUsers();
  }, []);

  async function onSubmit(data: Role) {
    try {
      await RolePut(data);
      const returnURL = `/roles`;
      navigate(returnURL);
    } catch (err) {
      await notifyErrResponse(err);
    }
  }

  async function onRemove(user: User) {
    try {
      if (
        !window.confirm(
          `Seguro de quitar al usuario ${user.name} del role ${role.name}?`
        )
      )
        return;
      await RoleRemoveUser(role.id, user.id);
      setUsers(users.filter(x => x.id !== user.id));
    } catch (err) {
      await notifyErrResponse(err);
    }
  }

  return (
    <React.Fragment>
      <RoleEditForm
        onSubmit={onSubmit}
        form={form}
        legend={role.name}
        role={role}
      />
      <br />
      <UserTable rows={users} onDelete={onRemove} />
    </React.Fragment>
  );
}
