import { CronJob } from 'cron';

export function initializeCron() {
  // cron job to run every minute
  new CronJob('* * * * *', function() {
    console.log('You will see this message every minute');
    console.log('Running at:', new Date());
    console.log('\n');
  }).start();
}