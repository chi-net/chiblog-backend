import express from 'express'

let app = express()

app.get('/app', (req, res) => {
  res.end(JSON.stringify({ status: 200, data: 'chihuo2104 needs 50$(x'}))
})

app.listen(12546, () => {
  console.log('APP listening on port 12546')
})
