import { Component, OnInit } from '@angular/core';
import { Http,Response,Headers } from '@angular/http';
import { Router } from '@angular/router';
import { ActivatedRoute} from '@angular/router';


@Component({
  selector: 'app-user-grant',
  templateUrl: './user-grant.component.html',
  styleUrls: ['./user-grant.component.css']
})
export class UserGrantComponent implements OnInit {

  constructor(private http:Http,private router:Router,private route:ActivatedRoute) { }

  grantObj:object={};
  a:any=[];
  grantUserDetails=function(grant)
  {
      if(this.a[2]==1)
      {
        this.grantObj=
        {
            "username":grant.user_name,
            "file_id" :parseInt(grant.resource_id),
            "given_by":this.a[1],
            "perm_type": parseInt(grant.select)
        } 
        if(!grant.user_name)
        {
            alert("Plz enter user name");
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
        this.http.post("/webapi/v1/userfilepermission", this.grantObj).subscribe((data)=>{
        //console.log(this.grantObj);
        var array=data.json();
        var arr=array["message"];
        if(arr=="Success")
        {
            this.router.navigate(['/manage-permission']);
        }
        else
        {
            alert("check user is exit or not or file_id is valid or not");
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
