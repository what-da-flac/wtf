import { useDisclosure } from '@mantine/hooks';
import { UseFormReturnType } from '@mantine/form';
import { Box, Button, Group, TextInput } from '@mantine/core';
import { Torrent } from '../models/torrent';
import React from 'react';

type params = {
  onSubmit: any;
  onDelete?: any | undefined;
  form: UseFormReturnType<Torrent>;
  torrent: Torrent;
  canEdit: boolean;
};

export default function TorrentEditForm({
  onSubmit,
  form,
  onDelete,
  torrent,
  canEdit,
}: params) {
  const [visible, { toggle }] = useDisclosure(false);
  return (
    <>
      <form
        onSubmit={form.onSubmit(async (data: any) => {
          toggle();
          await onSubmit(data);
          toggle();
        })}
        className="form-wrapper"
      >
        <Box pos="relative">
          <br />
          <TextInput
            label="Name"
            size="md"
            type="text"
            placeholder="File Name"
            disabled={true}
            {...form.getInputProps('name')}
          />
          <br />
          <TextInput
            label="Created"
            size="md"
            type="text"
            disabled={true}
            {...form.getInputProps('created')}
          />
          <br />
          <TextInput
            label="Status"
            size="md"
            type="text"
            disabled={true}
            {...form.getInputProps('status')}
          />
          <br />
          <TextInput
            label="User"
            size="md"
            type="text"
            disabled={true}
            {...form.getInputProps('user.name')}
          />
          <br />
          <TextInput
            label="Total Size"
            size="md"
            type="text"
            disabled={true}
            {...form.getInputProps('total_size')}
          />
          <br />
          <TextInput
            label="Magnet Link"
            size="md"
            type="text"
            placeholder="magnet link"
            disabled={true}
            {...form.getInputProps('magnet_link')}
          />
          <br />
          <Group>
            {canEdit && (
              <Button type="submit" variant="outline">
                Download
              </Button>
            )}
            {torrent?.id && (
              <Group>
                {torrent.id && (
                  <Button type="button" variant="outline" onClick={onDelete}>
                    Delete
                  </Button>
                )}
              </Group>
            )}
          </Group>
        </Box>
      </form>
    </>
  );
}
