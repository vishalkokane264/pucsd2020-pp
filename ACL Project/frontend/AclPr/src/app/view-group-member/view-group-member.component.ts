import { Component, OnInit } from '@angular/core';
import { Http,Response,Headers } from '@angular/http';
import { Router } from '@angular/router';
import { ActivatedRoute} from '@angular/router'

@Component({
  selector: 'app-view-group-member',
  templateUrl: './view-group-member.component.html',
  styleUrls: ['./view-group-member.component.css']
})
export class ViewGroupMemberComponent implements OnInit {

  constructor(private http:Http,private router:Router,private route:ActivatedRoute) { }

  users : any = [];
  	id:number;
    line:any=[];
    groupp :any={};
    a:any=[];
  	fetchData = function()
  	{
      this.http.get("/webapi/v1/groups").subscribe((res1:Response)=>
        {

          var a=res1.json();
          var list=a["data"];
          console.log(list);
          for(var j=0;j<list.length;j++)
          {
              if(list[j][0]==this.id)
              {
                  var par=list[j][1];
                  break;
              }
          }



    	   this.http.get(`${"/webapi/v1/usergroup/"}${par}`).subscribe(
    	   data=>
      {	
          var array=data.json();
          //console.log(array);
          var arr=array["data"];
          for(var i=0;i<arr.length;i++)
          {
              var obj={"user_name":arr[i][0],"gr_name":arr[i][1]};
              this.users.push(obj);
              //console.log(obj);
          }
          //console.log(this.users);
          })
      }
    )
    
}

  ngOnInit(): void 
  {
  		this.route.params.subscribe(params=>
   {
   		this.id=+params['id'];
   });
   		this.http.get("/webapi/v1/groups").subscribe((res:Response)=>
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
   this.fetchData();
   this.a=JSON.parse(localStorage.getItem("info"));
  }

}
