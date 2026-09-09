from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from advertisingOnDjango.models.KPI import KPI
from advertisingOnDjango.models.Campaign import Campaign
from advertisingOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model KPI
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class KPIDelegate Declaration
#======================================================================
class KPIDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, kPIId ):
		try:	
			kPI = KPI.objects.filter(id=kPIId)
			return kPI.first();
		except KPI.DoesNotExist:
			raise ProcessingError("KPI with id " + str(kPIId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, kPI):
		for model in serializers.deserialize("json", kPI):
			model.save()
			return model;

	def create(self, kPI):
		kPI.save()
		return kPI;

	def saveFromJson(self, kPI):
		for model in serializers.deserialize("json", kPI):
			model.save()
			return kPI;
	
	def save(self, kPI):
		kPI.save()
		return kPI;
	
	def delete(self, kPIId ):
		errMsg = "Failed to delete KPI from db using id " + str(kPIId)
		
		try:
			kPI = KPI.objects.get(id=kPIId)
			kPI.delete()
			return True
		except KPI.DoesNotExist:
			raise ProcessingError("KPI with id " + str(kPIId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = KPI.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all KPI from db")
		except Exception:
			return None;
		
	def assignCampaign( self, kPIId, campaignId ):
		# lazy importing avoids circular dependencies
		from advertisingOnDjango.delegates.CampaignDelegate import CampaignDelegate

		errMsg = "Failed to assign element " + str(campaignId) + " for Campaign on KPI"

		try:
			# get the KPI from db
			kPI = self.get( kPIId ).first()	
			
			# get the Campaign from db
			campaign = CampaignDelegate().get(campaignId).first();
			
			# assign the Campaign		
			kPI.campaign = campaign
			
			#save it
			kPI.save()

			# reload and return the appropriate version					
			return self.get( kPIId );
		except KPI.DoesNotExist:
			raise ProcessingError(errMsg + " : KPI with id " + str(kPIId) + " does not exist.")
		except Campaign.DoesNotExist:
			raise ProcessingError(errMsg + " : Campaign with id " + str(campaignId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCampaign( self, kPIId ):
		errMsg = "Failed to unassign element " + str(campaignId) + " for Campaign on KPI"

		try:
			# get the KPI from db
			kPI = self.get( kPIId ).first()	
			
			# assign to None for unassignment
			kPI.campaign = None			

			#save it
			kPI.save()

			# reload and return the appropriate version					
			return self.get( kPIId );
		except KPI.DoesNotExist:
			raise ProcessingError(errMsg + " : KPI with id " + str(kPIId) + " does not exist.")
		except Exception:
			return None;
		
