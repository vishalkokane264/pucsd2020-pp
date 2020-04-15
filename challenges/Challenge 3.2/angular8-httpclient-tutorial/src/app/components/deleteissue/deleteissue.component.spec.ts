import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { DeleteissueComponent } from './deleteissue.component';

describe('DeleteissueComponent', () => {
  let component: DeleteissueComponent;
  let fixture: ComponentFixture<DeleteissueComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ DeleteissueComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(DeleteissueComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
