import { Component } from '@angular/core';
import { AccommodationService } from 'src/app/services/accommodation/accommodation.service';
import { Router } from '@angular/router';
import { Accommodation } from 'src/app/model/accommodation';

@Component({
  selector: 'app-create-accommodation',
  templateUrl: './create-accommodation.component.html',
  styleUrls: ['./create-accommodation.component.css'],
})
export class CreateAccommodationComponent {
  public accommodation: Accommodation = new Accommodation();

  public benefits: string = '';
  public photos: string = '';
  public min_guest_num: string = '';
  public max_guest_num: string = '';

  constructor(
    private accommodationService: AccommodationService,
    private router: Router
  ) {}

  createAccommodation(): void {
    const splitBenefits = this.benefits.split(',').map((b) => b.trim());
    this.accommodation.benefits = splitBenefits;

    const splitPhotos = this.photos.split(',').map((b) => b.trim());
    this.accommodation.photos = splitPhotos;

    this.accommodationService
      .createAccommodation(this.accommodation)
      .subscribe(() => this.router.navigate(['/accommodations']));
  }
}
