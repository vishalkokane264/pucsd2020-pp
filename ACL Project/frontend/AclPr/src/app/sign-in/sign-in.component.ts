import { Component, OnInit } from '@angular/core';
import { Http,Response,Headers } from '@angular/http';
import { Router } from '@angular/router';

@Component({
  selector: 'app-sign-in',
  templateUrl: './sign-in.component.html',
  styleUrls: ['./sign-in.component.css']
})
export class SignInComponent implements OnInit {
  usersObj:object={};
  
  constructor(private http:Http,private router:Router) { }
  
  user : any = [];
  hi:string;
  signUpUser = function(usersss){
  this.usersObj=
  {
  		"username":usersss.username,
  		"password":usersss.psw
  }
    if(!usersss.username)
    {
      alert("Plz enter username");
      return false;
    }
    if(!usersss.psw)
    {
      alert("Plz enter password");
      return false;
    }
    this.http.get("/webapi/v1/user").subscribe(
    res=>
    {
      var data=res.json();
      //console.log(data);
      var array=data["data"];
      //console.log(array);
      var hasMatch=false;
      for(var i=0;i<array.length;i++)
      {
          console.log("Hello");
          if(array[i][3]==usersss.username && array[i][4]==usersss.psw)
          {
                //localStorage.setItem("id",usersss.id);
               //console.log(localStorage.getItem("id"));
               //this.router.navigate(['/after-sign-in']);
               hasMatch=true;
               break;
          }
      }
      if(hasMatch==true)
      {
          this.router.navigate(['/after-sign-in']);
          //console.log("Successful Login");
          //alert("successfully login");
          this.http.get("/webapi/v1/userpermission").subscribe(
    res1=>
    {
        //var info:object={};
        var list=res1.json();
        var datt=list["data"];
        //console.log(datt);
        for(var j=0;j<datt.length;j++)
        {
            if(usersss.username==datt[j][1])
            {
                    localStorage.setItem('info',JSON.stringify(datt[j]));
                    var a=JSON.parse(localStorage.getItem('info'));
                    console.log(a);
                    break;    
            }
        }


        //var info=[];
        //localStorage.setItem('info',JSON.stringify(datt));
        //var a=JSON.parse(localStorage.getItem('info'));
        //console.log(a);
    })

      }
      else
        {
          alert("username password is wrong");
          this.router.navigate(['/sign-in']);
          //return false;
        }
      /*var ar=data["message"];
      console.log(ar);

      if(ar=="Success")
      {
          localStorage.setItem("id",usersss.id);
           console.log(localStorage.getItem("id"));
         this.router.navigate(['/after-sign-in']);   
      }
      else
      {
          alert("plz enter valid username or password");
      }*/
   /* var hasMatch =false;

    for (var index = 0; index < data.length; ++index) {		
    if(data[index].id == usersss.id){
    hasMatch = true;
    console.log("Hello");
    //flag=1;
    break;
    }
}
	if(hasMatch==true)
	{
		this.router.navigate(['/after-sign-in']);
	}
	else
	{
		alert("username is wrong");
	}*/
}
    //console.log("Response",res);
    //console.log(this.users)}
  )
}

  
  ngOnInit(): void {
  /*.localStorage.setItem('id',this.hi);
  console.log(localStorage.getItem('id'));*/
  }

}
