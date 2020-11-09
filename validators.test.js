const { validateNickname } = require('./validators')

test('validates legitimate nickname', () => {
  expect(validateNickname('fossegrim')).toStrictEqual({
    valid: true
  })
})

test('invalidates non-string nickname', () => {
  expect(validateNickname(undefined)).toStrictEqual({
    valid: false,
    problem: 'nickname must be a string'
  })
})

test('invalidates zero-char nickname', () => {
  expect(validateNickname('')).toStrictEqual({
    valid: false,
    problem: 'nickname must be at least one character'
  })
})

test('invalidates 16-char nickname', () => {
  expect(validateNickname('123456789ABCDEFG')).toStrictEqual({
    valid: false,
    problem: 'nickname must be at most fifteen characters'
  })
})

test('validates 15-char nickname', () => {
  expect(validateNickname('123456789ABCDEF')).toStrictEqual({
    valid: true
  })
})

test('invalidates special-char nickname', () => {
  expect(validateNickname('_3biNHax0r_')).toStrictEqual({
    valid: false,
    problem: 'nickname cannot contain special characters'
  })
})

test('invalidates "liracer" nickname', () => {
  expect(validateNickname('liracer')).toStrictEqual({
    valid: false,
    problem: 'nickname cannot be "liracer"'
  })
})
