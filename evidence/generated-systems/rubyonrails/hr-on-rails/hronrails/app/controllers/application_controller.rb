class ApplicationController < ActionController::Base
    def health
        render plain: "Ruby on Rails application hronrails is running", status: :ok
    end
end