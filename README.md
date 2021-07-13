## 把本地项目推到自己的github
git remote add origin git@github.com:yao-yue/go_temp.git
git branch -M main
git push -u origin main


## 创建一个新的仓库并关联
echo "# go_temp" >> README.md
git init
git add README.md
git commit -m "first commit"
git branch -M main
git remote add origin git@github.com:yao-yue/go_temp.git
git push -u origin main