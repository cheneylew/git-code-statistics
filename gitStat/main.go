package main

import (
	"flag"
	"fmt"
	"github.com/cheneylew/git-code-statistics/gitStat/git"
	"github.com/cheneylew/goutil/utils"
	"regexp"
	"sort"
	"strings"
)

func main() {
	//cmd()
	calcRateOfContribution2()
}

func calcRateOfContribution2()  {
	startDate := "2021-04-01"
	endDate := "2021-07-01"
	var users []User
	//git.Shell("/Users/apple/Desktop/ehsy/opc", "git checkout master")
	users = append(users, start2("/Users/apple/Desktop/ehsy/opc", startDate, endDate)...)
	users = append(users, start2("/Users/apple/Desktop/ehsy/eis", startDate, endDate)...)
	users = append(users, start2("/Users/apple/Desktop/ehsy/uniapp", startDate, endDate)...)
	users = append(users, start2("/Users/apple/Desktop/ehsy/uni-spc", startDate, endDate)...)
	users = append(users, start2("/Users/apple/Desktop/ehsy/crm", startDate, endDate)...)
	users = append(users, start2("/Users/apple/Desktop/ehsy/raxwell-front-html", startDate, endDate)...)
	users = append(users, start2("/Users/apple/Desktop/ehsy/raxwell-eis", startDate, endDate)...)
	users = append(users, start2("/Users/apple/Desktop/ehsy/raxwell-opc", startDate, endDate)...)
	users = append(users, start2("/Users/apple/go/src/gitlab.ehsy.com/ehsy/frontend/go-service", startDate, endDate)...)
	users = append(users, start2("/Users/apple/Desktop/ehsy/flutter-warehouse", startDate, endDate)...)
	users = append(users, start2("/Users/apple/Desktop/ehsy/sso-api", startDate, endDate)...)
	users = append(users, start2("/Users/apple/Desktop/ehsy/questionare", startDate, endDate)...)
	merge(users)
}

func cmd()  {
	var dir        string
	var startDate    string
	var endDate        string
	flag.StringVar(&dir, "d", "", "git仓库所在目录")
	flag.StringVar(&startDate, "s", "", "开始日期(0000-00-00)")
	flag.StringVar(&endDate, "e", "", "结束日期(0000-00-00)")
	flag.Parse()
	if dir == "" || startDate == "" || endDate == "" {
		fmt.Println("参数输入错误！格式如下:")
		fmt.Println("./gitStat -d /path/to/project -s 2020-01-01 -e 2020-04-01")
		fmt.Println("参数含义:")
		flag.Usage()
		return
	}


	results := start(dir, startDate, endDate)
	merge(results)
}

type User struct {
	Name string
	Email string
	GitAddLines int
	GitRemoveLines int
	GitTotalLines int
	Nodes []git.Commit
}

func regexpGetNum(s string) int64 {
	r := regexp.MustCompile(`[\d-]+`).FindAllString(s, -1)
	if len(r) > 0{
		return utils.JKStrToInt64(r[0])
	}
	return 0
}

func showList(workDir string, author string, startDate string, endDate string) []string {
	author = strings.Replace(author, " ", "\\ ", -1)
	r := git.Shell(workDir, "git log --author='"+author+"' --since="+startDate+" --until="+endDate+"")
	ids := regexp.MustCompile(`[0-9a-z]{40}`).FindAllString(r, -1)
	return ids
}

func getUserStatics(workDir string, author string, email string, startDate string, endDate string) User {
	hashIds := showList(workDir, fmt.Sprintf("%s <%s",author,utils.TrimChars(email, "\t")), startDate, endDate)
	var user User
	var nodes []git.Commit
	for _, id := range hashIds {
		info := git.ShowDetail(workDir, id)
		if !info.IsMerge && !info.IsRevert {
			user.GitAddLines += info.AddLines
			user.GitRemoveLines += info.RemoveLines
			user.Email = info.Email
			info.WorkDir = workDir
			nodes = append(nodes, info)
		}
	}
	user.Name = author
	user.GitTotalLines = user.GitAddLines-user.GitRemoveLines
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].AddLines > nodes[j].AddLines
	})
	user.Nodes = nodes

	return user
}

func start(workDir string, startDate string, endDate string) []User {
	//workDir := "/Users/dejunliu/Desktop/ehys/uniapp"
	//info1 := showDetail(workDir, "4ed5c2e1d148e278a2d95c7a99d13f2576853ffe")
	//fmt.Println(info1)
	//return

	script := `git log --format='%aN <%ae' | sort -u | while read name; do echo -en "$name\t"; git log --author="$name" --since=2020-01-01 --until=2020-02-01 --pretty=tformat: --numstat | awk '{ add += $1; subs += $2; loc += $1 - $2 } END { printf ",added lines: %s, removed lines: %s, total lines: %s\n", add, subs, loc }' -; done`
	script = strings.Replace(script, "2020-01-01", startDate, -1)
	script = strings.Replace(script, "2020-02-01", endDate, -1)
	shellRes := git.Shell(workDir, script)
	//fmt.Println(git.git.ShellResult)
	slices := strings.Split(shellRes, "-en")
	var users []User
	for _, value := range slices {
		row := strings.Replace(utils.Trim(value), "\n", "", -1)
		items := strings.Split(row, ",")
		if len(items) > 1 {
			t := strings.Split(utils.TrimChars(items[0]," \t"), "<")
			author := utils.TrimChars(t[0], " \t")
			email := utils.TrimChars(t[1], " \t")
			var user User
			user.Name = author
			user.Email = email
			user.GitAddLines = int(regexpGetNum(items[1]))
			user.GitRemoveLines = int(regexpGetNum(items[2]))
			user.GitTotalLines = int(regexpGetNum(items[3]))
			if user.GitAddLines !=0 || user.GitRemoveLines != 0  {
				if user.Name != "" {
					users = append(users, user)
					//fmt.Println(user)
				}
			}
		}
	}

	//去除重复
	var ftUsers []User
	for _, user := range users {
		exist := false
		for ftKey, ftUser := range ftUsers {
			if ftUser.Name == user.Name {
				if strings.Contains(ftUser.Email, user.Email) {
					ftUsers[ftKey].Email = user.Email
					exist = true
				} else if strings.Contains(user.Email, ftUser.Email) {
					exist = true
				}
			}

		}
		if !exist {
			ftUsers = append(ftUsers, user)
		}
	}

	var results []User
	for _, user := range ftUsers {
		//fmt.Println(user)
		//fmt.Println("=======")
		tuser := getUserStatics(workDir, user.Name, user.Email, startDate, endDate)
		if tuser.GitAddLines == 0 && tuser.GitRemoveLines == 0 {
			continue
		}
		results = append(results, tuser)
	}

	return results
}

func start2(workDir string, startDate string, endDate string) []User {
	var users []User
	coms := git.Nodes(workDir, startDate, endDate)
	var infos []git.Commit
	for _, i2 := range coms {
		info := git.ShowDetail(workDir, i2.Id)
		info.Author = strings.Trim(info.Author, " \t")
		info.Email = strings.Trim(info.Email, " \t")
		infos = append(infos, info)
	}
	for _, info := range infos {
		uIndex := -1
		for idx, user := range users {
			if info.Author == user.Name && info.Email == user.Email {
				uIndex = idx
			}
		}
		if uIndex != -1{
			if info.IsMerge  || info.IsRevert  {
				continue
			} else {
				users[uIndex].Nodes = append(users[uIndex].Nodes, info)
				users[uIndex].GitAddLines += info.AddLines
				users[uIndex].GitRemoveLines += info.RemoveLines
				users[uIndex].GitTotalLines += (info.AddLines-info.RemoveLines)
			}
		} else {
			u := User{
				Name:           info.Author,
				Email:          info.Email,
				GitAddLines:    info.AddLines,
				GitRemoveLines: info.RemoveLines,
				GitTotalLines:  info.AddLines-info.RemoveLines,
				Nodes:          []git.Commit{info},
			}
			users = append(users, u)
		}
	}

	return users
}

func merge(users []User) {
	sort.Slice(users, func(i, j int) bool {
		return users[i].GitAddLines > users[j].GitAddLines
	})

	sameMap := make(map[string]string, 0)
	sameMap["color"] = "郑小丹"
	sameMap["Alan_liu"] = "刘德军"
	sameMap["sam_han"] = "韩庆满"
	sameMap["eric_he"] = "贺键"
	sameMap["tomas"] = "徐建鹏"
	sameMap["dan_han"] = "韩丹"
	sameMap["Claris Lu"] = "Claris_lu"
	//合并相同用户行数
	var results []User
	for _, user := range users {
		flag := false
		for key, rUser := range results {
			isSame := false
			for k, v := range sameMap {
				if (strings.TrimSpace(rUser.Name) == k && strings.TrimSpace(user.Name) == v) ||
					(strings.TrimSpace(rUser.Name) == v && strings.TrimSpace(user.Name) == k) {
					isSame = true
				}
			}
			if isSame ||
				(strings.TrimSpace(user.Name) == strings.TrimSpace(rUser.Name) ||
					strings.TrimSpace(user.Name) == strings.TrimSpace(rUser.Email) ||
					strings.TrimSpace(user.Email) == strings.TrimSpace(rUser.Name)) {
				flag = true
				results[key].GitTotalLines += user.GitTotalLines
				results[key].GitRemoveLines += user.GitRemoveLines
				results[key].GitAddLines += user.GitAddLines
				for _, node := range user.Nodes {
					results[key].Nodes = append(results[key].Nodes, node)
				}
			}
		}
		if !flag {
			results = append(results, user)
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].GitAddLines > results[j].GitAddLines
	})

	for _, tuser := range results {
		t := "\t\t"
		if len(tuser.Name) > 5 {
			t = "\t"
		}
		if strings.Contains(tuser.Name, "刘") ||
			strings.Contains(tuser.Name, "徐") ||
			strings.Contains(tuser.Name, "韩") {
			t = "\t\t"
		}
		fmt.Println(fmt.Sprintf("用户名：%s %s添加行数:%d \t\t删除行数:%d \t\t有效行数:%d  \t%s", tuser.Name , t, tuser.GitAddLines, tuser.GitRemoveLines, tuser.GitTotalLines, utils.TrimChars(tuser.Email, " \t")))
	}

	fmt.Println("提交或删除超过500行的异常节点:")
	for _, result := range results {
		for _, node := range result.Nodes {
			//fmt.Println(node.CommintID)
			if node.AddLines > 500 || node.RemoveLines > 500 {
				fmt.Println(fmt.Sprintf("%s \t %d \t %d \t %s \t %s \t %s", node.Id, node.AddLines, node.RemoveLines, node.Email, node.Author, node.WorkDir))
			}
		}
	}
}
