# Javascript


## Install / Upgrade
```shell
# NVM (node + npm)
# swap `v0.40.3` for whatever latest version of nvm-sh/nvm has been released
# to upgrade, run same command with a newer version
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.3/install.sh | bash
source ~/.zshrc

nvm install --lts
nvm use --lts

# PNPM
npm install -g pnpm
source ~/.zshrc
pnpm setup
source ~/.zshrc
```

## NestJS
```shell
# NestJS
pnpm install -g @nestjs/cli
source ~/.zshrc

# create scaffold project
nest new my-project-name
cd my-project-name
pnpm install
# if prompted
pnpm approve-builds
# run tests
pnpm run test
# run app
pnpm run start
```
