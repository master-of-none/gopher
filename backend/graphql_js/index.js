import { ApolloServer } from "@apollo/server";
import { startStandaloneServer } from "@apollo/server/standalone";

// Server Setup

const server = new ApolloServer({
    // TypeDef
    //Resolvers
});

const { url } = await startStandaloneServer(server, {
    listen: { port: 8080 },
});

console.log("Server is starting at Port", 8080);
