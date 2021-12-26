<p align="center">
  <a href="" rel="noopener">
 <img width=200px height=200px src="https://www.seekpng.com/png/full/971-9710801_axie-infinity-is-a-digital-pet-community-founded.png" alt="Bot logo"></a>
</p>

<h3 align="center">AXIE FINDER with LINE Messaging API SDK</h3>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![Platform](https://img.shields.io/badge/platform-LINE-brightgreen)](https://line.me/en)
[![Document](https://img.shields.io/badge/line--bot--sdk-documentation-blue)](https://developers.line.biz/en/docs/messaging-api/line-bot-sdk)
[![Programming](https://img.shields.io/github/go-mod/go-version/kenzch1r0/axie-finder)](https://go.dev/)
</div>

---

<p align="center"> 🤖 Bot for search axie in marketplace and filtering by LINE command
    <br> 
</p>

## 📝 Table of Contents

- [About](#about)
- [Demo / Working](#demo)
- [Usage](#usage)
- [Getting Started](#getting_started)
- [Deploying your own bot](#deployment)
- [Built Using](#built_using)
- [Authors](#authors)

## 🧐 About <a name = "about"></a>

Bot use for search axie and subcribe axie by LINE command then it's will fetch data every 30 minute.

## 🎥 Demo / Working <a name = "demo"></a>

<img width=240px height=480px src="https://github.com/kenzch1r0/axie-finder/blob/main/png/demo.jpg" alt="Bot logo"></a>

## 💭 How it works <a name = "working"></a>

The bot first extracts the command and filter from the message and then push filter into parameter and combie with graphql format to get axie with AXIE open API.

The response will be recive in json format, so this service will transform data into flex message to send back to LINE platform and display to users.

The entire bot is written in Go 1.16

## 🎈 Usage <a name = "usage"></a>

To use the bot, template:
```
#find class;part,part,...;limit
```

### Example:


> #find aquatic;mouth-risky-fish,horn-shoal-star,tail-gravel-ant,back-perch;10


The first command, i.e. "find" **is** case sensitive.

---

## 🏁 Getting Started <a name = "getting_started"></a>

### Prerequisites

You need to have line account create linebot service.

follwing this :

- Getting started with the Messaging API :
https://developers.line.biz/en/docs/messaging-api/getting-started/

- Building a bot :
https://developers.line.biz/en/docs/messaging-api/building-bot/

### Installing

A step by step series of examples that tell you how to get a development env running.

```
git clone https://github.com/kenzch1r0/axie-finder.git

go mod tidy

go run main.go

```


## 🚀 Deploying your own bot <a name = "deployment"></a>

To see an example project on how to deploy your bot, please see my own configuration:

- **Heroku**: https://devcenter.heroku.com/categories/deployment

## ⛏️ Built Using <a name = "built_using"></a>

- [LINE-BOT-SDK](https://developers.line.biz/en/docs/messaging-api/line-bot-sdk/) - Go Messaging API SDKs
- [Heroku](https://www.heroku.com/) - SaaS hosting platform

## ✍️ Authors <a name = "authors"></a>

- [@kenzch1r0](https://github.com/kenzch1r0) - Idea & Initial work

