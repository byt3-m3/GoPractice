package main

import (
	"bytes"
	"encoding/json"
	"errors"
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
	Name               string            `json:"name"`
	ValidationSteps    []ValidationStep  `json:"-"`
	ValidationStepsRaw []json.RawMessage `json:"validation_steps,omitempty"`
}

func (s *Step) UnmarshalJSON(b []byte) error {
	type step Step
	err := json.Unmarshal(b, (*step)(s))
	if err != nil {
		return err
	}

	for _, rawData := range s.ValidationStepsRaw {
		var data map[string]interface{}
		if err := json.Unmarshal(rawData, &data); err != nil {
			log.Println(err)
			return errors.New("Something Happened")
		}
		stepType := data["type"]
		var newStruct ValidationStep
		switch stepType {
		case "question_answer":
			newStruct = &QuestionAnswerValidationStep{}

		case "ssh_command":
			newStruct = &SSHValidationStep{}

		default:
			return errors.New("Unknown Type")
		}
		if err := json.Unmarshal(rawData, newStruct); err != nil {
			log.Println(err)
			return errors.New("Something Happened")
		}
		s.ValidationSteps = append(s.ValidationSteps, newStruct)

	}
	return nil

}

func (s *Step) MarshalJSON() ([]byte, error) {
	type step Step
	if s.ValidationSteps != nil {
		s.ValidationStepsRaw = []json.RawMessage{}
		for _, vStep := range s.ValidationSteps {
			var ntype string
			switch vStep.(type) {
			case *QuestionAnswerValidationStep:
				ntype = "question_answer"
			case *SSHValidationStep:
				ntype = "ssh_command"
			default:
				return nil, errors.New("Bad type")
			}
			vStepBytes, err := json.Marshal(vStep)
			if err != nil {
				fmt.Println("Somthing Happened")
				return nil, err
			}
			var nVstepRaw map[string]interface{}
			if err = json.Unmarshal(vStepBytes, &nVstepRaw); err != nil {

			}
			nVstepRaw["type"] = ntype
			vStepBytes, err = json.Marshal(nVstepRaw)
			if err != nil {
				fmt.Println("Somthing Happened")
				return nil, err
			}
			s.ValidationStepsRaw = append(s.ValidationStepsRaw, vStepBytes)
		}
	} else {
		s.ValidationStepsRaw = nil
	}
	return json.Marshal((*step)(s))
}

type QuestionAnswerValidationStep struct {
	Question string `json:"question,omitempty"`
	Answer   string `json:"answer,omitempty"`
}

func (s *QuestionAnswerValidationStep) Validate(response interface{}) {

}

type SSHValidationStep struct {
	Command        string `json:"command,omitempty"`
	ExpectedResult string `json:"expected_result,omitempty"`
}

func (s *SSHValidationStep) Validate(response interface{}) {

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

	myBytes, err := json.Marshal(s)
	var prettyprint bytes.Buffer

	err = json.Indent(&prettyprint, myBytes, "", "  ")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(prettyprint.Bytes()))
	//for _, step := range s.Steps {
	//	var newVSteps []ValidationStep
	//	for _, vstep := range step.ValidationSteps {
	//
	//		stepType := vstep.(map[string]interface{})["type"]
	//
	//		switch stepType {
	//		case "question_answer":
	//			fmt.Println("QuestionAnswer")
	//			var qa QuestionAnswerValidationStep
	//
	//			qa.Question = vstep.(map[string]interface{})["question"].(string)
	//			qa.Answer = vstep.(map[string]interface{})["answer"].(string)
	//			newVSteps = append(newVSteps, qa)
	//
	//		case "ssh_command":
	//			var sshVs SSHValidationStep
	//
	//			sshVs.Command = vstep.(map[string]interface{})["command"].(string)
	//			sshVs.ExpectedResult = vstep.(map[string]interface{})["expectedResult"].(string)
	//
	//			newVSteps = append(newVSteps, sshVs)
	//		}
	//
	//	}
	//
	//	fmt.Println(newVSteps)
	//}

	//fmt.Println(s)
}
