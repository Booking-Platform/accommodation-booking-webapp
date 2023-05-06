import { Component, OnInit } from '@angular/core';
import { ReservationService } from 'src/app/services/reservation/reservation.service';

@Component({
  selector: 'app-my-reservations',
  templateUrl: './my-reservations.component.html',
  styleUrls: ['./my-reservations.component.css']
})
export class MyReservationsComponent implements OnInit {
  public reservations: any[] = [];
  public userID: string = '';

  constructor(private reservationService: ReservationService){}

  ngOnInit(): void {
    this.userID = "511abbea5af7118727bf2312"
    this.reservationService.getAllReservationsByUserID(this.userID).subscribe((res: any) => {
      this.reservations = res.reservations;
      window.alert(JSON.stringify(this.reservations))
    
    });
  }

  cancelReservation(reservation: any): void {

  }

}
