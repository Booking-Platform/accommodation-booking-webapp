using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace accommodation_service.Core.Model.DTO
{
    public class AccomodationDTO
    {
        public string Name { get; set; }
        [Range(1, 10, ErrorMessage = "The number of guests should be between 1 and 10")]
        public int MinGuestNum { get; set; }
        [Range(1, 10, ErrorMessage = "The number of guests should be between 1 and 10")]
        public int MaxGuestNum { get; set; }
        public AddressDTO Address { get; set; }
        public bool AutomaticConfirmation { get; set; }
        public List<string> Photo { get; set; }
        public List<BenefitDTO> Benefits { get; set; }
    }
}
