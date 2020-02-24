import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CareerObjComponent } from './career-obj.component';

describe('CareerObjComponent', () => {
  let component: CareerObjComponent;
  let fixture: ComponentFixture<CareerObjComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CareerObjComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CareerObjComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
