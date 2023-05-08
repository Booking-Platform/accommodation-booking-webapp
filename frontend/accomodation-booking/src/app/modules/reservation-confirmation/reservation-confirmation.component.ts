import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { AccommodationService } from 'src/app/services/accommodation/accommodation.service';
import { ReservationService } from 'src/app/services/reservation/reservation.service';

@Component({
  selector: 'app-reservation-confirmation',
  templateUrl: './reservation-confirmation.component.html',
  styleUrls: ['./reservation-confirmation.component.css']
})
export class ReservationConfirmationComponent implements OnInit {

  public reservations: any[] = [];
  
  constructor(
    private reservationService: ReservationService,
  
    ) {}

  ngOnInit(): void {
    this.reservationService.getReservationsForConfirmation().subscribe((res: any) => {
      this.reservations = res;
      console.log(this.reservations);
    });
   }

   confirm(reservation: any) {
    window.alert(reservation.Id)
  }
    
}
