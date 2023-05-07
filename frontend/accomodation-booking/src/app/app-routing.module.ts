import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ViewComponent } from './modules/view/view.component';
import { ReservationDetailsComponent } from './modules/reservation-details/reservation-details.component';
import { ReservationConfirmationComponent } from './modules/reservation-confirmation/reservation-confirmation.component';
import { MyReservationsComponent } from './modules/my-reservations/my-reservations.component';
import { CreateAccommodationComponent } from './modules/create-accommodation/create-accommodation.component';
import { RegisterComponent } from './modules/register/register.component';
import { LoginComponent } from './modules/login/login.component';

const routes: Routes = [
  { path: 'view', component: ViewComponent },
  { path: 'accommodation-details', component: ReservationDetailsComponent },
  { path: 'reservations', component: ReservationConfirmationComponent },
  { path: 'myReservations', component: MyReservationsComponent },
  { path: 'createAccommodation', component: CreateAccommodationComponent },
  { path: 'register', component: RegisterComponent },
  { path: 'login', component: LoginComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
