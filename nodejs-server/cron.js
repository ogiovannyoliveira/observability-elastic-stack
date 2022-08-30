import { CronJob } from 'cron';
import getQueueConnection from './connectAmqp.js';

export function initializeCron() {
  // cron job to run every second
  new CronJob('* * * * *', async function() {
    console.log('Running at:', new Date());
    
    const queue = 'command';
    const msg = {
      name: `test-${Date.now()}`,
      date: new Date(),
    };

    await sendToQueue(queue, msg);
    
    console.log('\n');
  }).start();
}

export async function sendToQueue(queue, msg) {
  return new Promise((resolve, reject) => {
    const channel = getQueueConnection().then(conn => conn.createChannel());
    channel.then(ch => {
      ch.assertQueue(queue);
      ch.sendToQueue(queue, Buffer.from(JSON.stringify(msg)));
      console.log('Message', msg, 'sent to', queue);
      resolve();
    }).catch(err => reject(err));
  }).catch(err => console.log(err));
}