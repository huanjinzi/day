# git


//git tag 查看
git show release-1.5 --shortstat

// git 推送tags
git push origin --tags


//基于Tag创建分支
git checkout -b branch_name v1.0 

// git 删除远程分支 push tag 到远程仓库
git push origin :tag_name

// 删除tag 
git tag -d tag_name 

