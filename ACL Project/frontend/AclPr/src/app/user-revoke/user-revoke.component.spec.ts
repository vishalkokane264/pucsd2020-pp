import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { UserRevokeComponent } from './user-revoke.component';

describe('UserRevokeComponent', () => {
  let component: UserRevokeComponent;
  let fixture: ComponentFixture<UserRevokeComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ UserRevokeComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(UserRevokeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
