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
        game(_, args) {
            return db.games.find((game) => game.id == args.id);
        },
        author(_, args) {
            return db.authors.find((author) => author.id === args.id);
        },
    },
    Game: {
        reviews(parent) {
            return db.reviews.filter((r) => r.game_id === parent.id);
        },
    },
    Author: {
        reviews(parent) {
            return db.reviews.filter((r) => r.author_id === parent.id);
        },
    },
    Review: {
        author(parent) {
            return db.authors.find((author) => author.id == parent.author_id);
        },
        game(parent) {
            return db.games.find((game) => game.id === parent.game_id);
        },
    },
    Mutation: {
        deleteGame(_, args) {
            db.games = db.games.filter((g) => g.id != args.id);
            return db.games;
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
