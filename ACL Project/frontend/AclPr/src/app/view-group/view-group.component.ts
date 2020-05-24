  import { Component, OnInit } from '@angular/core';
import { Http,Response,Headers } from '@angular/http';
import { Router } from '@angular/router';

@Component({
  selector: 'app-view-group',
  templateUrl: './view-group.component.html',
  styleUrls: ['./view-group.component.css']
})
export class ViewGroupComponent implements OnInit {

  constructor(private http:Http,private router:Router) { }

  groups : any = [];
  	fetchData = function()
  	{

    	this.http.get("/webapi/v1/groups").subscribe(
    	res=>
    {	
    	var array=res.json("data");
      console.log(array["data"]);
      var arr = array["data"];
      
      for(var i=0;i<arr.length;i++)
      {
        var obj={"id":arr[i][0],
        "group_name":arr[i][1],
        "origin_name":arr[i][2]};
      this.groups.push(obj);
      console.log(obj);
      }
      console.log(this.users);
    }
  )
}

//group_name:string;
a:any=[];
private headers=new Headers({'content-type':'application/json'});
deleteRecord=function(group_name)
{
  if(this.a[2]==1)
  {
	if(confirm("Are You sure?"))
    {
    this.http.delete(`${"/webapi/v1/groups/"}${group_name}`).subscribe(
    (data: {})=>
    {
    //this.fetchData();
    this.router.navigate(['/create-group']);
    //console.log(data)

    }
  )}
  }
  else
  {
    alert("user is not a admin");
    return false;
  }
  }
  
  ngOnInit(): void 
  {
    this.fetchData();
    this.a=JSON.parse(localStorage.getItem("info"));
  }

}
