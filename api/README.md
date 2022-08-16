
```go
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := new(TitleLookupResponse)
	if err := xml.Unmarshal(b, response); err != nil {
		return nil, err
	}
	if response.Result.Code == 0 {
		xml.Unmarshal(b, &response.Result)
	}
	return response, nil

```