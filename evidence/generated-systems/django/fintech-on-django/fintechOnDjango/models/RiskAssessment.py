from django.db import models
from fintechOnDjango.models.DecisionOutcome import DecisionOutcome

#======================================================================
# 
# Encapsulates data for model RiskAssessment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RiskAssessment Declaration
#======================================================================
class RiskAssessment (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	score = RiskScore
	assessedAt = DateTime
	modelVersion = models.CharField(max_length=200, null=True)
	notes = models.CharField(max_length=200, null=True)
	application = models.ForeignKey('LoanApplication', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	decision = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DecisionOutcome])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.score
		str = str + self.assessedAt
		str = str + self.modelVersion
		str = str + self.notes
		str = str + self.decision
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "RiskAssessment";
    
	def objectType(self):
		return "RiskAssessment";
