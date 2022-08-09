import { createServer } from 'http';
import { config } from 'dotenv';
import { initializeCron } from './cron.js';

config();

const server = createServer((request, response) => {
  response.writeHead(200, { 'Content-Type': 'text/plain' });
  response.end('Hello World!\n');
});

const { APP_PORT } = process.env;

server.listen(+APP_PORT || 3333, () => {
  console.log('Server is up!\n\n');
  console.log('Initializing CRON job...');
  
  initializeCron();

  console.log('CRON job initialized.');
});