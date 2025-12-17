# pwman

**pwman** is a simple, secure, command-line password manager and generator for Linux users, written in Go. It allows you to generate strong passwords, store them encrypted locally, and retrieve them easily.

---

## Features

- Generate strong passwords of customizable length.
- Add passwords for websites with username and encrypted storage.
- Retrieve stored passwords securely using your master password.
- Local storage in `$HOME/.pwman/store.json`.
- Works entirely offline. No cloud storage required.

---

## Installation

You can install the latest version directly with Go:

```bash
go install github.com/alilo113/pwman@v0.1.2
```
Make sure ```$GOPATH/bin``` or ```$HOME/go/bin``` is in your ```PATH``` to run ```pwman``` from anywhere.


---

## Usages

### Generate a password
```
pwman -g
````
### Generate a password with a specified length
```
pwman -l 20
```
### Add a password to local storage
```
pwman -a
```
### Retrieve a password
```
pwman -r
```

---

## Local Storage
All passwords are encrypted with your master password and stored in:
```
$HOME/.pwman/store.json
```

---

## Donations
If you find pwman useful and want to support development: send donations to my PayPal at aliopdandan@gmail.com, donations are welcome. Every contribution helps keep the project alive and free for everyone.

## Contributing
Contributions are welcome! Feel free to open issues or submit pull requests on GitHub
