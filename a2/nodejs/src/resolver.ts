import { IResolvers } from "graphql-tools";
import fetch from "node-fetch";

const baseURL = `https://viacep.com.br/ws`

const resolvers: IResolvers = {
  Query: {
    viacepAPI: (parent, args) => {
      const { cep } = args
      return fetch(`${baseURL}/${cep}/json`).then(res => res.json())
    },
  },
};
export default resolvers;
