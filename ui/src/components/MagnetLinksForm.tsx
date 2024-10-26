import { useDisclosure } from '@mantine/hooks';
import { UseFormReturnType } from '@mantine/form';
import { Box, Button, Group, Textarea } from '@mantine/core';
import { Torrent } from '../models/torrent';

type params = {
  onSubmit: any;
  form: UseFormReturnType<Torrent>;
};

export default function MagnetLinksForm({ onSubmit, form }: params) {
  const [visible, { toggle }] = useDisclosure(false);
  return (
    <>
      <h1>Bulk create torrents</h1>
      <form
        onSubmit={form.onSubmit(async (data: any) => {
          toggle();
          // debugger // eslint-disable-line no-debugger
          await onSubmit(data.urls.split('\n').filter((item: any) => item));
          toggle();
        })}
        className="form-wrapper"
      >
        <Box pos="relative">
          <br />
          <Textarea
            label="Magnet Links"
            size="md"
            placeholder="Paste as many magnet links as you like, they will be queued for processing"
            {...form.getInputProps('urls')}
          ></Textarea>
          <br />
          <Group>
            <Button type="submit" variant="outline">
              Process
            </Button>
          </Group>
        </Box>
      </form>
    </>
  );
}
