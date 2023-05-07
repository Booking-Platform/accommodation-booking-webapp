import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Accommodation } from 'src/app/model/accommodation';
import { Address } from 'src/app/model/address';
import { Benefit } from 'src/app/model/benefit';
import { ReservationService } from 'src/app/services/reservation/reservation.service';

@Component({
  selector: 'app-reservation-details',
  templateUrl: './reservation-details.component.html',
  styleUrls: ['./reservation-details.component.css'],
})
export class ReservationDetailsComponent implements OnInit {
  public accommodationID: any = '644abbea0af7618727bf6512';
  public startDate: any = '';
  public endDate: any = '';
  public guestNum: any = '';
  public userID: any = '511abbea5af7118727bf2312';

  public address = new Address('Serbia', 'Belgrade', 'Main Street', '123');
  // public benefits = [
  //   new Benefit('Free Wi-Fi'),
  //   new Benefit('Swimming pool'),
  //   new Benefit('Gym'),
  // ];

  public accommodation = new Accommodation(
    '1',
    'Luxury Apartment in Belgrade',
    '2',
    '4',
    this.address,
    [
      'https://i.dummyjson.com/data/products/1/3.jpg',
      'https://i.dummyjson.com/data/products/1/4.jpg',
    ],
    ['pool']
  );

  constructor(
    private route: ActivatedRoute,
    private reservationService: ReservationService
  ) {}

  ngOnInit(): void {
    this.route.queryParams.subscribe((params) => {
      this.startDate = '2023-11-11';
      this.endDate = '2024-11-11';
      // this.startDate = params['startDate'];
      // this.endDate = params['endDate'];
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
  }
}
