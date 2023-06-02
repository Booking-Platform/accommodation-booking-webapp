import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { JwtHelperService } from '@auth0/angular-jwt';
import { ToastrService } from 'ngx-toastr';
import { HostServiceService } from 'src/app/services/host/host-service.service';
import { RatingService } from 'src/app/services/rating/rating.service';

@Component({
  selector: 'app-rate-host',
  templateUrl: './rate-host.component.html',
  styleUrls: ['./rate-host.component.css']
})
export class RateHostComponent implements OnInit{
  
  public hosts: any[] = [];
  public rating: any = '';
  public guestID: any = '';

  constructor(private hostService: HostServiceService, private jwtHelper: JwtHelperService, private ratingService: RatingService, private toastService: ToastrService, private router: Router) {}
  
  ngOnInit(): void {
    this.guestID = this.jwtHelper.decodeToken().userId;

  
    this.hostService.getHostsForRatingByUserID(this.guestID).subscribe((res: any) => {
      this.hosts = res;
    });

  
  }

  
  stars = [1, 2, 3, 4, 5];

  rateCard(host: any, rating: number) {
    host.selectedRating = rating;
    var newRating = {
      hostID: host.Id,
      guestID: this.guestID,
      rating: JSON.stringify(rating) 
    }

    this.ratingService.createRating(newRating).subscribe();
    this.toastService.success('You leave comment for ' + host.Name + ' ' + host.Surname + '!', 'Success');

  }

}
