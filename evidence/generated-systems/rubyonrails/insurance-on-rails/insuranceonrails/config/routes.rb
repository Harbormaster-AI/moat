Rails.application.routes.draw do
  root "application#health"
  resources :insurers do
    resources :products
    resources :distributionpartners
    resources :policies
    resources :claims
    resources :reinsuranceagreements
  end
  resources :insuranceproducts do
    resource :insurer
    resources :coveragedefinitions
  end
  resources :coveragedefinitions do
    resource :product
  end
  resources :distributors do
    resources :insurers
    resources :agents
    resources :policies
  end
  resources :agents do
    resource :distributor
    resources :policies
    resources :customers
  end
  resources :customers do
    resources :applications
    resources :policies
    resources :claims
    resources :agents
    resources :beneficiaries
  end
  resources :applications do
    resource :customer
    resource :product
    resource :distributor
    resources :quotes
    resource :selectedquote
  end
  resources :quotes do
    resource :application
    resources :underwritingdecisions
    resource :policy
  end
  resources :underwritingdecisions do
    resource :quote
    resource :underwriter
  end
  resources :underwriters do
    resources :decisions
    resource :insurer
  end
  resources :policys do
    resource :insurer
    resource :customer
    resource :product
    resource :agent
    resources :coverages
    resources :insuredobjects
    resources :endorsements
    resource :billingaccount
    resources :beneficiaries
    resources :claims
    resources :reinsuranceagreements
  end
  resources :endorsements do
    resource :policy
  end
  resources :policycoverages do
    resource :policy
    resources :insuredobjects
  end
  resources :insuredobjects do
    resource :policy
    resources :coverages
  end
  resources :beneficiarys do
    resource :policy
    resource :customer
  end
  resources :billingaccounts do
    resource :customer
    resources :policies
    resources :invoices
    resources :payments
  end
  resources :invoices do
    resource :billingaccount
    resource :policy
    resources :payments
  end
  resources :payments do
    resource :invoice
    resource :billingaccount
    resource :policy
  end
  resources :claims do
    resource :policy
    resource :customer
    resource :adjuster
    resource :incident
    resources :exposures
    resources :reserves
    resources :claimpayments
    resources :serviceproviders
    resources :subrogations
  end
  resources :incidents do
    resource :claim
    resources :insuredobjects
  end
  resources :exposures do
    resource :claim
    resource :policycoverage
    resource :insuredobject
    resources :reserves
    resources :payments
  end
  resources :adjusters do
    resources :claims
    resources :serviceproviders
  end
  resources :claimreserves do
    resource :claim
    resource :exposure
  end
  resources :claimpayments do
    resource :claim
    resource :exposure
    resource :beneficiary
    resource :serviceprovider
    resource :customer
  end
  resources :serviceproviders do
    resources :claims
  end
  resources :reinsuranceagreements do
    resource :insurer
    resources :policies
  end
  resources :subrogationrecoverys do
    resource :claim
    resource :exposure
    resource :counterparty
  end
  resources :thirdpartys do
    resources :subrogations
  end
  resources :documents do
    resource :policy
    resource :claim
    resource :application
    resource :customer
  end
end
