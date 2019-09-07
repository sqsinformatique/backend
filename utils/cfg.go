package utils

import (
	"errors"
	"fmt"
	"net"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	tagName = "environment"
)

var IP, _ = HostIP()

type envElement struct {
	Name         string
	DefaultValue string
}

func parseEnvTag(tag string) (res envElement) {
	args := strings.Split(tag, ",")
	res = envElement{
		Name: args[0],
	}

	if len(args) == 2 {
		res.DefaultValue = args[1]
	}

	return
}

func LoadFromEnv(cfg interface{}) {
	ps := reflect.ValueOf(cfg)

	s := ps.Elem()
	if s.Kind() == reflect.Struct {
		for i := 0; i < s.NumField(); i++ {
			field := s.Type().Field(i)
			// Get the field tag value
			tag := field.Tag.Get(tagName)
			tagElements := parseEnvTag(tag)

			// exported field
			f := s.FieldByName(field.Name)
			if f.IsValid() {
				if f.CanSet() {
					// change value
					if f.Kind() == reflect.Uint32 {
						var x int
						switch f.Interface().(type) {
						case Level:
							x = GetIntEnvVar(tagElements.Name, int(LevelByName(tagElements.DefaultValue)))
						}
						if !f.OverflowUint(uint64(x)) {
							f.SetUint(uint64(x))
						}
					}
					if f.Kind() == reflect.String {
						x := GetStrEnvVar(tagElements.Name, tagElements.DefaultValue)
						f.SetString(x)
					}
					if f.Kind() == reflect.Int64 {
						var x int
						switch f.Interface().(type) {
						case time.Duration:
							v, _ := strconv.ParseInt(tagElements.DefaultValue, 10, 64)
							x = int(GetTimeDurationEnvVar(tagElements.Name, time.Duration(v)))
						}
						if !f.OverflowInt(int64(x)) {
							f.SetInt(int64(x))
						}
					}
					if f.Kind() == reflect.Int {
						v, _ := strconv.ParseInt(tagElements.DefaultValue, 10, 64)
						x := GetIntEnvVar(tagElements.Name, int(v))
						if !f.OverflowInt(int64(x)) {
							f.SetInt(int64(x))
						}
					}
				}
			}
		}
	}
}

func GetStrEnvVar(name, defval string) string {
	val, res := os.LookupEnv(name)
	if !res {
		return defval
	}
	return val
}

func GetBoolEnvVar(name string, defval bool) bool {
	defStr := "false"
	if defval {
		defStr = "true"
	}
	str := GetStrEnvVar(name, defStr)
	val, err := strconv.ParseBool(str)
	if err != nil {
		return defval
	}
	return val
}

func GetIntEnvVar(name string, defval int) int {
	defStr := fmt.Sprint(defval)
	str := GetStrEnvVar(name, defStr)
	val, err := strconv.Atoi(str)
	if err != nil {
		return defval
	}
	return val
}

func GetTimeDurationEnvVar(name string, defval time.Duration) time.Duration {
	defStr := fmt.Sprint(defval.Nanoseconds())
	str := GetStrEnvVar(name, defStr)
	val, err := strconv.Atoi(str)
	if err != nil {
		return defval * time.Second
	}
	return time.Duration(val) * time.Second
}

func HostIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "0.0.0.0", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "0.0.0.0", errors.New("node is not connected")
}
