package redis

import (
	rg "github.com/garyburd/redigo/redis"
)

func (con *connection) Do(command string, args ...interface{}) (interface{}, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.Do(command, args...)
	} else {
		return con.c.Do(command, args...)
	}
}

/*
	Connection
*/
func (con *connection) Auth(password string) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.Auth(password)
	} else {
		res, err := rg.String(con.c.Do("AUTH", password))
		return getBool(res), err
	}
}

func (con *connection) Echo(message string) (string, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.Echo(message)
	} else {
		return rg.String(con.c.Do("ECHO", message))
	}
}

func (con *connection) Ping() (string, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.Ping()
	} else {
		return rg.String(con.c.Do("PING"))
	}
}

func (con *connection) Select(index int) (string, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.Select(index)
	} else {
		return rg.String(con.c.Do("SELECT", index))
	}
}

func (con *connection) Quit() (string, error) {
	return rg.String(con.c.Do("QUIT"))
}

/*
	Hashes
*/
func (con *connection) HDel(hashKey string, fields []string) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.HDel(hashKey, fields)
	} else {
		req := make([]interface{}, len(fields)+1)
		req[0] = hashKey
		for idx, val := range fields {
			req[idx+1] = val
		}

		return rg.Int(con.c.Do("HDEL", req...))
	}
}

func (con *connection) HExists(hashKey, field string) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.HExists(hashKey, field)
	} else {
		return rg.Bool(con.c.Do("HEXISTS", hashKey, field))
	}
}

func (con *connection) HGet(hashKey, field string) (string, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.HGet(hashKey, field)
	} else {
		return rg.String(con.c.Do("HGET", hashKey, field))
	}
}

func (con *connection) HGetFloat64(hashKey, field string) (float64, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.HGetFloat64(hashKey, field)
	} else {
		return rg.Float64(con.c.Do("HGET", hashKey, field))
	}
}

func (con *connection) HGetAll(hashKey string) (map[string]string, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.HGetAll(hashKey)
	} else {
		return rg.StringMap(con.c.Do("HGETALL", hashKey))
	}
}

func (con *connection) HGetAllFloat64(hashKey string) (map[string]float64, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.HGetAllFloat64(hashKey)
	} else {
		result, err := rg.StringMap(con.c.Do("HGETALL", hashKey))

		if err != nil {
			return map[string]float64{}, err
		}

		res := make(map[string]float64)

		for key, value := range result {
			res[key] = strToFloat64(value)
		}

		return res, nil
	}
}

func (con *connection) HIncrBy(hashKey, field string, increment int) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.HIncrBy(hashKey, field, increment)
	} else {
		return rg.Int(con.c.Do("HINCRBY", hashKey, field, increment))
	}
}

func (con *connection) HIncrByFloat(hashKey, field string, increment float64) (float64, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.HIncrByFloat(hashKey, field, increment)
	} else {
		return rg.Float64(con.c.Do("HINCRBYFLOAT", hashKey, field, increment))
	}
}

func (con *connection) HKeys(hashKey string) ([]string, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.HKeys(hashKey)
	} else {
		return rg.Strings(con.c.Do("HKEYS", hashKey))
	}
}

func (con *connection) HLen(hashKey string) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.HLen(hashKey)
	} else {
		return rg.Int(con.c.Do("HLEN", hashKey))
	}
}

func (con *connection) HMGet(hashKey string, fields []string) ([]string, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.HMGet(hashKey, fields)
	} else {
		req := make([]interface{}, len(fields)+1)
		req[0] = hashKey
		for idx, val := range fields {
			req[idx+1] = val
		}
		return rg.Strings(con.c.Do("HMGET", req...))
	}
}

func (con *connection) HMGetFloat64(hashKey string, fields []string) ([]float64, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.HMGetFloat64(hashKey, fields)
	} else {
		req := make([]interface{}, len(fields)+1)
		req[0] = hashKey
		for idx, val := range fields {
			req[idx+1] = val
		}

		result, err := rg.Strings(con.c.Do("HMGET", req...))

		if err != nil {
			return nil, err
		}

		res := make([]float64, len(result))

		for idx, value := range result {
			res[idx] = strToFloat64(value)
		}

		return res, nil
	}
}

func (con *connection) HMSet(hashKey string, fieldValue map[string]string) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.HMSet(hashKey, fieldValue)
	} else {
		req := make([]interface{}, len(fieldValue)*2+1)
		req[0] = hashKey
		idx := 1
		for name, value := range fieldValue {
			req[idx] = name
			idx++
			req[idx] = value
			idx++
		}

		res, err := rg.String(con.c.Do("HMSET", req...))
		return getBool(res), err
	}
}

func (con *connection) HMSetFloat64(hashKey string, fieldValue map[string]float64) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.HMSetFloat64(hashKey, fieldValue)
	} else {
		req := make([]interface{}, len(fieldValue)*2+1)
		req[0] = hashKey
		idx := 1
		for name, value := range fieldValue {
			req[idx] = name
			idx++
			req[idx] = value
			idx++
		}

		res, err := rg.String(con.c.Do("HMSET", req...))
		return getBool(res), err
	}
}

// func HScan() ()

func (con *connection) HSet(hashKey, field, value string) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.HSet(hashKey, field, value)
	} else {
		return rg.Int(con.c.Do("HSET", hashKey, field, value))
	}
}

func (con *connection) HSetFloat64(hashKey, field string, value float64) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.HSetFloat64(hashKey, field, value)
	} else {
		return rg.Int(con.c.Do("HSET", hashKey, field, value))
	}
}

func (con *connection) HSetNX(hashKey, field, value string) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.HSetNX(hashKey, field, value)
	} else {
		return rg.Int(con.c.Do("HSETNX", hashKey, field, value))
	}
}

func (con *connection) HStrLen(hashKey, field string) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.HStrLen(hashKey, field)
	} else {
		return rg.Int(con.c.Do("HSTRLEN", hashKey, field))
	}
}

func (con *connection) HVals(hashKey string) ([]string, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.HVals(hashKey)
	} else {
		return rg.Strings(con.c.Do("HVALS", hashKey))
	}
}

/*
	Keys
*/
func (con *connection) Del(keys []string) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.Del(keys)
	} else {
		req := make([]interface{}, len(keys))
		for idx, val := range keys {
			req[idx] = val
		}
		return rg.Int(con.c.Do("DEL", req...))
	}
}

func (con *connection) Dump(key string) (string, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.Dump(key)
	} else {
		return rg.String(con.c.Do("DUMP", key))
	}
}

func (con *connection) Exists(key string) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.Exists(key)
	} else {
		return rg.Bool(con.c.Do("EXISTS", key))
	}
}

func (con *connection) Expire(key string, seconds int) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.Expire(key, seconds)
	} else {
		return rg.Bool(con.c.Do("EXPIRE", key, seconds))
	}
}

func (con *connection) Expireat(key string, timestamp int64) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.Expireat(key, timestamp)
	} else {
		return rg.Bool(con.c.Do("EXPIREAT", key, timestamp))
	}
}

func (con *connection) Keys(pattern string) ([]string, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.Keys(pattern)
	} else {
		return rg.Strings(con.c.Do("KEYS", pattern))
	}
}

// Migrate()

func (con *connection) Move(key, db string) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.Move(key, db)
	} else {
		res, err := rg.String(con.c.Do("MOVE", key, db))
		return getBool(res), err
	}
}

// Object()

func (con *connection) Persist(key string) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.Persist(key)
	} else {
		return rg.Bool(con.c.Do("PERSIST", key))
	}
}

func (con *connection) PExpire(key string, millisec int64) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.PExpire(key, millisec)
	} else {
		return rg.Bool(con.c.Do("PEXPIRE", key, millisec))
	}
}

func (con *connection) PExpireat(key string, millisecTimestamp int64) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.PExpireat(key, millisecTimestamp)
	} else {
		return rg.Bool(con.c.Do("PEXPIREAT", key, millisecTimestamp))
	}
}

func (con *connection) PTTL(key string) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.PTTL(key)
	} else {
		return rg.Int(con.c.Do("PTTL", key))
	}
}

func (con *connection) RandomKey() (string, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.RandomKey()
	} else {
		return rg.String(con.c.Do("RANDOMKEY"))
	}
}

func (con *connection) Rename(key, newKey string) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.Rename(key, newKey)
	} else {
		res, err := rg.String(con.c.Do("Rename", key, newKey))
		return getBool(res), err
	}
}

func (con *connection) RenameNX(key, newKey string) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.RenameNX(key, newKey)
	} else {
		res, err := rg.String(con.c.Do("RenameNX", key, newKey))
		return getBool(res), err
	}
}

func (con *connection) Restore(key string, ttl int, serializedValue string) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.Restore(key, ttl, serializedValue)
	} else {
		res, err := rg.String(con.c.Do("RESTORE", key, ttl, serializedValue))
		return getBool(res), err
	}
}

func (con *connection) RestoreWithReplace(key string, ttl int, serializedValue, replace string) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.RestoreWithReplace(key, ttl, serializedValue, replace)
	} else {
		res, err := rg.String(con.c.Do("RESTORE", key, ttl, serializedValue, replace))
		return getBool(res), err
	}
}

// Scan()

func (con *connection) Sort(args ...interface{}) ([]string, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.Sort(args...)
	} else {
		return rg.Strings(con.c.Do("SORT", args...))
	}
}

func (con *connection) TTL(key string) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.TTL(key)
	} else {
		return rg.Int(con.c.Do("TTL", key))
	}
}

func (con *connection) Type(key string) (string, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.Type(key)
	} else {
		return rg.String(con.c.Do("TYPE", key))
	}
}

func (con *connection) Wait(numSlaves, ttl int) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.Wait(numSlaves, ttl)
	} else {
		return rg.Int(con.c.Do("Wait", numSlaves, ttl))
	}
}

/*
   Strings
*/
func (con *connection) Append(key, value string) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.Append(key, value)
	} else {
		return rg.Int(con.c.Do("APPEND", key, value))
	}
}

func (con *connection) BitCount(key string) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.BitCount(key)
	} else {
		return rg.Int(con.c.Do("BITCOUNT", key))
	}
}

func (con *connection) BitCountRange(key string, start, end int) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.BitCountRange(key, start, end)
	} else {
		return rg.Int(con.c.Do("BITCOUNTRANGE", key, start, end))
	}
}

func (con *connection) BitOP(key, destKey string, keys []interface{}) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.BitOP(key, destKey, keys)
	} else {
		return rg.Int(con.c.Do("BITOP", keys...))
	}
}

func (con *connection) BitPos(key string, start int) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.BitPos(key, start)
	} else {
		return rg.Int(con.c.Do("BITPOS", key, start))
	}
}

func (con *connection) BitPosRange(key string, start, end int) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.BitPosRange(key, start, end)
	} else {
		return rg.Int(con.c.Do("BITPOS", key, start, end))
	}
}

func (con *connection) Decr(key string) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.Decr(key)
	} else {
		return rg.Int(con.c.Do("DECR", key))
	}
}

func (con *connection) DecrBy(key string, decrement int) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.DecrBy(key, decrement)
	} else {
		return rg.Int(con.c.Do("DECRBY", key, decrement))
	}
}

func (con *connection) Get(key string) (string, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.Get(key)
	} else {
		return rg.String(con.c.Do("GET", key))
	}
}

func (con *connection) GetFloat64(key string) (float64, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.GetFloat64(key)
	} else {
		return rg.Float64(con.c.Do("GET", key))
	}
}

func (con *connection) GetBit(key string, offset int) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.GetBit(key, offset)
	} else {
		return rg.Int(con.c.Do("GETBIT", key, offset))
	}
}

func (con *connection) GetRange(key string, start, end int) (string, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.GetRange(key, start, end)
	} else {
		return rg.String(con.c.Do("GETRANGE", key, start, end))
	}
}

func (con *connection) GetSet(key, value string) (string, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.GetSet(key, value)
	} else {
		return rg.String(con.c.Do("GETSET", key, value))
	}
}

func (con *connection) Incr(key string) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.Incr(key)
	} else {
		return rg.Int(con.c.Do("INCR", key))
	}
}

func (con *connection) IncrBy(key string, increment int) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.IncrBy(key, increment)
	} else {
		return rg.Int(con.c.Do("INCRBY", key, increment))
	}
}

func (con *connection) IncrByFloat(key string, increment float64) (float64, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.IncrByFloat(key, increment)
	} else {
		return rg.Float64(con.c.Do("INCRBYFLOAT", key, increment))
	}
}

func (con *connection) MGet(keys []string) ([]string, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.MGet(keys)
	} else {
		req := make([]interface{}, len(keys))
		for idx, val := range keys {
			req[idx] = val
		}
		return rg.Strings(con.c.Do("MGET", req...))
	}
}

func (con *connection) MGetFloat64(keys []string) ([]float64, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.MGetFloat64(keys)
	} else {
		req := make([]interface{}, len(keys))
		for idx, val := range keys {
			req[idx] = val
		}

		result, err := rg.Strings(con.c.Do("MGET", req...))

		if err != nil {
			return nil, err
		}

		res := make([]float64, len(result))

		for idx, value := range result {
			res[idx] = strToFloat64(value)
		}

		return res, nil
	}
}

func (con *connection) MSet(keyValue map[string]string) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.MSet(keyValue)
	} else {
		req := make([]interface{}, len(keyValue)*2)
		idx := 0
		for name, value := range keyValue {
			req[idx] = name
			idx++
			req[idx] = value
			idx++
		}

		res, err := rg.String(con.c.Do("MSET", req...))
		return getBool(res), err
	}
}

func (con *connection) MSetFloat64(keyValue map[string]float64) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.MSetFloat64(keyValue)
	} else {
		req := make([]interface{}, len(keyValue)*2)
		idx := 0
		for name, value := range keyValue {
			req[idx] = name
			idx++
			req[idx] = value
			idx++
		}

		res, err := rg.String(con.c.Do("MSET", req...))
		return getBool(res), err
	}
}

func (con *connection) MSetNX(keyValue map[string]string) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.MSetNX(keyValue)
	} else {
		req := make([]interface{}, len(keyValue)*2)
		idx := 0
		for name, value := range keyValue {
			req[idx] = name
			idx++
			req[idx] = value
			idx++
		}

		return rg.Int(con.c.Do("MSETNX", req...))
	}
}

func (con *connection) PSetEX(key, value string, millisec int64) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer con.Release()
		return c.PSetEX(key, value, millisec)
	} else {
		res, err := rg.String(con.c.Do("PSETEX", key, millisec, value))
		return getBool(res), err
	}
}

func (con *connection) Set(key, value string) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.Set(key, value)
	} else {
		res, err := rg.String(con.c.Do("SET", key, value))
		return getBool(res), err
	}
}

func (con *connection) SetFloat64(key string, value float64) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.SetFloat64(key, value)
	} else {
		res, err := rg.String(con.c.Do("SET", key, value))
		return getBool(res), err
	}
}

func (con *connection) SetBit(key, value string, offset int) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.SetBit(key, value, offset)
	} else {
		return rg.Int(con.c.Do("SETBIT", key, value, offset))
	}
}

func (con *connection) SetEX(key, value string, seconds int) (bool, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.SetEX(key, value, seconds)
	} else {
		res, err := rg.String(con.c.Do("SETEX", key, seconds, value))
		return getBool(res), err
	}
}

func (con *connection) SetNX(key, value string) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.SetNX(key, value)
	} else {
		return rg.Int(con.c.Do("SETNX", key, value))
	}
}

func (con *connection) SetRange(key, value string, offset int) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.SetRange(key, value, offset)
	} else {
		return rg.Int(con.c.Do("SETRANGE", key, offset, value))
	}
}

func (con *connection) StrLen(key string) (int, error) {
	if con.p != nil {
		c, _ := con.GetConnection()
		defer c.Release()
		return c.StrLen(key)
	} else {
		return rg.Int(con.c.Do("STRLEN", key))
	}
}
