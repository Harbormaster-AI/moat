class Candidate < ApplicationRecord
  enum Source: [:Referral, :Agency, :JobBoard, :CareerSite, :Campus, :Social, :Internal]


  composed_of :personName,
    class_name: "PersonName",
    mapping: [
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      %w[personName_preferredName preferredName]
    ]

  composed_of :email,
    class_name: "Email",
    mapping: [
      %w[email_value value]
    ]

  composed_of :phoneNumber,
    class_name: "PhoneNumber",
    mapping: [
      ${$mapping}, 
      ${$mapping}, 
      %w[phoneNumber_extension extension]
    ]

  has_many :Applications, class_name: 'JobApplication'
  has_many :Interviews, class_name: 'Interview'
  has_many :Offers, class_name: 'Offer'
  has_many :Documents, class_name: 'Document'

end
