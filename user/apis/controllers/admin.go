package admin

import (
        "html/template"
        "log"
	"net/http"
        "fmt"
        "time"
        "strconv"
        "user/config"
        "user/apis/params"
        "user/models"
        "user/untils"
        "github.com/mholt/binding"
	"github.com/julienschmidt/httprouter"
)

    func Login(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
            t, err := template.ParseFiles("../../view/login.html")
            if (err != nil) {
                log.Println(err)
            }
            t.Execute(rw, nil)
    }

    func Reg(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
            t, err := template.ParseFiles("../../view/reg.html")
            if (err != nil) {
                log.Println(err)
            }
            t.Execute(rw, nil)
    }

    func Index(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        cookie, err := r.Cookie("id")
        if err != nil || cookie.Value == ""{
            http.Redirect(rw, r, "/user/login", http.StatusFound)
        }

        infoRep := new(params.InfoRep)

        p := new(models.Data)
        _ = p.GetRowById(cookie.Value)
        if p.ID == 0 {
            b, _ := infoRep.Result(500, "用户不存在")
            fmt.Fprintf(rw, "%s", b)
            return
        }

        t, err := template.ParseFiles("../../view/index.html")
        if (err != nil) {
            log.Println(err)
        }
        t.Execute(rw, p)
    }

    func DoReg(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        regReq := new(params.RegReq)
        infoRep := new(params.InfoRep)

        if err := binding.Bind(r, regReq); err != nil {
            log.Printf("binding, error:%v\n", err)
            b, _ := infoRep.Result(500, "invalidate")
            fmt.Fprintf(rw, "%v", b)
            return
        }
        if regReq.Pass != regReq.Comfirm {
            b, _ := infoRep.Result(500, "两次密码不一致")
            fmt.Fprintf(rw, "%s", b)
            return
        }

        if utils.CheckMobile(regReq.Mobile) != true {
            b, _ := infoRep.Result(500, "手机号格式不对")
            fmt.Fprintf(rw, "%s", b)
            return
        }

        if utils.CheckEmail(regReq.Email) != true {
            b, _ := infoRep.Result(500, "邮箱格式不对")
            fmt.Fprintf(rw, "%s", b)
            return
        }
        in := new(models.Data)

        in.Name = regReq.Name
        in.Password = regReq.Pass
        in.Email = regReq.Email
        in.Mobile = regReq.Mobile
        in.Createtime = time.Now().Unix()

        res := models.Create(in);
        if res < 1 {
            b, _ := infoRep.Result(500,"添加失败")
            fmt.Fprintf(rw, "%s", b)
            return
        }
        b, _ := infoRep.Result(200,"添加成功")
        fmt.Fprintf(rw, "%s", b)
    }

    func DoLogin(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        loginReq := new(params.LoginReq)
        infoRep := new(params.InfoRep)

        if err := binding.Bind(r, loginReq); err != nil {
            log.Printf("binding, error:%s\n", err)
            b, _ := infoRep.Result(500, "invalidate")
            fmt.Fprintf(rw, "%s", b)
            return
        }
        p := new(models.Data)
        _ = p.GetOne(loginReq.Name)
        if p.ID == 0 {
            b, _ := infoRep.Result(500, "用户不存在")
            fmt.Fprintf(rw, "%s", b)
            return
        }

        if loginReq.Pass != p.Password {
             b, _ := infoRep.Result(500, "密码错误")
            fmt.Fprintf(rw, "%s", b)
            return
        }
        cookie := http.Cookie{Name: "id", Value: strconv.Itoa(p.ID), Path: "/"}
        http.SetCookie(rw, &cookie)
        err := config.RedisHandle.Set("ulogin:"+strconv.Itoa(p.ID), time.Now().Unix(),0).Err()
        if err != nil {
            log.Printf("Error info:%s\n", err)
            return
        }
         b, _ := infoRep.Result(200,"success")
        fmt.Fprintf(rw, "%s", b)
    }

    func DoUpdata(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        cookie, err := r.Cookie("id")
        if err != nil || cookie.Value == ""{
            http.Redirect(rw, r, "/user/login", http.StatusFound)
        }
        updataReq := new(params.UpdataReq)
        infoRep := new(params.InfoRep)

        if err := binding.Bind(r, updataReq); err != nil {
            log.Printf("binding, error:%v\n", err)
            b, _ := infoRep.Result(500, "invalidate")
            fmt.Fprintf(rw, "%v", b)
            return
        }
        p := new(models.Data)
        p.Name = updataReq.Name
        p.Email = updataReq.Email
        p.Mobile = updataReq.Mobile
        p.ID,_ = strconv.Atoi(cookie.Value)
        err = models.Update(p)
        if err != nil {
            log.Printf("Error info:%s\n", err)
            return
        }
         b, _ := infoRep.Result(200,"success")
        fmt.Fprintf(rw, "%s", b)
    }




