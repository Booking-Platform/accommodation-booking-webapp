using AccomodationServiceLibrary.Core.Model;
using AccomodationServiceLibrary.Core.Model.DTO;
using AccomodationServiceLibrary.Core.Service;
using AccomodationServiceLibrary.Core.Service.Interfaces;
using Microsoft.AspNetCore.Mvc;

namespace accommodation_service.Controllers
{

    [ApiController]
    [Route("api/[controller]")]
    public class AccomodationController : ControllerBase
    {
        private readonly IAccomodationService _accomodationService;

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

            return CreatedAtAction(nameof(Get), new { id = accomodation.Name }, accomodation);
        }

        [HttpPut("{id:length(24)}")]
        public async Task<IActionResult> Update(string id, List<AppointmentDTO> appointments, bool automaticConfirmation)
        {
            var accomodation = await _accomodationService.GetAsync(id);
    
            if (accomodation is null)
            {
                return NotFound();
            }

            await _accomodationService.UpdateAsync(id, appointments, automaticConfirmation);

            return NoContent();
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

    }
}
