package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Chat struct {
	ID int64 `json:"id"`
}

type Message struct {
	MessageID int64  `json:"message_id"`
	Chat      Chat   `json:"chat"`
	Text      string `json:"text"`
}

type Update struct {
	UpdateID int64    `json:"update_id"`
	Message  *Message `json:"message"`
}

type UpdatesResponse struct {
	OK     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type SendMessageRequest struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

type TelegramAPIError struct {
	StatusCode int
}

func (e *TelegramAPIError) Error() string {
	return fmt.Sprintf(
		"telegram API returned HTTP status: %d",
		e.StatusCode,
	)
}

func main() {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN is not set or empty")
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	client := &http.Client{
		Timeout: 40 * time.Second,
	}

	var offset int64

	log.Println("Bot started")

	for {
		if ctx.Err() != nil {
			log.Println("Bot stopped")
			return
		}
		updates, err := getUpdates(ctx, client, token, offset)
		if err != nil {
			if ctx.Err() != nil {
				log.Println("Bot stopped")
				return
			}

			log.Println(err)

			select {
			case <-ctx.Done():
				log.Println("Bot stopped")
				return
			case <-time.After(3 * time.Second):
				continue
			}
		}

		for _, update := range updates {
			if ctx.Err() != nil {
				log.Println("Bot stopped")
				return
			}

			log.Println("Update ID:", update.UpdateID)

			if update.Message != nil && update.Message.Text == "/start" {
				err := sendMessage(
					ctx,
					client,
					token,
					update.Message.Chat.ID,
					"Добро пожаловать в TG Marketplace!",
				)

				if err != nil {
					if ctx.Err() != nil {
						log.Println("Bot stopped")
						return
					}

					var apiErr *TelegramAPIError

					if errors.As(err, &apiErr) &&
						apiErr.StatusCode == http.StatusForbidden {

						log.Println("Message forbidden, skipping update:", update.UpdateID)
						offset = update.UpdateID + 1
						continue
					}

					log.Println(err)

					select {
					case <-ctx.Done():
						log.Println("Bot stopped")
						return
					case <-time.After(3 * time.Second):
					}

					break
				}
			}

			offset = update.UpdateID + 1
		}
	}
}

func getUpdates(ctx context.Context, client *http.Client, token string, offset int64) ([]Update, error) {
	url := fmt.Sprintf(
		"https://api.telegram.org/bot%s/getUpdates?offset=%d&timeout=30",
		token,
		offset,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request")
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Telegram API")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected HTTP status: %d", resp.StatusCode)
	}

	var updates UpdatesResponse

	err = json.NewDecoder(resp.Body).Decode(&updates)
	if err != nil {
		return nil, fmt.Errorf("failed to decode Telegram response")
	}

	if !updates.OK {
		return nil, fmt.Errorf("Telegram API returned an error")
	}

	return updates.Result, nil
}

func sendMessage(ctx context.Context, client *http.Client, token string, chatID int64, text string) error {
	url := fmt.Sprintf(
		"https://api.telegram.org/bot%s/sendMessage",
		token,
	)

	message := SendMessageRequest{
		ChatID: chatID,
		Text:   text,
	}

	body, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to encode message")
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		url,
		bytes.NewReader(body),
	)
	if err != nil {
		return fmt.Errorf("failed to create request")
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send message")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return &TelegramAPIError{
			StatusCode: resp.StatusCode,
		}
	}
	var result struct {
		OK bool `json:"ok"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("failed to decode response")
	}

	if !result.OK {
		return fmt.Errorf("Telegram API returned an error")
	}

	return nil
}
