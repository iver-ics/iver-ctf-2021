# Magic Password

- Category: `web`
- Challenge author: **Pontus Norrström**

## Description

Santa locked himself out of his super-duper-website. Can you help find some magic password?

### Connection info

<http://2021.santahack.xyz:42203>

## Writeup

Knowing about "magic passwords" in PHP is considered one of the basic PHP
security holes to know about.

A big hint from the source code (<http://2021.santahack.xyz:42203/?source>) is
the mix of `==` and `===`.

Googling `magic password list PHP` lends you results such as this repository:
<https://github.com/spaze/hashes>.

Using any of the passwords from the `md5.md` file (<https://github.com/spaze/hashes/blob/master/md5.md>),
such as the first one `240610708`, will get you through the challenge.
