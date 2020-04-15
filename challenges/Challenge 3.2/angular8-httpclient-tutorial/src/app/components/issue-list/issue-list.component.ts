import { Component, OnInit, NgZone } from '@angular/core';

import { BugService } from '../../../shared/shared/bug.service';
import { FormBuilder, FormGroup, FormControl, FormControlName } from '@angular/forms';
import { Router } from '@angular/router';

import { Key } from 'protractor';
import {User} from '../../../shared/bug'
import { Observable } from 'rxjs';
@Component({
  selector: 'app-issue-list',
  templateUrl: './issue-list.component.html',
  styleUrls: ['./issue-list.component.css']
})
export class IssueListComponent implements OnInit {
  searchForm:FormGroup
  uid:FormControl
  uname:FormControl


  users=new Array<User>();
  IssuesList:any=[]
  UserList:any=[]

  constructor(
    public fb: FormBuilder,
    private ngZone: NgZone,
    private router: Router,
    public bugService: BugService
  ) { }
  ngOnInit(): void {
    this.loadForm();
    this.loadEmployees();
  }
  loadForm(){
    this.searchForm=this.fb.group({
      userId:['']
    })
  }

  loadEmployees(){
    return this.bugService.GetIssues().subscribe((data: {}) => {
      this.IssuesList = data;
      console.log(this.IssuesList);
    })
   }

   submitForm(){
     console.log(this.searchForm.value)
     this.getUserValue()
//     this.loadForm()
   }

   getUserValue() 
     {
     return this.bugService.GetIssue(this.searchForm.value.userId).subscribe(data=>{
       this.IssuesList=data;
       console.log(data)       
     });

   
  
  
  
  }
  deleteIusse(data){
    var index='1'
    //    var index = index = this.IssuesList.map(x => {return x.issue_name}).indexOf(data.issue_name);
    return this.bugService.DeleteBug(13).subscribe(res => {
     this.IssuesList.splice(index, 1)
      console.log('Issue deleted!')
    })
  }


}
