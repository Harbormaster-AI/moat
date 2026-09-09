from django.db import models
from healthcareOnDjango.models.DeviceType import DeviceType
from healthcareOnDjango.models.DeviceConnectivityStatus import DeviceConnectivityStatus

#======================================================================
# 
# Encapsulates data for model MedicalDevice
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class MedicalDevice Declaration
#======================================================================
class MedicalDevice (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	udi = models.CharField(max_length=200, null=True)
	manufacturer = models.CharField(max_length=200, null=True)
	patient = models.ForeignKey('Patient', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	observations = models.ManyToManyField('Observation',  blank=True, related_name='+')
	softwareUpdates = models.ManyToManyField('SoftwareUpdate',  blank=True, related_name='+')
	deviceType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DeviceType])
	connectivityStatus = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DeviceConnectivityStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.udi
		str = str + self.manufacturer
		str = str + self.deviceType
		str = str + self.connectivityStatus
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "MedicalDevice";
    
	def objectType(self):
		return "MedicalDevice";
