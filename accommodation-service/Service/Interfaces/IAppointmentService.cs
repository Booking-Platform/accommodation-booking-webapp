using accommodation_service.Core.Model.DTO;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace accommodation_service.Core.Service.Interfaces
{
    public interface IAppointmentService
    {

        Task<string> RemoveAppointmentAsync(string id);
        Task AddAppointmentsAsync(string id, List<AppointmentDTO> appointments, bool automaticConfirmation);
    }
}
