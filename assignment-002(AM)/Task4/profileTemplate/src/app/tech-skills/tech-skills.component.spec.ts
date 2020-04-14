import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { TechSkillsComponent } from './tech-skills.component';

describe('TechSkillsComponent', () => {
  let component: TechSkillsComponent;
  let fixture: ComponentFixture<TechSkillsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ TechSkillsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(TechSkillsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
