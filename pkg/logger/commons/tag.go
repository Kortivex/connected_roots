package commons

import "strings"

// Tag is used in WithTag method in commons.
//
//	commons.WithTag(TagPlatformApp)
//	commons.Info("server Up")
type Tag string

// NewTag return a new compatible Tag from `string`.
func NewTag(tag string) Tag {
	return Tag(tag)
}

// Add method return a composition of Tag.
//
//	tag :=NewTag("application")
//	tagApp := tag.Add("myApp") // tagApp is `application/myApp`
//	tagMethod := tagApp.Add("myMethod") // tagMethod is `application/myApp/myMethod`
//
// You can use this method por add a new level in any Tag.
//
//	myTag := TagPlatformHttpserver.Add("createHandler") // myTag is `platform/httpserver/createHandler`
func (t *Tag) Add(tag Tag) Tag {
	return *t + "/" + tag
}

// Split method return Tag split in two
//
// tag :=NewTag("application/MyApp")
// tagRoot, tagApp := tag.Split() // tagRoot is Tag("application") and tagApp is Tag("Myapp")
//
// tag :=NewTag("application/MyApp/Client")
// tagRoot, tagApp := tag.Split() // tagRoot is Tag("application") and tagApp is Tag("Myapp/Client").
func (t *Tag) Split() (Tag, Tag) {
	tags := strings.SplitN(string(*t), "/", 2)

	if len(tags) == 1 {
		return NewTag(tags[0]), ""
	}

	return NewTag(tags[0]), NewTag(tags[1])
}

var (
	TagPlatform           Tag = "platform"
	TagPlatformApp        Tag = "platform/app"
	TagPlatformClient     Tag = "platform/client"
	TagPlatformConfig     Tag = "platform/config"
	TagPlatformEcho       Tag = "platform/echo"
	TagPlatformGorm       Tag = "platform/gorm"
	TagPlatformHttpclient Tag = "platform/httpclient"
	TagPlatformHttpserver Tag = "platform/httpserver"
	TagPlatformMain       Tag = "platform/main"
	TagPlatformShell      Tag = "platform/shell"
	TagPlatformTelemetry  Tag = "platform/telemetry"

	TagService               Tag = "service"
	TagServiceConnectedRoots Tag = "service/connectedroots"

	TagDevelopment      Tag = "development"
	TagDevelopmentDebug Tag = "development/debug"
	TagDevelopmentTrace Tag = "development/trace"
)
