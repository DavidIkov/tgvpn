
from telebot import TeleBot
from dotenv import load_dotenv
import os
import threading

load_dotenv()

admin_ids = list(
    map(int, os.getenv("TG_ADMINS").split(",")))
token = os.getenv("TG_TOKEN")

bot = TeleBot(token)


@bot.message_handler(commands=['start'])
def start(message):
    bot.send_message(message.chat.id, "hello world")


bot_thread = threading.Thread(
    target=bot.infinity_polling, daemon=True)

bot_thread.start()
bot_thread.join()
