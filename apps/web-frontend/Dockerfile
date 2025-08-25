# Use a Node.js base image (e.g., Node 18 Alpine for a smaller image)
FROM node:18-alpine

# Create a working directory inside the container
WORKDIR /

# Copy package.json and package-lock.json to install dependencies
COPY package*.json ./

# Install dependencies
RUN npm install # or yarn install

# Copy the rest of the app
COPY . .

# Expose the port Next.js runs on (default is 3000)
EXPOSE 3000

# Command to run the Next.js development server
CMD ["npm", "run", "dev"]
