//   Account Generator Bot
//   Copyright (C) 2021 AnonyIndian (@xnony)

//   This program is distributed in the hope that it will be useful,
//   but WITHOUT ANY WARRANTY; without even the implied warranty of
//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//   GNU Affero General Public License for more details.


package main

import (
	"fmt"
	"math/rand"
	"strings"
	"strconv"
	"time"
	"os"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

var MSG string

func main() {
	// Put Your Bot Token via ENV Vars
	b, err := gotgbot.NewBot(os.Getenv("TOKEN"))
	if err != nil {
		panic("failed to create new bot: " + err.Error())
	}

	// Create updater and dispatcher.
	updater := ext.NewUpdater(b, nil)
	dispatcher := updater.Dispatcher

	// Handlers for runnning commands.
	dispatcher.AddHandler(handlers.NewCommand("start", start))
	dispatcher.AddHandler(handlers.NewCommand("gen", gen))

	err = updater.StartPolling(b, &ext.PollingOpts{Clean: true})
	if err != nil {
		panic("failed to start polling: " + err.Error())
	}
	fmt.Printf("%s has been started...\nMade with ❤️ by @xnony (@StarDevs).\n", b.User.Username)

	// Idle, to keep updates coming in, and avoid bot stopping.
	updater.Idle()
}


func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		// Put your accounts here
		// "email:pass",
		// Following are some demo accounts
		"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",
		"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",
		"sk_live_69GKI0saLB8uIEnxzv8VTvRX",
		"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",
		"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",
		"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",
		"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",
		"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",
	}

	return formats[rand.Intn(len(formats))]
}

func start(ctx *ext.Context) error {
	// To ensure bot does not reply outside of private chats
	if ctx.EffectiveChat.Type != "private" {
		return nil
	}
	// Following string is replied to cmd user on /start 
	MSG = "*Hi %v,\n" +
		"I am an Account Generator Bot\n" +
		"-------------------------------------------------\n" +
		"I can provide premium accounts of different services\n" +
		"--------------------------------------------------\n" +
		"Do /gen to generate an account\n" +
		"--------------------------------------------------\n" +
		"❤️Brought to You By @Stardevs❤️\n*"

	user := ctx.EffectiveUser
	channel_id, cerror := strconv.Atoi(os.Getenv("CHANNEL_ID"))
	if cerror != nil {fmt.Println("Please Provide me a valid Channel ID")}
	member, eror := ctx.Bot.GetChatMember(int64(channel_id), user.Id)
	if eror != nil {
		ctx.Bot.SendMessage(ctx.EffectiveChat.Id, "*Bot not admin in JoinCheck Channel.*", nil)
		return nil
	}
	// For Checking either user joined channel or not
	if member.Status == "member" || member.Status == "administrator" || member.Status == "creator" {
		_, err := ctx.EffectiveMessage.Reply(ctx.Bot, fmt.Sprintf(MSG, user.FirstName), &gotgbot.SendMessageOpts{
			ParseMode: "Markdown",
		})
		if err != nil {
			fmt.Println("failed to send: " + err.Error())
		}
	} else {
		// An Error message replied to command user if he's not member of the JoinCheck Channel
		ctx.EffectiveMessage.Reply(ctx.Bot, fmt.Sprintf("*You must join %v To use me.*", os.Getenv("CHANNEL_USERNAME")), &gotgbot.SendMessageOpts{ParseMode: "Markdown"})
	}
	if strings.ToLower(os.Getenv("LOGS")) != "false"{
		logs, log_err := strconv.Atoi(os.Getenv("LOGS"))
		if log_err != nil{fmt.Println(log_err.Error())}
		// Following message is sent in Logs Group (if set)
		ctx.Bot.SendMessage(int64(logs), fmt.Sprintf("#Start\n\nBot Started By %v(%v)", user.FirstName, user.Id), nil)
	}
	return nil
}

func gen(ctx *ext.Context) error {
	// To ensure bot does not reply outside of private chats
	if ctx.EffectiveChat.Type != "private" {
		return nil
	}
	Combo := strings.Split(randomFormat(), ":")
	// Following string is replied to cmd user on /gen 
	MSG = "𝙃𝙚𝙧𝙚 𝙄𝙨 𝙔𝙤𝙪𝙧 %v 𝘼𝙘𝙘𝙤𝙪𝙣𝙩" +
		"\n\nSK: `%v`" +
		"\n𝙂𝙚𝙣𝙚𝙧𝙖𝙩𝙚𝙙 𝘽𝙮: *%v*" +
		"\n\n𝙏𝙝𝙖𝙣𝙠 𝙮𝙤𝙪 𝙛𝙤𝙧 𝙪𝙨𝙞𝙣𝙜 𝙢𝙚!\n❤️𝙎𝙝𝙖𝙧𝙚 & 𝙎𝙪𝙥𝙥𝙤𝙧𝙩 *%v*❤️"

	user := ctx.EffectiveUser
	channel_id, cerror := strconv.Atoi(os.Getenv("CHANNEL_ID"))
	if cerror != nil {fmt.Println("Please Provide me a valid Channel ID")}
	member, eror := ctx.Bot.GetChatMember(int64(channel_id), user.Id)
	if eror != nil {
		ctx.Bot.SendMessage(ctx.EffectiveChat.Id, "Bot not admin in JoinCheck Channel", nil)
		return nil
	}
	// For Checking either user joined channel or not
	if member.Status == "member" || member.Status == "administrator" || member.Status == "creator" {
		_, err := ctx.EffectiveMessage.Reply(ctx.Bot, fmt.Sprintf(MSG, os.Getenv("ACC_NAME"), Combo[0], Combo[1], user.FirstName, os.Getenv("CHANNEL_USERNAME")), &gotgbot.SendMessageOpts{
			ParseMode: "Markdown",
		})
		if err != nil {
			fmt.Println("failed to send: " + err.Error())
		}
	} else {
		// An Error message replied to command user if he's not member of the JoinCheck Channel
		ctx.EffectiveMessage.Reply(ctx.Bot, fmt.Sprintf("*You must join %v to use me.*", os.Getenv("CHANNEL_USERNAME")), &gotgbot.SendMessageOpts{ParseMode: "Markdown"})
	}
	if strings.ToLower(os.Getenv("LOGS")) != "false"{
		logs, log_err := strconv.Atoi(os.Getenv("LOGS"))
		if log_err != nil{fmt.Println(log_err.Error())}
		// Following message is sent in Logs Group (if set)
		ctx.Bot.SendMessage(int64(logs), fmt.Sprintf("#AccClaimed\n\n%v generated a new account.", user.FirstName), nil)
	}
	return nil
}
