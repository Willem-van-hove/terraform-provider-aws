//go:generate go run ../../generate/listpages/main.go -ListOps=ListApplications
//go:generate go run ../../generate/tags/main.go -ListTags -ListTagsInIDElem=ResourceARN -ServiceTagsSlice -TagInIDElem=ResourceARN -UpdateTags
// ONLY generate directives and package declaration! Do not add anything else to this file.

package kinesisanalyticsv2
