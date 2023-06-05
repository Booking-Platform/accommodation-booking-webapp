import { Component, OnInit } from '@angular/core';
import { AccommodationService } from 'src/app/services/accommodation/accommodation.service';
import { MatDialog } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { Accommodation } from 'src/app/model/accommodation';

@Component({
  selector: 'app-view',
  templateUrl: './view.component.html',
  styleUrls: ['./view.component.css'],
})
export class ViewComponent implements OnInit {
  public accommodations: Accommodation[] = [];
  public filteredAccommodations: Accommodation[] = [];
  public from: string = '';
  public to: string = '';
  public numOfGuests: string = '';
  public city: string = '';
  public minPrice: number = 0;
  public maxPrice: number = Infinity;
  public featuredOnly: boolean = true;
  public selectedBenefits: string[] = [];
  public benefits: string[] = ['Gym', 'Pool', 'Wifi', 'Garage', 'Sauna'];
  constructor(
    private accommodationService: AccommodationService,
    public dialog: MatDialog,
    private router: Router
  ) {}

  ngOnInit() {}

  reserveAccommodation(acc: Accommodation) {
    this.router.navigate(['/accommodation-details'], {
      queryParams: {
        accommodationID: acc.id,
        startDate: this.from,
        endDate: this.to,
        numOfGuests: this.numOfGuests,
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
      this.filterAccommodations();
    });
  }

  filterAccommodations() {
    this.filteredAccommodations = this.accommodations.filter((acc: any) => {
      const price = acc.appointments[0].price;
      const fromDate = new Date(this.from);
      const toDate = new Date(this.to);
      const appointmentFromDate = new Date(acc.appointments[0].from);
      const appointmentToDate = new Date(acc.appointments[0].to);

      const selectedBenefits = this.selectedBenefits;

      // Proveri da li accommodation sadrži sve selektovane benefite
      const hasSelectedBenefits = selectedBenefits.every((benefit: string) =>
        acc.benefits.some((accBenefit: any) => accBenefit.name == benefit)
      );

      return (
        (this.minPrice === undefined || price >= this.minPrice) &&
        (this.maxPrice === undefined || price <= this.maxPrice) &&
        fromDate >= appointmentFromDate &&
        toDate <= appointmentToDate &&
        this.featuredOnly === acc.isFeaturedHost &&
        (selectedBenefits.length === 0 || hasSelectedBenefits)
      );
    });
  }

  toggleBenefitSelection(benefit: string) {
    const index = this.selectedBenefits.indexOf(benefit);
    if (index > -1) {
      // Ako je benefit već selektovan, ukloni ga iz niza selectedBenefits
      this.selectedBenefits.splice(index, 1);
    } else {
      // Ako benefit nije selektovan, dodaj ga u niz selectedBenefits
      this.selectedBenefits.push(benefit);
    }
  }
}
