using accommodation_service.Core.Service.Interfaces;
using accommodation_service.Protos;
using Grpc.Core;


namespace accommodation_service.Service
{
    public class TestService : AccommodationService.AccommodationServiceBase
    {

        private readonly IAccomodationService _accomodationService;

        public TestService(IAccomodationService accomodationService)
        {
            _accomodationService = accomodationService;
        }

        public override Task<GetResponse> Get(GetRequest request, ServerCallContext context)
        {
            return Task.FromResult(new GetResponse
            {
                Hello = "Hello from AccommodationService!"
            });
        }

        public override async Task<GetAllAccommodationsResponse> GetAll(GetAllRequest request, ServerCallContext context)
        {
            var accommodations = await _accomodationService.GetAsync();

            var protoAccommodations = accommodations.Select(a => new accommodation_service.Protos.Accommodation
            {
                Id = a.Id,
                Name = a.Name,
                MinGuestNum = a.MinGuestNum,
                MaxGuestNum = a.MaxGuestNum,
                AutomaticConfirmation = a.AutomaticConfirmation,
                Photo = { a.Photo },
                Address = new Address
                {
                    Id = a.Address.Id,
                    Country = a.Address.Country,
                    City = a.Address.City,
                    Street = a.Address.Street,
                    Number = a.Address.Number,
                },

                Benefits = { a.Benefits.Select(b => new Benefit
                {
                    Id = b.Id,
                    Name = b.Name,
                })},
                
                Appointments = { a.Appointment.Select(ap => new Appointment
                {
                    Id = ap.Id,
                    From = Google.Protobuf.WellKnownTypes.Timestamp.FromDateTime(ap.From),
                    To = Google.Protobuf.WellKnownTypes.Timestamp.FromDateTime(ap.To),
                    Status = (AppointmentStatus)ap.Status,
                    Price = ap.Price,
                    PerGuest = ap.PerGuest,
                })}
            });

            var response = new GetAllAccommodationsResponse
            {
                Accommodations = { protoAccommodations }
            };

            return response;
        }

        public override async Task<GetAccommodationByIdResponse> GetById(GetAccommodationByIdRequest request, ServerCallContext context)
        {
            var accommodation = await _accomodationService.GetAsync(request.Id);

            var protoAccommodation = new accommodation_service.Protos.Accommodation
            {
                Id = accommodation.Id,
                Name = accommodation.Name,
                MinGuestNum = accommodation.MinGuestNum,
                MaxGuestNum = accommodation.MaxGuestNum,
                AutomaticConfirmation = accommodation.AutomaticConfirmation,
                Photo = { accommodation.Photo },
                Address = new Address
                {
                    Id = accommodation.Address.Id,
                    Country = accommodation.Address.Country,
                    City = accommodation.Address.City,
                    Street = accommodation.Address.Street,
                    Number = accommodation.Address.Number,
                },

                Benefits = { accommodation.Benefits.Select(b => new Benefit
                {
                    Id = b.Id,
                    Name = b.Name,
                })},

                Appointments = { accommodation.Appointment.Select(ap => new Appointment
                {
                    Id = ap.Id,
                    From = Google.Protobuf.WellKnownTypes.Timestamp.FromDateTime(ap.From),
                    To = Google.Protobuf.WellKnownTypes.Timestamp.FromDateTime(ap.To),
                    Status = (AppointmentStatus)ap.Status,
                    Price = ap.Price,
                    PerGuest = ap.PerGuest,
                })}
            };

            var response = new GetAccommodationByIdResponse
            {
                Accommodation = protoAccommodation
            };

            return response;

        }
    }

}
