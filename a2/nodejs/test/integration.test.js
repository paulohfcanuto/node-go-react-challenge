const chai = require('chai');

const expect = chai.expect;
const url = `http://localhost:8080/`;
const request = require('supertest')(url);

describe('GraphQL integration test', () => {
    it('retorna os dados do cep 88807010', (done) => {
        request.post('graphql')
            .send({ query: '{ viacepAPI(cep: "88807010") { cep logradouro complemento bairro } }'})
            .expect(200)
            .end((err,res) => {
                if (err) return done(err);
                expect(res.body.data.viacepAPI.cep).to.equal('88807-010');
                expect(res.body.data.viacepAPI.logradouro).to.equal('Rua Osório de Lima');
                expect(res.body.data.viacepAPI.complemento).to.empty;
                expect(res.body.data.viacepAPI.bairro).to.equal('São Sebastião');
                done();
            })
    });
});
