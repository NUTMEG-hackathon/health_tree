FROM node:12.14.1

WORKDIR /app/vue-project
RUN npm install -g @vue/cli
COPY ./vue-project/ /app/vue-project
RUN npm install
RUN npm install --save axios
RUN npm install --save vuex-persistedstate
