# Password keeper v. 0.5.2

🌐 **Language**: [English](README.en.md ) / [Russian](README.ru.md )

## Instructions

1. When launching the program, create a new account by creating a username and password or PIN code.

    `Please note, if you forgot it, you will lose access to your saved data, write it down in your notebook`

    `Do not delete or modify the contents of the Database file.json " and " settings.json" in the "data" folder. The only thing they can do in case of data corruption is to use a backup of these files in the "data" folder to restore the state of the newly installed application with loss of access to all passwords`

    `In the 'settings.JSON' file, they can only change the 'password_generation_lenth' value, which is responsible for the length of the generated password. It is recommended to leave the value set to 16, but you can vary it from 8 to 16`

2. By creating a new account, they will be able to save passwords from Internet resources under it. For each new account, you will need to enter the master password for your current account. It is possible to create new accounts, as well as log in to other accounts if you know the password for them. Each account has access only to its own passwords, but not to anyone else's

3. When saving passwords, be sure to write them down on a physical medium, such as a notebook, because the application may contain flaws that lead to data loss (during the tests, information about the loss was not displayed, but it is strongly recommended to physically save passwords)

4. When you select the "copy password" action, after authentication, the password will be pasted to the clipboard, and you will only need to enter the password in the desired form using Ctrl + V

5. When creating a new password, it is possible to generate it randomly. After generation, it will be available for copying

6. It is used to encrypt passwords using the AES-256 algorithm with a 32-character generated encryption key. The encryption key is created based on the account's master password, a random public generated number, and the key. The master password is stored in hashed form, but it is almost impossible to find out the master password from the hash. The scrypto package is used to generate a password hash based on the SHA-256 hash function

> Confidentiality of data is ensured