using AccomodationServiceLibrary.Core.Model;
using AccomodationServiceLibrary.Core.Model.DTO;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace AccomodationServiceLibrary.Core.Service.Interfaces
{
    public interface IAccomodationService
    {
        Task<List<Accomodation>> GetAsync();
        Task<Accomodation?> GetAsync(string id);
        Task CreateAsync(AccomodationDTO newAccomodation);
        Task RemoveAsync(string id);

    }
}
