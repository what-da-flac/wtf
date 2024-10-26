import ReactPlayer from 'react-player';
import { Button, Group } from '@mantine/core';
import { IconArrowNarrowRight } from '@tabler/icons-react';

type Params = {
  userQuiz: any;
  nextStep: any;
};
export default function VideoPlayer({ userQuiz, nextStep }: Params) {
  return (
    <div className="form-wrapper video-player-continer">
      <ReactPlayer
        controls
        width="100%"
        height="calc(100vh - 420px)"
        url={userQuiz.quiz.video_url}
      />
      <Group mt="xl">
        <Button
          size="md"
          variant="outline"
          onClick={nextStep}
          className="button"
        >
          Próxima <IconArrowNarrowRight className="icon" />
        </Button>
      </Group>
    </div>
  );
}
