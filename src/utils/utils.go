package utils

import (
	"encoding/json"
	"fmt"
	"github.com/aljrubior/anyctl/clients/deployments/response"
	"github.com/aljrubior/anyctl/managers/entities"
	"github.com/google/uuid"
	"os"
	"strconv"
)

func WriteFile(filePath string, data []byte) error {

	err := os.WriteFile(filePath, data, 0644)
	return err
}

func TargetEntities2Map(fromTargets []entities.TargetEntity) map[string]entities.TargetEntity {

	result := make(map[string]entities.TargetEntity)

	for _, v := range fromTargets {
		result[v.GetId()] = v
	}

	return result
}

func mapPrivateSpaces(privateSpaces *[]entities.PrivateSpaceEntity) map[string]entities.PrivateSpaceEntity {

	result := make(map[string]entities.PrivateSpaceEntity)

	for _, v := range *privateSpaces {
		result[v.Id] = v
	}

	return result
}

func RtfTargetsEntities2Map(fromTargets []entities.RtfTargetEntity) map[string]entities.RtfTargetEntity {

	result := make(map[string]entities.RtfTargetEntity)

	for _, v := range fromTargets {
		result[v.Id] = v
	}

	return result
}

func BuildReplicasSummary(fromReplicas []response.Replica) string {
	total := len(fromReplicas)
	started := 0

	for _, v := range fromReplicas {
		if v.State == "STARTED" {
			started++
		}
	}

	return fmt.Sprintf("%s/%s", strconv.Itoa(started), strconv.Itoa(total))
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func PrintAsJson(object interface{}) {
	data, err := json.Marshal(object)

	if err != nil {
		println(err.Error())
	}

	println(string(data))

}
