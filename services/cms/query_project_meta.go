package cms

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

// QueryProjectMeta invokes the cms.QueryProjectMeta API synchronously
// api document: https://help.aliyun.com/api/cms/queryprojectmeta.html
func (client *Client) QueryProjectMeta(request *QueryProjectMetaRequest) (response *QueryProjectMetaResponse, err error) {
	response = CreateQueryProjectMetaResponse()
	err = client.DoAction(request, response)
	return
}

// QueryProjectMetaWithChan invokes the cms.QueryProjectMeta API asynchronously
// api document: https://help.aliyun.com/api/cms/queryprojectmeta.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryProjectMetaWithChan(request *QueryProjectMetaRequest) (<-chan *QueryProjectMetaResponse, <-chan error) {
	responseChan := make(chan *QueryProjectMetaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryProjectMeta(request)
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

// QueryProjectMetaWithCallback invokes the cms.QueryProjectMeta API asynchronously
// api document: https://help.aliyun.com/api/cms/queryprojectmeta.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryProjectMetaWithCallback(request *QueryProjectMetaRequest, callback func(response *QueryProjectMetaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryProjectMetaResponse
		var err error
		defer close(result)
		response, err = client.QueryProjectMeta(request)
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

// QueryProjectMetaRequest is the request struct for api QueryProjectMeta
type QueryProjectMetaRequest struct {
	*requests.RpcRequest
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	Labels     string           `position:"Query" name:"Labels"`
}

// QueryProjectMetaResponse is the response struct for api QueryProjectMeta
type QueryProjectMetaResponse struct {
	*responses.BaseResponse
	RequestId    string                      `json:"RequestId" xml:"RequestId"`
	Success      bool                        `json:"Success" xml:"Success"`
	ErrorCode    string                      `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string                      `json:"ErrorMessage" xml:"ErrorMessage"`
	Resources    ResourcesInQueryProjectMeta `json:"Resources" xml:"Resources"`
}

// CreateQueryProjectMetaRequest creates a request to invoke QueryProjectMeta API
func CreateQueryProjectMetaRequest() (request *QueryProjectMetaRequest) {
	request = &QueryProjectMetaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2018-03-08", "QueryProjectMeta", "cms", "openAPI")
	return
}

// CreateQueryProjectMetaResponse creates a response to parse from QueryProjectMeta response
func CreateQueryProjectMetaResponse() (response *QueryProjectMetaResponse) {
	response = &QueryProjectMetaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
