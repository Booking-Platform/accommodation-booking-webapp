import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Accommodation } from 'src/app/model/accommodation';
import { Address } from 'src/app/model/address';
import { Benefit } from 'src/app/model/benefit';
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
    var newReservation = {
      startDate: this.startDate,
      endDate: this.endDate,
      accommodationID: this.accommodationID,
      userID: this.userID,
    };
    this.reservationService.createReservation(newReservation).subscribe();
    this.router.navigate(['/myReservations']);
  }
}
