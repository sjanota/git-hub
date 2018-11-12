package git_pr

import (
	"github.com/jawher/mow.cli"
	"github.com/sjanota/git-pr/pkg/git"
	"github.com/sjanota/git-pr/pkg/ops"
)

type comment struct {
	comment *string
	append  *bool
	repo    git.Repo
}

func (c *comment) Configure(app *cli.Cli) {
	app.Command("comment cm", "Edit pull request comment", func(cmd *cli.Cmd) {
		c.comment = cmd.StringOpt("comment m", "", "Text of the comment to use instead of opening editor")
		c.append = cmd.BoolOpt("append a", false,
			"If set text will be appended to existing comment "+
				"in new line instead of overwriting it. Have no effect if edited in text editor")
		cmd.Action = c.action
	})
}

func (c comment) action() {
	var editor git.CommentEditor
	if *c.comment != "" {
		editor = c.repo.StaticCommentEditor(*c.comment, *c.append)
	} else {
		editor = c.repo.FileCommentEditor()
	}

	err := ops.Comment(c.repo, editor)
	if err != nil {
		panic(err)
	}
}
