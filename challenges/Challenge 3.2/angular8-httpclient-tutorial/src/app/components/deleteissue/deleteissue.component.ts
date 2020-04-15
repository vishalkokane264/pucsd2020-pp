import { Component, OnInit, NgZone } from '@angular/core';

import { BugService } from '../../../shared/shared/bug.service';
import { FormBuilder, FormGroup, FormControl, FormControlName } from '@angular/forms';
import { Router, ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-deleteissue',
  templateUrl: './deleteissue.component.html',
  styleUrls: ['./deleteissue.component.css']
})
export class DeleteissueComponent implements OnInit {
  IssuesList: any = [];
  deleteForm:FormGroup
userId:FormControl
  constructor(
    private actRoute: ActivatedRoute,    
    public bugService: BugService,
    public fb: FormBuilder,
    private ngZone: NgZone,
    private router: Router
  ) {
   }

  ngOnInit(): void {
    this.updateForm()
  }
  updateForm(){
    this.deleteForm = this.fb.group({
      userId: [''],
    })
  }

  deleteUser(){
    var index='1'
    alert('User deleted'+this.deleteForm.value.userId)
    //    var index = index = this.IssuesList.map(x => {return x.issue_name}).indexOf(data.issue_name);
    return this.bugService.DeleteBug(this.deleteForm.value.userId).subscribe(res => {
     this.IssuesList.splice(index, 1)
      console.log('User deleted!')
    });
  }
}
