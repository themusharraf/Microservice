const express = require('express');
const mongoose = require('mongoose');
const mongoString = "mongodb://mongodb:27017/docker-d";
const routes = require('./routes/routes');


mongoose.connect(mongoString, {useNewUrlParser: true});
const database = mongoose.connection;

console.log(mongoString)

database.on('error', (error) => {
    console.log(error)
})

database.once('connected', () => {
    console.log('Database Connected');
})
const app = express();

app.use(express.json());

app.listen(8080, () => {
    console.log(`Server Started at ${8080}`)
})

app.use('/api', routes)