using AccomodationServiceLibrary.Core.Model.DTO;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace AccomodationServiceLibrary.Core.Service.Interfaces
{
    public interface IAppointmentService
    {

        Task<string> RemoveAppointmentAsync(string id);
        Task AddAppointmentsAsync(string id, List<AppointmentDTO> appointments, bool automaticConfirmation);
    }
}
