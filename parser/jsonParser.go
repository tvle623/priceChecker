package main

func GetCostJson(content *interface{}, contentPath []interface{}) float64 {
	jsonObj := content
	cost := -1.0
	found := true
	converted := true
	name := ""
	index := -1

	logDebug("GetCost function\nJSON body is: %v\nJSON path to check: %v", *content, contentPath)

	for _, path := range contentPath {
		index, converted = path.(int)

		if converted {
			jsonObj, found = getJsonObjectByIndex(jsonObj, index)
		} else {
			name, converted = path.(string)
			if converted {
				jsonObj, found = getJsonObjectByName(jsonObj, name)
			}
		}

		if !converted || !found {
			logWarning("Was not able to find the cost from JSON.")
			return -1.0
		}
	}

	if found {
		cost, converted = (*jsonObj).(float64)
	}

	if !converted {
		cost = -1.0
	}

	logInfo("Found the cost: %v", cost)
	return cost
}

func getJsonObjectByName(content *interface{}, name string) (*interface{}, bool) {
	jsonObj, ok := (*content).(map[string]interface{})

	if !ok {
		return nil, ok
	}

	innerObj, ok := jsonObj[name].(interface{})
	return &innerObj, ok
}

func getJsonObjectByIndex(content *interface{}, index int) (*interface{}, bool) {
	jsonObj, ok := (*content).([]interface{})

	if !ok || len(jsonObj) < index+1 {
		return nil, false
	}

	innerObj, ok := jsonObj[index].(interface{})
	return &innerObj, ok
}
