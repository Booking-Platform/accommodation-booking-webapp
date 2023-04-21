using AccomodationServiceLibrary.Core.Model;
using AccomodationServiceLibrary.Core.Model.DTO;
using AccomodationServiceLibrary.Core.Repository.Interfaces;
using AccomodationServiceLibrary.Persistence;
using MongoDB.Driver;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace AccomodationServiceLibrary.Core.Repository
{
    public class AccomodationRepository : IAccomodationRepository
    {
        private readonly MongoDbContext _mongoDbContext;

        public AccomodationRepository(MongoDbContext mongoDbContext)
        {
            _mongoDbContext = mongoDbContext;
        }

        public async Task<List<Accomodation>> GetAsync()
        {
            return await _mongoDbContext.Accomodations.Find(_ => true).ToListAsync();
        }

        public async Task<Accomodation?> GetAsync(string id)
        {
            return await _mongoDbContext.Accomodations.Find(x => x.Id == id).FirstOrDefaultAsync();
        }
        public async Task CreateAsync(Accomodation newAccomodation)
        {
            await _mongoDbContext.Accomodations.InsertOneAsync(newAccomodation);
        }
        public async Task RemoveAsync(string id)
        {
            await _mongoDbContext.Accomodations.DeleteOneAsync(x => x.Id == id);
        }
    }
}
