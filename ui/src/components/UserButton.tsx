import { IconChevronRight } from '@tabler/icons-react';

// Mantine :
import { Avatar, Group, rem, Text, UnstyledButton } from '@mantine/core';

// Models :
import { UserProfileResponse } from '../models/user';

// CSS :
import classes from './UserButton.module.css';

type params = {
  user: UserProfileResponse;
};

export function UserButton({ user }: params) {
  const userEmail = user.email.length > 24;

  const truncatedName = userEmail
    ? `${user.email.slice(0, 12)}...${user.email.slice(-10)}`
    : user.email;

  return (
    <UnstyledButton className={classes.user}>
      <Group>
        <Avatar src={user.picture} radius="xl" />
        <div style={{ flex: 1 }}>
          <Text size="sm" fw={500}>
            {user.name}
          </Text>
          <Text c="dimmed" size="xs">
            {truncatedName}
          </Text>
        </div>
        <IconChevronRight
          style={{ width: rem(14), height: rem(14) }}
          stroke={1.5}
        />
      </Group>
    </UnstyledButton>
  );
}
