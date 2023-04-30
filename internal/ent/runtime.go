// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/go-faster/bot/internal/ent/gptdialog"
	"github.com/go-faster/bot/internal/ent/prnotification"
	"github.com/go-faster/bot/internal/ent/repository"
	"github.com/go-faster/bot/internal/ent/schema"
	"github.com/go-faster/bot/internal/ent/telegramchannelstate"
	"github.com/go-faster/bot/internal/ent/telegramuserstate"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	gptdialogFields := schema.GPTDialog{}.Fields()
	_ = gptdialogFields
	// gptdialogDescCreatedAt is the schema descriptor for created_at field.
	gptdialogDescCreatedAt := gptdialogFields[6].Descriptor()
	// gptdialog.DefaultCreatedAt holds the default value on creation for the created_at field.
	gptdialog.DefaultCreatedAt = gptdialogDescCreatedAt.Default.(func() time.Time)
	prnotificationFields := schema.PRNotification{}.Fields()
	_ = prnotificationFields
	// prnotificationDescPullRequestTitle is the schema descriptor for pull_request_title field.
	prnotificationDescPullRequestTitle := prnotificationFields[2].Descriptor()
	// prnotification.DefaultPullRequestTitle holds the default value on creation for the pull_request_title field.
	prnotification.DefaultPullRequestTitle = prnotificationDescPullRequestTitle.Default.(string)
	// prnotificationDescPullRequestBody is the schema descriptor for pull_request_body field.
	prnotificationDescPullRequestBody := prnotificationFields[3].Descriptor()
	// prnotification.DefaultPullRequestBody holds the default value on creation for the pull_request_body field.
	prnotification.DefaultPullRequestBody = prnotificationDescPullRequestBody.Default.(string)
	// prnotificationDescPullRequestAuthorLogin is the schema descriptor for pull_request_author_login field.
	prnotificationDescPullRequestAuthorLogin := prnotificationFields[4].Descriptor()
	// prnotification.DefaultPullRequestAuthorLogin holds the default value on creation for the pull_request_author_login field.
	prnotification.DefaultPullRequestAuthorLogin = prnotificationDescPullRequestAuthorLogin.Default.(string)
	repositoryFields := schema.Repository{}.Fields()
	_ = repositoryFields
	// repositoryDescDescription is the schema descriptor for description field.
	repositoryDescDescription := repositoryFields[4].Descriptor()
	// repository.DefaultDescription holds the default value on creation for the description field.
	repository.DefaultDescription = repositoryDescDescription.Default.(string)
	telegramchannelstateFields := schema.TelegramChannelState{}.Fields()
	_ = telegramchannelstateFields
	// telegramchannelstateDescPts is the schema descriptor for pts field.
	telegramchannelstateDescPts := telegramchannelstateFields[2].Descriptor()
	// telegramchannelstate.DefaultPts holds the default value on creation for the pts field.
	telegramchannelstate.DefaultPts = telegramchannelstateDescPts.Default.(int)
	telegramuserstateFields := schema.TelegramUserState{}.Fields()
	_ = telegramuserstateFields
	// telegramuserstateDescQts is the schema descriptor for qts field.
	telegramuserstateDescQts := telegramuserstateFields[1].Descriptor()
	// telegramuserstate.DefaultQts holds the default value on creation for the qts field.
	telegramuserstate.DefaultQts = telegramuserstateDescQts.Default.(int)
	// telegramuserstateDescPts is the schema descriptor for pts field.
	telegramuserstateDescPts := telegramuserstateFields[2].Descriptor()
	// telegramuserstate.DefaultPts holds the default value on creation for the pts field.
	telegramuserstate.DefaultPts = telegramuserstateDescPts.Default.(int)
	// telegramuserstateDescDate is the schema descriptor for date field.
	telegramuserstateDescDate := telegramuserstateFields[3].Descriptor()
	// telegramuserstate.DefaultDate holds the default value on creation for the date field.
	telegramuserstate.DefaultDate = telegramuserstateDescDate.Default.(int)
	// telegramuserstateDescSeq is the schema descriptor for seq field.
	telegramuserstateDescSeq := telegramuserstateFields[4].Descriptor()
	// telegramuserstate.DefaultSeq holds the default value on creation for the seq field.
	telegramuserstate.DefaultSeq = telegramuserstateDescSeq.Default.(int)
}
