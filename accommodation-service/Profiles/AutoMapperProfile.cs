using accommodation_service.Core.Model;
using accommodation_service.Core.Model.DTO;
using AutoMapper;
using MongoDB.Bson;
using static Google.Protobuf.Reflection.SourceCodeInfo.Types;

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
        }
    }
}
