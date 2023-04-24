using accommodation_service.Core.Model;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace accommodation_service.Core.Repository.Interfaces
{
    public interface IAppointmentRepository
    {
        Task<string> RemoveAppointmentAsync(string id);
        Task AddAppointmentsAsync(string id, List<Appointment> appointments, bool automaticConfirmation);
    }
}
