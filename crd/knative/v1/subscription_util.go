package v1

import (
	"fmt"
)

func (in Subscriber) String() string {
	if in.Uri != nil {
		return fmt.Sprintf(`URI:%s`, in.Uri.String())
	}
	if in.Service != nil {
		return fmt.Sprintf("SERVICE:%s", in.Service.String())
	}
	if in.Pod != nil {
		return fmt.Sprintf("POD:%s", in.Pod.String())
	}
	panic("Unknown Subscriber type")
}

func (in SubscriberUri) String() string {
	return fmt.Sprintf(`%s://%s`, in.Protocol, in.Uri)
}

func (in SubscriberService) String() string {
	return fmt.Sprintf("%s/%s", in.Namespace, in.Name)
}

func (in SubscriberPod) String() string {
	return fmt.Sprintf("%s/%s", in.Namespace, in.Selector.String())
}

func (in Subscriber) GetTarget() string {
	if in.Uri != nil {
		return in.Uri.GetTarget()
	}
	if in.Service != nil {
		return in.Service.GetTarget()
	}
	//TODO:实现pod的target
	//if in.Pod != nil {
	//	return in.Pod.String()
	//}
	panic("Unknown Subscriber type")
}

func (in SubscriberUri) GetTarget() string {
	return in.String()
}

func (in SubscriberService) GetTarget() string {
	return fmt.Sprintf(`http://%s.%s:%d%s`, in.Name, in.Namespace, in.Port, in.Path)
}
