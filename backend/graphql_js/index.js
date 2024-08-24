import { ApolloServer } from "@apollo/server";
import { startStandaloneServer } from "@apollo/server/standalone";
import { typeDefs } from "./schema";
// Server Setup

const server = new ApolloServer({
    // TypeDef
    typeDefs,
    //Resolvers
});

const { url } = await startStandaloneServer(server, {
    listen: { port: 8080 },
});

console.log("Server is starting at Port", 8080);
