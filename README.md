# git & go

- [git & go](#git--go)
  - [把本地项目推到自己的github](#把本地项目推到自己的github)
  - [创建一个新的仓库并关联](#创建一个新的仓库并关联)
  - [关于go的模块化](#关于go的模块化)
    - [清理包的缓存](#清理包的缓存)
  - [参考链接](#参考链接)

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


## 关于go的模块化
其实可以把他们和node的模块化结合起来学习
简单总结：  
1. 引入本地的包  
	"example.com/user/hello/morestrings"
2. 引入外部库的包
    "github.com/google/go-cmp/cmp"
    使用go mod tidy 进行拉取依赖
    ![](http://ww1.sinaimg.cn/large/006x4mSygy1gsf36ygagvj313c0wyaee.jpg)
3. 可以把自己写的库上传到github,然后本地可以用名字引入，然后再go mod tidy

### 清理包的缓存
you can pass the -modcache flag to go clean:
go clean -modcache


## 参考链接
[git文档](https://git-scm.com/docs/git-push) 
[go module化 参考文档](https://golang.org/doc/code)