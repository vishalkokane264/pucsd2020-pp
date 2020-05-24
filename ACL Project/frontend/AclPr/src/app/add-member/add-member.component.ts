import { Component, OnInit } from '@angular/core';
import { Http,Response,Headers } from '@angular/http';
import { Router } from '@angular/router';
import { ActivatedRoute} from '@angular/router'

@Component({
  selector: 'app-add-member',
  templateUrl: './add-member.component.html',
  styleUrls: ['./add-member.component.css']
})
export class AddMemberComponent implements OnInit {

  constructor(private http:Http,private router:Router,private route:ActivatedRoute) { }

  id:number;
  userObj:object={};
  //userName:string;
  a:any=[];
  addUserId=function(grp)
  {
      if(this.a[2]==1)
  {
      this.http.get("/webapi/v1/groups").subscribe((res1:Response)=>
     {

       var a=res1.json();
        var list=a["data"];
        console.log(list);
        var flag=0;
        for(var j=0;j<list.length;j++)
        {
            if(list[j][0]==this.id)
            {
                console.log("hello");
                var par=list[j][1];
                flag=1;
                break;
            }
        }
        if(flag==1)
        {
        alert("username is does not exist");
        return false;
        }
        this.userObj=
        {
            "user_name":grp.username,
            "gr_name": par
        } 
        if(!grp.username)
        {
            alert("Plz enter User name");
            return false;
        }
        this.http.post("/webapi/v1/usergroup", this.userObj).subscribe((data)=>{
        //console.log(data);
        var array=data.json();
        console.log(array);
        //console.log(this.id);
        this.router.navigate(['/after-sign-in']);
        })
        })

    }
    else
    {
        alert("user is not admin");
        return false;
    }
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
   		//var users=arr["data"];
      console.log(users);
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
      //console.log(this.userName);
   })
   this.a=JSON.parse(localStorage.getItem("info"));
  }

}
