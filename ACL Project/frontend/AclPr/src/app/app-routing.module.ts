






import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { UserComponent } from './user/user.component';
import { AdminComponent } from './admin/admin.component';
import { SignInComponent } from './sign-in/sign-in.component';
import { AfterSignInComponent } from './after-sign-in/after-sign-in.component';
import { AddUserComponent } from './add-user/add-user.component';
import { ViewUserComponent } from './view-user/view-user.component';
import { CreateGroupComponent } from './create-group/create-group.component';
import { ManagePermissionComponent } from './manage-permission/manage-permission.component';
import { UpdateUserComponent } from './update-user/update-user.component';
import { ViewGroupComponent } from './view-group/view-group.component';
import { AddMemberComponent } from './add-member/add-member.component';
import { ViewGroupMemberComponent } from './view-group-member/view-group-member.component';
import { AddResourceComponent } from './add-resource/add-resource.component';
import { RevokeComponent } from './revoke/revoke.component';
import { GrantComponent } from './grant/grant.component';
import { UserGrantComponent } from './user-grant/user-grant.component';
import { GroupGrantComponent } from './group-grant/group-grant.component';
import { GroupRevokeComponent } from './group-revoke/group-revoke.component';
import { UserRevokeComponent } from './user-revoke/user-revoke.component';
import { FileExplorerComponent } from './file-explorer/file-explorer.component';
import { ResViewComponent } from './res-view/res-view.component';

const routes: Routes = [{path:'',component:HomeComponent},
{path:'user',component:UserComponent},
{path:'admin',component:AdminComponent},
{path:'sign-in',component:SignInComponent},
{path:'after-sign-in',component:AfterSignInComponent},
{path:'add-user',component:AddUserComponent},
{path:'view-user',component:ViewUserComponent},
{path:'create-group',component:CreateGroupComponent},
{path:'manage-permission',component:ManagePermissionComponent},
{path:'update-user/:id',component:UpdateUserComponent},
{path:'view-group',component:ViewGroupComponent},
{path:'add-member/:id',component:AddMemberComponent},
{path:'view-group-member/:id',component:ViewGroupMemberComponent},
{path:'add-resource',component:AddResourceComponent},
{path:'revoke',component:RevokeComponent},
{path:'grant',component:GrantComponent},
{path:'user-grant',component:UserGrantComponent},
{path:'group-grant',component:GroupGrantComponent},
{path:'group-revoke',component:GroupRevokeComponent},
{path:'user-revoke',component:UserRevokeComponent},
{path:'file-explorer/:id',component:FileExplorerComponent},
{path:'res-view',component:ResViewComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
