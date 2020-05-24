import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { GroupGrantComponent } from './group-grant.component';

describe('GroupGrantComponent', () => {
  let component: GroupGrantComponent;
  let fixture: ComponentFixture<GroupGrantComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ GroupGrantComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(GroupGrantComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
