






import { Component, OnInit } from '@angular/core';
import { Http,Response,Headers } from '@angular/http';
import { Router } from '@angular/router';

@Component({
  selector: 'app-res-view',
  templateUrl: './res-view.component.html',
  styleUrls: ['./res-view.component.css']
})
export class ResViewComponent implements OnInit {

  constructor(private http:Http,private router:Router) { }

  resource:any=[];
  a:any=[];
  //b:3;
  fetchData=function()
  {
    this.http.get("/webapi/v1/fileserver").subscribe((data)=>{
       var array=data.json();
       var arr=array["data"];
       if(this.a[2]==1)
       {
          for(var i=0;i<arr.length;i++)
          {
              var obj={
              "id":arr[i][0],
              "file_name":arr[i][1],
              "owner_name":arr[i][2],
              "perm_type":3
              }
              this.resource.push(obj);
          }
       }
       else
       {
          this.http.get("/webapi/v1/userfilepermission").subscribe((data1)=>
          {
              var src=data1.json();
              var list=src["data"];
              console.log(list);
              var cnt=0;
              for(var j=0;j<list.length;j++)
              {
                  if(this.a[1]==list[j][0])
                  {
                      for(var k=0;k<arr.length;k++)
                      {
                          if(list[j][1]==arr[k][0])
                          {
                              var obj=
                              {
                                  "id":list[j][1],
                                  "file_name":arr[k][1],
                                  "owner_name":list[j][2],
                                  "perm_type":list[j][3],
                              }
                              this.resource.push(obj);
                              cnt++;
                          }
                      }
                  }
              }
              if(cnt==0)
              {
                  alert("user do not have any permission");
                  this.router.navigate(['/manage-permission']);
                  return false;
              }
          })
       }
    })
  }

  ngOnInit(): void {
    this.fetchData(); 
    this.a=JSON.parse(localStorage.getItem("info"));
    }

}
