package IRedis

import (
	"context"
	"time"
)

type RedisFactory interface {
	// Connection Resolve a redis connection instance.
	Connection(name ...string) RedisConnection
}

type GeoPos struct {
	Longitude, Latitude float64
}

type BitCount struct {
	Start, End int64
}

type GeoLocation struct {
	Name                      string
	Longitude, Latitude, Dist float64
	GeoHash                   int64
}

type GeoRadiusQuery struct {
	Radius float64
	// Can be m, km, ft, or mi. Default is km.
	Unit        string
	WithCoord   bool
	WithDist    bool
	WithGeoHash bool
	Count       int
	// Can be ASC or DESC. Default is no sort order.
	Sort      string
	Store     string
	StoreDist string
}

type ZStore struct {
	Keys    []string
	Weights []float64
	// Can be SUM, MIN or MAX.
	Aggregate string
}

type Z struct {
	Score  float64
	Member interface{}
}

type ZRangeBy struct {
	Min, Max      string
	Offset, Count int64
}

// RedisSubscribeFunc 订阅给定的消息频道
// Subscribe to a given message channel.
type RedisSubscribeFunc func(message, channel string)

type RedisConnection interface {
	RedisConnectionCtx

	// Subscribe  to a set of given channels for messages.
	Subscribe(channels []string, closure RedisSubscribeFunc) error

	// PSubscribe subscribe to a set of given channels with wildcards.
	PSubscribe(channels []string, closure RedisSubscribeFunc) error

	// Command Run a command against the Redis database.
	Command(method string, args ...interface{}) (interface{}, error)

	// PubSubChannels List the currently active channels. Active means that the channel contains one or more subscribers (excluding clients receiving subscriptions from the pattern). If pattern is not provided, all channels are listed, otherwise only lists matching the specified global- Channels of type mode are listed.
	PubSubChannels(pattern string) ([]string, error)

	// PubSubNumSub List the number of subscribers to the specified channel (excluding client subscribers in subscription mode).
	PubSubNumSub(channels ...string) (map[string]int64, error)

	// PubSubNumPat Returns the number of subscribed schemas (implemented using the command PSUBSCRIBE). Note that this command returns not the number of clients subscribed to the schema, but the sum of all schemas subscribed by the client.
	PubSubNumPat() (int64, error)

	// Publish Send the information message to the specified channel .
	Publish(channel string, message interface{}) (int64, error)

	// Get Returns the value of the given key.
	Get(key string) (string, error)

	// MGet get the values of all the given keys.
	MGet(keys ...string) ([]interface{}, error)

	// GetBit For the string value stored in key, get the bit at the specified offset.
	GetBit(key string, offset int64) (int64, error)

	// BitOpAnd Take a logical union of one or more keys and save the result to destkey.
	BitOpAnd(destKey string, keys ...string) (int64, error)

	// BitOpNot Take the logical negation of the given key and save the result to destkey.
	BitOpNot(destKey string, key string) (int64, error)

	// BitOpOr Logical OR of one or more keys and save the result to destkey.
	BitOpOr(destKey string, keys ...string) (int64, error)

	// BitOpXor XOR one or more keys and save the result to destkey.
	BitOpXor(destKey string, keys ...string) (int64, error)

	// GetDel Returns the string value of the string corresponding to the key, and deletes the key.
	GetDel(key string) (string, error)

	// GetEx An expiration of zero removes the TTL associated with the key (i.e. GETEX key persist).
	GetEx(key string, expiration time.Duration) (string, error)

	// GetRange Returns the substring of the string value corresponding to the key, which is determined by the start and end displacements.
	GetRange(key string, start, end int64) (string, error)

	// GetSet Automatically map the key to the value and return the value corresponding to the original key. If the key exists but the corresponding value is not a string, return an error.
	GetSet(key string, value interface{}) (string, error)

	// ClientGetName returns the name of the connection.
	ClientGetName() (string, error)

	// StrLen Returns the length of the string value of key. If the key corresponds to a non-string type, an error is returned.
	StrLen(key string) (int64, error)

	// getter end
	// keys start

	// Keys Find all keys matching the given pattern  (regular expression).
	Keys(pattern string) ([]string, error)

	// Del Delete the specified batch of keys, if some keys in the deletion do not exist, they will be ignored directly.
	Del(keys ...string) (int64, error)

	// FlushAll Delete all data in all databases, note that not the current database, but all databases.
	FlushAll() (string, error)

	// FlushDB Delete all data in the current database.
	FlushDB() (string, error)

	// Dump Serialize the given key and return the serialized value.
	Dump(key string) (string, error)

	// Exists Returns whether the key exists.
	Exists(keys ...string) (int64, error)

	// Expire Set the expiration time of the key. After the time expires, the key will be automatically deleted.
	Expire(key string, expiration time.Duration) (bool, error)

	// ExpireAt The role of EXPIREAT is similar to that of EXPIRE, both are used to set the time-to-live for the key. The difference is that the time parameter accepted by the EXPIREAT command is the UNIX timestamp.
	ExpireAt(key string, tm time.Time) (bool, error)

	// PExpire This command works like the EXPIRE command, but it sets the key's lifetime in milliseconds instead of seconds like the EXPIRE command.
	PExpire(key string, expiration time.Duration) (bool, error)

	// PExpireAt PEXPIREAT This command is similar to the EXPIREAT command, but it sets the expiry unix timestamp of the key in milliseconds instead of seconds like EXPIREAT.
	PExpireAt(key string, tm time.Time) (bool, error)

	// Migrate Atomically transfer the key from the current instance to the specified database of the target instance. Once the transfer is successful, the key is guaranteed to appear on the target instance, and the key on the current instance will be deleted.
	Migrate(host, port, key string, db int, timeout time.Duration) (string, error)

	// Move  the key of the current database to the given database db.
	Move(key string, db int) (bool, error)

	// Persist Removes the time-to-live for a given key, converting the key from "volatile" (a key with a time-to-live) to "durable" (a key with no time-to-live that never expires).
	Persist(key string) (bool, error)

	// PTTL This command is similar to the TTL command, but it returns the remaining time to live for the key in milliseconds instead of seconds like the TTL command.
	PTTL(key string) (time.Duration, error)

	// TTL Returns the remaining expiration time for the key. This reflection capability allows Redis clients to check the remaining validity period of a given key in the dataset.
	TTL(key string) (time.Duration, error)

	// RandomKey Returns a random key from the current database.
	RandomKey() (string, error)

	// Rename key to newkey, returns an error if key is the same as newkey. If newkey already exists, the value will be overwritten.
	Rename(key, newKey string) (string, error)

	// RenameNX If and only if newkey does not exist, rename key to newkey, if key does not exist, return an error.
	RenameNX(key, newKey string) (bool, error)

	// Type Returns the data structure type of the value stored by the key, which can return different types such as string, list, set, zset and hash.
	Type(key string) (string, error)

	// Wait This command blocks the current client until all previous write commands have been successfully transmitted and acknowledged by the specified slaves. If it times out, specify in milliseconds, even if the specified slaves have not arrived, the command still returns.
	Wait(numSlaves int, timeout time.Duration) (int64, error)

	// Scan Iterate over the key collection in the current database.
	Scan(cursor uint64, match string, count int64) ([]string, uint64, error)

	// BitCount The number of bits whose statistics string is set to 1.
	BitCount(key string, count *BitCount) (int64, error)

	// keys end

	// Set setter start
	// Set the key key to the specified "string" value.
	Set(key string, value interface{}, expiration time.Duration) (string, error)

	// Append If key already exists and the value is a string, then this command appends value to the end of the original value. If the key does not exist, then it will first create a key with an empty string, and then perform the append operation, in which case APPEND will be similar to the SET operation.
	Append(key, value string) (int64, error)

	// MSet Map the given keys to their corresponding values. MSET will replace the existing value with the new value, just like a normal SET command. If you do not want to overwrite existing values, see the command MSETNX.
	MSet(values ...interface{}) (string, error)

	// MSetNX Map the given keys to their corresponding values. As long as a key already exists, MSETNX will not perform an operation. Because of this feature, MSETNX can achieve either all operations succeed or none of them are executed, which can be used to set different keys to represent different fields of a unique object.
	MSetNX(values ...interface{}) (bool, error)

	// SetNX Set the key to value, if the key does not exist, this case is equivalent to the SET command. When the key exists, do nothing, SETNX is short for "SET if Not eXists".
	SetNX(key string, value interface{}, expiration time.Duration) (bool, error)

	// SetEx Set the key corresponding to the string value, and set the key to expire after the given seconds time.
	SetEx(key string, value interface{}, expiration time.Duration) (string, error)

	// SetBit set or clear the bit value of the key's value (string) at offset.
	SetBit(key string, offset int64, value int) (int64, error)

	// BitPos Returns the first bit in the string that is set to 1 or 0.
	BitPos(key string, bit int64, pos ...int64) (int64, error)

	// SetRange The function of this command is to overwrite the part of the string corresponding to the key, starting from the specified offset, overwriting the length of the value. If the offset is longer than the string corresponding to the current key, then the string is followed by 0 to achieve the offset. Non-existing keys are considered empty strings, so this command ensures that the key has a string large enough to set the value at offset.
	SetRange(key string, offset int64, value string) (int64, error)

	// Incr Performs an atomic increment operation on the value stored at the specified key. If the specified key does not exist, its value will be set to 0 before performing the incr operation.
	Incr(key string) (int64, error)

	// Decr Subtract 1 from the number corresponding to the key. If the key does not exist, the value corresponding to the key will be set to 0 before the operation. Returns an error if key has a value of the wrong type or is a string that cannot be represented as a number. This operation supports up to 64-bit signed integer numbers.
	Decr(key string) (int64, error)

	// IncrBy Add decrement to the number corresponding to the key. If the key does not exist, the key will be set to 0 before the operation. Returns an error if the value of key has the wrong type or is a string that cannot be represented as a number. This operation supports up to 64-bit signed positive numbers.
	IncrBy(key string, value int64) (int64, error)

	// DecrBy Subtract decrement from the number corresponding to the key. If the key does not exist, the key will be set to 0 before the operation. Returns an error if the value of key has the wrong type or is a string that cannot be represented as a number. This operation supports up to 64-bit signed positive numbers.
	DecrBy(key string, value int64) (int64, error)

	// IncrByFloat Increase the value of the floating-point number (stored in the string) by specifying the floating-point number key. When the key does not exist, set its value to 0 before operating.
	IncrByFloat(key string, value float64) (float64, error)

	// setter end

	// HGet hash start
	// Returns the value associated with the field in the hash set specified by key.
	HGet(key, field string) (string, error)

	// HGetAll Returns all fields and values in the hash set specified by key. In the return value, next to each field name is its value, so the length of the return value is twice the size of the hash set.
	HGetAll(key string) (map[string]string, error)

	// HMGet Returns the value of the specified field in the hash set specified by key.
	HMGet(key string, fields ...string) ([]interface{}, error)

	// HKeys Returns the names of all fields in the hash set specified by key.
	HKeys(key string) ([]string, error)

	// HLen Returns the number of fields contained in the hashset specified by key.
	HLen(key string) (int64, error)

	// HRandField A new command added in Redis 6.2 to randomly obtain the fields in the specified hash table.
	HRandField(key string, count int) ([]string, error)

	// HScan Used for incremental iteration to get all fields in the hash table and return their field names and their values.
	HScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error)

	// HValues Returns the values of all fields in the hashset specified by key.
	HValues(key string) ([]string, error)

	// HSet Sets the value of the specified field in the hash set specified by key.
	HSet(key string, values ...interface{}) (int64, error)

	// HSetNX Sets the field's value only if the specified field does not exist in the hash set specified by key. If the hash set specified by key does not exist, a new hash set is created and associated with key. If the field already exists, this operation has no effect.
	HSetNX(key, field string, value interface{}) (bool, error)

	// HMSet Sets the value of the specified field in the hash set specified by key. This command will rewrite all fields present in the hashset. If the hash set specified by key does not exist, a new hash set will be created and associated with key.
	HMSet(key string, values ...interface{}) (bool, error)

	// HDel Removes the specified domain from the hash set specified by key. Domains that do not exist in the hashset will be ignored.
	HDel(key string, fields ...string) (int64, error)

	// HExists Returns whether the field exists in the hash.
	HExists(key string, field string) (bool, error)

	// HIncrBy Increments the value of the specified field in the hash set specified by key. If key does not exist, a new hash set is created and associated with key. If the field does not exist, the value of the field is set to 0 before the operation is performed HINCRBY The range of supported values is limited to 64-bit signed integers.
	HIncrBy(key string, field string, value int64) (int64, error)

	// HIncrByFloat Performs an increment of float type for the field field value of the hash of the specified key. If field does not exist, set to 0 before performing the operation.
	HIncrByFloat(key string, field string, value float64) (float64, error)

	// hash end

	// SAdd set start
	// Add one or more specified member elements to the set key. If the specified one or more member elements already exist in the set key, it will be ignored. If the set key does not exist, create a new set key and add the member element to the set key. If the type of key is not a collection, an error is returned.
	SAdd(key string, members ...interface{}) (int64, error)

	// SCard Returns the cardinality of the key stored in the collection (the number of elements in the collection).
	SCard(key string) (int64, error)

	// SDiff Returns the elements of the difference between a set and the given set.
	SDiff(keys ...string) ([]string, error)

	// SDiffStore This command is similar to SDIFF, except that the command does not return a result set, but stores the result in the destination set. If the destination already exists, it will.
	SDiffStore(destination string, keys ...string) (int64, error)

	// SInter Returns the intersection of all members of the specified set.
	SInter(keys ...string) ([]string, error)

	// SInterStore This command is similar to the SINTER command, but instead of returning the result set directly, it stores the result in the destination collection. If the destination collection exists, it will be overwritten.
	SInterStore(destination string, keys ...string) (int64, error)

	// SIsMember Returns whether member member is a member of the stored collection key.
	SIsMember(key string, member interface{}) (bool, error)

	// SMembers Returns all elements of the key collection.
	SMembers(key string) ([]string, error)

	// SRem Removes the specified element from the key set. If the specified element is not an element in the key set, it is ignored. If the key set does not exist, it is treated as an empty set, and the command returns 0. If the type of key is not a set, then returns an error.
	SRem(key string, members ...interface{}) (int64, error)

	// SPopN Redis `SPOP key count` command.Remove and return multiple random elements from the collection stored at key.
	SPopN(key string, count int64) ([]string, error)

	// SPop Remove from the collection stored at key and return a.
	SPop(key string) (string, error)

	// SRandMemberN Redis `SRANDMEMBER key count` command. Randomly returns multiple elements in the key collection.
	SRandMemberN(key string, count int64) ([]string, error)

	// SMove Moves the member from the source collection to the destination collection. For other clients, the element will appear as a member of the source or destination collection at a specific time.
	SMove(source, destination string, member interface{}) (bool, error)

	// SRandMember Randomly returns an element in the key collection.
	SRandMember(key string) (string, error)

	// SUnion Returns all members of the union of the given collections.
	SUnion(keys ...string) ([]string, error)

	// SUnionStore his command works like the SUNION command, except that it does not return a result set, but stores the result in the destination set. If the destination already exists, it will be overwritten.
	SUnionStore(destination string, keys ...string) (int64, error)

	// set end

	// GeoAdd geo start
	// Adds the specified geospatial location (latitude, longitude, name) to the specified key.
	GeoAdd(key string, geoLocation ...*GeoLocation) (int64, error)

	// GeoHash Returns a Geohash representation of one or more location elements.
	GeoHash(key string, members ...string) ([]string, error)

	// GeoPos Returns the location (longitude and latitude) of all given location elements from key.
	GeoPos(key string, members ...string) ([]*GeoPos, error)

	// GeoDist Returns the distance between two given locations. If one of the two positions does not exist, the command returns null.
	GeoDist(key string, member1, member2, unit string) (float64, error)

	// GeoRadius Taking the given latitude and longitude as the center, among the position elements contained in the return key, all position elements whose distance from the center does not exceed the given maximum distance.
	GeoRadius(key string, longitude, latitude float64, query *GeoRadiusQuery) ([]GeoLocation, error)

	// GeoRadiusStore is a writing GEORADIUS command.
	GeoRadiusStore(key string, longitude, latitude float64, query *GeoRadiusQuery) (int64, error)

	// GeoRadiusByMember This command is the same as the GEORADIUS command, which can find the elements within the specified range, but the center point of GEORADIUSBYMEMBER is determined by the given position element, instead of using the input latitude and longitude to determine the center point specification like GEORADIUS The member's location is used as the center of the query.
	GeoRadiusByMember(key, member string, query *GeoRadiusQuery) ([]GeoLocation, error)

	// GeoRadiusByMemberStore is a writing GEORADIUSBYMEMBER command.
	GeoRadiusByMemberStore(key, member string, query *GeoRadiusQuery) (int64, error)

	// geo end

	// BLPop lists start
	// BLPOP is a popping primitive for blocking lists. It is the blocking version of the command LPOP because the connection will be blocked by the BLPOP command when there are no elements to pop in the given list. When multiple key parameters are given, each list is checked in order of the parameter keys, and the head element of the first non-empty list is popped.
	BLPop(timeout time.Duration, keys ...string) ([]string, error)

	// BRPop BRPOP is a blocking list pop primitive. It is a blocking version of RPOP, as this command blocks the connection if no elements can be popped from the given list. This command looks through the list in the order given by the keys and pops an element at the end of the first non-empty list found.
	BRPop(timeout time.Duration, keys ...string) ([]string, error)

	// BRPopLPush BRPOPLPUSH is a blocking version of RPOPLPUSH. When source contains elements, this command behaves exactly like RPOPLPUSH. When source is empty, Redis will block the connection until another client pushes an element in or the timeout expires. A timeout of 0 can be used to block the client indefinitely.
	BRPopLPush(source, destination string, timeout time.Duration) (string, error)

	// LIndex Returns the index of the element in the list index stored in key. Subscripts are 0-based, so 0 is the first element, 1 is the second, and so on. Negative indices are used to specify elements indexed from the end of the list. In this method, -1 means the last element, -2 means the second-to-last element, and so on. An error is returned when the value at the key position is not a list.
	LIndex(key string, index int64) (string, error)

	// LInsert Inserts the value in the list stored at key before or after the pivot value.
	LInsert(key, op string, pivot, value interface{}) (int64, error)

	// LLen Returns the length of the list stored in key. If the key does not exist, it is treated as an empty list, and the return length is 0. An error is returned when the value stored in key is not a list.
	LLen(key string) (int64, error)

	// LPop Removes and returns the first element of the list corresponding to key.
	LPop(key string) (string, error)

	// LPush Inserts all specified values at the head of the list stored at key. If the key does not exist, an empty list is created before the push operation. If the value corresponding to key is not a list, an error will be returned.
	LPush(key string, values ...interface{}) (int64, error)

	// LPushX Insert value at the head of the list below the key only if the key already exists and a list exists. Contrary to LPUSH, no operation is performed when the key does not exist.
	LPushX(key string, values ...interface{}) (int64, error)

	// LRange Returns the elements in the specified range stored in the list at key.
	LRange(key string, start, stop int64) ([]string, error)

	// LRem Removes the first count occurrences of the element whose value is value from the list stored at key.
	LRem(key string, count int64, value interface{}) (int64, error)

	// LSet Sets the value of the list element at index position to value.
	LSet(key string, index int64, value interface{}) (string, error)

	// LTrim Trim an existing list so that the list contains only the specified elements in the specified range.
	LTrim(key string, start, stop int64) (string, error)

	// RPop Push out and return the last element of the list stored at key.
	RPop(key string) (string, error)

	// RPopCount Extracts the specified number of elements on the right and returns the list stored at key.
	RPopCount(key string, count int) ([]string, error)

	// RPopLPush Atomically returns and removes the last element of the list stored in source (the tail of the list) and places that element in the position of the first element of the list stored in destination (the head of the list).
	RPopLPush(source, destination string) (string, error)

	// RPush Inserts all specified values to the end of the list stored at key. If the key does not exist, an empty list is created and then the push operation is performed. When key does not hold a list, an error is returned.
	RPush(key string, values ...interface{}) (int64, error)

	// RPushX Inserts the value value at the end of the list key if and only if key exists and is a list. Contrary to the RPUSH command, the RPUSHX command does nothing when the key does not exist.
	RPushX(key string, values ...interface{}) (int64, error)

	// lists end

	// scripting start

	// Eval Execute LUA scripts on the server side.
	Eval(script string, keys []string, args ...interface{}) (interface{}, error)

	// EvalSha Evaluate the script cached in the server against the given SHA1 checksum. Caching scripts to the server can be done with the SCRIPT LOAD command. The rest of this command, such as the way parameters are passed in, are the same as the EVAL command.
	EvalSha(sha1 string, keys []string, args ...interface{}) (interface{}, error)

	// ScriptExists Check if the script exists in the script cache.
	ScriptExists(hashes ...string) ([]bool, error)

	// ScriptFlush Flush the Lua scripts cache.
	ScriptFlush() (string, error)

	// ScriptKill Kills the currently running Lua script, if and only if the script has not performed any write operations, this command will take effect.
	ScriptKill() (string, error)

	// ScriptLoad Adds the script script to the script cache, but does not execute the script immediately. After the script has been added to the cache, the script can be called with the SHA1 checksum of the script via the EVALSHA command.
	ScriptLoad(script string) (string, error)

	// scripting end

	// zset start

	// ZAdd Adds all specified members to the keyed sorted set (sorted set).
	ZAdd(key string, members ...Z) (int64, error)

	// ZCard Returns the number of elements in the sorted set for key.
	ZCard(key string) (int64, error)

	// ZCount Returns the members of the sorted set key whose score value is between min and max (by default, the score value is equal to min or max). For details on how to use the parameters min and max, please refer to the ZRANGEBYSCORE command.
	ZCount(key, min, max string) (int64, error)

	// ZIncrBy Increment is added to the score value of the member member of the sorted set key. If a member does not exist in the key, add a member to the key with a score of increment (as if the previous score was 0.0). If the key does not exist, create an ordered set containing only the specified member members. Increment is added to the score value of the member member of the sorted set key. If a member does not exist in the key, add a member to the key with a score of increment (as if the previous score was 0.0). If the key does not exist, create an ordered set containing only the specified member members.
	ZIncrBy(key string, increment float64, member string) (float64, error)

	// ZInterStore Computes the intersection of the given numkeys sorted sets and places the result in destination. Before the key and other parameters to be calculated are given, the number of keys (numberkeys) must be given.
	ZInterStore(destination string, store *ZStore) (int64, error)

	// ZLexCount Counts the number of members between the specified members in a sorted set.
	ZLexCount(key, min, max string) (int64, error)

	// ZPopMax Removes and returns at most count members with the highest score in the sorted set key.
	ZPopMax(key string, count ...int64) ([]Z, error)

	// ZPopMin Removes and returns at most count members with the lowest score in the sorted set key.
	ZPopMin(key string, count ...int64) ([]Z, error)

	// ZRange Returns the specified range of elements stored in the sorted set key. The returned elements can be thought of as ordered from lowest to highest score. If the scores are the same, they will be sorted lexicographically.
	ZRange(key string, start, stop int64) ([]string, error)

	// ZRangeByLex Returns the members in the specified member range, sorted in the positive lexicographic order of the members, and the scores must be the same. In some business scenarios, when a string array needs to be sorted according to the lexicographical order of names, a data structure such as SortSet in Redis can be used for processing.
	ZRangeByLex(key string, opt *ZRangeBy) ([]string, error)

	// ZRevRangeByLex Returns the members in the specified member range, sorted in reverse lexicographic order of the members, and the scores must be the same.
	ZRevRangeByLex(key string, opt *ZRangeBy) ([]string, error)

	// ZRangeByScore Returns a list of members of the ordered collection with the specified score interval. The members of the ordered set are arranged in order of increasing score value (from small to large).
	ZRangeByScore(key string, opt *ZRangeBy) ([]string, error)

	// ZRank Returns the rank of the specified member in the sorted set. The ordered set members are arranged in order of increasing score value (from small to large).
	ZRank(key, member string) (int64, error)

	// ZRem Removes one or more members from the sorted set, non-existing members are ignored.
	ZRem(key string, members ...interface{}) (int64, error)

	// ZRemRangeByLex Removes all members of the given dictionary range in the sorted set.
	ZRemRangeByLex(key, min, max string) (int64, error)

	// ZRemRangeByRank Removes all members from an ordered set, specifying a rank interval.
	ZRemRangeByRank(key string, start, stop int64) (int64, error)

	// ZRemRangeByScore Removes all members from an ordered set with a specified score interval.
	ZRemRangeByScore(key, min, max string) (int64, error)

	// ZRevRange Returns a sorted set with members in the specified range. The positions of the members are arranged in decreasing order of score value (from largest to smallest). Members with the same score value are arranged in reverse lexicographical order.
	ZRevRange(key string, start, stop int64) ([]string, error)

	// ZRevRangeByScore Returns all members of the sorted set within the specified score interval. The ordered set members are arranged in order of decreasing score value (largest to smallest). Members with the same score value are arranged in reverse lexicographical order.
	ZRevRangeByScore(key string, opt *ZRangeBy) ([]string, error)

	// ZRevRank Returns the rank of the members in the sorted set. The members of the ordered set are sorted by decreasing score value (large to small).
	ZRevRank(key, member string) (int64, error)

	// ZScore Returns the fractional value of the members in the sorted set. Returns nil if the member element is not a member of the sorted set key, or if the key does not exist.
	ZScore(key, member string) (float64, error)

	// ZUnionStore Computes the union of one or more ordered sets given, where the number of given keys must be specified with the numkeys parameter, and stores the union (result set) in destination .
	ZUnionStore(key string, store *ZStore) (int64, error)

	// ZScan Used to iterate over elements in a sorted set (including element members and element scores).
	ZScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error)
}

type RedisConnectionCtx interface {

	// SubscribeWithContext to a set of given channels for messages.
	SubscribeWithContext(ctx context.Context, channels []string, closure RedisSubscribeFunc) error

	// PSubscribeWithContext SubscribeWithContext to a set of given channels with wildcards.
	PSubscribeWithContext(ctx context.Context, channels []string, closure RedisSubscribeFunc) error

	// CommandWithContext Run a CommandWithContext against the Redis database.
	CommandWithContext(ctx context.Context, method string, args ...interface{}) (interface{}, error)

	PubSubChannelsWithContext(ctx context.Context, pattern string) ([]string, error)

	PubSubNumSubWithContext(ctx context.Context, channels ...string) (map[string]int64, error)

	PubSubNumPatWithContext(ctx context.Context) (int64, error)

	PublishWithContext(ctx context.Context, channel string, message interface{}) (int64, error)

	// GetWithContext Returns the value of the given key.
	GetWithContext(ctx context.Context, key string) (string, error)

	// MGetWithContext get the values of all the given keys.
	MGetWithContext(ctx context.Context, keys ...string) ([]interface{}, error)

	// GetBitWithContext For the string value stored in key, get the bit at the specified offset.
	GetBitWithContext(ctx context.Context, key string, offset int64) (int64, error)

	BitOpAndWithContext(ctx context.Context, destKey string, keys ...string) (int64, error)

	BitOpNotWithContext(ctx context.Context, destKey string, key string) (int64, error)

	BitOpOrWithContext(ctx context.Context, destKey string, keys ...string) (int64, error)

	BitOpXorWithContext(ctx context.Context, destKey string, keys ...string) (int64, error)

	GetDelWithContext(ctx context.Context, key string) (string, error)

	GetExWithContext(ctx context.Context, key string, expiration time.Duration) (string, error)

	GetRangeWithContext(ctx context.Context, key string, start, end int64) (string, error)

	GetSetWithContext(ctx context.Context, key string, value interface{}) (string, error)

	ClientGetNameWithContext(ctx context.Context) (string, error)

	StrLenWithContext(ctx context.Context, key string) (int64, error)

	// getter end
	// keys start

	KeysWithContext(ctx context.Context, pattern string) ([]string, error)

	DelWithContext(ctx context.Context, keys ...string) (int64, error)

	FlushAllWithContext(ctx context.Context) (string, error)

	FlushDBWithContext(ctx context.Context) (string, error)

	DumpWithContext(ctx context.Context, key string) (string, error)

	ExistsWithContext(ctx context.Context, keys ...string) (int64, error)

	ExpireWithContext(ctx context.Context, key string, expiration time.Duration) (bool, error)

	ExpireAtWithContext(ctx context.Context, key string, tm time.Time) (bool, error)

	PExpireWithContext(ctx context.Context, key string, expiration time.Duration) (bool, error)

	PExpireAtWithContext(ctx context.Context, key string, tm time.Time) (bool, error)

	MigrateWithContext(ctx context.Context, host, port, key string, db int, timeout time.Duration) (string, error)

	MoveWithContext(ctx context.Context, key string, db int) (bool, error)

	PersistWithContext(ctx context.Context, key string) (bool, error)

	PTTLWithContext(ctx context.Context, key string) (time.Duration, error)

	TTLWithContext(ctx context.Context, key string) (time.Duration, error)

	RandomKeyWithContext(ctx context.Context) (string, error)

	RenameWithContext(ctx context.Context, key, newKey string) (string, error)

	RenameNXWithContext(ctx context.Context, key, newKey string) (bool, error)

	TypeWithContext(ctx context.Context, key string) (string, error)

	WaitWithContext(ctx context.Context, numSlaves int, timeout time.Duration) (int64, error)

	ScanWithContext(ctx context.Context, cursor uint64, match string, count int64) ([]string, uint64, error)

	BitCountWithContext(ctx context.Context, key string, count *BitCount) (int64, error)

	// keys end

	// SetWithContext setter start
	SetWithContext(ctx context.Context, key string, value interface{}, expiration time.Duration) (string, error)

	AppendWithContext(ctx context.Context, key, value string) (int64, error)

	MSetWithContext(ctx context.Context, values ...interface{}) (string, error)

	MSetNXWithContext(ctx context.Context, values ...interface{}) (bool, error)

	SetNXWithContext(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error)

	SetExWithContext(ctx context.Context, key string, value interface{}, expiration time.Duration) (string, error)

	SetBitWithContext(ctx context.Context, key string, offset int64, value int) (int64, error)

	BitPosWithContext(ctx context.Context, key string, bit int64, pos ...int64) (int64, error)

	SetRangeWithContext(ctx context.Context, key string, offset int64, value string) (int64, error)

	IncrWithContext(ctx context.Context, key string) (int64, error)

	DecrWithContext(ctx context.Context, key string) (int64, error)

	IncrByWithContext(ctx context.Context, key string, value int64) (int64, error)

	DecrByWithContext(ctx context.Context, key string, value int64) (int64, error)

	IncrByFloatWithContext(ctx context.Context, key string, value float64) (float64, error)

	// setter end

	// HGetWithContext hash start
	HGetWithContext(ctx context.Context, key, field string) (string, error)

	HGetAllWithContext(ctx context.Context, key string) (map[string]string, error)

	HMGetWithContext(ctx context.Context, key string, fields ...string) ([]interface{}, error)

	HKeysWithContext(ctx context.Context, key string) ([]string, error)

	HLenWithContext(ctx context.Context, key string) (int64, error)

	HRandFieldWithContext(ctx context.Context, key string, count int) ([]string, error)

	HScanWithContext(ctx context.Context, key string, cursor uint64, match string, count int64) ([]string, uint64, error)

	HValuesWithContext(ctx context.Context, key string) ([]string, error)

	HSetWithContext(ctx context.Context, key string, values ...interface{}) (int64, error)

	HSetNXWithContext(ctx context.Context, key, field string, value interface{}) (bool, error)

	HMSetWithContext(ctx context.Context, key string, values ...interface{}) (bool, error)

	HDelWithContext(ctx context.Context, key string, fields ...string) (int64, error)

	HExistsWithContext(ctx context.Context, key string, field string) (bool, error)

	HIncrByWithContext(ctx context.Context, key string, field string, value int64) (int64, error)

	HIncrByFloatWithContext(ctx context.Context, key string, field string, value float64) (float64, error)

	// hash end

	// SAddWithContext set start
	SAddWithContext(ctx context.Context, key string, members ...interface{}) (int64, error)

	SCardWithContext(ctx context.Context, key string) (int64, error)

	SDiffWithContext(ctx context.Context, keys ...string) ([]string, error)

	SDiffStoreWithContext(ctx context.Context, destination string, keys ...string) (int64, error)

	SInterWithContext(ctx context.Context, keys ...string) ([]string, error)

	SInterStoreWithContext(ctx context.Context, destination string, keys ...string) (int64, error)

	SIsMemberWithContext(ctx context.Context, key string, member interface{}) (bool, error)

	SMembersWithContext(ctx context.Context, key string) ([]string, error)

	SRemWithContext(ctx context.Context, key string, members ...interface{}) (int64, error)

	SPopNWithContext(ctx context.Context, key string, count int64) ([]string, error)

	SPopWithContext(ctx context.Context, key string) (string, error)

	SRandMemberNWithContext(ctx context.Context, key string, count int64) ([]string, error)

	SMoveWithContext(ctx context.Context, source, destination string, member interface{}) (bool, error)

	SRandMemberWithContext(ctx context.Context, key string) (string, error)

	SUnionWithContext(ctx context.Context, keys ...string) ([]string, error)

	SUnionStoreWithContext(ctx context.Context, destination string, keys ...string) (int64, error)

	// set end

	// geo start

	GeoAddWithContext(ctx context.Context, key string, geoLocation ...*GeoLocation) (int64, error)

	GeoHashWithContext(ctx context.Context, key string, members ...string) ([]string, error)

	GeoPosWithContext(ctx context.Context, key string, members ...string) ([]*GeoPos, error)

	GeoDistWithContext(ctx context.Context, key string, member1, member2, unit string) (float64, error)

	GeoRadiusWithContext(ctx context.Context, key string, longitude, latitude float64, query *GeoRadiusQuery) ([]GeoLocation, error)

	GeoRadiusStoreWithContext(ctx context.Context, key string, longitude, latitude float64, query *GeoRadiusQuery) (int64, error)

	GeoRadiusByMemberWithContext(ctx context.Context, key, member string, query *GeoRadiusQuery) ([]GeoLocation, error)

	GeoRadiusByMemberStoreWithContext(ctx context.Context, key, member string, query *GeoRadiusQuery) (int64, error)

	// geo end

	// lists start

	BLPopWithContext(ctx context.Context, timeout time.Duration, keys ...string) ([]string, error)

	BRPopWithContext(ctx context.Context, timeout time.Duration, keys ...string) ([]string, error)

	BRPopLPushWithContext(ctx context.Context, source, destination string, timeout time.Duration) (string, error)

	LIndexWithContext(ctx context.Context, key string, index int64) (string, error)

	LInsertWithContext(ctx context.Context, key, op string, pivot, value interface{}) (int64, error)

	LLenWithContext(ctx context.Context, key string) (int64, error)

	LPopWithContext(ctx context.Context, key string) (string, error)

	LPushWithContext(ctx context.Context, key string, values ...interface{}) (int64, error)

	LPushXWithContext(ctx context.Context, key string, values ...interface{}) (int64, error)

	LRangeWithContext(ctx context.Context, key string, start, stop int64) ([]string, error)

	LRemWithContext(ctx context.Context, key string, count int64, value interface{}) (int64, error)

	LSetWithContext(ctx context.Context, key string, index int64, value interface{}) (string, error)

	LTrimWithContext(ctx context.Context, key string, start, stop int64) (string, error)

	RPopWithContext(ctx context.Context, key string) (string, error)

	RPopCountWithContext(ctx context.Context, key string, count int) ([]string, error)

	RPopLPushWithContext(ctx context.Context, source, destination string) (string, error)

	RPushWithContext(ctx context.Context, key string, values ...interface{}) (int64, error)

	RPushXWithContext(ctx context.Context, key string, values ...interface{}) (int64, error)

	// lists end

	// EvalWithContext scripting start
	EvalWithContext(ctx context.Context, script string, keys []string, args ...interface{}) (interface{}, error)

	EvalShaWithContext(ctx context.Context, sha1 string, keys []string, args ...interface{}) (interface{}, error)

	ScriptExistsWithContext(ctx context.Context, hashes ...string) ([]bool, error)

	ScriptFlushWithContext(ctx context.Context) (string, error)

	ScriptKillWithContext(ctx context.Context) (string, error)

	ScriptLoadWithContext(ctx context.Context, script string) (string, error)

	// scripting end

	// zset start

	ZAddWithContext(ctx context.Context, key string, members ...Z) (int64, error)

	ZCardWithContext(ctx context.Context, key string) (int64, error)

	ZCountWithContext(ctx context.Context, key, min, max string) (int64, error)

	ZIncrByWithContext(ctx context.Context, key string, increment float64, member string) (float64, error)

	ZInterStoreWithContext(ctx context.Context, destination string, store *ZStore) (int64, error)

	ZLexCountWithContext(ctx context.Context, key, min, max string) (int64, error)

	ZPopMaxWithContext(ctx context.Context, key string, count ...int64) ([]Z, error)

	ZPopMinWithContext(ctx context.Context, key string, count ...int64) ([]Z, error)

	ZRangeWithContext(ctx context.Context, key string, start, stop int64) ([]string, error)

	ZRangeByLexWithContext(ctx context.Context, key string, opt *ZRangeBy) ([]string, error)

	ZRevRangeByLexWithContext(ctx context.Context, key string, opt *ZRangeBy) ([]string, error)

	ZRangeByScoreWithContext(ctx context.Context, key string, opt *ZRangeBy) ([]string, error)

	ZRankWithContext(ctx context.Context, key, member string) (int64, error)

	ZRemWithContext(ctx context.Context, key string, members ...interface{}) (int64, error)

	ZRemRangeByLexWithContext(ctx context.Context, key, min, max string) (int64, error)

	ZRemRangeByRankWithContext(ctx context.Context, key string, start, stop int64) (int64, error)

	ZRemRangeByScoreWithContext(ctx context.Context, key, min, max string) (int64, error)

	ZRevRangeWithContext(ctx context.Context, key string, start, stop int64) ([]string, error)

	ZRevRangeByScoreWithContext(ctx context.Context, key string, opt *ZRangeBy) ([]string, error)

	ZRevRankWithContext(ctx context.Context, key, member string) (int64, error)

	ZScoreWithContext(ctx context.Context, key, member string) (float64, error)

	ZUnionStoreWithContext(ctx context.Context, key string, store *ZStore) (int64, error)

	ZScanWithContext(ctx context.Context, key string, cursor uint64, match string, count int64) ([]string, uint64, error)
}
