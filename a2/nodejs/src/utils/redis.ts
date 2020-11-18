import redis, { RedisClient } from "redis";

const client = redis.createClient(process.env.HOST_REDIS!);

export function getRedisClient(): RedisClient {
  return client;
}
