using accommodation_service.Core.Model;
using accommodation_service.Core.Model.DTO;
using accommodation_service.Core.Service.Interfaces;
using Microsoft.AspNetCore.Mvc;

namespace accommodation_service.Controllers
{

    [ApiController]
    [Route("api/[controller]")]
    public class AccomodationController : ControllerBase
    {
        private readonly IAccomodationService _accomodationService;
        private readonly IAppointmentService _appointmentService;

        public AccomodationController(IAccomodationService accomodationService)
        {
            _accomodationService = accomodationService;
        }


        [HttpGet]
        public async Task<List<Accomodation>> Get()
        {
            return await _accomodationService.GetAsync();
        }

        [HttpGet("{id:length(24)}")]
        public async Task<ActionResult<Accomodation>> Get(string id)
        {
            var accomodation = await _accomodationService.GetAsync(id);

            if (accomodation is null)
            {
                return NotFound();
            }

            return accomodation;
        }

        [HttpPost]
        public async Task<IActionResult> Create([FromBody] AccomodationDTO accomodation)
        {
            if (!ModelState.IsValid)
            {
                return BadRequest(ModelState);
            }
            await _accomodationService.CreateAsync(accomodation);

            return CreatedAtAction(nameof(Get), new { name = accomodation.Name }, accomodation);
        }


        [HttpDelete("{id:length(24)}")]
        public async Task<IActionResult> Delete(string id)
        {
            var accomodation = await _accomodationService.GetAsync(id);

            if (accomodation is null)
            {
                return NotFound();
            }

            await _accomodationService.RemoveAsync(id);

            return NoContent();
        }

        [HttpPost("search")]
        public async Task<ActionResult<List<Accomodation>>> SearchAccomodations([FromBody] AccomodationSerachDTO searchParams)
        {
            return await _accomodationService.SearchAsync(searchParams.Location, searchParams.Guests, searchParams.StartDate, searchParams.EndDate);
        }
    }
}
