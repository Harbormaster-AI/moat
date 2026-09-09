from django.db import models
from governanceOnDjango.models.DispositionOutcome import DispositionOutcome

#======================================================================
# 
# Encapsulates data for model DispositionReview
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DispositionReview Declaration
#======================================================================
class DispositionReview (models.Model):

#======================================================================
# attribute declarations
#======================================================================
	reviewDate = models.DateField(null=True)
	reviewer = models.CharField(max_length=200, null=True)
	notes = models.CharField(max_length=200, null=True)
	record = models.ForeignKey('Record_', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	retentionSchedule = models.ForeignKey('RetentionSchedule', on_delete=models.CASCADE, null=True, blank=True, related_name='+')
	outcome = models.CharField(max_length=64, null=True, choices=[(tag.name, tag.value) for tag in DispositionOutcome])

#======================================================================
# function declarations
#======================================================================
	def toString(self):
		str = ""
		str = str + self.reviewDate
		str = str + self.reviewer
		str = str + self.notes
		str = str + self.outcome
		return str;
    
	def __str__(self):
		return self.toString();

	def identity(self):
		return "DispositionReview";
    
	def objectType(self):
		return "DispositionReview";
