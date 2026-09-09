from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.InvestmentAccount import InvestmentAccount
from fintechOnDjango.models.InvestmentPortfolio import InvestmentPortfolio
from fintechOnDjango.models.Trade import Trade
from fintechOnDjango.models.TradeOrder import TradeOrder
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model InvestmentAccount
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class InvestmentAccountDelegate Declaration
#======================================================================
class InvestmentAccountDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, investmentAccountId ):
		try:	
			investmentAccount = InvestmentAccount.objects.filter(id=investmentAccountId)
			return investmentAccount.first();
		except InvestmentAccount.DoesNotExist:
			raise ProcessingError("InvestmentAccount with id " + str(investmentAccountId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, investmentAccount):
		for model in serializers.deserialize("json", investmentAccount):
			model.save()
			return model;

	def create(self, investmentAccount):
		investmentAccount.save()
		return investmentAccount;

	def saveFromJson(self, investmentAccount):
		for model in serializers.deserialize("json", investmentAccount):
			model.save()
			return investmentAccount;
	
	def save(self, investmentAccount):
		investmentAccount.save()
		return investmentAccount;
	
	def delete(self, investmentAccountId ):
		errMsg = "Failed to delete InvestmentAccount from db using id " + str(investmentAccountId)
		
		try:
			investmentAccount = InvestmentAccount.objects.get(id=investmentAccountId)
			investmentAccount.delete()
			return True
		except InvestmentAccount.DoesNotExist:
			raise ProcessingError("InvestmentAccount with id " + str(investmentAccountId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = InvestmentAccount.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all InvestmentAccount from db")
		except Exception:
			return None;
		
	def assignPortfolio( self, investmentAccountId, portfolioId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.InvestmentPortfolioDelegate import InvestmentPortfolioDelegate

		errMsg = "Failed to assign element " + str(portfolioId) + " for Portfolio on InvestmentAccount"

		try:
			# get the InvestmentAccount from db
			investmentAccount = self.get( investmentAccountId ).first()	
			
			# get the InvestmentPortfolio from db
			investmentPortfolio = InvestmentPortfolioDelegate().get(portfolioId).first();
			
			# assign the Portfolio		
			investmentAccount.portfolio = investmentPortfolio
			
			#save it
			investmentAccount.save()

			# reload and return the appropriate version					
			return self.get( investmentAccountId );
		except InvestmentAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : InvestmentAccount with id " + str(investmentAccountId) + " does not exist.")
		except InvestmentPortfolio.DoesNotExist:
			raise ProcessingError(errMsg + " : InvestmentPortfolio with id " + str(portfolioId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPortfolio( self, investmentAccountId ):
		errMsg = "Failed to unassign element " + str(portfolioId) + " for Portfolio on InvestmentAccount"

		try:
			# get the InvestmentAccount from db
			investmentAccount = self.get( investmentAccountId ).first()	
			
			# assign to None for unassignment
			investmentAccount.investmentPortfolio = None			

			#save it
			investmentAccount.save()

			# reload and return the appropriate version					
			return self.get( investmentAccountId );
		except InvestmentAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : InvestmentAccount with id " + str(investmentAccountId) + " does not exist.")
		except Exception:
			return None;
		
	def addTrades( self, investmentAccountId, tradesIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.TradeDelegate import TradeDelegate

		errMsg = "Failed to add elements " + str(tradesIds) + " for Trades on InvestmentAccount"

		try:
			# get the InvestmentAccount
			investmentAccount = self.get( investmentAccountId ).first()
				
			# split on a comma with no spaces
			idList = tradesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Trade		
				trade = TradeDelegate().get(id).first();	
				# add the Trade
				investmentAccount.trades.add(trade)
				
			# save it		
			investmentAccount.save()
			
			# reload and return the appropriate version
			return self.get( investmentAccountId );
		except InvestmentAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : InvestmentAccount with id " + str(investmentAccountId) + " does not exist.")
		except Trade.DoesNotExist:
			raise ProcessingError(errMsg + " : Trade does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeTrades( self, investmentAccountId, tradesIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.TradeDelegate import TradeDelegate

		errMsg = "Failed to remove elements " + str(tradesIds) + " for Trades on InvestmentAccount"

		try:
			# get the InvestmentAccount
			investmentAccount = self.get( investmentAccountId ).first()
				
			# split on a comma with no spaces
			idList = tradesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Trade		
				trade = TradeDelegate().get(id).first();	
				# add the Trade
				investmentAccount.trades.remove(trade)
				
			# save it		
			investmentAccount.save()
			
			# reload and return the appropriate version
			return self.get( investmentAccountId );
		except InvestmentAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : InvestmentAccount with id " + str(investmentAccountId) + " does not exist.")
		except Trade.DoesNotExist:
			raise ProcessingError(errMsg + " : Trade does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addOrders( self, investmentAccountId, ordersIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.TradeOrderDelegate import TradeOrderDelegate

		errMsg = "Failed to add elements " + str(ordersIds) + " for Orders on InvestmentAccount"

		try:
			# get the InvestmentAccount
			investmentAccount = self.get( investmentAccountId ).first()
				
			# split on a comma with no spaces
			idList = ordersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the TradeOrder		
				tradeOrder = TradeOrderDelegate().get(id).first();	
				# add the TradeOrder
				investmentAccount.orders.add(tradeOrder)
				
			# save it		
			investmentAccount.save()
			
			# reload and return the appropriate version
			return self.get( investmentAccountId );
		except InvestmentAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : InvestmentAccount with id " + str(investmentAccountId) + " does not exist.")
		except TradeOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : TradeOrder does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOrders( self, investmentAccountId, ordersIds ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.TradeOrderDelegate import TradeOrderDelegate

		errMsg = "Failed to remove elements " + str(ordersIds) + " for Orders on InvestmentAccount"

		try:
			# get the InvestmentAccount
			investmentAccount = self.get( investmentAccountId ).first()
				
			# split on a comma with no spaces
			idList = ordersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the TradeOrder		
				tradeOrder = TradeOrderDelegate().get(id).first();	
				# add the TradeOrder
				investmentAccount.orders.remove(tradeOrder)
				
			# save it		
			investmentAccount.save()
			
			# reload and return the appropriate version
			return self.get( investmentAccountId );
		except InvestmentAccount.DoesNotExist:
			raise ProcessingError(errMsg + " : InvestmentAccount with id " + str(investmentAccountId) + " does not exist.")
		except TradeOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : TradeOrder does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
