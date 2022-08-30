// import 'elastic-apm-node/start.js';
import apm from 'elastic-apm-node';
// import { createServer } from 'http';

apm.start({
  // Override service name from package.json
  // Allowed characters: a-z, A-Z, 0-9, -, _, and space
  serviceName: 'nodejs-server',
  active: true,
  captureHeaders: true,
  traceContinuationStrategy: 'continue',

  transactionMaxSpans: 100,
  

  // Use if APM Server requires a token
  // secretToken: '',

  // Use if APM Server uses API keys for authentication
  // apiKey: '',

  // Set custom APM Server URL (default: http://localhost:8200)
  serverUrl: 'http://apm:8200',
})
.addErrorFilter(err => {
  // Do not send errors that have the following message
  console.log(err);
});


import express from 'express';
import { initializeCron, sendToQueue } from './cron.js';
import { config } from 'dotenv';
config();

const { APP_PORT, APM_SERVER_URL } = process.env;


// const server = createServer(async (request, response) => {
//   if (request.method === 'POST') {
//     request.on('data', async body => {
//       const data = JSON.parse(body.toString());

//       if (typeof data.msg != 'string') {
//         response.writeHead(422, { 'Content-Type': 'text/plain' });
//         response.end('Wrong type\n');
//         throw new Error('Invalid message');
//       }
        
//       const queue = 'command';
//       const msg = {
//         name: data.msg || `test-${Date.now()}`,
//         date: data.date || new Date(),
//       };
//       await sendToQueue(queue, msg);
  
//       response.writeHead(201, { 'Content-Type': 'text/plain' });
//       response.end('Hello World\n');
//     })
//     return;
//   }

//   response.writeHead(200, { 'Content-Type': 'text/plain' });
//   response.end('Hello World!\n');
// });

const app = express();

app.use(express.json());

app.get('/', (request, response) => {
  return response.send('Hello World!\n');
})

app.post('/', async (request, response) => {
  const data = request.body;

  if (typeof data.msg != 'string') {
    response.status(422).send('Wrong type');
    throw new Error('Invalid message');
  }
    
  const queue = 'command';
  const msg = {
    name: data.msg || `test-${Date.now()}`,
    date: data.date || new Date(),
  };

  await sendToQueue(queue, msg);

  return response.status(201).send('Hello World');
})

app.listen(+APP_PORT || 3333, () => {
  console.log('Server is up!\n\n');
  console.log('Initializing CRON job...');
  
  initializeCron();

  console.log('CRON job initialized.');
});