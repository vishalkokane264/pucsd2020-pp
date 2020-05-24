import { Component, OnInit } from '@angular/core';
import { Http,Response,Headers } from '@angular/http';
import { Router } from '@angular/router';

@Component({
  selector: 'app-add-resource',
  templateUrl: './add-resource.component.html',
  styleUrls: ['./add-resource.component.css']
})
export class AddResourceComponent implements OnInit {

  constructor(private http:Http,private router:Router) { }

   userObj:object={};
   a:any=[];
  addResourcePath=function(per)
  {

    if(this.a[2]==1)
    {
        this.userObj=
        {
            "file_path":per.file_path,
            "owner_name":this.a[1]
        } 
        if(!per.file_path)
        {
            alert("Plz enter Path");
            return false;
        }
        this.http.post("/webapi/v1/fileserver", this.userObj).subscribe((data)=>{
        console.log(data.json());
        //var array=data.json();
        //console.log(array);
        //console.log(this.id);
        this.router.navigate(['/manage-permission']);
        this.http.get("/webapi/v1/fileserver").subscribe((data1)=>{

          var arr=data1.json();
          console.log(arr);
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
