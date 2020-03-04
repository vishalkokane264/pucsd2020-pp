import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { CareerObjComponent } from './components/career-obj/career-obj.component';
import { TechSkillsComponent } from './components/tech-skills/tech-skills.component';
import { PersonalQualitiesComponent } from './components/personal-qualities/personal-qualities.component';
import { EducationQualificationComponent } from './components/education-qualification/education-qualification.component';
import { WorkExpComponent } from './components/work-exp/work-exp.component';
import { ProjectsComponent } from './components/projects/projects.component';
import { DeclarationComponent } from './components/declaration/declaration.component';


const routes: Routes = [
  {
    path: 'Career_Objective', 
    component: CareerObjComponent 
  },
  { path: 'Technical_skills', 
    component: TechSkillsComponent 
  },
  { path: 'Personal_Qualities', 
    component: PersonalQualitiesComponent 
  },
  { path: 'Education_Qualification', 
    component: EducationQualificationComponent 
  },
  { path: 'Work_Experience', 
    component: WorkExpComponent 
  },
  { path: 'Project_and_Assignments', 
    component: ProjectsComponent 
  },
  { path: 'Declaration', 
    component: DeclarationComponent 
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
