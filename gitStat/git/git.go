package git

import (
	"fmt"
	"github.com/cheneylew/gotools/tool"
	"github.com/cheneylew/goutil/utils"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

func Shell(workDir string, cmd string) string {
	//fmt.Println(cmd)
	utils.FileWriteString(workDir+"/tmp.sh", fmt.Sprintf("#!/usr/bin/env bash \n%v", cmd))
	c := exec.Command("sh", "tmp.sh")
	c.Dir = workDir
	out, err := c.Output()
	if err != nil {
		//panic(err)
	}
	os.Remove(workDir+"/tmp.sh")

	return string(out)
}

func RemoteBranches(dir string) []string {
	cmd := "git branch -r"
	res := Shell(dir, cmd)
	resArr := tool.ArrFilter(func(s interface{}) bool { return s.(string) != ""}, strings.Split(res, "\n")).([]string)
	branches := tool.ArrMap(func(s interface{}) string {return strings.TrimSpace(s.(string))}, resArr).([]string)
	return branches
}

func LocalBranches(dir string) []string {
	cmd := "git branch"
	res := Shell(dir, cmd)
	resArr := tool.ArrFilter(func(s interface{}) bool { return s.(string) != ""}, strings.Split(res, "\n")).([]string)
	branches := tool.ArrMap(func(s interface{}) string {return strings.TrimSpace(s.(string))}, resArr).([]string)
	return branches
}

func CheckoutRemoteBranch(dir string, branchName string)  {
	cmd := "git checkout -b "+branchName+" origin/"+branchName
	Shell(dir, cmd)
}

func CheckoutLocalBranch(dir string, branchName string)  {
	cmd := "git checkout "+branchName
	Shell(dir, cmd)
}

func DeleteLocalBranch(dir string, branchName string)  {
	Shell(dir, "git branch -D "+branchName)
}

func DeleteRemoteBranch(dir string, branchName string)  {
	Shell(dir, "git push origin --delete "+branchName)
}

func FetchOriginStatus(dir string)  {
	Shell(dir, "git fetch --all --prune")
}

func NodesAll(dir string) []Commit {
	return Nodes(dir, "1971-01-01", "2099-01-01")
}

func Nodes(dir string, startDate string, endDate string) []Commit {
	res := Shell(dir, "git log --since="+startDate+" --until="+endDate)
	if res == "" {
		return []Commit{}
	}
	strCommits := strings.Split(res, "\ncommit")
	items := tool.ArrMap(func(s string) Commit {
		c := Commit{}
		s = strings.TrimPrefix(s, "commit")
		s = strings.Trim(s, " \n\t")
		strs := strings.Split(s, "\n")
		isMerge := strings.HasPrefix(strs[1], "Merge:")
		c.Id = strs[0]
		if isMerge {
			c.MergeIds = strings.Split(strings.TrimPrefix(strs[1], "Merge: "), " ")
			author := strings.Split(strings.Trim(strings.TrimPrefix(strs[2], "Author: "), "> "), "<")
			c.Author = strings.Trim(author[0], " ")
			c.Email = strings.Trim(author[1], " ")
			c.IsMerge = true
			c.Message = strings.Trim(strings.Join(strs[4:], "\n"), " \n\t")
			date := strings.Trim(strings.TrimPrefix(strs[3], "Date:"), " ")
			location, _ := time.LoadLocation("Asia/Shanghai")
			dateTime, _ := time.ParseInLocation("Mon Jan 02 15:04:05 2006 +0800", date, location)
			c.Date = dateTime
			c.IsRevert = strings.Contains(c.Message, "Revert") ||
				strings.Contains(c.Message, "撤销")||
				strings.Contains(c.Message, "revert")
		} else {
			c.IsMerge = false
			author := strings.Split(strings.Trim(strings.TrimPrefix(strs[1], "Author: "), "> "), "<")
			c.Author = strings.Trim(author[0], " ")
			c.Email = strings.Trim(author[1], " ")
			c.Message = strings.Trim(strings.Join(strs[5:], "\n"), " \n\t")
			date := strings.Trim(strings.TrimPrefix(strs[2], "Date:"), " ")
			location, _ := time.LoadLocation("Asia/Shanghai")
			dateTime, _ := time.ParseInLocation("Mon Jan 02 15:04:05 2006 +0800", date, location)
			c.Date = dateTime
			c.IsRevert = strings.Contains(c.Message, "Revert") ||
				strings.Contains(c.Message, "撤销")||
				strings.Contains(c.Message, "revert")
		}
		return c
	}, strCommits).([]Commit)

	return items
}

func IsMergedTo(dir string, branch string, toBranch string) bool {
	CheckoutRemoteBranch(dir, toBranch)
	CheckoutLocalBranch(dir, toBranch)
	AllNodes := NodesAll(dir)
	CheckoutRemoteBranch(dir, branch)
	CheckoutLocalBranch(dir, branch)
	okNodes := NodesAll(dir)
	CheckoutLocalBranch(dir, toBranch)
	for _, node := range AllNodes {
		for _, id := range node.MergeIds {
			if strings.HasPrefix(okNodes[0].Id, id) {
				return true
			}
		}
	}

	return false
}

func RemoteHeads(dir string) []Commit {
	strs := strings.Split(Shell(dir, "git ls-remote"), "\n")
	items := tool.ArrFilter(func( s string) bool {
		return strings.Contains(s, "refs/heads")
	}, strs).([]string)
	items1 := tool.ArrMap(func(s string) Commit {
		c := Commit{}
		s = strings.TrimSpace(s)
		i := strings.Split(s, "\t")
		c.Id = strings.TrimSpace(i[0])
		c.Branch = strings.TrimPrefix(strings.TrimSpace(i[1]), "refs/heads/")
		return c
	}, items).([]Commit)
	return items1
}

func ShowDetail(workDir string, hash string) Commit {
	r := Shell(workDir, "git show "+hash+" --stat")
	//fmt.Println(r)
	matched := regexp.MustCompile(`Merge: \w+ \w+`).FindAllString(r, -1)
	isMerge := len(matched) > 0

	matched1 := regexp.MustCompile(`(\d+) file[s]? changed`).FindAllStringSubmatch(r, -1)
	changedFiles := 0
	if len(matched1) > 0 && len(matched1[0]) == 2 {
		changedFiles = utils.JKStrToInt(matched1[0][1])
	}

	matched1 = regexp.MustCompile(`(\d+) insertion`).FindAllStringSubmatch(r, -1)
	insertLines := 0
	if len(matched1) > 0 && len(matched1[0]) == 2 {
		insertLines = utils.JKStrToInt(matched1[0][1])
	}

	matched1 = regexp.MustCompile(`(\d+) deletion`).FindAllStringSubmatch(r, -1)
	deleteLines := 0
	if len(matched1) > 0 && len(matched1[0]) == 2 {
		deleteLines = utils.JKStrToInt(matched1[0][1])
	}

	matched1 = regexp.MustCompile(`<(\S+)>`).FindAllStringSubmatch(r, -1)
	email := ""
	if len(matched1) > 0 && len(matched1[0]) == 2 {
		email = matched1[0][1]
	}

	matched1 = regexp.MustCompile(`(.+) <(\S+)>`).FindAllStringSubmatch(r, -1)
	author := "not found"
	if len(matched1) > 0 && len(matched1[0]) > 1 {
		author = strings.TrimSpace(strings.Split(matched1[0][1], ":")[1])
	}

	isRevert := strings.Contains(r, "撤销") || strings.Contains(r, "Revert")
	if strings.Contains(r, "合并分支") || strings.Contains(r, "Merge") {
		isMerge = true
	}

	info := Commit{
		Id: hash,
		AddLines:insertLines,
		RemoveLines:deleteLines,
		TotalLines:changedFiles,
		IsMerge:isMerge,
		Email:email,
		IsRevert:isRevert,
		//Detail:r,
		Author:author,
		WorkDir:workDir,
	}
	return info
}
