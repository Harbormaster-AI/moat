from django.db import models
from healthcareOnDjango.models.AppointmentStatus import AppointmentStatus
from healthcareOnDjango.models.Priority import Priority

#======================================================================
# 
# Encapsulates data for model Appointment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Appointment Declaration
#======================================================================
class Appointment (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	appointmentDate = models.CharField(max_length=64, null=True)
	reason = models.CharField(max_length=200, null=True)
	patient = models.ForeignKey('Patient', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	clinician = models.ForeignKey('Clinician', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	facility = models.ForeignKey('Facility', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	encounter = models.OneToOneField('Encounter', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AppointmentStatus])
	priority = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in Priority])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.appointmentDate
		str = str + self.reason
		str = str + self.status
		str = str + self.priority
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Appointment";
    
	def objectType(self):
		return "Appointment";
