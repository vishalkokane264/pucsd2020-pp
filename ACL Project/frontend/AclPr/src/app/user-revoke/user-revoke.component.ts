import { Component, OnInit } from '@angular/core';
import { Http,Response,Headers } from '@angular/http';
import { Router } from '@angular/router';
import { ActivatedRoute} from '@angular/router';

@Component({
  selector: 'app-user-revoke',
  templateUrl: './user-revoke.component.html',
  styleUrls: ['./user-revoke.component.css']
})
export class UserRevokeComponent implements OnInit {

  constructor(private http:Http,private router:Router,private route:ActivatedRoute) { }
obj1=[];
  ress : any = [];
    hi:number;
    a:any=[];
    fetchData = function()
    {

      this.http.get("/webapi/v1/userfilepermission").subscribe(
      res=>
    { 
      
      
      var array=res.json("data");
      console.log(array["data"]);
      var arr = array["data"];
      
      for(var i=0;i<arr.length;i++)
      {
        var obj=
        {
          "id":arr[i][0],
          "username":arr[i][1],
          "file_id":arr[i][2],
          "given_by":arr[i][3],
          "perm_type":arr[i][4]};
      this.ress.push(obj);
      console.log(obj);
      }
      console.log(this.ress);
    /*  this.http.get("http://localhost:5000/addUser").subscribe(
      data=>
    { 
      this.http.get("http://localhost:9999/userResource").subscribe(
      res1=>
    { 

      var array=res.json();
      console.log(array);
      var arr=data.json();
      console.log(arr);
      var list=res1.json();
      console.log(list);
      for(var i=0;i<list.length;i++)
      {
        for(var j=0;j<arr.length;j++)
        {
          
            if(list[i].user_id==arr[j].id)
            {
              for(var k=0;k<array.length;k++)
              { 
                  if(list[i].resource_id==array[k].id)
                  { 
                    this.ress.push({"id":list[i].user_id,"first_name":arr[j].first_name,"resource_Path":array[k].resource_Path});
                  }
              }
            }
        }
      }
    })
    })*/
    }
  )
}



id:number;
private headers=new Headers({'content-type':'application/json'});
deleteRecord=function(id)
{
    if(confirm("Are You sure?"))
    {
    if(this.a[2]==1)
    {
    this.http.delete(`${"/webapi/v1/userfilepermission/"}${id}`).subscribe(
    (data: {})=>
    {
    // this.fetchData();
     this.router.navigate(['/revoke']);
    console.log(data)
    }
  )
  }
  else
  {
      alert("user is not a admin");
      return false;
  }

  }
}

  ngOnInit(): void {
   this.fetchData();
   this.a=JSON.parse(localStorage.getItem("info"));
  }

}
