FROM node:14

WORKDIR /apps
COPY package.json .
RUN npm install
COPY . .
EXPOSE 8080
CMD npm start