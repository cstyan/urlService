package dataStore

import (
	"github.com/go-redis/redis"
	"fmt"
	"log"
	"strings"
	"strconv"
	"os"
)

type RedisDataStore struct {
	client *redis.Client
}

func (rd RedisDataStore) Upload(urls string, malicious bool) error {
	eachUrl := strings.Split(urls, ",")
	for _, url := range eachUrl {
		log.Println("adding url ", url, "to data store.")
		err := rd.client.Set(url, malicious, 0).Err()
		if err != nil {
			responseString := fmt.Sprintf("Adding %s to redis failed", url)
			log.Println(responseString)
			return fmt.Errorf("%s", responseString)
		}
	}
	return nil
}

func (rd RedisDataStore) Query(url string) (bool, error) {
	val, err := rd.client.Get(url).Result()
	if err != nil {
		responseString := fmt.Sprintf("Querying redis for %s failed", url)
		log.Println(responseString)
		return false, fmt.Errorf("%s", responseString)
	} else if err == redis.Nil {
		responseString := fmt.Sprintf("url: %s not found in redis", url)
		log.Println(responseString)
		return false, fmt.Errorf("%s", responseString)
	}
	return strconv.ParseBool(val)
} 

func (rd RedisDataStore) Clear() error {
	val, err := rd.client.FlushAll().Result()
	log.Println(val)
	if err != nil {
		log.Println("error flushing redis")
		return err
	}
	return nil
}

// we don't really want/need this function for redis
func (rd RedisDataStore) String() {

}

func NewRedisDataStore() RedisDataStore {
	var dataStore RedisDataStore
	// just use the default redis conf for now
	// we get these env vars for free via linking docker containers
	address := fmt.Sprintf("%s:%s", 
						   os.Getenv("REDIS_PORT_6379_TCP_ADDR"), 
						   os.Getenv("REDIS_PORT_6379_TCP_PORT"))
	log.Println("trying to connect to redis at: ", address)
	dataStore.client = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	// we probably need some error checking in here

	log.Println("connected to redis")
	return dataStore
}