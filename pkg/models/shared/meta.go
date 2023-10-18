// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// MetaCursors - Cursors to navigate to previous or next pages through the API
type MetaCursors struct {
	// Cursor to navigate to the current page of results through the API
	Current *string `json:"current,omitempty"`
	// Cursor to navigate to the next page of results through the API
	Next *string `json:"next,omitempty"`
	// Cursor to navigate to the previous page of results through the API
	Previous *string `json:"previous,omitempty"`
}

func (o *MetaCursors) GetCurrent() *string {
	if o == nil {
		return nil
	}
	return o.Current
}

func (o *MetaCursors) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *MetaCursors) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

// Meta - Response metadata
type Meta struct {
	// Cursors to navigate to previous or next pages through the API
	Cursors *MetaCursors `json:"cursors,omitempty"`
	// Number of items returned in the data property of the response
	ItemsOnPage *int64 `json:"items_on_page,omitempty"`
}

func (o *Meta) GetCursors() *MetaCursors {
	if o == nil {
		return nil
	}
	return o.Cursors
}

func (o *Meta) GetItemsOnPage() *int64 {
	if o == nil {
		return nil
	}
	return o.ItemsOnPage
}