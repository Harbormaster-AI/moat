from django.db import models
from advertisingOnDjango.models.DeviceType import DeviceType
from advertisingOnDjango.models.PlatformType import PlatformType
from advertisingOnDjango.models.TargetingOperator import TargetingOperator

#======================================================================
# 
# Encapsulates data for model DeviceCriterion
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DeviceCriterion Declaration
#======================================================================
class DeviceCriterion (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	targetingProfile = models.ForeignKey('TargetingProfile', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	deviceType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DeviceType])
	platformType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in PlatformType])
	operator = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in TargetingOperator])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.deviceType
		str = str + self.platformType
		str = str + self.operator
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "DeviceCriterion";
    
	def objectType(self):
		return "DeviceCriterion";
