import { trimAll } from './text_utils';

test('trimAll', () => {
  const d = 'test  ';
  const res = trimAll(d);
  const expected = 'test';
  expect(res).toEqual(expected);
});
