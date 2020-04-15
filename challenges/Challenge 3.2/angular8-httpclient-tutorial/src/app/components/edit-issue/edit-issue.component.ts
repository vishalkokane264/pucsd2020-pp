import { Component, OnInit, NgZone } from '@angular/core';

import { BugService } from '../../../shared/shared/bug.service';
import { FormBuilder, FormGroup, FormControl, FormControlName } from '@angular/forms';
import { Router, ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-edit-issue',
  templateUrl: './edit-issue.component.html',
  styleUrls: ['./edit-issue.component.css']
})
export class EditIssueComponent implements OnInit {
  IssuesList: any = [];
  updateIssueForm: FormGroup;
  userIdval:FormControl;

  constructor(
    private actRoute: ActivatedRoute,    
    public bugService: BugService,
    public fb: FormBuilder,
    private ngZone: NgZone,
    private router: Router
  ) {
    var id = this.actRoute.snapshot.paramMap.get('userIdval');
    this.bugService.GetIssue(id).subscribe((data) => {
      this.updateIssueForm = this.fb.group({
        issue_name: [data.first_name],
        issue_message: [data.id]
      })
    })
   }

  ngOnInit(): void {
    this.updateForm()
  }

  updateForm(){
    this.updateIssueForm = this.fb.group({
      issue_name: [''],
      issue_message: ['']
    })
  }
  submitForm(){
    console.log("Hello all"+this.updateIssueForm.value.issue_name);
    var mid=this.updateIssueForm.value.issue_name;
    var mpass=this.updateIssueForm.value.issue_message;
    alert(mid+' '+mpass);
    var id = this.actRoute.snapshot.paramMap.get('userIdval');
    this.bugService.UpdateBug(id, this.updateIssueForm.value).subscribe(res => {
      this.ngZone.run(() => this.router.navigateByUrl('/issues-list'))
    })
  }
}
