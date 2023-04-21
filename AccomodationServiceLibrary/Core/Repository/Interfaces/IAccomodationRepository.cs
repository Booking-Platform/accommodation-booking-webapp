using AccomodationServiceLibrary.Core.Model;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace AccomodationServiceLibrary.Core.Repository.Interfaces
{
    public interface IAccomodationRepository
    {
        Task<List<Accomodation>> GetAsync();
        Task<Accomodation?> GetAsync(string id);
        Task CreateAsync(Accomodation newAccomodation);
        Task UpdateAsync(string id, List<Appointment> appointments, bool automaticConfirmation);
        Task RemoveAsync(string id);
    }
}
