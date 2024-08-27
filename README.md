# SimpleAuthenticator
simple authenticator like google authenticator 

### usage:
`password` `code_length`
```bash
go run main.go mypassword 5
Your code: 77857
```
The output changes every 25 seconds and you can get the old output using `calculateTime(25)`.
