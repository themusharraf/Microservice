const express = require('express');

const Model = require('../models/model');

const router = express.Router()

module.exports = router;


// router.post('/post', (req, res) => {
//     res.send('Post API');
// });

router.get('/create/:name', (req, res) => {
    const data = new Model({
        name: req.params.name
    })

    try {
        const dataToSave = data.save();
        res.status(200).json(req.params.name)
    }
    catch (error) {
        res.status(400).json({message: error.message})
    }
});


router.get('/all', async (req, res) => {
    try{
        const data = await Model.find();
        res.json(data)
    }
    catch(error){
        res.status(500).json({message: error.message})
    }
})
