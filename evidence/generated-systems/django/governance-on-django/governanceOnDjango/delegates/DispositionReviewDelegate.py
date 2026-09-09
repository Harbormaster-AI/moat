from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.DispositionReview import DispositionReview
from governanceOnDjango.models.Record_ import Record_
from governanceOnDjango.models.RetentionSchedule import RetentionSchedule
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model DispositionReview
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DispositionReviewDelegate Declaration
#======================================================================
class DispositionReviewDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, dispositionReviewId ):
		try:	
			dispositionReview = DispositionReview.objects.filter(id=dispositionReviewId)
			return dispositionReview.first();
		except DispositionReview.DoesNotExist:
			raise ProcessingError("DispositionReview with id " + str(dispositionReviewId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, dispositionReview):
		for model in serializers.deserialize("json", dispositionReview):
			model.save()
			return model;

	def create(self, dispositionReview):
		dispositionReview.save()
		return dispositionReview;

	def saveFromJson(self, dispositionReview):
		for model in serializers.deserialize("json", dispositionReview):
			model.save()
			return dispositionReview;
	
	def save(self, dispositionReview):
		dispositionReview.save()
		return dispositionReview;
	
	def delete(self, dispositionReviewId ):
		errMsg = "Failed to delete DispositionReview from db using id " + str(dispositionReviewId)
		
		try:
			dispositionReview = DispositionReview.objects.get(id=dispositionReviewId)
			dispositionReview.delete()
			return True
		except DispositionReview.DoesNotExist:
			raise ProcessingError("DispositionReview with id " + str(dispositionReviewId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = DispositionReview.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all DispositionReview from db")
		except Exception:
			return None;
		
	def assignRecord( self, dispositionReviewId, recordId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.Record_Delegate import Record_Delegate

		errMsg = "Failed to assign element " + str(recordId) + " for Record on DispositionReview"

		try:
			# get the DispositionReview from db
			dispositionReview = self.get( dispositionReviewId ).first()	
			
			# get the Record_ from db
			record_ = Record_Delegate().get(recordId).first();
			
			# assign the Record		
			dispositionReview.record = record_
			
			#save it
			dispositionReview.save()

			# reload and return the appropriate version					
			return self.get( dispositionReviewId );
		except DispositionReview.DoesNotExist:
			raise ProcessingError(errMsg + " : DispositionReview with id " + str(dispositionReviewId) + " does not exist.")
		except Record_.DoesNotExist:
			raise ProcessingError(errMsg + " : Record_ with id " + str(recordId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignRecord( self, dispositionReviewId ):
		errMsg = "Failed to unassign element " + str(recordId) + " for Record on DispositionReview"

		try:
			# get the DispositionReview from db
			dispositionReview = self.get( dispositionReviewId ).first()	
			
			# assign to None for unassignment
			dispositionReview.record_ = None			

			#save it
			dispositionReview.save()

			# reload and return the appropriate version					
			return self.get( dispositionReviewId );
		except DispositionReview.DoesNotExist:
			raise ProcessingError(errMsg + " : DispositionReview with id " + str(dispositionReviewId) + " does not exist.")
		except Exception:
			return None;
		
	def assignRetentionSchedule( self, dispositionReviewId, retentionScheduleId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.RetentionScheduleDelegate import RetentionScheduleDelegate

		errMsg = "Failed to assign element " + str(retentionScheduleId) + " for RetentionSchedule on DispositionReview"

		try:
			# get the DispositionReview from db
			dispositionReview = self.get( dispositionReviewId ).first()	
			
			# get the RetentionSchedule from db
			retentionSchedule = RetentionScheduleDelegate().get(retentionScheduleId).first();
			
			# assign the RetentionSchedule		
			dispositionReview.retentionSchedule = retentionSchedule
			
			#save it
			dispositionReview.save()

			# reload and return the appropriate version					
			return self.get( dispositionReviewId );
		except DispositionReview.DoesNotExist:
			raise ProcessingError(errMsg + " : DispositionReview with id " + str(dispositionReviewId) + " does not exist.")
		except RetentionSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : RetentionSchedule with id " + str(retentionScheduleId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignRetentionSchedule( self, dispositionReviewId ):
		errMsg = "Failed to unassign element " + str(retentionScheduleId) + " for RetentionSchedule on DispositionReview"

		try:
			# get the DispositionReview from db
			dispositionReview = self.get( dispositionReviewId ).first()	
			
			# assign to None for unassignment
			dispositionReview.retentionSchedule = None			

			#save it
			dispositionReview.save()

			# reload and return the appropriate version					
			return self.get( dispositionReviewId );
		except DispositionReview.DoesNotExist:
			raise ProcessingError(errMsg + " : DispositionReview with id " + str(dispositionReviewId) + " does not exist.")
		except Exception:
			return None;
		
