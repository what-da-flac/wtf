import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useForm } from '@mantine/form';
import { notifyErrResponse } from '../components/Errors';
import MagnetLinksForm from '../components/MagnetLinksForm';
import { ApiTorrentPost } from '../helpers/api';

export default function TorrentBulkForm() {
  const navigate = useNavigate();
  const [canEdit, setCanEdit] = useState<boolean>(true);
  const form = useForm<any>({
    initialValues: {
      urls: '',
    },
    validate: {
      urls: (value: string) => (!value ? 'type at least one url' : null),
    },
  });

  async function onSubmit(data: string[]) {
    try {
      const returnURL = '/torrents';
      setCanEdit(false);
      await ApiTorrentPost({
        urls: data,
      });
      navigate(returnURL);
    } catch (err) {
      await notifyErrResponse(err);
    } finally {
      setCanEdit(true);
    }
  }

  return <MagnetLinksForm onSubmit={onSubmit} form={form} />;
}
