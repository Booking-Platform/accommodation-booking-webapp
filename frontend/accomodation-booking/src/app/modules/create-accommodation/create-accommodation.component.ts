import { Address } from './../../model/address';
import { Component } from '@angular/core';
import { AccommodationService } from 'src/app/services/accommodation/accommodation.service';
import { Router } from '@angular/router';
import { Accommodation } from 'src/app/model/accommodation';
import { JwtHelperService } from '@auth0/angular-jwt';

@Component({
  selector: 'app-create-accommodation',
  templateUrl: './create-accommodation.component.html',
  styleUrls: ['./create-accommodation.component.css'],
})
export class CreateAccommodationComponent {
  public accommodation: Accommodation = new Accommodation();

  public city: string = '';
  public country: string = '';
  public street: string = '';
  public number: string = '';
  public benefits: string = '';
  public photos: string = '';
  public min_guest_num: string = '';
  public max_guest_num: string = '';

  constructor(
    private accommodationService: AccommodationService,
    private router: Router,
    private jwtHelper: JwtHelperService
  ) {}

  createAccommodation(): void {
    var userID = this.jwtHelper.decodeToken().userId;

    this.accommodation.host_id = userID
    
    const splitBenefits = this.benefits.split(',').map((b) => b.trim());
    this.accommodation.benefits = splitBenefits;

    const splitPhotos = this.photos.split(',').map((b) => b.trim());
    this.accommodation.photos = splitPhotos;

    this.accommodation.address.country = this.country;
    this.accommodation.address.city = this.city;
    this.accommodation.address.number = this.number;
    this.accommodation.address.street = this.street;

    this.accommodationService
      .createAccommodation(this.accommodation)
      .subscribe(() => this.router.navigate(['/allAccommodations']));
    this.router.navigate(['/allAccommodations']);
  }
}
