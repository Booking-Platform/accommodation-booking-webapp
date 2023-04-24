using accommodation_service.Persistence;
using accommodation_service.Core.Model;
using accommodation_service.Core.Repository.Interfaces;
using MongoDB.Driver;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace accommodation_service.Core.Repository
{
    public class AppointmentRepository : IAppointmentRepository
    {

        private readonly MongoDbContext _mongoDbContext;

        public AppointmentRepository(MongoDbContext mongoDbContext)
        {
            _mongoDbContext = mongoDbContext;
        }

        public async Task AddAppointmentsAsync(string id, List<Appointment> appointments, bool automaticConfirmation)
        {
            var filter = Builders<Accomodation>.Filter.Eq(x => x.Id, id);
            var accomodation = await _mongoDbContext.Accomodations.Find(filter).FirstOrDefaultAsync();
            if (accomodation == null) throw new Exception("Accomodation with given id does not exist.");
            else
            {
                if (accomodation.Appointment == null) accomodation.Appointment = new List<Appointment>();
                foreach (var appointment in appointments)
                {
                    if (accomodation.Appointment.Any(a => a.From <= appointment.To && a.To >= appointment.From)) throw new Exception("The appointment overlaps with an existing appointment.");
                    if (appointment.From < DateTime.UtcNow || appointment.To < DateTime.UtcNow) throw new Exception("The appointment cannot be in the past.");
                }
                accomodation.Appointment.AddRange(appointments);
                accomodation.AutomaticConfirmation = automaticConfirmation;
                var update = Builders<Accomodation>.Update.Set(x => x.Appointment, accomodation.Appointment)
                                                          .Set(x => x.AutomaticConfirmation, automaticConfirmation);
                await _mongoDbContext.Accomodations.UpdateOneAsync(filter, update);
            }
        }


        public async Task<string> RemoveAppointmentAsync(string id)
        {
            var filter = Builders<Accomodation>.Filter.ElemMatch(x => x.Appointment, a => a.Id == id);
            var update = Builders<Accomodation>.Update.PullFilter(x => x.Appointment, a => a.Id == id);
            var result = await _mongoDbContext.Accomodations.UpdateOneAsync(filter, update);
            var response = (result.ModifiedCount == 1) ?  $"Appointment with id {id} has been removed." : $"Failed to remove appointment with id {id}.";
            return response;
        }
    }
}
