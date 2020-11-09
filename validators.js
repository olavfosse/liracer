validateNickname = (nickname) => {
  if(typeof nickname !== 'string') {
    return {
      valid: false,
      problem: 'nickname must be a string'
    }
  } else if(nickname.length < 1) {
    return {
      valid: false,
      problem: 'nickname must be at least one character'
    }
  } else if(nickname.length > 15) {
    return {
      valid: false,
      problem: 'nickname must be at most fifteen characters'
    }
  } else if (nickname.split('').some(c => !/^([a-zA-Z0-9_]{1,15})$/.test(c))) {
    return {
      valid: false,
      problem: 'nickname cannot contain special characters'
    }
  } else if (nickname === 'liracer') {
    return {
      valid: false,
      problem: 'nickname cannot be "liracer"'
    }
  } else {
    return {
      valid: true
    }
  }
}

module.exports = {
  validateNickname
}
