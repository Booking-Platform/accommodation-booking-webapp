import { Component, OnInit } from '@angular/core';
import { JwtHelperService } from '@auth0/angular-jwt';
import { ReservationService } from 'src/app/services/reservation/reservation.service';

@Component({
  selector: 'app-my-reservations',
  templateUrl: './my-reservations.component.html',
  styleUrls: ['./my-reservations.component.css'],
})
export class MyReservationsComponent implements OnInit {
  public reservations: any[] = [];
  public userID: string = '';

  constructor(
    private reservationService: ReservationService,
    private jwtHelper: JwtHelperService
  ) {}

  ngOnInit(): void {
    this.userID = this.jwtHelper.decodeToken().userId;
    this.reservationService
      .getAllReservationsByUserID(this.userID)
      .subscribe((res: any) => {
        this.reservations = res;
      });
  }

  cancelReservation(reservation: any): void {
    this.reservationService
      .changeReservationStatus(reservation.Id, '3')
      .subscribe();
    location.reload();
  }
}
