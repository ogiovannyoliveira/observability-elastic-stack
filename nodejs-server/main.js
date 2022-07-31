import { createServer } from 'http';
import { initializeCron } from './cron.js';

const server = createServer((request, response) => {
  response.writeHead(200, { 'Content-Type': 'text/plain' });
  response.end('Hello World!\n');
});

server.listen(3333, () => {
  console.log('Server is up!\n\n');
  console.log('Initializing CRON job...');
  
  initializeCron();

  console.log('CRON job initialized.');
});