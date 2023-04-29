import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ReserveAccommodationComponent } from './reserve-accommodation.component';

describe('ReserveAccommodationComponent', () => {
  let component: ReserveAccommodationComponent;
  let fixture: ComponentFixture<ReserveAccommodationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ReserveAccommodationComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ReserveAccommodationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
