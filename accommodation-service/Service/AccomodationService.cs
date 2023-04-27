using accommodation_service.Core.Model;
using accommodation_service.Core.Model.DTO;
using accommodation_service.Core.Repository.Interfaces;
using accommodation_service.Core.Service.Interfaces;
using AutoMapper;
using MongoDB.Bson;

namespace accommodation_service.Core.Service
{
    public class AccomodationService : IAccomodationService
    {
        private readonly IAccomodationRepository _accomodationRepository;
        private readonly IMapper _mapper;
        private readonly IAppointmentService _appointmentService;

        public AccomodationService(IAccomodationRepository accomodationRepository, IMapper mapper, IAppointmentService appointmentService)
        {
            _accomodationRepository = accomodationRepository;
            _mapper = mapper;
            _appointmentService = appointmentService;
        }
        public async Task CreateAsync(AccomodationDTO newAccomodation)
        {
            var accomodation = _mapper.Map<Accomodation>(newAccomodation);
            accomodation.Benefits.ForEach(benefit => benefit.Id = ObjectId.GenerateNewId().ToString());
            accomodation.Address.Id = ObjectId.GenerateNewId().ToString();
            await _accomodationRepository.CreateAsync(accomodation);
        }

        public async Task<List<Accomodation>> GetAsync()
        {
            return await _accomodationRepository.GetAsync();
        }

        public async Task<Accomodation?> GetAsync(string id)
        {
            return await _accomodationRepository.GetAsync(id);
        }

        public async Task RemoveAsync(string id)
        {
            await _accomodationRepository.RemoveAsync(id);
        }

        public async Task<List<Accomodation>> SearchAsync(string location, int numGuests, DateTime startDate, DateTime endDate)
        {
           return await _accomodationRepository.SearchAsync(location, numGuests, startDate, endDate);
        }

    }
}
