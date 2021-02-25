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

// DeDupeAttachmentList removes duplicate from list
func DeDupeAttachmentList(
	attachmentList []*Attachment,
) []*Attachment {
	type key struct{ parentResourceType, parentResourceID, childResourceType, childResourceID string }
	var uniqueItems []*Attachment
	seenMap := make(map[key]bool)

	for _, v := range attachmentList {
		k := key{
			v.Parent.ResourceType.String(),
			*v.Parent.ResourceID,
			v.Child.ResourceType.String(),
			*v.Child.ResourceID,
		}
		if _, found := seenMap[k]; !found {
			seenMap[k] = true
			uniqueItems = append(uniqueItems, v)
		}
	}
	return uniqueItems
}
