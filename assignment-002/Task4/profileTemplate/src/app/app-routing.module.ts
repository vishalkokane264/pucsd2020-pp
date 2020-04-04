import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { CareerObjComponent } from './career-obj/career-obj.component';


const routes: Routes = [
  {
    path: 'Career_Objective', 
    component: CareerObjComponent 
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
