package comparators

import (
	"encoding/json"
	"fmt"
	"github.com/aljrubior/anyctl/clients/deployments/response"
	"github.com/aljrubior/anyctl/comparators/model"
	"gopkg.in/yaml.v2"
)

var emptyMap map[interface{}]interface{}

func NewDeploymentResponseComparator(leftHand, rightHand response.DeploymentResponse) (*DeploymentComparator, error) {

	leftHandAsMap, err := transformDeploymentResponseToMap(leftHand)

	if err != nil {
		return nil, err
	}

	rightHandAsMap, err := transformDeploymentResponseToMap(rightHand)

	if err != nil {
		return nil, err
	}

	_, exists := rightHandAsMap["desiredVersion"]

	if exists {
		delete(leftHandAsMap, "desiredVersion")
	}

	return NewDeploymentComparator(leftHandAsMap, rightHandAsMap), nil
}

func NewDeploymentSpecResponseComparator(leftHand, rightHand response.DeploymentSpecResponse) (*DeploymentComparator, error) {

	leftHandAsMap, err := transformDeploymentSpecResponseToMap(leftHand)

	if err != nil {
		return nil, err
	}

	rightHandAsMap, err := transformDeploymentSpecResponseToMap(rightHand)

	if err != nil {
		return nil, err
	}

	return NewDeploymentComparator(leftHandAsMap, rightHandAsMap), nil
}

func NewDeploymentComparator(leftHand, rightHand map[interface{}]interface{}) *DeploymentComparator {
	return &DeploymentComparator{
		leftHand:  leftHand,
		rightHand: rightHand,
	}
}

type DeploymentComparator struct {
	differences         []DeploymentDifference
	leftHand, rightHand map[interface{}]interface{}
}

func (this *DeploymentComparator) Compare() []DeploymentDifference {
	return this.diff(this.leftHand, this.rightHand, 0)
}

func (this *DeploymentComparator) addKeyChanged(differences []DeploymentDifference, depth int, name, leftValue, rightValue interface{}) []DeploymentDifference {

	switch leftValue.(type) {
	case map[interface{}]interface{}:
		differences = append(differences, DeploymentDifference{"~", depth, name, nil, nil})
	default:
		differences = append(differences, DeploymentDifference{"~", depth, name, leftValue, rightValue})
	}

	return differences
}

func (this *DeploymentComparator) addNewKey(differences []DeploymentDifference, depth int, name interface{}, value interface{}) []DeploymentDifference {

	switch value.(type) {
	case map[interface{}]interface{}:
		differences = append(differences, DeploymentDifference{"+", depth, name, nil, nil})
	default:
		differences = append(differences, DeploymentDifference{"+", depth, name, nil, value})
	}

	return differences
}

func (this *DeploymentComparator) addKeyDeleted(differences []DeploymentDifference, depth int, name interface{}, value interface{}) []DeploymentDifference {

	switch value.(type) {
	case map[interface{}]interface{}:
		differences = append(differences, DeploymentDifference{"-", depth, name, nil, nil})
	default:
		differences = append(differences, DeploymentDifference{"-", depth, name, nil, value})
	}

	return differences
}

func (this *DeploymentComparator) diffNewKeysOnly(leftHand, rightHand map[interface{}]interface{}, depth int) []DeploymentDifference {

	var diff []DeploymentDifference

	for k, rightValue := range rightHand {
		leftValue, exists := leftHand[k]

		switch rightValue.(type) {
		case map[interface{}]interface{}:
			if !exists {
				diff = this.addNewKey(diff, depth, k, nil)
			}

			if leftValue == nil {
				leftValue = emptyMap
			}

			this.diffNewKeysOnly(leftValue.(map[interface{}]interface{}), rightValue.(map[interface{}]interface{}), depth+1)
		default:
			if !exists {
				diff = this.addNewKey(diff, depth, k, rightValue)
			}
		}
	}

	return diff
}

func (this *DeploymentComparator) diff(leftHand, rightHand map[interface{}]interface{}, depth int) []DeploymentDifference {

	var diff []DeploymentDifference

	if len(leftHand) == 0 {
		return this.diffNewKeysOnly(emptyMap, rightHand, depth)
	}

	for k, leftValue := range leftHand {
		rightValue, exists := rightHand[k]

		switch leftValue.(type) {
		case map[interface{}]interface{}:

			if exists {
				subdiff := this.diff(leftValue.(map[interface{}]interface{}), rightValue.(map[interface{}]interface{}), depth+1)

				if len(subdiff) > 0 {
					diff = this.addKeyChanged(diff, depth, k, nil, nil)
					diff = append(diff, subdiff...)
				}

				subdiff = this.diffNewKeysOnly(leftValue.(map[interface{}]interface{}), rightValue.(map[interface{}]interface{}), depth+1)

				if len(subdiff) > 0 {
					diff = append(diff, subdiff...)
				}

				continue
			}

			diff = this.addKeyDeleted(diff, depth, k, nil)

			subdiff := this.diff(leftValue.(map[interface{}]interface{}), emptyMap, depth+1)

			if len(subdiff) > 0 {
				diff = append(diff, subdiff...)
			}

		default:
			if !exists {
				diff = this.addKeyDeleted(diff, depth, k, leftValue)
				continue
			}

			leftValueAsString := fmt.Sprintf("%v", leftValue)
			rightValueAsString := fmt.Sprintf("%v", rightValue)

			if leftValueAsString != rightValueAsString {
				diff = this.addKeyChanged(diff, depth, k, leftValueAsString, rightValueAsString)
				continue
			}
		}
	}

	return diff
}

func transformDeploymentResponseToMap(response response.DeploymentResponse) (map[interface{}]interface{}, error) {
	data, err := json.Marshal(response)

	if err != nil {
		return nil, err
	}

	var spec model.DeploymentSpec

	err = json.Unmarshal(data, &spec)

	if err != nil {
		return nil, err
	}

	data, err = json.Marshal(spec)

	if err != nil {
		return nil, err
	}

	var result map[interface{}]interface{}

	err = yaml.Unmarshal(data, &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func transformDeploymentSpecResponseToMap(response response.DeploymentSpecResponse) (map[interface{}]interface{}, error) {
	data, err := json.Marshal(response)

	if err != nil {
		return nil, err
	}

	var result map[interface{}]interface{}

	err = yaml.Unmarshal(data, &result)

	if err != nil {
		return nil, err
	}

	return result, nil

}
