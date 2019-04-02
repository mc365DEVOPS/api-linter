package lint

// Response describes the result returned by a rule.
type Response struct {
	Problems []Problem
}

// mrge merges another response.
func (resp *Response) merge(other Response) {
	resp.Problems = append(resp.Problems, other.Problems...)
}
