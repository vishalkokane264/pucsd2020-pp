import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewGroupMemberComponent } from './view-group-member.component';

describe('ViewGroupMemberComponent', () => {
  let component: ViewGroupMemberComponent;
  let fixture: ComponentFixture<ViewGroupMemberComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ViewGroupMemberComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ViewGroupMemberComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
