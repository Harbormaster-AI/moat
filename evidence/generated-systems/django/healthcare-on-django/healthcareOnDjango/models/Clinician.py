from django.db import models
from healthcareOnDjango.models.ClinicianType import ClinicianType
from healthcareOnDjango.models.ClinicianSpecialty import ClinicianSpecialty

#======================================================================
# 
# Encapsulates data for model Clinician
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Clinician Declaration
#======================================================================
class Clinician (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	firstName = models.CharField(max_length=200, null=True)
	lastName = models.CharField(max_length=200, null=True)
	licenseNumber = models.CharField(max_length=200, null=True)
	careTeams = models.ManyToManyField('CareTeam',  blank=True, related_name='+')
	appointments = models.ManyToManyField('Appointment',  blank=True, related_name='+')
	encounters = models.ManyToManyField('Encounter',  blank=True, related_name='+')
	procedures = models.ManyToManyField('Procedure',  blank=True, related_name='+')
	imagingReports = models.ManyToManyField('ImagingReport',  blank=True, related_name='+')
	clinicianType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ClinicianType])
	specialty = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ClinicianSpecialty])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.firstName
		str = str + self.lastName
		str = str + self.licenseNumber
		str = str + self.clinicianType
		str = str + self.specialty
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Clinician";
    
	def objectType(self):
		return "Clinician";
