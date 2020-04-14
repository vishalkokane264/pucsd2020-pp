import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { CareerObjComponent } from './career-obj/career-obj.component';
import { DeclarationComponent } from './declaration/declaration.component';
import { EducationQualificationComponent } from './education-qualification/education-qualification.component';
import { PersonalQualitiesComponent } from './personal-qualities/personal-qualities.component';
import { ProjectsComponent } from './projects/projects.component';
import { TechSkillsComponent } from './tech-skills/tech-skills.component';
import { WorkExpComponent } from './work-exp/work-exp.component';

@NgModule({
  declarations: [
    AppComponent,
    CareerObjComponent,
    DeclarationComponent,
    EducationQualificationComponent,
    PersonalQualitiesComponent,
    ProjectsComponent,
    TechSkillsComponent,
    WorkExpComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
