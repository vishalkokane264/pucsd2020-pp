import { Component, OnInit } from '@angular/core';
import { Http,Response,Headers } from '@angular/http';
import { Router } from '@angular/router';

@Component({
  selector: 'app-add-user',
  templateUrl: './add-user.component.html',
  styleUrls: ['./add-user.component.css']
})
export class AddUserComponent implements OnInit {
  userObj:object={};
  userType:object={};
  constructor(private http:Http,private router:Router) { }
  a:any=[];
  addUserDetails=function(userss)
  { 
      if(this.a[2]==1)
      {
        this.userObj=
        {
            "user_fname":userss.first_name,
            "user_lname" :userss.last_name,
            "username" : userss.username,
            "password":userss.psw
        }
        this.userType=
        {
            "username" : userss.username,
            "type_id"  : parseInt(userss.select)
        } 
        if(!userss.first_name)
        {
            alert("Plz enter First name");
            return false;
        }
        if(!userss.last_name)
        {
            alert("Plz enter Last name");
            return false;
        }
        if(userss.select=="0")
        {
            alert("plz select valid option");
            return false;
        }

        this.http.post("/webapi/v1/user", this.userObj).subscribe((data)=>
        {
          this.http.post("/webapi/v1/userpermission", this.userType).subscribe((res)=>
        {
            var array=data.json();
            console.log(array);
            var ar=array["message"];
            var list=res.json();
            var arr=list["message"];
            console.log(list);
            if(ar=="Success" && arr=="Success")
            {
                alert("added Successfully");
            }
            this.router.navigate(['/after-sign-in']);
      })
      })
      }
      else
      {
          alert("user is not a admin");
          return false;
      }
  }		

  ngOnInit(): void {
  this.a=JSON.parse(localStorage.getItem("info"));
  }

}
