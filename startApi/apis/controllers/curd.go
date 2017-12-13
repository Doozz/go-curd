package controllers

import (
	"fmt"
        //"encoding/json"
	"net/http"
        "log"
        "startApi/apis/params"
        "startApi/models"
        "github.com/mholt/binding"
	"github.com/julienschmidt/httprouter"
)

    func Create(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        infoReq := new(params.InfoReq)
        infoRep := new(params.InfoRep)
        //解析参数，默认是不会解析的
	if err := binding.Bind(r, infoReq); err != nil {
            log.Printf("binding, error:%v\n", err)
            b, _ := infoRep.Result(500, "invalidate")
            fmt.Fprintf(rw, "%v", b)
            return
        }
        p := new(models.Data)

        p.Name = infoReq.Name
        p.Age = infoReq.Age
        p.From = infoReq.From
        p.Sex = infoReq.Sex
        p.Memo = infoReq.Memo

        res := models.Create(p)
        if res < 1 {
            b, _ := infoRep.Result(500,"添加失败")
            fmt.Fprintf(rw, "%s", b)
            return
        }
        b, _ := infoRep.Result(200,"添加成功")
        fmt.Fprintf(rw, "%s", b)
    }

    func Update(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        infoReq := new(params.InfoReq)
        infoRep := new(params.InfoRep)

        if err := binding.Bind(r, infoReq); err != nil {
            log.Printf("binding, error:%v\n", err)
            b, _ := infoRep.Result(500, "invalidate")
            fmt.Fprintf(rw, "%s", b)
            return
        }

        p := new(models.Data)

        p.Name = infoReq.Name
        p.Age = infoReq.Age
        p.From = infoReq.From
        p.Sex = infoReq.Sex
        p.Memo = infoReq.Memo

        if infoReq.Id <= 0 {
            b, _ := infoRep.Result(500,"invalidate")
            fmt.Fprintf(rw, "%s", b)
            return
        }
         p.ID = infoReq.Id

        res := models.Create(p)
        if res < 0 {
            b, _ := infoRep.Result(500,"修改失败")
            fmt.Fprintf(rw, "%s", b)
            return
        }
        b, _ := infoRep.Result(200,"修改成功")
        fmt.Fprintf(rw, "%s", b)

    }

    func Query(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        infoRep := new(params.InfoRep)
        p := new(models.Data)
        res,_ := p.GetAll()
        da,_:=infoRep.Result(200, res)
        fmt.Fprintf(rw, "%s", da)
    }

    func QueryRow(rw http.ResponseWriter, r *http.Request, ps httprouter.Params ) {
        infoReq := new(params.IDReq)
        infoRep := new(params.InfoRep)
        if err := binding.Bind(r, infoReq); err != nil {
            log.Printf("binding, error:%v\n", err)
            b, _ := infoRep.Result(500, "invalidate")
            fmt.Fprintf(rw, "%s", b)
            return
        }
        p := new(models.Data)
        _ = p.GetOne(infoReq.Id)
        if p.ID == 0 {
            b, _ := infoRep.Result(500, "用户不存在")
            fmt.Fprintf(rw, "%s", b)
            return
        }
        da,_:=infoRep.Result(200, p)
        fmt.Fprintf(rw, "%s", da)

    }

    func Delete(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        infoReq := new(params.IDReq)
        infoRep := new(params.InfoRep)

        if err := binding.Bind(r, infoReq); err != nil {
            log.Printf("binding, error:%v\n", err)
            b, _ := infoRep.Result(500, "invalidate")
            fmt.Fprintf(rw, "%v", b)
            return
        }

        p := new(models.Data)

        p.ID = infoReq.Id
        res := models.Delete(p)
        if res < 1 {
            b, _ := infoRep.Result(500,"删除失败")
            fmt.Fprintf(rw, "%s", b)
            return
        }
        b, _ := infoRep.Result(200,"删除成功")
        fmt.Fprintf(rw, "%s", b)

    }
