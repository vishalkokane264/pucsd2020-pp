












import { Component, OnInit } from '@angular/core';
import { Http,Response,Headers } from '@angular/http';
import { Router } from '@angular/router';
import { ActivatedRoute} from '@angular/router';

@Component({
  selector: 'app-file-explorer',
  templateUrl: './file-explorer.component.html',
  styleUrls: ['./file-explorer.component.css']
})
export class FileExplorerComponent implements OnInit {
  textValue='initial value';
  constructor(private http:Http,private router:Router,private route:ActivatedRoute) { }
  id:number;
  fileOpen=function()
  {

  	this.http.get("/webapi/v1/fileserver").subscribe((data1)=>
    {
    	var arr=data1.json();
    	var array=arr["data"];
    	for(var i=0;i<array.length;i++)
    	{
    		if(this.id==array[i][0])
    		{
    			var obj={"file_path":array[i][1]};
    			var c=array[i][3]
    			console.log(obj);
    			break;
    		}
    	}
    })
  	this.http.post(`${"/webapi/v1/getfiledata/"}${this.id}`).subscribe((data)=>
  	{
  		console.log(data);
  	})	
  }

  ngOnInit(): void {
  this.route.params.subscribe(params=>
   {
   		this.id=+params['id'];
   });
   		this.http.get("/webapi/v1/userfilepermission").subscribe((res:Response)=>
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
   this.fileOpen();
  }

}
