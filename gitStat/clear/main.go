package clear

import (
	"fmt"
	"github.com/cheneylew/projects/gitStat/git"
	"strings"
)

func main() {
	clearBranches()
}

func clearBranches()  {
	dir := "/Users/apple/Desktop/ehsy/opc"
	git.CheckoutLocalBranch(dir, "master")
	nodes := git.NodesAll(dir)
	git.CheckoutLocalBranch(dir, "master-eis")
	nodes = append(nodes, git.NodesAll(dir)...)
	git.CheckoutLocalBranch(dir, "master-raxwell-eis")
	nodes = append(nodes, git.NodesAll(dir)...)
	git.CheckoutLocalBranch(dir, "master-raxwell-opc")
	nodes = append(nodes, git.NodesAll(dir)...)
	heads := git.RemoteHeads(dir)
	for i, c := range heads {
		for _, node := range nodes {
			for _, id := range node.MergeIds {
				if strings.HasPrefix(c.Id, id) {
					//已合并
					git.DeleteRemoteBranch(dir, c.Branch)
					fmt.Println("删除 ", c.Branch, " 成功!")
				}
			}
		}
		fmt.Println(fmt.Sprintf("已完成%v", float64(i+1)/float64(len(heads))))
	}
	git.FetchOriginStatus(dir)
}