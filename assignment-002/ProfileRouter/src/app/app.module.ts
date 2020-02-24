import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { CareerObjComponent } from './components/career-obj/career-obj.component';
import { TechSkillsComponent } from './components/tech-skills/tech-skills.component';
import { PersonalQualitiesComponent } from './components/personal-qualities/personal-qualities.component';
import { EducationQualificationComponent } from './components/education-qualification/education-qualification.component';
import { WorkExpComponent } from './components/work-exp/work-exp.component';
import { ProjectsComponent } from './components/projects/projects.component';
import { DeclarationComponent } from './components/declaration/declaration.component';

@NgModule({
  declarations: [
    AppComponent,
    CareerObjComponent,
    TechSkillsComponent,
    PersonalQualitiesComponent,
    EducationQualificationComponent,
    WorkExpComponent,
    ProjectsComponent,
    DeclarationComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
