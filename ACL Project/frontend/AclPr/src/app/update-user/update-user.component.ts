





import { Component, OnInit } from '@angular/core';
import { Http,Response,Headers } from '@angular/http';
import { Router } from '@angular/router';
import { ActivatedRoute} from '@angular/router';

@Component({
  selector: 'app-update-user',
  templateUrl: './update-user.component.html',
  styleUrls: ['./update-user.component.css']
})
export class UpdateUserComponent implements OnInit {

  constructor(private http :Http,private router:Router,private route:ActivatedRoute) { }



  id:number;
  updateObj : object={};
  a:any=[];
  updateUserDetails = function(update)
  {
    if(this.a[2]==1)
    {
    this.updateObj={
      "user_fname":update.first_name,
      "user_lname":update.last_name,
      "username"  :update.username,
      "password":update.psw
    }
    this.http.put(`${"/webapi/v1/user/"}${this.id}`,this.updateObj,JSON.stringify(this.updateObj)).subscribe(
    res=>{console.log(res);
    var arr=res.json();
    console.log(arr);
    var a=arr["ok"];
    this.router.navigate(['/view-user']);}
  )}
  else
  {
      alert("user is not a admin");
      return false;
  }
}

  ngOnInit(): void 
  {
   		this.route.params.subscribe(params=>
   {
   		this.id=+params['id'];
   });
   		this.http.get("/webapi/v1/user").subscribe((res:Response)=>
   {
   		var users=res.json();
   		var exist=false;
   		for(var i=0;i<users.length;i++)
   		{
   			if(parseInt(users[i].id)==this.id)
   			{
   				exist=true;
   				var data=users[i];
   				break;
   			}
   		}
   })
   this.a=JSON.parse(localStorage.getItem("info"));
  }

}
