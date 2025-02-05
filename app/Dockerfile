# Build stage
FROM node:20-alpine AS builder

WORKDIR /app

# Install pnpm globally
# RUN corepack enable && corepack prepare pnpm@latest --activate
RUN npm install -g pnpm@9.15.3

# Copy package files
COPY pnpm-lock.yaml ./
COPY package.json ./

# Install dependencies
RUN pnpm install

# Copy source code
COPY . .

# Build the application
RUN pnpm build

# Final stage
FROM node:20-alpine

WORKDIR /app

# Install pnpm globally
# RUN corepack enable && corepack prepare pnpm@latest --activate
RUN npm install -g pnpm@9.15.3

# Copy built application
COPY --from=builder /app/.output ./

# Expose port 3000
EXPOSE 3000

# Run the application
CMD ["node", "server/index.mjs"]