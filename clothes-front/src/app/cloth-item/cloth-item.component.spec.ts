import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ClothItemComponent } from './cloth-item.component';

describe('ClothItemComponent', () => {
  let component: ClothItemComponent;
  let fixture: ComponentFixture<ClothItemComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ClothItemComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ClothItemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
