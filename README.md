# churrer.xyz

[see what it looks like for yourself](https://churrer.xyz)

## Setup

### Frontend

install dependencies and start app using yarn

```bash
# install dependencies (doesn't include quarkus atm)
yarn install
# start nuxt frontend
yarn client:startdev
```

if you don't feel like running the yarn command to start the frontend:

don't worry, I got you covered. simply run the Intellij [StartClient](.run/StartClient.run.xml) run config.

> [!WARNING]
> `yarn install` and workspaces only work correctly in linux / macos. On windows this doesn't work.<br>
> it's a known [issue](https://github.com/yarnpkg/yarn/issues/4564) in yarn

for more nuxt stuff refer to [the client readme](./client/README.md)

### Backend

coming eventually
