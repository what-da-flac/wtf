# wtf-ui

UI for web app player.

## Local Development

Create a copy of `.env.sample`

```
cp .env.sample .env
```

```
npm i
npm run start
```

Application runs at http://localhost:3000

## Prettier

We try to avoid sending changes to git related to code format, and we want the code to be
readable at the same time.

Run this command before you push code to git, to detect any format inconsistency

```
npm run format:check
```

And run this command to actually fix the code

```
npm run format
```

If your code needs formatting, GitHub checks won't pass.
