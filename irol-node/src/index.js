import 'isomorphic-fetch'
import { GraphQLServer } from 'graphql-yoga'

import BaseSchema from './schema.graphql'

import UserApiResolver from './user-api'
import UserApiSchema from './user-api/schema.graphql'


const resolvers = {
  Mutation: {
    ...UserApiResolver.Mutation
  }
}

const server = new GraphQLServer({
  typeDefs: [ BaseSchema, UserApiSchema ],
  resolvers,
})

server.start(() => console.log(`Server is running on http://localhost:4000`))