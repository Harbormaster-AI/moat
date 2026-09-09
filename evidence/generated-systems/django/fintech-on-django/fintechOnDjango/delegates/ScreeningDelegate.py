from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.Screening import Screening
from fintechOnDjango.models.KYCProfile import KYCProfile
from fintechOnDjango.models.ComplianceAlert import ComplianceAlert
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Screening
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ScreeningDelegate Declaration
#======================================================================
class ScreeningDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, screeningId ):
		try:	
			screening = Screening.objects.filter(id=screeningId)
			return screening.first();
		except Screening.DoesNotExist:
			raise ProcessingError("Screening with id " + str(screeningId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, screening):
		for model in serializers.deserialize("json", screening):
			model.save()
			return model;

	def create(self, screening):
		screening.save()
		return screening;

	def saveFromJson(self, screening):
		for model in serializers.deserialize("json", screening):
			model.save()
			return screening;
	
	def save(self, screening):
		screening.save()
		return screening;
	
	def delete(self, screeningId ):
		errMsg = "Failed to delete Screening from db using id " + str(screeningId)
		
		try:
			screening = Screening.objects.get(id=screeningId)
			screening.delete()
			return True
		except Screening.DoesNotExist:
			raise ProcessingError("Screening with id " + str(screeningId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Screening.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Screening from db")
		except Exception:
			return None;
		
	def assignKycProfile( self, screeningId, kycProfileId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.KYCProfileDelegate import KYCProfileDelegate

		errMsg = "Failed to assign element " + str(kycProfileId) + " for KycProfile on Screening"

		try:
			# get the Screening from db
			screening = self.get( screeningId ).first()	
			
			# get the KYCProfile from db
			kYCProfile = KYCProfileDelegate().get(kycProfileId).first();
			
			# assign the KycProfile		
			screening.kycProfile = kYCProfile
			
			#save it
			screening.save()

			# reload and return the appropriate version					
			return self.get( screeningId );
		except Screening.DoesNotExist:
			raise ProcessingError(errMsg + " : Screening with id " + str(screeningId) + " does not exist.")
		except KYCProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : KYCProfile with id " + str(kycProfileId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignKycProfile( self, screeningId ):
		errMsg = "Failed to unassign element " + str(kycProfileId) + " for KycProfile on Screening"

		try:
			# get the Screening from db
			screening = self.get( screeningId ).first()	
			
			# assign to None for unassignment
			screening.kYCProfile = None			

			#save it
			screening.save()

			# reload and return the appropriate version					
			return self.get( screeningId );
		except Screening.DoesNotExist:
			raise ProcessingError(errMsg + " : Screening with id " + str(screeningId) + " does not exist.")
		except Exception:
			return None;
		
	def addAlerts( self, screeningId, alertsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.ComplianceAlertDelegate import ComplianceAlertDelegate

		errMsg = "Failed to add elements " + str(alertsIds) + " for Alerts on Screening"

		try:
			# get the Screening
			screening = self.get( screeningId ).first()
				
			# split on a comma with no spaces
			idList = alertsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ComplianceAlert		
				complianceAlert = ComplianceAlertDelegate().get(id).first();	
				# add the ComplianceAlert
				screening.alerts.add(complianceAlert)
				
			# save it		
			screening.save()
			
			# reload and return the appropriate version
			return self.get( screeningId );
		except Screening.DoesNotExist:
			raise ProcessingError(errMsg + " : Screening with id " + str(screeningId) + " does not exist.")
		except ComplianceAlert.DoesNotExist:
			raise ProcessingError(errMsg + " : ComplianceAlert does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAlerts( self, screeningId, alertsIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.ComplianceAlertDelegate import ComplianceAlertDelegate

		errMsg = "Failed to remove elements " + str(alertsIds) + " for Alerts on Screening"

		try:
			# get the Screening
			screening = self.get( screeningId ).first()
				
			# split on a comma with no spaces
			idList = alertsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ComplianceAlert		
				complianceAlert = ComplianceAlertDelegate().get(id).first();	
				# add the ComplianceAlert
				screening.alerts.remove(complianceAlert)
				
			# save it		
			screening.save()
			
			# reload and return the appropriate version
			return self.get( screeningId );
		except Screening.DoesNotExist:
			raise ProcessingError(errMsg + " : Screening with id " + str(screeningId) + " does not exist.")
		except ComplianceAlert.DoesNotExist:
			raise ProcessingError(errMsg + " : ComplianceAlert does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
