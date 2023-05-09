import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Appointment } from 'src/app/model/appointment';
import { AccommodationService } from 'src/app/services/accommodation/accommodation.service';
import { ReservationService } from 'src/app/services/reservation/reservation.service';

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
  public userID: any = '6457aa1726a4e9026520c831';
  public accommodation: any | undefined;
  public appointment!: Appointment;

  constructor(
    private route: ActivatedRoute,
    private reservationService: ReservationService,
    private accommodationService: AccommodationService,
    private router: Router
  ) {}

  ngOnInit(): void {
    this.route.queryParams.subscribe((params) => {
      this.startDate = params['startDate'];
      this.endDate = params['endDate'];
      this.accommodationID = params['accommodationID'];

      this.accommodationService
        .getAccommodationByID(this.accommodationID)
        .subscribe((res: any) => {
          this.accommodation = res.accommodation;
        });
    });
  }

  createReservation(): void {
    window.alert(JSON.stringify(this.accommodation));
    window.alert(this.accommodation.automaticConfirmation.toString());
    var newReservation = {
      startDate: this.startDate,
      endDate: this.endDate,
      accommodationID: this.accommodationID,
      userID: this.userID,
      automaticConfirmation:
        this.accommodation.automaticConfirmation.toString(),
    };
    this.reservationService.createReservation(newReservation).subscribe();
    this.router.navigate(['/myReservations']);
  }
}
