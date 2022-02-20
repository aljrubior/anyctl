package comparators

import "fmt"

var emptyMap map[interface{}]interface{}

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
