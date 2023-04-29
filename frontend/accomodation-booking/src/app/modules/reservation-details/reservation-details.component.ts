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
  public accommodationID: any = '123321';
  public startDate: any = '';
  public endDate: any = '';
  public guestNum: any = '';

  public address = new Address('1', 'Serbia', 'Belgrade', 'Main Street', '123');
  public benefits = [
    new Benefit('1', 'Free Wi-Fi'),
    new Benefit('2', 'Swimming pool'),
    new Benefit('3', 'Gym'),
  ];

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
    this.benefits
  );

  constructor(private route: ActivatedRoute, private reservationService: ReservationService) {}

  ngOnInit(): void {
    this.route.queryParams.subscribe((params) => {
      this.startDate = '2023-11-11'
      this.endDate = '2024-11-11'
      // this.startDate = params['startDate'];
      // this.endDate = params['endDate'];
    });
  }

  createReservation(): void {
    var newReservation = {
      "startDate": this.startDate,
      "endDate": this.endDate,
      "accommodationID": this.accommodationID
    }
    this.reservationService.createReservation(newReservation).subscribe()
  }
}
