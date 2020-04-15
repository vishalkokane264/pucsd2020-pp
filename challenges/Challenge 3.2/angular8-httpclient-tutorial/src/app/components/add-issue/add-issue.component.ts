import { Component, OnInit, NgZone } from '@angular/core';

import { BugService } from '../../../shared/shared/bug.service';
import { FormBuilder, FormGroup } from '@angular/forms';
import { Router } from '@angular/router';

@Component({
  selector: 'app-add-issue',
  templateUrl: './add-issue.component.html',
  styleUrls: ['./add-issue.component.css']
})
export class AddIssueComponent implements OnInit {
issueForm:FormGroup;
IssueArr:any=[];

  constructor(
    public fb: FormBuilder,
    private ngZone: NgZone,
    private router: Router,
    public bugService: BugService
  ) { }

  ngOnInit(): void {
    this.addIssue()
  }

  addIssue(){
    this.issueForm = this.fb.group({
      first_name:[''],
      last_name:[''],
      email:[''],
      contact_number:[''],
      password:[''],
      updated_by:[''],
    })
  }
  submitForm(){

    console.log(this.issueForm.value)
    var userVar = {
/*      first_name:"sneha1",
      last_name:"shete1",
      email:"sshete1@gmail.com",
      password:"sneha@123",
      contact_number:"1123456789",
      updated_by:"1" 
*/
      first_name: this.issueForm.value.first_name,
      last_name: this.issueForm.value.last_name,
      email:this.issueForm.value.email,
      password: this.issueForm.value.password,
      contact_number: this.issueForm.value.contact_number,

/*
      updated_by: this.issueForm.value.updated_by
    */
       }
    this.bugService.CreateBug(userVar).subscribe(res => {
      console.log('Issue added!')
      this.ngZone.run(() => this.router.navigateByUrl('/issues-list'))
    });
  }


}
