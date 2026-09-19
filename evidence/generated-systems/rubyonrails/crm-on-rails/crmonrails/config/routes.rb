Rails.application.routes.draw do
  root "application#health"
  resources :organizations do
    resources :users
    resources :accounts
    resources :teams
    resources :territories
    resources :products
    resources :pricebooks
    resources :campaigns
  end
  resources :users do
    resource :organization
    resources :teams
    resources :activities
    resources :ownedaccounts
    resources :ownedleads
    resources :ownedopportunities
    resources :ownedcases
    resources :quotes
    resources :orders
    resources :contracts
    resources :emailmessages
  end
  resources :teams do
    resource :organization
    resources :users
    resources :accounts
    resources :opportunities
    resources :cases
    resources :campaigns
  end
  resources :territorys do
    resource :organization
    resources :accounts
    resources :users
  end
  resources :accounts do
    resource :organization
    resource :parentaccount
    resources :childaccounts
    resources :contacts
    resources :opportunities
    resources :cases
    resource :owner
    resource :territory
    resources :activities
    resources :campaigns
    resources :quotes
    resources :orders
    resources :contracts
    resources :notes
    resources :emailmessages
  end
  resources :contacts do
    resource :organization
    resource :account
    resource :owner
    resources :activities
    resources :opportunities
    resources :cases
    resources :campaigns
    resources :notes
    resources :emailmessages
  end
  resources :leads do
    resource :organization
    resource :owner
    resources :activities
    resources :campaigns
    resource :convertedaccount
    resource :convertedcontact
    resource :convertedopportunity
    resources :notes
    resources :emailmessages
  end
  resources :opportunitys do
    resource :organization
    resource :account
    resource :owner
    resources :contacts
    resources :lineitems
    resources :stagehistory
    resources :quotes
    resources :orders
    resources :campaigns
    resources :activities
    resources :teams
  end
  resources :opportunitylineitems do
    resource :opportunity
    resource :product
    resource :pricebookentry
  end
  resources :opportunitystagehistorys do
    resource :opportunity
    resource :changedby
  end
  resources :products do
    resource :organization
    resources :pricebookentries
    resources :opportunitylineitems
    resources :quotelineitems
    resources :orderitems
  end
  resources :pricebooks do
    resource :organization
    resources :entries
    resources :quotes
    resources :orders
  end
  resources :pricebookentrys do
    resource :pricebook
    resource :product
  end
  resources :quotes do
    resource :organization
    resource :account
    resource :opportunity
    resource :owner
    resources :lineitems
    resource :pricebook
    resource :order
  end
  resources :quotelineitems do
    resource :quote
    resource :product
    resource :pricebookentry
    resource :opportunitylineitem
  end
  resources :orders do
    resource :organization
    resource :account
    resource :opportunity
    resource :quote
    resource :owner
    resources :items
    resource :contract
    resource :pricebook
  end
  resources :orderitems do
    resource :order
    resource :product
    resource :pricebookentry
  end
  resources :contracts do
    resource :organization
    resource :account
    resource :owner
    resources :orders
    resources :cases
  end
  resources :case_s do
    resource :organization
    resource :account
    resource :contact
    resource :owner
    resource :team
    resources :activities
    resources :casecomments
    resources :emails
    resources :relatedopportunities
  end
  resources :activitys do
    resource :organization
    resource :owner
    resource :account
    resource :contact
    resource :lead
    resource :opportunity
    resource :case
    resource :campaign
  end
  resources :campaigns do
    resource :organization
    resource :parentcampaign
    resources :childcampaigns
    resources :members
    resources :opportunities
    resources :accounts
    resources :leads
    resources :contacts
    resources :teams
    resources :activities
  end
  resources :campaignmembers do
    resource :campaign
    resource :lead
    resource :contact
  end
  resources :notes do
    resource :organization
    resource :owner
    resource :account
    resource :contact
    resource :opportunity
    resource :case
    resource :lead
  end
  resources :emailmessages do
    resource :organization
    resource :owner
    resource :account
    resource :contact
    resource :lead
    resource :case
    resource :opportunity
    resource :campaign
  end
end
