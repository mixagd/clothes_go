import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ClothFormComponent } from './cloth-form.component';

describe('ClothFormComponent', () => {
  let component: ClothFormComponent;
  let fixture: ComponentFixture<ClothFormComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ClothFormComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ClothFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
