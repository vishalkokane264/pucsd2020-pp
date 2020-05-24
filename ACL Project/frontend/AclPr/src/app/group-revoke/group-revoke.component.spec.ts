import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { GroupRevokeComponent } from './group-revoke.component';

describe('GroupRevokeComponent', () => {
  let component: GroupRevokeComponent;
  let fixture: ComponentFixture<GroupRevokeComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ GroupRevokeComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(GroupRevokeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
