import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ManagePermissionComponent } from './manage-permission.component';

describe('ManagePermissionComponent', () => {
  let component: ManagePermissionComponent;
  let fixture: ComponentFixture<ManagePermissionComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ManagePermissionComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ManagePermissionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
