package emr

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// OpsOperationListItem is a nested struct in emr response
type OpsOperationListItem struct {
	Id              int    `json:"Id" xml:"Id"`
	StartTime       string `json:"StartTime" xml:"StartTime"`
	EndTime         string `json:"EndTime" xml:"EndTime"`
	OpsCommandId    int    `json:"OpsCommandId" xml:"OpsCommandId"`
	OpsCommandName  string `json:"OpsCommandName" xml:"OpsCommandName"`
	Status          string `json:"Status" xml:"Status"`
	TotalTaskNum    int    `json:"TotalTaskNum" xml:"TotalTaskNum"`
	FailedTaskNum   int    `json:"FailedTaskNum" xml:"FailedTaskNum"`
	FinishedTaskNum int    `json:"FinishedTaskNum" xml:"FinishedTaskNum"`
	ClusterId       string `json:"ClusterId" xml:"ClusterId"`
	RegionId        string `json:"RegionId" xml:"RegionId"`
	Params          string `json:"Params" xml:"Params"`
	Remark          string `json:"Remark" xml:"Remark"`
	RunningTime     int    `json:"RunningTime" xml:"RunningTime"`
	Category        string `json:"Category" xml:"Category"`
}
