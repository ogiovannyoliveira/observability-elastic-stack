import { connect } from "amqplib";

export default async function getQueueConnection() {
  const connection = await connect(process.env.AMQP_URL);

  async function createChannel() {
    return connection.createChannel();
  }

  this.connection = connection;
  this.createChannel = createChannel;

  return this;
}