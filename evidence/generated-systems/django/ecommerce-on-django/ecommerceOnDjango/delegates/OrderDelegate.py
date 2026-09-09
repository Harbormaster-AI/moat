from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from ecommerceOnDjango.models.Order import Order
from ecommerceOnDjango.models.Customer import Customer
from ecommerceOnDjango.models.Channel import Channel
from ecommerceOnDjango.models.OrderLine import OrderLine
from ecommerceOnDjango.models.Payment import Payment
from ecommerceOnDjango.models.Shipment import Shipment
from ecommerceOnDjango.models.Refund import Refund
from ecommerceOnDjango.models.Promotion import Promotion
from ecommerceOnDjango.models.Seller import Seller
from ecommerceOnDjango.models.GiftCardRedemption import GiftCardRedemption
from ecommerceOnDjango.models.CouponRedemption import CouponRedemption
from ecommerceOnDjango.models.ReturnRequest import ReturnRequest
from ecommerceOnDjango.models.Invoice import Invoice
from ecommerceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Order
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OrderDelegate Declaration
#======================================================================
class OrderDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, orderId ):
		try:	
			order = Order.objects.filter(id=orderId)
			return order.first();
		except Order.DoesNotExist:
			raise ProcessingError("Order with id " + str(orderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, order):
		for model in serializers.deserialize("json", order):
			model.save()
			return model;

	def create(self, order):
		order.save()
		return order;

	def saveFromJson(self, order):
		for model in serializers.deserialize("json", order):
			model.save()
			return order;
	
	def save(self, order):
		order.save()
		return order;
	
	def delete(self, orderId ):
		errMsg = "Failed to delete Order from db using id " + str(orderId)
		
		try:
			order = Order.objects.get(id=orderId)
			order.delete()
			return True
		except Order.DoesNotExist:
			raise ProcessingError("Order with id " + str(orderId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Order.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Order from db")
		except Exception:
			return None;
		
	def assignCustomer( self, orderId, customerId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CustomerDelegate import CustomerDelegate

		errMsg = "Failed to assign element " + str(customerId) + " for Customer on Order"

		try:
			# get the Order from db
			order = self.get( orderId ).first()	
			
			# get the Customer from db
			customer = CustomerDelegate().get(customerId).first();
			
			# assign the Customer		
			order.customer = customer
			
			#save it
			order.save()

			# reload and return the appropriate version					
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Customer.DoesNotExist:
			raise ProcessingError(errMsg + " : Customer with id " + str(customerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCustomer( self, orderId ):
		errMsg = "Failed to unassign element " + str(customerId) + " for Customer on Order"

		try:
			# get the Order from db
			order = self.get( orderId ).first()	
			
			# assign to None for unassignment
			order.customer = None			

			#save it
			order.save()

			# reload and return the appropriate version					
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignChannel( self, orderId, channelId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ChannelDelegate import ChannelDelegate

		errMsg = "Failed to assign element " + str(channelId) + " for Channel on Order"

		try:
			# get the Order from db
			order = self.get( orderId ).first()	
			
			# get the Channel from db
			channel = ChannelDelegate().get(channelId).first();
			
			# assign the Channel		
			order.channel = channel
			
			#save it
			order.save()

			# reload and return the appropriate version					
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Channel.DoesNotExist:
			raise ProcessingError(errMsg + " : Channel with id " + str(channelId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignChannel( self, orderId ):
		errMsg = "Failed to unassign element " + str(channelId) + " for Channel on Order"

		try:
			# get the Order from db
			order = self.get( orderId ).first()	
			
			# assign to None for unassignment
			order.channel = None			

			#save it
			order.save()

			# reload and return the appropriate version					
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignSeller( self, orderId, sellerId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.SellerDelegate import SellerDelegate

		errMsg = "Failed to assign element " + str(sellerId) + " for Seller on Order"

		try:
			# get the Order from db
			order = self.get( orderId ).first()	
			
			# get the Seller from db
			seller = SellerDelegate().get(sellerId).first();
			
			# assign the Seller		
			order.seller = seller
			
			#save it
			order.save()

			# reload and return the appropriate version					
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Seller.DoesNotExist:
			raise ProcessingError(errMsg + " : Seller with id " + str(sellerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignSeller( self, orderId ):
		errMsg = "Failed to unassign element " + str(sellerId) + " for Seller on Order"

		try:
			# get the Order from db
			order = self.get( orderId ).first()	
			
			# assign to None for unassignment
			order.seller = None			

			#save it
			order.save()

			# reload and return the appropriate version					
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Exception:
			return None;
		
	def assignInvoice( self, orderId, invoiceId ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.InvoiceDelegate import InvoiceDelegate

		errMsg = "Failed to assign element " + str(invoiceId) + " for Invoice on Order"

		try:
			# get the Order from db
			order = self.get( orderId ).first()	
			
			# get the Invoice from db
			invoice = InvoiceDelegate().get(invoiceId).first();
			
			# assign the Invoice		
			order.invoice = invoice
			
			#save it
			order.save()

			# reload and return the appropriate version					
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Invoice.DoesNotExist:
			raise ProcessingError(errMsg + " : Invoice with id " + str(invoiceId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignInvoice( self, orderId ):
		errMsg = "Failed to unassign element " + str(invoiceId) + " for Invoice on Order"

		try:
			# get the Order from db
			order = self.get( orderId ).first()	
			
			# assign to None for unassignment
			order.invoice = None			

			#save it
			order.save()

			# reload and return the appropriate version					
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Exception:
			return None;
		
	def addOrderLines( self, orderId, orderLinesIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.OrderLineDelegate import OrderLineDelegate

		errMsg = "Failed to add elements " + str(orderLinesIds) + " for OrderLines on Order"

		try:
			# get the Order
			order = self.get( orderId ).first()
				
			# split on a comma with no spaces
			idList = orderLinesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the OrderLine		
				orderLine = OrderLineDelegate().get(id).first();	
				# add the OrderLine
				order.orderLines.add(orderLine)
				
			# save it		
			order.save()
			
			# reload and return the appropriate version
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except OrderLine.DoesNotExist:
			raise ProcessingError(errMsg + " : OrderLine does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOrderLines( self, orderId, orderLinesIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.OrderLineDelegate import OrderLineDelegate

		errMsg = "Failed to remove elements " + str(orderLinesIds) + " for OrderLines on Order"

		try:
			# get the Order
			order = self.get( orderId ).first()
				
			# split on a comma with no spaces
			idList = orderLinesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the OrderLine		
				orderLine = OrderLineDelegate().get(id).first();	
				# add the OrderLine
				order.orderLines.remove(orderLine)
				
			# save it		
			order.save()
			
			# reload and return the appropriate version
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except OrderLine.DoesNotExist:
			raise ProcessingError(errMsg + " : OrderLine does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPayments( self, orderId, paymentsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.PaymentDelegate import PaymentDelegate

		errMsg = "Failed to add elements " + str(paymentsIds) + " for Payments on Order"

		try:
			# get the Order
			order = self.get( orderId ).first()
				
			# split on a comma with no spaces
			idList = paymentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Payment		
				payment = PaymentDelegate().get(id).first();	
				# add the Payment
				order.payments.add(payment)
				
			# save it		
			order.save()
			
			# reload and return the appropriate version
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Payment.DoesNotExist:
			raise ProcessingError(errMsg + " : Payment does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePayments( self, orderId, paymentsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.PaymentDelegate import PaymentDelegate

		errMsg = "Failed to remove elements " + str(paymentsIds) + " for Payments on Order"

		try:
			# get the Order
			order = self.get( orderId ).first()
				
			# split on a comma with no spaces
			idList = paymentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Payment		
				payment = PaymentDelegate().get(id).first();	
				# add the Payment
				order.payments.remove(payment)
				
			# save it		
			order.save()
			
			# reload and return the appropriate version
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Payment.DoesNotExist:
			raise ProcessingError(errMsg + " : Payment does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addShipments( self, orderId, shipmentsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ShipmentDelegate import ShipmentDelegate

		errMsg = "Failed to add elements " + str(shipmentsIds) + " for Shipments on Order"

		try:
			# get the Order
			order = self.get( orderId ).first()
				
			# split on a comma with no spaces
			idList = shipmentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Shipment		
				shipment = ShipmentDelegate().get(id).first();	
				# add the Shipment
				order.shipments.add(shipment)
				
			# save it		
			order.save()
			
			# reload and return the appropriate version
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Shipment.DoesNotExist:
			raise ProcessingError(errMsg + " : Shipment does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeShipments( self, orderId, shipmentsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ShipmentDelegate import ShipmentDelegate

		errMsg = "Failed to remove elements " + str(shipmentsIds) + " for Shipments on Order"

		try:
			# get the Order
			order = self.get( orderId ).first()
				
			# split on a comma with no spaces
			idList = shipmentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Shipment		
				shipment = ShipmentDelegate().get(id).first();	
				# add the Shipment
				order.shipments.remove(shipment)
				
			# save it		
			order.save()
			
			# reload and return the appropriate version
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Shipment.DoesNotExist:
			raise ProcessingError(errMsg + " : Shipment does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addRefunds( self, orderId, refundsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.RefundDelegate import RefundDelegate

		errMsg = "Failed to add elements " + str(refundsIds) + " for Refunds on Order"

		try:
			# get the Order
			order = self.get( orderId ).first()
				
			# split on a comma with no spaces
			idList = refundsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Refund		
				refund = RefundDelegate().get(id).first();	
				# add the Refund
				order.refunds.add(refund)
				
			# save it		
			order.save()
			
			# reload and return the appropriate version
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Refund.DoesNotExist:
			raise ProcessingError(errMsg + " : Refund does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeRefunds( self, orderId, refundsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.RefundDelegate import RefundDelegate

		errMsg = "Failed to remove elements " + str(refundsIds) + " for Refunds on Order"

		try:
			# get the Order
			order = self.get( orderId ).first()
				
			# split on a comma with no spaces
			idList = refundsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Refund		
				refund = RefundDelegate().get(id).first();	
				# add the Refund
				order.refunds.remove(refund)
				
			# save it		
			order.save()
			
			# reload and return the appropriate version
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Refund.DoesNotExist:
			raise ProcessingError(errMsg + " : Refund does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addAppliedPromotions( self, orderId, appliedPromotionsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.PromotionDelegate import PromotionDelegate

		errMsg = "Failed to add elements " + str(appliedPromotionsIds) + " for AppliedPromotions on Order"

		try:
			# get the Order
			order = self.get( orderId ).first()
				
			# split on a comma with no spaces
			idList = appliedPromotionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Promotion		
				promotion = PromotionDelegate().get(id).first();	
				# add the Promotion
				order.appliedPromotions.add(promotion)
				
			# save it		
			order.save()
			
			# reload and return the appropriate version
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Promotion.DoesNotExist:
			raise ProcessingError(errMsg + " : Promotion does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAppliedPromotions( self, orderId, appliedPromotionsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.PromotionDelegate import PromotionDelegate

		errMsg = "Failed to remove elements " + str(appliedPromotionsIds) + " for AppliedPromotions on Order"

		try:
			# get the Order
			order = self.get( orderId ).first()
				
			# split on a comma with no spaces
			idList = appliedPromotionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Promotion		
				promotion = PromotionDelegate().get(id).first();	
				# add the Promotion
				order.appliedPromotions.remove(promotion)
				
			# save it		
			order.save()
			
			# reload and return the appropriate version
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except Promotion.DoesNotExist:
			raise ProcessingError(errMsg + " : Promotion does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addGiftCardRedemptions( self, orderId, giftCardRedemptionsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.GiftCardRedemptionDelegate import GiftCardRedemptionDelegate

		errMsg = "Failed to add elements " + str(giftCardRedemptionsIds) + " for GiftCardRedemptions on Order"

		try:
			# get the Order
			order = self.get( orderId ).first()
				
			# split on a comma with no spaces
			idList = giftCardRedemptionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the GiftCardRedemption		
				giftCardRedemption = GiftCardRedemptionDelegate().get(id).first();	
				# add the GiftCardRedemption
				order.giftCardRedemptions.add(giftCardRedemption)
				
			# save it		
			order.save()
			
			# reload and return the appropriate version
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except GiftCardRedemption.DoesNotExist:
			raise ProcessingError(errMsg + " : GiftCardRedemption does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeGiftCardRedemptions( self, orderId, giftCardRedemptionsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.GiftCardRedemptionDelegate import GiftCardRedemptionDelegate

		errMsg = "Failed to remove elements " + str(giftCardRedemptionsIds) + " for GiftCardRedemptions on Order"

		try:
			# get the Order
			order = self.get( orderId ).first()
				
			# split on a comma with no spaces
			idList = giftCardRedemptionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the GiftCardRedemption		
				giftCardRedemption = GiftCardRedemptionDelegate().get(id).first();	
				# add the GiftCardRedemption
				order.giftCardRedemptions.remove(giftCardRedemption)
				
			# save it		
			order.save()
			
			# reload and return the appropriate version
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except GiftCardRedemption.DoesNotExist:
			raise ProcessingError(errMsg + " : GiftCardRedemption does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addCouponRedemptions( self, orderId, couponRedemptionsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CouponRedemptionDelegate import CouponRedemptionDelegate

		errMsg = "Failed to add elements " + str(couponRedemptionsIds) + " for CouponRedemptions on Order"

		try:
			# get the Order
			order = self.get( orderId ).first()
				
			# split on a comma with no spaces
			idList = couponRedemptionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the CouponRedemption		
				couponRedemption = CouponRedemptionDelegate().get(id).first();	
				# add the CouponRedemption
				order.couponRedemptions.add(couponRedemption)
				
			# save it		
			order.save()
			
			# reload and return the appropriate version
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except CouponRedemption.DoesNotExist:
			raise ProcessingError(errMsg + " : CouponRedemption does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCouponRedemptions( self, orderId, couponRedemptionsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.CouponRedemptionDelegate import CouponRedemptionDelegate

		errMsg = "Failed to remove elements " + str(couponRedemptionsIds) + " for CouponRedemptions on Order"

		try:
			# get the Order
			order = self.get( orderId ).first()
				
			# split on a comma with no spaces
			idList = couponRedemptionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the CouponRedemption		
				couponRedemption = CouponRedemptionDelegate().get(id).first();	
				# add the CouponRedemption
				order.couponRedemptions.remove(couponRedemption)
				
			# save it		
			order.save()
			
			# reload and return the appropriate version
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except CouponRedemption.DoesNotExist:
			raise ProcessingError(errMsg + " : CouponRedemption does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addReturnRequests( self, orderId, returnRequestsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ReturnRequestDelegate import ReturnRequestDelegate

		errMsg = "Failed to add elements " + str(returnRequestsIds) + " for ReturnRequests on Order"

		try:
			# get the Order
			order = self.get( orderId ).first()
				
			# split on a comma with no spaces
			idList = returnRequestsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ReturnRequest		
				returnRequest = ReturnRequestDelegate().get(id).first();	
				# add the ReturnRequest
				order.returnRequests.add(returnRequest)
				
			# save it		
			order.save()
			
			# reload and return the appropriate version
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except ReturnRequest.DoesNotExist:
			raise ProcessingError(errMsg + " : ReturnRequest does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeReturnRequests( self, orderId, returnRequestsIds ):
		# lazy importing avoids circular dependencies
		from ecommerceOnDjango.delegates.ReturnRequestDelegate import ReturnRequestDelegate

		errMsg = "Failed to remove elements " + str(returnRequestsIds) + " for ReturnRequests on Order"

		try:
			# get the Order
			order = self.get( orderId ).first()
				
			# split on a comma with no spaces
			idList = returnRequestsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ReturnRequest		
				returnRequest = ReturnRequestDelegate().get(id).first();	
				# add the ReturnRequest
				order.returnRequests.remove(returnRequest)
				
			# save it		
			order.save()
			
			# reload and return the appropriate version
			return self.get( orderId );
		except Order.DoesNotExist:
			raise ProcessingError(errMsg + " : Order with id " + str(orderId) + " does not exist.")
		except ReturnRequest.DoesNotExist:
			raise ProcessingError(errMsg + " : ReturnRequest does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
