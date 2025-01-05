package main

import (
	"errors"
	"strconv"
	"strings"
)

const (
	errmsg = "invalid port specification"
)

func parsedash(in string, out *[]int) error {
	strs := strings.Split(in, "-")
	if len(strs) != 2 {
		return errors.New(errmsg)
	}
	start, err := strconv.Atoi(strs[0])
	if err != nil {
		return errors.New(errmsg)
	}
	end, err := strconv.Atoi(strs[1])
	if err != nil {
		return errors.New(errmsg)
	}
	if start < 1 || end > 65535 || start > end {
		return errors.New(errmsg)
	}
	for ; start < end ; start++ {
		*out  = append(*out, start)
	}
	return nil
}

func addport(port string, out *[]int) error {
	portint, err := strconv.Atoi(port)
	if err != nil {
		return errors.New(errmsg)
	}
	if portint < 1 || portint > 65535 {
		return errors.New(errmsg)
	}
	*out = append(*out, portint)
	return nil
}

func parse(port string) ([]int, error) {
	var ret []int
	if strings.Contains(port, ",") && strings.Contains(port, "-"){
		port := strings.Split(port, ",")
		for _, i := range port{
			if strings.Contains(i, "-"){
				if err := parsedash(i, &ret); err != nil{
					return nil, errors.New(errmsg)
				}
			} else {
				if err := addport(i, &ret); err != nil{
					return nil, errors.New(errmsg)
				}
			}
		}
	} else if strings.Contains(port, ","){
		port := strings.Split(port, ",")
		for _, i := range port{
			if err := addport(i, &ret); err != nil{
				return nil, errors.New(errmsg)
			}
		}
	} else if strings.Contains(port, "-"){
		if err := parsedash(port, &ret); err != nil{
			return nil, errors.New(errmsg)
		}
	} else {
		if err := addport(port, &ret); err != nil{
			return nil, errors.New(errmsg)
		}
	}
	return ret, nil
}