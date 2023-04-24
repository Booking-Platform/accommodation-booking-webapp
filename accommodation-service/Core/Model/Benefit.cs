using MongoDB.Bson;
using MongoDB.Bson.Serialization.Attributes;

namespace accommodation_service.Core.Model
{
    public class Benefit
    {
        [BsonId]
        [BsonRepresentation(BsonType.ObjectId)]
        public string Id { get; set; }
        public string Name { get; set; }
    }
}
