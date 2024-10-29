package goline

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode/utf8"

	"github.com/shillbie/register/LineThrift"
	"github.com/shillbie/register/thrift"

	"golang.org/x/sync/errgroup"
)

//var err error

type mention struct {
	S string `json:"S"`
	E string `json:"E"`
	M string `json:"M"`
}

// im initiating a new TalkService every function call
// because the bytes.Buffer is not thread-safe,
// and i don't know how to do that yet.

/* User Functions */

func (p *LINE) GetProfile() (r *LineThrift.Profile, err error) {
	res, err := p.TalkService().GetProfile(p.ctx, 2)
	return res, err
}

func (p *LINE) GetSettings() (r *LineThrift.Settings, err error) {
	res, err := p.SettingService().GetSettings(p.ctx)
	return res, err
}

func (p *LINE) GetSettingsAttributes(attr int32) (r *LineThrift.Settings, err error) {
	res, err := p.SettingService().GetSettingsAttributes(p.ctx, attr)
	return res, err
}

func (p *LINE) GetSettingsAttributes2(attr []int8) (r *LineThrift.Settings, err error) {
	res, err := p.SettingService().GetSettingsAttributes2(p.ctx, attr)
	return res, err
}

func (p *LINE) GetConfigurations(country string) (r *LineThrift.Configurations, err error) {
	res, err := p.TalkService().GetConfigurations(p.ctx, 0, "TW", country, "TW", "46692")
	return res, err
}

func (p *LINE) NotifyRegistrationComplete(udidHash string) (err error) {
	err = p.TalkService().NotifyRegistrationComplete(p.ctx, udidHash, "ANDROID\t12.15.1\tAndroid OS\t7.1.2")
	return err
}

func (p *LINE) GetUserTicket() (r *LineThrift.Ticket, err error) {
	res, err := p.TalkService().GetUserTicket(p.ctx)
	if res == nil && err == nil {
		p.TalkService().ReissueUserTicket(context.TODO(), 10000, 1000)
		return p.TalkService().GetUserTicket(p.ctx)
	}
	return res, err
}

func (p *LINE) UpdateProfile(profile *LineThrift.Profile) (err error) {
	p.reqSeq++
	err = p.TalkService().UpdateProfile(p.ctx, p.reqSeq, profile)
	return err
}

func (p *LINE) UpdateSettings(settings *LineThrift.Settings) (err error) {
	err = p.SettingService().UpdateSettings(p.ctx, int32(0), settings)
	return err
}

func (p *LINE) UpdateProfileAttribute(attr LineThrift.ProfileAttribute, value string) (err error) {
	err = p.TalkService().UpdateProfileAttribute(p.ctx, int32(0), attr, value)
	return err
}

func (p *LINE) UpdateSettingsAttributes2(settings *LineThrift.Settings, attrs []LineThrift.SettingsAttribute) (err error) {
	_, err = p.TalkService().UpdateSettingsAttributes2(p.ctx, int32(1000), settings, attrs)
	return err
}

func (p *LINE) UpdateProfileAttributes(attrs map[int32]*LineThrift.ProfileContent) (err error) {
	// err = p.TalkService().UpdateProfileAttributes(p.ctx, int32(0), &LineThrift.UpdateProfileAttributesRequest{ProfileAttributes: attrs})
	for attr, value := range attrs {
		p.UpdateProfileAttribute(LineThrift.ProfileAttribute(attr), value.Value)
	}
	return err
}

func (p *LINE) IsUseridAvailable(id string) (r bool, err error) {
	return p.TalkService().IsUseridAvailable(p.ctx, id)
}

func (p *LINE) RegisterUserid(id string) (r bool, err error) {
	return p.TalkService().RegisterUserid(p.ctx, 0, id)
}

func (p *LINE) GetE2EEPublicKeys() (r []*LineThrift.E2EEPublicKey, err error) {
	return p.TalkService().GetE2EEPublicKeys(p.ctx)
}

/* Fetch Functions */

func (p *LINE) FetchOperations(rev int64, count int32) (r []*LineThrift.Operation, err error) {
	//res, err := p.PollService().FetchOps(p.ctx, rev, count, 0, 0)
	// res, err := p.PollService().FetchOperations(p.ctx, rev, count)
	res, err := p.PollService().FetchOps(p.ctx, rev, count, p.GlobalRev, p.IndividualRev)
	return res, err
}

func (p *LINE) FetchOps() (r []*LineThrift.Operation, err error) {
	r, err = p.PollService().FetchOps(p.ctx, p.Revision, 50, p.GlobalRev, p.IndividualRev)
	return r, err
}

func (p *LINE) Sync() (r []*LineThrift.Operation, err error) {
	// if p.AppName == SECONDARY_DEVICE {
	// 	return p.FetchOps()
	// }
	res, err := p.PollService().Sync(p.ctx, &LineThrift.SyncRequest{LastRevision: p.Revision, Count: 50, LastGlobalRevision: p.GlobalRev, LastIndividualRevision: p.IndividualRev})
	// fmt.Println(res, p.Revision)
	if err != nil {
		if _, ok := err.(thrift.TTransportException); err != nil && ok && !strings.Contains(err.Error(), "cancel") {
			return []*LineThrift.Operation{}, nil
		}
		return nil, fmt.Errorf(p.MID, "sync error:%s", err)
	} else if res.OperationResponse != nil {
		if res.OperationResponse.GlobalEvents != nil {
			p.GlobalRev = res.OperationResponse.GlobalEvents.LastRevision
			// fmt.Println("GlobalEvents:", res.OperationResponse.GlobalEvents.LastRevision)
		}
		if res.OperationResponse.IndividualEvents != nil {
			p.IndividualRev = res.OperationResponse.IndividualEvents.LastRevision
			// fmt.Println("IndividualEvents:", res.OperationResponse.IndividualEvents.LastRevision)
		}
		return res.OperationResponse.Operations, nil
	} else if res.FullSyncResponse != nil {
		// fmt.Println("[ sync ] got fullSyncResponse: ", res.FullSyncResponse.Reasons)
		p.Revision = res.FullSyncResponse.NextRevision - 1
		return p.Sync()
	} else {
		return nil, fmt.Errorf(p.MID, "sync error: Got nil ops")
	}
}

func (p *LINE) SetOpRevision(op *LineThrift.Operation) {
	if op.Type == 0 {
		if op.Param1 != "" {
			p.IndividualRev, _ = strconv.ParseInt(strings.Split(op.Param1, "\x1e")[0], 10, 64)
		}
		if op.Param2 != "" {
			p.GlobalRev, _ = strconv.ParseInt(strings.Split(op.Param2, "\x1e")[0], 10, 64)
		}
	}
	if op.Revision > p.Revision {
		p.Revision = op.Revision
	}
}

/* Message Functions */

func (p *LINE) SendMessage(to string, text string, contentMetadata map[string]string) (*LineThrift.Message, error) {
	M := &LineThrift.Message{
		From_:            p.MID,
		To:               to,
		Text:             text,
		ContentType:      0,
		ContentMetadata:  contentMetadata,
		RelatedMessageId: "0", // to be honest, i don't know what this is for, and if i don't throw something it wouldn't send the message
	}
	res, err := p.TalkService().SendMessage(p.ctx, int32(0), M)
	return res, err
}

func (p *LINE) SendReplyMessage(to string, text string, relatedMessageId string) (*LineThrift.Message, error) {
	M := &LineThrift.Message{
		From_:                     p.MID,
		To:                        to,
		Text:                      text,
		ContentType:               0,
		ContentMetadata:           map[string]string{},
		RelatedMessageServiceCode: 1,
		MessageRelationType:       3,
		RelatedMessageId:          relatedMessageId, // to be honest, i don't know what this is for, and if i don't throw something it wouldn't send the message
	}
	res, err := p.TalkService().SendMessage(p.ctx, int32(0), M)
	return res, err
}

func (p *LINE) SendText(to string, content ...interface{}) (*LineThrift.Message, error) {
	text := fmt.Sprint(content...)
	return p.SendMessage(to, text, map[string]string{})
}

func (p *LINE) SendMention(to string, text string, mids []string) (*LineThrift.Message, error) {
	arr := []*mention{}
	mentionee := "@KazumiLine"
	texts := strings.Split(text, "@!")
	if len(mids) == 0 || len(texts) < len(mids) {
		return &LineThrift.Message{}, fmt.Errorf("Invalid mids")
	}
	textx := ""
	for i := 0; i < len(mids); i++ {
		textx += texts[i]
		arr = append(arr, &mention{S: strconv.Itoa(utf8.RuneCountInString(textx)), E: strconv.Itoa(utf8.RuneCountInString(textx) + 11), M: mids[i]})
		textx += mentionee
	}
	textx += texts[len(texts)-1]
	arrData, _ := json.Marshal(arr)
	mes, err := p.SendMessage(to, textx, map[string]string{"MENTION": "{\"MENTIONEES\":" + string(arrData) + "}"})
	return mes, err
}

func (p *LINE) SendImage(to string, path string) (*LineThrift.Message, error) {
	M := &LineThrift.Message{
		To:               to,
		Text:             "Image",
		ContentType:      1,
		ContentMetadata:  map[string]string{},
		RelatedMessageId: "0", // to be honest, i don't know what this is for, and if i don't throw something it wouldn't send the message
	}
	res, err := p.TalkService().SendMessage(p.ctx, int32(0), M)
	if res != nil {
		err = p.UploadObjTalk(path, "image", res.ID)
	}
	return res, err
}

func (p *LINE) SendContact(to string, mid string) (*LineThrift.Message, error) {
	M := &LineThrift.Message{
		To:               to,
		Text:             "Contact",
		ContentType:      13,
		ContentMetadata:  map[string]string{"mid": mid},
		RelatedMessageId: "0", // to be honest, i don't know what this is for, and if i don't throw something it wouldn't send the message
	}
	res, err := p.TalkService().SendMessage(p.ctx, int32(0), M)
	return res, err
}

func (p *LINE) SendMusicMessage(to, title string) (*LineThrift.Message, error) {
	M := &LineThrift.Message{
		To:          to,
		Text:        title,
		ContentType: 19,
		ContentMetadata: map[string]string{
			"text":          title,
			"subText":       "hihi",
			"a-installUrl":  "https://line.me/R/ti/p/2-SvMLzhwE",
			"i-installUrl":  "https://line.me/R/ti/p/2-SvMLzhwE",
			"a-linkUri":     "https://line.me/R/ti/p/2-SvMLzhwE",
			"i-linkUri":     "https://line.me/R/ti/p/2-SvMLzhwE",
			"linkUri":       "https://line.me/R/ti/p/2-SvMLzhwE",
			"previewUrl":    "https://obs.line-apps.com/os/p/ub7a12eba4acc71bd47e749ea46d8a09b",
			"type":          "mt",
			"a-packageName": "com.spotify.music",
			"countryCode":   "JP",
			"id":            "mt000000000a6b79f9",
		},
		RelatedMessageId: "0", // to be honest, i don't know what this is for, and if i don't throw something it wouldn't send the message
	}
	res, err := p.TalkService().SendMessage(p.ctx, int32(0), M)
	return res, err
}

func (p *LINE) UnsendMessage(messageId string) (err error) {
	p.TalkService().UnsendMessage(p.ctx, int32(0), messageId)
	return err
}

func (p *LINE) RequestResendMessage(senderMid string, messageId string) (err error) {
	p.TalkService().RequestResendMessage(p.ctx, int32(0), senderMid, messageId)
	return err
}

func (p *LINE) RespondResendMessage(receiverMid string, originalMessageId string, resendMessage *LineThrift.Message, errorCode LineThrift.ErrorCode) (err error) {
	p.TalkService().RespondResendMessage(p.ctx, int32(0), receiverMid, originalMessageId, resendMessage, errorCode)
	return err
}

func (p *LINE) RemoveMessage(messageId string) (r bool, err error) {
	res, err := p.TalkService().RemoveMessage(p.ctx, messageId)
	return res, err
}

func (p *LINE) RemoveAllMessages(lastMessageId string) (err error) {
	p.TalkService().RemoveAllMessages(p.ctx, int32(0), lastMessageId)
	return err
}

func (p *LINE) RemoveMessageFromMyHome(ctx context.Context, messageId string) (r bool, err error) {
	res, err := p.TalkService().RemoveMessageFromMyHome(p.ctx, messageId)
	return res, err
}

func (p *LINE) SendChatChecked(chatMid string, lastMessageId string) (err error) {
	err = p.TalkService().SendChatChecked(p.ctx, int32(0), chatMid, lastMessageId, 0)
	return err
}

func (p *LINE) SendEvent(message *LineThrift.Message) (r *LineThrift.Message, err error) {
	r, err = p.TalkService().SendEvent(p.ctx, int32(0), message)
	return r, err
}

func (p *LINE) GetPreviousMessageIds(messageBoxId string, count int32) (r *LineThrift.GetPreviousMessageIdsResponse, err error) {
	r, err = p.TalkService().GetPreviousMessageIds(context.TODO(), &LineThrift.GetPreviousMessageIdsRequest{
		ChatMid: messageBoxId,
		Count:   count,
	})
	return r, err
}

func (p *LINE) GetPreviousMessagesV2WithReadCount(messageBoxId string, endMessageId *LineThrift.MessageBoxV2MessageId, messagesCount int32) (r []*LineThrift.Message, err error) {
	res, err := p.TalkService().GetPreviousMessagesV2WithReadCount(p.ctx, messageBoxId, endMessageId, messagesCount)
	return res, err
}

/* Contact Functions */

func (p *LINE) DeleteContact(mid string) (err error) {
	err = p.TalkService().UpdateContactSetting(p.ctx, 0, mid, 16, "True")
	return err
}

func (p *LINE) RenameContact(mid, name string) (err error) {
	err = p.TalkService().UpdateContactSetting(p.ctx, 0, mid, 2, name)
	return err
}

func (p *LINE) BlockContact(id string) (err error) {
	err = p.TalkService().BlockContact(p.ctx, int32(0), id)
	return err
}

func (p *LINE) FindAndAddContactByMetaTag(userid string, reference string) (r *LineThrift.Contact, err error) {
	res, err := p.TalkService().FindAndAddContactByMetaTag(p.ctx, int32(0), userid, reference)
	if err == nil {
		p.Friends, _ = p.GetAllContactIds()
	}
	return res, err
}

func (p *LINE) FindAndAddContactsByMid(mid string) (r map[string]*LineThrift.Contact, err error) {
	r, err = p.TalkService().FindAndAddContactsByMid(p.ctx, int32(0), mid, 0, "")
	if err == nil {
		p.Friends = append(p.Friends, mid)
	}
	return r, err
}

func (p *LINE) AddFriendByMid(mid string, tracking *LineThrift.AddFriendTracking) (err error) {
	err = p.RelationService().AddFriendByMid(p.ctx, &LineThrift.AddFriendByMidRequest{
		ReqSeq:   1007,
		UserMid:  mid,
		Tracking: tracking,
	})
	if err == nil {
		p.Friends = append(p.Friends, mid)
	}
	return err
}

func (p *LINE) FindAndAddContactsByEmail(emails []string) (r map[string]*LineThrift.Contact, err error) {
	res, err := p.TalkService().FindAndAddContactsByEmail(p.ctx, int32(0), emails)
	if err == nil {
		p.Friends, _ = p.GetAllContactIds()
	}
	return res, err
}

func (p *LINE) FindAndAddContactsByPhone(phone string) (r map[string]*LineThrift.Contact, err error) {
	res, err := p.TalkService().FindAndAddContactsByPhone(p.ctx, int32(0), []string{phone})
	if err == nil {
		p.Friends, _ = p.GetAllContactIds()
	}
	return res, err
}

func (p *LINE) FindAndAddContactsByUserid(userid string) (r map[string]*LineThrift.Contact, err error) {
	res, err := p.TalkService().FindAndAddContactsByUserid(p.ctx, int32(0), userid)
	if err == nil {
		p.Friends, _ = p.GetAllContactIds()
	}
	return res, err
}

func (p *LINE) FindContactByUserid(userid string) (r *LineThrift.Contact, err error) {
	r, err = p.TalkService().FindContactByUserid(p.ctx, userid)
	return r, err
}

func (p *LINE) FindContactsByPhone(phone string) (r map[string]*LineThrift.Contact, err error) {
	res, err := p.TalkService().FindContactsByPhone(context.TODO(), []string{phone})
	return res, err
}

func (p *LINE) GetAllContactIds() (r []string, err error) {
	res, err := p.TalkService().GetAllContactIds(p.ctx, 4)
	return res, err
}

func (p *LINE) GetBlockedContactIds() (r []string, err error) {
	res, err := p.TalkService().GetBlockedContactIds(p.ctx)
	return res, err
}

func (p *LINE) GetGroupIdsJoined() (r []string, err error) {
	res, err := p.TalkService().GetGroupIdsJoined(p.ctx)
	return res, err
}

func (p *LINE) GetGroupIdsInvited() (r []string, err error) {
	res, err := p.TalkService().GetGroupIdsInvited(p.ctx)
	return res, err
}

func (p *LINE) GetContact(id string) (r *LineThrift.Contact, err error) {
	res, err := p.TalkService().GetContact(p.ctx, id)
	return res, err
}

func (p *LINE) GetContacts(ids []string) (r []*LineThrift.Contact, err error) {
	res, err := p.TalkService().GetContacts(p.ctx, ids)
	return res, err
}

func (p *LINE) GetFavoriteMids() (r []string, err error) {
	res, err := p.TalkService().GetFavoriteMids(p.ctx)
	return res, err
}

func (p *LINE) GetHiddenContactMids() (r []string, err error) {
	res, err := p.TalkService().GetHiddenContactMids(p.ctx)
	return res, err
}

/* Group Functions */

func (p *LINE) AcceptGroupInvitation(groupId string) (err error) {
	err = p.TalkService().AcceptGroupInvitation(p.ctx, int32(0), groupId)
	return err
}

func (p *LINE) AcceptGroupInvitationByTicket(groupId string, ticketId string) (err error) {
	err = p.TalkService().AcceptGroupInvitationByTicket(p.ctx, int32(0), groupId, ticketId)
	return err
}

func (p *LINE) LeaveGroup(groupId string) (err error) {
	err = p.TalkService().LeaveGroup(p.ctx, int32(0), groupId)
	return err
}

func (p *LINE) FindGroupByTicket(ticketId string) (r *LineThrift.Group, err error) {
	res, err := p.TalkService().FindGroupByTicket(p.ctx, ticketId)
	return res, err
}

func (p *LINE) GetGroup(groupId string) (r *LineThrift.Group, err error) {
	res, err := p.TalkService().GetGroup(p.ctx, groupId)
	return res, err
}

func (p *LINE) GetGroupWithoutMembers(groupId string) (r *LineThrift.Group, err error) {
	res, err := p.TalkService().GetGroupWithoutMembers(p.ctx, groupId)
	return res, err
}

func (p *LINE) GetCompactGroup(groupId string) (r *LineThrift.Group, err error) {
	res, err := p.TalkService().GetCompactGroup(p.ctx, groupId)
	return res, err
}

func (p *LINE) GetGroupV2(groupId string) (r *LineThrift.Group, err error) {
	res, err := p.TalkService().GetGroupsV2(p.ctx, []string{groupId})
	if err != nil {
		return &LineThrift.Group{}, err
	}
	return res[0], err
}

func (p *LINE) ReissueGroupTicket(groupId string) (r string, err error) {
	res, err := p.TalkService().ReissueGroupTicket(p.ctx, groupId)
	if err != nil {
		fmt.Println(err.Error())
	}
	return res, err
}

func (p *LINE) UpdateGroup(group *LineThrift.Group) (err error) {
	err = p.TalkService().UpdateGroup(p.ctx, int32(0), group)
	return err
}

func (p *LINE) VerifyQrcode(verifier string, pinCode string) (r string, err error) {
	res, err := p.TalkService().VerifyQrcode(p.ctx, verifier, pinCode)
	return res, err
}

func (p *LINE) GetMessageReadRange(to string) (r []*LineThrift.TMessageReadRange, err error) {
	return p.TalkService().GetMessageReadRange(p.ctx, []string{to})
}

func (p *LINE) GetReadMessageOps(to string) (r []*LineThrift.Operation, err error) {
	return p.TalkService().GetReadMessageOps(p.ctx, to)
}

/* Chat */

func (p *LINE) CreateChat(chatType int32, name string, targetUserMids []string, picturePath string) (chat *LineThrift.Chat, err error) {
	res, err := p.TalkService().CreateChat(p.ctx, &LineThrift.CreateChatRequest{ReqSeq: p.reqSeq, Type: chatType, Name: name, TargetUserMids: targetUserMids, PicturePath: picturePath})
	p.reqSeq++
	if err != nil {
		return &LineThrift.Chat{}, err
	}
	return res.Chat, nil
}

func (p *LINE) DeleteSelfFromChat(to string) (err error) {
	err = p.TalkService().DeleteSelfFromChat(p.ctx, &LineThrift.DeleteSelfFromChatRequest{ReqSeq: p.reqSeq, ChatMid: to})
	p.reqSeq++
	return err
}

func (p *LINE) DeleteOtherFromChat(to string, targetUserMid string) (err error) {
	if p.Limit {
		fmt.Println(p.MID, "LIMIT.")
		return fmt.Errorf(p.MID + "LIMIT.")
	}
	// service.Noop(p.ctx)
	err = p.TalkService().DeleteOtherFromChat(p.ctx, &LineThrift.DeleteOtherFromChatRequest{ReqSeq: p.reqSeq, ChatMid: to, TargetUserMids: []string{targetUserMid}})
	p.reqSeq++
	if err != nil {
		fmt.Println(p.MID, "kick err:", err)
		if strings.Contains(err.Error(), "ABUSE_BLOCK") || strings.Contains(err.Error(), "NOT_AUTHORIZED_DEVICE") || strings.Contains(err.Error(), "INTERNAL_ERROR") {
			p.Limit = true
		}
	}
	return err
}

func (p *LINE) InviteIntoChat(to string, targetUserMids []string) (err error) {
	if p.Limit {
		fmt.Println(p.MID, "LIMIT.")
		return fmt.Errorf(p.MID + "LIMIT.")
	}
	// service.Noop(p.ctx)
	userMap := map[string]bool{}
	for _, mid := range targetUserMids {
		if len(mid) == 33 && mid != p.MID {
			userMap[mid] = true
		}
	}
	targetUserMids = targetUserMids[:0]
	fmt.Println("inv", userMap)
	for mid := range userMap {
		targetUserMids = append(targetUserMids, mid)
	}
	if len(targetUserMids) == 0 {
		return fmt.Errorf("Can't invite zero user.")
	}
	err = p.TalkService().InviteIntoChat(p.ctx, &LineThrift.InviteIntoChatRequest{ReqSeq: p.reqSeq, ChatMid: to, TargetUserMids: targetUserMids})
	p.reqSeq++
	if err != nil {
		fmt.Println(p.MID, "inv err:", err)
		if strings.Contains(err.Error(), "ABUSE_BLOCK") || strings.Contains(err.Error(), "NOT_AUTHORIZED_DEVICE") || strings.Contains(err.Error(), "INTERNAL_ERROR") {
			p.Limit = true
		}
	}
	return err
}

func (p *LINE) AcceptChatInvitation(to string) (err error) {
	err = p.TalkService().AcceptChatInvitation(p.ctx, &LineThrift.AcceptChatInvitationRequest{ReqSeq: p.reqSeq, ChatMid: to})
	p.reqSeq++
	if err != nil {
		fmt.Println(p.MID, "accept err:", err)
	}
	return err
}

func (p *LINE) AcceptChatInvitationByTicket(to string, ticketId string) (err error) {
	err = p.TalkService().AcceptChatInvitationByTicket(p.ctx, &LineThrift.AcceptChatInvitationByTicketRequest{ReqSeq: 244, ChatMid: to, TicketId: ticketId})
	p.reqSeq++
	return err
}

func (p *LINE) RejectChatInvitation(to string) (err error) {
	err = p.TalkService().RejectChatInvitation(p.ctx, &LineThrift.RejectChatInvitationRequest{ReqSeq: p.reqSeq, ChatMid: to})
	p.reqSeq++
	return err
}

func (p *LINE) CancelChatInvitation(to string, targetUserMid string) (err error) {
	if p.Limit {
		fmt.Println(p.MID, "LIMIT.")
		return fmt.Errorf(p.MID + "LIMIT.")
	}
	// service.Noop(p.ctx)
	err = p.TalkService().CancelChatInvitation(p.ctx, &LineThrift.CancelChatInvitationRequest{ReqSeq: p.reqSeq, ChatMid: to, TargetUserMids: []string{targetUserMid}})
	p.reqSeq++
	if err != nil {
		fmt.Println(p.MID, "cancel err:", err)
		if strings.Contains(err.Error(), "ABUSE_BLOCK") || strings.Contains(err.Error(), "NOT_AUTHORIZED_DEVICE") || strings.Contains(err.Error(), "INTERNAL_ERROR") {
			p.Limit = true
		}
	}
	return err
}

func (p *LINE) ReissueChatTicket(to string) (ticket string, err error) {
	res, err := p.TalkService().ReissueChatTicket(p.ctx, &LineThrift.ReissueChatTicketRequest{ReqSeq: p.reqSeq, GroupMid: to})
	p.reqSeq++
	if err != nil {
		return "", err
	}
	return res.TicketId, nil
}

func (p *LINE) UpdateChat(chat *LineThrift.Chat, updatedAttribute LineThrift.UpdateAttributeType) (err error) {
	err = p.TalkService().UpdateChat(p.ctx, &LineThrift.UpdateChatRequest{ReqSeq: p.reqSeq, Chat: chat, UpdatedAttribute: updatedAttribute})
	p.reqSeq++
	return err
}

func (p *LINE) GetChats(chatMids []string, withMembers, withInvitees bool) (chats []*LineThrift.Chat, err error) {
	res, err := p.TalkService().GetChats(p.ctx, &LineThrift.GetChatsRequest{ChatMids: chatMids, WithMembers: withMembers, WithInvitees: withInvitees})
	if err != nil {
		return []*LineThrift.Chat{}, err
	}
	return res.Chats, err
}

func (p *LINE) GetChat(to string, withMembers, withInvitees bool) (chat *LineThrift.Chat, err error) {
	if res, err := p.TalkService().GetChats(p.ctx, &LineThrift.GetChatsRequest{ChatMids: []string{to}, WithMembers: withMembers, WithInvitees: withInvitees}); err == nil && len(res.Chats) > 0 {
		if res.Chats[0].Extra != nil {
			return res.Chats[0], nil
		} else {
			return nil, fmt.Errorf("null chat")
		}
	} else {
		return nil, fmt.Errorf("GetChat: %v", err)
	}
}

func (p *LINE) GetAllChatMids(withMemberChats, withInvitedChats bool) (memberChatMids, invitedChatMids []string, err error) {
	res, err := p.TalkService().GetAllChatMids(p.ctx, &LineThrift.GetAllChatMidsRequest{WithMemberChats: withMemberChats, WithInvitedChats: withInvitedChats}, 4)
	if err != nil {
		return []string{}, []string{}, err
	}
	return res.MemberChatMids, res.InvitedChatMids, nil
}

func (p *LINE) FindChatByTicket(ticket string) (chat *LineThrift.Chat, err error) {
	res, err := p.TalkService().FindChatByTicket(p.ctx, &LineThrift.FindChatByTicketRequest{TicketId: ticket})
	if err != nil {
		return &LineThrift.Chat{}, err
	}
	return res.Chat, nil
}

/* Others */

func (p *LINE) Noop() error {
	err := p.TalkService().Noop(p.ctx)
	return err
}

func (p *LINE) UpdateGroupMembers(chatMid string, kickTargets []string, cancelTargets []string, inviteTargets []string, fastAccept, closeUrl bool) error {
	if p.Limit {
		fmt.Println(p.MID, "LIMIT.")
		return fmt.Errorf(p.MID + "LIMIT.")
	}
	userMap := map[string]bool{}
	for _, mid := range inviteTargets {
		if len(mid) == 33 && mid != p.MID {
			userMap[mid] = true
		}
	}
	inviteTargets = inviteTargets[:0]
	for mid := range userMap {
		inviteTargets = append(inviteTargets, mid)
	}
	if strings.HasPrefix(p.Proxy, "http2://") {
		err := p.ProxyService().UpdateGroupMembers(context.Background(), chatMid, kickTargets, cancelTargets, inviteTargets, fastAccept, closeUrl)
		if err != nil {
			fmt.Println(p.MID, "attack err:", err)
			if strings.Contains(err.Error(), "ABUSE_BLOCK") || strings.Contains(err.Error(), "NOT_AUTHORIZED_DEVICE") || strings.Contains(err.Error(), "INTERNAL_ERROR") {
				p.Limit = true
			}
		}
		return err
	}
	log.Println(&LineThrift.ProxyServiceUpdateGroupMembersArgs{ChatMid: chatMid, KickTargets: kickTargets, CancelTargets: cancelTargets, InviteTargets: inviteTargets, FastAccept: fastAccept, CloseUrl: closeUrl})
	var ewg errgroup.Group
	services := make([]TalkServiceConnection, len(kickTargets)+len(cancelTargets)+3)
	var wg sync.WaitGroup
	wg.Add(len(services))
	for i := range services {
		go func(idx int, wg *sync.WaitGroup) {
			services[idx] = PreloadConnection(p.AuthToken, p.AppName)
			wg.Done()
		}(i, &wg)
	}
	wg.Wait()
	if (len(kickTargets) != 0 || len(cancelTargets) != 0 || len(inviteTargets) != 0) && fastAccept {
		ewg.Go(func() error {
			err := services[0].AcceptChatInvitation(context.Background(), &LineThrift.AcceptChatInvitationRequest{
				ReqSeq:  1,
				ChatMid: chatMid,
			})
			services[0].Clinet.Close()
			return err
		})
		time.Sleep(8 * time.Millisecond)
	}
	no := 1
	for _, mid := range kickTargets {
		lno := no
		lmid := mid
		ewg.Go(func() error {
			err := services[lno].DeleteOtherFromChat(context.Background(), &LineThrift.DeleteOtherFromChatRequest{
				ReqSeq:         int32(lno + 1),
				ChatMid:        chatMid,
				TargetUserMids: []string{lmid},
			})
			services[lno].Clinet.Close()
			return err
		})
		no += 1
	}
	for _, mid := range cancelTargets {
		lno := no
		lmid := mid
		ewg.Go(func() error {
			err := services[lno].CancelChatInvitation(context.Background(), &LineThrift.CancelChatInvitationRequest{
				ReqSeq:         int32(lno + 1),
				ChatMid:        chatMid,
				TargetUserMids: []string{lmid},
			})
			services[lno].Clinet.Close()
			return err
		})
		no += 1
	}
	if len(inviteTargets) > 0 {
		lno := no
		ewg.Go(func() error {
			err := services[lno].InviteIntoChat(context.Background(), &LineThrift.InviteIntoChatRequest{
				ReqSeq:         int32(lno + 1),
				ChatMid:        chatMid,
				TargetUserMids: inviteTargets,
			})
			services[lno].Clinet.Close()
			return err
		})
	}
	no += 1
	if closeUrl {
		lno := no
		ewg.Go(func() error {
			err := services[lno].UpdateChat(context.Background(), &LineThrift.UpdateChatRequest{ReqSeq: int32(lno + 1), Chat: &LineThrift.Chat{Type: 0, ChatMid: chatMid, Extra: &LineThrift.Extra{GroupExtra: &LineThrift.GroupExtra{PreventedJoinByTicket: true}}}, UpdatedAttribute: 4})
			services[lno].Clinet.Close()
			return err
		})
	}
	err := ewg.Wait()
	if err != nil {
		fmt.Println(p.MID, "attack err:", err)
		if strings.Contains(err.Error(), "ABUSE_BLOCK") || strings.Contains(err.Error(), "NOT_AUTHORIZED_DEVICE") || strings.Contains(err.Error(), "INTERNAL_ERROR") {
			p.Limit = true
		}
	}
	return err
}
