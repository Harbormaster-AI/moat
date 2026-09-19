class ApplicationController < ActionController::Base
    def health
        render plain: "Ruby on Rails application ecommerceonrails is running", status: :ok
    end
end