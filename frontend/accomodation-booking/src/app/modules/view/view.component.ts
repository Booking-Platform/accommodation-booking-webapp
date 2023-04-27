import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Accommodation } from 'src/app/model/accommodation';
import { Address } from 'src/app/model/address';
import { Benefit } from 'src/app/model/benefit';
import { AccommodationService } from 'src/app/services/accommodation/accommodation.service';

@Component({
  selector: 'app-view',
  templateUrl: './view.component.html',
  styleUrls: ['./view.component.css'],
})
export class ViewComponent implements OnInit {
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

  public accommodations: Accommodation[] = [];

  constructor(private accommodationService: AccommodationService) {}

  ngOnInit(): void {
    this.accommodationService.getAccommodations().subscribe((res: any) => {
      this.accommodations = res;
    });

    this.accommodations.push(this.accommodation);
    this.accommodations.push(this.accommodation);
    this.accommodations.push(this.accommodation);
    this.accommodations.push(this.accommodation);
    this.accommodations.push(this.accommodation);
    this.accommodations.push(this.accommodation);
    this.accommodations.push(this.accommodation);
  }
}
