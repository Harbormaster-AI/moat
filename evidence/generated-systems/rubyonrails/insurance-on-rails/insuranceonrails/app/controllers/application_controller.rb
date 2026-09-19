class ApplicationController < ActionController::Base
    def health
        render plain: "Ruby on Rails application insuranceonrails is running", status: :ok
    end
end