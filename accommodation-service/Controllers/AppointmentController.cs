using accommodation_service.Core.Model.DTO;
using accommodation_service.Core.Service.Interfaces;
using Microsoft.AspNetCore.Mvc;

namespace accommodation_service.Controllers
{

    [ApiController]
    [Route("api/[controller]")]
    public class AppointmentController : ControllerBase
    {


        private readonly IAppointmentService _appointmentService;
        private readonly IAccomodationService _accomodationService;

        public AppointmentController(IAppointmentService appointmentService, IAccomodationService accomodationService)
        {
            _appointmentService = appointmentService;
            _accomodationService = accomodationService;
        }

        [HttpPut("{id:length(24)}")]
        public async Task<IActionResult> AddAppointmentsToAccomodation(string id, List<AppointmentDTO> appointments, bool automaticConfirmation)
        {
            var accomodation = await _accomodationService.GetAsync(id);
            if (accomodation is null)
                return NotFound();
            try
            {
                await _appointmentService.AddAppointmentsAsync(id, appointments, automaticConfirmation);
            }
            catch (Exception ex)
            {
                return BadRequest(ex.Message);
            }
            return Ok();
        }


        [HttpDelete("{id:length(24)}")]
        public async Task<string> RemoveAppointmentFromAccomodation(string id)
        {
            /* var accomodation = await _accomodationService.GetAsync(id);

             if (accomodation is null)
             {
                 return NotFound();
             }*/
            return await _appointmentService.RemoveAppointmentAsync(id); ;
        }

    }
}
