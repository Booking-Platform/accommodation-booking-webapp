import { AfterViewInit, Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Accommodation } from 'src/app/model/accommodation';
import { Address } from 'src/app/model/address';
import { Benefit } from 'src/app/model/benefit';
import { AccommodationService } from 'src/app/services/accommodation/accommodation.service';
import { MatDialog } from '@angular/material/dialog';
import { Router } from '@angular/router';

@Component({
  selector: 'app-view',
  templateUrl: './view.component.html',
  styleUrls: ['./view.component.css'],
})
export class ViewComponent implements OnInit {
  public accommodations: Accommodation[] = [];
  public from: string = '';
  public to: string = '';
  public numOfGuests: string = '';
  public city: string = '';

  constructor(
    private accommodationService: AccommodationService,
    public dialog: MatDialog,
    private router: Router
  ) {}

  ngOnInit() {
    this.accommodationService.getAccommodations().subscribe((res: any) => {
      this.accommodations = res.accommodations;
    });
  }

  reserveAccommodation(acc: any) {
    this.router.navigate(['/accommodation-details'], {
      queryParams: {
        accommodationID: 'accID',
        startDate: '2023-10-10',
        endDate: '2023-10-11',
      },
    });
  }

  search() {
    var searchParams = {
      from: this.from,
      to: this.to,
      numOfGuests: this.numOfGuests,
      city: this.city,
    };
    this.accommodationService.search(searchParams).subscribe((res: any) => {
      this.accommodations = res.accommodations;
    });
  }
}
