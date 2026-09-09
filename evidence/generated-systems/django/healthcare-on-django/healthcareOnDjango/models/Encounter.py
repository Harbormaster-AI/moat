from django.db import models
from healthcareOnDjango.models.EncounterStatus import EncounterStatus
from healthcareOnDjango.models.EncounterType import EncounterType

#======================================================================
# 
# Encapsulates data for model Encounter
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Encounter Declaration
#======================================================================
class Encounter (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	encounterNumber = models.CharField(max_length=200, null=True)
	startDateTime = models.CharField(max_length=64, null=True)
	endDateTime = models.CharField(max_length=64, null=True)
	patient = models.ForeignKey('Patient', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	clinician = models.ForeignKey('Clinician', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	facility = models.ForeignKey('Facility', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	appointment = models.OneToOneField('Appointment', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	diagnoses = models.ManyToManyField('Diagnosis',  blank=True, related_name='+')
	procedures = models.ManyToManyField('Procedure',  blank=True, related_name='+')
	observations = models.ManyToManyField('Observation',  blank=True, related_name='+')
	orders = models.ManyToManyField('ClinicalOrder',  blank=True, related_name='+')
	admission = models.OneToOneField('Admission', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	discharge = models.OneToOneField('Discharge', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in EncounterStatus])
	encounterType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in EncounterType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.encounterNumber
		str = str + self.startDateTime
		str = str + self.endDateTime
		str = str + self.status
		str = str + self.encounterType
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Encounter";
    
	def objectType(self):
		return "Encounter";
