package examples

type (
	// @GetFeature
	// Rectangle @request
	Rectangle struct {
		lo Point `json:"lo" binding:"required"` // One corner of the rectangle.
		hi Point `json:"hi" binding:"required"` // The other corner of the rectangle.
	}
	Point struct {
		latitude  int32 `json:"latitude" binding:"required"`
		longitude int32 `json:"longitude" binding:"required"`
	}

	// Feature @response
	Feature struct {
		name     string `json:"name"`     // The name of the feature.
		location Point  `json:"location"` // The point where the feature is detected.
	}

	// @GetRecordRoute
	// RouteNote @request
	RouteNote struct {
		location Point  `json:"location" binding:"required"` // The location from which the message is sent.
		message  string `json:"message"`                     // The message to be sent.
	}

	// RouteNoteDetails @response
	RouteNoteDetails struct {
		routes []routeSummary `json:"routes"`
	}

	routeSummary struct {
		point_count   int32 `json:"point_count"`   // The number of points received.
		feature_count int32 `json:"feature_count"` // The number of known features passed while traversing the route.
		distance      int32 `json:"distance"`      // The distance covered in metres.
		elapsed_time  int32 `json:"elapsed_time"`  // The duration of the traversal in seconds.
	}
)

// 对外接口
// GetFeature
func GetFeature() {
	// TODO
}

// GetRecordRoute
func GetRecordRoute() {
	// TODO
}
