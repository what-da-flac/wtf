import { Link } from 'react-router-dom';
import { Edit, Map, User } from 'tabler-icons-react';

// Mantine :
import { Badge, Button, Card, Flex, Group, Image, Text } from '@mantine/core';

// Models :

// Helpers :
import { trimAndCapitalize, trimBigText } from '../helpers/text_utils';
import { Movie } from '../models/movie';

type Params = {
  m: Movie;
};

const maxHeadingLength = 30;
const maxDescriptionLength = 150;

export default function MovieCard({ m }: Params) {
  return (
    <Card
      withBorder
      shadow="md"
      radius="lg"
      padding="lg"
      className="campaign-card-container"
    >
      <Card.Section>
        <Image alt="logo" height={200} src={m.image_url} />
      </Card.Section>
      <Flex gap="sm" mt="sm">
        <Map />
        <Text fw={700}>
          {m.title && trimAndCapitalize(m.title, maxHeadingLength)}
        </Text>
      </Flex>
      <Text size="sm" mt="sm" c="dimmed" className="campaign-card-description">
        {m.description ? (
          trimBigText(m.description, maxDescriptionLength)
        ) : (
          <>No description</>
        )}
      </Text>
      <Group justify="space-between" mt="md">
        <Flex gap="5px" align="center">
          <User size={20} color="#4dabf7" />
          <Text size="lg" pt="3px">
            Uploaded by Mau
          </Text>
        </Flex>
        <Badge color="pink" variant="light">
          Draft
        </Badge>
      </Group>
      <Flex gap="md">
        <Button
          fullWidth
          mt="sm"
          radius="lg"
          variant="outline"
          component={Link}
          to={`/movies/${m.id}`}
        >
          Edit
          <Edit size={20} />
        </Button>
      </Flex>
    </Card>
  );
}
