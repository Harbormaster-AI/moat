class ApplicationController < ActionController::Base
    def health
        render plain: "Ruby on Rails application fintechonrails is running", status: :ok
    end
end