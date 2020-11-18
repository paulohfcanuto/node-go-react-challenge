import fetch from "node-fetch";

import { promisify } from "util";

import { getRedisClient } from "../../utils/redis";
import { Cep } from "./schema";

const baseURL = process.env.HOST_VIA_CEP!;

const redisClient = getRedisClient();

const getCache = promisify(redisClient.get).bind(redisClient);

function putCache(cep: string, content: Cep): Cep {
  //não espera o async concluir pois se der problema no cache a aplicação pode continuar sem problemas
  redisClient.set(cep, JSON.stringify(content));
  return content;
}

class Repository {
  async findByCep(cep: string): Promise<Cep> {
    const cache = await getCache(cep);

    if (cache) {
      return JSON.parse(cache);
    }

    return fetch(`${baseURL}/${cep}/json`)
      .then((res) => res.json())
      .then((json) => putCache(cep, json));
  }
}

export default new Repository();
