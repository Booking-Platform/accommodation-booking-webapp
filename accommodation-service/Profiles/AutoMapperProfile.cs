using AccomodationServiceLibrary.Core.Model;
using AccomodationServiceLibrary.Core.Model.DTO;
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
