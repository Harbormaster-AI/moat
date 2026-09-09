from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.Consent import Consent
from governanceOnDjango.models.DataProcessingActivity import DataProcessingActivity
from governanceOnDjango.models.PrivacyNotice import PrivacyNotice
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Consent
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ConsentDelegate Declaration
#======================================================================
class ConsentDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, consentId ):
		try:	
			consent = Consent.objects.filter(id=consentId)
			return consent.first();
		except Consent.DoesNotExist:
			raise ProcessingError("Consent with id " + str(consentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, consent):
		for model in serializers.deserialize("json", consent):
			model.save()
			return model;

	def create(self, consent):
		consent.save()
		return consent;

	def saveFromJson(self, consent):
		for model in serializers.deserialize("json", consent):
			model.save()
			return consent;
	
	def save(self, consent):
		consent.save()
		return consent;
	
	def delete(self, consentId ):
		errMsg = "Failed to delete Consent from db using id " + str(consentId)
		
		try:
			consent = Consent.objects.get(id=consentId)
			consent.delete()
			return True
		except Consent.DoesNotExist:
			raise ProcessingError("Consent with id " + str(consentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Consent.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Consent from db")
		except Exception:
			return None;
		
	def assignPrivacyNotice( self, consentId, privacyNoticeId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.PrivacyNoticeDelegate import PrivacyNoticeDelegate

		errMsg = "Failed to assign element " + str(privacyNoticeId) + " for PrivacyNotice on Consent"

		try:
			# get the Consent from db
			consent = self.get( consentId ).first()	
			
			# get the PrivacyNotice from db
			privacyNotice = PrivacyNoticeDelegate().get(privacyNoticeId).first();
			
			# assign the PrivacyNotice		
			consent.privacyNotice = privacyNotice
			
			#save it
			consent.save()

			# reload and return the appropriate version					
			return self.get( consentId );
		except Consent.DoesNotExist:
			raise ProcessingError(errMsg + " : Consent with id " + str(consentId) + " does not exist.")
		except PrivacyNotice.DoesNotExist:
			raise ProcessingError(errMsg + " : PrivacyNotice with id " + str(privacyNoticeId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPrivacyNotice( self, consentId ):
		errMsg = "Failed to unassign element " + str(privacyNoticeId) + " for PrivacyNotice on Consent"

		try:
			# get the Consent from db
			consent = self.get( consentId ).first()	
			
			# assign to None for unassignment
			consent.privacyNotice = None			

			#save it
			consent.save()

			# reload and return the appropriate version					
			return self.get( consentId );
		except Consent.DoesNotExist:
			raise ProcessingError(errMsg + " : Consent with id " + str(consentId) + " does not exist.")
		except Exception:
			return None;
		
	def addProcessingActivities( self, consentId, processingActivitiesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataProcessingActivityDelegate import DataProcessingActivityDelegate

		errMsg = "Failed to add elements " + str(processingActivitiesIds) + " for ProcessingActivities on Consent"

		try:
			# get the Consent
			consent = self.get( consentId ).first()
				
			# split on a comma with no spaces
			idList = processingActivitiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataProcessingActivity		
				dataProcessingActivity = DataProcessingActivityDelegate().get(id).first();	
				# add the DataProcessingActivity
				consent.processingActivities.add(dataProcessingActivity)
				
			# save it		
			consent.save()
			
			# reload and return the appropriate version
			return self.get( consentId );
		except Consent.DoesNotExist:
			raise ProcessingError(errMsg + " : Consent with id " + str(consentId) + " does not exist.")
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProcessingActivity does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeProcessingActivities( self, consentId, processingActivitiesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataProcessingActivityDelegate import DataProcessingActivityDelegate

		errMsg = "Failed to remove elements " + str(processingActivitiesIds) + " for ProcessingActivities on Consent"

		try:
			# get the Consent
			consent = self.get( consentId ).first()
				
			# split on a comma with no spaces
			idList = processingActivitiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataProcessingActivity		
				dataProcessingActivity = DataProcessingActivityDelegate().get(id).first();	
				# add the DataProcessingActivity
				consent.processingActivities.remove(dataProcessingActivity)
				
			# save it		
			consent.save()
			
			# reload and return the appropriate version
			return self.get( consentId );
		except Consent.DoesNotExist:
			raise ProcessingError(errMsg + " : Consent with id " + str(consentId) + " does not exist.")
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProcessingActivity does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
