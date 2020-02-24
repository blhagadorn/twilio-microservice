const express = require('express');
const bodyParser = require('body-parser');
const app = express();
const dotenv = require('dotenv').config();
app.use(bodyParser.json());
const PORT = process.env.PORT || 8087;
app.listen(PORT, () => console.log(`Example app listening on port ${PORT}!`))

if (!process.env.ACCOUNT_SID || !process.env.AUTH_TOKEN) {
  console.log("Couldn't find environmental variables account_sid or auth_token")
  process.exit(1)
}
console.log("Twilio microservice listening and initialized");
const accountSid = process.env.ACCOUNT_SID;
const authToken = process.env.AUTH_TOKEN;
const client = require('twilio')(accountSid, authToken);

app.post('/text', function (req, res) {
  sendMessage(req.body.from, req.body.to, req.body.body)
  console.log(req.body)
  res.send('Message sent')
})

// Health check for k8s liveness
app.get('/healthz', function (req, res) {
  res.send('Healthy')
})

function sendMessage(from, to, body) {
  client.messages
  .create({
     body: body,
     from: from,
     to: to
   })
  .then(message => console.log(message.sid))
  .catch(error => console.log(error));
}


