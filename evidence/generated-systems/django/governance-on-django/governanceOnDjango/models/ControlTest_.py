from django.db import models
from governanceOnDjango.models.TestType import TestType
from governanceOnDjango.models.ControlEffectiveness import ControlEffectiveness
from governanceOnDjango.models.TestStatus import TestStatus

#======================================================================
# 
# Encapsulates data for model ControlTest_
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ControlTest_ Declaration
#======================================================================
class ControlTest_ (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	testPeriodStart = models.DateField(null=True)
	testPeriodEnd = models.DateField(null=True)
	sampleSize = models.IntegerField(null=True)
	control = models.ForeignKey('Control', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	evidence = models.ManyToManyField('Evidence',  blank=True, related_name='+')
	engagement = models.ForeignKey('AuditEngagement', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	testType = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in TestType])
	effectiveness = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in ControlEffectiveness])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in TestStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.testPeriodStart
		str = str + self.testPeriodEnd
		str = str + self.sampleSize
		str = str + self.testType
		str = str + self.effectiveness
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "ControlTest_";
    
	def objectType(self):
		return "ControlTest_";
