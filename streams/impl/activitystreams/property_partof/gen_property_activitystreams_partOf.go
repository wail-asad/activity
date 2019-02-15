package propertypartof

import (
	"fmt"
	vocab "github.com/go-fed/activity/streams/vocab"
	"net/url"
)

// ActivityStreamsPartOfProperty is the functional property "partOf". It is
// permitted to be one of multiple value types. At most, one type of value can
// be present, or none at all. Setting a value will clear the other types of
// values so that only one of the 'Is' methods will return true. It is
// possible to clear all values, so that this property is empty.
type ActivityStreamsPartOfProperty struct {
	activitystreamsLinkMember                  vocab.ActivityStreamsLink
	activitystreamsCollectionMember            vocab.ActivityStreamsCollection
	activitystreamsCollectionPageMember        vocab.ActivityStreamsCollectionPage
	activitystreamsMentionMember               vocab.ActivityStreamsMention
	activitystreamsOrderedCollectionMember     vocab.ActivityStreamsOrderedCollection
	activitystreamsOrderedCollectionPageMember vocab.ActivityStreamsOrderedCollectionPage
	unknown                                    interface{}
	iri                                        *url.URL
	alias                                      string
}

// DeserializePartOfProperty creates a "partOf" property from an interface
// representation that has been unmarshalled from a text or binary format.
func DeserializePartOfProperty(m map[string]interface{}, aliasMap map[string]string) (*ActivityStreamsPartOfProperty, error) {
	alias := ""
	if a, ok := aliasMap["https://www.w3.org/TR/activitystreams-vocabulary"]; ok {
		alias = a
	}
	propName := "partOf"
	if len(alias) > 0 {
		// Use alias both to find the property, and set within the property.
		propName = fmt.Sprintf("%s:%s", alias, "partOf")
	}
	i, ok := m[propName]

	if ok {
		if s, ok := i.(string); ok {
			u, err := url.Parse(s)
			// If error exists, don't error out -- skip this and treat as unknown string ([]byte) at worst
			// Also, if no scheme exists, don't treat it as a URL -- net/url is greedy
			if err == nil && len(u.Scheme) > 0 {
				this := &ActivityStreamsPartOfProperty{
					alias: alias,
					iri:   u,
				}
				return this, nil
			}
		}
		if m, ok := i.(map[string]interface{}); ok {
			if v, err := mgr.DeserializeLinkActivityStreams()(m, aliasMap); err == nil {
				this := &ActivityStreamsPartOfProperty{
					activitystreamsLinkMember: v,
					alias:                     alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeCollectionActivityStreams()(m, aliasMap); err == nil {
				this := &ActivityStreamsPartOfProperty{
					activitystreamsCollectionMember: v,
					alias:                           alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeCollectionPageActivityStreams()(m, aliasMap); err == nil {
				this := &ActivityStreamsPartOfProperty{
					activitystreamsCollectionPageMember: v,
					alias:                               alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeMentionActivityStreams()(m, aliasMap); err == nil {
				this := &ActivityStreamsPartOfProperty{
					activitystreamsMentionMember: v,
					alias:                        alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeOrderedCollectionActivityStreams()(m, aliasMap); err == nil {
				this := &ActivityStreamsPartOfProperty{
					activitystreamsOrderedCollectionMember: v,
					alias:                                  alias,
				}
				return this, nil
			} else if v, err := mgr.DeserializeOrderedCollectionPageActivityStreams()(m, aliasMap); err == nil {
				this := &ActivityStreamsPartOfProperty{
					activitystreamsOrderedCollectionPageMember: v,
					alias: alias,
				}
				return this, nil
			}
		}
		this := &ActivityStreamsPartOfProperty{
			alias:   alias,
			unknown: i,
		}
		return this, nil
	}
	return nil, nil
}

// NewActivityStreamsPartOfProperty creates a new partOf property.
func NewActivityStreamsPartOfProperty() *ActivityStreamsPartOfProperty {
	return &ActivityStreamsPartOfProperty{alias: ""}
}

// Clear ensures no value of this property is set. Calling HasAny or any of the
// 'Is' methods afterwards will return false.
func (this *ActivityStreamsPartOfProperty) Clear() {
	this.activitystreamsLinkMember = nil
	this.activitystreamsCollectionMember = nil
	this.activitystreamsCollectionPageMember = nil
	this.activitystreamsMentionMember = nil
	this.activitystreamsOrderedCollectionMember = nil
	this.activitystreamsOrderedCollectionPageMember = nil
	this.unknown = nil
	this.iri = nil
}

// GetActivityStreamsCollection returns the value of this property. When
// IsActivityStreamsCollection returns false, GetActivityStreamsCollection
// will return an arbitrary value.
func (this ActivityStreamsPartOfProperty) GetActivityStreamsCollection() vocab.ActivityStreamsCollection {
	return this.activitystreamsCollectionMember
}

// GetActivityStreamsCollectionPage returns the value of this property. When
// IsActivityStreamsCollectionPage returns false,
// GetActivityStreamsCollectionPage will return an arbitrary value.
func (this ActivityStreamsPartOfProperty) GetActivityStreamsCollectionPage() vocab.ActivityStreamsCollectionPage {
	return this.activitystreamsCollectionPageMember
}

// GetActivityStreamsLink returns the value of this property. When
// IsActivityStreamsLink returns false, GetActivityStreamsLink will return an
// arbitrary value.
func (this ActivityStreamsPartOfProperty) GetActivityStreamsLink() vocab.ActivityStreamsLink {
	return this.activitystreamsLinkMember
}

// GetActivityStreamsMention returns the value of this property. When
// IsActivityStreamsMention returns false, GetActivityStreamsMention will
// return an arbitrary value.
func (this ActivityStreamsPartOfProperty) GetActivityStreamsMention() vocab.ActivityStreamsMention {
	return this.activitystreamsMentionMember
}

// GetActivityStreamsOrderedCollection returns the value of this property. When
// IsActivityStreamsOrderedCollection returns false,
// GetActivityStreamsOrderedCollection will return an arbitrary value.
func (this ActivityStreamsPartOfProperty) GetActivityStreamsOrderedCollection() vocab.ActivityStreamsOrderedCollection {
	return this.activitystreamsOrderedCollectionMember
}

// GetActivityStreamsOrderedCollectionPage returns the value of this property.
// When IsActivityStreamsOrderedCollectionPage returns false,
// GetActivityStreamsOrderedCollectionPage will return an arbitrary value.
func (this ActivityStreamsPartOfProperty) GetActivityStreamsOrderedCollectionPage() vocab.ActivityStreamsOrderedCollectionPage {
	return this.activitystreamsOrderedCollectionPageMember
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return an arbitrary value.
func (this ActivityStreamsPartOfProperty) GetIRI() *url.URL {
	return this.iri
}

// GetType returns the value in this property as a Type. Returns nil if the value
// is not an ActivityStreams type, such as an IRI or another value.
func (this ActivityStreamsPartOfProperty) GetType() vocab.Type {
	if this.IsActivityStreamsLink() {
		return this.GetActivityStreamsLink()
	}
	if this.IsActivityStreamsCollection() {
		return this.GetActivityStreamsCollection()
	}
	if this.IsActivityStreamsCollectionPage() {
		return this.GetActivityStreamsCollectionPage()
	}
	if this.IsActivityStreamsMention() {
		return this.GetActivityStreamsMention()
	}
	if this.IsActivityStreamsOrderedCollection() {
		return this.GetActivityStreamsOrderedCollection()
	}
	if this.IsActivityStreamsOrderedCollectionPage() {
		return this.GetActivityStreamsOrderedCollectionPage()
	}

	return nil
}

// HasAny returns true if any of the different values is set.
func (this ActivityStreamsPartOfProperty) HasAny() bool {
	return this.IsActivityStreamsLink() ||
		this.IsActivityStreamsCollection() ||
		this.IsActivityStreamsCollectionPage() ||
		this.IsActivityStreamsMention() ||
		this.IsActivityStreamsOrderedCollection() ||
		this.IsActivityStreamsOrderedCollectionPage() ||
		this.iri != nil
}

// IsActivityStreamsCollection returns true if this property has a type of
// "Collection". When true, use the GetActivityStreamsCollection and
// SetActivityStreamsCollection methods to access and set this property.
func (this ActivityStreamsPartOfProperty) IsActivityStreamsCollection() bool {
	return this.activitystreamsCollectionMember != nil
}

// IsActivityStreamsCollectionPage returns true if this property has a type of
// "CollectionPage". When true, use the GetActivityStreamsCollectionPage and
// SetActivityStreamsCollectionPage methods to access and set this property.
func (this ActivityStreamsPartOfProperty) IsActivityStreamsCollectionPage() bool {
	return this.activitystreamsCollectionPageMember != nil
}

// IsActivityStreamsLink returns true if this property has a type of "Link". When
// true, use the GetActivityStreamsLink and SetActivityStreamsLink methods to
// access and set this property.
func (this ActivityStreamsPartOfProperty) IsActivityStreamsLink() bool {
	return this.activitystreamsLinkMember != nil
}

// IsActivityStreamsMention returns true if this property has a type of "Mention".
// When true, use the GetActivityStreamsMention and SetActivityStreamsMention
// methods to access and set this property.
func (this ActivityStreamsPartOfProperty) IsActivityStreamsMention() bool {
	return this.activitystreamsMentionMember != nil
}

// IsActivityStreamsOrderedCollection returns true if this property has a type of
// "OrderedCollection". When true, use the GetActivityStreamsOrderedCollection
// and SetActivityStreamsOrderedCollection methods to access and set this
// property.
func (this ActivityStreamsPartOfProperty) IsActivityStreamsOrderedCollection() bool {
	return this.activitystreamsOrderedCollectionMember != nil
}

// IsActivityStreamsOrderedCollectionPage returns true if this property has a type
// of "OrderedCollectionPage". When true, use the
// GetActivityStreamsOrderedCollectionPage and
// SetActivityStreamsOrderedCollectionPage methods to access and set this
// property.
func (this ActivityStreamsPartOfProperty) IsActivityStreamsOrderedCollectionPage() bool {
	return this.activitystreamsOrderedCollectionPageMember != nil
}

// IsIRI returns true if this property is an IRI. When true, use GetIRI and SetIRI
// to access and set this property
func (this ActivityStreamsPartOfProperty) IsIRI() bool {
	return this.iri != nil
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this ActivityStreamsPartOfProperty) JSONLDContext() map[string]string {
	m := map[string]string{"https://www.w3.org/TR/activitystreams-vocabulary": this.alias}
	var child map[string]string
	if this.IsActivityStreamsLink() {
		child = this.GetActivityStreamsLink().JSONLDContext()
	} else if this.IsActivityStreamsCollection() {
		child = this.GetActivityStreamsCollection().JSONLDContext()
	} else if this.IsActivityStreamsCollectionPage() {
		child = this.GetActivityStreamsCollectionPage().JSONLDContext()
	} else if this.IsActivityStreamsMention() {
		child = this.GetActivityStreamsMention().JSONLDContext()
	} else if this.IsActivityStreamsOrderedCollection() {
		child = this.GetActivityStreamsOrderedCollection().JSONLDContext()
	} else if this.IsActivityStreamsOrderedCollectionPage() {
		child = this.GetActivityStreamsOrderedCollectionPage().JSONLDContext()
	}
	/*
	   Since the literal maps in this function are determined at
	   code-generation time, this loop should not overwrite an existing key with a
	   new value.
	*/
	for k, v := range child {
		m[k] = v
	}
	return m
}

// KindIndex computes an arbitrary value for indexing this kind of value. This is
// a leaky API detail only for folks looking to replace the go-fed
// implementation. Applications should not use this method.
func (this ActivityStreamsPartOfProperty) KindIndex() int {
	if this.IsActivityStreamsLink() {
		return 0
	}
	if this.IsActivityStreamsCollection() {
		return 1
	}
	if this.IsActivityStreamsCollectionPage() {
		return 2
	}
	if this.IsActivityStreamsMention() {
		return 3
	}
	if this.IsActivityStreamsOrderedCollection() {
		return 4
	}
	if this.IsActivityStreamsOrderedCollectionPage() {
		return 5
	}
	if this.IsIRI() {
		return -2
	}
	return -1
}

// LessThan compares two instances of this property with an arbitrary but stable
// comparison. Applications should not use this because it is only meant to
// help alternative implementations to go-fed to be able to normalize
// nonfunctional properties.
func (this ActivityStreamsPartOfProperty) LessThan(o vocab.ActivityStreamsPartOfProperty) bool {
	idx1 := this.KindIndex()
	idx2 := o.KindIndex()
	if idx1 < idx2 {
		return true
	} else if idx1 > idx2 {
		return false
	} else if this.IsActivityStreamsLink() {
		return this.GetActivityStreamsLink().LessThan(o.GetActivityStreamsLink())
	} else if this.IsActivityStreamsCollection() {
		return this.GetActivityStreamsCollection().LessThan(o.GetActivityStreamsCollection())
	} else if this.IsActivityStreamsCollectionPage() {
		return this.GetActivityStreamsCollectionPage().LessThan(o.GetActivityStreamsCollectionPage())
	} else if this.IsActivityStreamsMention() {
		return this.GetActivityStreamsMention().LessThan(o.GetActivityStreamsMention())
	} else if this.IsActivityStreamsOrderedCollection() {
		return this.GetActivityStreamsOrderedCollection().LessThan(o.GetActivityStreamsOrderedCollection())
	} else if this.IsActivityStreamsOrderedCollectionPage() {
		return this.GetActivityStreamsOrderedCollectionPage().LessThan(o.GetActivityStreamsOrderedCollectionPage())
	} else if this.IsIRI() {
		return this.iri.String() < o.GetIRI().String()
	}
	return false
}

// Name returns the name of this property: "partOf".
func (this ActivityStreamsPartOfProperty) Name() string {
	return "partOf"
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this ActivityStreamsPartOfProperty) Serialize() (interface{}, error) {
	if this.IsActivityStreamsLink() {
		return this.GetActivityStreamsLink().Serialize()
	} else if this.IsActivityStreamsCollection() {
		return this.GetActivityStreamsCollection().Serialize()
	} else if this.IsActivityStreamsCollectionPage() {
		return this.GetActivityStreamsCollectionPage().Serialize()
	} else if this.IsActivityStreamsMention() {
		return this.GetActivityStreamsMention().Serialize()
	} else if this.IsActivityStreamsOrderedCollection() {
		return this.GetActivityStreamsOrderedCollection().Serialize()
	} else if this.IsActivityStreamsOrderedCollectionPage() {
		return this.GetActivityStreamsOrderedCollectionPage().Serialize()
	} else if this.IsIRI() {
		return this.iri.String(), nil
	}
	return this.unknown, nil
}

// SetActivityStreamsCollection sets the value of this property. Calling
// IsActivityStreamsCollection afterwards returns true.
func (this *ActivityStreamsPartOfProperty) SetActivityStreamsCollection(v vocab.ActivityStreamsCollection) {
	this.Clear()
	this.activitystreamsCollectionMember = v
}

// SetActivityStreamsCollectionPage sets the value of this property. Calling
// IsActivityStreamsCollectionPage afterwards returns true.
func (this *ActivityStreamsPartOfProperty) SetActivityStreamsCollectionPage(v vocab.ActivityStreamsCollectionPage) {
	this.Clear()
	this.activitystreamsCollectionPageMember = v
}

// SetActivityStreamsLink sets the value of this property. Calling
// IsActivityStreamsLink afterwards returns true.
func (this *ActivityStreamsPartOfProperty) SetActivityStreamsLink(v vocab.ActivityStreamsLink) {
	this.Clear()
	this.activitystreamsLinkMember = v
}

// SetActivityStreamsMention sets the value of this property. Calling
// IsActivityStreamsMention afterwards returns true.
func (this *ActivityStreamsPartOfProperty) SetActivityStreamsMention(v vocab.ActivityStreamsMention) {
	this.Clear()
	this.activitystreamsMentionMember = v
}

// SetActivityStreamsOrderedCollection sets the value of this property. Calling
// IsActivityStreamsOrderedCollection afterwards returns true.
func (this *ActivityStreamsPartOfProperty) SetActivityStreamsOrderedCollection(v vocab.ActivityStreamsOrderedCollection) {
	this.Clear()
	this.activitystreamsOrderedCollectionMember = v
}

// SetActivityStreamsOrderedCollectionPage sets the value of this property.
// Calling IsActivityStreamsOrderedCollectionPage afterwards returns true.
func (this *ActivityStreamsPartOfProperty) SetActivityStreamsOrderedCollectionPage(v vocab.ActivityStreamsOrderedCollectionPage) {
	this.Clear()
	this.activitystreamsOrderedCollectionPageMember = v
}

// SetIRI sets the value of this property. Calling IsIRI afterwards returns true.
func (this *ActivityStreamsPartOfProperty) SetIRI(v *url.URL) {
	this.Clear()
	this.iri = v
}

// SetType attempts to set the property for the arbitrary type. Returns an error
// if it is not a valid type to set on this property.
func (this *ActivityStreamsPartOfProperty) SetType(t vocab.Type) error {
	if v, ok := t.(vocab.ActivityStreamsLink); ok {
		this.SetActivityStreamsLink(v)
		return nil
	}
	if v, ok := t.(vocab.ActivityStreamsCollection); ok {
		this.SetActivityStreamsCollection(v)
		return nil
	}
	if v, ok := t.(vocab.ActivityStreamsCollectionPage); ok {
		this.SetActivityStreamsCollectionPage(v)
		return nil
	}
	if v, ok := t.(vocab.ActivityStreamsMention); ok {
		this.SetActivityStreamsMention(v)
		return nil
	}
	if v, ok := t.(vocab.ActivityStreamsOrderedCollection); ok {
		this.SetActivityStreamsOrderedCollection(v)
		return nil
	}
	if v, ok := t.(vocab.ActivityStreamsOrderedCollectionPage); ok {
		this.SetActivityStreamsOrderedCollectionPage(v)
		return nil
	}

	return fmt.Errorf("illegal type to set on partOf property: %T", t)
}
