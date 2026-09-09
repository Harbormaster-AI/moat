from django.contrib import admin

# Register your models here.
from .models.Organization import Organization
from .models.User import User
from .models.Team import Team
from .models.Territory import Territory
from .models.Account import Account
from .models.Contact import Contact
from .models.Lead import Lead
from .models.Opportunity import Opportunity
from .models.OpportunityLineItem import OpportunityLineItem
from .models.OpportunityStageHistory import OpportunityStageHistory
from .models.Product import Product
from .models.PriceBook import PriceBook
from .models.PriceBookEntry import PriceBookEntry
from .models.Quote import Quote
from .models.QuoteLineItem import QuoteLineItem
from .models.Order import Order
from .models.OrderItem import OrderItem
from .models.Contract import Contract
from .models.Case_ import Case_
from .models.Activity import Activity
from .models.Campaign import Campaign
from .models.CampaignMember import CampaignMember
from .models.Note import Note
from .models.EmailMessage import EmailMessage

# Need to add this for each model that requires managing

admin.site.register(Organization)
admin.site.register(User)
admin.site.register(Team)
admin.site.register(Territory)
admin.site.register(Account)
admin.site.register(Contact)
admin.site.register(Lead)
admin.site.register(Opportunity)
admin.site.register(OpportunityLineItem)
admin.site.register(OpportunityStageHistory)
admin.site.register(Product)
admin.site.register(PriceBook)
admin.site.register(PriceBookEntry)
admin.site.register(Quote)
admin.site.register(QuoteLineItem)
admin.site.register(Order)
admin.site.register(OrderItem)
admin.site.register(Contract)
admin.site.register(Case_)
admin.site.register(Activity)
admin.site.register(Campaign)
admin.site.register(CampaignMember)
admin.site.register(Note)
admin.site.register(EmailMessage)
