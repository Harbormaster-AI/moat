from django.db import models
from aerospaceOnDjango.models.AppointmentStatus import AppointmentStatus

#======================================================================
# 
# Encapsulates data for model MaintenanceAppointment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MaintenanceAppointment Declaration
#======================================================================
class MaintenanceAppointment (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	appointmentDate = models.DateField(null=True)
	aircraft = models.ForeignKey('Aircraft', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	mroFacility = models.ForeignKey('MROFacility', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	workOrder = models.OneToOneField('MaintenanceWorkOrder', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AppointmentStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.appointmentDate
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "MaintenanceAppointment";
    
	def objectType(self):
		return "MaintenanceAppointment";
