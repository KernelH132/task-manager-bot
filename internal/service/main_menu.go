package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/KernelH132/pingme/internal/repository"
)

func HandleMainMenu(ctx context.Context, chatID int64, input string) {
	switch {
	case input == "/start":
		SendChatAction(chatID, "typing...")
		welcomeMessage := `
🌸 ━━━━━━━━━━━━━━ 🌸
             ᴡᴇʟᴄᴏᴍᴇ ᴛᴏ 𝚁𝚢𝚞𝚔 𝙱𝚘𝚝 ⚚
 
  ╰┈➤   𝙵𝚞𝚗 𝚌𝚘𝚖𝚖𝚊𝚗𝚍𝚜 ‎ꫂ᭪݁
  ╰┈➤   𝙲𝚘𝚘𝚕 𝚏𝚎𝚊𝚝𝚞𝚛𝚎𝚜 .☘︎ ݁˖
  ╰┈➤   𝙴𝚗𝚍𝚕𝚎𝚜𝚜 𝚟𝚒𝚋𝚎𝚜 𖦹

𝚁𝚢𝚞𝚔 𝙱𝚘𝚝 𝚒𝚜 𝚑𝚎𝚛𝚎 𝚝𝚘 𝚔𝚎𝚎𝚙 𝚝𝚑𝚒𝚗𝚐𝚜 𝚒𝚗𝚝𝚎𝚛𝚎𝚜𝚝𝚒𝚗𝚐.

🌸 ━━━━━━━━━━━━━━ 🌸

      /help — ᴠɪᴇᴡ ᴀᴠᴀɪʟᴀʙʟᴇ ᴄᴏᴍᴍᴀɴᴅs

  🦋    ᴍᴀɪɴ ɢʀᴏᴜᴘ →  t.me/ryuk_bott  🦋

🌸 ━━━━━━━━━━━━━━ 🌸
                ʙᴏᴛ sᴛᴀᴛᴜs: ᴏɴʟɪɴᴇ ᯓ★
`
		err := SendPhotoWithCaption(chatID, "https://pin.it/6dAfyFQff", welcomeMessage)
		if err != nil {
			fmt.Println("Error sending welcome message:", err)
			SendMessage(chatID, "Welcome to Ryuk Bot! Use /help to see available commands.")
		}

	case strings.ToLower(input) == "/register":
		err := SetUserState(ctx, repository.DB, chatID, "awaiting_username")
		if err != nil {
			SendMessage(chatID, "Connection error. Please try /register again.")
			return
		}

		SendMessage(chatID, "Please enter a username 😊.")

	case strings.ToLower(input) == "/help":
		SendMessage(chatID, `
            𝚁𝚢𝚞𝚔  𝚌𝚘𝚖𝚖𝚊𝚗𝚍𝚜 & 𝙵𝚎𝚊𝚝𝚞𝚛𝚎𝚜 ⚚

                    ────୨ৎ────

   ╰┈➤ /register   →   create/edit profile
   ╰┈➤ /help         →    show menu
   ╰┈➤ /profile     →    view your profile
   ╰┈➤ /ping         →    check if bot is up  

                      ────୨ৎ────


𝚁𝚢𝚞𝚔 𝚒𝚜 𝚊 𝚠𝚘𝚛𝚔 𝚒𝚗 𝚙𝚛𝚘𝚐𝚛𝚎𝚜𝚜, 𝚠𝚒𝚝𝚑 𝚗𝚎𝚠 𝚏𝚎𝚊𝚝𝚞𝚛𝚎𝚜 𝚋𝚎𝚒𝚗𝚐 𝚊𝚍𝚍𝚎𝚍 𝚛𝚎𝚐𝚞𝚕𝚊𝚛𝚕𝚢.

                                       ୨ৎ

ᴄʜᴇᴄᴋ ᴏᴜᴛ ᴛʜᴇ ᴍᴀɪɴ ɢʀᴏᴜᴘ ᴛᴏ ᴄᴏɴɴᴇᴄᴛ ᴡɪᴛʜ ᴏᴛʜᴇʀs → https://t.me/ryuk_bott
`)
	case input == "/ping":
		SendMessage(chatID, "pong")

	case input == "/group":
		SendMessage(chatID, "Join our main group to connect with other users, get updates and support:\n\nhttps://t.me/ryuk_bott")

	case input == "/profile":
		username, err := GetProfile(ctx, chatID)
		if err != nil {
			SendMessage(chatID, "Error fetching profile. Please register first using /register.")
			return
		}
		SendMessage(chatID, fmt.Sprintf("Your username is: %s", username))
	}

}

func HandleUsernameCreation(ctx context.Context, chatID int64, username string) {

	if len(username) >= 32 {
		SendMessage(chatID, "That username is too long! Keep it under 32 characters.")
		return

	}
	err := SaveUsernameToDB(ctx, repository.DB, chatID, username)
	if err != nil {
		SendMessage(chatID, "System error. Try again later.")
		return
	}

	// Step 2: RESET to idle so they can use other commands again
	err = SetUserState(ctx, repository.DB, chatID, "idle")
	if err != nil {
		fmt.Printf("Warning: failed to reset user state for %d: %v\n", chatID, err)
		return
	}

	createdMessage := fmt.Sprintf("Hi %s, your username has been created!🚀", username)

	// Step 3: Success message
	SendMessage(chatID, createdMessage)

}
