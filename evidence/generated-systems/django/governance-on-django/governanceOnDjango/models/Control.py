from django.db import models
from governanceOnDjango.models.ControlType import ControlType
from governanceOnDjango.models.ControlFrequency import ControlFrequency
from governanceOnDjango.models.ControlStatus import ControlStatus

#======================================================================
# 
# Encapsulates data for model Control
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Control Declaration
#======================================================================
class Control (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	objective = models.CharField(max_length=200, null=True)
	ownerDepartment = models.CharField(max_length=200, null=True)
	policy = models.ForeignKey('Policy', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	controlTests = models.ManyToManyField('ControlTest_',  blank=True, related_name='+')
	evidence = models.ManyToManyField('Evidence',  blank=True, related_name='+')
	risks = models.ManyToManyField('Risk',  blank=True, related_name='+')
	obligations = models.ManyToManyField('Obligation',  blank=True, related_name='+')
	procedures = models.ManyToManyField('Procedure',  blank=True, related_name='+')
	issues = models.ManyToManyField('Issue',  blank=True, related_name='+')
	controlType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ControlType])
	frequency = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ControlFrequency])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ControlStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.objective
		str = str + self.ownerDepartment
		str = str + self.controlType
		str = str + self.frequency
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Control";
    
	def objectType(self):
		return "Control";
