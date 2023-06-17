import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Appointment } from 'src/app/model/appointment';
import { RatingService } from 'src/app/services/rating/rating.service';

@Component({
  selector: 'app-ratings',
  templateUrl: './ratings.component.html',
  styleUrls: ['./ratings.component.css'],
})
export class RatingsComponent implements OnInit {
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
    private ratingService: RatingService,
    private router: Router
  ) {}

  ngOnInit(): void {
    this.route.queryParams.subscribe((params) => {
      this.startDate = params['startDate'];
      this.endDate = params['endDate'];
      this.accommodationID = params['accommodationID'];
      this.numOfGuests = params['numOfGuests'];
    });

    this.ratingService.getRattingsForAccommodation(this.accommodationID).subscribe((res) => {

    });
  
  }
}
