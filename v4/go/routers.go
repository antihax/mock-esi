package esiv4

import (
	"net/http"
	"reflect"
	"strconv"

	"github.com/antihax/mock-esi/mockesi"
)

func init() {

	mockesi.NewRoute(
		"GetCharactersCharacterId",
		"Get",
		"/v4/characters/{character_id}/",
		GetCharactersCharacterId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdNotifications",
		"Get",
		"/v4/characters/{character_id}/notifications/",
		GetCharactersCharacterIdNotifications,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdCspa",
		"Post",
		"/v4/characters/{character_id}/cspa/",
		PostCharactersCharacterIdCspa,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationId",
		"Get",
		"/v4/corporations/{corporation_id}/",
		GetCorporationsCorporationId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdSkills",
		"Get",
		"/v4/characters/{character_id}/skills/",
		GetCharactersCharacterIdSkills,
	)

	mockesi.NewRoute(
		"GetUniverseSystemsSystemId",
		"Get",
		"/v4/universe/systems/{system_id}/",
		GetUniverseSystemsSystemId,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdWalletsDivisionJournal",
		"Get",
		"/v4/corporations/{corporation_id}/wallets/{division}/journal/",
		GetCorporationsCorporationIdWalletsDivisionJournal,
	)

}

func errorOut(w http.ResponseWriter, r *http.Request, e error) {
	w.WriteHeader(http.StatusInternalServerError)

	w.Write([]byte(e.Error()))
}

func processParameters(data interface{}, input string) (v interface{}, err error) {
	switch reflect.TypeOf(data).String() {
	case "int":
		{
			v, err = strconv.Atoi(input)
			v = v.(int)
		}
	case "int8":
		{
			v, err = strconv.ParseInt(input, 10, 8)
			v = (int8)(v.(int64))
		}
	case "int16":
		{
			v, err = strconv.ParseInt(input, 10, 16)
			v = (int16)(v.(int64))
		}
	case "rune":
		{
			v, err = strconv.ParseInt(input, 10, 32)
			v = (rune)(v.(int64))
		}
	case "int32":
		{
			v, err = strconv.ParseInt(input, 10, 32)
			v = (int32)(v.(int64))
		}
	case "int64":
		{
			v, err = strconv.ParseInt(input, 10, 64)
			v = v.(int64)
		}
	case "uint8":
		{
			v, err = strconv.ParseUint(input, 10, 8)
			v = (uint8)(v.(uint64))
		}
	case "byte":
		{
			v, err = strconv.ParseUint(input, 10, 8)
			v = (byte)(v.(uint64))
		}
	case "uint16":
		{
			v, err = strconv.ParseUint(input, 10, 16)
			v = (uint16)(v.(uint64))
		}
	case "uint32":
		{
			v, err = strconv.ParseUint(input, 10, 32)
			v = (uint32)(v.(uint64))
		}
	case "uint64":
		{
			v, err = strconv.ParseUint(input, 10, 64)
			v = v.(uint64)
		}
	case "float32":
		{
			v, err = strconv.ParseFloat(input, 32)
			v = (float32)(v.(float64))
		}
	case "float64":
		{
			v, err = strconv.ParseFloat(input, 64)
			v = v.(float64)
		}
	case "bool":
		{
			v, err = strconv.ParseBool(input)
			v = v.(bool)
		}
	case "string":
		{
			v = input
		}
	}

	return
}
