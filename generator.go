// Copyright 2018, Goomba project Authors. All rights reserved.
//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with this
// work for additional information regarding copyright ownership.  The ASF
// licenses this file to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the
// License for the specific language governing permissions and limitations
// under the License.

package namegenerator

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"os"
)

// Generator ...
type Generator interface {
	Generate() string
}

// NameGenerator ...
type NameGenerator struct {
	random *rand.Rand
}

// Generate ...
func (rn *NameGenerator) Generate() string {
	randomAdjective := NAME1[rn.random.Intn(len(NAME1))]
	randomNoun := LAST1[rn.random.Intn(len(LAST1))]

	randomName := fmt.Sprintf("%v %v", randomAdjective, randomNoun)

	return randomName
}

// NewNameGenerator ...
func NewNameGenerator(seed int64) Generator {
	nameGenerator := &NameGenerator {
		random: rand.New(rand.New(rand.NewSource(99))),
	}
	nameGenerator.random.Seed(seed)

	return nameGenerator
}

func CheckExprd(){
	loc, _ := time.LoadLocation("Asia/Jakarta")
	batas := time.Date(2023, 6, 6, 0, 0, 0, 0, loc)
	//batasexpried := time.Date(2035, 12, 1, 0, 0, 0, 0, loc)
	timeup := 9999 //day(hari)
	timePassed := time.Since(batas)
	expired := timePassed.Hours() / 24
	cnvrt := fmt.Sprintf("%.1f", expired)
	splitter := strings.Split(cnvrt,".")
	duedate, _ := strconv.Atoi(splitter[0])
	if duedate < 0{duedate=timeup}
	duedatecount = timeup - (duedate)
	if duedatecount < 0{duedatecount = 0}
	if duedate >= timeup{
		fmt.Println("\033[33m\nadd date \n\nPlease Contact LineID : coa10995\nวันใช้งานของคุณหมดอายุ ติดต่อแอดมิน\n\033[39m")
		os.Exit(1)
	}
}
