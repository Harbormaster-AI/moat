from django.db import models
from hrOnDjango.models.RequisitionStatus import RequisitionStatus
from hrOnDjango.models.RequisitionPriority import RequisitionPriority

#======================================================================
# 
# Encapsulates data for model JobRequisition
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class JobRequisition Declaration
#======================================================================
class JobRequisition (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	requisitionNumber = models.CharField(max_length=200, null=True)
	title = models.CharField(max_length=200, null=True)
	openings = models.IntegerField(null=True)
	targetStartDate = models.DateField(null=True)
	department = models.ForeignKey('Department', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	hiringManager = models.ForeignKey('Employee', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	recruiter = models.ForeignKey('Employee', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	jobProfile = models.ForeignKey('JobProfile', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	candidates = models.ManyToManyField('Candidate',  blank=True, related_name='+')
	interviews = models.ManyToManyField('Interview',  blank=True, related_name='+')
	offers = models.ManyToManyField('Offer',  blank=True, related_name='+')
	status = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in RequisitionStatus])
	priority = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in RequisitionPriority])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.requisitionNumber
		str = str + self.title
		str = str + self.openings
		str = str + self.targetStartDate
		str = str + self.status
		str = str + self.priority
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "JobRequisition";
    
	def objectType(self):
		return "JobRequisition";
