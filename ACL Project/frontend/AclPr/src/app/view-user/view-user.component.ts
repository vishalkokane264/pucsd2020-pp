





import { Component, OnInit } from '@angular/core';
import { Http,Response,Headers } from '@angular/http';
import { Router } from '@angular/router';

//import 'rxjs/add/operator/toPromise';
//import { toPromise } from 'rxjs/operators/toPromise';
//import { Observable } from 'rxjs/Observable';

@Component({
  selector: 'app-view-user',
  templateUrl: './view-user.component.html',
  styleUrls: ['./view-user.component.css']
})
export class ViewUserComponent implements OnInit {
  userObj:object={};

  constructor(private http:Http,private router:Router) { }
  users : any = [];
    fetchData = function()
    {
      this.http.get("/webapi/v1/user").subscribe(
      res=>
    {
      var array=res.json("data");
      console.log(array["data"]);
      var arr = array["data"];
      
      for(var i=0;i<arr.length;i++)
      {
        var obj={"id":arr[i][0],
        "user_fname":arr[i][1],
        "user_lname":arr[i][2],
        "username":arr[i][3]};
      this.users.push(obj);
      console.log(obj);
      }
      console.log(this.users);
    }
  )
}

a:any=[];
private headers=new Headers({'content-type':'application/json'});
deleteRecord=function(group_name)
{
  if(this.a[2]==1)
  {
  if(confirm("Are You sure?"))
    {
    this.http.delete(`${"/webapi/v1/user/"}${group_name}`).subscribe(
    (data: {})=>
    {
    //this.fetchData();
    //this.deletePermission(group_name);
    //console.log(data)
    this.router.navigate(['/after-sign-in']);

    })
  
  }
    


  }
  else
  {
    alert("user is not a admin");
    return false;
  }
  }

  deletePermission=function(group_name)
  {
      this.http.delete(`${"/webapi/v1/userpermission/"}${group_name}`).subscribe(
      (data: {})=>
      {
      //this.fetchData();
      this.router.navigate(['/view-user']);
      //console.log(data)
  
      })
    }
  

  ngOnInit(): void {
  this.fetchData();
  this.a=JSON.parse(localStorage.getItem("info"));
  }

}