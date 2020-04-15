import { NgModule, Component } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { AddIssueComponent } from './components/add-issue/add-issue.component';
import { EditIssueComponent } from './components/edit-issue/edit-issue.component';
import { IssueListComponent } from './components/issue-list/issue-list.component';
import { AddUserComponent } from './add-user/add-user.component';


const routes: Routes = [

  {
    path:'addissue',
    component:AddIssueComponent
  },
  {
    path:'editissue',
    component:EditIssueComponent
  },
  {
    path:'issuelist',
    component:IssueListComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
