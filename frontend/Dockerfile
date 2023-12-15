FROM node:18-slim
RUN mkdir /app
COPY package.json /app/
WORKDIR /app
COPY . ./

# ENV MY_FIRST_ENV=helloworld

RUN npm install
RUN npm run build
EXPOSE 4000
CMD ["npm", "run","start"]