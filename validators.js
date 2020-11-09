export const validateNickname = (nickname) => {
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
  } else if (nick.split('').some(c => !/^([a-zA-Z0-9_]{1,15})$/.test(c))) {
    return {
      valid: false,
      problem: 'nickname must not contain special characters'
    }
  } else if (nick === 'liracer') {
    return {
      valid: false,
      problem: 'nickname must not be "liracer"'
    }
  } else {
    return {
      valid: true
    }
  }
}
