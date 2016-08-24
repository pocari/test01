# go test

## install

```
PROJECT_DIR="project_dir_path"
mkdir $PROJECT_DIR
cd $PROJECT_DIR
gvm pkgset create --local
gvm pkgset use --local
git clone git@github.com:pocari/test01.git src/github.com/pocari/test01
go get github.com/Masterminds/glide
go install github.com/Masterminds/glide
cd src/github.com/pocari/test01
glide up
```
