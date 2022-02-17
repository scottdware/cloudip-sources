# cloudip-sources
Repo containing easy access to cloud vendors public IP ranges

This repo is a web service which allows access to the three major vendors' (AWS, Google, Azure) public IP ranges. There is the ability to pull just the IPv4 or IPv4 addresses, or all of them.

## Usage

You can replace the `<vendor>` param with any of the following:
* aws
* google
* azure

**IPv4 Ranges**

`https://cloudip-sources.sdubs.org/<vendor>/ranges/4`

**IPv6 Ranges**

`https://cloudip-sources.sdubs.org/<vendor>/ranges/6`

**IPv4 and IPv6 Ranges**

`https://cloudip-sources.sdubs.org/<vendor>/ranges/all`

## CLI Version

There is also a CLI tool which retrieves the IP ranges, and can also export them to CSV if you choose. That can be downloaded here:

[cloudip](https://github.com/scottdware/cloudip/releases/latest)
