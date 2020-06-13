package config

import "github.com/pkg/errors"

// WorkflowDefinition is the consumer friendly data structure that hosts the loaded workflow definition
type WorkflowDefinition struct {
	Flowit Flowit
}

// Flowit is the consumer friendly data structure that hosts the loaded workflow definition main body
type Flowit struct {
	Version       string
	Config        Config
	Variables     Variables
	Branches      []Branch
	Tags          []Tag
	StateMachines []StateMachine
	Workflows     []Workflow
}

// Config is the consumer friendly data structure that hosts the loaded workflow definition configuration
type Config struct {
	AbortOnFailedAction bool
	Strict              bool
	Shell               string
}

// Variables is the consumer friendly data structure that hosts the loaded workflow definition variables
type Variables map[string]interface{}

// Branch is the consumer friendly data structure that hosts the loaded workflow definition branch
type Branch struct {
	ID          string
	Name        string
	Prefix      string
	Suffix      string
	Eternal     bool
	Protected   bool
	Transitions []Transition
}

// Tag is the consumer friendly data structure that hosts the loaded workflow definition tag
type Tag struct {
	ID       string
	Format   string
	Stages   Stages
	Branches []string
}

// StateMachine is the consumer friendly data structure that hosts
// the loaded workflow definition state machine
type StateMachine struct {
	ID           string
	Stages       []string
	InitialStage string
	FinalStages  []string
	Transitions  []StateMachineTransition
}

// StateMachineTransition is the consumer friendly data structure that hosts
// the loaded workflow definition state machine transition
type StateMachineTransition struct {
	From []string
	To   []string
}

// Stages is the consumer friendly data structure that hosts
// the loaded workflow definition tag stages
type Stages map[string][]string

// Workflow is the consumer friendly data structure that hosts
// the loaded workflow definition workflow
type Workflow struct {
	ID           string
	StateMachine string
	Stages       []Stage
}

// Stage is the consumer friendly data structure that hosts
// the loaded workflow definition workflow stage
type Stage struct {
	ID         string
	Args       []string
	Conditions []string
	Actions    []string
}

// Transition is the consumer friendly data structure that hosts
// the loaded workflow definition branch transition
type Transition struct {
	From string
	To   []string
}

// Stages returns the loaded workflow definition stages for the specified workflowID
func (wd WorkflowDefinition) Stages(workflowID string) ([]Stage, error) {
	for _, workflow := range wd.Flowit.Workflows {
		if workflow.ID == workflowID {
			return workflow.Stages, nil
		}
	}
	return nil, errors.New("Invalid workflowID: " + workflowID)
}

// Stage returns the loaded workflow definition stage for the specified workflowID and stageID
func (wd WorkflowDefinition) Stage(workflowID, stageID string) (Stage, error) {
	for _, workflow := range wd.Flowit.Workflows {
		if workflow.ID == workflowID {
			for _, stage := range workflow.Stages {
				if stage.ID == stageID {
					return stage, nil
				}
			}
			return Stage{}, errors.New("Invalid stageID: " + stageID)
		}
	}
	return Stage{}, errors.New("Invalid workflowID: " + workflowID)
}
