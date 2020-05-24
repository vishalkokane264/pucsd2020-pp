import { Component, OnInit } from '@angular/core';
import { Http,Response,Headers } from '@angular/http';
import { Router } from '@angular/router';

@Component({
  selector: 'app-group-grant',
  templateUrl: './group-grant.component.html',
  styleUrls: ['./group-grant.component.css']
})
export class GroupGrantComponent implements OnInit {

  constructor(private http:Http,private router:Router) { }

  groupObj:object={};
  a:any=[];
  grantGroupDetails=function(grant)
  {
      if(this.a[2]==1)
      {
        this.groupObj=
        {
            "gr_name":grant.group_name,
            "file_id" :parseInt(grant.resource_id),
            "given_by":this.a[1],
            "perm_type": parseInt(grant.select)
        } 
        if(!grant.group_name)
        {
            alert("Plz enter group id");
            return false;
        }
        if(!grant.resource_id)
        {
            alert("Plz enter resource id");
            return false;
        }
        if(grant.select=="1")
        {
            alert("Plz select Drop Down Option");
            return false;
        }
    	//console.log(grant.select);
        this.http.post("/webapi/v1/groupfilepermission", this.groupObj).subscribe((data)=>{
        //console.log(data);
        var array=data.json();
        var arr=array["message"];
        if(arr=="Success")
        {
            this.router.navigate(['/manage-permission']);
        }
        else
        {
              alert("check group is exit or not or file id is valid or not");
              return false;
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
  }

}
