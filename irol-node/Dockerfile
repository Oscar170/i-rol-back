# Create image based on the official Node LTS image from dockerhub
FROM node:8.9.4

# Create a directory where our app will be placed
RUN mkdir -p /usr/src/app

ENV PORT 4000

# Change directory so that our commands run inside this new directory
WORKDIR /usr/src/app

# Copy dependency definitions
COPY package.json /usr/src/app

# Install node dependecies
RUN npm install

# Get all the code needed to run the app
COPY . /usr/src/app

CMD ["npm", "start"]
