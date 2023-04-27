using accommodation_service.Core.Model;
using accommodation_service.Core.Model.DTO;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace accommodation_service.Core.Service.Interfaces
{
    public interface IAccomodationService
    {
        Task<List<Accomodation>> GetAsync();
        Task<Accomodation?> GetAsync(string id);
        Task CreateAsync(AccomodationDTO newAccomodation);
        Task RemoveAsync(string id);
        Task<List<Accomodation>> SearchAsync(string location, int numGuests, DateTime startDate, DateTime endDate);

    }
}
