const fs = require('fs')
const path = require('path')

const mapCarriageReturnsToNewlines = (s) => {
  return s.split('').map(c => c === '\r' ? '\n' : c).join('')
}

const getSnippets = () => {
  const snippets = []
  const rootPath = path.join(__dirname, 'snippets')

  const languages = fs.readdirSync(rootPath)
  languages.forEach(language => {
    const programsPath = path.join(rootPath, language)
    const programPaths = fs.readdirSync(programsPath)

    programPaths.forEach(program => {
      const programPath = path.join(programsPath, program)
      const options = { encoding: 'utf8', flag:'r' }
      let code = fs.readFileSync(programPath, options)
      // This might fix #50
      // Currently this is not tested on windows, where the problem occured because I(fossegrim) don't have a windows box at hand
      code = mapCarriageReturnsToNewlines(code)
      const snippet = {
        language,
        code,
        name: program,
      }

      snippets.push(snippet)
    })
  })

  return snippets
}

const getProgramPath = (language, program) => path.join(__dirname, 'snippets', language, program)

const enforceNoSpaceIndentation = (snippets) => {
  snippets.forEach(snippet => {
    snippet.code.split('\n').forEach((line, index) => {
      if(line.startsWith(' ')) {
        const path = getProgramPath(snippet.language, snippet.name)
        console.error(`ERROR: Found space indentation at ${path}:${index + 1}`)
        process.exit(1)
      }
    })
  })
}

const enforceMaxium79CharactersPerLine = (snippets) => {
  snippets.forEach(snippet => {
    snippet.code.split('\n').forEach((line, index) => {
      if(line.length > 79) {
        const path = getProgramPath(snippet.language, snippet.name)
        console.error(`ERROR: Found line with more than 79 characters at ${path}:${index + 1}`)
        process.exit(1)
      }
    })
  })
}

const enforceNoExcessNewline = (snippets) => {
  snippets.forEach(snippet => {
    const index = snippet.code.length - 1
    if(snippet.code[index] === '\n'){
      const path = getProgramPath(snippet.language, snippet.name)
      console.error(`ERROR: Found excess newline at ${path}:${index + 1}`)
      process.exit(1)
    }
  })
}

const snippets = getSnippets()
enforceNoSpaceIndentation(snippets)
enforceMaxium79CharactersPerLine(snippets)
enforceNoExcessNewline(snippets)

module.exports = snippets