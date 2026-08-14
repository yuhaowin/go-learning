package engine

type SimpleEngine struct {
	ItemChan chan any
}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		result, err := Worker(r)
		if err != nil {
			continue
		}

		for _, item := range result.Items {
			go func() { e.ItemChan <- item }()
		}

		requests = append(requests, result.Requests...)
	}
}
