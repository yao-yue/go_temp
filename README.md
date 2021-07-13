## 把本地项目推到自己的github
git remote add origin git@github.com:yao-yue/go_temp.git<br/> 
git branch -M main <br/> 
git push -u origin main <br/> 
> 关于-M的解释  --move --force  <br/> 
> --move  Move/rename a branch and the corresponding reflog. <br/> 
> --force  allow renaming the branch even if the new branch name already exists <br/> 
> 理解就是强行新建一个分支main，即使已经有了 <br/> 

>关于-u的解释 <br/> 
> -u <br/> 
>--set-upstream  的简写 <br/> 
>For every branch that is up to date or successfully pushed, add upstream  (tracking) reference, used by argument-less git-pull[1] and other commands. For  more information, see branch.<name>.merge in git-config[1]. 


## 创建一个新的仓库并关联 
echo "# go_temp" >> README.md  <br/> 
git init  <br/> 
git add README.md  <br/> 
git commit -m "first commit"  <br/> 
git branch -M main  <br/> 
git remote add origin git@github.com:yao-yue/go_temp.git  <br/> 
git push -u origin main  <br/> 


## 参考链接
[git文档](https://git-scm.com/docs/git-push) 