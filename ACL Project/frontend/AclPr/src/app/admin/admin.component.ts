import { Component, OnInit } from '@angular/core';
import { Http,Response,Headers } from '@angular/http';
import { Router } from '@angular/router';

@Component({
  selector: 'app-admin',
  templateUrl: './admin.component.html',
  styleUrls: ['./admin.component.css']
})
export class AdminComponent implements OnInit {

  constructor(private http:Http,private router:Router) { }
 
  /*userObj:object={};

  addnewUser=function()
  {
	this.userObj={
	"user_fname":userss.first_name,
	"user_lname" :userss.last_name,
  "username"  : userss.username,
	"password":userss.psw,
  "user_type":userss.select
  } 
	if(!userss.first_name)
    {
      alert("Plz enter First name");
      //form.inputfield.focus();
      return false;
    }
    if(!userss.last_name)
    {
      alert("Plz enter Last name");
      return false;
    }
    if(userss.select=="1")
    {
      alert("plz select valid option");
      return false;
    }

	this.http.post("/webapi/v1/user", this.userObj).subscribe((data)=>{
	console.log(data.json());
    //this.router.navigate(['/sign-in']);
	})
  }*/
  ngOnInit(): void {
  }

}
