from django.db import models
from advertisingOnDjango.models.CreativeApprovalStatus import CreativeApprovalStatus

#======================================================================
# 
# Encapsulates data for model CreativeApproval
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CreativeApproval Declaration
#======================================================================
class CreativeApproval (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	reviewer = models.CharField(max_length=200, null=True)
	reviewedAt = models.DateField(null=True)
	creativeAsset = models.ForeignKey('CreativeAsset', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	publisher = models.ForeignKey('Publisher', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in CreativeApprovalStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.reviewer
		str = str + self.reviewedAt
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "CreativeApproval";
    
	def objectType(self):
		return "CreativeApproval";
