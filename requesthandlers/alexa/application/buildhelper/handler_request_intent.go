package buildhelper

import (
	"fmt"

	"github.com/personalbuildhelper/protocol/alexa"
)

type handlerIntentRequest struct {
}

func (handlerIntentRequest) handleRequest(r *alexa.AlexaRequest) (*alexa.AlexaResponse, error) {
	return &alexa.AlexaResponse{
		Version: r.Version,
		Response: &alexa.Response{
			OutputSpeech: getOutputSpeech(r.Request.Intent),
		},
	}, nil
}

func getOutputSpeech(intent *alexa.Intent) *alexa.OutputSpeech {
	slots := intent.Slots
	buildType, buildTypeGiven := slots["buildtypeslot"]
	sourceType, sourceTypeGiven := slots["buildsourcetypeslot"]

	switch {
	case buildTypeGiven && sourceTypeGiven:
		return &alexa.OutputSpeech{
			Type: "PlainText",
			Text: fmt.Sprintf("%s is build for %s", sourceType.Value, buildType.Value),
		}
	case !buildTypeGiven && sourceTypeGiven:
		return &alexa.OutputSpeech{
			Type: "PlainText",
			Text: "Please tell the build type",
		}
	case buildTypeGiven && !sourceTypeGiven:
		return &alexa.OutputSpeech{
			Type: "PlainText",
			Text: "Please tell the source type",
		}
	case !buildTypeGiven && !sourceTypeGiven:
		return &alexa.OutputSpeech{
			Type: "PlainText",
			Text: "Please tell source and type for build",
		}
	}

	return &alexa.OutputSpeech{
		Type: "PlainText",
		Text: "This is not possible",
	}

}

/*amzn1.ask.skill.462a6bc5-8525-48cd-9b6a-bc862fc1b667
amzn1.ask.skill.462a6bc5-8525-48cd-9b6a-bc862fc1b667*/
