import { Component, OnInit } from '@angular/core';
import { Http,Response,Headers } from '@angular/http';
import { Router } from '@angular/router';

@Component({
  selector: 'app-create-group',
  templateUrl: './create-group.component.html',
  styleUrls: ['./create-group.component.css']
})
export class CreateGroupComponent implements OnInit {

  constructor(private http:Http,private router:Router) { }

  groupObj:object={};
  a:any=[];
  addGroupName=function(grp)
  {
      if(this.a[2]==1)
      {
  		  this.groupObj=
  		{
  			"group_name":grp.group_name,
        "group_owner":this.a[1]
  		}
  		if(!grp.group_name)
        {
            alert("Plz enter group name");
            return false;
        }
        this.http.post("/webapi/v1/groups", this.groupObj).subscribe((data)=>{
        var array=data.json();
        console.log(array);
        var arr=array["message"];
        console.log(arr);
        if(arr=="Success")
        {
            alert("group created");
            this.router.navigate(['/view-group']);
        }
        else
        {
            alert("something went worng")
        }
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
  console.log(this.a);
  }

}
