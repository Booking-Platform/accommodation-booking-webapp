using AccomodationServiceLibrary.Core.Model;
using AccomodationServiceLibrary.Core.Model.DTO;
using AccomodationServiceLibrary.Core.Repository.Interfaces;
using AccomodationServiceLibrary.Core.Service.Interfaces;
using AutoMapper;
using MongoDB.Bson;

namespace AccomodationServiceLibrary.Core.Service
{
    public class AccomodationService : IAccomodationService
    {
        private readonly IAccomodationRepository _accomodationRepository;
        private readonly IMapper _mapper;

        public AccomodationService(IAccomodationRepository accomodationRepository, IMapper mapper)
        {
            _accomodationRepository = accomodationRepository;
            _mapper = mapper;
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

        public async Task UpdateAsync(string id, List<AppointmentDTO> appointments, bool automaticConfirmation)
        {
            var newAppointments = _mapper.Map<List<Appointment>>(appointments);
            newAppointments.ForEach(app => app.Id = ObjectId.GenerateNewId().ToString());
            await _accomodationRepository.UpdateAsync(id, newAppointments, automaticConfirmation);
        }
    }
}
