using accommodation_service.Core.Model;
using accommodation_service.Core.Repository.Interfaces;
using accommodation_service.Persistence;
using MongoDB.Bson;
using MongoDB.Driver;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace accommodation_service.Core.Repository
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

        public async Task<List<Accomodation>> SearchAsync(string location, int numGuests, DateTime startDate, DateTime endDate)
        {
            var filter = Builders<Accomodation>.Filter.And(
                Builders<Accomodation>.Filter.Regex(x => x.Address.City, new BsonRegularExpression(location, "i")),
                Builders<Accomodation>.Filter.Gte(x => x.MaxGuestNum, numGuests),
                Builders<Accomodation>.Filter.Lte(x => x.MinGuestNum, numGuests),
                Builders<Accomodation>.Filter.ElemMatch(x => x.Appointment,
                    x => x.From <= endDate && x.To >= startDate && x.Status == 0));
            return await _mongoDbContext.Accomodations.Find(filter).ToListAsync();
        }


    }
}
