module.exports = {
  apps: [
    {
      name: 'go-community',
      port: '3002',
      exec_mode: 'cluster',
      instances: 1,
      script: './.output/server/index.mjs',
    },
  ],
}
