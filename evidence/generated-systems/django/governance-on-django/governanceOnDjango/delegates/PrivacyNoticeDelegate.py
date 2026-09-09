from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.PrivacyNotice import PrivacyNotice
from governanceOnDjango.models.DataProcessingActivity import DataProcessingActivity
from governanceOnDjango.models.Organization import Organization
from governanceOnDjango.models.Consent import Consent
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model PrivacyNotice
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PrivacyNoticeDelegate Declaration
#======================================================================
class PrivacyNoticeDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, privacyNoticeId ):
		try:	
			privacyNotice = PrivacyNotice.objects.filter(id=privacyNoticeId)
			return privacyNotice.first();
		except PrivacyNotice.DoesNotExist:
			raise ProcessingError("PrivacyNotice with id " + str(privacyNoticeId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, privacyNotice):
		for model in serializers.deserialize("json", privacyNotice):
			model.save()
			return model;

	def create(self, privacyNotice):
		privacyNotice.save()
		return privacyNotice;

	def saveFromJson(self, privacyNotice):
		for model in serializers.deserialize("json", privacyNotice):
			model.save()
			return privacyNotice;
	
	def save(self, privacyNotice):
		privacyNotice.save()
		return privacyNotice;
	
	def delete(self, privacyNoticeId ):
		errMsg = "Failed to delete PrivacyNotice from db using id " + str(privacyNoticeId)
		
		try:
			privacyNotice = PrivacyNotice.objects.get(id=privacyNoticeId)
			privacyNotice.delete()
			return True
		except PrivacyNotice.DoesNotExist:
			raise ProcessingError("PrivacyNotice with id " + str(privacyNoticeId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = PrivacyNotice.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all PrivacyNotice from db")
		except Exception:
			return None;
		
	def assignOrganization( self, privacyNoticeId, organizationId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on PrivacyNotice"

		try:
			# get the PrivacyNotice from db
			privacyNotice = self.get( privacyNoticeId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			privacyNotice.organization = organization
			
			#save it
			privacyNotice.save()

			# reload and return the appropriate version					
			return self.get( privacyNoticeId );
		except PrivacyNotice.DoesNotExist:
			raise ProcessingError(errMsg + " : PrivacyNotice with id " + str(privacyNoticeId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, privacyNoticeId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on PrivacyNotice"

		try:
			# get the PrivacyNotice from db
			privacyNotice = self.get( privacyNoticeId ).first()	
			
			# assign to None for unassignment
			privacyNotice.organization = None			

			#save it
			privacyNotice.save()

			# reload and return the appropriate version					
			return self.get( privacyNoticeId );
		except PrivacyNotice.DoesNotExist:
			raise ProcessingError(errMsg + " : PrivacyNotice with id " + str(privacyNoticeId) + " does not exist.")
		except Exception:
			return None;
		
	def addProcessingActivities( self, privacyNoticeId, processingActivitiesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataProcessingActivityDelegate import DataProcessingActivityDelegate

		errMsg = "Failed to add elements " + str(processingActivitiesIds) + " for ProcessingActivities on PrivacyNotice"

		try:
			# get the PrivacyNotice
			privacyNotice = self.get( privacyNoticeId ).first()
				
			# split on a comma with no spaces
			idList = processingActivitiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataProcessingActivity		
				dataProcessingActivity = DataProcessingActivityDelegate().get(id).first();	
				# add the DataProcessingActivity
				privacyNotice.processingActivities.add(dataProcessingActivity)
				
			# save it		
			privacyNotice.save()
			
			# reload and return the appropriate version
			return self.get( privacyNoticeId );
		except PrivacyNotice.DoesNotExist:
			raise ProcessingError(errMsg + " : PrivacyNotice with id " + str(privacyNoticeId) + " does not exist.")
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProcessingActivity does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeProcessingActivities( self, privacyNoticeId, processingActivitiesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.DataProcessingActivityDelegate import DataProcessingActivityDelegate

		errMsg = "Failed to remove elements " + str(processingActivitiesIds) + " for ProcessingActivities on PrivacyNotice"

		try:
			# get the PrivacyNotice
			privacyNotice = self.get( privacyNoticeId ).first()
				
			# split on a comma with no spaces
			idList = processingActivitiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataProcessingActivity		
				dataProcessingActivity = DataProcessingActivityDelegate().get(id).first();	
				# add the DataProcessingActivity
				privacyNotice.processingActivities.remove(dataProcessingActivity)
				
			# save it		
			privacyNotice.save()
			
			# reload and return the appropriate version
			return self.get( privacyNoticeId );
		except PrivacyNotice.DoesNotExist:
			raise ProcessingError(errMsg + " : PrivacyNotice with id " + str(privacyNoticeId) + " does not exist.")
		except DataProcessingActivity.DoesNotExist:
			raise ProcessingError(errMsg + " : DataProcessingActivity does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addConsents( self, privacyNoticeId, consentsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ConsentDelegate import ConsentDelegate

		errMsg = "Failed to add elements " + str(consentsIds) + " for Consents on PrivacyNotice"

		try:
			# get the PrivacyNotice
			privacyNotice = self.get( privacyNoticeId ).first()
				
			# split on a comma with no spaces
			idList = consentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Consent		
				consent = ConsentDelegate().get(id).first();	
				# add the Consent
				privacyNotice.consents.add(consent)
				
			# save it		
			privacyNotice.save()
			
			# reload and return the appropriate version
			return self.get( privacyNoticeId );
		except PrivacyNotice.DoesNotExist:
			raise ProcessingError(errMsg + " : PrivacyNotice with id " + str(privacyNoticeId) + " does not exist.")
		except Consent.DoesNotExist:
			raise ProcessingError(errMsg + " : Consent does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeConsents( self, privacyNoticeId, consentsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ConsentDelegate import ConsentDelegate

		errMsg = "Failed to remove elements " + str(consentsIds) + " for Consents on PrivacyNotice"

		try:
			# get the PrivacyNotice
			privacyNotice = self.get( privacyNoticeId ).first()
				
			# split on a comma with no spaces
			idList = consentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Consent		
				consent = ConsentDelegate().get(id).first();	
				# add the Consent
				privacyNotice.consents.remove(consent)
				
			# save it		
			privacyNotice.save()
			
			# reload and return the appropriate version
			return self.get( privacyNoticeId );
		except PrivacyNotice.DoesNotExist:
			raise ProcessingError(errMsg + " : PrivacyNotice with id " + str(privacyNoticeId) + " does not exist.")
		except Consent.DoesNotExist:
			raise ProcessingError(errMsg + " : Consent does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
