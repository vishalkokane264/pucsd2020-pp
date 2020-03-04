import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { PersonalQualitiesComponent } from './personal-qualities.component';

describe('PersonalQualitiesComponent', () => {
  let component: PersonalQualitiesComponent;
  let fixture: ComponentFixture<PersonalQualitiesComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ PersonalQualitiesComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(PersonalQualitiesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
