# Hidden Present

- Category: `misc`
- Challenge author: **Björn Österman**

## Description

Where is it?

### Connection info

```sh
nc 2021.santahack.xyz 42022
```

## Writeup

This challenge is a dash-prompt where you can find a hidden file name '.flag'

```sh
nc 2021.santahack.xyz 42022
ls -la
cat .flag
```
