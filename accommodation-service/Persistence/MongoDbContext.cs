using accommodation_service.Core.Model;
using MongoDB.Driver;

namespace accommodation_service.Persistence
{
    public class MongoDbContext
    {

        private readonly IMongoDatabase _database;

        public MongoDbContext(string connectionString, string databaseName)
        {
            var client = new MongoClient(connectionString);
            _database = client.GetDatabase(databaseName);
        }

        public IMongoCollection<Accomodation> Accomodations => _database.GetCollection<Accomodation>("accomodations");
    }

}
