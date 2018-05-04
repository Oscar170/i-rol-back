export default {
  Mutation: {
    signIn: (parent, { username, password }) => {
      return fetch('http://user-api:8080/sign-in', {
        method: 'post',
        headers: { 'content-type': 'application/json' },
        body: JSON.stringify({
          username,
          password,
        })
      })
      .then(r => r.json())
    },
    signUp: (parent, { username, email, password }) => {
      return fetch('http://user-api:8080/sign-up', {
        method: 'post',
        headers: { 'content-type': 'application/json' },
        body: JSON.stringify({
          username,
          email,
          password,
        })
      })
      .then(r => r.json())
    },
  }
}
