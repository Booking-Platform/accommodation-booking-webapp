using AccomodationServiceLibrary.Core.Model;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace AccomodationServiceLibrary.Core.Repository.Interfaces
{
    public interface IAppointmentRepository
    {
        Task<string> RemoveAppointmentAsync(string id);
        Task AddAppointmentsAsync(string id, List<Appointment> appointments, bool automaticConfirmation);
    }
}
