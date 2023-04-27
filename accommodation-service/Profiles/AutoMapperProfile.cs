using accommodation_service.Core.Model;
using accommodation_service.Core.Model.DTO;
using AutoMapper;


namespace accommodation_service.Profiles
{
    public class AutoMapperProfile : Profile
    {
        public AutoMapperProfile()
        {
            CreateMap<AddressDTO, Address>();
            CreateMap<BenefitDTO, Benefit>();
            CreateMap<AppointmentDTO, Appointment>();
            CreateMap<AccomodationDTO, Accomodation>();

            CreateMap<accommodation_service.Core.Model.Accomodation, accommodation_service.Protos.Accommodation>();
        }
    }
}
