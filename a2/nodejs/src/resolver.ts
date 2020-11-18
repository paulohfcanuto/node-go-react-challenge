import { IResolvers } from "graphql-tools";

import viaCepRepository from "../src/integrations/viacep/repository";

const resolvers: IResolvers = {
  Query: {
    viacepAPI: (parent, args) => {
      const { cep } = args;
      return viaCepRepository.findByCep(cep);
    },
  },
};
export default resolvers;
