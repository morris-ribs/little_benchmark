import express from 'express';
import bodyParser from 'body-parser';
import DummyInfo from './dummyinfo';

let app = express();
app.use(express.static('public'));

// create application/json parser 
app.use(bodyParser.json());
 
app.post('/bombhere', function(req, res) {
    try{
        if (!req.body) return res.sendStatus(400);

        var dmi = new DummyInfo(req.body); 
        return res.status(200).contentType('application/json').send(JSON.stringify(dmi));
    }
    catch(err) {
        console.log(err);
        return res.sendStatus(500);
    }
});

app.listen(11000, function(){
  console.log('Server running in port 11000');
  setTimeout(function(){process.exit();}, 65000);
});

