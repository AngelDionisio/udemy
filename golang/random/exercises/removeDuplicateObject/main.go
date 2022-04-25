package main

import "fmt"

// ResourceType something that can be shared
type ResourceType int32

func (e ResourceType) String() string { return ResourceTypeByValue[e] }

// map
var (
	ResourceTypeByValue = map[ResourceType]string{
		ResourceTypeAppHouse: "HOUSE",
		ResourceTypeBuyer:    "BUYER",
	}
)

// types
const (
	ResourceTypeAppHouse ResourceType = 0
	ResourceTypeBuyer    ResourceType = 1
)

// ResourceRef instace of a resource
type ResourceRef struct {
	ResourceType *ResourceType
	ResourceID   *string
}

// Attachment link between parent and child resources
type Attachment struct {
	Parent *ResourceRef
	Child  *ResourceRef
}

func pointerToString(s string) *string {
	return &s
}

func main() {
	parentType := ResourceTypeBuyer
	parent := ResourceRef{
		ResourceType: &parentType,
		ResourceID:   pointerToString("parent_testID_1"),
	}

	childType := ResourceTypeAppHouse
	child := ResourceRef{
		ResourceType: &childType,
		ResourceID:   pointerToString("child_testID_1"),
	}

	childDuped := ResourceRef{
		ResourceType: &childType,
		ResourceID:   pointerToString("child_testID_1"),
	}

	attachment := &Attachment{
		Parent: &parent,
		Child:  &child,
	}
	attachmentDuped := &Attachment{
		Parent: &parent,
		Child:  &childDuped,
	}

	var listWithDupes []*Attachment
	listWithDupes = append(listWithDupes, attachment)
	listWithDupes = append(listWithDupes, attachmentDuped)

	deDuped := DeDupeAttachmentList(listWithDupes)

	fmt.Println("Length of original:", len(listWithDupes))
	fmt.Println("Length of deDuped :", len(deDuped))
}

type attachmentKey struct{ parentResourceType, parentResourceID, childResourceType, childResourceID string }

func makeAttachmentKey(attachment *Attachment) *attachmentKey {
	return &attachmentKey{
		parentResourceType: attachment.Parent.ResourceType.String(),
		parentResourceID:   *attachment.Parent.ResourceID,
		childResourceType:  attachment.Child.ResourceType.String(),
		childResourceID:    *attachment.Child.ResourceID,
	}
}

// DeDupeAttachmentList removes duplicate from list
func DeDupeAttachmentList(attachmentList []*Attachment) []*Attachment {
	var uniqueList []*Attachment
	seenMap := make(map[attachmentKey]bool)

	for _, attachment := range attachmentList {
		k := *makeAttachmentKey(attachment)
		beforeLen := len(seenMap)
		seenMap[k] = true
		if len(seenMap) > beforeLen {
			uniqueList = append(uniqueList, attachment)
		}
	}
	return uniqueList
}
