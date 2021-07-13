## 把本地项目推到自己的github
git remote add origin git@github.com:yao-yue/go_temp.git 
git branch -M main 
git push -u origin main 
> 关于-M的解释  --move --force  
> --move  Move/rename a branch and the corresponding reflog. 
> --force  allow renaming the branch even if the new branch name already exists 
> 理解就是强行新建一个分支main，即使已经有了 

>关于-u的解释 
> -u 
>--set-upstream  的简写 
>For every branch that is up to date or successfully pushed, add upstream  (tracking) reference, used by argument-less git-pull[1] and other commands. For  more information, see branch.<name>.merge in git-config[1]. 


## 创建一个新的仓库并关联 
echo "# go_temp" >> README.md 
git init 
git add README.md 
git commit -m "first commit" 
git branch -M main 
git remote add origin git@github.com:yao-yue/go_temp.git 
git push -u origin main 


## 参考链接
[git文档](https://git-scm.com/docs/git-push) 