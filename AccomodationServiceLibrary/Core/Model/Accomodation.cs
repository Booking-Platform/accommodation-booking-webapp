using DataAnnotationsExtensions;
using MongoDB.Bson;
using MongoDB.Bson.Serialization.Attributes;
using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.Linq;
using System.Text;
using System.Text.Json.Serialization;
using System.Threading.Tasks;

namespace AccomodationServiceLibrary.Core.Model
{
    public class Accomodation
    {
        [BsonId]
        [BsonRepresentation(BsonType.ObjectId)]
        public string Id { get; set; }
        public string Name { get; set; }

        [Range(1, 10, ErrorMessage = "The number of guests should be between 1 and 10")]
        public int MinGuestNum { get; set; }

        [Range(1, 10, ErrorMessage = "The number of guests should be between 1 and 10")]
        public int MaxGuestNum { get; set; }
        public Address Address { get; set; }
        public AccomodationStatus Status { get; set; }
        public bool AutomaticConfirmation { get; set; }
        public List<string> Photo { get; set; }

        private List<Benefit>? _benefits;
        private List<Appointment>? _appointment;

        public List<Benefit>? Benefits
        {
            get => _benefits;
            set => _benefits = value;
        }

        public List<Appointment>? Appointment
        {
            get => _appointment;
            set => _appointment = value;
        }

    }
}
