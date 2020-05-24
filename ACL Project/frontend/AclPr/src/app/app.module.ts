







import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HomeComponent } from './home/home.component';
import { UserComponent } from './user/user.component';
import { AdminComponent } from './admin/admin.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { SignInComponent } from './sign-in/sign-in.component';
import { HttpModule } from '@angular/http';
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

@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    UserComponent,
    AdminComponent,
    SignInComponent,
    AfterSignInComponent,
    AddUserComponent,
    ViewUserComponent,
    CreateGroupComponent,
    ManagePermissionComponent,
    UpdateUserComponent,
    ViewGroupComponent,
    AddMemberComponent,
    ViewGroupMemberComponent,
    AddResourceComponent,
    RevokeComponent,
    GrantComponent,
    UserGrantComponent,
    GroupGrantComponent,
    GroupRevokeComponent,
    UserRevokeComponent,
    FileExplorerComponent,
    ResViewComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,FormsModule,HttpModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
