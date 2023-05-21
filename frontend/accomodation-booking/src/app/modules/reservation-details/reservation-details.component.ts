import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Appointment } from 'src/app/model/appointment';
import { AccommodationService } from 'src/app/services/accommodation/accommodation.service';
import { ReservationService } from 'src/app/services/reservation/reservation.service';
import { JwtHelperService } from '@auth0/angular-jwt';

@Component({
  selector: 'app-reservation-details',
  templateUrl: './reservation-details.component.html',
  styleUrls: ['./reservation-details.component.css'],
})
export class ReservationDetailsComponent implements OnInit {
  public accommodationID: any = '';
  public startDate: any = '';
  public endDate: any = '';
  public guestNum: any = '';
  public userID: any = '';
  public accommodation: any | undefined;
  public appointment!: Appointment;
  public numOfGuests: any = '';
  public pricePerNight: any = '';
  public perGuest: any = '';
  public totalDays: any = '';
  public totalPrice: any = '';

  constructor(
    private route: ActivatedRoute,
    private reservationService: ReservationService,
    private accommodationService: AccommodationService,
    private router: Router,
    private jwtHelper: JwtHelperService
  ) {}

  ngOnInit(): void {
    this.route.queryParams.subscribe((params) => {
      this.startDate = params['startDate'];
      this.endDate = params['endDate'];
      this.accommodationID = params['accommodationID'];
      this.numOfGuests = params['numOfGuests'];

      this.accommodationService
        .getAccommodationByID(this.accommodationID)
        .subscribe((res: any) => {
          this.accommodation = res.accommodation;

          this.perGuest = this.accommodation.appointments[0].perGuest;
          this.pricePerNight = this.accommodation.appointments[0].price;
          this.totalDays = this.calculateDaysBetweenDates(
            this.startDate,
            this.endDate
          );
          if (this.perGuest) {
            this.totalPrice =
              this.totalDays * this.pricePerNight * this.numOfGuests;
          } else {
            this.totalPrice = this.totalDays * this.pricePerNight;
          }
        });
    });
  }

  createReservation(): void {
    var newReservation = {
      startDate: this.startDate,
      endDate: this.endDate,
      accommodationID: this.accommodationID,
      userID: this.jwtHelper.decodeToken().userId,
      guestNum: this.numOfGuests,
      automaticConfirmation:
        this.accommodation.automaticConfirmation.toString(),
    };
    this.reservationService.createReservation(newReservation).subscribe();
    this.router.navigate(['/myReservations']);
  }

  calculateDaysBetweenDates(from: string, to: string): number {
    const fromDate = new Date(from);
    const toDate = new Date(to);

    // Calculate the time difference in milliseconds
    const timeDiff = toDate.getTime() - fromDate.getTime();

    // Convert the time difference from milliseconds to days
    const daysDiff = Math.ceil(timeDiff / (1000 * 3600 * 24));

    return daysDiff;
  }
}
