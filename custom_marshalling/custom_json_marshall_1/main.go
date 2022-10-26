package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type ValidationType string

var (
	QuestionAnswerValidationType ValidationType = "question_answer"
	SSHCommandValidationType     ValidationType = "ssh_command"
)

type Scenario struct {
	Name  string `json:"name"`
	Steps []Step `json:"steps,omitempty"`
}

type Step struct {
	Name            string        `json:"name"`
	ValidationSteps []interface{} `json:"validation_steps,omitempty"`
}

type QuestionAnswerValidationStep struct {
	Question string `json:"question,omitempty"`
	Answer   string `json:"answer,omitempty"`
}

func (s QuestionAnswerValidationStep) Validate(response interface{}) {

}

type SSHValidationStep struct {
	Command        string `json:"command,omitempty"`
	ExpectedResult string `json:"expected_result,omitempty"`
}

func (s SSHValidationStep) Validate(response interface{}) {

}

type ValidationStep interface {
	Validate(response interface{})
}

func main() {
	file, err := os.Open("payload.json")
	if err != nil {
		log.Println(err)
	}

	var s Scenario
	if err := json.NewDecoder(file).Decode(&s); err != nil {
		log.Println(err)
	}
	for _, step := range s.Steps {
		var newVSteps []ValidationStep
		for _, vstep := range step.ValidationSteps {

			stepType := vstep.(map[string]interface{})["type"]

			switch stepType {
			case "question_answer":
				fmt.Println("QuestionAnswer")
				var qa QuestionAnswerValidationStep

				qa.Question = vstep.(map[string]interface{})["question"].(string)
				qa.Answer = vstep.(map[string]interface{})["answer"].(string)
				newVSteps = append(newVSteps, qa)

			case "ssh_command":
				var sshVs SSHValidationStep

				sshVs.Command = vstep.(map[string]interface{})["command"].(string)
				sshVs.ExpectedResult = vstep.(map[string]interface{})["expectedResult"].(string)

				newVSteps = append(newVSteps, sshVs)
			}

		}

		fmt.Println(newVSteps)
	}

	//fmt.Println(s)
}
