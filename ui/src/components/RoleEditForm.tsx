import { Button, Title, Group, TextInput } from '@mantine/core';
import { Role } from '../models/role';

type params = {
  onSubmit: any;
  onDelete?: any | undefined;
  form: any;
  legend: string;
  role: Role;
  showDelete?: boolean;
};

export function RoleEditForm({
  onSubmit,
  form,
  legend,
  onDelete,
  showDelete = false,
  role,
}: params) {
  return (
    <>
      <form
        onSubmit={form.onSubmit((data: any) => {
          onSubmit(data);
        })}
      >
        <Title>{legend}</Title>
        <br />
        <TextInput
          label="Nombre"
          placeholder="Nombre"
          disabled={true}
          {...form.getInputProps('name')}
        />
        <br />
        <TextInput
          label="Descripcion"
          placeholder="Descripcion"
          {...form.getInputProps('description')}
        />
        <br />
        <Group>
          <Button type="submit" variant="outline">
            Guardar
          </Button>
          {role.id && (
            <Group>
              {showDelete && (
                <Button type="button" variant="outline" onClick={onDelete}>
                  Eliminar Role
                </Button>
              )}
            </Group>
          )}
        </Group>
      </form>
    </>
  );
}
