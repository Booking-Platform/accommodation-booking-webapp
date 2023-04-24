
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using AutoMapper;
using MongoDB.Bson;
using accommodation_service.Core.Repository.Interfaces;
using accommodation_service.Core.Model;
using accommodation_service.Core.Model.DTO;
using accommodation_service.Core.Service.Interfaces;

namespace accommodation_service.Core.Service
{
    public class AppointmentService : IAppointmentService
    {

        private readonly IAppointmentRepository _appointmentRepository;
        private readonly IMapper _mapper;
        private readonly IAppointmentService _appointmentService;

        public AppointmentService(IAppointmentRepository appointmentRepository, IMapper mapper)
        {
            _appointmentRepository = appointmentRepository;
            _mapper = mapper;
        }

        public async Task AddAppointmentsAsync(string id, List<AppointmentDTO> appointments, bool automaticConfirmation)
        {
            var newAppointments = _mapper.Map<List<Appointment>>(appointments);
            newAppointments.ForEach(app => app.Id = ObjectId.GenerateNewId().ToString());
            await _appointmentRepository.AddAppointmentsAsync(id, newAppointments, automaticConfirmation);
        }

        public async Task<string> RemoveAppointmentAsync(string id)
        {
            return await _appointmentRepository.RemoveAppointmentAsync(id);
        }
    }
}
