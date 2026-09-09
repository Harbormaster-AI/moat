from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.SalesRegion import SalesRegion
from aerospaceOnDjango.models.Operator import Operator
from aerospaceOnDjango.models.SalesCampaign import SalesCampaign
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model SalesRegion
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SalesRegionDelegate Declaration
#======================================================================
class SalesRegionDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, salesRegionId ):
		try:	
			salesRegion = SalesRegion.objects.filter(id=salesRegionId)
			return salesRegion.first();
		except SalesRegion.DoesNotExist:
			raise ProcessingError("SalesRegion with id " + str(salesRegionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, salesRegion):
		for model in serializers.deserialize("json", salesRegion):
			model.save()
			return model;

	def create(self, salesRegion):
		salesRegion.save()
		return salesRegion;

	def saveFromJson(self, salesRegion):
		for model in serializers.deserialize("json", salesRegion):
			model.save()
			return salesRegion;
	
	def save(self, salesRegion):
		salesRegion.save()
		return salesRegion;
	
	def delete(self, salesRegionId ):
		errMsg = "Failed to delete SalesRegion from db using id " + str(salesRegionId)
		
		try:
			salesRegion = SalesRegion.objects.get(id=salesRegionId)
			salesRegion.delete()
			return True
		except SalesRegion.DoesNotExist:
			raise ProcessingError("SalesRegion with id " + str(salesRegionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = SalesRegion.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all SalesRegion from db")
		except Exception:
			return None;
		
	def addOperators( self, salesRegionId, operatorsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.OperatorDelegate import OperatorDelegate

		errMsg = "Failed to add elements " + str(operatorsIds) + " for Operators on SalesRegion"

		try:
			# get the SalesRegion
			salesRegion = self.get( salesRegionId ).first()
				
			# split on a comma with no spaces
			idList = operatorsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Operator		
				operator = OperatorDelegate().get(id).first();	
				# add the Operator
				salesRegion.operators.add(operator)
				
			# save it		
			salesRegion.save()
			
			# reload and return the appropriate version
			return self.get( salesRegionId );
		except SalesRegion.DoesNotExist:
			raise ProcessingError(errMsg + " : SalesRegion with id " + str(salesRegionId) + " does not exist.")
		except Operator.DoesNotExist:
			raise ProcessingError(errMsg + " : Operator does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOperators( self, salesRegionId, operatorsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.OperatorDelegate import OperatorDelegate

		errMsg = "Failed to remove elements " + str(operatorsIds) + " for Operators on SalesRegion"

		try:
			# get the SalesRegion
			salesRegion = self.get( salesRegionId ).first()
				
			# split on a comma with no spaces
			idList = operatorsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Operator		
				operator = OperatorDelegate().get(id).first();	
				# add the Operator
				salesRegion.operators.remove(operator)
				
			# save it		
			salesRegion.save()
			
			# reload and return the appropriate version
			return self.get( salesRegionId );
		except SalesRegion.DoesNotExist:
			raise ProcessingError(errMsg + " : SalesRegion with id " + str(salesRegionId) + " does not exist.")
		except Operator.DoesNotExist:
			raise ProcessingError(errMsg + " : Operator does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addSalesCampaigns( self, salesRegionId, salesCampaignsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.SalesCampaignDelegate import SalesCampaignDelegate

		errMsg = "Failed to add elements " + str(salesCampaignsIds) + " for SalesCampaigns on SalesRegion"

		try:
			# get the SalesRegion
			salesRegion = self.get( salesRegionId ).first()
				
			# split on a comma with no spaces
			idList = salesCampaignsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the SalesCampaign		
				salesCampaign = SalesCampaignDelegate().get(id).first();	
				# add the SalesCampaign
				salesRegion.salesCampaigns.add(salesCampaign)
				
			# save it		
			salesRegion.save()
			
			# reload and return the appropriate version
			return self.get( salesRegionId );
		except SalesRegion.DoesNotExist:
			raise ProcessingError(errMsg + " : SalesRegion with id " + str(salesRegionId) + " does not exist.")
		except SalesCampaign.DoesNotExist:
			raise ProcessingError(errMsg + " : SalesCampaign does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeSalesCampaigns( self, salesRegionId, salesCampaignsIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.SalesCampaignDelegate import SalesCampaignDelegate

		errMsg = "Failed to remove elements " + str(salesCampaignsIds) + " for SalesCampaigns on SalesRegion"

		try:
			# get the SalesRegion
			salesRegion = self.get( salesRegionId ).first()
				
			# split on a comma with no spaces
			idList = salesCampaignsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the SalesCampaign		
				salesCampaign = SalesCampaignDelegate().get(id).first();	
				# add the SalesCampaign
				salesRegion.salesCampaigns.remove(salesCampaign)
				
			# save it		
			salesRegion.save()
			
			# reload and return the appropriate version
			return self.get( salesRegionId );
		except SalesRegion.DoesNotExist:
			raise ProcessingError(errMsg + " : SalesRegion with id " + str(salesRegionId) + " does not exist.")
		except SalesCampaign.DoesNotExist:
			raise ProcessingError(errMsg + " : SalesCampaign does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
