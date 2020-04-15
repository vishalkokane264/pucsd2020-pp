import { Component, OnInit, NgZone } from '@angular/core';

import { BugService } from '../../shared/shared/bug.service';
import { FormBuilder, FormGroup } from '@angular/forms';
import { Router } from '@angular/router';

@Component({
  selector: 'app-add-user',
  templateUrl: './add-user.component.html',
  styleUrls: ['./add-user.component.css']
})
export class AddUserComponent implements OnInit {
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
        issue_name: [''],
        issue_message: ['']
      })
    }
  
    submitForm(){
      var userVar = { 
        first_name: "VishalSakharam",
        last_name: "Kokane",
        email:"vishalkokane24@gmail.com",
        password: "123123",
        contact_number: "8698557586",
        updated_by: 1
      }
      this.bugService.CreateBug(userVar).subscribe(res => {
        console.log('Issue added!')
        this.ngZone.run(() => this.router.navigateByUrl('/issues-list'))
      });
    }
  
  
  }
  