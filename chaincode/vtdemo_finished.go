/*
Copyright IBM Corp 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"errors"
	"fmt"	
	

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

//custom data models
type CompanyInfo struct {
	Companyname string `json:"companyname"`
	Companycontact  string `json:"companycontact"`
	Companybudget  int `json:"companybudget"`
	CompanyID string `json:"companyid"`
	ContractorIDs string `json:"contractorids"`
	ContraIDs string `json:"contractorids"`
	taskids string `json:"taskids"`

}

type ContractorInfo struct {
	Contractorname string `json:"Contractorname"`	
	Contractorassignedto string `json:"contractorassignedto"`		// assigned to which project
	ContractorHourlyrate  string `json:"contractorHourlyrate"`
	ContractorID string `json:"contractorid"`
	CompanyID string `json:"companyid"`
	hoursworked int `json:"hoursworked"`
}


type ManagerInfo struct {
	Managername string `json:"Contractorname"`	
	Managerassignedto string `json:"managerassignedto"`		// assigned to which project
	ManagerID string `json:"managerid"`
	CompanyID string `json:"companyid"`
	EstimatedHours int `json:"estimatedhours"`
}

type TimecardInfo struct {
	ContractorID string `json:"contractorid"`	
	TimecardweekEnding string `json:"timecardweekending"`		// assigned to which project
	TimecardHours int `json:"timecardhours"`
	TimecardDescription string `json:"timecarddescription"`
	TimecardID string `json:"timecardid"`
}

type TaskInfo struct {
	TaskID string `json:"taskid"`	
	TaskName string `json:"taskname"`
	TaskDescription int `json:"taskdescription"`
	EstimatedHours string `json:"estimatedhours"`
	EstimatedBudget int `json:"estimatedbudget"`
	TaskCreatedBy string `json:"taskcreatedBy"`
	CompanyID string `json:"comppanyid"`
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	err := stub.PutState("hello_world", []byte(args[0]))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Invoke isur entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {
		return t.Init(stub, "init", args)
	} else if function == "write" {
		return t.write (stub, args)
	} else if function == "createcompany" {
		return t.createcompany(stub, args)
	} else if function == "createcontractor" {
		return t.createcontractor(stub, args)
	} else if function == "creattask" {
		return t.creattask(stub, args)
	} else if function == "write" {
		return t.write(stub, args)
	}

	fmt.Println("invoke did not find func: " + function)

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "read" { //read a variable
		return t.read(stub, args)
	}
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query: " + function)
}


// insert company info
func (t *SimpleChaincode) createcompany(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	//var key, value string
	var err error
	fmt.Println("running createcompany()")

	if len(args) != 6 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3. name of the key and value to set")
	}

	companyid := args[0]
	companyname := args[1]
	companycontact := args[2]
	companybudget := args[3]
	contractorids := args[4]
	taskids := args[5]
	
	str := `{"companyname": "` + companyname + `", "companycontact": "` + companycontact + `", "companybudget": "` + companybudget + `", "companyid": "` + companyid + `","contractorids": "` + contractorids + `","taskids": "` + taskids + `"}`

	//fmt.Println ("company parms" + companyid + "::" + companyname + "::" + companycontact + "::"+ companybudget + "::" + contractorids +"::"+ taskids + "::"+ str)
	
	fmt.Println(str)
	err = stub.PutState(companyid, []byte(str))									//store company with id as key

	if err != nil {
		return nil, err
	}
	return nil, nil
}

// insert contractor info
func (t *SimpleChaincode) createcontractor(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	//var key, value string
	var err error
	fmt.Println("running createcontractor()")

	if len(args) != 5 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3. name of the key and value to set")
	}

	contractorid := args[0]
	contractorname := args[1]
	contractorassignedto := args[2]	// assigned to task
	contractorhourlyrate := args[3]
	hoursworked := args[4]
	
	str := `{"contractorname": "` + contractorname + `", "contractorassignedto": "` + contractorassignedto + `", "contractorid": "` + contractorid + `","contractorhourlyrate": "` + contractorhourlyrate + `","hoursworked": "` + hoursworked + `"}`
	//fmt.Println ("contractor parms" + companyid + "::" + contractorname + "::" + contractorassignedto + "::"+ contractorid + "::" + contractorhourlyrate + "::"+ str)
	fmt.Println(str)

	err = stub.PutState(contractorid, []byte(str))									//store contractor with id as key

	if err != nil {
		return nil, err
	}
	return nil, nil
}



// insert task info
func (t *SimpleChaincode) creattask(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	//var key, value string
	var err error
	fmt.Println("running createtask()")

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3. name of the key and value to set")
	}

	taskid := args[0]
	taskname := args[1]
	managername := args[2]
	estimatedhours := args[3]
	
	str := `{"taskid": "` + taskid + `", "taskname": "` + taskname + `", "estimatedhours": "` + estimatedhours + `", "managername": "` + managername + `"}`

	fmt.Println(str)

	//fmt.Println ("task parms" + taskid + "::" + taskname + "::" + taskdescription + "::"+ estimatedhours + "::" + estimatedbudget + "::" + taskcreatedBy + "::"+ companyid +"::"+ str)
	
	
	err = stub.PutState(taskid, []byte(str))									//store manager with id as key

	if err != nil {
		return nil, err
	}
	return nil, nil
}

// insert manager info
func (t *SimpleChaincode) createmanager(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	//var key, value string
	var err error
	fmt.Println("running createmanager()")

	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3. name of the key and value to set")
	}

	companyid := args[0]
	managername := args[1]
	managerID := args[2]
	managerassignedto := args[2]
	
	str := `{"managername": "` + managername + `", "managerID": "` + managerID + `", "managerassignedto": ` + managerassignedto + `, "companyid": "` + companyid + `"}`

	fmt.Println ("manager parms" + companyid + "::" + managername + "::" + managerID + "::"+ managerassignedto + "::" + str)
	
	
	err = stub.PutState(managerID, []byte(str))									//store manager with id as key

	if err != nil {
		return nil, err
	}
	return nil, nil
}



// write - invoke function to write key/value pair
func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, value string
	var err error
	fmt.Println("running write VT..3")

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
	}

	key = args[0] //rename for funsies
	value = args[1]
	err = stub.PutState(key, []byte(value)) //write the variable into the chaincode state
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// read - query function to read key/value pair
func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, jsonResp string
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

	key = args[0]
	valAsbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}

	return valAsbytes, nil
}
