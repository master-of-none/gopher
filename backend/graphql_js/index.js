import { ApolloServer } from "@apollo/server";
import { startStandaloneServer } from "@apollo/server/standalone";
import { typeDefs } from "./schema.js";
import db from "./_db.js";
// Server Setup

// Resolver Constant
const resolvers = {
    Query: {
        games() {
            return db.games;
        },
        reviews() {
            return db.reviews;
        },
        authors() {
            return db.authors;
        },
        review(_, args) {
            return db.reviews.find((review) => review.id === args.id);
        },
    },
};
const server = new ApolloServer({
    // TypeDef
    typeDefs,
    //Resolvers
    resolvers,
});

const { url } = await startStandaloneServer(server, {
    listen: { port: 8080 },
});

console.log("Server is starting at Port", 8080);
