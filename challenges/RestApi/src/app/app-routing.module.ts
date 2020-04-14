import { NgModule, Component } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { CreateComponent } from './create/create.component';
import { UpdateComponent } from './update/update.component';
import { SearchComponent } from './search/search.component';


const routes: Routes = [
  {
    path:'create',
    component:CreateComponent
  },
  {
    path:'update',
    component:UpdateComponent
  },
  {
    path:'search',
    component:SearchComponent
  }

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
