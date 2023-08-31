FROM node:18-alpine

WORKDIR /app

COPY node_modules node_modules
COPY package.json package-lock.json ./
COPY dist dist
COPY public public

# ON_DEPLOY_ADD_ENV_VARS
ENV NODE_ENV=production

RUN npm install

COPY . ./

CMD ["npm", "run", "dev"]
