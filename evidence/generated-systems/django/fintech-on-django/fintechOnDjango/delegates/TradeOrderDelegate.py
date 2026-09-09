from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.TradeOrder import TradeOrder
from fintechOnDjango.models.InvestmentPortfolio import InvestmentPortfolio
from fintechOnDjango.models.Security import Security
from fintechOnDjango.models.Trade import Trade
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model TradeOrder
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TradeOrderDelegate Declaration
#======================================================================
class TradeOrderDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, tradeOrderId ):
		try:	
			tradeOrder = TradeOrder.objects.filter(id=tradeOrderId)
			return tradeOrder.first();
		except TradeOrder.DoesNotExist:
			raise ProcessingError("TradeOrder with id " + str(tradeOrderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, tradeOrder):
		for model in serializers.deserialize("json", tradeOrder):
			model.save()
			return model;

	def create(self, tradeOrder):
		tradeOrder.save()
		return tradeOrder;

	def saveFromJson(self, tradeOrder):
		for model in serializers.deserialize("json", tradeOrder):
			model.save()
			return tradeOrder;
	
	def save(self, tradeOrder):
		tradeOrder.save()
		return tradeOrder;
	
	def delete(self, tradeOrderId ):
		errMsg = "Failed to delete TradeOrder from db using id " + str(tradeOrderId)
		
		try:
			tradeOrder = TradeOrder.objects.get(id=tradeOrderId)
			tradeOrder.delete()
			return True
		except TradeOrder.DoesNotExist:
			raise ProcessingError("TradeOrder with id " + str(tradeOrderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = TradeOrder.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all TradeOrder from db")
		except Exception:
			return None;
		
	def assignPortfolio( self, tradeOrderId, portfolioId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.InvestmentPortfolioDelegate import InvestmentPortfolioDelegate

		errMsg = "Failed to assign element " + str(portfolioId) + " for Portfolio on TradeOrder"

		try:
			# get the TradeOrder from db
			tradeOrder = self.get( tradeOrderId ).first()	
			
			# get the InvestmentPortfolio from db
			investmentPortfolio = InvestmentPortfolioDelegate().get(portfolioId).first();
			
			# assign the Portfolio		
			tradeOrder.portfolio = investmentPortfolio
			
			#save it
			tradeOrder.save()

			# reload and return the appropriate version					
			return self.get( tradeOrderId );
		except TradeOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : TradeOrder with id " + str(tradeOrderId) + " does not exist.")
		except InvestmentPortfolio.DoesNotExist:
			raise ProcessingError(errMsg + " : InvestmentPortfolio with id " + str(portfolioId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPortfolio( self, tradeOrderId ):
		errMsg = "Failed to unassign element " + str(portfolioId) + " for Portfolio on TradeOrder"

		try:
			# get the TradeOrder from db
			tradeOrder = self.get( tradeOrderId ).first()	
			
			# assign to None for unassignment
			tradeOrder.investmentPortfolio = None			

			#save it
			tradeOrder.save()

			# reload and return the appropriate version					
			return self.get( tradeOrderId );
		except TradeOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : TradeOrder with id " + str(tradeOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignSecurity( self, tradeOrderId, securityId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.SecurityDelegate import SecurityDelegate

		errMsg = "Failed to assign element " + str(securityId) + " for Security on TradeOrder"

		try:
			# get the TradeOrder from db
			tradeOrder = self.get( tradeOrderId ).first()	
			
			# get the Security from db
			security = SecurityDelegate().get(securityId).first();
			
			# assign the Security		
			tradeOrder.security = security
			
			#save it
			tradeOrder.save()

			# reload and return the appropriate version					
			return self.get( tradeOrderId );
		except TradeOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : TradeOrder with id " + str(tradeOrderId) + " does not exist.")
		except Security.DoesNotExist:
			raise ProcessingError(errMsg + " : Security with id " + str(securityId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSecurity( self, tradeOrderId ):
		errMsg = "Failed to unassign element " + str(securityId) + " for Security on TradeOrder"

		try:
			# get the TradeOrder from db
			tradeOrder = self.get( tradeOrderId ).first()	
			
			# assign to None for unassignment
			tradeOrder.security = None			

			#save it
			tradeOrder.save()

			# reload and return the appropriate version					
			return self.get( tradeOrderId );
		except TradeOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : TradeOrder with id " + str(tradeOrderId) + " does not exist.")
		except Exception:
			return None;
		
	def addTrades( self, tradeOrderId, tradesIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.TradeDelegate import TradeDelegate

		errMsg = "Failed to add elements " + str(tradesIds) + " for Trades on TradeOrder"

		try:
			# get the TradeOrder
			tradeOrder = self.get( tradeOrderId ).first()
				
			# split on a comma with no spaces
			idList = tradesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Trade		
				trade = TradeDelegate().get(id).first();	
				# add the Trade
				tradeOrder.trades.add(trade)
				
			# save it		
			tradeOrder.save()
			
			# reload and return the appropriate version
			return self.get( tradeOrderId );
		except TradeOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : TradeOrder with id " + str(tradeOrderId) + " does not exist.")
		except Trade.DoesNotExist:
			raise ProcessingError(errMsg + " : Trade does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeTrades( self, tradeOrderId, tradesIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.TradeDelegate import TradeDelegate

		errMsg = "Failed to remove elements " + str(tradesIds) + " for Trades on TradeOrder"

		try:
			# get the TradeOrder
			tradeOrder = self.get( tradeOrderId ).first()
				
			# split on a comma with no spaces
			idList = tradesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Trade		
				trade = TradeDelegate().get(id).first();	
				# add the Trade
				tradeOrder.trades.remove(trade)
				
			# save it		
			tradeOrder.save()
			
			# reload and return the appropriate version
			return self.get( tradeOrderId );
		except TradeOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : TradeOrder with id " + str(tradeOrderId) + " does not exist.")
		except Trade.DoesNotExist:
			raise ProcessingError(errMsg + " : Trade does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
