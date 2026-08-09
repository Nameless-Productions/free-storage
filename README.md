# Free storage
A CLI tool that stores your files in Discord

### How it works?

On Discord a bot can upload files up to 10MB per message but there is no limit how many messages you send.

This splits your file into 10MB chunks and sends it in different messages then stores the message IDs in a json file

When downloading it gets the message IDs from the json file, gets the chunks and connects them

### Setup and usage

Firstly get the file for linux from the latest release

On Windows you will have to clone the repo and build it manually

Then move it into your /bin dir

```bash
sudo mv Downloads/free-storage /bin
```

After that you sucessfuly installed it

**Usage:**

To upload a file run:

```bash
free-storage fileName
```

This shows the progress for it

To download just run:

```bash
free-storage fileName
```

If the file exists then it will download it