using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace accommodation_service.Core.Model.DTO
{
    public class AccomodationSerachDTO
    { 
            public string Location { get; set; }
            public int Guests { get; set; }
            public DateTime StartDate { get; set; }
            public DateTime EndDate { get; set; }
    }
}
