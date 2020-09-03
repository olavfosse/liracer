const express = require('express')

const port = 3101
const buildPath = `${__dirname}/front/build`
const app = express()

console.log(`Running http server on ${port}`)
app.use(express.static(buildPath))
app.get('/*', (_request, response) => {
  response.sendFile(`${buildPath}/index.html`)
})

app.listen(port)