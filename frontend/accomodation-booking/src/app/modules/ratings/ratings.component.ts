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
  public accommodationName: any = '';
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
  public hostID = '';
  public accommodationID: any = '';

  public accommodationComments: any[] = [];
  public hostComments: any[] = [];


  constructor(
    private route: ActivatedRoute,
    private ratingService: RatingService,
    private router: Router
  ) {}

  continue(): void {
    this.router.navigate(['/accommodation-details'], {
      queryParams: {
        accommodationID: this.accommodationID,
        startDate: this.startDate,
        endDate: this.endDate,
        numOfGuests: this.numOfGuests,
      },
    });
  }

  ngOnInit(): void {
    this.route.queryParams.subscribe((params) => {
      this.startDate = params['startDate'];
      this.endDate = params['endDate'];
      this.accommodationName = params['accommodationName'];
      this.numOfGuests = params['numOfGuests'];
      this.hostID = params['hostID'];
      this.accommodationID = params['accommodationID'];
    });
  
    window.alert(this.accommodationID)
    this.ratingService.getRattingsForAccommodation(this.accommodationName).subscribe((res) => {
      this.accommodationComments = res;    
    });
  
    this.ratingService.getRattingsForHost(this.hostID).subscribe((res) => {
      this.hostComments = res;    
    });
  
  }
}
