from django.db import models
from aerospaceOnDjango.models.ProgramStatus import ProgramStatus

#======================================================================
# 
# Encapsulates data for model AircraftProgram
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AircraftProgram Declaration
#======================================================================
class AircraftProgram (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	programCode = models.CharField(max_length=200, null=True)
	entryIntoServiceYear = models.IntegerField(null=True)
	manufacturer = models.ForeignKey('AerospaceManufacturer', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	aircraftFamilies = models.ManyToManyField('AircraftFamily',  blank=True, related_name='+')
	typeCertificate = models.OneToOneField('TypeCertificate', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	keySuppliers = models.ManyToManyField('Supplier',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ProgramStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.programCode
		str = str + self.entryIntoServiceYear
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "AircraftProgram";
    
	def objectType(self):
		return "AircraftProgram";
