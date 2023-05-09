import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AllAccommodationsComponent } from './all-accommodations.component';

describe('AllAccommodationsComponent', () => {
  let component: AllAccommodationsComponent;
  let fixture: ComponentFixture<AllAccommodationsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AllAccommodationsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AllAccommodationsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
