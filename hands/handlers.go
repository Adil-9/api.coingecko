package hands

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	"net/http"
	"os"
	"time"

	"github.com/Adil-9/api.coingecko/structures"
	"github.com/joho/godotenv"
)

func HandleCoins(w http.ResponseWriter, r *http.Request) {
	log.SetFlags(log.Lshortfile)

	if r.Method != http.MethodGet {
		w.Write([]byte("Method not allowed"))
		return
	}

	loadEnv()
	coinName := r.URL.Query().Get("coin")
	var coin []byte
	api_link := os.Getenv("API_KEY")
	if coinName == "" {
		coins, err := requestAllCoins(api_link)
		if err != nil {
			w.Write([]byte("internal server error"))
			return
		} else {
			w.Write(coins)
			return
		}
	} else {
		coins, err := requestAllCoins(api_link)
		if err != nil {
			w.Write([]byte("internal server error"))
			return
		} else {
			var Coins []structures.Coin
			err = json.Unmarshal(coins, &Coins)
			if err != nil {
				w.Write([]byte("internal server error"))
				return
			}
			for _, v := range Coins {
				if v.Name == coinName {
					coin, err = json.MarshalIndent(v, "", "  ")
					if err != nil {
						w.Write([]byte("internal server error"))
						return
					}
					break
				}
			}
			if len(coin) == 0 {
				w.Write([]byte(http.StatusText(http.StatusBadRequest)))
				return
			}
			w.Write(coin)
			return
		}
	}
}

func requestAllCoins(link string) ([]byte, error) {
	Coins := make([]structures.Coin, 0, 255)

	path := "coins.json"

	err := checkFile(path)
	if err != nil {
		log.Println("Error:", err)
		return nil, err
	} else {
		Coins, err = readFile(Coins, path)
		if err != nil {
			log.Println("can not read file")
			loadEnv()
			api_link := os.Getenv("API_KEY")

			Coins, err = getCoins(api_link)
			if err != nil {
				log.Println("Error:", err)
				return nil, err
			}

			marshaled, err := json.MarshalIndent(Coins, "", "  ")
			if err != nil {
				return nil, err
			}
			return marshaled, nil
		} else {
			marshal, err := json.MarshalIndent(Coins, "", "  ")
			if err != nil {
				// log.Println("Error marshalling")
			} else {
				return marshal, nil
			}
		}
	}

	loadEnv()
	api_link := os.Getenv("API_KEY")
	Coins, err = getCoins(api_link)
	if err != nil {
		return nil, err
	}

	marshaled, err := json.MarshalIndent(Coins, "", "  ")
	if err != nil {
		return nil, err
	}
	return marshaled, nil
}

func readFile(coins []structures.Coin, path string) ([]structures.Coin, error) {
	var CoinsFile structures.Coins

	data, err := os.ReadFile(path)
	if err != nil {
		log.Println("Error reading file:", err)
		return coins, err
	}

	if len(data) != 0 {
		err = json.Unmarshal(data, &CoinsFile)
		if err != nil {
			log.Println("Error unmarshaling person:", err)
			return coins, err
		}
	}

	if time.Since(CoinsFile.LastUpdate) > time.Duration(time.Minute*10) || len(data) == 0 {
		log.Println("Updating coins data")
		loadEnv()
		api_link := os.Getenv("API_KEY")

		coins, err = getCoins(api_link)
		if err != nil {
			return coins, err
		}

		CoinsFile.LastUpdate = time.Now()
		CoinsFile.AllCoins = coins

		marshaled, err := json.MarshalIndent(CoinsFile, "", "  ")
		if err != nil {
			return coins, err
		} else {
			err = os.WriteFile(path, marshaled, 0644)
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return coins, err
			}
			return coins, nil
		}
	}

	coins = CoinsFile.AllCoins

	return coins, nil
}

func loadEnv() { // load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		// log.Fatal("Error loading .env file")
	}
}

func checkFile(path string) error {

	// check if file exists
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		// if file doesn't exist, create it
		file, err := os.Create(path)
		if err != nil {
			log.Println("Error creating file:", err)
			return err
		}
		defer file.Close()
		log.Println("File created:", path)
		return nil
	} else if err != nil {
		log.Println("Error checking file:", err)
		return err
	} else {
		log.Println("File exists:", path)
		return nil
	}
}

func getCoins(link string) ([]structures.Coin, error) {
	var Coins []structures.Coin

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*7)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)
	if err != nil {
		log.Println("Encountered error sending request:", req, "\nError:", err)
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Encountered error getting response:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Encountered error reading response body:", err)
		return nil, err
	}
	err = json.Unmarshal(body, &Coins)
	if err != nil {
		log.Println("Encountered error unmarshalling data:", err)
		return nil, err
	}

	return Coins, nil
}
