import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NavbarComponent } from './modules/navbar/navbar.component';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { MatCardModule } from '@angular/material/card';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { ToastrModule } from 'ngx-toastr';
import { ViewComponent } from './modules/view/view.component';
import { MatDialogModule } from '@angular/material/dialog';
import { ReservationDetailsComponent } from './modules/reservation-details/reservation-details.component';
import { ReservationConfirmationComponent } from './modules/reservation-confirmation/reservation-confirmation.component';
import { MyReservationsComponent } from './modules/my-reservations/my-reservations.component';
import { CreateAccommodationComponent } from './modules/create-accommodation/create-accommodation.component';
import { RegisterComponent } from './modules/register/register.component';
import { LoginComponent } from './modules/login/login.component';
import { AddAppointmentComponent } from './modules/add-appointment/add-appointment.component';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { AllAccommodationsComponent } from './modules/all-accommodations/all-accommodations.component';
import { JwtInterceptor, JwtModule } from '@auth0/angular-jwt';
import { ProfileComponent } from './modules/profile/profile.component';
import { RateHostComponent } from './modules/rate-host/rate-host.component';
import { MyRatingsComponent } from './my-ratings/my-ratings.component';

@NgModule({
  declarations: [
    AppComponent,
    NavbarComponent,
    ViewComponent,
    ReservationDetailsComponent,
    ReservationConfirmationComponent,
    MyReservationsComponent,
    CreateAccommodationComponent,
    RegisterComponent,
    LoginComponent,
    AddAppointmentComponent,
    AllAccommodationsComponent,
    ProfileComponent,
    RateHostComponent,
    MyRatingsComponent,
  ],
  imports: [
    MatSlideToggleModule,
    BrowserAnimationsModule,
    ToastrModule.forRoot({
      positionClass: 'toast-top-right',
      preventDuplicates: true,
    }),
    FormsModule,
    BrowserModule,
    NgbModule,
    AppRoutingModule,
    MatCardModule,
    MatDialogModule,
    HttpClientModule,
    JwtModule.forRoot({
      config: {
        tokenGetter: () => localStorage.getItem('jwt-token'),
      },
    }),
  ],
  providers: [
    {
      provide: HTTP_INTERCEPTORS,
      useClass: JwtInterceptor,
      multi: true,
    },
  ],

  bootstrap: [AppComponent],
})
export class AppModule {}
