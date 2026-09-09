from django.db import models
from hrOnDjango.models.AcknowledgementStatus import AcknowledgementStatus

#======================================================================
# 
# Encapsulates data for model PolicyAcknowledgement
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PolicyAcknowledgement Declaration
#======================================================================
class PolicyAcknowledgement (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	acknowledgementDate = models.DateField(null=True)
	policy = models.ForeignKey('Policy', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	employee = models.ForeignKey('Employee', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in AcknowledgementStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.acknowledgementDate
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "PolicyAcknowledgement";
    
	def objectType(self):
		return "PolicyAcknowledgement";
