

/auth/apple

1. validateRequestBody
2. validateIdentityToken
3. get `sub` & `email`
4. check if user exists by `sub`
5. create user
6. generate access token, generate refresh token
7. store all the information about the user in db

body: {
    identityToken: "<SUPER_LONG_IDENTITY_TOKEN>"
}


/auth/google


