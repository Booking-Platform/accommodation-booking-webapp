import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ReservationConfirmationComponent } from './reservation-confirmation.component';

describe('ReservationConfirmationComponent', () => {
  let component: ReservationConfirmationComponent;
  let fixture: ComponentFixture<ReservationConfirmationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ReservationConfirmationComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ReservationConfirmationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
