import { CronJob } from 'cron';
import getQueueConnection from './connectAmqp.js';

export function initializeCron() {
  // cron job to run every second
  new CronJob('* * * * * *', async function() {
    console.log('You will see this message every minute');
    console.log('Running at:', new Date());
    
    const channel = await (await getQueueConnection()).createChannel();
    const queue = 'command';
    const msg = `Hello World! ${Date.now()}`;
    await channel.assertQueue(queue);
    channel.sendToQueue(queue, Buffer.from(msg));
    
    console.log('Message', msg, 'sent to', queue);
    console.log('\n');
  }).start();
}