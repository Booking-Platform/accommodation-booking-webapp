using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace accommodation_service.Core.Model.DTO
{
    public class AppointmentDTO
    {
        public DateTime From { get; set; }
        public DateTime To { get; set; }
        public double Price { get; set; }
        public bool PerGuest { get; set; }
    }
}
