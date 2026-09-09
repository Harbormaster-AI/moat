from django.db import models
from insuranceOnDjango.models.UnderwritingDecisionType import UnderwritingDecisionType

#======================================================================
# 
# Encapsulates data for model UnderwritingDecision
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class UnderwritingDecision Declaration
#======================================================================
class UnderwritingDecision (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	notes = models.CharField(max_length=200, null=True)
	decisionDate = models.DateField(null=True)
	quote = models.ForeignKey('Quote', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	underwriter = models.ForeignKey('Underwriter', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	decision = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in UnderwritingDecisionType])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.notes
		str = str + self.decisionDate
		str = str + self.decision
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "UnderwritingDecision";
    
	def objectType(self):
		return "UnderwritingDecision";
