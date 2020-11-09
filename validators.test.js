const { validateNickname } = require('./validators')

test('invalidates special-char nickname', () => {
  expect(validateNickname('_3binhax0r_')).toStrictEqual({
    valid: false,
    problem: 'nickname cannot contain special characters'
  })
})
