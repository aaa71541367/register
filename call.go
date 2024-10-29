package goline

import (
	"context"

	"github.com/shillbie/register/LineThrift"
	// "fmt"
)

func (p *LINE) AcquireCallRoute(to string) (r []string, err error) {
	res, err := p.TalkService().AcquireCallRoute(p.ctx, to)
	return res, err
}

func (p *LINE) AcquireGroupCallRoute(chatMid string, mediaType LineThrift.GroupCallMediaType) (r *LineThrift.GroupCallRoute, err error) {
	res, err := p.CallService().AcquireGroupCallRoute(p.ctx, chatMid, mediaType)
	return res, err
}

func (p *LINE) GetGroupCall(chatMid string) (r *LineThrift.GroupCall, err error) {
	res, err := p.CallService().GetGroupCall(p.ctx, chatMid)
	return res, err
}

func (p *LINE) InviteIntoGroupCall(chatMid string, memberMids []string, mediaType LineThrift.GroupCallMediaType) (err error) {
	err = p.CallService().InviteIntoGroupCall(p.ctx, chatMid, memberMids, mediaType)
	return err
}

func (p *LINE) FindContactByUseridWithoutAbuseBlockForChannel(userid string) (r *LineThrift.Contact, err error) {
	res, err := p.CallService().FindContactByUseridWithoutAbuseBlockForChannel(context.TODO(), userid)
	return res, err
}
