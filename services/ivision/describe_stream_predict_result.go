package ivision

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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeStreamPredictResult invokes the ivision.DescribeStreamPredictResult API synchronously
// api document: https://help.aliyun.com/api/ivision/describestreampredictresult.html
func (client *Client) DescribeStreamPredictResult(request *DescribeStreamPredictResultRequest) (response *DescribeStreamPredictResultResponse, err error) {
	response = CreateDescribeStreamPredictResultResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeStreamPredictResultWithChan invokes the ivision.DescribeStreamPredictResult API asynchronously
// api document: https://help.aliyun.com/api/ivision/describestreampredictresult.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeStreamPredictResultWithChan(request *DescribeStreamPredictResultRequest) (<-chan *DescribeStreamPredictResultResponse, <-chan error) {
	responseChan := make(chan *DescribeStreamPredictResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeStreamPredictResult(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeStreamPredictResultWithCallback invokes the ivision.DescribeStreamPredictResult API asynchronously
// api document: https://help.aliyun.com/api/ivision/describestreampredictresult.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeStreamPredictResultWithCallback(request *DescribeStreamPredictResultRequest, callback func(response *DescribeStreamPredictResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeStreamPredictResultResponse
		var err error
		defer close(result)
		response, err = client.DescribeStreamPredictResult(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeStreamPredictResultRequest is the request struct for api DescribeStreamPredictResult
type DescribeStreamPredictResultRequest struct {
	*requests.RpcRequest
	NextPageToken        string           `position:"Query" name:"NextPageToken"`
	StartTime            string           `position:"Query" name:"StartTime"`
	PredictId            string           `position:"Query" name:"PredictId"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	ProbabilityThreshold string           `position:"Query" name:"ProbabilityThreshold"`
	ShowLog              string           `position:"Query" name:"ShowLog"`
	ModelId              string           `position:"Query" name:"ModelId"`
	EndTime              string           `position:"Query" name:"EndTime"`
	CurrentPage          requests.Integer `position:"Query" name:"CurrentPage"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeStreamPredictResultResponse is the response struct for api DescribeStreamPredictResult
type DescribeStreamPredictResultResponse struct {
	*responses.BaseResponse
	RequestId          string             `json:"RequestId" xml:"RequestId"`
	TotalNum           int                `json:"TotalNum" xml:"TotalNum"`
	CurrentPage        int                `json:"CurrentPage" xml:"CurrentPage"`
	PageSize           int                `json:"PageSize" xml:"PageSize"`
	NextPageToken      string             `json:"NextPageToken" xml:"NextPageToken"`
	StreamPredictDatas StreamPredictDatas `json:"StreamPredictDatas" xml:"StreamPredictDatas"`
}

// CreateDescribeStreamPredictResultRequest creates a request to invoke DescribeStreamPredictResult API
func CreateDescribeStreamPredictResultRequest() (request *DescribeStreamPredictResultRequest) {
	request = &DescribeStreamPredictResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ivision", "2019-03-08", "DescribeStreamPredictResult", "ivision", "openAPI")
	return
}

// CreateDescribeStreamPredictResultResponse creates a response to parse from DescribeStreamPredictResult response
func CreateDescribeStreamPredictResultResponse() (response *DescribeStreamPredictResultResponse) {
	response = &DescribeStreamPredictResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
