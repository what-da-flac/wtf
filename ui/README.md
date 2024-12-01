# ui

UI for wtf suite.

## Local Development

Create a copy of `sample.env`

```
cp sample.env .env
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

If your code needs formatting, checks won't pass.
