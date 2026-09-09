from django.db import models
from governanceOnDjango.models.RiskCategory import RiskCategory
from governanceOnDjango.models.RiskImpact import RiskImpact
from governanceOnDjango.models.RiskLikelihood import RiskLikelihood
from governanceOnDjango.models.RiskStatus import RiskStatus

#======================================================================
# 
# Encapsulates data for model Risk
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Risk Declaration
#======================================================================
class Risk (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	name = models.CharField(max_length=200, null=True)
	description = models.CharField(max_length=200, null=True)
	inherentRiskScore = models.IntegerField(null=True)
	residualRiskScore = models.IntegerField(null=True)
	organization = models.ForeignKey('Organization', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	controls = models.ManyToManyField('Control',  blank=True, related_name='+')
	assessments = models.ManyToManyField('RiskAssessment',  blank=True, related_name='+')
	issues = models.ManyToManyField('Issue',  blank=True, related_name='+')
	findings = models.ManyToManyField('AuditFinding',  blank=True, related_name='+')
	category = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in RiskCategory])
	impact = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in RiskImpact])
	likelihood = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in RiskLikelihood])
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in RiskStatus])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.name
		str = str + self.description
		str = str + self.inherentRiskScore
		str = str + self.residualRiskScore
		str = str + self.category
		str = str + self.impact
		str = str + self.likelihood
		str = str + self.status
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "Risk";
    
	def objectType(self):
		return "Risk";
