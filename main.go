package main

import (
	"math/rand"
	//import a packet to use function from it
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// use pseudonym for *tgbotapi.BotAPI
var bot *tgbotapi.BotAPI
var chatId int64

// connectWithTelegram create a connection with telegram
func connectWithTelegram() {
	var err error
	if bot, err = tgbotapi.NewBotAPI(TOKEN); err != nil {
		panic("Cannot connect telegram")

	}
}

// sendMessage send message to chat
func sendMessage(msg string) {
	msConfig := tgbotapi.NewMessage(chatId, msg)
	bot.Send(msConfig)
}

// getAnswer generates answer randomly from a slice
func getAnswer() string {
	index := rand.Intn(len(answers))
	return answers[index]
}

// sendAnswer send answer generated with getAnswer() to chat with to concrete message
func sendAnswer(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(chatId, getAnswer())
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)

}

var answers = []string{
	"Да",
	"Нет",
	"Прославь своё прошлое, живи настоящим, строй будущее.",
	"Загляни внутрь себя, прежде чем принимать важные решения.",
	"Самый трудный путь часто оказывается самым кратким.",
	"Каждая неудача - шанс сделать следующую попытку ещё лучше.",
	"Слушай советы, но принимай решения самостоятельно.",
	"Внимательно взвесь плюсы и минусы, прежде чем действовать.",
	"Не бойся ошибок, они учат больше, чем успех.",
	"Доверься своей интуиции, она часто не подводит.",
	"Лучше сделать и пожалеть, чем не сделать и жалеть.",
	"Повторяй свои успехи, изучай свои ошибки.",
	"Проявляй терпение в деле, но не в принятии решений.",
	"Не останавливайся на одном пути, ищи новые возможности.",
	"Важно уметь распознавать шансы и использовать их.",
	"Сложное - это просто, если разбить на маленькие шаги.",
	"Доверься себе, ты сильнее, чем думаешь.",
	"Избегай поспешных решений, но не откладывай их слишком долго.",
	"Проявляй гибкость в планировании, но строгость в выполнении.",
	"Не зацикливайся на неудачах, посмотри вперёд с оптимизмом.",
	"Честность перед собой - основа успешных решений.",
	"Учись на ошибках других, это сократит твой путь к успеху.",
	"Запомни: сомнения - это твоё подсознание, готовящее тебя к решению.",
	"Лучшие решения часто приходят после тщательного обдумывания.",
	"Легче принимать решения, когда ты понимаешь свои ценности.",
	"Вдохни глубже, расслабься и послушай своё сердце перед решением.",
	"Не сравнивай свои решения с другими, ты - уникален.",
	"Избегай решений во время ярости или горя, подожди спокойствия.",
	"Страх не должен быть мотивом для принятия решений.",
	"Проще открывать двери, чем их закрывать, оставляй возможности открытыми.",
	"Будь отважным, даже если не уверен, важно двигаться вперёд.",
	"Жизнь состоит из решений, не бойся их, будь капитаном своего судна.",
}

func main() {
	connectWithTelegram()
	//for getting all updates without limit
	updateconfig := tgbotapi.NewUpdate(0)
	// create and assign a new channel to update
	for update := range bot.GetUpdatesChan(updateconfig) {
		if update.Message != nil && update.Message.Text == "/start" {
			chatId = update.Message.Chat.ID
			sendMessage("Здравствуйте, меня зовут Атлант. Вы можете задать закрытый вопрос, т.е. вопрос, на который можно ответить ДА или НЕТ. Что вы хотите спросить?")
		}

		if update.Message != nil && update.Message.Text != "" {
			if update.Message.Text != "/start" {
				sendAnswer(&update)
			}
		}

	}
}
