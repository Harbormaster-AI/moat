from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from aerospaceOnDjango.models.SalesCampaign import SalesCampaign
from aerospaceOnDjango.models.SalesRegion import SalesRegion
from aerospaceOnDjango.models.Operator import Operator
from aerospaceOnDjango.models.Quote import Quote
from aerospaceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model SalesCampaign
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class SalesCampaignDelegate Declaration
#======================================================================
class SalesCampaignDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, salesCampaignId ):
		try:	
			salesCampaign = SalesCampaign.objects.filter(id=salesCampaignId)
			return salesCampaign.first();
		except SalesCampaign.DoesNotExist:
			raise ProcessingError("SalesCampaign with id " + str(salesCampaignId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, salesCampaign):
		for model in serializers.deserialize("json", salesCampaign):
			model.save()
			return model;

	def create(self, salesCampaign):
		salesCampaign.save()
		return salesCampaign;

	def saveFromJson(self, salesCampaign):
		for model in serializers.deserialize("json", salesCampaign):
			model.save()
			return salesCampaign;
	
	def save(self, salesCampaign):
		salesCampaign.save()
		return salesCampaign;
	
	def delete(self, salesCampaignId ):
		errMsg = "Failed to delete SalesCampaign from db using id " + str(salesCampaignId)
		
		try:
			salesCampaign = SalesCampaign.objects.get(id=salesCampaignId)
			salesCampaign.delete()
			return True
		except SalesCampaign.DoesNotExist:
			raise ProcessingError("SalesCampaign with id " + str(salesCampaignId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = SalesCampaign.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all SalesCampaign from db")
		except Exception:
			return None;
		
	def assignRegion( self, salesCampaignId, regionId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.SalesRegionDelegate import SalesRegionDelegate

		errMsg = "Failed to assign element " + str(regionId) + " for Region on SalesCampaign"

		try:
			# get the SalesCampaign from db
			salesCampaign = self.get( salesCampaignId ).first()	
			
			# get the SalesRegion from db
			salesRegion = SalesRegionDelegate().get(regionId).first();
			
			# assign the Region		
			salesCampaign.region = salesRegion
			
			#save it
			salesCampaign.save()

			# reload and return the appropriate version					
			return self.get( salesCampaignId );
		except SalesCampaign.DoesNotExist:
			raise ProcessingError(errMsg + " : SalesCampaign with id " + str(salesCampaignId) + " does not exist.")
		except SalesRegion.DoesNotExist:
			raise ProcessingError(errMsg + " : SalesRegion with id " + str(regionId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignRegion( self, salesCampaignId ):
		errMsg = "Failed to unassign element " + str(regionId) + " for Region on SalesCampaign"

		try:
			# get the SalesCampaign from db
			salesCampaign = self.get( salesCampaignId ).first()	
			
			# assign to None for unassignment
			salesCampaign.salesRegion = None			

			#save it
			salesCampaign.save()

			# reload and return the appropriate version					
			return self.get( salesCampaignId );
		except SalesCampaign.DoesNotExist:
			raise ProcessingError(errMsg + " : SalesCampaign with id " + str(salesCampaignId) + " does not exist.")
		except Exception:
			return None;
		
	def assignOperator( self, salesCampaignId, operatorId ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.OperatorDelegate import OperatorDelegate

		errMsg = "Failed to assign element " + str(operatorId) + " for Operator on SalesCampaign"

		try:
			# get the SalesCampaign from db
			salesCampaign = self.get( salesCampaignId ).first()	
			
			# get the Operator from db
			operator = OperatorDelegate().get(operatorId).first();
			
			# assign the Operator		
			salesCampaign.operator = operator
			
			#save it
			salesCampaign.save()

			# reload and return the appropriate version					
			return self.get( salesCampaignId );
		except SalesCampaign.DoesNotExist:
			raise ProcessingError(errMsg + " : SalesCampaign with id " + str(salesCampaignId) + " does not exist.")
		except Operator.DoesNotExist:
			raise ProcessingError(errMsg + " : Operator with id " + str(operatorId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOperator( self, salesCampaignId ):
		errMsg = "Failed to unassign element " + str(operatorId) + " for Operator on SalesCampaign"

		try:
			# get the SalesCampaign from db
			salesCampaign = self.get( salesCampaignId ).first()	
			
			# assign to None for unassignment
			salesCampaign.operator = None			

			#save it
			salesCampaign.save()

			# reload and return the appropriate version					
			return self.get( salesCampaignId );
		except SalesCampaign.DoesNotExist:
			raise ProcessingError(errMsg + " : SalesCampaign with id " + str(salesCampaignId) + " does not exist.")
		except Exception:
			return None;
		
	def addQuotes( self, salesCampaignId, quotesIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.QuoteDelegate import QuoteDelegate

		errMsg = "Failed to add elements " + str(quotesIds) + " for Quotes on SalesCampaign"

		try:
			# get the SalesCampaign
			salesCampaign = self.get( salesCampaignId ).first()
				
			# split on a comma with no spaces
			idList = quotesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Quote		
				quote = QuoteDelegate().get(id).first();	
				# add the Quote
				salesCampaign.quotes.add(quote)
				
			# save it		
			salesCampaign.save()
			
			# reload and return the appropriate version
			return self.get( salesCampaignId );
		except SalesCampaign.DoesNotExist:
			raise ProcessingError(errMsg + " : SalesCampaign with id " + str(salesCampaignId) + " does not exist.")
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeQuotes( self, salesCampaignId, quotesIds ):
		# lazy importing avoids circular dependencies
		from aerospaceOnDjango.delegates.QuoteDelegate import QuoteDelegate

		errMsg = "Failed to remove elements " + str(quotesIds) + " for Quotes on SalesCampaign"

		try:
			# get the SalesCampaign
			salesCampaign = self.get( salesCampaignId ).first()
				
			# split on a comma with no spaces
			idList = quotesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Quote		
				quote = QuoteDelegate().get(id).first();	
				# add the Quote
				salesCampaign.quotes.remove(quote)
				
			# save it		
			salesCampaign.save()
			
			# reload and return the appropriate version
			return self.get( salesCampaignId );
		except SalesCampaign.DoesNotExist:
			raise ProcessingError(errMsg + " : SalesCampaign with id " + str(salesCampaignId) + " does not exist.")
		except Quote.DoesNotExist:
			raise ProcessingError(errMsg + " : Quote does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
