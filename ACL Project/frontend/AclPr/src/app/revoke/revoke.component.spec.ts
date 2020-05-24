import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { RevokeComponent } from './revoke.component';

describe('RevokeComponent', () => {
  let component: RevokeComponent;
  let fixture: ComponentFixture<RevokeComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ RevokeComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(RevokeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
