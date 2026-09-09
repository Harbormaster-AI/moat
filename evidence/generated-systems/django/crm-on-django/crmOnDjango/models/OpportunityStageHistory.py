from django.db import models
from crmOnDjango.models.OpportunityStage import OpportunityStage
from crmOnDjango.models.OpportunityStage import OpportunityStage

#======================================================================
# 
# Encapsulates data for model OpportunityStageHistory
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OpportunityStageHistory Declaration
#======================================================================
class OpportunityStageHistory (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	changedAt = models.CharField(max_length=64, null=True)
	comment = models.CharField(max_length=200, null=True)
	opportunity = models.ForeignKey('Opportunity', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	changedBy = models.ForeignKey('User', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	fromStage = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in OpportunityStage])
	toStage = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in OpportunityStage])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.changedAt
		str = str + self.comment
		str = str + self.fromStage
		str = str + self.toStage
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "OpportunityStageHistory";
    
	def objectType(self):
		return "OpportunityStageHistory";
