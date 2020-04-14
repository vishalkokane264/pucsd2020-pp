import { Component, OnInit, NgModule } from '@angular/core';
import {PostService} from '../services/post.service';

@Component({
  selector: 'app-create',
  templateUrl: './create.component.html',
  styleUrls: ['./create.component.scss']
})


export class CreateComponent implements OnInit {
products=[]
  constructor(private service:PostService) {  }
  ngOnInit(): void {
    this.service.getSendRequest().subscribe((data:any[])=>{console.log(data);this.products=data;})
  }

  onClickSubmit(formData) {
    var userVar = {
      first_name: formData.firstname,
      last_name: formData.lastname,
      email: formData.emailId,
      password: formData.passwordVal,
      contact_number: formData.contact,
      updated_by: formData.updateBy
    }
//  this.service.create(userVar);

    }
}

