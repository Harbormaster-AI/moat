class ApplicationController < ActionController::Base
    def health
        render plain: "Ruby on Rails application healthcareonrails is running", status: :ok
    end
end